package controllers

import "github.com/robfig/revel"
import (
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
	c.actionSupport = TaxTypeSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c TaxType) DeleteData() revel.Result {
	c.actionSupport = TaxTypeSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c TaxType) EditData() revel.Result {
	c.actionSupport = TaxTypeSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c TaxType) NewData() revel.Result {
	c.actionSupport = TaxTypeSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c TaxType) GetData() revel.Result {
	c.actionSupport = TaxTypeSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c TaxType) CopyData() revel.Result {
	c.actionSupport = TaxTypeSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c TaxType) GiveUpData() revel.Result {
	c.actionSupport = TaxTypeSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c TaxType) RefreshData() revel.Result {
	c.actionSupport = TaxTypeSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c TaxType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}
