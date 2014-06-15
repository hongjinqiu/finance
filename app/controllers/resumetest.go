package controllers

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

type ResumeTest struct {
	*revel.Controller
}

func (c ResumeTest) Schema() revel.Result {
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

func (c ResumeTest) CopyTest() revel.Result {
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

func (c ResumeTest) SeqTest() revel.Result {
	//	db *mgo.Database, sequenceName string
	connectionFactory := mongo.GetInstance()
	session, db := connectionFactory.GetConnection()
	defer session.Close()
	sequenceName := "sysUserId"
	seq := mongo.GetSequenceNo(db, sequenceName)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText(strconv.Itoa(seq))
}

func (c ResumeTest) BeginTransaction() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnManager.BeginTransaction()

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("BeginTransaction")
}

func (c ResumeTest) Commit1() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{"name": "insertTest1"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"name": "insertTest2"}
	txnManager.Insert(txnId, "Test2", test2)

	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit1")
}

func (c ResumeTest) Commit2() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{"name": "insertTest2"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"name": "insertTest2"}
	txnManager.Insert(txnId, "Test2", test2)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit2")
}

func (c ResumeTest) Commit3() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{"name": "insertTest3"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"_id": 3, "name": "insertTest3"}
	txnManager.Update(txnId, "Test2", test2)

	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit3")
}

func (c ResumeTest) Commit4() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{"name": "insertTest4"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"_id": 3, "name": "insertTest4"}
	txnManager.Update(txnId, "Test2", test2)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit4")
}

func (c ResumeTest) Commit5() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

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

func (c ResumeTest) Commit6() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{"name": "insertTest6"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"_id": 4, "name": "insertTest6"}
	txnManager.Update(txnId, "Test2", test2)

	test3 := map[string]interface{}{"_id": 2, "name": "insertTest6"}
	txnManager.Remove(txnId, "Test2", test3)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit6")
}

func (c ResumeTest) Commit7() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{"name": "insertTest7"}
	txnManager.Insert(txnId, "Test1", test1)

	test3 := map[string]interface{}{"_id": 2, "name": "insertTest7"}
	txnManager.Remove(txnId, "Test2", test3)

	test2 := map[string]interface{}{"_id": 4, "name": "insertTest7"}
	txnManager.Update(txnId, "Test2", test2)

	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit7")
}

func (c ResumeTest) Commit8() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{"name": "insertTest8"}
	txnManager.Insert(txnId, "Test1", test1)

	test3 := map[string]interface{}{"_id": 3, "name": "insertTest8"}
	txnManager.Remove(txnId, "Test2", test3)

	test2 := map[string]interface{}{"_id": 4, "name": "insertTest8"}
	txnManager.Update(txnId, "Test2", test2)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit8")
}

func (c ResumeTest) Commit9() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{"name": "insertTest9"}
	txnManager.Insert(txnId, "Test1", test1)

	test3 := map[string]interface{}{"_id": 3, "name": "insertTest9"}
	txnManager.Remove(txnId, "Test2", test3)

	test2 := map[string]interface{}{"_id": 4, "name": "insertTest9"}
	txnManager.Update(txnId, "Test2", test2)

	test4 := map[string]interface{}{"_id": 5, "name": "insertTest9"}
	txnManager.Remove(txnId, "Test2", test4)

	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-9"
	return c.RenderText("Commit9")
}

func (c ResumeTest) Commit10() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{"name": "insertTest10"}
	txnManager.Insert(txnId, "Test1", test1)

	test3 := map[string]interface{}{"_id": 4, "name": "insertTest10"}
	txnManager.Remove(txnId, "Test2", test3)

	test2 := map[string]interface{}{"_id": 6, "name": "insertTest10"}
	txnManager.Update(txnId, "Test2", test2)

	test4 := map[string]interface{}{"_id": 7, "name": "insertTest10"}
	txnManager.Remove(txnId, "Test2", test4)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit10")
}

