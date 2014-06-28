package controllers

import (
	. "com/papersns/accountinout"
	. "com/papersns/common"
	. "com/papersns/component"
	"com/papersns/global"
	. "com/papersns/model"
	"encoding/json"
	"fmt"
)

func (c AccountInOutDisplay) getDataCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	dataSourceModelId := c.Params.Get("dataSourceModelId")
	formTemplateId := c.Params.Get("formTemplateId")
	jsonData := c.Params.Get("jsonData")

	jsonMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonData), &jsonMap)
	if err != nil {
		panic(err)
	}
	queryMap := jsonMap["A"].(map[string]interface{})

	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)

	accountLi := c.getAccountList(sessionId, queryMap)
	origBalanceLi := c.getOrigBalanceList(sessionId, queryMap)
	increaseReduceBalanceLi := c.getIncreaseReduceBalanceList(sessionId, queryMap)

	dataSetLi := c.mergeAndCalceFinalBalance(accountLi, origBalanceLi, increaseReduceBalanceLi)
	dataSetLi = c.filterEmpty(dataSetLi, queryMap)

	bo := map[string]interface{}{
		"_id": 0,
		"id":  0,
		"A": map[string]interface{}{
			"id": 0,
		},
		"B": dataSetLi,
	}

	//	usedCheck := UsedCheck{}
	//	usedCheckBo := usedCheck.GetFormUsedCheck(sessionId, dataSource, bo)
	usedCheckBo := map[string]interface{}{}

	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	columnModelData := templateManager.GetColumnModelDataForFormTemplate(formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	modelTemplateFactory.ConvertDataType(dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	return ModelRenderVO{
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

// 过滤无余额无发生额不显示,查询结果中，若选项为是，上期结余＝期末结余＝0时，对应帐户+币别在结果中不显示 。
func (c AccountInOutDisplay) filterEmpty(balanceLi []interface{}, queryMap map[string]interface{}) []interface{} {
	if fmt.Sprint(queryMap["displayMode"]) == "1" {
		result := []interface{}{}
		commonUtil := CommonUtil{}
		for _, item := range balanceLi {
			itemMap := item.(map[string]interface{})
			origBalance := commonUtil.GetFloat64FromMap(itemMap, "origBalance")
			finalBalance := commonUtil.GetFloat64FromMap(itemMap, "finalBalance")
			if origBalance != 0 || finalBalance != 0 {
				result = append(result, item)
			}
		}

		return result
	} else {
		return balanceLi
	}
}

/**
@param accountLi
	(cashAccount + bankAccount)
@param origBalanceLi
	amtEarly 期初
@param increaseReduceBalanceLi
	origBalance 最近一个会计期之前的期初
	amtIncrease 本期增加
	amtReduce 本期减少
	increaseCount 增加笔数
	reduceCount 减少笔数
@return [{
	origBalance 最近一个会计期之前的期初
	amtIncrease 本期增加
	amtReduce 本期减少
	finalBalance 期末结余,需要计算出来
	increaseCount 增加笔数
	reduceCount 减少笔数
}, ...]
*/
func (c AccountInOutDisplay) mergeAndCalceFinalBalance(accountLi []interface{}, origBalanceLi []interface{}, increaseReduceBalanceLi []interface{}) []interface{} {
	result := []interface{}{}

	allLi := []interface{}{}
	for _, item := range accountLi {
		allLi = append(allLi, item)
	}
	//	for _, item := range origBalanceLi {
	//		allLi = append(allLi, item)
	//	}
	//	for _, item := range increaseReduceBalanceLi {
	//		allLi = append(allLi, item)
	//	}
	uniqueMap := map[string]interface{}{}
	for _, item := range allLi {
		master := item.(map[string]interface{})
		key := fmt.Sprint(master["accountType"]) + "_" + fmt.Sprint(master["accountId"]) + "_" + fmt.Sprint(master["currencyTypeId"])
		if uniqueMap[key] == nil {
			uniqueMap[key] = map[string]interface{}{
				"id":                        0,
				"accountType":               master["accountType"],
				"accountId":                 master["accountId"],
				"currencyTypeId":            master["currencyTypeId"],
				"bankAccountCurrencyTypeId": 0,
				"origBalance":               "0",
				"amtIncrease":               "0",
				"amtReduce":                 "0",
				"increaseCount":             0,
				"reduceCount":               0,
			}
		}
	}
	// 设置id,bankAccountCurrencyTypeId,
	for _, item := range accountLi {
		master := item.(map[string]interface{})
		key := fmt.Sprint(master["accountType"]) + "_" + fmt.Sprint(master["accountId"]) + "_" + fmt.Sprint(master["currencyTypeId"])
		if uniqueMap[key] != nil {
			data := uniqueMap[key].(map[string]interface{})
			uniqueMap[key] = data

			data["id"] = master["id"]
			data["bankAccountCurrencyTypeId"] = master["bankAccountCurrencyTypeId"]
		}
	}
	commonUtil := CommonUtil{}
	// 累加origBalanceLi中的 amtEarly 字段到 origBalance 中,同时设到期末中,防止本期增加,本期减少没查询到数据时,期末结余被放空
	for _, item := range origBalanceLi {
		master := item.(map[string]interface{})
		key := fmt.Sprint(master["accountType"]) + "_" + fmt.Sprint(master["accountId"]) + "_" + fmt.Sprint(master["currencyTypeId"])
		if uniqueMap[key] != nil {
			data := uniqueMap[key].(map[string]interface{})
			uniqueMap[key] = data

			amtEarly := commonUtil.GetFloat64FromMap(master, "amtEarly")
			origBalance := commonUtil.GetFloat64FromMap(data, "origBalance")

			data["origBalance"] = fmt.Sprint(amtEarly + origBalance)
			data["finalBalance"] = fmt.Sprint(amtEarly + origBalance)
		}
	}
	// 根据 increaseReduceBalanceLi,累加 origBalance 等字段计算出 finalBalance
	for _, item := range increaseReduceBalanceLi {
		master := item.(map[string]interface{})
		key := fmt.Sprint(master["accountType"]) + "_" + fmt.Sprint(master["accountId"]) + "_" + fmt.Sprint(master["currencyTypeId"])
		if uniqueMap[key] != nil {
			data := uniqueMap[key].(map[string]interface{})
			uniqueMap[key] = data

			origBalanceInMaster := commonUtil.GetFloat64FromMap(master, "origBalance")
			amtIncreaseInMaster := commonUtil.GetFloat64FromMap(master, "amtIncrease")
			amtReduceInMaster := commonUtil.GetFloat64FromMap(master, "amtReduce")
			increaseCountInMaster := commonUtil.GetIntFromMap(master, "increaseCount")
			reduceCountInMaster := commonUtil.GetIntFromMap(master, "reduceCount")

			origBalanceInData := commonUtil.GetFloat64FromMap(data, "origBalance")
			amtIncreaseInData := commonUtil.GetFloat64FromMap(data, "amtIncrease")
			amtReduceInData := commonUtil.GetFloat64FromMap(data, "amtReduce")
			increaseCountInData := commonUtil.GetIntFromMap(data, "increaseCount")
			reduceCountInData := commonUtil.GetIntFromMap(data, "reduceCount")

			data["origBalance"] = fmt.Sprint(origBalanceInMaster + origBalanceInData)
			data["amtIncrease"] = fmt.Sprint(amtIncreaseInMaster + amtIncreaseInData)
			data["amtReduce"] = fmt.Sprint(amtReduceInMaster + amtReduceInData)

			data["finalBalance"] = fmt.Sprint(origBalanceInMaster + origBalanceInData + amtIncreaseInMaster + amtIncreaseInData - (amtReduceInMaster + amtReduceInData))

			data["increaseCount"] = fmt.Sprint(increaseCountInMaster + increaseCountInData)
			data["reduceCount"] = fmt.Sprint(reduceCountInMaster + reduceCountInData)
		}
	}
	for _, item := range uniqueMap {
		result = append(result, item)
	}

	return result
}

/**
获取账户,从Cash,BankAccountCurrencyType中,根据查询条件,查询账户出来,
queryMap: 银行账户,现金账户,币别,银行账户属性
*/
func (c AccountInOutDisplay) getAccountList(sessionId int, queryMap map[string]interface{}) []interface{} {
	session, _ := global.GetConnection(sessionId)
	result := []interface{}{}
	commonUtil := CommonUtil{}
	id := 1
	cashAccountIdLi := commonUtil.GetIntLiFromMap(queryMap, "cashAccountId")
	bankAccountIdLi := commonUtil.GetIntLiFromMap(queryMap, "bankAccountId")
	isQueryCashAccount := false
	isQueryBankAccount := false
	if len(cashAccountIdLi) == 0 && len(bankAccountIdLi) == 0 {
		isQueryCashAccount = true
		isQueryBankAccount = true
	} else if len(cashAccountIdLi) == 0 && len(bankAccountIdLi) != 0 {
		isQueryCashAccount = false
		isQueryBankAccount = true
	} else if len(cashAccountIdLi) != 0 && len(bankAccountIdLi) == 0 {
		isQueryCashAccount = true
		isQueryBankAccount = false
	} else if len(cashAccountIdLi) != 0 && len(bankAccountIdLi) != 0 {
		isQueryCashAccount = true
		isQueryBankAccount = true
	}
	if isQueryCashAccount { // 现金账户查询
		cashAccountQuery := map[string]interface{}{}
		//			cashAccountIdLi := commonUtil.GetIntLiFromMap(queryMap, "cashAccountId")
		if len(cashAccountIdLi) > 0 {
			cashAccountQuery["A.id"] = map[string]interface{}{
				"$in": cashAccountIdLi,
			}
		}
		currencyTypeIdLi := commonUtil.GetIntLiFromMap(queryMap, "currencyTypeId")
		if len(currencyTypeIdLi) > 0 {
			cashAccountQuery["A.currencyTypeId"] = map[string]interface{}{
				"$in": currencyTypeIdLi,
			}
		}

		querySupport := QuerySupport{}
		collectionName := "CashAccount"
		pageNo := 1
		pageSize := 1000
		orderBy := ""
		cashAccountResult := querySupport.IndexWithSession(session, collectionName, cashAccountQuery, pageNo, pageSize, orderBy)
		cashAccountItems := cashAccountResult["items"].([]interface{})
		for _, item := range cashAccountItems {
			id += 1
			itemMap := item.(map[string]interface{})
			master := itemMap["A"].(map[string]interface{})
			result = append(result, map[string]interface{}{
				"id":                        id,
				"accountType":               1,
				"accountId":                 master["id"],
				"bankAccountCurrencyTypeId": 0,
				"currencyTypeId":            master["currencyTypeId"],
			})
		}
	}
	if isQueryBankAccount { // 银行账户查询
		bankAccountQuery := map[string]interface{}{}

		if len(bankAccountIdLi) > 0 {
			bankAccountQuery["A.bankAccountId"] = map[string]interface{}{
				"$in": bankAccountIdLi,
			}
		}
		currencyTypeIdLi := commonUtil.GetIntLiFromMap(queryMap, "currencyTypeId")
		if len(currencyTypeIdLi) > 0 {
			bankAccountQuery["A.currencyTypeId"] = map[string]interface{}{
				"$in": currencyTypeIdLi,
			}
		}
		property := commonUtil.GetIntFromMap(queryMap, "property")
		if property > 0 {
			bankAccountQuery["A.accountProperty"] = property
		}

		querySupport := QuerySupport{}
		collectionName := "BankAccountCurrencyType"
		pageNo := 1
		pageSize := 1000
		orderBy := ""
		bankAccountResult := querySupport.IndexWithSession(session, collectionName, bankAccountQuery, pageNo, pageSize, orderBy)
		bankAccountItems := bankAccountResult["items"].([]interface{})
		for _, item := range bankAccountItems {
			id += 1
			itemMap := item.(map[string]interface{})
			master := itemMap["A"].(map[string]interface{})
			result = append(result, map[string]interface{}{
				"id":                        id,
				"accountType":               2,
				"accountId":                 master["bankAccountId"],
				"bankAccountCurrencyTypeId": master["id"],
				"currencyTypeId":            master["currencyTypeId"],
			})
		}
	}

	return result
}

/**
获取账户的期初数据
*/
func (c AccountInOutDisplay) getOrigBalanceList(sessionId int, queryMap map[string]interface{}) []interface{} {
	session, _ := global.GetConnection(sessionId)
	result := []interface{}{}
	accountInitQuery := map[string]interface{}{}
	orQuery := []interface{}{}
	commonUtil := CommonUtil{}
	if queryMap["cashAccountId"] != nil {
		cashAccountIdLi := commonUtil.GetIntLiFromMap(queryMap, "cashAccountId")
		if len(cashAccountIdLi) > 0 {
			orQuery = append(orQuery, map[string]interface{}{
				"A.accountId": map[string]interface{}{
					"$in": cashAccountIdLi,
				},
				"A.accountType": 1,
			})
		}
	}
	if queryMap["bankAccountId"] != nil {
		bankAccountIdLi := commonUtil.GetIntLiFromMap(queryMap, "bankAccountId")
		if len(bankAccountIdLi) > 0 {
			orQuery = append(orQuery, map[string]interface{}{
				"A.accountId": map[string]interface{}{
					"$in": bankAccountIdLi,
				},
				"A.accountType": 2,
			})
		}
	}
	if len(orQuery) > 0 {
		accountInitQuery["$or"] = orQuery
	}
	if queryMap["currencyTypeId"] != nil {
		currencyTypeIdLi := commonUtil.GetIntLiFromMap(queryMap, "currencyTypeId")
		if len(currencyTypeIdLi) > 0 {
			accountInitQuery["A.currencyTypeId"] = map[string]interface{}{
				"$in": currencyTypeIdLi,
			}
		}
	}

	querySupport := QuerySupport{}
	collectionName := "AccountInit"
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	queryResult := querySupport.IndexWithSession(session, collectionName, accountInitQuery, pageNo, pageSize, orderBy)
	queryResultItems := queryResult["items"].([]interface{})
	for _, item := range queryResultItems {
		itemMap := item.(map[string]interface{})
		master := itemMap["A"].(map[string]interface{})
		result = append(result, map[string]interface{}{
			"accountType":               master["accountType"],
			"accountId":                 master["accountId"],
			"bankAccountCurrencyTypeId": master["bankAccountCurrencyTypeId"],
			"currencyTypeId":            master["currencyTypeId"],
			"amtEarly":                  master["amtEarly"],
		})
	}
	return result
}

/**
获取本期增加,本期减少
*/
func (c AccountInOutDisplay) getIncreaseReduceBalanceList(sessionId int, queryMap map[string]interface{}) []interface{} {
	result := []interface{}{}
	commonUtil := CommonUtil{}
	if fmt.Sprint(queryMap["queryMode"]) == "1" { // 按日期查询
		billDateBegin := commonUtil.GetIntFromMap(queryMap, "billDateBegin")
		billDateEnd := commonUtil.GetIntFromMap(queryMap, "billDateEnd")
		if billDateBegin == 0 {
			billDateBegin = 11110101
		}
		if billDateEnd == 0 {
			billDateEnd = 99991212
		}
		lastAccountingYear, lastSequenceNo := c.getLastAccountingYearSequenceNo(sessionId, billDateBegin)
		// 当为0时,如何处理,TODO,
		// 计算期初和本期增加,本期减少,
		origBalanceLi := c.getOrigBalanceFromAccountInOut(sessionId, queryMap, lastAccountingYear, lastSequenceNo)
		// 累加最近一个会计期结束日期+1 至 billDateBegin 之间的本期增加,本期减少到期初,
		{
			var lastEndDate int
			if lastAccountingYear > 0 && lastSequenceNo > 0 {
				accountInOutService := AccountInOutService{}
				_, lastEndDate = accountInOutService.GetAccountingPeriodStartEndDate(sessionId, lastAccountingYear, lastSequenceNo)
			} else {// 没有找到最近的一个会计期时,将lastEndDate置为19700101
				lastEndDate = 19700101
			}
			dateUtil := DateUtil{}
			nextOfLastEndDate := dateUtil.GetNextDate(lastEndDate)
			preOfBillDateBegin := dateUtil.GetPreDate(billDateBegin)
			amtIncreaseReduceLi := c.getAmtIncreaseReduceByDate(sessionId, queryMap, nextOfLastEndDate, preOfBillDateBegin)
			origBalanceLi = c.addIncreaseReduceToOrigBalance(origBalanceLi, amtIncreaseReduceLi)
		}
		amtIncreaseReduceLi := c.getAmtIncreaseReduceByDate(sessionId, queryMap, billDateBegin, billDateEnd)
		// 合并期初,本期增加,本期减少
		result = c.mergeOrigBalanceAndIncreaseReduce(origBalanceLi, amtIncreaseReduceLi)
	} else if fmt.Sprint(queryMap["queryMode"]) == "2" { // 按期间查询
		accountingYearStart := commonUtil.GetIntFromMap(queryMap, "accountingYearStart")
		accountingYearEnd := commonUtil.GetIntFromMap(queryMap, "accountingYearEnd")
		accountingPeriodStart := commonUtil.GetIntFromMap(queryMap, "accountingPeriodStart")
		accountingPeriodEnd := commonUtil.GetIntFromMap(queryMap, "accountingPeriodEnd")
		if accountingYearStart == 0 {
			accountingYearStart = 1111
		}
		if accountingYearEnd == 0 {
			accountingYearEnd = 9999
		}
		if accountingPeriodEnd == 0 {
			accountingPeriodEnd = 100
		}
		lastAccountingYear, lastSequenceNo := c.getLastAccountingYearSequenceNoByYearMonth(sessionId, accountingYearStart, accountingPeriodStart)

		// 计算期初和本期增加,本期减少,
		origBalanceLi := c.getOrigBalanceFromAccountInOut(sessionId, queryMap, lastAccountingYear, lastSequenceNo)
		amtIncreaseReduceLi := c.getAmtIncreaseReduceByYearMonth(sessionId, queryMap, accountingYearStart, accountingPeriodStart, accountingYearEnd, accountingPeriodEnd)
		// 合并期初,本期增加,本期减少
		result = c.mergeOrigBalanceAndIncreaseReduce(origBalanceLi, amtIncreaseReduceLi)
	}
	return result
}

func (c AccountInOutDisplay) addIncreaseReduceToOrigBalance(origBalanceLi []interface{}, increaseReduceLi []interface{}) []interface{} {
	allLi := []interface{}{}
	for _, item := range origBalanceLi {
		allLi = append(allLi, item)
	}
	for _, item := range increaseReduceLi {
		allLi = append(allLi, item)
	}
	uniqueMap := map[string]interface{}{}
	commonUtil := CommonUtil{}
	for _, item := range allLi {
		master := item.(map[string]interface{})
		key := fmt.Sprint(master["accountType"]) + "_" + fmt.Sprint(master["accountId"]) + "_" + fmt.Sprint(master["currencyTypeId"])
		if uniqueMap[key] == nil {
			uniqueMap[key] = map[string]interface{}{
				"accountType":    master["accountType"],
				"accountId":      master["accountId"],
				"currencyTypeId": master["currencyTypeId"],
				"origBalance":    "0",
			}
		}
		data := uniqueMap[key].(map[string]interface{})
		origBalanceInData := commonUtil.GetFloat64FromMap(data, "origBalance")

		origBalance := commonUtil.GetFloat64FromMap(master, "origBalance")
		amtIncrease := commonUtil.GetFloat64FromMap(master, "amtIncrease")
		amtReduce := commonUtil.GetFloat64FromMap(master, "amtReduce")

		data["origBalance"] = fmt.Sprint(origBalanceInData + origBalance + amtIncrease - amtReduce)

		uniqueMap[key] = data
	}
	result := []interface{}{}
	for _, item := range uniqueMap {
		result = append(result, item)
	}
	return result
}

func (c AccountInOutDisplay) mergeOrigBalanceAndIncreaseReduce(origBalanceLi []interface{}, increaseReduceLi []interface{}) []interface{} {
	allLi := []interface{}{}
	for _, item := range origBalanceLi {
		allLi = append(allLi, item)
	}
	for _, item := range increaseReduceLi {
		allLi = append(allLi, item)
	}
	uniqueMap := map[string]interface{}{}
	commonUtil := CommonUtil{}
	for _, item := range allLi {
		master := item.(map[string]interface{})
		key := fmt.Sprint(master["accountType"]) + "_" + fmt.Sprint(master["accountId"]) + "_" + fmt.Sprint(master["currencyTypeId"])
		if uniqueMap[key] == nil {
			uniqueMap[key] = map[string]interface{}{
				"accountType":    master["accountType"],
				"accountId":      master["accountId"],
				"currencyTypeId": master["currencyTypeId"],
				"origBalance":    "0",
				"amtIncrease":    "0",
				"amtReduce":      "0",
				"increaseCount":  0,
				"reduceCount":    0,
			}
		}
		data := uniqueMap[key].(map[string]interface{})
		origBalanceInData := commonUtil.GetFloat64FromMap(data, "origBalance")
		amtIncreaseInData := commonUtil.GetFloat64FromMap(data, "amtIncrease")
		amtReduceInData := commonUtil.GetFloat64FromMap(data, "amtReduce")
		increaseCountInData := commonUtil.GetIntFromMap(data, "increaseCount")
		reduceCountInData := commonUtil.GetIntFromMap(data, "reduceCount")

		origBalance := commonUtil.GetFloat64FromMap(master, "origBalance")
		amtIncrease := commonUtil.GetFloat64FromMap(master, "amtIncrease")
		amtReduce := commonUtil.GetFloat64FromMap(master, "amtReduce")
		increaseCount := commonUtil.GetIntFromMap(master, "increaseCount")
		reduceCount := commonUtil.GetIntFromMap(master, "reduceCount")

		data["origBalance"] = fmt.Sprint(origBalanceInData + origBalance)
		data["amtIncrease"] = fmt.Sprint(amtIncreaseInData + amtIncrease)
		data["amtReduce"] = fmt.Sprint(amtReduceInData + amtReduce)
		data["increaseCount"] = increaseCountInData + increaseCount
		data["reduceCount"] = reduceCountInData + reduceCount

		uniqueMap[key] = data
	}
	result := []interface{}{}
	for _, item := range uniqueMap {
		result = append(result, item)
	}
	return result
}

func (c AccountInOutDisplay) getOrigBalanceFromAccountInOut(sessionId int, queryMap map[string]interface{}, lastAccountingYear int, lastSequenceNo int) []interface{} {
	session, _ := global.GetConnection(sessionId)
	result := []interface{}{}
	accountInitQuery := map[string]interface{}{}
	orQuery := []interface{}{}
	commonUtil := CommonUtil{}
	if queryMap["cashAccountId"] != nil {
		cashAccountIdLi := commonUtil.GetIntLiFromMap(queryMap, "cashAccountId")
		if len(cashAccountIdLi) > 0 {
			orQuery = append(orQuery, map[string]interface{}{
				"A.accountId": map[string]interface{}{
					"$in": cashAccountIdLi,
				},
				"A.accountType": 1,
			})
		}
	}
	if queryMap["bankAccountId"] != nil {
		bankAccountIdLi := commonUtil.GetIntLiFromMap(queryMap, "bankAccountId")
		if len(bankAccountIdLi) > 0 {
			orQuery = append(orQuery, map[string]interface{}{
				"A.accountId": map[string]interface{}{
					"$in": bankAccountIdLi,
				},
				"A.accountType": 2,
			})
		}
	}
	yearMonthOrQuery := []interface{}{}
	yearMonthOrQuery = append(yearMonthOrQuery, map[string]interface{}{
		"A.accountingPeriodYear": lastAccountingYear,
		"A.accountingPeriodMonth": map[string]interface{}{
			"$lte": lastSequenceNo,
		},
	})
	yearMonthOrQuery = append(yearMonthOrQuery, map[string]interface{}{
		"A.accountingPeriodYear": map[string]interface{}{
			"$lt": lastAccountingYear,
		},
	})
	if len(orQuery) > 0 {
		accountInitQuery["$and"] = []interface{}{
			map[string]interface{}{
				"$or": orQuery,
			},
			map[string]interface{}{
				"$or": yearMonthOrQuery,
			},
		}
	} else {
		accountInitQuery["$or"] = yearMonthOrQuery
	}
	if queryMap["currencyTypeId"] != nil {
		currencyTypeIdLi := commonUtil.GetIntLiFromMap(queryMap, "currencyTypeId")
		if len(currencyTypeIdLi) > 0 {
			accountInitQuery["A.currencyTypeId"] = map[string]interface{}{
				"$in": currencyTypeIdLi,
			}
		}
	}

	querySupport := QuerySupport{}
	collectionName := "AccountInOut"
	pageNo := 1
	pageSize := 10000
	orderBy := ""
	queryResult := querySupport.IndexWithSession(session, collectionName, accountInitQuery, pageNo, pageSize, orderBy)
	queryResultItems := queryResult["items"].([]interface{})
	uniqueMap := map[string]interface{}{}
	for _, item := range queryResultItems {
		itemMap := item.(map[string]interface{})
		master := itemMap["A"].(map[string]interface{})
		key := fmt.Sprint(master["accountType"]) + "_" + fmt.Sprint(master["accountId"]) + "_" + fmt.Sprint(master["currencyTypeId"])
		if uniqueMap[key] == nil {
			uniqueMap[key] = map[string]interface{}{
				"accountType":    master["accountType"],
				"accountId":      master["accountId"],
				"currencyTypeId": master["currencyTypeId"],
				"origBalance":    "0",
			}
		}
		data := uniqueMap[key].(map[string]interface{})
		origBalance := commonUtil.GetFloat64FromString(fmt.Sprint(data["origBalance"]))
		amtIncrease := commonUtil.GetFloat64FromString(fmt.Sprint(master["amtIncrease"]))
		amtReduce := commonUtil.GetFloat64FromString(fmt.Sprint(master["amtReduce"]))
		data["origBalance"] = fmt.Sprint(origBalance + amtIncrease - amtReduce)
		uniqueMap[key] = data
	}
	for _, item := range uniqueMap {
		result = append(result, item)
	}
	return result
}

func (c AccountInOutDisplay) getAmtIncreaseReduceByDate(sessionId int, queryMap map[string]interface{}, billDateBegin int, billDateEnd int) []interface{} {
	session, _ := global.GetConnection(sessionId)
	result := []interface{}{}
	accountInitQuery := map[string]interface{}{}
	orQuery := []interface{}{}
	commonUtil := CommonUtil{}
	if queryMap["cashAccountId"] != nil {
		cashAccountIdLi := commonUtil.GetIntLiFromMap(queryMap, "cashAccountId")
		if len(cashAccountIdLi) > 0 {
			orQuery = append(orQuery, map[string]interface{}{
				"A.accountId": map[string]interface{}{
					"$in": cashAccountIdLi,
				},
				"A.accountType": 1,
			})
		}
	}
	if queryMap["bankAccountId"] != nil {
		bankAccountIdLi := commonUtil.GetIntLiFromMap(queryMap, "bankAccountId")
		if len(bankAccountIdLi) > 0 {
			orQuery = append(orQuery, map[string]interface{}{
				"A.accountId": map[string]interface{}{
					"$in": bankAccountIdLi,
				},
				"A.accountType": 2,
			})
		}
	}
	yearMonthOrQuery := map[string]interface{}{
		"$and": []interface{}{
			map[string]interface{}{
				"A.billDate": map[string]interface{}{
					"$gte": billDateBegin,
				},
			},
			map[string]interface{}{
				"A.billDate": map[string]interface{}{
					"$lte": billDateEnd,
				},
			},
		},
	}
	if len(orQuery) > 0 {
		accountInitQuery["$and"] = []interface{}{
			map[string]interface{}{
				"$or": orQuery,
			},
			yearMonthOrQuery,
		}
	} else {
		for k, v := range yearMonthOrQuery {
			accountInitQuery[k] = v
		}
	}
	if queryMap["currencyTypeId"] != nil {
		currencyTypeIdLi := commonUtil.GetIntLiFromMap(queryMap, "currencyTypeId")
		if len(currencyTypeIdLi) > 0 {
			accountInitQuery["A.currencyTypeId"] = map[string]interface{}{
				"$in": currencyTypeIdLi,
			}
		}
	}

	querySupport := QuerySupport{}
	collectionName := "AccountInOutItem"
	pageNo := 1
	pageSize := 10000
	orderBy := ""
	queryResult := querySupport.IndexWithSession(session, collectionName, accountInitQuery, pageNo, pageSize, orderBy)
	queryResultItems := queryResult["items"].([]interface{})
	uniqueMap := map[string]interface{}{}
	for _, item := range queryResultItems {
		itemMap := item.(map[string]interface{})
		master := itemMap["A"].(map[string]interface{})
		key := fmt.Sprint(master["accountType"]) + "_" + fmt.Sprint(master["accountId"]) + "_" + fmt.Sprint(master["currencyTypeId"])
		if uniqueMap[key] == nil {
			uniqueMap[key] = map[string]interface{}{
				"accountType":    master["accountType"],
				"accountId":      master["accountId"],
				"currencyTypeId": master["currencyTypeId"],
				"amtIncrease":    "0",
				"amtReduce":      "0",
				"increaseCount":  0,
				"reduceCount":    0,
			}
		}
		data := uniqueMap[key].(map[string]interface{})
		amtIncreaseInData := commonUtil.GetFloat64FromString(fmt.Sprint(data["amtIncrease"]))
		amtReduceInData := commonUtil.GetFloat64FromString(fmt.Sprint(data["amtReduce"]))
		increaseCountInData := commonUtil.GetIntFromMap(data, "increaseCount")
		reduceCountInData := commonUtil.GetIntFromMap(data, "reduceCount")

		amtIncrease := commonUtil.GetFloat64FromString(fmt.Sprint(master["amtIncrease"]))
		amtReduce := commonUtil.GetFloat64FromString(fmt.Sprint(master["amtReduce"]))
		data["amtIncrease"] = fmt.Sprint(amtIncreaseInData + amtIncrease)
		data["amtReduce"] = fmt.Sprint(amtReduceInData + amtReduce)
		if amtIncrease > 0 {
			data["increaseCount"] = increaseCountInData + 1
		}
		if amtReduce > 0 {
			data["reduceCount"] = reduceCountInData + 1
		}
		uniqueMap[key] = data
	}
	for _, item := range uniqueMap {
		result = append(result, item)
	}
	return result
}

func (c AccountInOutDisplay) getAmtIncreaseReduceByYearMonth(sessionId int, queryMap map[string]interface{}, accountingYearStart int, accountingPeriodStart int, accountingYearEnd int, accountingPeriodEnd int) []interface{} {
	session, _ := global.GetConnection(sessionId)
	result := []interface{}{}
	accountInitQuery := map[string]interface{}{}
	orQuery := []interface{}{}
	commonUtil := CommonUtil{}
	if queryMap["cashAccountId"] != nil {
		cashAccountIdLi := commonUtil.GetIntLiFromMap(queryMap, "cashAccountId")
		if len(cashAccountIdLi) > 0 {
			orQuery = append(orQuery, map[string]interface{}{
				"A.accountId": map[string]interface{}{
					"$in": cashAccountIdLi,
				},
				"A.accountType": 1,
			})
		}
	}
	if queryMap["bankAccountId"] != nil {
		bankAccountIdLi := commonUtil.GetIntLiFromMap(queryMap, "bankAccountId")
		if len(bankAccountIdLi) > 0 {
			orQuery = append(orQuery, map[string]interface{}{
				"A.accountId": map[string]interface{}{
					"$in": bankAccountIdLi,
				},
				"A.accountType": 2,
			})
		}
	}
	// 年月的东东,
	yearMonthOrQuery := map[string]interface{}{
		"$and": []interface{}{
			map[string]interface{}{
				"$or": []interface{}{
					map[string]interface{}{
						"A.accountingPeriodYear": accountingYearStart,
						"A.accountingPeriodMonth": map[string]interface{}{
							"$gte": accountingPeriodStart,
						},
					},
					map[string]interface{}{
						"A.accountingPeriodYear": map[string]interface{}{
							"$gt": accountingYearStart,
						},
					},
				},
			},
			map[string]interface{}{
				"$or": []interface{}{
					map[string]interface{}{
						"A.accountingPeriodYear": accountingYearEnd,
						"A.accountingPeriodMonth": map[string]interface{}{
							"$lte": accountingPeriodEnd,
						},
					},
					map[string]interface{}{
						"A.accountingPeriodYear": map[string]interface{}{
							"$lt": accountingYearEnd,
						},
					},
				},
			},
		},
	}
	if len(orQuery) > 0 {
		accountInitQuery["$and"] = []interface{}{
			map[string]interface{}{
				"$or": orQuery,
			},
			yearMonthOrQuery,
		}
	} else {
		for k, v := range yearMonthOrQuery {
			accountInitQuery[k] = v
		}
	}
	if queryMap["currencyTypeId"] != nil {
		currencyTypeIdLi := commonUtil.GetIntLiFromMap(queryMap, "currencyTypeId")
		if len(currencyTypeIdLi) > 0 {
			accountInitQuery["A.currencyTypeId"] = map[string]interface{}{
				"$in": currencyTypeIdLi,
			}
		}
	}

	querySupport := QuerySupport{}
	collectionName := "AccountInOut"
	pageNo := 1
	pageSize := 10000
	orderBy := ""
	queryResult := querySupport.IndexWithSession(session, collectionName, accountInitQuery, pageNo, pageSize, orderBy)
	queryResultItems := queryResult["items"].([]interface{})
	uniqueMap := map[string]interface{}{}
	for _, item := range queryResultItems {
		itemMap := item.(map[string]interface{})
		master := itemMap["A"].(map[string]interface{})
		key := fmt.Sprint(master["accountType"]) + "_" + fmt.Sprint(master["accountId"]) + "_" + fmt.Sprint(master["currencyTypeId"])
		if uniqueMap[key] == nil {
			uniqueMap[key] = map[string]interface{}{
				"accountType":    master["accountType"],
				"accountId":      master["accountId"],
				"currencyTypeId": master["currencyTypeId"],
				"amtIncrease":    "0",
				"amtReduce":      "0",
				"increaseCount":  0,
				"reduceCount":    0,
			}
		}
		data := uniqueMap[key].(map[string]interface{})
		amtIncreaseInData := commonUtil.GetFloat64FromString(fmt.Sprint(data["amtIncrease"]))
		amtReduceInData := commonUtil.GetFloat64FromString(fmt.Sprint(data["amtReduce"]))
		increaseCountInData := commonUtil.GetIntFromMap(data, "increaseCount")
		reduceCountInData := commonUtil.GetIntFromMap(data, "reduceCount")

		amtIncrease := commonUtil.GetFloat64FromString(fmt.Sprint(master["amtIncrease"]))
		amtReduce := commonUtil.GetFloat64FromString(fmt.Sprint(master["amtReduce"]))
		increaseCount := commonUtil.GetIntFromMap(master, "increaseCount")
		reduceCount := commonUtil.GetIntFromMap(master, "reduceCount")

		data["amtIncrease"] = fmt.Sprint(amtIncreaseInData + amtIncrease)
		data["amtReduce"] = fmt.Sprint(amtReduceInData + amtReduce)
		data["increaseCount"] = increaseCountInData + increaseCount
		data["reduceCount"] = reduceCountInData + reduceCount
		uniqueMap[key] = data
	}
	for _, item := range uniqueMap {
		result = append(result, item)
	}
	return result
}

/**
取得max(会计期结束日期<date)的会计期年月
@param date yyyyMMdd
*/
func (c AccountInOutDisplay) getLastAccountingYearSequenceNoByYearMonth(sessionId int, accountingYear int, sequenceNo int) (int, int) {
	session, _ := global.GetConnection(sessionId)
	queryMap := map[string]interface{}{
		"$or": []interface{}{
			map[string]interface{}{
				"A.accountingYear": accountingYear,
				"B.sequenceNo": map[string]interface{}{
					"$lt": sequenceNo,
				},
			},
			map[string]interface{}{
				"A.accountingYear": map[string]interface{}{
					"$lt": accountingYear,
				},
			},
		},
	}
	querySupport := QuerySupport{}
	collectionName := "AccountingPeriod"
	pageNo := 1
	pageSize := 1
	orderBy := "-A.accountingYear"
	queryResult := querySupport.IndexWithSession(session, collectionName, queryMap, pageNo, pageSize, orderBy)
	items := queryResult["items"].([]interface{})
	if len(items) > 0 {
		commonUtil := CommonUtil{}
		accountingPeriod := items[0].(map[string]interface{})
		master := accountingPeriod["A"].(map[string]interface{})
		accountingYearToReturn := commonUtil.GetIntFromMap(master, "accountingYear")
		sequenceNoToReturn := 0
		detailB := accountingPeriod["B"].([]interface{})
		detailLength := len(detailB)
		i := detailLength - 1
		for i >= 0 {
			line := detailB[i].(map[string]interface{})
			sequenceNoInDB := commonUtil.GetIntFromMap(line, "sequenceNo")

			if accountingYearToReturn < accountingYear {
				sequenceNoToReturn = sequenceNoInDB
				break
			}

			if sequenceNoInDB < sequenceNo {
				sequenceNoToReturn = sequenceNoInDB
				break
			}
			i -= 1
		}
		return accountingYearToReturn, sequenceNoToReturn
	} else {
		return 0, 0
	}
}

/**
取得max(会计期结束日期<date)的会计期年月
@param date yyyyMMdd
*/
func (c AccountInOutDisplay) getLastAccountingYearSequenceNo(sessionId int, date int) (int, int) {
	session, _ := global.GetConnection(sessionId)
	queryMap := map[string]interface{}{
		"B.endDate": map[string]interface{}{
			"$lt": date,
		},
	}
	querySupport := QuerySupport{}
	collectionName := "AccountingPeriod"
	pageNo := 1
	pageSize := 1
	orderBy := "-A.accountingYear"
	queryResult := querySupport.IndexWithSession(session, collectionName, queryMap, pageNo, pageSize, orderBy)
	items := queryResult["items"].([]interface{})
	if len(items) > 0 {
		commonUtil := CommonUtil{}
		accountingPeriod := items[0].(map[string]interface{})
		master := accountingPeriod["A"].(map[string]interface{})
		accountingYear := commonUtil.GetIntFromMap(master, "accountingYear")
		sequenceNo := 0
		detailB := accountingPeriod["B"].([]interface{})
		detailLength := len(detailB)
		i := detailLength - 1
		for i >= 0 {
			line := detailB[i].(map[string]interface{})
			endDate := commonUtil.GetIntFromMap(line, "endDate")
			if endDate < date {
				sequenceNo = commonUtil.GetIntFromMap(line, "sequenceNo")
				break
			}
			i -= 1
		}
		return accountingYear, sequenceNo
	} else {
		return 0, 0
	}
}
