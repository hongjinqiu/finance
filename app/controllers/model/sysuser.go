package model

import "github.com/robfig/revel"
import (
	. "com/papersns/model"
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func init() {
}

type ModelTest struct {
	*revel.Controller
}

func (c ModelTest) Schema() revel.Result {
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource, bo := modelTemplateFactory.GetInstance("SysUser")
	println(bo)
	// 1.query data,
	// from data-provider
	// from query-parameters

	xmlDataArray, err := xml.MarshalIndent(&dataSource, "", "\t")
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	jsonDataArray, err := json.MarshalIndent(&bo, "", "\t")
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}
	
	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText(string(xmlDataArray))
	return c.RenderText(string(jsonDataArray))
}
