package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	. "com/papersns/component"
	. "com/papersns/model"
	. "com/papersns/model/handler"
	"strings"
)

func init() {
}

type ActionTestSupport struct {
	ActionSupport
}

func (c ActionTestSupport) RBeforeSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	println("ActionTestSupport RBeforeSaveData")
}

func (c ActionTestSupport) RAfterSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	println("ActionTestSupport RAfterSaveData")
}

type ActionTest struct {
	*revel.Controller
	BillAction
}

func (c ActionTest) SaveData() revel.Result {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ActionTest) DeleteData() revel.Result {
	c.RActionSupport = ActionTestSupport{}

	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ActionTest) EditData() revel.Result {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ActionTest) NewData() revel.Result {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ActionTest) GetData() revel.Result {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c ActionTest) CopyData() revel.Result {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ActionTest) GiveUpData() revel.Result {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c ActionTest) RefreshData() revel.Result {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ActionTest) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

func (c ActionTest) CancelData() revel.Result {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.RCancelDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c ActionTest) UnCancelData() revel.Result {
	c.RActionSupport = ActionSupport{}
	modelRenderVO := c.RUnCancelDataCommon()
	return c.RRenderCommon(modelRenderVO)
}
