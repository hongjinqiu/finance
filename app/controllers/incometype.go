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
	bo, relationBo, dataSource := c.saveCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c IncomeType) DeleteData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	
	bo, relationBo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c IncomeType) EditData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, relationBo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c IncomeType) NewData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, relationBo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c IncomeType) GetData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, relationBo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 复制
 */
func (c IncomeType) CopyData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, relationBo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c IncomeType) GiveUpData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, relationBo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 刷新
 */
func (c IncomeType) RefreshData() revel.Result {
	c.actionSupport = IncomeTypeSupport{}
	bo, relationBo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c IncomeType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

