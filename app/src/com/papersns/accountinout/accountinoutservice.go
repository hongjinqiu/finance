package accountinout

import (
	. "com/papersns/common"
	. "com/papersns/component"
	. "com/papersns/error"
	"com/papersns/global"
	. "com/papersns/model"
	. "com/papersns/model/handler"
	. "com/papersns/mongo"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type AccountInOutService struct{}

const (
	LIMIT_CONTROL_SUCCESS = 1 // 检查通过
	LIMIT_CONTROL_FORBID  = 2 // 赤字禁止
	LIMIT_CONTROL_WARN    = 3 // 赤字警告
)

func (o AccountInOutService) GetFirstAccountingPeriodStartEndDate(sessionId int, year int) (int, int) {
	session, _ := global.GetConnection(sessionId)
	dataSourceModelId := "AccountingPeriod"
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	qb := QuerySupport{}
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	queryMap := map[string]interface{}{
		"A.accountingYear": year,
		"A.createUnit": qb.GetCreateUnitByUserId(session, userId),
	}
	accountingPeriod, found := qb.FindByMapWithSession(session, collectionName, queryMap)
	if !found {
		//		panic(BusinessError{Message: "会计年度:" + fmt.Sprint(year) + ",会计期序号:" + fmt.Sprint(sequenceNo) + "未找到对应会计期"})
		log.Println("会计年度:" + fmt.Sprint(year) + "未找到对应会计期")
		return 0, 0
	}
	var startDate int
	var endDate int
	bDataSetLi := accountingPeriod["B"].([]interface{})
	commonUtil := CommonUtil{}
	for _, item := range bDataSetLi {
		line := item.(map[string]interface{})
		startDate = commonUtil.GetIntFromMap(line, "startDate")
		endDate = commonUtil.GetIntFromMap(line, "endDate")
		break
	}
	return startDate, endDate
}

func (o AccountInOutService) GetAccountingPeriodStartEndDate(sessionId int, year int, sequenceNo int) (int, int) {
	session, _ := global.GetConnection(sessionId)
	dataSourceModelId := "AccountingPeriod"
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	qb := QuerySupport{}
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	queryMap := map[string]interface{}{
		"A.accountingYear": year,
		"B.sequenceNo":     sequenceNo,
		"A.createUnit": qb.GetCreateUnitByUserId(session, userId),
	}
	accountingPeriod, found := qb.FindByMapWithSession(session, collectionName, queryMap)
	if !found {
		//		panic(BusinessError{Message: "会计年度:" + fmt.Sprint(year) + ",会计期序号:" + fmt.Sprint(sequenceNo) + "未找到对应会计期"})
		log.Println("会计年度:" + fmt.Sprint(year) + ",会计期序号:" + fmt.Sprint(sequenceNo) + "未找到对应会计期")
		return 0, 0
	}
	var startDate int
	var endDate int
	bDataSetLi := accountingPeriod["B"].([]interface{})
	commonUtil := CommonUtil{}
	for _, item := range bDataSetLi {
		line := item.(map[string]interface{})
		if fmt.Sprint(line["sequenceNo"]) == fmt.Sprint(sequenceNo) {
			startDate = commonUtil.GetIntFromMap(line, "startDate")
			endDate = commonUtil.GetIntFromMap(line, "endDate")
			break
		}
	}
	return startDate, endDate
}

func (o AccountInOutService) GetAccountingPeriodYearSequenceNo(sessionId int, ymd int) (int, int) {
	session, _ := global.GetConnection(sessionId)
	dataSourceModelId := "AccountingPeriod"
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	qb := QuerySupport{}
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	queryMap := map[string]interface{}{
		"B.startDate": map[string]interface{}{
			"$lte": ymd,
		},
		"B.endDate": map[string]interface{}{
			"$gte": ymd,
		},
		"A.createUnit": qb.GetCreateUnitByUserId(session, userId),
	}
	accountingPeriod, found := qb.FindByMapWithSession(session, collectionName, queryMap)
	if !found {
		billDate, err := time.Parse("20060102", fmt.Sprint(ymd))
		if err != nil {
			panic(err)
		}
		log.Println("单据日期" + billDate.Format("2006-01-02") + "未找到对应会计期")
		return 0,0
//		panic(BusinessError{Message: "单据日期" + billDate.Format("2006-01-02") + "未找到对应会计期"})
	}
	masterData := accountingPeriod["A"].(map[string]interface{})
	year, err := strconv.Atoi(fmt.Sprint(masterData["accountingYear"]))
	if err != nil {
		panic(err)
	}
	var sequenceNo int
	bDataSetLi := accountingPeriod["B"].([]interface{})
	for _, item := range bDataSetLi {
		line := item.(map[string]interface{})
		if fmt.Sprint(line["startDate"]) <= fmt.Sprint(ymd) && fmt.Sprint(ymd) <= fmt.Sprint(line["endDate"]) {
			sequenceNo, err = strconv.Atoi(fmt.Sprint(line["sequenceNo"]))
			if err != nil {
				panic(err)
			}
			break
		}
	}
//	if sequenceNo == 0 {
//		billDate, err := time.Parse("20060102", fmt.Sprint(ymd))
//		if err != nil {
//			panic(err)
//		}
//		panic(BusinessError{Message: "单据日期" + billDate.Format("2006-01-02") + "找到年度" + fmt.Sprint(year) + ",未找到会计期"})
//	}

	return year, sequenceNo
}

/**
 * 通用的现金帐户赤字检查,假设账户字段名都为:accountId
 * @param session
 * @param diffDateRowAllLi
 * @return [[string,...], [string,...]]
 */
func (o AccountInOutService) CheckCashAccountDiffDataRowLimitControl(sessionId int, diffDateRowAllLi []DiffDataRow) ([]string, []string) {
	accountIdLi := []int{}
	for _, item := range diffDateRowAllLi {
		if item.SrcData != nil {
			accountId, err := strconv.Atoi(fmt.Sprint((item.SrcData)["accountId"]))
			if err != nil {
				panic(err)
			}
			if !o.isAccountIdExist(accountIdLi, accountId) {
				accountIdLi = append(accountIdLi, accountId)
			}
		}
		if item.DestData != nil {
			accountId, err := strconv.Atoi(fmt.Sprint((*item.DestData)["accountId"]))
			if err != nil {
				panic(err)
			}
			if !o.isAccountIdExist(accountIdLi, accountId) {
				accountIdLi = append(accountIdLi, accountId)
			}
		}
	}
	forbidLi := []string{} // 帐户赤字报错信息
	warnLi := []string{}   // 帐户赤字警告信息
	accountInOutService := AccountInOutService{}
	for _, accountId := range accountIdLi {
		checkResult := accountInOutService.CheckCashAccountLimitControl(sessionId, accountId)
		code := checkResult["code"].(int)
		if code == LIMIT_CONTROL_FORBID {
			forbidLi = append(forbidLi, checkResult["message"].(string))
		} else if code == LIMIT_CONTROL_WARN {
			warnLi = append(warnLi, checkResult["message"].(string))
		}
	}
	return forbidLi, warnLi
}

func (o AccountInOutService) isAccountIdExist(accountIdLi []int, accountId int) bool {
	for _, accountIdInArray := range accountIdLi {
		if accountIdInArray == accountId {
			return true
		}
	}
	return false
}

type AccountIdCurrencyTypeIdVO struct {
	AccountId      int
	CurrencyTypeId int
}

func (o AccountInOutService) CheckBankAccountDiffDataRowLimitControl(sessionId int, diffDateRowAllLi []DiffDataRow) ([]string, []string) {
	accountIdCurrencyTypeIdVOLi := []AccountIdCurrencyTypeIdVO{}
	for _, item := range diffDateRowAllLi {
		if item.SrcData != nil {
			accountId, err := strconv.Atoi(fmt.Sprint((item.SrcData)["accountId"]))
			if err != nil {
				panic(err)
			}
			currencyTypeId, err := strconv.Atoi(fmt.Sprint((item.SrcData)["currencyTypeId"]))
			if err != nil {
				panic(err)
			}
			if !o.isAccountIdCurrencyTypeIdExist(accountIdCurrencyTypeIdVOLi, accountId, currencyTypeId) {
				accountIdCurrencyTypeIdVOLi = append(accountIdCurrencyTypeIdVOLi, AccountIdCurrencyTypeIdVO{
					AccountId:      accountId,
					CurrencyTypeId: currencyTypeId,
				})
			}
		}
		if item.DestData != nil {
			accountId, err := strconv.Atoi(fmt.Sprint((*item.DestData)["accountId"]))
			if err != nil {
				panic(err)
			}
			currencyTypeId, err := strconv.Atoi(fmt.Sprint((*item.DestData)["currencyTypeId"]))
			if err != nil {
				panic(err)
			}
			if !o.isAccountIdCurrencyTypeIdExist(accountIdCurrencyTypeIdVOLi, accountId, currencyTypeId) {
				accountIdCurrencyTypeIdVOLi = append(accountIdCurrencyTypeIdVOLi, AccountIdCurrencyTypeIdVO{
					AccountId:      accountId,
					CurrencyTypeId: currencyTypeId,
				})
			}
		}
	}
	forbidLi := []string{} // 帐户赤字报错信息
	warnLi := []string{}   // 帐户赤字警告信息
	accountInOutService := AccountInOutService{}
	for _, accountIdCurrencyTypeIdVO := range accountIdCurrencyTypeIdVOLi {
		checkResult := accountInOutService.CheckBankAccountLimitControl(sessionId, accountIdCurrencyTypeIdVO.AccountId, accountIdCurrencyTypeIdVO.CurrencyTypeId)
		code := checkResult["code"].(int)
		if code == LIMIT_CONTROL_FORBID {
			forbidLi = append(forbidLi, checkResult["message"].(string))
		} else if code == LIMIT_CONTROL_WARN {
			warnLi = append(warnLi, checkResult["message"].(string))
		}
	}
	return forbidLi, warnLi
}

func (c AccountInOutService) isAccountIdCurrencyTypeIdExist(accountIdCurrencyTypeIdVOLi []AccountIdCurrencyTypeIdVO, accountId int, currencyTypeId int) bool {
	for _, accountIdCurrencyTypeIdVO := range accountIdCurrencyTypeIdVOLi {
		if accountIdCurrencyTypeIdVO.AccountId == accountId && accountIdCurrencyTypeIdVO.CurrencyTypeId == currencyTypeId {
			return true
		}
	}
	return false
}

/**
 * 检查现金帐户赤字,供过帐完毕后调用
 * @param session
 * @param accountId
 * @return {code: int, message: string}
 * code LIMIT_CONTROL_SUCCESS:检查通过
 * code LIMIT_CONTROL_FORBID:赤字字段为禁止,保存时金额 < 0,需要调用方抛出异常通知客户端
 * code LIMIT_CONTROL_WARN:赤字字段为警告,保存时金额 < 0,需要调用方提供警告信息通知客户端
 */
func (o AccountInOutService) CheckCashAccountLimitControl(sessionId int, accountId int) map[string]interface{} {
	qb := QuerySupport{}
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	session, _ := global.GetConnection(sessionId)
	queryMap := map[string]interface{}{
		"_id": accountId,
		"A.createUnit": qb.GetCreateUnitByUserId(session, userId),
	}
	collectionName := "CashAccount"
	cashAccountBo, found := qb.FindByMapWithSession(session, collectionName, queryMap)
	if !found {
		queryMapByte, err := json.Marshal(&queryMap)
		if err != nil {
			panic(err)
		}
		panic(BusinessError{Message: "现金账户没找到，查询条件为:" + string(queryMapByte)})
	}
	cashAccount := cashAccountBo["A"].(map[string]interface{})
	limitsControl := fmt.Sprint(cashAccount["limitsControl"])
	commonUtil := CommonUtil{}
	amtOriginalCurrencyBalance := commonUtil.GetFloat64FromString(fmt.Sprint(cashAccount["amtOriginalCurrencyBalance"]))
	if limitsControl == "1" { // 禁止
		if amtOriginalCurrencyBalance < 0 {
			result := map[string]interface{}{
				"code":    LIMIT_CONTROL_FORBID,
				"message": "帐户:" + fmt.Sprint(cashAccount["name"]) + "出现赤字！",
			}
			return result
		}
	} else if limitsControl == "2" { // 警告
		if amtOriginalCurrencyBalance < 0 {
			result := map[string]interface{}{
				"code":    LIMIT_CONTROL_WARN,
				"message": "帐户:" + fmt.Sprint(cashAccount["name"]) + "出现赤字！",
			}
			return result
		}
	}
	return map[string]interface{}{
		"code": LIMIT_CONTROL_SUCCESS,
	}
}

/**
 * 银行帐户赤字检查
 * @param sessionId
 * @param accountId
 * @param currencyTypeId
 * @return {code: int, message: string}
 * code AccountLimitControlStatus.SUCCESS:检查通过
 * code AccountLimitControlStatus.FORBID:赤字字段为禁止,保存时金额 < 0,需要调用方抛出异常通知客户端
 * code AccountLimitControlStatus.WARN:赤字字段为警告,保存时金额 < 0,需要调用方提供警告信息通知客户端
 */
func (o AccountInOutService) CheckBankAccountLimitControl(sessionId int, accountId int, currencyTypeId int) map[string]interface{} {
	session, _ := global.GetConnection(sessionId)
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	qb := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": accountId,
		"A.createUnit": qb.GetCreateUnitByUserId(session, userId),
	}
	collectionName := "BankAccount"
	bankAccountBo, found := qb.FindByMapWithSession(session, collectionName, queryMap)
	if !found {
		queryMapByte, err := json.Marshal(&queryMap)
		if err != nil {
			panic(err)
		}
		panic(BusinessError{Message: "银行账户没找到，查询条件为:" + string(queryMapByte)})
	}
	bankAccountMaster := bankAccountBo["A"].(map[string]interface{})
	bDetailDataLi := bankAccountBo["B"].([]interface{})
	commonUtil := CommonUtil{}
	for _, item := range bDetailDataLi {
		bankAccountCurrencyType := item.(map[string]interface{})
		if fmt.Sprint(bankAccountCurrencyType["currencyTypeId"]) == fmt.Sprint(currencyTypeId) {
			limitsControl := fmt.Sprint(bankAccountCurrencyType["limitsControl"])
			amtOriginalCurrencyBalanceStr := fmt.Sprint(bankAccountCurrencyType["amtOriginalCurrencyBalance"])
			amtOriginalCurrencyBalance := commonUtil.GetFloat64FromString(amtOriginalCurrencyBalanceStr)

			if limitsControl == "1" { // 禁止
				if amtOriginalCurrencyBalance < 0 {
					result := map[string]interface{}{
						"code":    LIMIT_CONTROL_FORBID,
						"message": "帐户:" + fmt.Sprint(bankAccountMaster["name"]) + "出现赤字！",
					}
					return result
				}
			} else if limitsControl == "2" { // 警告
				if amtOriginalCurrencyBalance < 0 {
					result := map[string]interface{}{
						"code":    LIMIT_CONTROL_WARN,
						"message": "帐户:" + fmt.Sprint(bankAccountMaster["name"]) + "出现赤字！",
					}
					return result
				}
			}
			break
		}
	}
	return map[string]interface{}{
		"code": LIMIT_CONTROL_SUCCESS,
	}
}

