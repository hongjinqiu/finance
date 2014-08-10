package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type FundSummarySupport struct {
	ActionSupport
}

type FundSummary struct {
	*revel.Controller
	BaseDataAction
}

func (c FundSummary) SaveData() revel.Result {
	c.RActionSupport = FundSummarySupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c FundSummary) DeleteData() revel.Result {
	c.RActionSupport = FundSummarySupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c FundSummary) EditData() revel.Result {
	c.RActionSupport = FundSummarySupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c FundSummary) NewData() revel.Result {
	c.RActionSupport = FundSummarySupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c FundSummary) GetData() revel.Result {
	c.RActionSupport = FundSummarySupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c FundSummary) CopyData() revel.Result {
	c.RActionSupport = FundSummarySupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c FundSummary) GiveUpData() revel.Result {
	c.RActionSupport = FundSummarySupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c FundSummary) RefreshData() revel.Result {
	c.RActionSupport = FundSummarySupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c FundSummary) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

