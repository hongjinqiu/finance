package controllers

import "github.com/robfig/revel"
import (
	"com/papersns/mongo"
)

func init() {
}

type Tree struct {
	*revel.Controller
}

func (o Tree) SysUserTest() revel.Result {
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()

	collection := "SysUser"
	c := db.C(collection)
	
	queryMap := map[string]interface{}{}
	
	sysUserResult := []map[string]interface{}{}
	err := c.Find(queryMap).Limit(10).All(&sysUserResult)
	if err != nil {
		panic(err)
	}
	
	items := []map[string]interface{}{}
	for _, item := range sysUserResult {
		items = append(items, map[string]interface{}{
			"code": item["_id"],
			"name": item["nick"],
		})
	}
	
	o.Response.ContentType = "application/json; charset=utf-8"
	return o.RenderJson(&items)
}

