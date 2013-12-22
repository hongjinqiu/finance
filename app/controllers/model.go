package controllers

import "github.com/robfig/revel"
import (
	"os"
	"io/ioutil"
	"encoding/xml"
	"encoding/json"
	. "com/papersns/model"
)

func init() {
}

type Model struct {
	*revel.Controller
}

func (m Model) FieldTest() revel.Result {
	file, err := os.Open("/home/hongjinqiu/goworkspace/src/finance/app/src/com/papersns/model/xml/fieldpool.xml")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	
	fields := Fields{}
	err = xml.Unmarshal(data, &fields)
	if err != nil {
		panic(err)
	}
	
	m.Response.ContentType = "application/json; charset=utf-8"
	return m.RenderJson(fields)
}

func (m Model) DataSourceTest() revel.Result {
	file, err := os.Open("/home/hongjinqiu/goworkspace/src/finance/app/src/com/papersns/model/demo/datasource/pc_ds_billtype.xml")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	
	dataSource := DataSource{}
	err = xml.Unmarshal(data, &dataSource)
	if err != nil {
		panic(err)
	}
	
	dataSource2 := DataSource{}
	jsonByte, err := json.Marshal(dataSource)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonByte, &dataSource2)
	if err != nil {
		panic(err)
	}
	
	m.Response.ContentType = "application/json; charset=utf-8"
	return m.RenderJson(dataSource2)
}
