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
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c MeasureUnit) DeleteData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c MeasureUnit) EditData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c MeasureUnit) NewData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c MeasureUnit) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c MeasureUnit) CopyData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c MeasureUnit) GiveUpData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c MeasureUnit) RefreshData() revel.Result {
	c.actionSupport = MeasureUnitSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c MeasureUnit) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

