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
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c Provider) DeleteData() revel.Result {
	c.actionSupport = ProviderSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c Provider) EditData() revel.Result {
	c.actionSupport = ProviderSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c Provider) NewData() revel.Result {
	c.actionSupport = ProviderSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c Provider) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c Provider) CopyData() revel.Result {
	c.actionSupport = ProviderSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Provider) GiveUpData() revel.Result {
	c.actionSupport = ProviderSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c Provider) RefreshData() revel.Result {
	c.actionSupport = ProviderSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
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

