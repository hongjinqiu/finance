package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/component"
	//	"github.com/sbinet/go-python"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"labix.org/v2/mgo"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func init() {
	/*
		err := python.Initialize()
		if err != nil {
			panic(err)
		}

		sys_path := python.PySys_GetObject("path")
		if sys_path == nil {
			panic("get sys.path return nil")
		}

		path := python.PyString_FromString("/home/hongjinqiu/goworkspace/src/finance")
		if path == nil {
			panic("get path return nil")
		}

		err = python.PyList_Append(sys_path, path)
		if err != nil {
			panic(err)
		}
	*/
	revel.TemplateFuncs["gt"] = func(a int, b int) bool {
		return a > b
	}
	revel.TemplateFuncs["gte"] = func(a int, b int) bool {
		return a >= b
	}
	revel.TemplateFuncs["lt"] = func(a int, b int) bool {
		return a < b
	}
	revel.TemplateFuncs["lte"] = func(a int, b int) bool {
		return a <= b
	}
	revel.TemplateFuncs["residue"] = func(a int, b int, c int) bool {
		return a%b == c
	}
	revel.TemplateFuncs["last"] = func(a int, b int) bool {
		return a-1 == b
	}
}

type Component struct {
	*revel.Controller
}

func (c Component) Schema() revel.Result {
	file, err := os.Open("/home/hongjinqiu/goworkspace/src/finance/app/src/com/papersns/component/schema/查询示例.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	listTemplate := ListTemplate{}
	err = xml.Unmarshal(data, &listTemplate)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	// 1.query data,
	// from data-provider
	// from query-parameters

	xmlDataArray, err := xml.MarshalIndent(listTemplate, "", "\t")
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}
	c.Response.Status = http.StatusOK
	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText(string(xmlDataArray))
}

func (c Component) ListTemplate() revel.Result {
	file, err := os.Open("/home/hongjinqiu/goworkspace/src/finance/app/src/com/papersns/component/schema/SysUser.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	listTemplate := ListTemplate{}
	err = xml.Unmarshal(data, &listTemplate)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	// 1.toolbar bo
	templateManager := TemplateManager{}
	toolbarBo := templateManager.GetToolbarForListTemplate(&listTemplate)
	paramMap := map[string]string{}
	for k, v := range c.Params.Form {
		value := strings.Join(v, ",")
		if value != "" {
			paramMap[k] = value
		}
	}
	pageNo := 1
	pageSize := 10
	if c.Params.Get("pageNo") != "" {
		pageNoInt, _ := strconv.ParseInt(c.Params.Get("pageNo"), 10, 0)
		if pageNoInt > 1 {
			pageNo = int(pageNoInt)
		}
	}
	if c.Params.Get("pageSize") != "" {
		pageSizeInt, _ := strconv.ParseInt(c.Params.Get("pageSize"), 10, 0)
		if pageSizeInt > 10 {
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
		// 1.query data,
		// from data-provider
		// from query-parameters
		result := map[string]interface{}{
			"listTemplate": listTemplate,
			"toolbarBo":    toolbarBo,
			//		"dataBo":       dataBo,
			//		"columns":       columns,
			"dataBoJson":       template.JS(string(dataBoByte)),
			"listTemplateJson": template.JS(string(listTemplateByte)),
			//		"columnsJson":   string(columnsByte),
		}
		return c.Render(result)
	}
}

func (c Component) MapReduce() revel.Result {
	qb := QuerySupport{}

	collection := "SysUser"
	query := map[string]interface{}{
		"_id": map[string]interface{}{
			"$lt": 20,
		},
	}
	mapReduce := mgo.MapReduce{
		Map: `
function() {
    emit(this.nick, {
        count: 1,
        nick: this.nick
    })
}
`,
		Reduce: `
function(key, values) {
    var count = 0;
    values.forEach(function(item){
        count += item.count;
    })
    return {
        count: count,
        nick: values[0].nick
    }
}
`,
	}
	results := qb.MapReduce(collection, query, mapReduce, 1, 2)
	c.Response.ContentType = "application/json; charset=utf-8"
	return c.RenderJson(&results)
}

func (c Component) ScriptTest() revel.Result {
	o := ExpressionParser{}
	//	classMethod := "SysUser.beforeBuildQuery"
	//	paramMap := map[string]string{
	//		"age": "20",
	//	}
	//	result := o.ParseBeforeBuildQuery(classMethod, paramMap)
	classMethod := "SysUser.afterBuildQuery"
	queryLi := []map[string]interface{}{
		map[string]interface{}{"name": "test"},
		map[string]interface{}{"nick": "测试"},
	}
	result := o.ParseAfterBuildQuery(classMethod, queryLi)
	c.Response.ContentType = "application/json; charset=utf-8"
	return c.RenderJson(result)
}

func (c Component) MongoTest() revel.Result {
	querySupport := QuerySupport{}
	m, isFind := querySupport.Find("SysUser", `{"_id": 15}`)
	if !isFind {
		panic("not found")
	}

	m["bool"] = true
	m["nil"] = nil
	c.Response.Status = http.StatusOK
	c.Response.ContentType = "text/plain; charset=utf-8"
	data, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		panic(err)
	}
	return c.RenderText(string(data))
}

func (c Component) IndexTest() revel.Result {
	querySupport := QuerySupport{}
	collection := "ScheduleSetting"
	query := map[string]interface{}{}
	pageNo := 1
	pageSize := 10
	result := querySupport.Index(collection, query, pageNo, pageSize)

	c.Response.Status = http.StatusOK
	c.Response.ContentType = "text/plain; charset=utf-8"
	data, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		panic(err)
	}
	return c.RenderText(string(data))
}

