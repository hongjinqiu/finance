package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type FundSummarySupport struct {
	ActionSupport
}

type FundSummary struct {
	BaseDataAction
}

func (c FundSummary) SaveData() revel.Result {
	c.actionSupport = FundSummarySupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c FundSummary) DeleteData() revel.Result {
	c.actionSupport = FundSummarySupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c FundSummary) EditData() revel.Result {
	c.actionSupport = FundSummarySupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c FundSummary) NewData() revel.Result {
	c.actionSupport = FundSummarySupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c FundSummary) GetData() revel.Result {
	c.actionSupport = FundSummarySupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c FundSummary) CopyData() revel.Result {
	c.actionSupport = FundSummarySupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c FundSummary) GiveUpData() revel.Result {
	c.actionSupport = FundSummarySupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c FundSummary) RefreshData() revel.Result {
	c.actionSupport = FundSummarySupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c FundSummary) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

