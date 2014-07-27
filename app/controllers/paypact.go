package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type PayPactSupport struct {
	ActionSupport
}

type PayPact struct {
	BaseDataAction
}

func (c PayPact) SaveData() revel.Result {
	c.actionSupport = PayPactSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PayPact) DeleteData() revel.Result {
	c.actionSupport = PayPactSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PayPact) EditData() revel.Result {
	c.actionSupport = PayPactSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PayPact) NewData() revel.Result {
	c.actionSupport = PayPactSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PayPact) GetData() revel.Result {
	c.actionSupport = PayPactSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c PayPact) CopyData() revel.Result {
	c.actionSupport = PayPactSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c PayPact) GiveUpData() revel.Result {
	c.actionSupport = PayPactSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c PayPact) RefreshData() revel.Result {
	c.actionSupport = PayPactSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PayPact) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

