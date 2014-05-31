package controllers

import "github.com/robfig/revel"
import (
	"strings"
	"time"
	"fmt"
	"strconv"
	. "com/papersns/model"
)

func init() {
}

type AccountingPeriodSupport struct {
	ActionSupport
}

func (o AccountingPeriodSupport) afterNewData(sessionId int, dataSource DataSource, bo *map[string]interface{}) {
	masterData := (*bo)["A"].(map[string]interface{})

	year := time.Now().Year()
	masterData["accountingYear"] = year
	
	(*bo)["A"] = masterData
	
	numAccountingPeriod, err := strconv.Atoi(fmt.Sprint(masterData["numAccountingPeriod"]))
	if err != nil {
		panic(err)
	}
	detailDataLi := []interface{}{}
	
	modelTemplateFactory := ModelTemplateFactory{}
	dataSetId := "B"
	for i := 0; i < numAccountingPeriod; i++ {
		data := modelTemplateFactory.GetDataSetNewData(dataSource, dataSetId, *bo)
		data["id"] = "afterNewData" + fmt.Sprint(i)
		data["sequenceNo"] = i + 1
		numStr := fmt.Sprint(i + 1)
		if i + 1 < 10 {
			numStr = "0" + numStr
		}
		startDateStr := fmt.Sprint(year) + numStr + "01"
		startDate, err := strconv.Atoi(startDateStr)
		if err != nil {
			panic(err)
		}
		data["startDate"] = startDate
		startTime, err := time.Parse("20060102", startDateStr)
		if err != nil {
			panic(err)
		}
		nextMonthTime := startTime.AddDate(0, 1, -1)
		data["endDate"], err = strconv.Atoi(nextMonthTime.Format("20060102"))
		if err != nil {
			panic(err)
		}
		detailDataLi = append(detailDataLi, data)
	}
	
	(*bo)["B"] = detailDataLi
}

type AccountingPeriod struct {
	BaseDataAction
}

func (c AccountingPeriod) SaveData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountingPeriod) DeleteData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	
	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountingPeriod) EditData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountingPeriod) NewData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountingPeriod) GetData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c AccountingPeriod) CopyData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountingPeriod) GiveUpData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountingPeriod) RefreshData() revel.Result {
	c.actionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c AccountingPeriod) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

