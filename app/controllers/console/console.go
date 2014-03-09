package sysuser

import "github.com/robfig/revel"
import (
	. "com/papersns/common"
	. "com/papersns/component"
	"com/papersns/global"
	. "com/papersns/model"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	//	"net/http"
	"html/template"
	"os"
	"strconv"
	"strings"
)

func init() {

}

type Console struct {
	*revel.Controller
}

func (c Console) Summary() revel.Result {
	println("session is:", c.Session["userId"])

	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate("Console")

	//	if true {
	//		xmlDataArray, err := xml.Marshal(&formTemplate)
	//		if err != nil {
	//			panic(err)
	//		}
	//		return c.RenderXml(&formTemplate)
	//	}

	formTemplateJsonDataArray, err := json.Marshal(&formTemplate)
	if err != nil {
		panic(err)
	}

	toolbarBo := map[string]interface{}{}

	dataBo := map[string]interface{}{}
	{
		listTemplateInfoLi := templateManager.GetListTemplateInfoLi()
		dataBo["Component"] = getSummaryListTemplateInfoLi(listTemplateInfoLi)
	}
	{
		selectorTemplateInfoLi := templateManager.GetSelectorTemplateInfoLi()
		dataBo["Selector"] = getSummarySelectorTemplateInfoLi(selectorTemplateInfoLi)
	}
	{
		formTemplateInfoLi := templateManager.GetFormTemplateInfoLi()
		dataBo["Form"] = getSummaryFormTemplateInfoLi(formTemplateInfoLi)
	}
	{
		modelTemplateFactory := ModelTemplateFactory{}
		dataSourceInfoLi := modelTemplateFactory.GetDataSourceInfoLi()
		dataBo["DataSource"] = getSummaryDataSourceInfoLi(dataSourceInfoLi)
	}
	for _, item := range formTemplate.FormElemLi {
		if item.XMLName.Local == "column-model" {
			if dataBo[item.ColumnModel.Name] == nil {
				dataBo[item.ColumnModel.Name] = []interface{}{}
			}
			items := dataBo[item.ColumnModel.Name].([]interface{})
			items = templateManager.GetColumnModelDataForColumnModel(item.ColumnModel, items)
			dataBo[item.ColumnModel.Name] = items
		} else if item.XMLName.Local == "toolbar" {
			toolbarBo[item.Toolbar.Name] = templateManager.GetToolbarBo(item.Toolbar)
		}
	}

	dataBoByte, err := json.Marshal(dataBo)
	if err != nil {
		panic(err)
	}

	//	c.Response.Status = http.StatusOK
	//	c.Response.ContentType = "text/plain; charset=utf-8"
	result := map[string]interface{}{
		"formTemplate":         formTemplate,
		"toolbarBo":            toolbarBo,
		"dataBo":               dataBo,
		"formTemplateJsonData": template.JS(string(formTemplateJsonDataArray)),
		"dataBoJson":           template.JS(string(dataBoByte)),
	}
	// formTemplate.ViewTemplate.View
	//	RenderText(text string, objs ...interface{}) Result

	viewPath := revel.Config.StringDefault("REVEL_VIEW_PATH", "")
	file, err := os.Open(viewPath + "/" + formTemplate.ViewTemplate.View)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	//	c.Response.Out
	//	return c.RenderTemplate(string(fileContent))
	funcMap := map[string]interface{}{
		"eq": func(a, b interface{}) bool {
			return a == b
		},
	}
	c.Response.ContentType = "text/html; charset=utf-8"
	tmpl, err := template.New("summary").Funcs(funcMap).Parse(string(fileContent))
	if err != nil {
		panic(err)
	}
	tmplResult := map[string]interface{}{
		"result": result,
	}
	//tmpl.Execute(c.Response.Out, result)
	tmpl.Execute(c.Response.Out, tmplResult)
	return nil
	//	return c.Render(string(fileContent), result)
}

