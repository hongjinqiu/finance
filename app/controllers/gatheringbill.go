package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/accountinout"
	. "com/papersns/component"
	. "com/papersns/error"
	"com/papersns/global"
	. "com/papersns/common"
	. "com/papersns/model"
	. "com/papersns/model/handler"
	"fmt"
	"strconv"
	"strings"
	"regexp"
)

func init() {
}

type GatheringBillSupport struct {
	ActionSupport
}

func (c GatheringBillSupport) afterNewData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	master := (*bo)["A"].(map[string]interface{})
	(*bo)["A"] = master
	modelTemplateFactory := ModelTemplateFactory{}
	billTypeParameterDataSource := modelTemplateFactory.GetDataSource("BillReceiveTypeParameter")
	collectionName := modelTemplateFactory.GetCollectionName(billTypeParameterDataSource)
	session, _ := global.GetConnection(sessionId)
	qb := QuerySupport{}
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	queryMap := map[string]interface{}{
		"A.billTypeId": 1,
		"A.createUnit": qb.GetCreateUnitByUserId(session, userId),
	}
	
	billTypeParameter, found := qb.FindByMapWithSession(session, collectionName, queryMap)
	if !found {
		panic(BusinessError{
			Message: "未找到收款单类型参数",
		})
	}
	billTypeParameterMaster := billTypeParameter["A"].(map[string]interface{})
	master["property"] = billTypeParameterMaster["property"]
	
	// 币别默认值
	currencyTypeQuery := map[string]interface{}{
		"A.code": "RMB",
		"A.createUnit": qb.GetCreateUnitByUserId(session, userId),
	}
	
	currencyTypeCollectionName := "CurrencyType"
	result, found := qb.FindByMapWithSession(session, currencyTypeCollectionName, currencyTypeQuery)
	if found {
		master["currencyTypeId"] = result["id"]
	}
	
	// 单据编号
	c.setBillNo(sessionId, bo)
}

func (o GatheringBillSupport) afterCopyData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	// 单据编号
	o.setBillNo(sessionId, bo)
}

func (o GatheringBillSupport) setBillNo(sessionId int, bo *map[string]interface{}) {
	master := (*bo)["A"].(map[string]interface{})
	(*bo)["A"] = master
	session, _ := global.GetConnection(sessionId)

	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}	
	// 单据编号
	qb := QuerySupport{}
	gatheringBillCollectionName := "GatheringBill"
	billNoQuery := map[string]interface{}{
		"A.createUnit": qb.GetCreateUnitByUserId(session, userId),
	}
	pageNo := 1
	pageSize := 1
	orderBy := "-A.billNo"
	result := qb.IndexWithSession(session, gatheringBillCollectionName, billNoQuery, pageNo, pageSize, orderBy)
	items := result["items"].([]interface{})
	dateUtil := DateUtil{}
	if len(items) > 0 {
		maxItem := items[0].(map[string]interface{})
		maxItemA := maxItem["A"].(map[string]interface{})
		maxBillNo := fmt.Sprint(maxItemA["billNo"])
		
		regx := regexp.MustCompile(`^.*?(\d+)$`)
		matchResult := regx.FindStringSubmatch(maxBillNo)
		if len(matchResult) >= 2 {
			matchNum := matchResult[1]
			matchInt, err := strconv.Atoi(matchNum)
			if err != nil {
				master["billNo"] = fmt.Sprint(dateUtil.GetCurrentYyyyMMdd()) + "_001"
			} else {
				matchStr := fmt.Sprint(matchInt + 1)
				if len(matchStr) < 3 {
					matchStr = "000"[:(3 - len(matchStr))] + matchStr
				}
				master["billNo"] = fmt.Sprint(dateUtil.GetCurrentYyyyMMdd()) + "_" + matchStr
			}
		} else {
			master["billNo"] = fmt.Sprint(dateUtil.GetCurrentYyyyMMdd()) + "_001"
		}
	} else {
		master["billNo"] = fmt.Sprint(dateUtil.GetCurrentYyyyMMdd()) + "_001"
	}
}

