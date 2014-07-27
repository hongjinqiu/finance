package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type CustomerTypeSupport struct {
	ActionSupport
}

type CustomerType struct {
	BaseDataAction
}

func (c CustomerType) SaveData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CustomerType) DeleteData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CustomerType) EditData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CustomerType) NewData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CustomerType) GetData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c CustomerType) CopyData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CustomerType) GiveUpData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c CustomerType) RefreshData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c CustomerType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