func getSummaryListTemplateInfoLi(listTemplateInfoLi []ListTemplateInfo) []interface{} {
	componentItems := []interface{}{}
	for _, item := range listTemplateInfoLi {
		module := "组件模型"
		if item.ListTemplate.DataSourceModelId != "" && item.ListTemplate.Adapter.Name != "" {
			module = "数据源模型适配"
		}
		componentItems = append(componentItems, map[string]interface{}{
			"id":     item.ListTemplate.Id,
			"name":   item.ListTemplate.Description,
			"module": module,
			"path":   item.Path,
		})
	}
	return componentItems
}

func getSummarySelectorTemplateInfoLi(selectorTemplateInfoLi []SelectorTemplateInfo) []interface{} {
	componentItems := []interface{}{}
	for _, item := range selectorTemplateInfoLi {
		module := "组件模型选择器"
		if item.ListTemplate.DataSourceModelId != "" && item.ListTemplate.Adapter.Name != "" {
			module = "数据源模型选择器适配"
		}
		id := item.ListTemplate.SelectorId
		if id == "" {
			id = item.ListTemplate.Id
		}
		componentItems = append(componentItems, map[string]interface{}{
			"id":     id,
			"name":   item.ListTemplate.Description,
			"module": module,
			"path":   item.Path,
		})
	}
	return componentItems
}

func getSummaryFormTemplateInfoLi(formTemplateInfoLi []FormTemplateInfo) []interface{} {
	formItems := []interface{}{}
	for _, item := range formTemplateInfoLi {
		module := "form模型"
		if item.FormTemplate.DataSourceModelId != "" && item.FormTemplate.Adapter.Name != "" {
			module = "数据源模型适配"
		}
		formItems = append(formItems, map[string]interface{}{
			"id":     item.FormTemplate.Id,
			"name":   item.FormTemplate.Description,
			"module": module,
			"path":   item.Path,
		})
	}
	return formItems
}

func getSummaryDataSourceInfoLi(dataSourceInfoLi []DataSourceInfo) []interface{} {
	dataSourceItems := []interface{}{}
	for _, item := range dataSourceInfoLi {
		dataSourceItems = append(dataSourceItems, map[string]interface{}{
			"id":     item.DataSource.Id,
			"name":   item.DataSource.DisplayName,
			"module": "数据源模型",
			"path":   item.Path,
		})
	}
	return dataSourceItems
}

func (c Console) ListSchema() revel.Result {
	schemaName := c.Params.Get("@name")

	templateManager := TemplateManager{}
	listTemplate := templateManager.GetListTemplate(schemaName)

	result := c.listSelectorCommon(&listTemplate, true)
	
	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		callback := c.Params.Get("callback")
		if callback == "" {
			dataBo := result["dataBo"]
			c.Response.ContentType = "application/json; charset=utf-8"
			return c.RenderJson(&dataBo)
		}
		dataBoText := result["dataBoText"].(string)
		c.Response.ContentType = "text/javascript; charset=utf-8"
		return c.RenderText(callback + "(" + dataBoText + ");")
	} else {
		return c.Render(result)
	}
}

