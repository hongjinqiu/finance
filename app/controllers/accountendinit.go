package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type AccountEndInitSupport struct {
	ActionSupport
}

type AccountEndInit struct {
	BaseDataAction
}

func (c AccountEndInit) SaveData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c AccountEndInit) DeleteData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountEndInit) EditData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountEndInit) NewData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountEndInit) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c AccountEndInit) CopyData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountEndInit) GiveUpData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c AccountEndInit) RefreshData() revel.Result {
	c.actionSupport = AccountEndInitSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountEndInit) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

