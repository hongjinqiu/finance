package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type BillReceiveTypeParameterSupport struct {
	ActionSupport
}

type BillReceiveTypeParameter struct {
	BaseDataAction
}

func (c BillReceiveTypeParameter) SaveData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillReceiveTypeParameter) DeleteData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillReceiveTypeParameter) EditData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillReceiveTypeParameter) NewData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillReceiveTypeParameter) GetData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BillReceiveTypeParameter) CopyData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillReceiveTypeParameter) GiveUpData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BillReceiveTypeParameter) RefreshData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillReceiveTypeParameter) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}
