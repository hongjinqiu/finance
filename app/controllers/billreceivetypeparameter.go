package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type BillReceiveTypeParameterSupport struct {
	ActionSupport
}

type BillReceiveTypeParameter struct {
	*revel.Controller
	BaseDataAction
}

func (c BillReceiveTypeParameter) SaveData() revel.Result {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BillReceiveTypeParameter) DeleteData() revel.Result {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BillReceiveTypeParameter) EditData() revel.Result {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BillReceiveTypeParameter) NewData() revel.Result {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BillReceiveTypeParameter) GetData() revel.Result {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BillReceiveTypeParameter) CopyData() revel.Result {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillReceiveTypeParameter) GiveUpData() revel.Result {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BillReceiveTypeParameter) RefreshData() revel.Result {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BillReceiveTypeParameter) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
