package controllers

import "github.com/robfig/revel"
import (
	"strings"
	. "com/papersns/model"
	. "com/papersns/component"
	. "com/papersns/error"
	"com/papersns/global"
)

func init() {
}

type BankAccountSupport struct {
	ActionSupport
}

func (o BankAccountSupport) afterNewData(sessionId int, dataSource DataSource, bo *map[string]interface{}) {
	modelTemplateFactory := ModelTemplateFactory{}
	dataSetId := "B"
	data := modelTemplateFactory.GetDataSetNewData(dataSource, dataSetId, *bo)
	
	// 设置默认的币别
	qb := QuerySupport{}
	session, _ := global.GetConnection(sessionId)
	collection := "CurrencyType"
	query := map[string]interface{}{
		"A.code": "RMB",
	}
	currencyType, found := qb.FindByMapWithSession(session, collection, query)
	if !found {
		panic(BusinessError{"没有找到币别人民币，请先配置默认币别"})
	}
	data["currencyTypeId"] = currencyType["id"]
	
	(*bo)["B"] = []interface{}{
		data,
	}
}

type BankAccount struct {
	BaseDataAction
}

func (c BankAccount) SaveData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankAccount) DeleteData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankAccount) EditData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankAccount) NewData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankAccount) GetData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BankAccount) CopyData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BankAccount) GiveUpData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BankAccount) RefreshData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankAccount) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