func (c ResumeTest) Commit11() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	{
		query := map[string]interface{}{
			"_id": map[string]interface{}{
				"$lt": 7,
			},
		}
		update := map[string]interface{}{
			"$set": map[string]interface{}{
				"age": 11,
			},
		}
		unModify := map[string]interface{}{
			"$unset": map[string]interface{}{
				"age": 1,
			},
		}
		txnManager.UpdateAll(txnId, "Test1", query, update, unModify)
	}
	{
		query := map[string]interface{}{
			"_id": map[string]interface{}{
				"$gte": 8,
			},
		}
		update := map[string]interface{}{
			"$set": map[string]interface{}{
				"age": 12,
			},
		}
		unModify := map[string]interface{}{
			"$unset": map[string]interface{}{
				"age": 1,
			},
		}
		txnManager.UpdateAll(txnId, "Test2", query, update, unModify)
	}
	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit11")
}

func (c ResumeTest) Commit12() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	{
		query := map[string]interface{}{
			"_id": map[string]interface{}{
				"$lt": 7,
			},
		}
		update := map[string]interface{}{
			"$set": map[string]interface{}{
				"age": 9,
			},
		}
		unModify := map[string]interface{}{
			"$unset": map[string]interface{}{
				"age": 1,
			},
		}
		txnManager.UpdateAll(txnId, "Test1", query, update, unModify)
	}
	{
		query := map[string]interface{}{
			"_id": map[string]interface{}{
				"$gte": 8,
			},
		}
		update := map[string]interface{}{
			"$set": map[string]interface{}{
				"age": 9,
			},
		}
		unModify := map[string]interface{}{
			"$unset": map[string]interface{}{
				"age": 1,
			},
		}
		txnManager.UpdateAll(txnId, "Test2", query, update, unModify)
	}

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit12")
}

func (c ResumeTest) Commit13() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{
		"name": "insertTest13",
	}
	txnManager.Insert(txnId, "Test1", test1)

	test1["age"] = 20
	txnManager.Update(txnId, "Test1", test1)

	txnManager.Remove(txnId, "Test1", test1)

	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit13")
}

func (c ResumeTest) Commit14() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{
		"name": "insertTest14",
	}
	txnManager.Insert(txnId, "Test1", test1)

	test1["age"] = 20
	txnManager.Update(txnId, "Test1", test1)

	txnManager.Remove(txnId, "Test1", test1)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit14")
}

func (c ResumeTest) Commit15() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{
		"name": "insertTest15",
	}
	test1 = txnManager.Insert(txnId, "Test1", test1)
	test1, _ = txnManager.Remove(txnId, "Test1", test1)
	test1["age"] = 20
	test1, _ = txnManager.Update(txnId, "Test1", test1)

	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit15")
}

func (c ResumeTest) Commit16() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{
		"name": "insertTest16",
	}
	test1 = txnManager.Insert(txnId, "Test1", test1)
	test1, _ = txnManager.Remove(txnId, "Test1", test1)
	test1["age"] = 20
	test1, _ = txnManager.Update(txnId, "Test1", test1)

	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit16")
}

func (c ResumeTest) Commit17() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{
		"name": "insertTest17",
	}
	test1 = txnManager.Insert(txnId, "Test1", test1)
	test1, _ = txnManager.Remove(txnId, "Test1", test1)
	test1["age"] = 20
	test1, _ = txnManager.Remove(txnId, "Test1", test1)

	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit17")
}

func (c ResumeTest) Commit18() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{
		"name": "insertTest18",
	}
	test1 = txnManager.Insert(txnId, "Test1", test1)
	test1, _ = txnManager.Remove(txnId, "Test1", test1)
	test1["age"] = 20
	test1, _ = txnManager.Remove(txnId, "Test1", test1)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit18")
}

func (c ResumeTest) Commit19() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{
		"name": "insertTest19",
	}
	test1 = txnManager.Insert(txnId, "Test1", test1)
	test1["address"] = "xiamen"
	test1, _ = txnManager.Update(txnId, "Test1", test1)
	test1["age"] = 20
	test1, _ = txnManager.Update(txnId, "Test1", test1)

	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit19")
}

func (c ResumeTest) Commit20() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{
		"name": "insertTest20",
	}
	test1 = txnManager.Insert(txnId, "Test1", test1)
	test1["address"] = "xiamen"
	test1, _ = txnManager.Update(txnId, "Test1", test1)
	test1["age"] = 20
	test1, _ = txnManager.Update(txnId, "Test1", test1)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit20")
}

