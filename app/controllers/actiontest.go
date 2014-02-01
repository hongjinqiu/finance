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
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c ActionTest) DeleteData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c ActionTest) EditData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c ActionTest) NewData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c ActionTest) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 复制
 */
func (c ActionTest) CopyData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ActionTest) GiveUpData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

/**
 * 刷新
 */
func (c ActionTest) RefreshData() revel.Result {
	c.actionSupport = ActionTestSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
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
	bo, dataSource := c.cancelDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c ActionTest) UnCancelData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, dataSource := c.unCancelDataCommon()
	
	return c.renderCommon(bo, dataSource)
}
