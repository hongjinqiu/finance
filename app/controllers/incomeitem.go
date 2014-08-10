package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
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
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c IncomeItem) DeleteData() revel.Result {
	c.RActionSupport = IncomeItemSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c IncomeItem) EditData() revel.Result {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c IncomeItem) NewData() revel.Result {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c IncomeItem) GetData() revel.Result {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c IncomeItem) CopyData() revel.Result {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c IncomeItem) GiveUpData() revel.Result {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c IncomeItem) RefreshData() revel.Result {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c IncomeItem) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

