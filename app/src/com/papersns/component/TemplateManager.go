package component

import "github.com/robfig/revel"

import (
	"com/papersns/dictionary"
	. "com/papersns/interceptor"
	. "com/papersns/script"
	"com/papersns/mongo"
	"com/papersns/tree"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"labix.org/v2/mgo"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

var templaterwlock sync.RWMutex = sync.RWMutex{}
var gListTemplateDict map[string]ListTemplateInfo = map[string]ListTemplateInfo{}
var gFormTemplateDict map[string]FormTemplateInfo = map[string]FormTemplateInfo{}
var gSelectorTemplateDict map[string]SelectorTemplateInfo = map[string]SelectorTemplateInfo{}

type ListTemplateInfo struct {
	Path         string
	ListTemplate ListTemplate
}

type FormTemplateInfo struct {
	Path         string
	FormTemplate FormTemplate
}

type SelectorTemplateInfo struct {
	Path         string
	ListTemplate ListTemplate
}

type TemplateManager struct{}

// TODO, byTest
func (o TemplateManager) GetListTemplateInfo() []ListTemplateInfo {
	listTemplateInfo := []ListTemplateInfo{}
	for _, item := range gListTemplateDict {
		listTemplateInfo = append(listTemplateInfo, item)
	}
	return listTemplateInfo
}

// TODO, byTest
func (o TemplateManager) RefretorListTemplateInfo() []ListTemplateInfo {
	o.clearListTemplate()
	o.loadListTemplate()
	listTemplateInfo := []ListTemplateInfo{}
	for _, item := range gListTemplateDict {
		listTemplateInfo = append(listTemplateInfo, item)
	}
	return listTemplateInfo
}

// TODO, byTest
func (o TemplateManager) GetListTemplate(id string) ListTemplate {
	if revel.Config.StringDefault("mode.dev", "true") == "true" {
		listTemplateInfo, found := o.findListTemplate(id)
		if found {
			listTemplateInfo, err := o.loadSingleListTemplateWithLock(listTemplateInfo.Path)
			if err != nil {
				panic(err)
			}
			if listTemplateInfo.ListTemplate.Id == id {
				return listTemplateInfo.ListTemplate
			}
		}
		o.clearListTemplate()
		o.loadListTemplate()
		listTemplateInfo, found = o.findListTemplate(id)
		if found {
			return listTemplateInfo.ListTemplate
		}
		panic(id + " not exists in ListTemplate list")
	}

	if len(gListTemplateDict) == 0 {
		o.loadListTemplate()
	}
	listTemplateInfo, found := o.findListTemplate(id)
	if found {
		return listTemplateInfo.ListTemplate
	}
	panic(id + " not exists in ListTemplate list")
}

// TODO bytest,
func (o TemplateManager) findListTemplate(id string) (ListTemplateInfo, bool) {
	templaterwlock.RLock()
	defer templaterwlock.RUnlock()

	for _, item := range gListTemplateDict {
		if item.ListTemplate.Id == id {
			return item, true
		}
	}
	return ListTemplateInfo{}, false
}

// TODO, byTest
func (o TemplateManager) clearListTemplate() {
	templaterwlock.Lock()
	defer templaterwlock.Unlock()

	gListTemplateDict = map[string]ListTemplateInfo{}
}

// TODO, byTest
func (o TemplateManager) loadListTemplate() {
	templaterwlock.Lock()
	defer templaterwlock.Unlock()

	path := revel.Config.StringDefault("LIST_TEMPLATE_PATH", "")
	if path != "" {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Index(path, "list_") > -1 && strings.Index(path, ".xml") > -1 && !info.IsDir() {
				_, err = o.loadSingleListTemplate(path)
				if err != nil {
					return err
				}
			}

			return nil
		})
	}
}

// TODO, byTest
func (o TemplateManager) loadSingleListTemplateWithLock(path string) (ListTemplateInfo, error) {
	templaterwlock.Lock()
	defer templaterwlock.Unlock()

	return o.loadSingleListTemplate(path)
}

