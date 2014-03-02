package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type AccountingPeriodSupport struct {
	ActionSupport
}

type AccountingPeriod struct {
	BaseDataAction
}

func (c AccountingPeriod) SaveData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c AccountingPeriod) DeleteData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountingPeriod) EditData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountingPeriod) NewData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountingPeriod) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c AccountingPeriod) CopyData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountingPeriod) GiveUpData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c AccountingPeriod) RefreshData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c AccountingPeriod) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

