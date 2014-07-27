package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type IncomeTypeSupport struct {
	ActionSupport
}

type IncomeType struct {
	BaseDataAction
}

func (c IncomeType) SaveData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c IncomeType) DeleteData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c IncomeType) EditData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c IncomeType) NewData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c IncomeType) GetData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c IncomeType) CopyData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c IncomeType) GiveUpData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c IncomeType) RefreshData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c IncomeType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

