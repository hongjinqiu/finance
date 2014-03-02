package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type IncomeTypeSupport struct {
	ActionSupport
}

type IncomeType struct {
	BaseDataAction
}

func (c IncomeType) SaveData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c IncomeType) DeleteData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c IncomeType) EditData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c IncomeType) NewData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c IncomeType) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c IncomeType) CopyData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c IncomeType) GiveUpData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c IncomeType) RefreshData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c IncomeType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