// TODO, byTest
func (o TemplateManager) loadSingleListTemplate(path string) (ListTemplateInfo, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return ListTemplateInfo{}, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return ListTemplateInfo{}, err
	}

	listTemplate := ListTemplate{}
	err = xml.Unmarshal(data, &listTemplate)
	if err != nil {
		return ListTemplateInfo{}, err
	}

	// TODO byTest
	if listTemplate.Adapter.Name != "" {
		classMethod := listTemplate.Adapter.Name + ".ApplyAdapter"
		commonMethod := CommonMethod{}
		paramLi := []*interface{}{}
		var param interface{} = listTemplate
		paramLi = append(paramLi, &param)
		commonMethod.Parse(classMethod, &paramLi)
	}

	listTemplateInfo := ListTemplateInfo{
		Path:         path,
		ListTemplate: listTemplate,
	}
	gListTemplateDict[listTemplate.Id] = listTemplateInfo
	return listTemplateInfo, nil
}

// TODO, byTest
func (o TemplateManager) GetSelectorTemplateInfo() []SelectorTemplateInfo {
	selectorTemplateInfo := []SelectorTemplateInfo{}
	for _, item := range gSelectorTemplateDict {
		selectorTemplateInfo = append(selectorTemplateInfo, item)
	}
	return selectorTemplateInfo
}

// TODO, byTest
func (o TemplateManager) RefretorSelectorTemplateInfo() []SelectorTemplateInfo {
	o.clearSelectorTemplate()
	o.loadSelectorTemplate()
	selectorTemplateInfo := []SelectorTemplateInfo{}
	for _, item := range gSelectorTemplateDict {
		selectorTemplateInfo = append(selectorTemplateInfo, item)
	}
	return selectorTemplateInfo
}

// TODO, byTest
func (o TemplateManager) GetSelectorTemplate(id string) ListTemplate {
	if revel.Config.StringDefault("mode.dev", "true") == "true" {
		selectorTemplateInfo, found := o.findSelectorTemplate(id)
		if found {
			selectorTemplateInfo, err := o.loadSingleSelectorTemplateWithLock(selectorTemplateInfo.Path)
			if err != nil {
				panic(err)
			}
			if strings.Index(selectorTemplateInfo.Path, "list_") == -1 {
				if selectorTemplateInfo.ListTemplate.Id == id {
					return selectorTemplateInfo.ListTemplate
				}
			}
		}
		o.clearSelectorTemplate()
		o.loadSelectorTemplate()
		selectorTemplateInfo, found = o.findSelectorTemplate(id)
		if found {
			return selectorTemplateInfo.ListTemplate
		}
		panic(id + " not exists in ListTemplate list")
	}

	if len(gSelectorTemplateDict) == 0 {
		o.loadSelectorTemplate()
	}
	selectorTemplateInfo, found := o.findSelectorTemplate(id)
	if found {
		return selectorTemplateInfo.ListTemplate
	}
	panic(id + " not exists in ListTemplate list")
}

// TODO bytest,
func (o TemplateManager) findSelectorTemplate(id string) (SelectorTemplateInfo, bool) {
	templaterwlock.RLock()
	defer templaterwlock.RUnlock()

	for _, item := range gSelectorTemplateDict {
		if item.ListTemplate.Id == id {
			return item, true
		}
	}
	return SelectorTemplateInfo{}, false
}

// TODO, byTest
func (o TemplateManager) clearSelectorTemplate() {
	templaterwlock.Lock()
	defer templaterwlock.Unlock()

	gSelectorTemplateDict = map[string]SelectorTemplateInfo{}
}

// TODO, byTest
func (o TemplateManager) loadSelectorTemplate() {
	templaterwlock.Lock()
	defer templaterwlock.Unlock()

	path := revel.Config.StringDefault("SELECTOR_TEMPLATE_PATH", "")
	if path != "" {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Index(path, "list_") > -1 && strings.Index(path, ".xml") > -1 && !info.IsDir() {
				_, err = o.loadSingleSelectorTemplate(path)
				if err != nil {
					return err
				}
			}

			return nil
		})
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Index(path, "selector_") > -1 && strings.Index(path, ".xml") > -1 && !info.IsDir() {
				_, err = o.loadSingleSelectorTemplate(path)
				if err != nil {
					return err
				}
			}

			return nil
		})
	}
}

