package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type ArticleSupport struct {
	ActionSupport
}

type Article struct {
	BaseDataAction
}

func (c Article) SaveData() revel.Result {
	c.actionSupport = ArticleSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Article) DeleteData() revel.Result {
	c.actionSupport = ArticleSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Article) EditData() revel.Result {
	c.actionSupport = ArticleSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Article) NewData() revel.Result {
	c.actionSupport = ArticleSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Article) GetData() revel.Result {
	c.actionSupport = ArticleSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c Article) CopyData() revel.Result {
	c.actionSupport = ArticleSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Article) GiveUpData() revel.Result {
	c.actionSupport = ArticleSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c Article) RefreshData() revel.Result {
	c.actionSupport = ArticleSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Article) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

