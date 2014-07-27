package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type IncomeItemSupport struct {
	ActionSupport
}

type IncomeItem struct {
	BaseDataAction
}

func (c IncomeItem) SaveData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c IncomeItem) DeleteData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c IncomeItem) EditData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c IncomeItem) NewData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c IncomeItem) GetData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c IncomeItem) CopyData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c IncomeItem) GiveUpData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c IncomeItem) RefreshData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c IncomeItem) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

