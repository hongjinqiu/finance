package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type ProviderTypeSupport struct {
	ActionSupport
}

type ProviderType struct {
	BaseDataAction
}

func (c ProviderType) SaveData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ProviderType) DeleteData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ProviderType) EditData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ProviderType) NewData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ProviderType) GetData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c ProviderType) CopyData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ProviderType) GiveUpData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c ProviderType) RefreshData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ProviderType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

