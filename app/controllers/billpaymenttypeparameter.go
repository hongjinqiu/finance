package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
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
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BillPaymentTypeParameter) DeleteData() revel.Result {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BillPaymentTypeParameter) EditData() revel.Result {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BillPaymentTypeParameter) NewData() revel.Result {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BillPaymentTypeParameter) GetData() revel.Result {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BillPaymentTypeParameter) CopyData() revel.Result {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillPaymentTypeParameter) GiveUpData() revel.Result {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BillPaymentTypeParameter) RefreshData() revel.Result {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BillPaymentTypeParameter) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