func (c GatheringBillSupport) getSrcDiffDataRowItem(diffDataRow DiffDataRow) DiffDataRow {
	tmpItem := diffDataRow
	tmpItem.DestBo = nil
	tmpItem.DestData = nil
	return tmpItem
}

func (c GatheringBillSupport) getDestDiffDataRowItem(diffDataRow DiffDataRow) DiffDataRow {
	tmpItem := diffDataRow
	tmpItem.SrcBo = nil
	tmpItem.SrcData = nil
	return tmpItem
}

func (c GatheringBillSupport) afterSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}, diffDataRowLi *[]DiffDataRow) {
	for _, item := range *diffDataRowLi {
		if item.SrcData != nil && item.DestData != nil { // 修改
			// 旧数据反过账,新数据正过账
			c.logAccount(sessionId, dataSource, *bo, item, item.SrcData, BEFORE_UPDATE)
			c.logAccount(sessionId, dataSource, *bo, item, *item.DestData, AFTER_UPDATE)
		} else if item.SrcData == nil && item.DestData != nil { // 新增
			// 新数据正过账
			c.logAccount(sessionId, dataSource, *bo, item, *item.DestData, ADD)
		} else if item.SrcData != nil && item.DestData == nil { // 删除
			c.logAccount(sessionId, dataSource, *bo, item, item.SrcData, DELETE)
		}
	}

	c.checkLimitsControlByDiffDataRowLi(sessionId, *bo, *diffDataRowLi)
}

