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
	bo, relationBo, dataSource := c.saveCommon()

	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BillReceiveTypeParameter) DeleteData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	
	bo, relationBo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BillReceiveTypeParameter) EditData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	bo, relationBo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BillReceiveTypeParameter) NewData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	bo, relationBo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

func (c BillReceiveTypeParameter) GetData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	bo, relationBo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 复制
 */
func (c BillReceiveTypeParameter) CopyData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	bo, relationBo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillReceiveTypeParameter) GiveUpData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	bo, relationBo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
}

/**
 * 刷新
 */
func (c BillReceiveTypeParameter) RefreshData() revel.Result {
	c.actionSupport = BillReceiveTypeParameterSupport{}
	bo, relationBo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, relationBo, dataSource)
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
