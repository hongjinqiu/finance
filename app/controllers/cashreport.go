package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type CashReportSupport struct {
	ActionSupport
}

type CashReport struct {
	BaseDataAction
}

func (c CashReport) SaveData() revel.Result {
	c.actionSupport = CashReportSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashReport) DeleteData() revel.Result {
	c.actionSupport = CashReportSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashReport) EditData() revel.Result {
	c.actionSupport = CashReportSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashReport) NewData() revel.Result {
	c.actionSupport = CashReportSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashReport) GetData() revel.Result {
	c.actionSupport = CashReportSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c CashReport) CopyData() revel.Result {
	c.actionSupport = CashReportSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CashReport) GiveUpData() revel.Result {
	c.actionSupport = CashReportSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c CashReport) RefreshData() revel.Result {
	c.actionSupport = CashReportSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CashReport) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

