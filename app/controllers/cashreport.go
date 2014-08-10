package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type CashReportSupport struct {
	ActionSupport
}

type CashReport struct {
	*revel.Controller
	BaseDataAction
}

func (c CashReport) SaveData() revel.Result {
	c.RActionSupport = CashReportSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CashReport) DeleteData() revel.Result {
	c.RActionSupport = CashReportSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CashReport) EditData() revel.Result {
	c.RActionSupport = CashReportSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CashReport) NewData() revel.Result {
	c.RActionSupport = CashReportSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CashReport) GetData() revel.Result {
	c.RActionSupport = CashReportSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c CashReport) CopyData() revel.Result {
	c.RActionSupport = CashReportSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CashReport) GiveUpData() revel.Result {
	c.RActionSupport = CashReportSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c CashReport) RefreshData() revel.Result {
	c.RActionSupport = CashReportSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CashReport) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

