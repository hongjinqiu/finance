package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/model"
	. "com/papersns/component"
	"com/papersns/mongo"
	"strings"
	"strconv"
)

func init() {
}

type BaseDataAction struct {
	*revel.Controller
}

/**
 * 列表页
 */
//func (baseData BaseDataAction) ListData() revel.Result {
//	
//}

/**
 * 新增
 */
func (c BaseDataAction) NewData() revel.Result {
	dataSourceModelId := c.Params.Get("dataSourceModelId")
	c.beforeNewData(dataSourceModelId)
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource, bo := modelTemplateFactory.GetInstance(dataSourceModelId)
	c.afterNewData(&dataSource, &bo)
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

func (c BaseDataAction) beforeNewData(dataSourceModelId string) {
	
}

func (c BaseDataAction) afterNewData(dataSource *DataSource, bo *map[string]interface{}) {
	
}

/**
 * 复制
 */
func (c BaseDataAction) CopyData() revel.Result {
	dataSourceModelId := c.Params.Get("dataSourceModelId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	
	c.beforeCopyData(dataSourceModelId, id)
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	srcBo, found := querySupport.FindByMap(dataSourceModelId, queryMap)
	if !found {
		panic("CopyData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource, bo := modelTemplateFactory.GetCopyInstance(dataSourceModelId, srcBo)
	c.afterCopyData(&dataSource, &bo)
	
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

func (c BaseDataAction) beforeCopyData(dataSourceModelId string, id int) {
	
}

func (c BaseDataAction) afterCopyData(dataSource *DataSource, bo *map[string]interface{}) {
	
}

/**
 * 修改验证
 * 业务为:已作废不可编辑
 */
func (c BaseDataAction) editValidate(dataSourceModelId string, id int) bool {
	return true
}

/**
 * 修改
 */
func (c BaseDataAction) EditData() revel.Result {
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
		panic("EditData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	c.beforeEditData(dataSourceModelId, &bo)
	
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()
	err = db.C(dataSourceModelId).Insert(bo)
	if err != nil {
		panic(err)
	}
	
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource, _ := modelTemplateFactory.GetInstance(dataSourceModelId)
	c.afterEditData(&dataSource, &bo)
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

func (c BaseDataAction) beforeEditData(dataSourceModelId string, bo *map[string]interface{}) {
	
}

func (c BaseDataAction) afterEditData(dataSource *DataSource, bo *map[string]interface{}) {
	
}

/**
 * 保存
 */
func (c BaseDataAction) SaveData() revel.Result {
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
		panic("SaveData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	c.beforeSaveData(dataSourceModelId, &bo)
	
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()
	err = db.C(dataSourceModelId).Insert(bo)
	if err != nil {
		panic(err)
	}
	
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource, _ := modelTemplateFactory.GetInstance(dataSourceModelId)
	c.afterSaveData(&dataSource, &bo)
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

func (c BaseDataAction) beforeSaveData(dataSourceModelId string, bo *map[string]interface{}) {
	
}

func (c BaseDataAction) afterSaveData(dataSource *DataSource, bo *map[string]interface{}) {
	
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BaseDataAction) GiveUpData() revel.Result {
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
		panic("CopyData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	c.beforeGiveUpData(dataSourceModelId, &bo)
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource, _ := modelTemplateFactory.GetInstance(dataSourceModelId)
	c.afterGiveUpData(&dataSource, &bo)
	
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

func (c BaseDataAction) beforeGiveUpData(dataSourceModelId string, bo *map[string]interface{}) {
	
}

func (c BaseDataAction) afterGiveUpData(dataSource *DataSource, bo *map[string]interface{}) {
	
}

/**
 * 删除
 */
func (c BaseDataAction) DeleteData() revel.Result {
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
		panic("CopyData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	c.beforeDeleteData(dataSourceModelId, &bo)
	
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()
	err = db.C(dataSourceModelId).Remove(queryMap)
	if err != nil {
		panic(err)
	}
	
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource, _ := modelTemplateFactory.GetInstance(dataSourceModelId)
	c.afterDeleteData(&dataSource, &bo)
	
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

func (c BaseDataAction) beforeDeleteData(dataSourceModelId string, bo *map[string]interface{}) {
	
}

func (c BaseDataAction) afterDeleteData(dataSource *DataSource, bo *map[string]interface{}) {
	
}

/**
 * 刷新
 */
func (c BaseDataAction) RefreshData() revel.Result {
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
		panic("CopyData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	c.beforeRefreshData(dataSourceModelId, &bo)
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource, _ := modelTemplateFactory.GetInstance(dataSourceModelId)
	c.afterRefreshData(&dataSource, &bo)
	
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

func (c BaseDataAction) beforeRefreshData(dataSourceModelId string, bo *map[string]interface{}) {
	
}

func (c BaseDataAction) afterRefreshData(dataSource *DataSource, bo *map[string]interface{}) {
	
}

/**
 * 修改日志
 */
func (c BaseDataAction) LogList() revel.Result {
	dataSourceModelId := c.Params.Get("dataSourceModelId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	
	collectionName := "PubReferenceLog"
	// reference,beReference
	querySupport := QuerySupport{}
	query := map[string]interface{}{
		"beReference": []interface{}{dataSourceModelId,id},
	}
	pageNo := 1
	pageSize := 10
	orderBy := ""
	result := querySupport.Index(collectionName, query, pageNo, pageSize, orderBy)
	
	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}



