package component

import (
	"com/papersns/dictionary"
	"com/papersns/mongo"
	"com/papersns/tree"
	"encoding/json"
	"fmt"
	"labix.org/v2/mgo"
	"log"
)

type TemplateManager struct{}

func (o TemplateManager) QueryDataForListTemplate(listTemplate *ListTemplate, paramMap map[string]string, pageNo int, pageSize int) map[string]interface{} {
	expressionParser := ExpressionParser{}
	paramMap = expressionParser.ParseBeforeBuildQuery(listTemplate.BeforeBuildQuery, paramMap)

	queryMap := map[string]interface{}{}
	queryLi := []map[string]interface{}{}

	collection := listTemplate.DataProvider.Collection
	mapStr := listTemplate.DataProvider.Map
	reduce := listTemplate.DataProvider.Reduce
	fixBsonQuery := listTemplate.DataProvider.FixBsonQuery

	fixBsonQueryMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(fixBsonQuery), &fixBsonQueryMap)
	if err != nil {
		panic(err)
	}

	queryLi = append(queryLi, fixBsonQueryMap)

	queryParameters := listTemplate.QueryParameterGroup.QueryParameterLi
	queryParameterBuilder := QueryParameterBuilder{}
	for _, queryParameter := range queryParameters {
		if queryParameter.Editor != "" && queryParameter.Restriction != "" {
			name := queryParameter.Name
			if paramMap[name] != "" {
				queryParameterMap := queryParameterBuilder.buildQuery(queryParameter, paramMap[name])
				queryLi = append(queryLi, queryParameterMap)
			}
		}
	}

	queryLi = expressionParser.ParseAfterBuildQuery(listTemplate.AfterBuildQuery, queryLi)

	querySupport := QuerySupport{}
	queryMap["$and"] = queryLi
	if len(queryLi) == 1 {
		queryMap = queryLi[0]
	}

	queryByte, err := json.MarshalIndent(queryMap, "", "\t")
	if err != nil {
		panic(err)
	}
	if mapStr == "" {
		orderBy := listTemplate.ColumnModel.BsonOrderBy
		log.Println("QueryDataForListTemplate,collection:" + collection + ",query is:" + string(queryByte) + ",orderBy is:" + orderBy)
		result := querySupport.Index(collection, queryMap, pageNo, pageSize, orderBy)
		items := result["items"].([]interface{})
		result["items"] = expressionParser.ParseAfterQueryData(listTemplate.AfterQueryData, items)
		return result
	}
	mapReduce := mgo.MapReduce{
		Map:    mapStr,
		Reduce: reduce,
	}

	mapReduceByte, err := json.MarshalIndent(mapReduce, "", "\t")
	if err != nil {
		panic(err)
	}

	log.Println("QueryDataForListTemplate,collection:" + collection + ",query is:" + string(queryByte) + ",mapReduce:" + string(mapReduceByte))
	mapReduceLi := querySupport.MapReduceAll(collection, queryMap, mapReduce)
	items := []interface{}{}
	for _, item := range mapReduceLi {
		item["id"] = item["_id"]
		items = append(items, item)
	}
	items = expressionParser.ParseAfterQueryData(listTemplate.AfterQueryData, items)
	return map[string]interface{}{
		"totalResults": len(mapReduceLi),
		"items":        items,
	}
}

func (o TemplateManager) GetColumnModelDataForListTemplate(listTemplate *ListTemplate, items []interface{}) []interface{} {
	columnModelItems := []interface{}{}
	expressionParser := ExpressionParser{}
	for _, item := range items {
		record := item.(map[string]interface{})
		recordJsonByte, err := json.Marshal(record)
		if err != nil {
			panic(err)
		}
		recordJson := string(recordJsonByte)

		loopItem := map[string]interface{}{}
		loopItem[listTemplate.ColumnModel.CheckboxColumn.Name] = expressionParser.Parse(recordJson, listTemplate.ColumnModel.CheckboxColumn.Expression)
		loopItem[listTemplate.ColumnModel.IdColumn.Name] = record[listTemplate.ColumnModel.IdColumn.Name]
		loopItem["id"] = record[listTemplate.ColumnModel.IdColumn.Name]
		loopItem["_id"] = record[listTemplate.ColumnModel.IdColumn.Name]
		for _, columnItem := range listTemplate.ColumnModel.ColumnLi {
			o.GetColumnModelDataForColumnItem(columnItem, record, &loopItem)
		}

		columnModelItems = append(columnModelItems, loopItem)
	}

	return columnModelItems
}

func (o TemplateManager) GetColumnModelDataForColumnItem(columnItem Column, record map[string]interface{}, loopItem *map[string]interface{}) {
	if columnItem.XMLName.Local != "virtual-column" {
		if columnItem.ColumnModel.ColumnLi != nil {
			for _, columnItemItem := range columnItem.ColumnModel.ColumnLi {
				o.GetColumnModelDataForColumnItem(columnItemItem, record, loopItem)
			}
		} else {
			(*loopItem)[columnItem.Name] = record[columnItem.Name]
			o.ApplyDictionaryColumnData(loopItem, columnItem)
			o.ApplyScriptColumnData(loopItem, record, columnItem)
		}
	} else {
		if (*loopItem)[columnItem.Name] == nil {
			virtualColumn := map[string]interface{}{}
			buttons := []interface{}{}
			virtualColumn["buttons"] = buttons
			(*loopItem)[columnItem.Name] = virtualColumn
		}

		recordJsonByte, err := json.Marshal(record)
		if err != nil {
			panic(err)
		}
		recordJson := string(recordJsonByte)
		
		expressionParser := ExpressionParser{}
		for _, buttonItem := range columnItem.Buttons.ButtonLi {
			button := map[string]interface{}{}
			button["isShow"] = expressionParser.Parse(recordJson, buttonItem.Expression)
			virtualColumn := (*loopItem)[columnItem.Name].(map[string]interface{})
			buttons := virtualColumn["buttons"].([]interface{})
			buttons = append(buttons, button)
			virtualColumn["buttons"] = buttons // append will generate a new reference, so must reset value
		}
	}
}

