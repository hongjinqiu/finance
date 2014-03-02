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
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c Article) DeleteData() revel.Result {
	c.actionSupport = ArticleSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c Article) EditData() revel.Result {
	c.actionSupport = ArticleSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c Article) NewData() revel.Result {
	c.actionSupport = ArticleSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c Article) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c Article) CopyData() revel.Result {
	c.actionSupport = ArticleSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Article) GiveUpData() revel.Result {
	c.actionSupport = ArticleSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c Article) RefreshData() revel.Result {
	c.actionSupport = ArticleSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c Article) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

