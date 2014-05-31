package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type ProviderSupport struct {
	ActionSupport
}

type Provider struct {
	BaseDataAction
}

func (c Provider) SaveData() revel.Result {
	c.actionSupport = ProviderSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Provider) DeleteData() revel.Result {
	c.actionSupport = ProviderSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Provider) EditData() revel.Result {
	c.actionSupport = ProviderSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Provider) NewData() revel.Result {
	c.actionSupport = ProviderSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Provider) GetData() revel.Result {
	c.actionSupport = ProviderSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c Provider) CopyData() revel.Result {
	c.actionSupport = ProviderSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Provider) GiveUpData() revel.Result {
	c.actionSupport = ProviderSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c Provider) RefreshData() revel.Result {
	c.actionSupport = ProviderSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c Provider) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

