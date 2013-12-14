package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/model"
	"com/papersns/mongo"
	"strconv"
	"strings"
	. "com/papersns/component"
)

func init() {
}

type BillAction struct {
	BaseDataAction
}

/**
 * 作废
 */
func (c BillAction) CancelData() revel.Result {
	dataSourceModelId := c.Params.Get("dataSourceModelId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	bo, found := querySupport.FindByMap(dataSourceModelId, queryMap)
	if !found {
		panic("CancelData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	c.beforeCancelData(dataSource, &bo)
	mainData := bo["A"].(map[string]interface{})
	mainData["billStatus"] = 4
	
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()
	db.C(dataSourceModelId).Update(queryMap, bo)
	
	c.afterCancelData(dataSource, &bo)
	
	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(map[string]interface{}{
			"bo": bo,
			"dataSource": dataSource,
		})
	}
	return c.Render()
}

func (c BillAction) beforeCancelData(dataSource DataSource, bo *map[string]interface{}) {
	
}

func (c BillAction) afterCancelData(dataSource DataSource, bo *map[string]interface{}) {
	
}

/**
 * 反作废
 */
func (c BillAction) UnCancelData() revel.Result {
	dataSourceModelId := c.Params.Get("dataSourceModelId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	bo, found := querySupport.FindByMap(dataSourceModelId, queryMap)
	if !found {
		panic("UnCancelData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	c.beforeUnCancelData(dataSource, &bo)
	mainData := bo["A"].(map[string]interface{})
	mainData["billStatus"] = 0
	
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	db.C(dataSourceModelId).Update(queryMap, bo)
	defer session.Close()
	
	c.afterUnCancelData(dataSource, &bo)
	
	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(map[string]interface{}{
			"bo": bo,
			"dataSource": dataSource,
		})
	}
	return c.Render()
}

func (c BillAction) beforeUnCancelData(dataSource DataSource, bo *map[string]interface{}) {
	
}

func (c BillAction) afterUnCancelData(dataSource DataSource, bo *map[string]interface{}) {
	
}