/**
 * 现金帐户,日记帐,月档过帐
 * @param sessionId
 * @param accountInOutParam 过帐参数对象
 */
func (o AccountInOutService) LogAllCashDeposit(sessionId int, accountInOutParam AccountInOutParam) {
	o.LogCashAccountInOut(sessionId, accountInOutParam)
	o.LogCashDailyInOut(sessionId, accountInOutParam.AccountInOutItemParam, accountInOutParam.DiffDataType)
	o.LogCashMonthInOut(sessionId, accountInOutParam)
}

/**
 * 现金帐户过帐,
 * @param depositLogParam 过帐参数对象
 */
func (o AccountInOutService) LogCashAccountInOut(sessionId int, accountInOutParam AccountInOutParam) {
	session, db := global.GetConnection(sessionId)
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	collectionName := "CashAccount"
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": accountInOutParam.AccountId,
		"A.createUnit": querySupport.GetCreateUnitByUserId(session, userId),
	}
	cashAccountBo, found := querySupport.FindByMapWithSession(session, collectionName, queryMap)
	if !found {
		queryMapByte, err := json.Marshal(&queryMap)
		if err != nil {
			panic(err)
		}
		panic(BusinessError{Message: "现金账户没找到，查询条件为:" + string(queryMapByte)})
	}

	commonUtil := CommonUtil{}
	cashAccount := cashAccountBo["A"].(map[string]interface{})
	if fmt.Sprint(cashAccount["currencyTypeId"]) != fmt.Sprint(accountInOutParam.CurrencyTypeId) {
		panic(BusinessError{Message: "过账现金账户:" + fmt.Sprint(cashAccount["name"]) + "未找到对应币别"})
	}

	amtOriginalCurrencyBalanceStr := fmt.Sprint(cashAccount["amtOriginalCurrencyBalance"])
	amtOriginalCurrencyBalance := commonUtil.GetFloat64FromString(amtOriginalCurrencyBalanceStr)

	amtIncrease := commonUtil.GetFloat64FromString(accountInOutParam.AmtIncrease)
	amtReduce := commonUtil.GetFloat64FromString(accountInOutParam.AmtReduce)

	if accountInOutParam.DiffDataType == ADD || accountInOutParam.DiffDataType == AFTER_UPDATE { // 正过账
		amtOriginalCurrencyBalance += (amtIncrease - amtReduce) // TODO AMT_INCREASE - AMT_REDUCE
	} else if accountInOutParam.DiffDataType == BEFORE_UPDATE || accountInOutParam.DiffDataType == DELETE { // 反过账
		amtOriginalCurrencyBalance -= (amtIncrease - amtReduce) /// TODO AMT_INCREASE - AMT_REDUCE
	}
	cashAccount["amtOriginalCurrencyBalance"] = fmt.Sprint(amtOriginalCurrencyBalance)
	cashAccountBo["A"] = cashAccount
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	_, found = txnManager.Update(txnId, collectionName, cashAccountBo)
	if !found {
		panic(BusinessError{Message: "现金账户更新失败"})
	}
}

