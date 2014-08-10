package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type CustomerTypeSupport struct {
	ActionSupport
}

type CustomerType struct {
	BaseDataAction
}

func (c CustomerType) SaveData() revel.Result {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CustomerType) DeleteData() revel.Result {
	c.RActionSupport = CustomerTypeSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CustomerType) EditData() revel.Result {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CustomerType) NewData() revel.Result {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CustomerType) GetData() revel.Result {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c CustomerType) CopyData() revel.Result {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CustomerType) GiveUpData() revel.Result {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c CustomerType) RefreshData() revel.Result {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CustomerType) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

