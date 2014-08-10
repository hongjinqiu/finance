package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	. "com/papersns/accountinout"
	. "com/papersns/common"
	. "com/papersns/component"
	"com/papersns/global"
	. "com/papersns/model"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func init() {
}

type AccountInOutDisplaySupport struct {
	ActionSupport
}

func (o AccountInOutDisplaySupport) RAfterNewData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	masterData := (*bo)["A"].(map[string]interface{})
	(*bo)["A"] = masterData

	session, _ := global.GetConnection(sessionId)
	qb := QuerySupport{}
	query := map[string]interface{}{
		"A.code": "RMB",
	}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		query[k] = v
	}

	collectionName := "CurrencyType"
	{
		queryByte, err := json.MarshalIndent(&query, "", "\t")
		if err != nil {
			panic(err)
		}
		log.Println("RAfterNewData,collectionName:" + collectionName + ", query:" + string(queryByte))
	}
	result, found := qb.FindByMapWithSession(session, collectionName, query)
	if found {
		masterData["currencyTypeId"] = result["id"]
	}
}

type AccountInOutDisplay struct {
	BaseDataAction
}

func (c AccountInOutDisplay) RGetDataCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)
	defer c.RRollbackTxn(sessionId)

	userId, err := strconv.Atoi(c.Session["userId"])
	if err != nil {
		panic(err)
	}

	dataSourceModelId := c.Params.Get("dataSourceModelId")
	formTemplateId := c.Params.Get("formTemplateId")
	jsonData := c.Params.Get("jsonData")

	jsonMap := map[string]interface{}{}
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	if err != nil {
		panic(err)
	}
	queryMap := jsonMap["A"].(map[string]interface{})

	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	accountLi := c.getAccountList(sessionId, formTemplate, queryMap)
	origBalanceLi := c.getOrigBalanceList(sessionId, formTemplate, queryMap)
	increaseReduceBalanceLi := c.getIncreaseReduceBalanceList(sessionId, formTemplate, queryMap)

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

	columnModelData := templateManager.GetColumnModelDataForFormTemplate(sessionId, formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	modelTemplateFactory.ConvertDataType(dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	return ModelRenderVO{
		UserId:      userId,
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
	mathUtil := MathUtil{}
	// 累加origBalanceLi中的 amtEarly 字段到 origBalance 中,同时设到期末中,防止本期增加,本期减少没查询到数据时,期末结余被放空
	for _, item := range origBalanceLi {
		master := item.(map[string]interface{})
		key := fmt.Sprint(master["accountType"]) + "_" + fmt.Sprint(master["accountId"]) + "_" + fmt.Sprint(master["currencyTypeId"])
		if uniqueMap[key] != nil {
			data := uniqueMap[key].(map[string]interface{})
			uniqueMap[key] = data

			amtEarly := fmt.Sprint(master["amtEarly"])
			origBalance := fmt.Sprint(data["origBalance"])

			data["origBalance"] = mathUtil.Add(amtEarly, origBalance)
			data["finalBalance"] = mathUtil.Add(amtEarly, origBalance)
		}
	}
	// 根据 increaseReduceBalanceLi,累加 origBalance 等字段计算出 finalBalance
	for _, item := range increaseReduceBalanceLi {
		master := item.(map[string]interface{})
		key := fmt.Sprint(master["accountType"]) + "_" + fmt.Sprint(master["accountId"]) + "_" + fmt.Sprint(master["currencyTypeId"])
		if uniqueMap[key] != nil {
			data := uniqueMap[key].(map[string]interface{})
			uniqueMap[key] = data

			origBalanceInMaster := fmt.Sprint(master["origBalance"])
			amtIncreaseInMaster := fmt.Sprint(master["amtIncrease"])
			amtReduceInMaster := fmt.Sprint(master["amtReduce"])
			increaseCountInMaster := fmt.Sprint(master["increaseCount"])
			reduceCountInMaster := fmt.Sprint(master["reduceCount"])

			origBalanceInData := fmt.Sprint(data["origBalance"])
			amtIncreaseInData := fmt.Sprint(data["amtIncrease"])
			amtReduceInData := fmt.Sprint(data["amtReduce"])
			increaseCountInData := fmt.Sprint(data["increaseCount"])
			reduceCountInData := fmt.Sprint(data["reduceCount"])

			data["origBalance"] = mathUtil.Add(origBalanceInMaster, origBalanceInData)
			data["amtIncrease"] = mathUtil.Add(amtIncreaseInMaster, amtIncreaseInData)
			data["amtReduce"] = mathUtil.Add(amtReduceInMaster, amtReduceInData)

			finalBalance := mathUtil.Add(origBalanceInMaster, origBalanceInData)
			finalBalance = mathUtil.Add(finalBalance, amtIncreaseInMaster)
			finalBalance = mathUtil.Add(finalBalance, amtIncreaseInData)
			finalBalance = mathUtil.Sub(finalBalance, amtReduceInMaster)
			finalBalance = mathUtil.Sub(finalBalance, amtReduceInData)
			data["finalBalance"] = finalBalance

			data["increaseCount"] = mathUtil.Add(increaseCountInMaster, increaseCountInData)
			data["reduceCount"] = mathUtil.Add(reduceCountInMaster, reduceCountInData)
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
func (c AccountInOutDisplay) getAccountList(sessionId int, formTemplate FormTemplate, queryMap map[string]interface{}) []interface{} {
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
		querySupport := QuerySupport{}
		cashAccountQuery := map[string]interface{}{}
		permissionSupport := PermissionSupport{}
		permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
		for k, v := range permissionQueryDict {
			cashAccountQuery[k] = v
		}
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
		querySupport := QuerySupport{}
		bankAccountQuery := map[string]interface{}{}
		permissionSupport := PermissionSupport{}
		permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
		for k, v := range permissionQueryDict {
			bankAccountQuery[k] = v
		}

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
func (c AccountInOutDisplay) getOrigBalanceList(sessionId int, formTemplate FormTemplate, queryMap map[string]interface{}) []interface{} {
	session, _ := global.GetConnection(sessionId)
	result := []interface{}{}
	querySupport := QuerySupport{}
	accountInitQuery := map[string]interface{}{}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		accountInitQuery[k] = v
	}
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
func (c AccountInOutDisplay) getIncreaseReduceBalanceList(sessionId int, formTemplate FormTemplate, queryMap map[string]interface{}) []interface{} {
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
		lastAccountingYear, lastSequenceNo := c.getLastAccountingYearSequenceNo(sessionId, formTemplate, billDateBegin)
		// 当为0时,如何处理,TODO,
		// 计算期初和本期增加,本期减少,
		origBalanceLi := c.getOrigBalanceFromAccountInOut(sessionId, formTemplate, queryMap, lastAccountingYear, lastSequenceNo)
		// 累加最近一个会计期结束日期+1 至 billDateBegin 之间的本期增加,本期减少到期初,
		{
			var lastEndDate int
			if lastAccountingYear > 0 && lastSequenceNo > 0 {
				accountInOutService := AccountInOutService{}
				_, lastEndDate = accountInOutService.GetAccountingPeriodStartEndDate(sessionId, lastAccountingYear, lastSequenceNo)
			} else { // 没有找到最近的一个会计期时,将lastEndDate置为19700101
				lastEndDate = 19700101
			}
			dateUtil := DateUtil{}
			nextOfLastEndDate := dateUtil.GetNextDate(lastEndDate)
			preOfBillDateBegin := dateUtil.GetPreDate(billDateBegin)
			amtIncreaseReduceLi := c.getAmtIncreaseReduceByDate(sessionId, formTemplate, queryMap, nextOfLastEndDate, preOfBillDateBegin)
			origBalanceLi = c.addIncreaseReduceToOrigBalance(origBalanceLi, amtIncreaseReduceLi)
		}
		amtIncreaseReduceLi := c.getAmtIncreaseReduceByDate(sessionId, formTemplate, queryMap, billDateBegin, billDateEnd)
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
		lastAccountingYear, lastSequenceNo := c.getLastAccountingYearSequenceNoByYearMonth(sessionId, formTemplate, accountingYearStart, accountingPeriodStart)

		// 计算期初和本期增加,本期减少,
		origBalanceLi := c.getOrigBalanceFromAccountInOut(sessionId, formTemplate, queryMap, lastAccountingYear, lastSequenceNo)
		amtIncreaseReduceLi := c.getAmtIncreaseReduceByYearMonth(sessionId, formTemplate, queryMap, accountingYearStart, accountingPeriodStart, accountingYearEnd, accountingPeriodEnd)
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
	mathUtil := MathUtil{}
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
		origBalanceInData := fmt.Sprint(data["origBalance"])
		origBalance := fmt.Sprint(master["origBalance"])
		amtIncrease := fmt.Sprint(master["amtIncrease"])
		amtReduce := fmt.Sprint(master["amtReduce"])

		origBalanceResult := mathUtil.Add(origBalanceInData, origBalance)
		origBalanceResult = mathUtil.Add(origBalanceResult, amtIncrease)
		origBalanceResult = mathUtil.Sub(origBalanceResult, amtReduce)
		data["origBalance"] = origBalanceResult

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
	mathUtil := MathUtil{}
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
		origBalanceInData := fmt.Sprint(data["origBalance"])
		amtIncreaseInData := fmt.Sprint(data["amtIncrease"])
		amtReduceInData := fmt.Sprint(data["amtReduce"])
		increaseCountInData := commonUtil.GetIntFromMap(data, "increaseCount")
		reduceCountInData := commonUtil.GetIntFromMap(data, "reduceCount")

		origBalance := fmt.Sprint(master["origBalance"])
		amtIncrease := fmt.Sprint(master["amtIncrease"])
		amtReduce := fmt.Sprint(master["amtReduce"])
		increaseCount := commonUtil.GetIntFromMap(master, "increaseCount")
		reduceCount := commonUtil.GetIntFromMap(master, "reduceCount")

		data["origBalance"] = mathUtil.Add(origBalanceInData, origBalance)
		data["amtIncrease"] = mathUtil.Add(amtIncreaseInData, amtIncrease)
		data["amtReduce"] = mathUtil.Add(amtReduceInData, amtReduce)
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

func (c AccountInOutDisplay) getOrigBalanceFromAccountInOut(sessionId int, formTemplate FormTemplate, queryMap map[string]interface{}, lastAccountingYear int, lastSequenceNo int) []interface{} {
	session, _ := global.GetConnection(sessionId)
	result := []interface{}{}
	querySupport := QuerySupport{}
	accountInitQuery := map[string]interface{}{}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		accountInitQuery[k] = v
	}

	orQuery := []interface{}{}
	commonUtil := CommonUtil{}
	mathUtil := MathUtil{}
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
		origBalance := fmt.Sprint(data["origBalance"])
		amtIncrease := fmt.Sprint(master["amtIncrease"])
		amtReduce := fmt.Sprint(master["amtReduce"])
		origBalanceResult := mathUtil.Add(origBalance, amtIncrease)
		origBalanceResult = mathUtil.Sub(origBalanceResult, amtReduce)
		data["origBalance"] = origBalanceResult
		uniqueMap[key] = data
	}
	for _, item := range uniqueMap {
		result = append(result, item)
	}
	return result
}

func (c AccountInOutDisplay) getAmtIncreaseReduceByDate(sessionId int, formTemplate FormTemplate, queryMap map[string]interface{}, billDateBegin int, billDateEnd int) []interface{} {
	session, _ := global.GetConnection(sessionId)
	result := []interface{}{}
	querySupport := QuerySupport{}
	accountInitQuery := map[string]interface{}{}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		accountInitQuery[k] = v
	}

	orQuery := []interface{}{}
	commonUtil := CommonUtil{}
	mathUtil := MathUtil{}
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
		amtIncreaseInData := fmt.Sprint(data["amtIncrease"])
		amtReduceInData := fmt.Sprint(data["amtReduce"])
		increaseCountInData := commonUtil.GetIntFromMap(data, "increaseCount")
		reduceCountInData := commonUtil.GetIntFromMap(data, "reduceCount")

		amtIncrease := fmt.Sprint(master["amtIncrease"])
		amtReduce := fmt.Sprint(master["amtReduce"])
		data["amtIncrease"] = mathUtil.Add(amtIncreaseInData, amtIncrease)
		data["amtReduce"] = mathUtil.Add(amtReduceInData, amtReduce)
		if commonUtil.GetFloat64FromString(amtIncrease) > 0 {
			data["increaseCount"] = increaseCountInData + 1
		}
		if commonUtil.GetFloat64FromString(amtReduce) > 0 {
			data["reduceCount"] = reduceCountInData + 1
		}
		uniqueMap[key] = data
	}
	for _, item := range uniqueMap {
		result = append(result, item)
	}
	return result
}

func (c AccountInOutDisplay) getAmtIncreaseReduceByYearMonth(sessionId int, formTemplate FormTemplate, queryMap map[string]interface{}, accountingYearStart int, accountingPeriodStart int, accountingYearEnd int, accountingPeriodEnd int) []interface{} {
	session, _ := global.GetConnection(sessionId)
	result := []interface{}{}
	querySupport := QuerySupport{}
	accountInitQuery := map[string]interface{}{}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		accountInitQuery[k] = v
	}

	orQuery := []interface{}{}
	commonUtil := CommonUtil{}
	mathUtil := MathUtil{}
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
		amtIncreaseInData := fmt.Sprint(data["amtIncrease"])
		amtReduceInData := fmt.Sprint(data["amtReduce"])
		increaseCountInData := commonUtil.GetIntFromMap(data, "increaseCount")
		reduceCountInData := commonUtil.GetIntFromMap(data, "reduceCount")

		amtIncrease := fmt.Sprint(master["amtIncrease"])
		amtReduce := fmt.Sprint(master["amtReduce"])
		increaseCount := commonUtil.GetIntFromMap(master, "increaseCount")
		reduceCount := commonUtil.GetIntFromMap(master, "reduceCount")

		data["amtIncrease"] = mathUtil.Add(amtIncreaseInData, amtIncrease)
		data["amtReduce"] = mathUtil.Add(amtReduceInData, amtReduce)
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
func (c AccountInOutDisplay) getLastAccountingYearSequenceNoByYearMonth(sessionId int, formTemplate FormTemplate, accountingYear int, sequenceNo int) (int, int) {
	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
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
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
	}

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
func (c AccountInOutDisplay) getLastAccountingYearSequenceNo(sessionId int, formTemplate FormTemplate, date int) (int, int) {
	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"B.endDate": map[string]interface{}{
			"$lt": date,
		},
	}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
	}

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

func (c AccountInOutDisplay) SaveData() revel.Result {
	c.RActionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountInOutDisplay) DeleteData() revel.Result {
	c.RActionSupport = AccountInOutDisplaySupport{}

	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountInOutDisplay) EditData() revel.Result {
	c.RActionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountInOutDisplay) NewData() revel.Result {
	c.RActionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountInOutDisplay) GetData() revel.Result {
	c.RActionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c AccountInOutDisplay) CopyData() revel.Result {
	c.RActionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountInOutDisplay) GiveUpData() revel.Result {
	c.RActionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountInOutDisplay) RefreshData() revel.Result {
	c.RActionSupport = AccountInOutDisplaySupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c AccountInOutDisplay) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
