package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type PayBillSupport struct {
	ActionSupport
}

type PayBill struct {
	BillAction
}

func (c PayBill) SaveData() revel.Result {
	c.actionSupport = PayBillSupport{}
	bo, relationBo, dataSource := c.saveCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c PayBill) DeleteData() revel.Result {
	c.actionSupport = PayBillSupport{}
	
	bo, relationBo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c PayBill) EditData() revel.Result {
	c.actionSupport = PayBillSupport{}
	bo, relationBo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c PayBill) NewData() revel.Result {
	c.actionSupport = PayBillSupport{}
	bo, relationBo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c PayBill) GetData() revel.Result {
	c.actionSupport = PayBillSupport{}
	bo, relationBo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 复制
 */
func (c PayBill) CopyData() revel.Result {
	c.actionSupport = PayBillSupport{}
	bo, relationBo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c PayBill) GiveUpData() revel.Result {
	c.actionSupport = PayBillSupport{}
	bo, relationBo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 刷新
 */
func (c PayBill) RefreshData() revel.Result {
	c.actionSupport = PayBillSupport{}
	bo, relationBo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c PayBill) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

/**
 * 作废
 */
func (c PayBill) CancelData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, relationBo, dataSource := c.cancelDataCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 反作废
 */
func (c PayBill) UnCancelData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, relationBo, dataSource := c.unCancelDataCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

