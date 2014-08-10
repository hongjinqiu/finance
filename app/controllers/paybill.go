package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	. "com/papersns/accountinout"
	. "com/papersns/common"
	. "com/papersns/component"
	. "com/papersns/error"
	"com/papersns/global"
	. "com/papersns/model"
	. "com/papersns/model/handler"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func init() {
}

type PayBillSupport struct {
	ActionSupport
}

func (c PayBillSupport) RAfterNewData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	master := (*bo)["A"].(map[string]interface{})
	modelTemplateFactory := ModelTemplateFactory{}
	billTypeParameterDataSource := modelTemplateFactory.GetDataSource("BillPaymentTypeParameter")
	collectionName := modelTemplateFactory.GetCollectionName(billTypeParameterDataSource)
	session, _ := global.GetConnection(sessionId)
	qb := QuerySupport{}
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	queryMap := map[string]interface{}{
		"A.billTypeId": 2,
		"A.createUnit": qb.GetCreateUnitByUserId(session, userId),
	}
	//	FindByMapWithSession(session *mgo.Session, collection string, query map[string]interface{}) (result map[string]interface{}, found bool) {
	billTypeParameter, found := qb.FindByMapWithSession(session, collectionName, queryMap)
	if !found {
		panic(BusinessError{
			Message: "未找到付款单类型参数",
		})
	}
	billTypeParameterMaster := billTypeParameter["A"].(map[string]interface{})
	master["property"] = billTypeParameterMaster["property"]
	(*bo)["A"] = master

	// 币别默认值
	currencyTypeQuery := map[string]interface{}{
		"A.code":       "RMB",
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

func (o PayBillSupport) RAfterCopyData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	// 单据编号
	o.setBillNo(sessionId, bo)
}

