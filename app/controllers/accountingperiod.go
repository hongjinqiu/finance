package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/common"
	. "com/papersns/component"
	"com/papersns/global"
	. "com/papersns/model"
	"fmt"
	"strconv"
	"strings"
	"time"
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
		if i+1 < 10 {
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

func (c AccountingPeriod) renderCommon(modelRenderVO ModelRenderVO) revel.Result {
	bo := modelRenderVO.Bo
	relationBo := modelRenderVO.RelationBo
	dataSource := modelRenderVO.DataSource
	usedCheckBo := modelRenderVO.UsedCheckBo
	// 重新修改usedCheckBo,改为查询单据,只要在会计期内存在单据,则视为被用
	modelTemplateFactory := ModelTemplateFactory{}
	strId := modelTemplateFactory.GetStrId(bo)
	if strId != "" && strId != "0" {
		bDataSetLi := bo["B"].([]interface{})
		firstLineData := bDataSetLi[0].(map[string]interface{})
		lastLineData := bDataSetLi[len(bDataSetLi)-1].(map[string]interface{})

		commonUtil := CommonUtil{}
		firstStartDate := commonUtil.GetIntFromMap(firstLineData, "startDate")
		lastEndDate := commonUtil.GetIntFromMap(lastLineData, "endDate")

		queryMap := map[string]interface{}{
			"A.billDate": map[string]interface{}{
				"$gte": firstStartDate,
				"$lt":  lastEndDate,
			},
		}

		qb := QuerySupport{}
		//GatheringBill,PayBill
		sessionId := global.GetSessionId()
		defer global.CloseSession(sessionId)
		session, _ := global.GetConnection(sessionId)
		dataSourceIdLi := []string{"GatheringBill", "PayBill"}
		for _, dataSourceId := range dataSourceIdLi {
			tmpDataSource := modelTemplateFactory.GetDataSource(dataSourceId)
			collectionName := modelTemplateFactory.GetCollectionName(tmpDataSource)
			// TODO
			//			func (qb QuerySupport) FindByMapWithSession(session *mgo.Session, collection string, query map[string]interface{}) (result map[string]interface{}, found bool) {
			_, found := qb.FindByMapWithSession(session, collectionName, queryMap)
			if found {
				// 主数据集设置被用标记
				if usedCheckBo["A"] == nil {
					usedCheckBo["A"] = map[string]interface{}{}
				}
				masterUsedCheck := usedCheckBo["A"].(map[string]interface{})
				usedCheckBo["A"] = masterUsedCheck
				masterUsedCheck[strId] = true
				
				// 分录数据集设置被用标记
				if usedCheckBo["B"] == nil {
					usedCheckBo["B"] = map[string]interface{}{}
				}
				detailUsedCheck := usedCheckBo["B"].(map[string]interface{})
				usedCheckBo["B"] = usedCheckBo["B"]
				for _, detailData := range bDataSetLi {
					detailDataDict := detailData.(map[string]interface{})
					detailUsedCheck[fmt.Sprint(detailDataDict["id"])] = true
				}
				break
			}
		}
	}

	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateAllFieldBo(dataSource, &bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		if (*data)[fieldGroup.Id] != nil {
			(*data)[fieldGroup.Id] = fmt.Sprint((*data)[fieldGroup.Id])
		}
	})
	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(map[string]interface{}{
			"bo":          bo,
			"relationBo":  relationBo,
			"usedCheckBo": usedCheckBo,
			//"dataSource": dataSource,
		})
	}
	return c.Render()
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
