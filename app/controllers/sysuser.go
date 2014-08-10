package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type SysUserSupport struct {
	ActionSupport
}

type SysUser struct {
	*revel.Controller
	BaseDataAction
}

func (c SysUser) SaveData() revel.Result {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c SysUser) DeleteData() revel.Result {
	c.RActionSupport = SysUserSupport{}

	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c SysUser) EditData() revel.Result {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c SysUser) NewData() revel.Result {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c SysUser) GetData() revel.Result {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c SysUser) CopyData() revel.Result {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c SysUser) GiveUpData() revel.Result {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c SysUser) RefreshData() revel.Result {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c SysUser) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
