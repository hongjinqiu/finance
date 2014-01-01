package model

import "github.com/robfig/revel"
import (
	. "com/papersns/model"
	. "com/papersns/mongo"
	"com/papersns/mongo"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
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
	seq := mongo.GetSequenceNo(db, sequenceName)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText(strconv.Itoa(seq))
}

func (c ModelTest) BeginTransaction() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	collections := []string{"Test1", "Test2"}
	txnManager.BeginTransaction(collections)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("BeginTransaction")
}

func (c ModelTest) Commit1() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	collections := []string{"Test1", "Test2"}
	txnId := txnManager.BeginTransaction(collections)
	
	test1 := map[string]interface{}{"name": "insertTest1"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"name": "insertTest2"}
	txnManager.Insert(txnId, "Test2", test2)
	
	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit1")
}

func (c ModelTest) Commit2() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	collections := []string{"Test1", "Test2"}
	txnId := txnManager.BeginTransaction(collections)
	
	test1 := map[string]interface{}{"name": "insertTest2"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"name": "insertTest2"}
	txnManager.Insert(txnId, "Test2", test2)
	
	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit2")
}

func (c ModelTest) Commit3() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	collections := []string{"Test1", "Test2"}
	txnId := txnManager.BeginTransaction(collections)
	
	test1 := map[string]interface{}{"name": "insertTest3"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"_id": 3, "name": "insertTest3"}
	txnManager.Update(txnId, "Test2", test2)
	
	txnManager.Commit(txnId)
	
	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit3")
}

func (c ModelTest) Commit4() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	collections := []string{"Test1", "Test2"}
	txnId := txnManager.BeginTransaction(collections)
	
	test1 := map[string]interface{}{"name": "insertTest4"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"_id": 3, "name": "insertTest4"}
	txnManager.Update(txnId, "Test2", test2)
	
	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit4")
}

func (c ModelTest) Commit5() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	collections := []string{"Test1", "Test2"}
	txnId := txnManager.BeginTransaction(collections)
	
	test1 := map[string]interface{}{"name": "insertTest5"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"_id": 3, "name": "insertTest5"}
	txnManager.Update(txnId, "Test2", test2)
	
	test3 := map[string]interface{}{"_id": 2, "name": "insertTest5"}
	txnManager.Remove(txnId, "Test2", test3)
	
	txnManager.Commit(txnId)
	
	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit5")
}


