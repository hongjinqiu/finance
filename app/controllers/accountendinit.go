package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type AccountEndInitSupport struct {
	ActionSupport
}

type AccountEndInit struct {
	BaseDataAction
}

func (c AccountEndInit) SaveData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountEndInit) DeleteData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountEndInit) EditData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountEndInit) NewData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountEndInit) GetData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c AccountEndInit) CopyData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountEndInit) GiveUpData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountEndInit) RefreshData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountEndInit) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

