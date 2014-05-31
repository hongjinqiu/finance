package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type GatheringBillSupport struct {
	ActionSupport
}

type GatheringBill struct {
	BillAction
}

func (c GatheringBill) SaveData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c GatheringBill) DeleteData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c GatheringBill) EditData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c GatheringBill) NewData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c GatheringBill) GetData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c GatheringBill) CopyData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c GatheringBill) GiveUpData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c GatheringBill) RefreshData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c GatheringBill) LogList() revel.Result {
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
func (c GatheringBill) CancelData() revel.Result {
	c.actionSupport = ActionSupport{}
	modelRenderVO := c.cancelDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 反作废
 */
func (c GatheringBill) UnCancelData() revel.Result {
	c.actionSupport = ActionSupport{}
	modelRenderVO := c.unCancelDataCommon()
	return c.renderCommon(modelRenderVO)
}