func (o PayBillSupport) setBillNo(sessionId int, bo *map[string]interface{}) {
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
					matchStr = "000"[:(3-len(matchStr))] + matchStr
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

func (c PayBillSupport) getSrcDiffDataRowItem(diffDataRow DiffDataRow) DiffDataRow {
	tmpItem := diffDataRow
	tmpItem.DestBo = nil
	tmpItem.DestData = nil
	return tmpItem
}

func (c PayBillSupport) getDestDiffDataRowItem(diffDataRow DiffDataRow) DiffDataRow {
	tmpItem := diffDataRow
	tmpItem.SrcBo = nil
	tmpItem.SrcData = nil
	return tmpItem
}

func (c PayBillSupport) RAfterSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}, diffDataRowLi *[]DiffDataRow) {
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

func (c PayBillSupport) checkLimitsControlByDiffDataRowLi(sessionId int, bo map[string]interface{}, diffDataRowLi []DiffDataRow) {
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

func (c PayBillSupport) checkLimitsControlPanicMessage(sessionId int, bo map[string]interface{}, cashAccountDiffDataRowLi []DiffDataRow, bankAccountDiffDataRowLi []DiffDataRow) {
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

func (c PayBillSupport) RAfterDeleteData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	masterData := (*bo)["A"].(map[string]interface{})
	if fmt.Sprint(masterData["billStatus"]) == "4" { // 4为已作废,已作废单据不过账,不检查赤字
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

func (c PayBillSupport) logAccount(sessionId int, dataSource DataSource, bo map[string]interface{}, diffDataRow DiffDataRow, contextData map[string]interface{}, diffDataType int) {
	if diffDataRow.FieldGroupLi[0].IsMasterField() {
		c.logAccountForMaster(sessionId, dataSource, bo, diffDataRow, contextData, diffDataType)
	} else {
		c.logAccountForDetailB(sessionId, dataSource, bo, diffDataRow, contextData, diffDataType)
	}
}

func (c PayBillSupport) logAccountForMaster(sessionId int, dataSource DataSource, bo map[string]interface{}, diffDataRow DiffDataRow, contextData map[string]interface{}, diffDataType int) {
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

	amtPay, err := strconv.ParseFloat(fmt.Sprint(contextData["amtPay"]), 64)
	if err != nil {
		panic(err)
	}

	if amtPay >= 0 {
		accountInOutParam.AmtReduce = fmt.Sprint(contextData["amtPay"])
	} else {
		oldStr := "-"
		newStr := ""
		limit := -1
		accountInOutParam.AmtIncrease = strings.Replace(fmt.Sprint(contextData["amtPay"]), oldStr, newStr, limit)
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

		if amtPay >= 0 {
			accountInOutItemParam.AmtReduce = fmt.Sprint(contextData["amtPay"])
		} else {
			accountInOutItemParam.AmtReduce = fmt.Sprint(-amtPay)
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
		accountInOutItemParam.ChamberlainType = c.getIntData(contextData, "payerType")
		accountInOutItemParam.ChamberlainId = c.getIntData(contextData, "payerId")
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

func (c PayBillSupport) getIntData(data map[string]interface{}, fieldName string) int {
	fieldValue, err := strconv.Atoi(fmt.Sprint(data[fieldName]))
	if err != nil {
		panic(err)
	}
	return fieldValue
}

func (c PayBillSupport) getInt64Data(data map[string]interface{}, fieldName string) int64 {
	fieldValue, err := strconv.ParseInt(fmt.Sprint(data[fieldName]), 0, 64)
	if err != nil {
		panic(err)
	}
	return fieldValue
}

func (c PayBillSupport) logAccountForDetailB(sessionId int, dataSource DataSource, bo map[string]interface{}, diffDataRow DiffDataRow, contextData map[string]interface{}, diffDataType int) {
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
		accountInOutItemParam.ChamberlainType = c.getIntData(masterData, "payerType")
		accountInOutItemParam.ChamberlainId = c.getIntData(masterData, "payerId")
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
func (c PayBillSupport) RAfterCancelData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
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
func (c PayBillSupport) RAfterUnCancelData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	// 过账,
	modelIterator.IterateDataBo(dataSource, bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		diffDataRow := DiffDataRow{
			FieldGroupLi: fieldGroupLi,
			DestBo:       bo,
			DestData:     data,
		}
		c.logAccount(sessionId, dataSource, *bo, diffDataRow, *diffDataRow.DestData, ADD)
	})

	diffDataRowLi := []DiffDataRow{}
	modelIterator.IterateDataBo(dataSource, bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		diffDataRow := DiffDataRow{
			FieldGroupLi: fieldGroupLi,
			DestBo:       bo,
			DestData:     data,
		}
		diffDataRowLi = append(diffDataRowLi, diffDataRow)
	})
	c.checkLimitsControlByDiffDataRowLi(sessionId, *bo, diffDataRowLi)
}

type PayBill struct {
	BillAction
}

func (c PayBill) SaveData() revel.Result {
	c.RActionSupport = PayBillSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PayBill) DeleteData() revel.Result {
	c.RActionSupport = PayBillSupport{}

	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PayBill) EditData() revel.Result {
	c.RActionSupport = PayBillSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PayBill) NewData() revel.Result {
	c.RActionSupport = PayBillSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PayBill) GetData() revel.Result {
	c.RActionSupport = PayBillSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c PayBill) CopyData() revel.Result {
	c.RActionSupport = PayBillSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c PayBill) GiveUpData() revel.Result {
	c.RActionSupport = PayBillSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c PayBill) RefreshData() revel.Result {
	c.RActionSupport = PayBillSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c PayBill) LogList() revel.Result {
	result := c.RLogListCommon()

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
func (c PayBill) CancelData() revel.Result {
	c.RActionSupport = ActionSupport{}
	modelRenderVO := c.RCancelDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 反作废
 */
func (c PayBill) UnCancelData() revel.Result {
	c.RActionSupport = ActionSupport{}
	modelRenderVO := c.RUnCancelDataCommon()
	return c.RRenderCommon(modelRenderVO)
}
