package sysuser

import "github.com/robfig/revel"
import (
	. "com/papersns/component"
	. "com/papersns/model"
	"encoding/json"
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

	formTemplateJsonDataArray, err := json.Marshal(&formTemplate)
	if err != nil {
		panic(err)
	}

	toolbarBo := map[string]interface{}{}

	dataBo := map[string]interface{}{}
	//	GetFormTemplateInfo
	{
		listTemplateInfoLi := templateManager.GetListTemplateInfo()

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
		dataBo["Component"] = componentItems
	}
	{
		formTemplateInfoLi := templateManager.GetFormTemplateInfo()

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
		dataBo["Form"] = formItems
	}
	{
		modelTemplateFactory := ModelTemplateFactory{}
		dataSourceInfoLi := modelTemplateFactory.GetDataSourceInfo()

		dataSourceItems := []interface{}{}
		for _, item := range dataSourceInfoLi {
			dataSourceItems = append(dataSourceItems, map[string]interface{}{
				"id":     item.DataSource.Id,
				"name":   item.DataSource.DisplayName,
				"module": "数据源模型",
				"path":   item.Path,
			})
		}
		dataBo["DataSource"] = dataSourceItems
	}
	for _, item := range formTemplate.FormElemLi {
		if item.XMLName.Local == "column-model" {
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

	return c.Render(string(fileContent), result)
	//	return c.RenderText(string(jsonDataArray))
	//	return c.RenderText(string(xmlDataArray))
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

	if refretorType == "Component" {
		listTemplateInfoLi := templateManager.RefretorListTemplateInfo()
		dataBo := map[string]interface{}{
			"items": listTemplateInfoLi,
		}
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataBo)
	}
	if refretorType == "Selector" {
		listTemplateInfoLi := templateManager.RefretorSelectorTemplateInfo()
		dataBo := map[string]interface{}{
			"items": listTemplateInfoLi,
		}
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataBo)
	}
	if refretorType == "Form" {
		formTemplateInfoLi := templateManager.RefretorFormTemplateInfo()
		dataBo := map[string]interface{}{
			"items": formTemplateInfoLi,
		}
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataBo)
	}
	if refretorType == "DataSource" {
		modelTemplateFactory := ModelTemplateFactory{}
		dataSourceTemplateInfoLi := modelTemplateFactory.RefretorDataSourceInfo()
		dataBo := map[string]interface{}{
			"items": dataSourceTemplateInfoLi,
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
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&listTemplate)
	}
	if refretorType == "Selector" {
		listTemplate := templateManager.GetSelectorTemplate(id)
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&listTemplate)
	}
	if refretorType == "Form" {
		formTemplate := templateManager.GetFormTemplate(id)
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&formTemplate)
	}
	if refretorType == "DataSource" {
		modelTemplateFactory := ModelTemplateFactory{}
		dataSourceTemplate := modelTemplateFactory.GetDataSource(id)
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataSourceTemplate)
	}
	c.Response.ContentType = "application/json; charset=utf-8"
	return c.RenderJson(map[string]interface{}{
		"message": "可能传入了错误的refretorType:" + refretorType,
	})
}
