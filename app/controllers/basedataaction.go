package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/model"
	. "com/papersns/component"
	"com/papersns/mongo"
	"strings"
	"strconv"
	"encoding/json"
	. "com/papersns/model/handler"
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
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	c.beforeNewData(dataSource)
	bo := modelTemplateFactory.GetInstanceByDS(dataSource)
	c.afterNewData(dataSource, &bo)
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

func (c BaseDataAction) beforeNewData(dataSource DataSource) {
	
}

func (c BaseDataAction) afterNewData(dataSource DataSource, bo *map[string]interface{}) {
	
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
	
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	srcBo, found := querySupport.FindByMap(dataSourceModelId, queryMap)
	if !found {
		panic("CopyData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	modelTemplateFactory.ConvertDataType(dataSource, &srcBo)
	c.beforeCopyData(dataSource, srcBo)
	dataSource, bo := modelTemplateFactory.GetCopyInstance(dataSourceModelId, srcBo)
	c.afterCopyData(dataSource, &bo)
	
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

func (c BaseDataAction) beforeCopyData(dataSource DataSource, srcBo map[string]interface{}) {
	
}

func (c BaseDataAction) afterCopyData(dataSource DataSource, bo *map[string]interface{}) {
	
}

/**
 * 修改验证
 */
func (c BaseDataAction) editValidate(dataSource DataSource, bo map[string]interface{}) (string, bool) {
	return "", true
}

/**
 * 修改
 */
func (c BaseDataAction) EditData() revel.Result {
	dataSourceModelId := c.Params.Get("dataSourceModelId")
	jsonBo := c.Params.Get("jsonData")
	
	bo := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonBo), &bo)
	if err != nil {
		panic(err)
	}

	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	
	editMessage, isValid := c.editValidate(dataSource, bo)
	if !isValid {
		panic(editMessage)
	}
	
	c.beforeEditData(dataSource, &bo)
	
	financeService := FinanceService{}
	diffDataRowLi := financeService.SaveData(dataSource, &bo)
	
	c.afterEditData(dataSource, &bo, diffDataRowLi)
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

func (c BaseDataAction) beforeEditData(dataSource DataSource, bo *map[string]interface{}) {
	
}

func (c BaseDataAction) afterEditData(dataSource DataSource, bo *map[string]interface{}, diffDataRowLi *[]DiffDataRow) {
	
}

/**
 * 保存
 */
func (c BaseDataAction) SaveData() revel.Result {
	dataSourceModelId := c.Params.Get("dataSourceModelId")
	jsonBo := c.Params.Get("jsonData")
	
	bo := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonBo), &bo)
	if err != nil {
		panic(err)
	}

	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	c.beforeSaveData(dataSource, &bo)

	financeService := FinanceService{}
	diffDataRowLi := financeService.SaveData(dataSource, &bo)	
	
	c.afterSaveData(dataSource, &bo, diffDataRowLi)
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

func (c BaseDataAction) beforeSaveData(dataSource DataSource, bo *map[string]interface{}) {
	
}

func (c BaseDataAction) afterSaveData(dataSource DataSource, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	
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
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	c.beforeGiveUpData(dataSource, &bo)
	c.afterGiveUpData(dataSource, &bo)
	
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

func (c BaseDataAction) beforeGiveUpData(dataSource DataSource, bo *map[string]interface{}) {
	
}

func (c BaseDataAction) afterGiveUpData(dataSource DataSource, bo *map[string]interface{}) {
	
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
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	c.beforeDeleteData(dataSource, &bo)
	
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()

	usedCheck := UsedCheck{}
	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateDataBo(dataSource, &bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}){
		if fieldGroupLi[0].IsMasterField() {
			usedCheck.Delete(db, fieldGroupLi, *data)
		}
	})
	
	err = db.C(dataSourceModelId).Remove(queryMap)
	if err != nil {
		panic(err)
	}
	
	c.afterDeleteData(dataSource, &bo)
	
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

func (c BaseDataAction) beforeDeleteData(dataSource DataSource, bo *map[string]interface{}) {
	
}

func (c BaseDataAction) afterDeleteData(dataSource DataSource, bo *map[string]interface{}) {
	
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
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	c.beforeRefreshData(dataSource, &bo)
	c.afterRefreshData(dataSource, &bo)
	
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

func (c BaseDataAction) beforeRefreshData(dataSource DataSource, bo *map[string]interface{}) {
	
}

func (c BaseDataAction) afterRefreshData(dataSource DataSource, bo *map[string]interface{}) {
	
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
		"beReference": []interface{}{dataSourceModelId,"A", "id",id},
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
