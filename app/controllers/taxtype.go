package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type TaxTypeSupport struct {
	ActionSupport
}

type TaxType struct {
	BaseDataAction
}

func (c TaxType) SaveData() revel.Result {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c TaxType) DeleteData() revel.Result {
	c.RActionSupport = TaxTypeSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c TaxType) EditData() revel.Result {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c TaxType) NewData() revel.Result {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c TaxType) GetData() revel.Result {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c TaxType) CopyData() revel.Result {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c TaxType) GiveUpData() revel.Result {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c TaxType) RefreshData() revel.Result {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c TaxType) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
