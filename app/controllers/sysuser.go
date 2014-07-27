package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type SysUserSupport struct {
	ActionSupport
}

type SysUser struct {
	BaseDataAction
}

func (c SysUser) SaveData() revel.Result {
	c.actionSupport = SysUserSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c SysUser) DeleteData() revel.Result {
	c.actionSupport = SysUserSupport{}

	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c SysUser) EditData() revel.Result {
	c.actionSupport = SysUserSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c SysUser) NewData() revel.Result {
	c.actionSupport = SysUserSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c SysUser) GetData() revel.Result {
	c.actionSupport = SysUserSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c SysUser) CopyData() revel.Result {
	c.actionSupport = SysUserSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c SysUser) GiveUpData() revel.Result {
	c.actionSupport = SysUserSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c SysUser) RefreshData() revel.Result {
	c.actionSupport = SysUserSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c SysUser) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
