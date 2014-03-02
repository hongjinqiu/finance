package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type IncomeItemSupport struct {
	ActionSupport
}

type IncomeItem struct {
	BaseDataAction
}

func (c IncomeItem) SaveData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c IncomeItem) DeleteData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c IncomeItem) EditData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c IncomeItem) NewData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c IncomeItem) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c IncomeItem) CopyData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c IncomeItem) GiveUpData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c IncomeItem) RefreshData() revel.Result {
	c.actionSupport = IncomeItemSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c IncomeItem) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

