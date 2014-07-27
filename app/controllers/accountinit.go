package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type AccountInitSupport struct {
	ActionSupport
}

type AccountInit struct {
	BaseDataAction
}

func (c AccountInit) SaveData() revel.Result {
	c.actionSupport = AccountInitSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInit) DeleteData() revel.Result {
	c.actionSupport = AccountInitSupport{}

	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInit) EditData() revel.Result {
	c.actionSupport = AccountInitSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInit) NewData() revel.Result {
	c.actionSupport = AccountInitSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInit) GetData() revel.Result {
	c.actionSupport = AccountInitSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c AccountInit) CopyData() revel.Result {
	c.actionSupport = AccountInitSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountInit) GiveUpData() revel.Result {
	c.actionSupport = AccountInitSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountInit) RefreshData() revel.Result {
	c.actionSupport = AccountInitSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountInit) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