/**
 * 现金日记账过帐,全增全删实现
 * @param accountInOutItemParam 日记帐明细业务记录
 * @param diffDataType 差异数据类型
 */
func (o AccountInOutService) LogCashDailyInOut(sessionId int, accountInOutItemParam AccountInOutItemParam, diffDataType int) {
	if diffDataType == ADD {
		o.addCashBankDailyInOut(sessionId, accountInOutItemParam)
	} else if diffDataType == AFTER_UPDATE {
		o.addCashBankDailyInOut(sessionId, accountInOutItemParam)
	} else if diffDataType == BEFORE_UPDATE {
		o.deleteCashBankDailyInOut(sessionId, accountInOutItemParam)
	} else {
		o.deleteCashBankDailyInOut(sessionId, accountInOutItemParam)
	}
}

/**
* 添加日记账明细
 */
func (o AccountInOutService) addCashBankDailyInOut(sessionId int, accountInOutItemParam AccountInOutItemParam) {
	_, db := global.GetConnection(sessionId)
	dataSourceModelId := "AccountInOutItem"
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	masterSeqName := GetMasterSequenceName(dataSource)
	masterSeqId := GetSequenceNo(db, masterSeqName)
	accountInOutItem := accountInOutItemParam.ToMap()
	accountInOutItem["id"] = masterSeqId
	bo := map[string]interface{}{
		"_id": masterSeqId,
		"id":  masterSeqId,
		"A":   accountInOutItem,
	}
	modelTemplateFactory.ConvertDataType(dataSource, &bo)

	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	txnManager.Insert(txnId, collectionName, bo)
}

