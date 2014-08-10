package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type PubReferenceLogSupport struct {
	ActionSupport
}

type PubReferenceLog struct {
	BaseDataAction
}

func (c PubReferenceLog) SaveData() revel.Result {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PubReferenceLog) DeleteData() revel.Result {
	c.RActionSupport = PubReferenceLogSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PubReferenceLog) EditData() revel.Result {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PubReferenceLog) NewData() revel.Result {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PubReferenceLog) GetData() revel.Result {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c PubReferenceLog) CopyData() revel.Result {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c PubReferenceLog) GiveUpData() revel.Result {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c PubReferenceLog) RefreshData() revel.Result {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PubReferenceLog) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

