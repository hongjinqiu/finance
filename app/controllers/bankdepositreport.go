package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type BankDepositReportSupport struct {
	ActionSupport
}

type BankDepositReport struct {
	BaseDataAction
}

func (c BankDepositReport) SaveData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankDepositReport) DeleteData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankDepositReport) EditData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankDepositReport) NewData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankDepositReport) GetData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BankDepositReport) CopyData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BankDepositReport) GiveUpData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BankDepositReport) RefreshData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankDepositReport) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

