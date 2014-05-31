package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type BankSupport struct {
	ActionSupport
}

type Bank struct {
	BaseDataAction
}

func (c Bank) SaveData() revel.Result {
	c.actionSupport = BankSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Bank) DeleteData() revel.Result {
	c.actionSupport = BankSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Bank) EditData() revel.Result {
	c.actionSupport = BankSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Bank) NewData() revel.Result {
	c.actionSupport = BankSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Bank) GetData() revel.Result {
	c.actionSupport = BankSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c Bank) CopyData() revel.Result {
	c.actionSupport = BankSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Bank) GiveUpData() revel.Result {
	c.actionSupport = BankSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c Bank) RefreshData() revel.Result {
	c.actionSupport = BankSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Bank) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

