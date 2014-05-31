package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type AccountUnInitSupport struct {
	ActionSupport
}

type AccountUnInit struct {
	BaseDataAction
}

func (c AccountUnInit) SaveData() revel.Result {
	c.actionSupport = AccountUnInitSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountUnInit) DeleteData() revel.Result {
	c.actionSupport = AccountUnInitSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountUnInit) EditData() revel.Result {
	c.actionSupport = AccountUnInitSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountUnInit) NewData() revel.Result {
	c.actionSupport = AccountUnInitSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountUnInit) GetData() revel.Result {
	c.actionSupport = AccountUnInitSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c AccountUnInit) CopyData() revel.Result {
	c.actionSupport = AccountUnInitSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountUnInit) GiveUpData() revel.Result {
	c.actionSupport = AccountUnInitSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountUnInit) RefreshData() revel.Result {
	c.actionSupport = AccountUnInitSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountUnInit) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