// TODO, byTest
func (o TemplateManager) loadSingleSelectorTemplateWithLock(path string) (SelectorTemplateInfo, error) {
	templaterwlock.Lock()
	defer templaterwlock.Unlock()

	return o.loadSingleSelectorTemplate(path)
}

// TODO, byTest
func (o TemplateManager) loadSingleSelectorTemplate(path string) (SelectorTemplateInfo, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return SelectorTemplateInfo{}, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return SelectorTemplateInfo{}, err
	}

	listTemplate := ListTemplate{}
	err = xml.Unmarshal(data, &listTemplate)
	if err != nil {
		return SelectorTemplateInfo{}, err
	}

	isAdd := true
	if strings.Index(path, "list_") > -1 {
		if listTemplate.SelectorId == "" {
			isAdd = false
		}
	}
	if isAdd {
		if listTemplate.Adapter.Name != "" {
			classMethod := listTemplate.Adapter.Name + ".ApplyAdapter"
			commonMethod := CommonMethod{}
			paramLi := []*interface{}{}
			var param interface{} = listTemplate
			paramLi = append(paramLi, &param)
			commonMethod.Parse(classMethod, &paramLi)
		}

		selectorTemplateInfo := SelectorTemplateInfo{
			Path:         path,
			ListTemplate: listTemplate,
		}
		if strings.Index(path, "list_") > -1 {
			gSelectorTemplateDict[listTemplate.SelectorId] = selectorTemplateInfo
		} else {
			gSelectorTemplateDict[listTemplate.Id] = selectorTemplateInfo
		}
	}
	return SelectorTemplateInfo{}, nil
}


// TODO, byTest
func (o TemplateManager) GetFormTemplateInfo() []FormTemplateInfo {
	formTemplateInfo := []FormTemplateInfo{}
	for _, item := range gFormTemplateDict {
		formTemplateInfo = append(formTemplateInfo, item)
	}
	return formTemplateInfo
}

// TODO, byTest
func (o TemplateManager) RefretorFormTemplateInfo() []FormTemplateInfo {
	o.clearFormTemplate()
	o.loadFormTemplate()
	formTemplateInfo := []FormTemplateInfo{}
	for _, item := range gFormTemplateDict {
		formTemplateInfo = append(formTemplateInfo, item)
	}
	return formTemplateInfo
}

// TODO, byTest
func (o TemplateManager) GetFormTemplate(id string) FormTemplate {
	if revel.Config.StringDefault("mode.dev", "true") == "true" {
		formTemplateInfo, found := o.findFormTemplate(id)
		if found {
			formTemplateInfo, err := o.loadSingleFormTemplateWithLock(formTemplateInfo.Path)
			if err != nil {
				panic(err)
			}
			if formTemplateInfo.FormTemplate.Id == id {
				return formTemplateInfo.FormTemplate
			}
		}
		o.clearFormTemplate()
		o.loadFormTemplate()
		formTemplateInfo, found = o.findFormTemplate(id)
		if found {
			return formTemplateInfo.FormTemplate
		}
		panic(id + " not exists in FormTemplate list")
	}

	if len(gFormTemplateDict) == 0 {
		o.loadFormTemplate()
	}
	formTemplateInfo, found := o.findFormTemplate(id)
	if found {
		return formTemplateInfo.FormTemplate
	}
	panic(id + " not exists in FormTemplate list")
}

// TODO bytest,
func (o TemplateManager) findFormTemplate(id string) (FormTemplateInfo, bool) {
	templaterwlock.RLock()
	defer templaterwlock.RUnlock()

	for _, item := range gFormTemplateDict {
		if item.FormTemplate.Id == id {
			return item, true
		}
	}
	return FormTemplateInfo{}, false
}

// TODO, byTest
func (o TemplateManager) clearFormTemplate() {
	templaterwlock.Lock()
	defer templaterwlock.Unlock()

	gFormTemplateDict = map[string]FormTemplateInfo{}
}

