package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type BalanceTypeSupport struct {
	ActionSupport
}

type BalanceType struct {
	BaseDataAction
}

func (c BalanceType) SaveData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BalanceType) DeleteData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BalanceType) EditData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BalanceType) NewData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BalanceType) GetData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BalanceType) CopyData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BalanceType) GiveUpData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BalanceType) RefreshData() revel.Result {
	c.actionSupport = BalanceTypeSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BalanceType) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

