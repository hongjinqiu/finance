package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/component"
	"com/papersns/global"
	. "com/papersns/model"
	. "com/papersns/mongo"
	"strconv"
	"strings"
)

func init() {
}

type BillAction struct {
	BaseDataAction
}

func (c BillAction) SaveData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, relationBo, dataSource := c.saveCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BillAction) DeleteData() revel.Result {
	c.actionSupport = ActionSupport{}

	bo, relationBo, dataSource := c.deleteDataCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BillAction) EditData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, relationBo, dataSource := c.editDataCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BillAction) NewData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, relationBo, dataSource := c.newDataCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BillAction) GetData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, relationBo, dataSource := c.getDataCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 复制
 */
func (c BillAction) CopyData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, relationBo, dataSource := c.copyDataCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillAction) GiveUpData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, relationBo, dataSource := c.giveUpDataCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 刷新
 */
func (c BillAction) RefreshData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, relationBo, dataSource := c.refreshDataCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BillAction) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

/**
 * 作废
 */
func (c BillAction) CancelData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, relationBo, dataSource := c.cancelDataCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BillAction) cancelDataCommon() (map[string]interface{}, map[string]interface{}, DataSource) {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	dataSourceModelId := c.Params.Get("dataSourceModelId")
	formTemplateId := c.Params.Get("formTemplateId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	bo, found := querySupport.FindByMap(collectionName, queryMap)
	if !found {
		panic("CancelData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}

	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	c.setModifyFixFieldValue(sessionId, dataSource, &bo)
	c.actionSupport.beforeCancelData(sessionId, dataSource, &bo)
	mainData := bo["A"].(map[string]interface{})
	mainData["billStatus"] = 4

	_, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	_, updateResult := txnManager.Update(txnId, dataSourceModelId, bo)
	if !updateResult {
		panic("作废失败")
	}

	c.actionSupport.afterCancelData(sessionId, dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	
	bo, _ = querySupport.FindByMap(collectionName, queryMap)
	
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	columnModelData := templateManager.GetColumnModelDataForFormTemplate(formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})
	
	return bo, relationBo, dataSource
}

/**
 * 反作废
 */
func (c BillAction) UnCancelData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, relationBo, dataSource := c.unCancelDataCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BillAction) unCancelDataCommon() (map[string]interface{}, map[string]interface{}, DataSource) {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	dataSourceModelId := c.Params.Get("dataSourceModelId")
	formTemplateId := c.Params.Get("formTemplateId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	bo, found := querySupport.FindByMap(collectionName, queryMap)
	if !found {
		panic("UnCancelData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}

	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	c.setModifyFixFieldValue(sessionId, dataSource, &bo)
	c.actionSupport.beforeUnCancelData(sessionId, dataSource, &bo)
	mainData := bo["A"].(map[string]interface{})
	mainData["billStatus"] = 0

	_, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	_, updateResult := txnManager.Update(txnId, dataSourceModelId, bo)
	if !updateResult {
		panic("反作废失败")
	}

	c.actionSupport.afterUnCancelData(sessionId, dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	
	bo, _ = querySupport.FindByMap(collectionName, queryMap)
	
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	columnModelData := templateManager.GetColumnModelDataForFormTemplate(formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})
	
	return bo, relationBo, dataSource
}
