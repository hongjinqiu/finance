package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
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

func (o CashAccountSupport) RAfterNewData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
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
func (o CashAccountSupport) RBeforeSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
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
	*revel.Controller
	BaseDataAction
}

func (c CashAccount) SaveData() revel.Result {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CashAccount) DeleteData() revel.Result {
	c.RActionSupport = CashAccountSupport{}

	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CashAccount) EditData() revel.Result {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CashAccount) NewData() revel.Result {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CashAccount) GetData() revel.Result {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c CashAccount) CopyData() revel.Result {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CashAccount) GiveUpData() revel.Result {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c CashAccount) RefreshData() revel.Result {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CashAccount) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
