package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type AccountInOutSupport struct {
	ActionSupport
}

type AccountInOut struct {
	BaseDataAction
}

func (c AccountInOut) SaveData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c AccountInOut) DeleteData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountInOut) EditData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountInOut) NewData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountInOut) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c AccountInOut) CopyData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountInOut) GiveUpData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c AccountInOut) RefreshData() revel.Result {
	c.actionSupport = AccountInOutSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountInOut) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

