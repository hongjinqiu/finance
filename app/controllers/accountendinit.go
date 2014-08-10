package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type AccountEndInitSupport struct {
	ActionSupport
}

type AccountEndInit struct {
	BaseDataAction
}

func (c AccountEndInit) SaveData() revel.Result {
	c.RActionSupport = AccountEndInitSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountEndInit) DeleteData() revel.Result {
	c.RActionSupport = AccountEndInitSupport{}

	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountEndInit) EditData() revel.Result {
	c.RActionSupport = AccountEndInitSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountEndInit) NewData() revel.Result {
	c.RActionSupport = AccountEndInitSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountEndInit) GetData() revel.Result {
	c.RActionSupport = AccountEndInitSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c AccountEndInit) CopyData() revel.Result {
	c.RActionSupport = AccountEndInitSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountEndInit) GiveUpData() revel.Result {
	c.RActionSupport = AccountEndInitSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountEndInit) RefreshData() revel.Result {
	c.RActionSupport = AccountEndInitSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountEndInit) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