func (c Console) SelectorSchema() revel.Result {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)

	schemaName := c.Params.Get("@name")

	templateManager := TemplateManager{}
	listTemplate := templateManager.GetSelectorTemplate(schemaName)
	c.setSelectionMode(&listTemplate)
	c.setDisplayField(&listTemplate)
	isGetBo := false
	if c.Params.Get("format") != "" {
		isGetBo = true
	}
	result := c.listSelectorCommon(&listTemplate, isGetBo)

	selectionBo := map[string]interface{}{
		"url": templateManager.GetViewUrl(listTemplate),
	}
	ids := c.Params.Get("@id")
	if ids != "" {
		relationLi := []map[string]interface{}{}
		strIdLi := strings.Split(ids, ",")
		selectorId := listTemplate.SelectorId
		if selectorId == "" {
			selectorId = listTemplate.Id
		}
		for _, item := range strIdLi {
			if item != "" {
				id, err := strconv.Atoi(item)
				if err != nil {
					panic(err)
				}
				relationLi = append(relationLi, map[string]interface{}{
					"relationId": id,
					"selectorId": selectorId,
				})
			}
		}
		templateManager := TemplateManager{}
		relationBo := templateManager.GetRelationBo(sessionId, relationLi)
		if relationBo[selectorId] != nil {
			selectionBo = relationBo[selectorId].(map[string]interface{})
		}
	}
	selectionBoByte, err := json.Marshal(&selectionBo)
	if err != nil {
		panic(err)
	}

	commonUtil := CommonUtil{}
	selectionBoJson := string(selectionBoByte)
	selectionBoJson = commonUtil.FilterJsonEmptyAttr(selectionBoJson)
	result["selectionBoJson"] = template.JS(selectionBoJson)

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		callback := c.Params.Get("callback")
		if callback == "" {
			dataBo := result["dataBo"]
			c.Response.ContentType = "application/json; charset=utf-8"
			return c.RenderJson(&dataBo)
		}
		dataBoText := result["dataBoText"].(string)
		c.Response.ContentType = "text/javascript; charset=utf-8"
		return c.RenderText(callback + "(" + dataBoText + ");")
	} else {
		return c.Render(result)
	}
}

func (c Console) setSelectionMode(listTemplate *ListTemplate) {
	multi := c.Params.Get("@multi")
	if multi != "" {
		if multi == "true" {
			listTemplate.ColumnModel.SelectionMode = "checkbox"
		} else {
			listTemplate.ColumnModel.SelectionMode = "radio"
		}
	}
}

func (c Console) setDisplayField(listTemplate *ListTemplate) {
	displayField := c.Params.Get("@displayField")
	if displayField != "" {
		if strings.Contains(displayField, "{") {
			listTemplate.ColumnModel.SelectionTemplate = displayField
		} else {
			strFieldLi := strings.Split(displayField, ",")
			fieldLi := []string{}
			for _, item := range strFieldLi {
				fieldLi = append(fieldLi, "{"+item+"}")
			}
			listTemplate.ColumnModel.SelectionTemplate = strings.Join(fieldLi, ",")
		}
	}
}

func (c Console) listSelectorCommon(listTemplate *ListTemplate, isGetBo bool) map[string]interface{} {
	// 1.toolbar bo
	templateManager := TemplateManager{}
	templateManager.ApplyDictionaryForQueryParameter(listTemplate)
	templateManager.ApplyTreeForQueryParameter(listTemplate)
	toolbarBo := templateManager.GetToolbarForListTemplate(*listTemplate)
	paramMap := map[string]string{}
	//	c.Request.URL
	for k, v := range c.Params.Form {
		value := strings.Join(v, ",")
		if value != "" { // && !strings.Contains(k, "@")
			paramMap[k] = value
		}
	}
	pageNo := 1
	pageSize := 10
	if listTemplate.DataProvider.Size != "" {
		pageSizeInt, err := strconv.Atoi(listTemplate.DataProvider.Size)
		if err != nil {
			panic(err)
		}
		pageSize = pageSizeInt
	}
	if c.Params.Get("pageNo") != "" {
		pageNoInt, _ := strconv.ParseInt(c.Params.Get("pageNo"), 10, 0)
		if pageNoInt > 1 {
			pageNo = int(pageNoInt)
		}
	}
	if c.Params.Get("pageSize") != "" {
		pageSizeInt, _ := strconv.ParseInt(c.Params.Get("pageSize"), 10, 0)
		if pageSizeInt >= 10 {
			pageSize = int(pageSizeInt)
		}
	}
	dataBo := map[string]interface{}{
		"totalResults": 0,
		"items":        []interface{}{},
	}
	//if c.Params.Get("@entrance") != "true" {
	if isGetBo {
		dataBo = templateManager.GetBoForListTemplate(listTemplate, paramMap, pageNo, pageSize)
	}

	dataBoByte, err := json.Marshal(&dataBo)
	if err != nil {
		panic(err)
	}

	listTemplateByte, err := json.Marshal(listTemplate)
	if err != nil {
		panic(err)
	}
	
	queryParameterRenderLi := c.getQueryParameterRenderLi(*listTemplate)

	//showParameterLi := templateManager.GetShowParameterLiForListTemplate(listTemplate)
	showParameterLi := []QueryParameter{}
	hiddenParameterLi := templateManager.GetHiddenParameterLiForListTemplate(listTemplate)
	result := map[string]interface{}{
		"pageSize":          pageSize,
		"listTemplate":      listTemplate,
		"toolbarBo":         toolbarBo,
		"showParameterLi":   showParameterLi,
		"hiddenParameterLi": hiddenParameterLi,
		"queryParameterRenderLi": queryParameterRenderLi,
		"dataBo":            dataBo,
		//		"columns":       columns,
		"dataBoText":       string(dataBoByte),
		"dataBoJson":       template.JS(string(dataBoByte)),
		"listTemplateJson": template.JS(string(listTemplateByte)),
		//		"columnsJson":   string(columnsByte),
	}
	return result
}

