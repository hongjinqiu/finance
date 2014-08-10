package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type AccountInitSupport struct {
	ActionSupport
}

type AccountInit struct {
	*revel.Controller
	BaseDataAction
}

func (c AccountInit) SaveData() revel.Result {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountInit) DeleteData() revel.Result {
	c.RActionSupport = AccountInitSupport{}

	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountInit) EditData() revel.Result {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountInit) NewData() revel.Result {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountInit) GetData() revel.Result {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c AccountInit) CopyData() revel.Result {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountInit) GiveUpData() revel.Result {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountInit) RefreshData() revel.Result {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountInit) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
