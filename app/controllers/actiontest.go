package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/model"
	. "com/papersns/model/handler"
	"strings"
)

func init() {
}

type ActionTestSupport struct {
	ActionSupport
}

func (c ActionTestSupport) beforeSaveData(sessionId int, dataSource DataSource, bo *map[string]interface{}) {
	println("ActionTestSupport beforeSaveData")
}

func (c ActionTestSupport) afterSaveData(sessionId int, dataSource DataSource, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	println("ActionTestSupport afterSaveData")
}


type ActionTest struct {
	BillAction
}

func (c ActionTest) SaveData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ActionTest) DeleteData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ActionTest) EditData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ActionTest) NewData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ActionTest) GetData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c ActionTest) CopyData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ActionTest) GiveUpData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c ActionTest) RefreshData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ActionTest) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

func (c ActionTest) CancelData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	modelRenderVO := c.cancelDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c ActionTest) UnCancelData() revel.Result {
	c.actionSupport = ActionSupport{}
	modelRenderVO := c.unCancelDataCommon()
	return c.renderCommon(modelRenderVO)
}