// TODO,by test
func (c Console) FormSchema() revel.Result {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)

	schemaName := c.Params.Get("@name")
	strId := c.Params.Get("id")
	formStatus := c.Params.Get("formStatus")

	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(schemaName)

	result := map[string]interface{}{
		"formTemplate": formTemplate,
		"id":           strId,
		"formStatus":   formStatus,
	}
	if formTemplate.DataSourceModelId != "" {
		// 光有formTemplate不行,还要有model的内容,才可以渲染数据
		modelTemplateFactory := ModelTemplateFactory{}
		dataSource := modelTemplateFactory.GetDataSource(formTemplate.DataSourceModelId)
		modelTemplateFactory.ClearReverseRelation(&dataSource)
		dataSourceByte, err := json.Marshal(&dataSource)
		if err != nil {
			panic(err)
		}
		result["dataSource"] = dataSource
		commonUtil := CommonUtil{}
		dataSourceJson := string(dataSourceByte)
		dataSourceJson = commonUtil.FilterJsonEmptyAttr(dataSourceJson)
		result["dataSourceJson"] = template.JS(dataSourceJson)
	}
	//toolbarBo
	toolbarBo := map[string]interface{}{}
	for i, item := range formTemplate.FormElemLi {
		if item.XMLName.Local == "toolbar" {
			toolbarBo[item.Toolbar.Name] = templateManager.GetToolbarBo(item.Toolbar)
		}
		// 加入主数据集tag,页面渲染用
		if item.XMLName.Local == "column-model" && item.ColumnModel.DataSetId == "A" {
			formTemplate.FormElemLi[i].RenderTag = item.ColumnModel.DataSetId + "_" + fmt.Sprint(i)
		}
	}
	result["toolbarBo"] = toolbarBo
	dataBo := map[string]interface{}{}
	if strId != "" && formTemplate.DataSourceModelId != "" {
		dataSourceModelId := formTemplate.DataSourceModelId
		id, err := strconv.Atoi(strId)
		if err != nil {
			panic(err)
		}
		querySupport := QuerySupport{}
		queryMap := map[string]interface{}{
			"_id": id,
		}
		modelTemplateFactory := ModelTemplateFactory{}
		dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
		collectionName := modelTemplateFactory.GetCollectionName(dataSource)
		bo, found := querySupport.FindByMap(collectionName, queryMap)
		if !found {
			panic("FormSchema, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
		}
		dataBo = bo
	}
	result["dataBo"] = dataBo

	relationBo := map[string]interface{}{}
	if formTemplate.DataSourceModelId != "" {
		if strId != "" && formTemplate.DataSourceModelId != "" {
			modelTemplateFactory := ModelTemplateFactory{}
			dataSource := modelTemplateFactory.GetDataSource(formTemplate.DataSourceModelId)
			relationLi := modelTemplateFactory.GetRelationLi(sessionId, dataSource, dataBo)
			relationBo = templateManager.GetRelationBo(sessionId, relationLi)
		}
	}
	result["relationBo"] = relationBo
	relationBoByte, err := json.Marshal(&relationBo)
	if err != nil {
		panic(err)
	}

	// 主数据集的后台渲染
	result["masterRenderLi"] = c.getMasterRenderLi(formTemplate)
	//	{
	//		dataBoByte, err := json.Marshal(result["masterRenderLi"])
	//		if err != nil {
	//			panic(err)
	//		}
	//		commonUtil := CommonUtil{}
	//		masterRenderLiJson := string(dataBoByte)
	//		masterRenderLiJson = commonUtil.FilterJsonEmptyAttr(masterRenderLiJson)
	//		result["masterRenderLiJson"] = template.JS(masterRenderLiJson)
	//	}

	formTemplateJsonDataArray, err := json.Marshal(&formTemplate)
	if err != nil {
		panic(err)
	}

	dataBoByte, err := json.Marshal(&dataBo)
	if err != nil {
		panic(err)
	}

	layerBo := templateManager.GetLayerForFormTemplate(sessionId, formTemplate)
	layerBoByte, err := json.Marshal(&layerBo)
	if err != nil {
		panic(err)
	}

	commonUtil := CommonUtil{}
	formTemplateJsonData := string(formTemplateJsonDataArray)
	formTemplateJsonData = commonUtil.FilterJsonEmptyAttr(formTemplateJsonData)
	result["formTemplateJsonData"] = template.JS(formTemplateJsonData)
	dataBoJson := string(dataBoByte)
	dataBoJson = commonUtil.FilterJsonEmptyAttr(dataBoJson)
	result["dataBoJson"] = template.JS(dataBoJson)
	layerBoJson := string(layerBoByte)
	layerBoJson = commonUtil.FilterJsonEmptyAttr(layerBoJson)
	result["layerBoJson"] = template.JS(layerBoJson)
	relationBoJson := string(relationBoByte)
	relationBoJson = commonUtil.FilterJsonEmptyAttr(relationBoJson)
	result["relationBoJson"] = template.JS(relationBoJson)

	viewPath := revel.Config.StringDefault("REVEL_VIEW_PATH", "")
	file, err := os.Open(viewPath + "/" + formTemplate.ViewTemplate.View)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	funcMap := map[string]interface{}{
		"eq": func(a, b interface{}) bool {
			return a == b
		},
	}
	c.Response.ContentType = "text/html; charset=utf-8"
	tmpl, err := template.New("formSchema").Funcs(funcMap).Parse(string(fileContent))
	if err != nil {
		panic(err)
	}
	tmplResult := map[string]interface{}{
		"result": result,
	}
	tmpl.Execute(c.Response.Out, tmplResult)
	return nil
}