/**
 * 删除日记帐明细
 * @param sessionId
 * @param accountInOutItemParam 日记帐明细业务参数
 */
func (o AccountInOutService) deleteCashBankDailyInOut(sessionId int, accountInOutItemParam AccountInOutItemParam) {
	_, db := global.GetConnection(sessionId)
	query := map[string]interface{}{
		"A.accountId":      accountInOutItemParam.AccountId,
		"A.currencyTypeId": accountInOutItemParam.CurrencyTypeId,
		"A.accountType":    accountInOutItemParam.AccountType,
		"A.billId":         accountInOutItemParam.BillId,
		"A.billTypeId":     accountInOutItemParam.BillTypeId,
	}
	if accountInOutItemParam.BillDetailId != 0 {
		query["A.billDetailId"] = accountInOutItemParam.BillDetailId
		query["A.billDetailName"] = accountInOutItemParam.BillDetailName
	} else {
		query["A.billDetailId"] = 0 // 主数据集记录
	}
	dataSourceModelId := "AccountInOutItem"
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	_, result := txnManager.RemoveAll(txnId, collectionName, query)
	if !result {
		queryByte, err := json.MarshalIndent(&query, "", "\t")
		if err != nil {
			panic(err)
		}
		panic("删除日记账明细失败，查询语句为:" + string(queryByte))
	}
}

