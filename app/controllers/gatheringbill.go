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
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c GatheringBill) DeleteData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c GatheringBill) EditData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c GatheringBill) NewData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c GatheringBill) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c GatheringBill) CopyData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c GatheringBill) GiveUpData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c GatheringBill) RefreshData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
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
	bo, dataSource := c.cancelDataCommon()

	return c.renderCommon(bo, dataSource)
}

/**
 * 反作废
 */
func (c GatheringBill) UnCancelData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, dataSource := c.unCancelDataCommon()

	return c.renderCommon(bo, dataSource)
}