func (c Console) getMasterRenderLi(formTemplate FormTemplate) map[string]interface{} {
	if formTemplate.DataSourceModelId == "" {
		return nil
	}
	result := map[string]interface{}{}

	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(formTemplate.DataSourceModelId)

	modelIterator := ModelIterator{}
	var message interface{} = ""
	for i, item := range formTemplate.FormElemLi {
		if item.XMLName.Local == "column-model" && item.ColumnModel.DataSetId == "A" {
			lineColSpan, err := strconv.Atoi(item.ColumnModel.ColSpan)
			if err != nil {
				lineColSpan = 1
			}
			container := [][]map[string]interface{}{}
			containerItem := []map[string]interface{}{}
			lineColSpanSum := 0
			for _, column := range item.ColumnModel.ColumnLi {
				if column.XMLName.Local == "html" {
					columnColSpan, err := strconv.Atoi(column.ColSpan)
					if err != nil {
						columnColSpan = 1
					}
					containerItem = append(containerItem, map[string]interface{}{
						"isHtml": "true",
						"html":   column.Html,
					})
					lineColSpanSum += columnColSpan
					if lineColSpanSum >= lineColSpan {
						container = append(container, containerItem)
						containerItem = []map[string]interface{}{}
						lineColSpanSum = lineColSpanSum - lineColSpan
					}
				} else {
					modelIterator.IterateAllField(&dataSource, &message, func(fieldGroup *FieldGroup, result *interface{}) {
						if fieldGroup.IsMasterField() && fieldGroup.Id == column.Name {
							if column.Hideable != "true" {
								columnColSpan, err := strconv.Atoi(column.ColSpan)
								if err != nil {
									columnColSpan = 1
								}
								containerItem = append(containerItem, map[string]interface{}{
									"isHtml":      "false",
									"required":    fmt.Sprint(fieldGroup.AllowEmpty == "false"),
									"label":       column.Text,
									"name":        column.Name,
									"columnWidth": column.ColumnWidth,
									"columnSpan":  columnColSpan - 1,
									"labelWidth":  column.LabelWidth,
								})
								lineColSpanSum += columnColSpan
								if lineColSpanSum >= lineColSpan {
									container = append(container, containerItem)
									containerItem = []map[string]interface{}{}
									lineColSpanSum = lineColSpanSum - lineColSpan
								}
							}
						}
					})
				}
			}
			if 0 < lineColSpanSum && lineColSpanSum < lineColSpan {
				container = append(container, containerItem)
			}
			result[item.DataSetId+"_"+fmt.Sprint(i)] = container
		}
	}

	return result
}

