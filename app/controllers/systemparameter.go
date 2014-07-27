package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type SystemParameterSupport struct {
	ActionSupport
}

type SystemParameter struct {
	BaseDataAction
}

func (c SystemParameter) SaveData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c SystemParameter) DeleteData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c SystemParameter) EditData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c SystemParameter) NewData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c SystemParameter) GetData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c SystemParameter) CopyData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c SystemParameter) GiveUpData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c SystemParameter) RefreshData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c SystemParameter) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
