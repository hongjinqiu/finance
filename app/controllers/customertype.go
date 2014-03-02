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
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c CustomerType) DeleteData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c CustomerType) EditData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c CustomerType) NewData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c CustomerType) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c CustomerType) CopyData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CustomerType) GiveUpData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c CustomerType) RefreshData() revel.Result {
	c.actionSupport = CustomerTypeSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c CustomerType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