func (c GatheringBillSupport) checkLimitsControlByDiffDataRowLi(sessionId int, bo map[string]interface{}, diffDataRowLi []DiffDataRow) {
	// 判断赤字,先收集账户名称,并进行字段适配
	cashAccountDiffDataRowLi := []DiffDataRow{}
	bankAccountDiffDataRowLi := []DiffDataRow{}
	for _, item := range diffDataRowLi {
		if item.FieldGroupLi[0].IsMasterField() {
			if item.SrcData != nil && item.DestData != nil { // 修改
				srcProperty := fmt.Sprint(item.SrcData["property"])
				if srcProperty == "1" { // 银行存款
					bankAccountDiffDataRowLi = append(bankAccountDiffDataRowLi, c.getSrcDiffDataRowItem(item))
				} else if srcProperty == "2" { // 现金
					cashAccountDiffDataRowLi = append(cashAccountDiffDataRowLi, c.getSrcDiffDataRowItem(item))
				}
				destProperty := fmt.Sprint((*item.DestData)["property"])
				if destProperty == "1" { // 银行存款
					bankAccountDiffDataRowLi = append(bankAccountDiffDataRowLi, c.getDestDiffDataRowItem(item))
				} else if destProperty == "2" { // 现金
					cashAccountDiffDataRowLi = append(cashAccountDiffDataRowLi, c.getDestDiffDataRowItem(item))
				}
			} else if item.SrcData == nil && item.DestData != nil { // 新增
				destProperty := fmt.Sprint((*item.DestData)["property"])
				if destProperty == "1" { // 银行存款
					bankAccountDiffDataRowLi = append(bankAccountDiffDataRowLi, c.getDestDiffDataRowItem(item))
				} else if destProperty == "2" { // 现金
					cashAccountDiffDataRowLi = append(cashAccountDiffDataRowLi, c.getDestDiffDataRowItem(item))
				}
			} else if item.SrcData != nil && item.DestData == nil { // 删除
				srcProperty := fmt.Sprint(item.SrcData["property"])
				if srcProperty == "1" { // 银行存款
					bankAccountDiffDataRowLi = append(bankAccountDiffDataRowLi, c.getSrcDiffDataRowItem(item))
				} else if srcProperty == "2" { // 现金
					cashAccountDiffDataRowLi = append(cashAccountDiffDataRowLi, c.getSrcDiffDataRowItem(item))
				}
			}
		} else {
			if item.SrcData != nil && item.DestData != nil { // 修改
				srcAccountType := fmt.Sprint(item.SrcData["accountType"])
				srcDiffDataRowItem := c.getSrcDiffDataRowItem(item)
				srcMasterData := srcDiffDataRowItem.SrcBo["A"].(map[string]interface{})
				srcDiffDataRowItem.SrcData["currencyTypeId"] = srcMasterData["currencyTypeId"]
				if srcAccountType == "1" { // 现金
					cashAccountDiffDataRowLi = append(cashAccountDiffDataRowLi, srcDiffDataRowItem)
				} else if srcAccountType == "2" { // 银行
					bankAccountDiffDataRowLi = append(bankAccountDiffDataRowLi, srcDiffDataRowItem)
				}
				destAccountType := fmt.Sprint((*item.DestData)["accountType"])
				destDiffDataRowItem := c.getDestDiffDataRowItem(item)
				destMasterData := (*destDiffDataRowItem.DestBo)["A"].(map[string]interface{})
				(*destDiffDataRowItem.DestData)["currencyTypeId"] = destMasterData["currencyTypeId"]
				if destAccountType == "1" { // 现金
					cashAccountDiffDataRowLi = append(cashAccountDiffDataRowLi, destDiffDataRowItem)
				} else if destAccountType == "2" { // 银行
					bankAccountDiffDataRowLi = append(bankAccountDiffDataRowLi, destDiffDataRowItem)
				}
			} else if item.SrcData == nil && item.DestData != nil { // 新增
				destAccountType := fmt.Sprint((*item.DestData)["accountType"])
				destDiffDataRowItem := c.getDestDiffDataRowItem(item)
				destMasterData := (*destDiffDataRowItem.DestBo)["A"].(map[string]interface{})
				(*destDiffDataRowItem.DestData)["currencyTypeId"] = destMasterData["currencyTypeId"]
				if destAccountType == "1" { // 现金
					cashAccountDiffDataRowLi = append(cashAccountDiffDataRowLi, destDiffDataRowItem)
				} else if destAccountType == "2" { // 银行
					bankAccountDiffDataRowLi = append(bankAccountDiffDataRowLi, destDiffDataRowItem)
				}
			} else if item.SrcData != nil && item.DestData == nil { // 删除
				srcAccountType := fmt.Sprint(item.SrcData["accountType"])
				srcDiffDataRowItem := c.getSrcDiffDataRowItem(item)
				srcMasterData := srcDiffDataRowItem.SrcBo["A"].(map[string]interface{})
				srcDiffDataRowItem.SrcData["currencyTypeId"] = srcMasterData["currencyTypeId"]
				if srcAccountType == "1" { // 现金
					cashAccountDiffDataRowLi = append(cashAccountDiffDataRowLi, srcDiffDataRowItem)
				} else if srcAccountType == "2" { // 银行
					bankAccountDiffDataRowLi = append(bankAccountDiffDataRowLi, srcDiffDataRowItem)
				}
			}
		}
	}
	c.checkLimitsControlPanicMessage(sessionId, bo, cashAccountDiffDataRowLi, bankAccountDiffDataRowLi)
}

func (c GatheringBillSupport) checkLimitsControlPanicMessage(sessionId int, bo map[string]interface{}, cashAccountDiffDataRowLi []DiffDataRow, bankAccountDiffDataRowLi []DiffDataRow) {
	accountInOutService := AccountInOutService{}
	cashForbidLi, cashWarnLi := accountInOutService.CheckCashAccountDiffDataRowLimitControl(sessionId, cashAccountDiffDataRowLi)
	bankForbidLi, bankWarnLi := accountInOutService.CheckBankAccountDiffDataRowLimitControl(sessionId, bankAccountDiffDataRowLi)

	forbidLi := []string{}
	warnLi := []string{}
	for _, item := range cashForbidLi {
		forbidLi = append(forbidLi, item)
	}
	for _, item := range bankForbidLi {
		forbidLi = append(forbidLi, item)
	}
	for _, item := range cashWarnLi {
		warnLi = append(warnLi, item)
	}
	for _, item := range bankWarnLi {
		warnLi = append(warnLi, item)
	}
	if len(forbidLi) > 0 {
		panic(BusinessError{
			Code:    LIMIT_CONTROL_FORBID,
			Message: strings.Join(forbidLi, "<br />"),
		})
	}
	continueAnyAll := "false"
	if bo["continueAnyAll"] != nil && fmt.Sprint(bo["continueAnyAll"]) != "" {
		continueAnyAll = bo["continueAnyAll"].(string)
	}
	if len(warnLi) > 0 && continueAnyAll == "false" {
		panic(BusinessError{
			Code:    LIMIT_CONTROL_WARN,
			Message: strings.Join(warnLi, "<br />"),
		})
	}
}

