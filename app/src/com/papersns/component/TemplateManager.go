package component

import (
	"encoding/json"
	"log"
	"labix.org/v2/mgo"
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
		log.Println("QueryDataForListTemplate,collection:" + collection + ",query is:" + string(queryByte))
		result := querySupport.Index(collection, queryMap, pageNo, pageSize)
		items := result["items"].([]interface{})
		result["items"] = expressionParser.ParseAfterQueryData(listTemplate.AfterQueryData, items)
		return result
	}
	mapReduce := mgo.MapReduce{
		Map: mapStr,
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
			if columnItem.XMLName.Local != "virtual-column" {
				loopItem[columnItem.Name] = record[columnItem.Name]
			} else {
				if loopItem[columnItem.Name] == nil {
					virtualColumn := map[string]interface{}{}
					buttons := []interface{}{}
					virtualColumn["buttons"] = buttons
					loopItem[columnItem.Name] = virtualColumn
				}

				for _, buttonItem := range columnItem.Buttons.ButtonLi {
					button := map[string]interface{}{}
					button["isShow"] = expressionParser.Parse(recordJson, buttonItem.Expression)
					virtualColumn := loopItem[columnItem.Name].(map[string]interface{})
					buttons := virtualColumn["buttons"].([]interface{})
					buttons = append(buttons, button)
					virtualColumn["buttons"] = buttons // append will generate a new reference, so must reset value
				}
			}
		}

		columnModelItems = append(columnModelItems, loopItem)
	}

	return columnModelItems
}

func (o TemplateManager) GetToolbarForListTemplate(listTemplate *ListTemplate) []interface{} {
	toolbar := []interface{}{}

	expressionParser := ExpressionParser{}
	for _, buttonItem := range listTemplate.Toolbar.ButtonLi {
		button := map[string]interface{}{}
		//		button["isShow"] = expressionParser.Parse(buttonItem.Expression)
		expression := buttonItem.Expression
		expression = ""
		button["isShow"] = expressionParser.Parse("", expression)
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