/**
 * 现金月档过帐,
 * 1.月档记录不存在,则给其新增一条
 * 2.正反过帐
 * @param accountInOutParam	参数对象
 */
func (o AccountInOutService) LogCashMonthInOut(sessionId int, accountInOutParam AccountInOutParam) {
	o.logMonthInOut(sessionId, accountInOutParam)
}

/**
 * 现金月档过帐,
 * 1.月档记录不存在,则给其新增一条
 * 2.正反过帐
 * @param accountInOutParam	参数对象
 */
func (o AccountInOutService) logMonthInOut(sessionId int, accountInOutParam AccountInOutParam) {
	session, db := global.GetConnection(sessionId)
	// 参数中加入 currencyTypeId
	commonUtil := CommonUtil{}
	amtIncrease := commonUtil.GetFloat64FromString(accountInOutParam.AmtIncrease)
	amtReduce := commonUtil.GetFloat64FromString(accountInOutParam.AmtReduce)

	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}

	qb := QuerySupport{}
	query := map[string]interface{}{
		"A.accountId":             accountInOutParam.AccountId,
		"A.currencyTypeId":        accountInOutParam.CurrencyTypeId,
		"A.accountType":           accountInOutParam.AccountType,
		"A.accountingPeriodYear":  accountInOutParam.AccountingPeriodYear,
		"A.accountingPeriodMonth": accountInOutParam.AccountingPeriodMonth,
		"A.createUnit": qb.GetCreateUnitByUserId(session, userId),
	}

	dataSourceModelId := "AccountInOut"
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)

	accountInOut, found := qb.FindByMapWithSession(session, collectionName, query)
	if !found {
		accountInOut = o.addFinAccountInOutForCash(sessionId, accountInOutParam)
	} else { // 设置modifyBy等字段
		masterData := accountInOut["A"].(map[string]interface{})
		masterData["modifyBy"] = accountInOutParam.ModifyBy
		masterData["modifyUnit"] = accountInOutParam.ModifyUnit
		masterData["modifyTime"] = accountInOutParam.ModifyTime
		accountInOut["A"] = masterData
	}

	masterData := accountInOut["A"].(map[string]interface{})
	accountInOut["A"] = masterData
	amtIncreaseInDb := commonUtil.GetFloat64FromString(fmt.Sprint(masterData["amtIncrease"]))
	amtReduceInDb := commonUtil.GetFloat64FromString(fmt.Sprint(masterData["amtReduce"]))

	increaseCountInDb := commonUtil.GetIntFromString(fmt.Sprint(masterData["increaseCount"]))
	reduceCountInDb := commonUtil.GetIntFromString(fmt.Sprint(masterData["reduceCount"]))

	if accountInOutParam.DiffDataType == ADD || accountInOutParam.DiffDataType == AFTER_UPDATE {
		amtIncreaseInDb += amtIncrease
		amtReduceInDb += amtReduce

		if amtIncrease > 0 {
			increaseCountInDb += 1
		}
		if amtReduce > 0 {
			reduceCountInDb += 1
		}
	} else {
		amtIncreaseInDb -= amtIncrease
		amtReduceInDb -= amtReduce

		if amtIncrease > 0 {
			increaseCountInDb -= 1
		}
		if amtReduce > 0 {
			reduceCountInDb -= 1
		}
	}
	masterData["amtIncrease"] = fmt.Sprint(amtIncreaseInDb)
	masterData["amtReduce"] = fmt.Sprint(amtReduceInDb)
	masterData["increaseCount"] = increaseCountInDb
	masterData["reduceCount"] = reduceCountInDb

	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	_, found = txnManager.Update(txnId, collectionName, accountInOut)
	if !found {
		panic(BusinessError{Message: "月档更新失败"})
	}
}