func (c Console) getQueryParameterRenderLi(listTemplate ListTemplate) [][]map[string]interface{} {
	lineColSpan := 6
	container := [][]map[string]interface{}{}
	containerItem := []map[string]interface{}{}
	lineColSpanSum := 0
	listTemplateIterator := ListTemplateIterator{}
	var result interface{} = ""
	listTemplateIterator.IterateTemplateQueryParameter(listTemplate, &result, func(queryParameter QueryParameter, result *interface{}){
		if queryParameter.Editor != "hidden" {
			columnColSpan := 2
			containerItem = append(containerItem, map[string]interface{}{
				"label":       queryParameter.Text,
				"name":        queryParameter.Name,
			})
			lineColSpanSum += columnColSpan
			if lineColSpanSum >= lineColSpan {
				container = append(container, containerItem)
				containerItem = []map[string]interface{}{}
				lineColSpanSum = lineColSpanSum - lineColSpan
			}
		}
	})
	if 0 < lineColSpanSum && lineColSpanSum < lineColSpan {
		container = append(container, containerItem)
	}
	return container
}

func (c Console) Relation() revel.Result {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)

	selectorId := c.Params.Get("selectorId")
	id := c.Params.Get("id")

	templateManager := TemplateManager{}
	relationLi := []map[string]interface{}{
		map[string]interface{}{
			"selectorId": selectorId,
			"relationId": id,
		},
	}
	relationBo := templateManager.GetRelationBo(sessionId, relationLi)
	var result interface{} = nil
	var url interface{} = nil
	if relationBo[selectorId] != nil {
		selRelationBo := relationBo[selectorId].(map[string]interface{})
		if selRelationBo[id] != nil {
			result = selRelationBo[id]
			url = selRelationBo["url"]
		}
	}
	c.Response.ContentType = "application/json; charset=utf-8"
	return c.RenderJson(map[string]interface{}{
		"result": result,
		"url":    url,
	})
}