// TODO, byTest
func (o TemplateManager) loadFormTemplate() {
	templaterwlock.Lock()
	defer templaterwlock.Unlock()

	path := revel.Config.StringDefault("FORM_TEMPLATE_PATH", "")
	if path != "" {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Index(path, "form_") > -1 && strings.Index(path, ".xml") > -1 && !info.IsDir() {
				_, err = o.loadSingleFormTemplate(path)
				if err != nil {
					return err
				}
			}

			return nil
		})
	}
}

// TODO, byTest
func (o TemplateManager) loadSingleFormTemplateWithLock(path string) (FormTemplateInfo, error) {
	templaterwlock.Lock()
	defer templaterwlock.Unlock()

	return o.loadSingleFormTemplate(path)
}

// TODO, byTest
func (o TemplateManager) loadSingleFormTemplate(path string) (FormTemplateInfo, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return FormTemplateInfo{}, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return FormTemplateInfo{}, err
	}

	formTemplate := FormTemplate{}
	err = xml.Unmarshal(data, &formTemplate)
	if err != nil {
		return FormTemplateInfo{}, err
	}

	// TODO byTest
	if formTemplate.Adapter.Name != "" {
		classMethod := formTemplate.Adapter.Name + ".ApplyAdapter"
		commonMethod := CommonMethod{}
		paramLi := []*interface{}{}
		var param interface{} = formTemplate
		paramLi = append(paramLi, &param)
		commonMethod.Parse(classMethod, &paramLi)
	}

	for i, _ := range formTemplate.FormElemLi {
		formElem := &formTemplate.FormElemLi[i]
		if formElem.XMLName.Local == "html" {
			formElemXmlData, err := xml.Marshal(formElem)
			if err != nil {
				panic(err)
			}
			err = xml.Unmarshal(formElemXmlData, &(formElem.Html))
			if err != nil {
				panic(err)
			}
		} else if formElem.XMLName.Local == "toolbar" {
			formElemXmlData, err := xml.Marshal(formElem)
			if err != nil {
				panic(err)
			}
			err = xml.Unmarshal(formElemXmlData, &(formElem.Toolbar))
			if err != nil {
				panic(err)
			}
		} else if formElem.XMLName.Local == "column-model" {
			formElemXmlData, err := xml.Marshal(formElem)
			if err != nil {
				panic(err)
			}
			err = xml.Unmarshal(formElemXmlData, &(formElem.ColumnModel))
			if err != nil {
				panic(err)
			}
		}
	}

	formTemplateInfo := FormTemplateInfo{
		Path:         path,
		FormTemplate: formTemplate,
	}
	gFormTemplateDict[formTemplate.Id] = formTemplateInfo
	return formTemplateInfo, nil
}

func (o TemplateManager) QueryDataForListTemplate(listTemplate *ListTemplate, paramMap map[string]string, pageNo int, pageSize int) map[string]interface{} {
	interceptorManager := InterceptorManager{}
	interceptorManager.ParseBeforeBuildQuery(listTemplate.BeforeBuildQuery, &paramMap)

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
				if listTemplate.Adapter.Name != "" {
					classMethod := listTemplate.Adapter.Name + ".ApplyQueryParameter"
					commonMethod := CommonMethod{}
					paramLi := []*interface{}{}
					var param interface{} = listTemplate
					paramLi = append(paramLi, &param)
					param = queryParameter
					paramLi = append(paramLi, &param)
					commonMethod.Parse(classMethod, &paramLi)
				}
				queryParameterMap := queryParameterBuilder.buildQuery(queryParameter, paramMap[name])
				queryLi = append(queryLi, queryParameterMap)
			}
		}
	}

	interceptorManager.ParseAfterBuildQuery(listTemplate.AfterBuildQuery, &queryLi)

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
		interceptorManager.ParseAfterQueryData(listTemplate.AfterQueryData, &items)
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
	interceptorManager.ParseAfterQueryData(listTemplate.AfterQueryData, &items)
	return map[string]interface{}{
		"totalResults": len(mapReduceLi),
		"items":        items,
	}
}

