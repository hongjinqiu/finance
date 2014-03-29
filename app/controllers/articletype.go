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
	bo, relationBo, dataSource := c.saveCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c ArticleType) DeleteData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	
	bo, relationBo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c ArticleType) EditData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	bo, relationBo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c ArticleType) NewData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	bo, relationBo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c ArticleType) GetData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	bo, relationBo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 复制
 */
func (c ArticleType) CopyData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	bo, relationBo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ArticleType) GiveUpData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	bo, relationBo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 刷新
 */
func (c ArticleType) RefreshData() revel.Result {
	c.actionSupport = ArticleTypeSupport{}
	bo, relationBo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c ArticleType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

