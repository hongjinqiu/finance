package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type AccountInOutDisplaySupport struct {
	ActionSupport
}

type AccountInOutDisplay struct {
	BaseDataAction
}

func (c AccountInOutDisplay) SaveData() revel.Result {
	c.actionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInOutDisplay) DeleteData() revel.Result {
	c.actionSupport = AccountInOutDisplaySupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInOutDisplay) EditData() revel.Result {
	c.actionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInOutDisplay) NewData() revel.Result {
	c.actionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInOutDisplay) GetData() revel.Result {
	c.actionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c AccountInOutDisplay) CopyData() revel.Result {
	c.actionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountInOutDisplay) GiveUpData() revel.Result {
	c.actionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountInOutDisplay) RefreshData() revel.Result {
	c.actionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInOutDisplay) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}
