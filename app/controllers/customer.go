package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type CustomerSupport struct {
	ActionSupport
}

type Customer struct {
	*revel.Controller
	BaseDataAction
}

func (c Customer) SaveData() revel.Result {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c Customer) DeleteData() revel.Result {
	c.RActionSupport = CustomerSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c Customer) EditData() revel.Result {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c Customer) NewData() revel.Result {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c Customer) GetData() revel.Result {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c Customer) CopyData() revel.Result {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Customer) GiveUpData() revel.Result {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c Customer) RefreshData() revel.Result {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c Customer) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

