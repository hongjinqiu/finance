package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type AccountUnInitSupport struct {
	ActionSupport
}

type AccountUnInit struct {
	BaseDataAction
}

func (c AccountUnInit) SaveData() revel.Result {
	c.RActionSupport = AccountUnInitSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountUnInit) DeleteData() revel.Result {
	c.RActionSupport = AccountUnInitSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountUnInit) EditData() revel.Result {
	c.RActionSupport = AccountUnInitSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountUnInit) NewData() revel.Result {
	c.RActionSupport = AccountUnInitSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountUnInit) GetData() revel.Result {
	c.RActionSupport = AccountUnInitSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c AccountUnInit) CopyData() revel.Result {
	c.RActionSupport = AccountUnInitSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountUnInit) GiveUpData() revel.Result {
	c.RActionSupport = AccountUnInitSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountUnInit) RefreshData() revel.Result {
	c.RActionSupport = AccountUnInitSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountUnInit) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

