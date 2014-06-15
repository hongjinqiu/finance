package controllers

import (
	. "com/papersns/component"
	. "com/papersns/error"
	"com/papersns/global"
	. "com/papersns/model"
	. "com/papersns/model/handler"
	. "com/papersns/mongo"
	"fmt"
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

func (c CashAccountInit) saveCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	// 自己拆,再循环保存
	//	dataSourceModelId := c.Params.Form.Get("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := c.Params.Get("formTemplateId")
	jsonBo := c.Params.Form.Get("jsonData")

	bo := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonBo), &bo)
	if err != nil {
		panic(err)
	}
	continueAnyAll := "false"
	if bo["continueAnyAll"] != nil && fmt.Sprint(bo["continueAnyAll"]) != "" {
		continueAnyAll = bo["continueAnyAll"].(string)
	}

	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	//	modelTemplateFactory.ConvertDataType(dataSource, &bo)

	cashAccountInitLi := []map[string]interface{}{}
	bDataSetLi := bo["B"].([]interface{})
	nowIdLi := []int{}
	for _, item := range bDataSetLi {
		mapItem := item.(map[string]interface{})
		line := map[string]interface{}{
			"A": mapItem,
		}
		if mapItem["id"] != nil {
			line["id"] = mapItem["id"]
			line["_id"] = mapItem["id"]
		} else {
			line["id"] = ""
			line["_id"] = ""
		}
		modelTemplateFactory.ConvertDataType(dataSource, &line)
		strId := fmt.Sprint(line["id"])
		if strId != "" {
			intId, _ := strconv.Atoi(strId)
			nowIdLi = append(nowIdLi, intId)
		}
		cashAccountInitLi = append(cashAccountInitLi, line)
	}
	queryData := bo["A"].(map[string]interface{})
	strAccountId := fmt.Sprint(queryData["accountId"])
	// 先处理删除的数据,并弄到差异数据中
	diffDataRowAllLi := []DiffDataRow{}
	toDeleteLi := c.dealDelete(sessionId, dataSource, queryData, nowIdLi)
	modelIterator := ModelIterator{}
	for _, item := range toDeleteLi {
		var result interface{} = ""
		modelIterator.IterateDataBo(dataSource, &item, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
			diffDataRow := DiffDataRow{
				FieldGroupLi: fieldGroupLi,
				SrcData: *data,
				SrcBo: item,
			}
			diffDataRowAllLi = append(diffDataRowAllLi, diffDataRow)
		})
	}

	bo = map[string]interface{}{
		"_id": 0,
		"id":  0,
		"A": map[string]interface{}{
			"id":        0,
			"accountId": strAccountId,
		},
	}
	bDataSetLi = []interface{}{}
	for i, _ := range cashAccountInitLi {
		cashAccountInit := &cashAccountInitLi[i]
		strId := modelTemplateFactory.GetStrId(*cashAccountInit)
		if strId == "" || strId == "0" {
			c.setCreateFixFieldValue(sessionId, dataSource, cashAccountInit)
		} else {
			c.setModifyFixFieldValue(sessionId, dataSource, cashAccountInit)
			editMessage, isValid := c.actionSupport.editValidate(sessionId, dataSource, *cashAccountInit)
			if !isValid {
				panic(editMessage)
			}
		}
		// 这样只会是新增和修改的数据
		c.actionSupport.beforeSaveData(sessionId, dataSource, cashAccountInit)
		financeService := FinanceService{}
		diffDataRowLi := financeService.SaveData(sessionId, dataSource, cashAccountInit)
		c.actionSupport.afterSaveData(sessionId, dataSource, cashAccountInit, diffDataRowLi)

		bDataSetLi = append(bDataSetLi, (*cashAccountInit)["A"])
		
		for _, diffDataRowItem := range *diffDataRowLi {
			diffDataRowAllLi = append(diffDataRowAllLi, diffDataRowItem)
		}
	}
	cashAccountInitSupport := c.actionSupport.(CashAccountInitSupport)
	cashAccountInitSupport.checkLimitsControl(sessionId, diffDataRowAllLi, continueAnyAll)
	bo["B"] = bDataSetLi

	usedCheckBo := map[string]interface{}{}

	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	columnModelData := templateManager.GetColumnModelDataForFormTemplate(formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return ModelRenderVO{
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

func (c CashAccountInit) dealDelete(sessionId int, dataSource DataSource, queryData map[string]interface{}, nowIdLi []int) []map[string]interface{} {
	modelTemplateFactory := ModelTemplateFactory{}
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	_, db := global.GetConnection(sessionId)
	queryMap := map[string]interface{}{
		"_id": map[string]interface{}{
			"$nin": nowIdLi,
		},
		"A.accountType": 1,
	}
	strAccountId := fmt.Sprint(queryData["accountId"])
	if strAccountId != "" && strAccountId != "0" {
		accountIdLi := []int{}
		for _, item := range strings.Split(strAccountId, ",") {
			accountId, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			accountIdLi = append(accountIdLi, accountId)
		}
		queryMap["A.accountId"] = map[string]interface{}{
			"$in": accountIdLi,
		}
	}
	queryMapByte, err := json.Marshal(&queryMap)
	if err != nil {
		panic(err)
	}
	log.Println("dealDelete,collectionName:" + collectionName + ", query:" + string(queryMapByte))
	toDeleteLi := []map[string]interface{}{}
	coll := db.C(collectionName)
	err = coll.Find(queryMap).All(&toDeleteLi)
	if err != nil {
		panic(err)
	}
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	modelIterator := ModelIterator{}
	usedCheck := UsedCheck{}
	for _, item := range toDeleteLi {
		if usedCheck.CheckUsed(sessionId, dataSource, item) {
			panic(BusinessError{Message: "已被用，不能删除"})
		}

		// 删除被用,帐户初始化不存在被用,不判断被用
		var result interface{} = ""
		modelIterator.IterateDataBo(dataSource, &item, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
			usedCheck.DeleteAll(sessionId, fieldGroupLi, *data)
		})

		c.actionSupport.beforeDeleteData(sessionId, dataSource, &item)
		_, removeResult := txnManager.Remove(txnId, collectionName, item)
		if !removeResult {
			panic("删除失败")
		}
		c.actionSupport.afterDeleteData(sessionId, dataSource, &item)
	}
	return toDeleteLi
}

func (c CashAccountInit) getDataCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	//	dataSourceModelId := c.Params.Get("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := c.Params.Get("formTemplateId")

	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 1,
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	c.actionSupport.beforeGetData(sessionId, dataSource)
	// 需要进行循环处理,再转回来,
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	result := querySupport.Index(collectionName, queryMap, pageNo, pageSize, orderBy)
	items := result["items"].([]interface{})
	dataSetLi := []interface{}{}
	for _, item := range items {
		line := item.(map[string]interface{})
		bo := map[string]interface{}{
			"_id": line["id"],
			"id":  line["id"],
			"A":   line["A"],
		}
		modelTemplateFactory.ConvertDataType(dataSource, &bo)
		c.actionSupport.afterGetData(sessionId, dataSource, &bo)
		dataSetLi = append(dataSetLi, bo["A"])
	}
	bo := map[string]interface{}{
		"_id": 0,
		"id":  0,
		"A": map[string]interface{}{
			"id":        0,
			"accountId": 0,
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

	//	modelTemplateFactory.ConvertDataType(dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	return ModelRenderVO{
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

func (c CashAccountInit) editDataCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	//	dataSourceModelId := c.Params.Get("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := c.Params.Get("formTemplateId")
	queryDataStr := c.Params.Get("queryData")
	queryData := map[string]interface{}{}
	err := json.Unmarshal([]byte(queryDataStr), &queryData)
	if err != nil {
		panic(err)
	}
	strAccountId := fmt.Sprint(queryData["accountId"])
	//	strId := c.Params.Get("id")
	//	id, err := strconv.Atoi(strId)
	//	if err != nil {
	//		panic(err)
	//	}

	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 1,
	}
	if strAccountId != "" && strAccountId != "0" {
		accountIdLi := []int{}
		for _, item := range strings.Split(strAccountId, ",") {
			accountId, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			accountIdLi = append(accountIdLi, accountId)
		}
		queryMap["A.accountId"] = map[string]interface{}{
			"$in": accountIdLi,
		}
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	//	bo, found := querySupport.FindByMap(collectionName, queryMap)
	//	if !found {
	//		panic("RefreshData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	//	}
	// 需要进行循环处理,再转回来,
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	result := querySupport.Index(collectionName, queryMap, pageNo, pageSize, orderBy)
	items := result["items"].([]interface{})
	dataSetLi := []interface{}{}
	for _, item := range items {
		line := item.(map[string]interface{})
		bo := map[string]interface{}{
			"_id": line["id"],
			"id":  line["id"],
			"A":   line["A"],
		}

		modelTemplateFactory.ConvertDataType(dataSource, &bo)
		editMessage, isValid := c.actionSupport.editValidate(sessionId, dataSource, bo)
		if !isValid {
			panic(editMessage)
		}

		c.actionSupport.beforeEditData(sessionId, dataSource, &bo)
		c.actionSupport.afterEditData(sessionId, dataSource, &bo)
		dataSetLi = append(dataSetLi, bo["A"])
	}

	bo := map[string]interface{}{
		"_id": 0,
		"id":  0,
		"A": map[string]interface{}{
			"id":        0,
			"accountId": strAccountId,
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

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return ModelRenderVO{
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

func (c CashAccountInit) giveUpDataCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	//	dataSourceModelId := c.Params.Get("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := c.Params.Get("formTemplateId")
	queryDataStr := c.Params.Get("queryData")
	queryData := map[string]interface{}{}
	err := json.Unmarshal([]byte(queryDataStr), &queryData)
	if err != nil {
		panic(err)
	}
	strAccountId := fmt.Sprint(queryData["accountId"])
	//	strId := c.Params.Get("id")
	//	id, err := strconv.Atoi(strId)
	//	if err != nil {
	//		panic(err)
	//	}

	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 1,
	}
	if strAccountId != "" && strAccountId != "0" {
		accountIdLi := []int{}
		for _, item := range strings.Split(strAccountId, ",") {
			accountId, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			accountIdLi = append(accountIdLi, accountId)
		}
		queryMap["A.accountId"] = map[string]interface{}{
			"$in": accountIdLi,
		}
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	//	bo, found := querySupport.FindByMap(collectionName, queryMap)
	//	if !found {
	//		panic("giveUpData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	//	}
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	result := querySupport.Index(collectionName, queryMap, pageNo, pageSize, orderBy)
	items := result["items"].([]interface{})
	dataSetLi := []interface{}{}
	for _, item := range items {
		line := item.(map[string]interface{})
		bo := map[string]interface{}{
			"_id": line["id"],
			"id":  line["id"],
			"A":   line["A"],
		}
		modelTemplateFactory.ConvertDataType(dataSource, &bo)
		c.actionSupport.beforeGiveUpData(sessionId, dataSource, &bo)
		c.actionSupport.afterGiveUpData(sessionId, dataSource, &bo)
		dataSetLi = append(dataSetLi, bo["A"])
	}
	bo := map[string]interface{}{
		"_id": 0,
		"id":  0,
		"A": map[string]interface{}{
			"id":        0,
			"accountId": strAccountId,
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

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return ModelRenderVO{
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

func (c CashAccountInit) refreshDataCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	//	dataSourceModelId := c.Params.Get("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := c.Params.Get("formTemplateId")
	queryDataStr := c.Params.Get("queryData")
	queryData := map[string]interface{}{}
	err := json.Unmarshal([]byte(queryDataStr), &queryData)
	if err != nil {
		panic(err)
	}
	strAccountId := fmt.Sprint(queryData["accountId"])
	//	strId := c.Params.Get("id")
	//	id, err := strconv.Atoi(strId)
	//	if err != nil {
	//		panic(err)
	//	}

	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 1,
	}
	if strAccountId != "" && strAccountId != "0" {
		accountIdLi := []int{}
		for _, item := range strings.Split(strAccountId, ",") {
			accountId, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			accountIdLi = append(accountIdLi, accountId)
		}
		queryMap["A.accountId"] = map[string]interface{}{
			"$in": accountIdLi,
		}
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	//	bo, found := querySupport.FindByMap(collectionName, queryMap)
	//	if !found {
	//		panic("RefreshData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	//	}
	// 需要进行循环处理,再转回来,
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	result := querySupport.Index(collectionName, queryMap, pageNo, pageSize, orderBy)
	items := result["items"].([]interface{})
	dataSetLi := []interface{}{}
	for _, item := range items {
		line := item.(map[string]interface{})
		bo := map[string]interface{}{
			"_id": line["id"],
			"id":  line["id"],
			"A":   line["A"],
		}
		modelTemplateFactory.ConvertDataType(dataSource, &bo)
		c.actionSupport.beforeRefreshData(sessionId, dataSource, &bo)
		c.actionSupport.afterRefreshData(sessionId, dataSource, &bo)
		dataSetLi = append(dataSetLi, bo["A"])
	}
	bo := map[string]interface{}{
		"_id": 0,
		"id":  0,
		"A": map[string]interface{}{
			"id":        0,
			"accountId": strAccountId,
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

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return ModelRenderVO{
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}