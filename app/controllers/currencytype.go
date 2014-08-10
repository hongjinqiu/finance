package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type CurrencyTypeSupport struct {
	ActionSupport
}

type CurrencyType struct {
	BaseDataAction
}

func (c CurrencyType) SaveData() revel.Result {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CurrencyType) DeleteData() revel.Result {
	c.RActionSupport = CurrencyTypeSupport{}

	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CurrencyType) EditData() revel.Result {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CurrencyType) NewData() revel.Result {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CurrencyType) GetData() revel.Result {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c CurrencyType) CopyData() revel.Result {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CurrencyType) GiveUpData() revel.Result {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c CurrencyType) RefreshData() revel.Result {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c CurrencyType) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