func (c ResumeTest) SelectForUpdate1() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{
		"_id":  36,
		"name": "SelectForUpdate1",
	}
	txnManager.SelectForUpdate(txnId, "Test1", test1)
	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("SelectForUpdate1")
}

func (c ResumeTest) SelectForUpdate2() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{
		"_id":  36,
		"name": "SelectForUpdate2",
	}
	txnManager.SelectForUpdate(txnId, "Test1", test1)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("SelectForUpdate2")
}

func (c ResumeTest) SelectForUpdate3() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{
		"_id":  36,
		"name": "SelectForUpdate3",
	}
	txnManager.SelectForUpdate(txnId, "Test1", test1)
	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("SelectForUpdate3")
}

func (c ResumeTest) UpdateAll1() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	query := map[string]interface{}{
		"_id": 36,
	}
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"age": 20,
		},
	}
	unModify := map[string]interface{}{
		"$unset": map[string]interface{}{
			"age": 1,
		},
	}
	txnManager.UpdateAll(txnId, "Test1", query, update, unModify)
	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("UpdateAll1")
}

func (c ResumeTest) UpdateAll2() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	query := map[string]interface{}{
		"_id": 36,
	}
	update := map[string]interface{}{
		"$inc": map[string]interface{}{
			"age": 5,
		},
	}
	unModify := map[string]interface{}{
		"$inc": map[string]interface{}{
			"age": -5,
		},
	}
	txnManager.UpdateAll(txnId, "Test1", query, update, unModify)
	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("UpdateAll1")
}

func (c ResumeTest) UpdateAll3() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	query := map[string]interface{}{
		"_id": map[string]interface{}{
			"$lte": 100,
		},
	}
	update := map[string]interface{}{
		"$inc": map[string]interface{}{
			"age": 5,
		},
	}
	unModify := map[string]interface{}{
		"$inc": map[string]interface{}{
			"age": -5,
		},
	}
	obj, _ := txnManager.UpdateAll(txnId, "Test1", query, update, unModify)
	fmt.Println("obj3", obj)
	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("UpdateAll3")
}

func (c ResumeTest) RemoveAll1() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	query := map[string]interface{}{
		"_id": 5,
	}
	obj, _ := txnManager.RemoveAll(txnId, "Test1", query)
	fmt.Println("removeAll1", obj)
	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("RemoveAll1")
}

func (c ResumeTest) RemoveAll2() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnId := txnManager.BeginTransaction()

	query := map[string]interface{}{
		"_id": map[string]interface{}{
			"$lte": 100,
		},
	}
	obj, _ := txnManager.RemoveAll(txnId, "Test1", query)
	fmt.Println("RemoveAll2", obj)
	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("RemoveAll2")
}

func (c ResumeTest) Rollback1() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()

	test1 := map[string]interface{}{"name": "rollbackTest1"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"name": "rollbackTest2"}
	txnManager.Insert(txnId, "Test2", test2)

	txnManager.Rollback(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Rollback1")
}

func (c ResumeTest) ResumeBeginCommit1() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnManager.Resume()

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("ResumeBeginCommit1")
}

func (c ResumeTest) ResumeBeginCommit2() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}
	txnManager.ResumePeriod()

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("ResumeBeginCommit2")
}

func (c ResumeTest) RecoverTest() revel.Result {
	connectionFactory := mongo.ConnectionFactory{}
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	txnManager := TxnManager{db}

	txnId := txnManager.BeginTransaction()
	defer func() {
		if x := recover(); x != nil {
			txnManager.Rollback(txnId)
			panic(x)
		}
	}()

	test1 := map[string]interface{}{"name": "insertTest1"}
	txnManager.Insert(txnId, "Test1", test1)
	test2 := map[string]interface{}{"name": "insertTest2"}
	txnManager.Insert(txnId, "Test2", test2)

	txnManager.Commit(txnId)

	c.Response.ContentType = "text/plain; charset=utf-8"
	return c.RenderText("Commit1")
}
