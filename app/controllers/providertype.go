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
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c ProviderType) DeleteData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c ProviderType) EditData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c ProviderType) NewData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c ProviderType) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c ProviderType) CopyData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ProviderType) GiveUpData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c ProviderType) RefreshData() revel.Result {
	c.actionSupport = ProviderTypeSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
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

