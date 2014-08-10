package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type ProviderSupport struct {
	ActionSupport
}

type Provider struct {
	BaseDataAction
}

func (c Provider) SaveData() revel.Result {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c Provider) DeleteData() revel.Result {
	c.RActionSupport = ProviderSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c Provider) EditData() revel.Result {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c Provider) NewData() revel.Result {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c Provider) GetData() revel.Result {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c Provider) CopyData() revel.Result {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Provider) GiveUpData() revel.Result {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c Provider) RefreshData() revel.Result {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c Provider) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

