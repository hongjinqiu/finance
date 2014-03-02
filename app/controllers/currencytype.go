package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type CurrencyTypeSupport struct {
	ActionSupport
}

type CurrencyType struct {
	BaseDataAction
}

func (c CurrencyType) SaveData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c CurrencyType) DeleteData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c CurrencyType) EditData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c CurrencyType) NewData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c CurrencyType) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c CurrencyType) CopyData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CurrencyType) GiveUpData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c CurrencyType) RefreshData() revel.Result {
	c.actionSupport = CurrencyTypeSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c CurrencyType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

