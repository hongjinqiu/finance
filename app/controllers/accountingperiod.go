package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/common"
	. "com/papersns/component"
	. "com/papersns/error"
	"com/papersns/global"
	. "com/papersns/model"
	. "com/papersns/revel"
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

func (o AccountingPeriodSupport) RAfterNewData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
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
		data["id"] = "RAfterNewData" + fmt.Sprint(i)
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

/**
删除前判断被用，会计期内有单据则视为被用
*/
func (o AccountingPeriodSupport) RBeforeDeleteData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	bDataSetLi := (*bo)["B"].([]interface{})
	firstLineData := bDataSetLi[0].(map[string]interface{})
	lastLineData := bDataSetLi[len(bDataSetLi)-1].(map[string]interface{})

	commonUtil := CommonUtil{}
	firstStartDate := commonUtil.GetIntFromMap(firstLineData, "startDate")
	lastEndDate := commonUtil.GetIntFromMap(lastLineData, "endDate")

	qb := QuerySupport{}
	session, _ := global.GetConnection(sessionId)
	queryMap := map[string]interface{}{
		"A.billDate": map[string]interface{}{
			"$gte": firstStartDate,
			"$lt":  lastEndDate,
		},
	}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
	}

	modelTemplateFactory := ModelTemplateFactory{}
	dataSourceIdLi := []string{"GatheringBill", "PayBill"}
	for _, dataSourceId := range dataSourceIdLi {
		tmpDataSource := modelTemplateFactory.GetDataSource(dataSourceId)
		collectionName := modelTemplateFactory.GetCollectionName(tmpDataSource)
		_, found := qb.FindByMapWithSession(session, collectionName, queryMap)
		if found {
			panic(BusinessError{Message: "会计期范围内存在单据，不能删除"})
		}
	}
}

type AccountingPeriod struct {
	BaseDataAction
}

func (c AccountingPeriod) RRenderCommon(modelRenderVO ModelRenderVO) revel.Result {
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

		qb := QuerySupport{}
		//GatheringBill,PayBill
		sessionId := global.GetSessionId()
		global.SetGlobalAttr(sessionId, "userId", fmt.Sprint(modelRenderVO.UserId))
		defer global.CloseSession(sessionId)
		session, _ := global.GetConnection(sessionId)

		queryMap := map[string]interface{}{
			"A.billDate": map[string]interface{}{
				"$gte": firstStartDate,
				"$lt":  lastEndDate,
			},
		}
		permissionSupport := PermissionSupport{}
		permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, modelRenderVO.FormTemplate.Security)
		for k, v := range permissionQueryDict {
			queryMap[k] = v
		}

		dataSourceIdLi := []string{"GatheringBill", "PayBill"}
		for _, dataSourceId := range dataSourceIdLi {
			tmpDataSource := modelTemplateFactory.GetDataSource(dataSourceId)
			collectionName := modelTemplateFactory.GetCollectionName(tmpDataSource)
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
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

func (c AccountingPeriod) SaveData() revel.Result {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountingPeriod) DeleteData() revel.Result {
	c.RActionSupport = AccountingPeriodSupport{}

	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountingPeriod) EditData() revel.Result {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountingPeriod) NewData() revel.Result {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountingPeriod) GetData() revel.Result {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c AccountingPeriod) CopyData() revel.Result {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountingPeriod) GiveUpData() revel.Result {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountingPeriod) RefreshData() revel.Result {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountingPeriod) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
