package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/component"
	. "com/papersns/error"
	"com/papersns/global"
	. "com/papersns/model"
	"strconv"
	"strings"
)

func init() {
}

type CashAccountSupport struct {
	ActionSupport
}

func (o CashAccountSupport) afterNewData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	masterData := (*bo)["A"].(map[string]interface{})
	(*bo)["A"] = masterData
	
	session, _ := global.GetConnection(sessionId)
	qb := QuerySupport{}
	query := map[string]interface{}{
		"A.code": "RMB",
	}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		query[k] = v
	}
	
	collectionName := "CurrencyType"
	result, found := qb.FindByMapWithSession(session, collectionName, query)
	if found {
		masterData["currencyTypeId"] = result["id"]
	}
}

/**
* 为避免并发问题,重设amtOriginalCurrencyBalance为数据库中值
 */
func (o CashAccountSupport) beforeSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	session, _ := global.GetConnection(sessionId)
	modelTemplateFactory := ModelTemplateFactory{}
	strId := modelTemplateFactory.GetStrId(*bo)
	if strId != "" && strId != "0" {
		id, err := strconv.Atoi(strId)
		if err != nil {
			panic(err)
		}
		qb := QuerySupport{}
		queryMap := map[string]interface{}{
			"_id":          id,
		}
		permissionSupport := PermissionSupport{}
		permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
		for k, v := range permissionQueryDict {
			queryMap[k] = v
		}
		
		collectionName := "CashAccount"
		boInDb, found := qb.FindByMapWithSession(session, collectionName, queryMap)
		if !found {
			panic(BusinessError{Message: "现金账户保存前，现金账户未找到"})
		}

		masterData := (*bo)["A"].(map[string]interface{})
		(*bo)["A"] = masterData

		masterDataInDb := boInDb["A"].(map[string]interface{})
		boInDb["A"] = masterDataInDb

		masterData["amtOriginalCurrencyBalance"] = masterDataInDb["amtOriginalCurrencyBalance"]
	}
}

type CashAccount struct {
	BaseDataAction
}

func (c CashAccount) SaveData() revel.Result {
	c.actionSupport = CashAccountSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashAccount) DeleteData() revel.Result {
	c.actionSupport = CashAccountSupport{}

	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashAccount) EditData() revel.Result {
	c.actionSupport = CashAccountSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashAccount) NewData() revel.Result {
	c.actionSupport = CashAccountSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashAccount) GetData() revel.Result {
	c.actionSupport = CashAccountSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c CashAccount) CopyData() revel.Result {
	c.actionSupport = CashAccountSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CashAccount) GiveUpData() revel.Result {
	c.actionSupport = CashAccountSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c CashAccount) RefreshData() revel.Result {
	c.actionSupport = CashAccountSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashAccount) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
