package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type BankAccountInitSupport struct {
	ActionSupport
}

type BankAccountInit struct {
	BaseDataAction
}

func (c BankAccountInit) SaveData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c BankAccountInit) DeleteData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BankAccountInit) EditData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BankAccountInit) NewData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BankAccountInit) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c BankAccountInit) CopyData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BankAccountInit) GiveUpData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c BankAccountInit) RefreshData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BankAccountInit) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

