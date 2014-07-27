package controllers

import "github.com/robfig/revel"
import (
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
	c.actionSupport = MeasureUnitSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c MeasureUnit) DeleteData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c MeasureUnit) EditData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c MeasureUnit) NewData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c MeasureUnit) GetData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c MeasureUnit) CopyData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c MeasureUnit) GiveUpData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c MeasureUnit) RefreshData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c MeasureUnit) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