func (c Console) Refretor() revel.Result {
	refretorType := c.Params.Get("type")
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate("Console")

	if refretorType == "Component" {
		listTemplateInfoLi := templateManager.RefretorListTemplateInfo()
		items := getSummaryListTemplateInfoLi(listTemplateInfoLi)
		for _, item := range formTemplate.FormElemLi {
			if item.XMLName.Local == "column-model" && item.ColumnModel.Name == "Component" {
				items = templateManager.GetColumnModelDataForColumnModel(item.ColumnModel, items)
				break
			}
		}

		dataBo := map[string]interface{}{
			"items": items,
		}

		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataBo)
	}
	if refretorType == "Selector" {
		selectorTemplateInfoLi := templateManager.RefretorSelectorTemplateInfo()
		items := getSummarySelectorTemplateInfoLi(selectorTemplateInfoLi)
		for _, item := range formTemplate.FormElemLi {
			if item.XMLName.Local == "column-model" && item.ColumnModel.Name == "Selector" {
				items = templateManager.GetColumnModelDataForColumnModel(item.ColumnModel, items)
				break
			}
		}

		dataBo := map[string]interface{}{
			"items": items,
		}
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataBo)
	}
	if refretorType == "Form" {
		formTemplateInfoLi := templateManager.RefretorFormTemplateInfo()
		items := getSummaryFormTemplateInfoLi(formTemplateInfoLi)
		for _, item := range formTemplate.FormElemLi {
			if item.XMLName.Local == "column-model" && item.ColumnModel.Name == "Form" {
				items = templateManager.GetColumnModelDataForColumnModel(item.ColumnModel, items)
				break
			}
		}

		dataBo := map[string]interface{}{
			"items": items,
		}
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataBo)
	}
	if refretorType == "DataSource" {
		modelTemplateFactory := ModelTemplateFactory{}
		dataSourceTemplateInfoLi := modelTemplateFactory.RefretorDataSourceInfo()
		items := getSummaryDataSourceInfoLi(dataSourceTemplateInfoLi)
		for _, item := range formTemplate.FormElemLi {
			if item.XMLName.Local == "column-model" && item.ColumnModel.Name == "DataSource" {
				items = templateManager.GetColumnModelDataForColumnModel(item.ColumnModel, items)
				break
			}
		}

		dataBo := map[string]interface{}{
			"items": items,
		}
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataBo)
	}
	c.Response.ContentType = "application/json; charset=utf-8"
	return c.RenderJson(map[string]interface{}{
		"message": "可能传入了错误的refretorType:" + refretorType,
	})
}

func (c Console) Xml() revel.Result {
	refretorType := c.Params.Get("type")
	id := c.Params.Get("@name")
	templateManager := TemplateManager{}

	if refretorType == "Component" {
		listTemplate := templateManager.GetListTemplate(id)
		return c.RenderXml(&listTemplate)
	}
	if refretorType == "Selector" {
		selectorTemplate := templateManager.GetSelectorTemplate(id)
		return c.RenderXml(&selectorTemplate)
	}
	if refretorType == "Form" {
		formTemplate := templateManager.GetFormTemplate(id)
		return c.RenderXml(&formTemplate)
	}
	if refretorType == "DataSource" {
		modelTemplateFactory := ModelTemplateFactory{}
		dataSourceTemplate := modelTemplateFactory.GetDataSource(id)
		return c.RenderXml(&dataSourceTemplate)
	}
	c.Response.ContentType = "application/json; charset=utf-8"
	return c.RenderJson(map[string]interface{}{
		"message": "可能传入了错误的refretorType:" + refretorType,
	})
}

func (c Console) RawXml() revel.Result {
	refretorType := c.Params.Get("type")
	id := c.Params.Get("@name")
	templateManager := TemplateManager{}

	if refretorType == "Component" {
		listTemplateInfo := templateManager.GetListTemplateInfo(id)
		listTemplate := ListTemplate{}
		file, err := os.Open(listTemplateInfo.Path)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = xml.Unmarshal(data, &listTemplate)
		if err != nil {
			panic(err)
		}

		return c.RenderXml(&listTemplate)
	}
	if refretorType == "Selector" {
		selectorTemplateInfo := templateManager.GetSelectorTemplateInfo(id)
		selectorTemplate := ListTemplate{}
		file, err := os.Open(selectorTemplateInfo.Path)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = xml.Unmarshal(data, &selectorTemplate)
		if err != nil {
			panic(err)
		}

		return c.RenderXml(&selectorTemplate)
	}
	if refretorType == "Form" {
		formTemplateInfo := templateManager.GetFormTemplateInfo(id)
		formTemplate := FormTemplate{}
		file, err := os.Open(formTemplateInfo.Path)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = xml.Unmarshal(data, &formTemplate)
		if err != nil {
			panic(err)
		}

		return c.RenderXml(&formTemplate)
	}
	c.Response.ContentType = "application/json; charset=utf-8"
	return c.RenderJson(map[string]interface{}{
		"message": "可能传入了错误的refretorType:" + refretorType,
	})
}
