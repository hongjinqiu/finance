package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type BalanceTypeSupport struct {
	ActionSupport
}

type BalanceType struct {
	BaseDataAction
}

func (c BalanceType) SaveData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c BalanceType) DeleteData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BalanceType) EditData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BalanceType) NewData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BalanceType) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c BalanceType) CopyData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BalanceType) GiveUpData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c BalanceType) RefreshData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BalanceType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