/**
 * 添加现金账户会计期汇总(月档)记录
 * @param sessionId
 * @param accountInOutParam 参数对象
 */
func (o AccountInOutService) addFinAccountInOutForCash(sessionId int, accountInOutParam AccountInOutParam) map[string]interface{} {
	accountInOut := map[string]interface{}{
		"accountType":           accountInOutParam.AccountType,
		"accountId":             accountInOutParam.AccountId,
		"currencyTypeId":        accountInOutParam.CurrencyTypeId,
		"exchangeRateShow":      accountInOutParam.ExchangeRateShow,
		"exchangeRate":          accountInOutParam.ExchangeRate,
		"accountingPeriodYear":  accountInOutParam.AccountingPeriodYear,
		"accountingPeriodMonth": accountInOutParam.AccountingPeriodMonth,
		"amtIncrease":           "0",
		"amtReduce":             "0",
		"createBy":              accountInOutParam.CreateBy,
		"createTime":            accountInOutParam.CreateTime,
		"createUnit":            accountInOutParam.CreateUnit,
	}
	_, db := global.GetConnection(sessionId)
	dataSourceModelId := "AccountInOut"
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	masterSeqName := GetMasterSequenceName(dataSource)
	masterSeqId := GetSequenceNo(db, masterSeqName)
	accountInOut["id"] = masterSeqId
	bo := map[string]interface{}{
		"_id": masterSeqId,
		"id":  masterSeqId,
		"A":   accountInOut,
	}
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	return txnManager.Insert(txnId, collectionName, bo)
}

