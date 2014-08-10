package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type ProviderTypeSupport struct {
	ActionSupport
}

type ProviderType struct {
	BaseDataAction
}

func (c ProviderType) SaveData() revel.Result {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ProviderType) DeleteData() revel.Result {
	c.RActionSupport = ProviderTypeSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ProviderType) EditData() revel.Result {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ProviderType) NewData() revel.Result {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ProviderType) GetData() revel.Result {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c ProviderType) CopyData() revel.Result {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ProviderType) GiveUpData() revel.Result {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c ProviderType) RefreshData() revel.Result {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ProviderType) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

