package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/component"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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

func (c Component) MongoTest() revel.Result {
	querySupport := QuerySupport{}
	m, isFind := querySupport.Find("SysUser", `{"_id": 15}`)
	if !isFind {
		panic("not found")
	}

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
	query := `{}`
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
