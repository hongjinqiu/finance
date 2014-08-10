package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type MeasureUnitSupport struct {
	ActionSupport
}

type MeasureUnit struct {
	BaseDataAction
}

func (c MeasureUnit) SaveData() revel.Result {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c MeasureUnit) DeleteData() revel.Result {
	c.RActionSupport = MeasureUnitSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c MeasureUnit) EditData() revel.Result {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c MeasureUnit) NewData() revel.Result {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c MeasureUnit) GetData() revel.Result {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c MeasureUnit) CopyData() revel.Result {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c MeasureUnit) GiveUpData() revel.Result {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c MeasureUnit) RefreshData() revel.Result {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c MeasureUnit) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

