package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type SystemParameterSupport struct {
	ActionSupport
}

type SystemParameter struct {
	*revel.Controller
	BaseDataAction
}

func (c SystemParameter) SaveData() revel.Result {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c SystemParameter) DeleteData() revel.Result {
	c.RActionSupport = SystemParameterSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c SystemParameter) EditData() revel.Result {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c SystemParameter) NewData() revel.Result {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c SystemParameter) GetData() revel.Result {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c SystemParameter) CopyData() revel.Result {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c SystemParameter) GiveUpData() revel.Result {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c SystemParameter) RefreshData() revel.Result {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c SystemParameter) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
