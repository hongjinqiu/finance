package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type IncomeTypeSupport struct {
	ActionSupport
}

type IncomeType struct {
	*revel.Controller
	BaseDataAction
}

func (c IncomeType) SaveData() revel.Result {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c IncomeType) DeleteData() revel.Result {
	c.RActionSupport = IncomeTypeSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c IncomeType) EditData() revel.Result {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c IncomeType) NewData() revel.Result {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c IncomeType) GetData() revel.Result {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c IncomeType) CopyData() revel.Result {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c IncomeType) GiveUpData() revel.Result {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c IncomeType) RefreshData() revel.Result {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c IncomeType) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

