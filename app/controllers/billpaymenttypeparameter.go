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
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillPaymentTypeParameter) DeleteData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillPaymentTypeParameter) EditData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillPaymentTypeParameter) NewData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillPaymentTypeParameter) GetData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BillPaymentTypeParameter) CopyData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillPaymentTypeParameter) GiveUpData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BillPaymentTypeParameter) RefreshData() revel.Result {
	c.actionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillPaymentTypeParameter) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
