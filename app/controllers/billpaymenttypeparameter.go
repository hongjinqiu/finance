package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type BillPaymentTypeParameterSupport struct {
	ActionSupport
}

type BillPaymentTypeParameter struct {
	BaseDataAction
}

func (c BillPaymentTypeParameter) SaveData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c BillPaymentTypeParameter) DeleteData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BillPaymentTypeParameter) EditData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BillPaymentTypeParameter) NewData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BillPaymentTypeParameter) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c BillPaymentTypeParameter) CopyData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillPaymentTypeParameter) GiveUpData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c BillPaymentTypeParameter) RefreshData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BillPaymentTypeParameter) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}
