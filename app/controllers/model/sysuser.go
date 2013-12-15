package model

import "github.com/robfig/revel"
import (
	. "com/papersns/model"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"com/papersns/mongo"
)

func init() {
}

type ModelTest struct {
	*revel.Controller
}

func (c ModelTest) Schema() revel.Result {
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource, bo := modelTemplateFactory.GetInstance("SysUser")
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

func (c ModelTest) CopyTest() revel.Result {
	modelTemplateFactory := ModelTemplateFactory{}
//		"billStatus": 5,
	srcJson := `{
	"A": {
		"attachCount": "3",
		"code": "45",
		"createBy": 0,
		"createTime": 0,
		"endDate": 0,
		"id": 99,
		"modifyBy": 0,
		"modifyTime": 0,
		"nick": "名称",
		"nick_ref": {
			"DisplayField": "nick",
			"RelationDataSetId": "A",
			"RelationExpr": true,
			"RelationModelId": "SysUser"
		},
		"remark": "",
		"startDate": 0
	},
	"B": [{"id": 3, "attachCount": 5}]
}`
	srcBo := map[string]interface{}{}
	err := json.Unmarshal([]byte(srcJson), &srcBo)
	if err != nil {
		panic(err)
	}
	dataSource, bo := modelTemplateFactory.GetCopyInstance("SysUser", srcBo)
	
	
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
	return c.RenderText(string(jsonDataArray))
	return c.RenderText(string(xmlDataArray))
}

func (c ModelTest) SeqTest() revel.Result {
//	db *mgo.Database, sequenceName string
	connectionFactory := mongo.GetInstance()
	session, db := connectionFactory.GetConnection()
	defer session.Close()
	sequenceName := "sysUserId"
	seq := GetSequenceNo(db, sequenceName)
	
	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText(strconv.Itoa(seq))
}

