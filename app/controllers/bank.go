package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type BankSupport struct {
	ActionSupport
}

type Bank struct {
	BaseDataAction
}

func (c Bank) SaveData() revel.Result {
	c.actionSupport = BankSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c Bank) DeleteData() revel.Result {
	c.actionSupport = BankSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c Bank) EditData() revel.Result {
	c.actionSupport = BankSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c Bank) NewData() revel.Result {
	c.actionSupport = BankSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c Bank) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c Bank) CopyData() revel.Result {
	c.actionSupport = BankSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Bank) GiveUpData() revel.Result {
	c.actionSupport = BankSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c Bank) RefreshData() revel.Result {
	c.actionSupport = BankSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c Bank) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}
