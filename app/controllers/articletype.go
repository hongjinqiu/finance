package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type ArticleTypeSupport struct {
	ActionSupport
}

type ArticleType struct {
	BaseDataAction
}

func (c ArticleType) SaveData() revel.Result {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ArticleType) DeleteData() revel.Result {
	c.RActionSupport = ArticleTypeSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ArticleType) EditData() revel.Result {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ArticleType) NewData() revel.Result {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ArticleType) GetData() revel.Result {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c ArticleType) CopyData() revel.Result {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ArticleType) GiveUpData() revel.Result {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c ArticleType) RefreshData() revel.Result {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ArticleType) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