func (o TemplateManager) GetColumnModelDataForListTemplate(listTemplate ListTemplate, items []interface{}) []interface{} {
	o.applyAdapterColumnName(&listTemplate)
	return o.GetColumnModelDataForColumnModel(listTemplate.ColumnModel, items)
}

func (o TemplateManager) applyAdapterColumnName(listTemplate *ListTemplate) {
	if listTemplate.Adapter.Name != "" {
		//ApplyColumnName(listTemplate *ListTemplate, column *Column) {
		classMethod := listTemplate.Adapter.Name + ".ApplyColumnName"
		commonMethod := CommonMethod{}
		paramLi := []*interface{}{}
		var param interface{} = listTemplate
		paramLi = append(paramLi, &param)
		param = listTemplate.ColumnModel.IdColumn
		paramLi = append(paramLi, &param)
		commonMethod.Parse(classMethod, &paramLi)
		for i, _ := range listTemplate.ColumnModel.ColumnLi {
			o.recursionApplyAdapterColumnName(*listTemplate, &listTemplate.ColumnModel.ColumnLi[i])
		}
	}
}

// TODO, bytest
func (o TemplateManager) recursionApplyAdapterColumnName(listTemplate ListTemplate, columnItem *Column) {
	if columnItem.XMLName.Local != "virtual-column" {
		if columnItem.ColumnModel.ColumnLi != nil {
			for i, _ := range columnItem.ColumnModel.ColumnLi {
				o.recursionApplyAdapterColumnName(listTemplate, &columnItem.ColumnModel.ColumnLi[i])
			}
		} else {
			commonMethod := CommonMethod{}
			classMethod := listTemplate.Adapter.Name + ".ApplyColumnName"
			paramLi := []*interface{}{}
			var param interface{} = listTemplate
			paramLi = append(paramLi, &param)
			param = *columnItem
			paramLi = append(paramLi, &param)
			commonMethod.Parse(classMethod, &paramLi)
		}
	}
}

func (o TemplateManager) GetColumnModelDataForColumnModel(columnModel ColumnModel, items []interface{}) []interface{} {
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
		loopItem[columnModel.CheckboxColumn.Name] = expressionParser.Parse(recordJson, columnModel.CheckboxColumn.Expression)
		loopItem[columnModel.IdColumn.Name] = o.getValueBySpot(record, columnModel.IdColumn.Name)
		loopItem["id"] = o.getValueBySpot(record, columnModel.IdColumn.Name)
		loopItem["_id"] = o.getValueBySpot(record, columnModel.IdColumn.Name)
		for _, columnItem := range columnModel.ColumnLi {
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
			(*loopItem)[columnItem.Name] = o.getValueBySpot(record, columnItem.Name)
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

// TODO, byTest
func (o TemplateManager) getValueBySpot(record map[string]interface{}, name string) interface{} {
	current := record
	nameLi := strings.Split(name, ".")
	for i, _ := range nameLi {
		if i < len(nameLi)-1 {
			if current[name] == nil {
				return nil
			}
			if reflect.ValueOf(current[name]).Kind() == reflect.Map {
				current = current[name].(map[string]interface{})
			} else {
				return nil
			}
		} else {
			return current[name]
		}
	}
	return nil
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

func (o TemplateManager) GetToolbarForListTemplate(listTemplate ListTemplate) []interface{} {
	return o.GetToolbarBo(listTemplate.Toolbar)
}

func (o TemplateManager) GetToolbarBo(toolbar Toolbar) []interface{} {
	result := []interface{}{}

	expressionParser := ExpressionParser{}
	for _, buttonItem := range toolbar.ButtonLi {
		button := map[string]interface{}{}
		button["isShow"] = expressionParser.Parse("", buttonItem.Expression)
		result = append(result, button)
	}

	return result
}

/**
 * 获取模版业务对象
 */
func (o TemplateManager) GetBoForListTemplate(listTemplate *ListTemplate, paramMap map[string]string, pageNo int, pageSize int) map[string]interface{} {
	queryResult := o.QueryDataForListTemplate(listTemplate, paramMap, pageNo, pageSize)
	items := queryResult["items"].([]interface{})
	bo := o.GetColumnModelDataForListTemplate(*listTemplate, items)
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
