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
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c SystemParameter) DeleteData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c SystemParameter) EditData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c SystemParameter) NewData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c SystemParameter) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c SystemParameter) CopyData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c SystemParameter) GiveUpData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c SystemParameter) RefreshData() revel.Result {
	c.actionSupport = SystemParameterSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c SystemParameter) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}