/**
 * 银行帐户,日记帐,月档过帐
 * @param sessionId
 * @param accountInOutParam 过帐参数对象
 */
func (o AccountInOutService) LogAllBankDeposit(sessionId int, accountInOutParam AccountInOutParam) {
	o.LogBankAccountInOut(sessionId, accountInOutParam)
	o.LogBankDailyInOut(sessionId, accountInOutParam.AccountInOutItemParam, accountInOutParam.DiffDataType)
	o.LogBankMonthInOut(sessionId, accountInOutParam)
}

/**
 * 银行帐户过帐,
 * @param sessionId
 * @param accountInOutParam 过帐参数对象
 */
func (o AccountInOutService) LogBankAccountInOut(sessionId int, accountInOutParam AccountInOutParam) {
	// 从银行账户初始化那里面拷贝一份出来,
	session, db := global.GetConnection(sessionId)
	
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	
	collectionName := "BankAccount"
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": accountInOutParam.AccountId,
		"A.createUnit": querySupport.GetCreateUnitByUserId(session, userId),
	}
	bankAccountBo, found := querySupport.FindByMapWithSession(session, collectionName, queryMap)
	if !found {
		queryMapByte, err := json.Marshal(&queryMap)
		if err != nil {
			panic(err)
		}
		panic(BusinessError{Message: "银行账户没找到，查询条件为:" + string(queryMapByte)})
	}
	currencyTypeId := accountInOutParam.CurrencyTypeId
	bDetailDataLi := bankAccountBo["B"].([]interface{})
	commonUtil := CommonUtil{}
	isFound := false
	for i, item := range bDetailDataLi {
		bankAccountCurrencyType := item.(map[string]interface{})
		if fmt.Sprint(bankAccountCurrencyType["currencyTypeId"]) == fmt.Sprint(currencyTypeId) {
			amtOriginalCurrencyBalanceStr := fmt.Sprint(bankAccountCurrencyType["amtOriginalCurrencyBalance"])
			amtOriginalCurrencyBalance := commonUtil.GetFloat64FromString(amtOriginalCurrencyBalanceStr)
			amtIncrease := commonUtil.GetFloat64FromString(accountInOutParam.AmtIncrease)
			amtReduce := commonUtil.GetFloat64FromString(accountInOutParam.AmtReduce)

			if accountInOutParam.DiffDataType == ADD || accountInOutParam.DiffDataType == AFTER_UPDATE { // 正过账
				amtOriginalCurrencyBalance += (amtIncrease - amtReduce)
			} else if accountInOutParam.DiffDataType == BEFORE_UPDATE || accountInOutParam.DiffDataType == DELETE { // 反过账
				amtOriginalCurrencyBalance -= (amtIncrease - amtReduce)
			}
			bankAccountCurrencyType["amtOriginalCurrencyBalance"] = fmt.Sprint(amtOriginalCurrencyBalance)
			bDetailDataLi[i] = bankAccountCurrencyType
			isFound = true
			break
		}
	}
	if !isFound {
		bankAccountMaster := bankAccountBo["A"].(map[string]interface{})
		panic(BusinessError{Message: "过账银行账户:" + fmt.Sprint(bankAccountMaster["name"]) + "未找到对应币别"})
	}

	bankAccountBo["B"] = bDetailDataLi
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	_, found = txnManager.Update(txnId, collectionName, bankAccountBo)
	if !found {
		panic(BusinessError{Message: "银行账户更新失败"})
	}
}

/**
 * 银行日记帐过帐,全增全删实现
 * @param sessionId
 * @param accountInOutItemParam 日记帐明细业务参数
 * @param diffDataType 差异数据类型
 */
func (o AccountInOutService) LogBankDailyInOut(sessionId int, accountInOutItemParam AccountInOutItemParam, diffDataType int) {
	if diffDataType == ADD {
		o.addCashBankDailyInOut(sessionId, accountInOutItemParam)
	} else if diffDataType == AFTER_UPDATE {
		o.addCashBankDailyInOut(sessionId, accountInOutItemParam)
	} else if diffDataType == BEFORE_UPDATE {
		o.deleteCashBankDailyInOut(sessionId, accountInOutItemParam)
	} else {
		o.deleteCashBankDailyInOut(sessionId, accountInOutItemParam)
	}
}

/**
 * 银行月档过帐,
 * 1.月档记录不存在,则给其新增一条
 * 2.正反过帐
 * @param accountInOutParam	参数对象
 */
func (o AccountInOutService) LogBankMonthInOut(sessionId int, accountInOutParam AccountInOutParam) {
	o.logMonthInOut(sessionId, accountInOutParam)
}
