package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type CurrencyTypeSupport struct {
	ActionSupport
}

type CurrencyType struct {
	BaseDataAction
}

func (c CurrencyType) SaveData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CurrencyType) DeleteData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}

	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CurrencyType) EditData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CurrencyType) NewData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CurrencyType) GetData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c CurrencyType) CopyData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CurrencyType) GiveUpData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c CurrencyType) RefreshData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CurrencyType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
