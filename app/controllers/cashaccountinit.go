package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type CashAccountInitSupport struct {
	ActionSupport
}

type CashAccountInit struct {
	BaseDataAction
}

func (c CashAccountInit) SaveData() revel.Result {
	c.actionSupport = CashAccountInitSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashAccountInit) DeleteData() revel.Result {
	c.actionSupport = CashAccountInitSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashAccountInit) EditData() revel.Result {
	c.actionSupport = CashAccountInitSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashAccountInit) NewData() revel.Result {
	c.actionSupport = CashAccountInitSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashAccountInit) GetData() revel.Result {
	c.actionSupport = CashAccountInitSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c CashAccountInit) CopyData() revel.Result {
	c.actionSupport = CashAccountInitSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CashAccountInit) GiveUpData() revel.Result {
	c.actionSupport = CashAccountInitSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c CashAccountInit) RefreshData() revel.Result {
	c.actionSupport = CashAccountInitSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashAccountInit) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