func (c GatheringBillSupport) afterDeleteData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	masterData := (*bo)["A"].(map[string]interface{})
	if fmt.Sprint(masterData["billStatus"]) == "4" {// 4为已作废,已作废单据不过账,不检查赤字
		return
	}
	modelIterator := ModelIterator{}
	var result interface{} = ""
	// 过账,
	modelIterator.IterateDataBo(dataSource, bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		diffDataRow := DiffDataRow{
			FieldGroupLi: fieldGroupLi,
			SrcBo:        *bo,
			SrcData:      *data,
		}
		c.logAccount(sessionId, dataSource, *bo, diffDataRow, diffDataRow.SrcData, DELETE)
	})

	diffDataRowLi := []DiffDataRow{}
	modelIterator.IterateDataBo(dataSource, bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		diffDataRow := DiffDataRow{
			FieldGroupLi: fieldGroupLi,
			SrcBo:        *bo,
			SrcData:      *data,
		}
		diffDataRowLi = append(diffDataRowLi, diffDataRow)
	})
	c.checkLimitsControlByDiffDataRowLi(sessionId, *bo, diffDataRowLi)
}

func (c GatheringBillSupport) logAccount(sessionId int, dataSource DataSource, bo map[string]interface{}, diffDataRow DiffDataRow, contextData map[string]interface{}, diffDataType int) {
	if diffDataRow.FieldGroupLi[0].IsMasterField() {
		c.logAccountForMaster(sessionId, dataSource, bo, diffDataRow, contextData, diffDataType)
	} else {
		c.logAccountForDetailB(sessionId, dataSource, bo, diffDataRow, contextData, diffDataType)
	}
}

