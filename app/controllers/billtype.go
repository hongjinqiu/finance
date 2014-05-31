package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type BillTypeSupport struct {
	ActionSupport
}

type BillType struct {
	BaseDataAction
}

func (c BillType) SaveData() revel.Result {
	c.actionSupport = BillTypeSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillType) DeleteData() revel.Result {
	c.actionSupport = BillTypeSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillType) EditData() revel.Result {
	c.actionSupport = BillTypeSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillType) NewData() revel.Result {
	c.actionSupport = BillTypeSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillType) GetData() revel.Result {
	c.actionSupport = BillTypeSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BillType) CopyData() revel.Result {
	c.actionSupport = BillTypeSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillType) GiveUpData() revel.Result {
	c.actionSupport = BillTypeSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BillType) RefreshData() revel.Result {
	c.actionSupport = BillTypeSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BillType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}
