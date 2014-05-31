package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type AccountInOutSupport struct {
	ActionSupport
}

type AccountInOut struct {
	BaseDataAction
}

func (c AccountInOut) SaveData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInOut) DeleteData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInOut) EditData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInOut) NewData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInOut) GetData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c AccountInOut) CopyData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountInOut) GiveUpData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountInOut) RefreshData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInOut) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

