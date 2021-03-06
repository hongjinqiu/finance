package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
)

func init() {
}

type BalanceTypeSupport struct {
	ActionSupport
}

type BalanceType struct {
	*revel.Controller
	BaseDataAction
}

func (c BalanceType) SaveData() revel.Result {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BalanceType) DeleteData() revel.Result {
	c.RActionSupport = BalanceTypeSupport{}
	
	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BalanceType) EditData() revel.Result {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BalanceType) NewData() revel.Result {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BalanceType) GetData() revel.Result {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BalanceType) CopyData() revel.Result {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BalanceType) GiveUpData() revel.Result {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BalanceType) RefreshData() revel.Result {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BalanceType) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

