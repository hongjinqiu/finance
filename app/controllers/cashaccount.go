package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type CashAccountSupport struct {
	ActionSupport
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
	return c.Render()
}