func (o TemplateManager) ApplyDictionaryColumnData(loopItem *map[string]interface{}, columnItem Column) {
	dictionaryManager := dictionary.GetInstance()
	if columnItem.XMLName.Local == "dictionary-column" {
		dictionaryItem := dictionaryManager.GetDictionary(columnItem.Dictionary)
		items := dictionaryItem["items"]
		if items != nil {
			itemsLi := items.([]map[string]interface{})
			columnValue := fmt.Sprint((*loopItem)[columnItem.Name])
			for _, codeNameItem := range itemsLi {
				code := fmt.Sprint(codeNameItem["code"])
				if code == columnValue {
					(*loopItem)[columnItem.Name+"_DICTIONARY_NAME"] = codeNameItem["name"]
					break
				}
			}
		}
	}
}

func (o TemplateManager) ApplyScriptColumnData(loopItem *map[string]interface{}, record map[string]interface{}, columnItem Column) {
	if columnItem.XMLName.Local == "script-column" {
		data, err := json.Marshal(record)
		if err != nil {
			panic(err)
		}

		expressionParser := ExpressionParser{}
		scriptValue := expressionParser.ParseString(string(data), columnItem.Script)
		(*loopItem)[columnItem.Name] = scriptValue
	}
}

func (o TemplateManager) GetToolbarForListTemplate(listTemplate *ListTemplate) []interface{} {
	toolbar := []interface{}{}

	expressionParser := ExpressionParser{}
	for _, buttonItem := range listTemplate.Toolbar.ButtonLi {
		button := map[string]interface{}{}
		button["isShow"] = expressionParser.Parse("", buttonItem.Expression)
		toolbar = append(toolbar, button)
	}

	return toolbar
}

/**
 * 获取模版业务对象
 */
func (o TemplateManager) GetBoForListTemplate(listTemplate *ListTemplate, paramMap map[string]string, pageNo int, pageSize int) map[string]interface{} {
	queryResult := o.QueryDataForListTemplate(listTemplate, paramMap, pageNo, pageSize)
	items := queryResult["items"].([]interface{})
	bo := o.GetColumnModelDataForListTemplate(listTemplate, items)
	return map[string]interface{}{
		"totalResults": queryResult["totalResults"],
		"items":        bo,
	}
}

func (o TemplateManager) GetColumns(listTemplate *ListTemplate) []string {
	fields := []string{}
	//	loopItem["isShowCheckbox"] = expressionParser.Parse(recordJson, listTemplate.ColumnModel.CheckboxColumn.Expression)
	//		loopItem["id"] = record[listTemplate.ColumnModel.IdColumn.Name]
	//		for _, columnItem := range listTemplate.ColumnModel.ColumnLi {
	fields = append(fields, listTemplate.ColumnModel.IdColumn.Name)
	for _, columnItem := range listTemplate.ColumnModel.ColumnLi {
		//		fields = append(fields, columnItem.Name)
		fields = append(fields, columnItem.Name)
	}
	return fields
}

func (o TemplateManager) GetShowParameterLiForListTemplate(listTemplate *ListTemplate) []QueryParameter {
	queryParameterLi := []QueryParameter{}
	for _, item := range listTemplate.QueryParameterGroup.QueryParameterLi {
		if item.Editor != "hidden" {
			queryParameterLi = append(queryParameterLi, item)
		}
	}
	return queryParameterLi
}

func (o TemplateManager) GetHiddenParameterLiForListTemplate(listTemplate *ListTemplate) []QueryParameter {
	queryParameterLi := []QueryParameter{}
	for _, item := range listTemplate.QueryParameterGroup.QueryParameterLi {
		if item.Editor == "hidden" {
			queryParameterLi = append(queryParameterLi, item)
		}
	}
	return queryParameterLi
}

func (o TemplateManager) ApplyDictionaryForQueryParameter(listTemplate *ListTemplate) {
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()

	dictionaryManager := dictionary.GetInstance()
	for i, _ := range listTemplate.QueryParameterGroup.QueryParameterLi {
		item := &(listTemplate.QueryParameterGroup.QueryParameterLi[i])
		for _, parameterAttribute := range item.ParameterAttributeLi {
			if parameterAttribute.Name == "dictionary" {
				item.Dictionary = dictionaryManager.GetDictionaryBySession(db, parameterAttribute.Value)
				break
			}
		}
	}
}

func (o TemplateManager) ApplyTreeForQueryParameter(listTemplate *ListTemplate) {
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()

	treeManager := tree.GetInstance()
	for i, _ := range listTemplate.QueryParameterGroup.QueryParameterLi {
		item := &(listTemplate.QueryParameterGroup.QueryParameterLi[i])
		for _, parameterAttribute := range item.ParameterAttributeLi {
			if parameterAttribute.Name == "tree" {
				item.Tree = treeManager.GetTreeBySession(db, parameterAttribute.Value)
				break
			}
		}
	}
}
