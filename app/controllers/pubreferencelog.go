package controllers

import "github.com/robfig/revel"
import (
	"strings"
)

func init() {
}

type PubReferenceLogSupport struct {
	ActionSupport
}

type PubReferenceLog struct {
	BaseDataAction
}

func (c PubReferenceLog) SaveData() revel.Result {
	c.actionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PubReferenceLog) DeleteData() revel.Result {
	c.actionSupport = PubReferenceLogSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PubReferenceLog) EditData() revel.Result {
	c.actionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PubReferenceLog) NewData() revel.Result {
	c.actionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PubReferenceLog) GetData() revel.Result {
	c.actionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c PubReferenceLog) CopyData() revel.Result {
	c.actionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c PubReferenceLog) GiveUpData() revel.Result {
	c.actionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c PubReferenceLog) RefreshData() revel.Result {
	c.actionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c PubReferenceLog) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