func (c GatheringBillSupport) logAccountForMaster(sessionId int, dataSource DataSource, bo map[string]interface{}, diffDataRow DiffDataRow, contextData map[string]interface{}, diffDataType int) {
	modelTemplateFactory := ModelTemplateFactory{}
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)

	accountInOutService := AccountInOutService{}
	ymd := c.getIntData(contextData, "billDate")
	year, sequenceNo := accountInOutService.GetAccountingPeriodYearSequenceNo(sessionId, ymd)

	accountInOutParam := AccountInOutParam{}
	property := c.getIntData(contextData, "property") // 1:银行存款,2:现金,3:其它
	if property != 1 && property != 2 {
		return
	}
	var accountType int
	if property == 1 { // 银行存款
		accountType = 2 // 银行
	} else if property == 2 { // 现金
		accountType = 1 // 现金
	}

	accountInOutParam.AccountType = accountType
	accountInOutParam.AccountId = c.getIntData(contextData, "accountId")
	accountInOutParam.CurrencyTypeId = c.getIntData(contextData, "currencyTypeId")
	accountInOutParam.ExchangeRateShow = fmt.Sprint(contextData["exchangeRateShow"])
	accountInOutParam.ExchangeRate = fmt.Sprint(contextData["exchangeRate"])
	accountInOutParam.AccountingPeriodYear = year
	accountInOutParam.AccountingPeriodMonth = sequenceNo

	amtGathering, err := strconv.ParseFloat(fmt.Sprint(contextData["amtGathering"]), 64)
	if err != nil {
		panic(err)
	}

	if amtGathering >= 0 {
		accountInOutParam.AmtIncrease = fmt.Sprint(contextData["amtGathering"])
	} else {
		oldStr := "-"
		newStr := ""
		limit := -1
		accountInOutParam.AmtReduce = strings.Replace(fmt.Sprint(contextData["amtGathering"]), oldStr, newStr, limit)
	}
	accountInOutParam.DiffDataType = diffDataType
	accountInOutParam.CreateBy = c.getIntData(contextData, "createBy")
	accountInOutParam.CreateTime = c.getInt64Data(contextData, "createTime")
	accountInOutParam.CreateUnit = c.getIntData(contextData, "createUnit")
	// 删除数据时,有可能数据没有modifyBy,modifyUnit,modifyTime,因此,做空判断
	if contextData["modifyBy"] != nil {
		accountInOutParam.ModifyBy = c.getIntData(contextData, "modifyBy")
	}
	if contextData["modifyUnit"] != nil {
		accountInOutParam.ModifyUnit = c.getIntData(contextData, "modifyUnit")
	}
	if contextData["modifyTime"] != nil {
		accountInOutParam.ModifyTime = c.getInt64Data(contextData, "modifyTime")
	}

	accountInOutItemParam := AccountInOutItemParam{}
	{
		accountInOutItemParam.AccountType = accountType
		accountInOutItemParam.AccountId = accountInOutParam.AccountId
		accountInOutItemParam.CurrencyTypeId = accountInOutParam.CurrencyTypeId
		accountInOutItemParam.ExchangeRateShow = accountInOutParam.ExchangeRateShow
		accountInOutItemParam.ExchangeRate = accountInOutParam.ExchangeRate

		if amtGathering >= 0 {
			accountInOutItemParam.AmtIncrease = fmt.Sprint(contextData["amtGathering"])
		} else {
			accountInOutItemParam.AmtReduce = fmt.Sprint(-amtGathering)
		}
		accountInOutItemParam.BillTypeId = c.getIntData(contextData, "billTypeId")
		accountInOutItemParam.BillDataSourceName = dataSource.Id
		accountInOutItemParam.BillCollectionName = collectionName
		accountInOutItemParam.BillDetailName = diffDataRow.FieldGroupLi[0].GetDataSetId()
		accountInOutItemParam.BillId = c.getIntData(contextData, "id")
		//		accountInOutItemParam.BillDetailId
		accountInOutItemParam.BillNo = fmt.Sprint(contextData["billNo"])
		accountInOutItemParam.BillDate = c.getIntData(contextData, "billDate")
		accountInOutItemParam.BalanceDate = c.getIntData(contextData, "balanceDate")
		accountInOutItemParam.BalanceTypeId = c.getIntData(contextData, "balanceTypeId")
		accountInOutItemParam.BalanceNo = fmt.Sprint(contextData["balanceNo"])
		accountInOutItemParam.ChamberlainType = c.getIntData(contextData, "chamberlainType")
		accountInOutItemParam.ChamberlainId = c.getIntData(contextData, "chamberlainId")
		accountInOutItemParam.CreateBy = accountInOutParam.CreateBy
		accountInOutItemParam.CreateTime = accountInOutParam.CreateTime
		accountInOutItemParam.CreateUnit = accountInOutParam.CreateUnit
		accountInOutItemParam.ModifyBy = accountInOutParam.ModifyBy
		accountInOutItemParam.ModifyUnit = accountInOutParam.ModifyUnit
		accountInOutItemParam.ModifyTime = accountInOutParam.ModifyTime
	}

	accountInOutParam.AccountInOutItemParam = accountInOutItemParam
	if accountType == 1 { // 现金
		accountInOutService.LogAllCashDeposit(sessionId, accountInOutParam)
	} else if accountType == 2 { // 银行
		accountInOutService.LogAllBankDeposit(sessionId, accountInOutParam)
	}
}

func (c GatheringBillSupport) getIntData(data map[string]interface{}, fieldName string) int {
	fieldValue, err := strconv.Atoi(fmt.Sprint(data[fieldName]))
	if err != nil {
		panic(err)
	}
	return fieldValue
}

