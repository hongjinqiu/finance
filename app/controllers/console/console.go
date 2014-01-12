package sysuser

import "github.com/robfig/revel"
import (
	. "com/papersns/component"
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
		componentItems = append(componentItems, map[string]interface{}{
			"id":     item.ListTemplate.Id,
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

	// 1.toolbar bo
	templateManager.ApplyDictionaryForQueryParameter(&listTemplate)
	templateManager.ApplyTreeForQueryParameter(&listTemplate)
	toolbarBo := templateManager.GetToolbarForListTemplate(listTemplate)
	paramMap := map[string]string{}
	for k, v := range c.Params.Form {
		value := strings.Join(v, ",")
		if value != "" {
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
	dataBo := templateManager.GetBoForListTemplate(&listTemplate, paramMap, pageNo, pageSize)

	//	columns := templateManager.GetColumns(&listTemplate)

	//	columnsByte, err := json.Marshal(columns)
	//	if err != nil {
	//		fmt.Printf("error: %v", err)
	//		return c.Render(err)
	//	}

	dataBoByte, err := json.Marshal(&dataBo)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	listTemplateByte, err := json.Marshal(&listTemplate)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		callback := c.Params.Get("callback")
		if callback == "" {
			c.Response.ContentType = "application/json; charset=utf-8"
			return c.RenderJson(&dataBo)
		}
		c.Response.ContentType = "text/javascript; charset=utf-8"
		return c.RenderText(callback + "(" + string(dataBoByte) + ");")
	} else {
		showParameterLi := templateManager.GetShowParameterLiForListTemplate(&listTemplate)
		hiddenParameterLi := templateManager.GetHiddenParameterLiForListTemplate(&listTemplate)
		result := map[string]interface{}{
			"pageSize":          pageSize,
			"listTemplate":      listTemplate,
			"toolbarBo":         toolbarBo,
			"showParameterLi":   showParameterLi,
			"hiddenParameterLi": hiddenParameterLi,
			//		"dataBo":       dataBo,
			//		"columns":       columns,
			"dataBoJson":       template.JS(string(dataBoByte)),
			"listTemplateJson": template.JS(string(listTemplateByte)),
			//		"columnsJson":   string(columnsByte),
		}
		return c.Render(result)
	}
}

func (c Console) SelectorSchema() revel.Result {
	schemaName := c.Params.Get("@name")

	templateManager := TemplateManager{}
	listTemplate := templateManager.GetSelectorTemplate(schemaName)

	// 1.toolbar bo
	templateManager.ApplyDictionaryForQueryParameter(&listTemplate)
	templateManager.ApplyTreeForQueryParameter(&listTemplate)
	toolbarBo := templateManager.GetToolbarForListTemplate(listTemplate)
	paramMap := map[string]string{}
	for k, v := range c.Params.Form {
		value := strings.Join(v, ",")
		if value != "" {
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
	dataBo := templateManager.GetBoForListTemplate(&listTemplate, paramMap, pageNo, pageSize)

	//	columns := templateManager.GetColumns(&listTemplate)

	//	columnsByte, err := json.Marshal(columns)
	//	if err != nil {
	//		fmt.Printf("error: %v", err)
	//		return c.Render(err)
	//	}

	dataBoByte, err := json.Marshal(&dataBo)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	listTemplateByte, err := json.Marshal(&listTemplate)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		callback := c.Params.Get("callback")
		if callback == "" {
			c.Response.ContentType = "application/json; charset=utf-8"
			return c.RenderJson(&dataBo)
		}
		c.Response.ContentType = "text/javascript; charset=utf-8"
		return c.RenderText(callback + "(" + string(dataBoByte) + ");")
	} else {
		showParameterLi := templateManager.GetShowParameterLiForListTemplate(&listTemplate)
		hiddenParameterLi := templateManager.GetHiddenParameterLiForListTemplate(&listTemplate)
		result := map[string]interface{}{
			"pageSize":          pageSize,
			"listTemplate":      listTemplate,
			"toolbarBo":         toolbarBo,
			"showParameterLi":   showParameterLi,
			"hiddenParameterLi": hiddenParameterLi,
			//		"dataBo":       dataBo,
			//		"columns":       columns,
			"dataBoJson":       template.JS(string(dataBoByte)),
			"listTemplateJson": template.JS(string(listTemplateByte)),
			//		"columnsJson":   string(columnsByte),
		}
		return c.Render(result)
	}
}

// TODO,by test
func (c Console) FormSchema() revel.Result {
	schemaName := c.Params.Get("@name")

	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(schemaName)

	result := map[string]interface{}{
		"formTemplate": formTemplate,
	}
	if formTemplate.DataSourceModelId != "" {
		// 光有formTemplate不行,还要有model的内容,才可以渲染数据
		modelTemplateFactory := ModelTemplateFactory{}
		dataSource := modelTemplateFactory.GetDataSource(formTemplate.DataSourceModelId)
		dataSourceByte, err := json.Marshal(&dataSource)
		if err != nil {
			panic(err)
		}
		result["dataSource"] = dataSource
		result["dataSourceJson"] = template.JS(string(dataSourceByte))
	}
	//toolbarBo
	toolbarBo := map[string]interface{}{}
	for _, item := range formTemplate.FormElemLi {
		if item.XMLName.Local == "toolbar" {
			toolbarBo[item.Toolbar.Name] = templateManager.GetToolbarBo(item.Toolbar)
		}
	}
	result["toolbarBo"] = toolbarBo
	dataBo := map[string]interface{}{}
	result["dataBo"] = dataBo

	formTemplateJsonDataArray, err := json.Marshal(&formTemplate)
	if err != nil {
		panic(err)
	}

	dataBoByte, err := json.Marshal(&dataBo)
	if err != nil {
		panic(err)
	}

	result["formTemplateJsonData"] = template.JS(string(formTemplateJsonDataArray))
	result["dataBoJson"] = template.JS(string(dataBoByte))

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

	return c.Render(string(fileContent), result)
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