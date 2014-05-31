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
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PayBill) DeleteData() revel.Result {
	c.actionSupport = PayBillSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PayBill) EditData() revel.Result {
	c.actionSupport = PayBillSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PayBill) NewData() revel.Result {
	c.actionSupport = PayBillSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PayBill) GetData() revel.Result {
	c.actionSupport = PayBillSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c PayBill) CopyData() revel.Result {
	c.actionSupport = PayBillSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c PayBill) GiveUpData() revel.Result {
	c.actionSupport = PayBillSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c PayBill) RefreshData() revel.Result {
	c.actionSupport = PayBillSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
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
	modelRenderVO := c.cancelDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 反作废
 */
func (c PayBill) UnCancelData() revel.Result {
	c.actionSupport = ActionSupport{}
	modelRenderVO := c.unCancelDataCommon()
	return c.renderCommon(modelRenderVO)
}