func (c GatheringBillSupport) getInt64Data(data map[string]interface{}, fieldName string) int64 {
	fieldValue, err := strconv.ParseInt(fmt.Sprint(data[fieldName]), 0, 64)
	if err != nil {
		panic(err)
	}
	return fieldValue
}

func (c GatheringBillSupport) logAccountForDetailB(sessionId int, dataSource DataSource, bo map[string]interface{}, diffDataRow DiffDataRow, contextData map[string]interface{}, diffDataType int) {
	masterData := bo["A"].(map[string]interface{})
	modelTemplateFactory := ModelTemplateFactory{}
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)

	accountInOutService := AccountInOutService{}
	ymd := c.getIntData(masterData, "billDate")
	year, sequenceNo := accountInOutService.GetAccountingPeriodYearSequenceNo(sessionId, ymd)

	accountInOutParam := AccountInOutParam{}
	accountType := c.getIntData(contextData, "accountType")

	accountInOutParam.AccountType = accountType
	accountInOutParam.AccountId = c.getIntData(contextData, "accountId")
	accountInOutParam.CurrencyTypeId = c.getIntData(masterData, "currencyTypeId") // 分录里面没有币别,依表头币别
	accountInOutParam.ExchangeRateShow = fmt.Sprint(masterData["exchangeRateShow"])
	accountInOutParam.ExchangeRate = fmt.Sprint(masterData["exchangeRate"])
	accountInOutParam.AccountingPeriodYear = year
	accountInOutParam.AccountingPeriodMonth = sequenceNo

	amtFee, err := strconv.ParseFloat(fmt.Sprint(contextData["amtFee"]), 64)
	if err != nil {
		panic(err)
	}

	if amtFee >= 0 { // 费用>0时,写到减少字段
		accountInOutParam.AmtReduce = fmt.Sprint(contextData["amtFee"])
	} else {
		oldStr := "-"
		newStr := ""
		limit := -1
		accountInOutParam.AmtIncrease = strings.Replace(fmt.Sprint(contextData["amtFee"]), oldStr, newStr, limit)
	}
	accountInOutParam.DiffDataType = diffDataType
	accountInOutParam.CreateBy = c.getIntData(contextData, "createBy")
	accountInOutParam.CreateTime = c.getInt64Data(contextData, "createTime")
	accountInOutParam.CreateUnit = c.getIntData(contextData, "createUnit")
	accountInOutParam.ModifyBy = c.getIntData(contextData, "modifyBy")
	accountInOutParam.ModifyUnit = c.getIntData(contextData, "modifyUnit")
	accountInOutParam.ModifyTime = c.getInt64Data(contextData, "modifyTime")

	accountInOutItemParam := AccountInOutItemParam{}
	{
		accountInOutItemParam.AccountType = accountType
		accountInOutItemParam.AccountId = accountInOutParam.AccountId
		accountInOutItemParam.CurrencyTypeId = accountInOutParam.CurrencyTypeId
		accountInOutItemParam.ExchangeRateShow = accountInOutParam.ExchangeRateShow
		accountInOutItemParam.ExchangeRate = accountInOutParam.ExchangeRate

		if amtFee >= 0 {
			accountInOutItemParam.AmtReduce = fmt.Sprint(contextData["amtFee"])
		} else {
			accountInOutItemParam.AmtIncrease = fmt.Sprint(-amtFee)
		}
		accountInOutItemParam.BillTypeId = c.getIntData(masterData, "billTypeId")
		accountInOutItemParam.BillDataSourceName = dataSource.Id
		accountInOutItemParam.BillCollectionName = collectionName
		accountInOutItemParam.BillDetailName = diffDataRow.FieldGroupLi[0].GetDataSetId()
		accountInOutItemParam.BillId = c.getIntData(masterData, "id")
		accountInOutItemParam.BillDetailId = c.getIntData(contextData, "id")
		accountInOutItemParam.BillNo = fmt.Sprint(masterData["billNo"])
		accountInOutItemParam.BillDate = c.getIntData(masterData, "billDate")
		accountInOutItemParam.BalanceDate = c.getIntData(masterData, "balanceDate")
		accountInOutItemParam.BalanceTypeId = c.getIntData(masterData, "balanceTypeId")
		accountInOutItemParam.BalanceNo = fmt.Sprint(masterData["balanceNo"])
		accountInOutItemParam.ChamberlainType = c.getIntData(masterData, "chamberlainType")
		accountInOutItemParam.ChamberlainId = c.getIntData(masterData, "chamberlainId")
		accountInOutItemParam.CreateBy = accountInOutParam.CreateBy
		accountInOutItemParam.CreateTime = accountInOutParam.CreateTime
		accountInOutItemParam.CreateUnit = accountInOutParam.CreateUnit
		accountInOutItemParam.ModifyBy = accountInOutParam.ModifyBy
		accountInOutItemParam.ModifyUnit = accountInOutParam.ModifyUnit
		accountInOutItemParam.ModifyTime = accountInOutParam.ModifyTime
	}

	accountInOutParam.AccountInOutItemParam = accountInOutItemParam
	if accountType == 1 { // 现金
		accountInOutService.LogAllCashDeposit(sessionId, accountInOutParam)
	} else if accountType == 2 { // 银行
		accountInOutService.LogAllBankDeposit(sessionId, accountInOutParam)
	}
}

