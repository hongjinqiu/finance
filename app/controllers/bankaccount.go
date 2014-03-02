package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type BankAccountSupport struct {
	ActionSupport
}

type BankAccount struct {
	BaseDataAction
}

func (c BankAccount) SaveData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c BankAccount) DeleteData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BankAccount) EditData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BankAccount) NewData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BankAccount) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c BankAccount) CopyData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BankAccount) GiveUpData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c BankAccount) RefreshData() revel.Result {
	c.actionSupport = BankAccountSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BankAccount) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