func (c Component) SchemaTest() revel.Result {
	file, err := os.Open(`/home/hongjinqiu/goworkspace/src/finance/app/src/com/papersns/component/schema/查询示例.xml`)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	listTemplate := ListTemplate{}
	err = xml.Unmarshal(data, &listTemplate)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	// 1.query data,
	// from data-provider
	// from query-parameters

	xmlDataArray, err := xml.MarshalIndent(listTemplate, "", "\t")
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}
	println(xmlDataArray)

	jsonArray, err := json.MarshalIndent(listTemplate, "", "\t")
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	c.Response.Status = http.StatusOK
	c.Response.ContentType = "text/plain; charset=utf-8"
	//	return c.RenderText(string(xmlDataArray))
	return c.RenderText(string(jsonArray))
}

func (c Component) SchemaQueryParameterTest() revel.Result {
	file, err := os.Open(`/home/hongjinqiu/goworkspace/src/finance/app/src/com/papersns/component/schema/SysUser.xml`)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	listTemplate := ListTemplate{}
	err = xml.Unmarshal(data, &listTemplate)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	// 1.query data,
	// from data-provider
	// from query-parameters
	templateManager := TemplateManager{}
	paramMap := map[string]string{}
	{
		//		paramMap["nick"] = "abc"
		//		paramMap["dept_id"] = "2"
		//		paramMap["type"] = "0,1,2.5,3.5,abc"
		//		paramMap["createTimeBegin"] = "2013-05-07"
		//		paramMap["createTimeEnd"] = "2014-06-03"
	}
	pageNo := 1
	pageSize := 10
	queryResult := templateManager.QueryDataForListTemplate(&listTemplate, paramMap, pageNo, pageSize)
	items := queryResult["items"].([]interface{})
	if len(items) > 1 {
		queryResult["items"] = items[:1]
	}

	jsonByte, err := json.MarshalIndent(queryResult, "", "\t")
	if err != nil {
		panic(err)
	}

	c.Response.Status = http.StatusOK
	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText(string(jsonByte))
}

func (c Component) GetColumnModelDataForListTemplate() revel.Result {
	file, err := os.Open(`/home/hongjinqiu/goworkspace/src/finance/app/src/com/papersns/component/schema/SysUser.xml`)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	listTemplate := ListTemplate{}
	err = xml.Unmarshal(data, &listTemplate)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	templateManager := TemplateManager{}
	paramMap := map[string]string{}
	pageNo := 1
	pageSize := 10
	queryResult := templateManager.QueryDataForListTemplate(&listTemplate, paramMap, pageNo, pageSize)
	items := queryResult["items"].([]interface{})
	if len(items) > 1 {
		queryResult["items"] = items[:1]
	}

	columnResult := templateManager.GetColumnModelDataForListTemplate(&listTemplate, items[:1])
	jsonByte, err := json.MarshalIndent(columnResult, "", "\t")
	if err != nil {
		panic(err)
	}

	c.Response.Status = http.StatusOK
	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText(string(jsonByte))
}

func (c Component) GetToolbarForListTemplate() revel.Result {
	file, err := os.Open(`/home/hongjinqiu/goworkspace/src/finance/app/src/com/papersns/component/schema/SysUser.xml`)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	listTemplate := ListTemplate{}
	err = xml.Unmarshal(data, &listTemplate)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	templateManager := TemplateManager{}
	queryResult := templateManager.GetToolbarForListTemplate(&listTemplate)

	jsonByte, err := json.MarshalIndent(queryResult, "", "\t")
	if err != nil {
		panic(err)
	}

	c.Response.Status = http.StatusOK
	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText(string(jsonByte))
}

func (c Component) YUI() revel.Result {
	return c.Render()
}

func (c Component) Layout() revel.Result {
	return c.Render()
}
