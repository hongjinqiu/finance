package controllers

import "github.com/robfig/revel"
import (
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
	c.actionSupport = ArticleTypeSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ArticleType) DeleteData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ArticleType) EditData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ArticleType) NewData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ArticleType) GetData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c ArticleType) CopyData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ArticleType) GiveUpData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c ArticleType) RefreshData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ArticleType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

