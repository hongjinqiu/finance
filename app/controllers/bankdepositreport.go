package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type BankDepositReportSupport struct {
	ActionSupport
}

type BankDepositReport struct {
	*revel.Controller
	BaseDataAction
}

func (c BankDepositReport) SaveData() revel.Result {
	c.RActionSupport = BankDepositReportSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BankDepositReport) DeleteData() revel.Result {
	c.RActionSupport = BankDepositReportSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BankDepositReport) EditData() revel.Result {
	c.RActionSupport = BankDepositReportSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BankDepositReport) NewData() revel.Result {
	c.RActionSupport = BankDepositReportSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BankDepositReport) GetData() revel.Result {
	c.RActionSupport = BankDepositReportSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BankDepositReport) CopyData() revel.Result {
	c.RActionSupport = BankDepositReportSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BankDepositReport) GiveUpData() revel.Result {
	c.RActionSupport = BankDepositReportSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BankDepositReport) RefreshData() revel.Result {
	c.RActionSupport = BankDepositReportSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BankDepositReport) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

