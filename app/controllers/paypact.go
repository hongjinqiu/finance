package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type PayPactSupport struct {
	ActionSupport
}

type PayPact struct {
	BaseDataAction
}

func (c PayPact) SaveData() revel.Result {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PayPact) DeleteData() revel.Result {
	c.RActionSupport = PayPactSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PayPact) EditData() revel.Result {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PayPact) NewData() revel.Result {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PayPact) GetData() revel.Result {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c PayPact) CopyData() revel.Result {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c PayPact) GiveUpData() revel.Result {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c PayPact) RefreshData() revel.Result {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PayPact) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