/**
 * 作废过账,赤字判断
*/
func (c GatheringBillSupport) afterCancelData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{})   {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	// 过账,
	modelIterator.IterateDataBo(dataSource, bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		diffDataRow := DiffDataRow{
			FieldGroupLi: fieldGroupLi,
			SrcBo:        *bo,
			SrcData:      *data,
		}
		c.logAccount(sessionId, dataSource, *bo, diffDataRow, diffDataRow.SrcData, DELETE)
	})

	diffDataRowLi := []DiffDataRow{}
	modelIterator.IterateDataBo(dataSource, bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		diffDataRow := DiffDataRow{
			FieldGroupLi: fieldGroupLi,
			SrcBo:        *bo,
			SrcData:      *data,
		}
		diffDataRowLi = append(diffDataRowLi, diffDataRow)
	})
	c.checkLimitsControlByDiffDataRowLi(sessionId, *bo, diffDataRowLi)
}

/**
 * 反作废过账,赤字判断
*/
func (c GatheringBillSupport) afterUnCancelData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{})  {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	// 过账,
	modelIterator.IterateDataBo(dataSource, bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		diffDataRow := DiffDataRow{
			FieldGroupLi: fieldGroupLi,
			DestBo:        bo,
			DestData:      data,
		}
		c.logAccount(sessionId, dataSource, *bo, diffDataRow, *diffDataRow.DestData, ADD)
	})

	diffDataRowLi := []DiffDataRow{}
	modelIterator.IterateDataBo(dataSource, bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		diffDataRow := DiffDataRow{
			FieldGroupLi: fieldGroupLi,
			DestBo:        bo,
			DestData:      data,
		}
		diffDataRowLi = append(diffDataRowLi, diffDataRow)
	})
	c.checkLimitsControlByDiffDataRowLi(sessionId, *bo, diffDataRowLi)
}

type GatheringBill struct {
	BillAction
}

func (c GatheringBill) SaveData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c GatheringBill) DeleteData() revel.Result {
	c.actionSupport = GatheringBillSupport{}

	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c GatheringBill) EditData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c GatheringBill) NewData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c GatheringBill) GetData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c GatheringBill) CopyData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c GatheringBill) GiveUpData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c GatheringBill) RefreshData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c GatheringBill) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}

/**
 * 作废
 */
func (c GatheringBill) CancelData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.cancelDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 反作废
 */
func (c GatheringBill) UnCancelData() revel.Result {
	c.actionSupport = GatheringBillSupport{}
	modelRenderVO := c.unCancelDataCommon()
	return c.renderCommon(modelRenderVO)
}
