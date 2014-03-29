package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type BankDepositReportSupport struct {
	ActionSupport
}

type BankDepositReport struct {
	BaseDataAction
}

func (c BankDepositReport) SaveData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	bo, relationBo, dataSource := c.saveCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BankDepositReport) DeleteData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	
	bo, relationBo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BankDepositReport) EditData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	bo, relationBo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BankDepositReport) NewData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	bo, relationBo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BankDepositReport) GetData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	bo, relationBo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 复制
 */
func (c BankDepositReport) CopyData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	bo, relationBo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BankDepositReport) GiveUpData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	bo, relationBo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 刷新
 */
func (c BankDepositReport) RefreshData() revel.Result {
	c.actionSupport = BankDepositReportSupport{}
	bo, relationBo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BankDepositReport) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

