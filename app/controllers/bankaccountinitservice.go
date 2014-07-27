package controllers

import (
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
	"strings"
)

func (c BankAccountInit) saveCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	userId, err := strconv.Atoi(c.Session["userId"])
	if err != nil {
		panic(err)
	}

	// 自己拆,再循环保存
	//	dataSourceModelId := c.Params.Form.Get("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := c.Params.Get("formTemplateId")
	jsonBo := c.Params.Form.Get("jsonData")

	bo := map[string]interface{}{}
	err = json.Unmarshal([]byte(jsonBo), &bo)
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
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)

	bankAccountInitLi := []map[string]interface{}{}
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
		bankAccountInitLi = append(bankAccountInitLi, line)
	}
	queryData := bo["A"].(map[string]interface{})
	strAccountId := fmt.Sprint(queryData["accountId"])
	// 先处理删除的数据,并弄到差异数据中
	diffDataRowAllLi := []DiffDataRow{}
	toDeleteLi := c.dealDelete(sessionId, dataSource, formTemplate, queryData, nowIdLi)
	modelIterator := ModelIterator{}
	for _, item := range toDeleteLi {
		var result interface{} = ""
		modelIterator.IterateDataBo(dataSource, &item, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
			diffDataRow := DiffDataRow{
				FieldGroupLi: fieldGroupLi,
				SrcData:      *data,
				SrcBo:        item,
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
	for i, _ := range bankAccountInitLi {
		bankAccountInit := &bankAccountInitLi[i]
		strId := modelTemplateFactory.GetStrId(*bankAccountInit)
		if strId == "" || strId == "0" {
			c.setCreateFixFieldValue(sessionId, dataSource, bankAccountInit)
		} else {
			c.setModifyFixFieldValue(sessionId, dataSource, bankAccountInit)
			editMessage, isValid := c.actionSupport.editValidate(sessionId, dataSource, formTemplate, *bankAccountInit)
			if !isValid {
				panic(editMessage)
			}
		}
		// 这样只会是新增和修改的数据
		c.actionSupport.beforeSaveData(sessionId, dataSource, formTemplate, bankAccountInit)
		financeService := FinanceService{}
		diffDataRowLi := financeService.SaveData(sessionId, dataSource, bankAccountInit)
		c.actionSupport.afterSaveData(sessionId, dataSource, formTemplate, bankAccountInit, diffDataRowLi)

		bDataSetLi = append(bDataSetLi, (*bankAccountInit)["A"])

		for _, diffDataRowItem := range *diffDataRowLi {
			diffDataRowAllLi = append(diffDataRowAllLi, diffDataRowItem)
		}
	}
	bankAccountInitSupport := c.actionSupport.(BankAccountInitSupport)
	bankAccountInitSupport.checkLimitsControl(sessionId, diffDataRowAllLi, continueAnyAll)
	bo["B"] = bDataSetLi

	usedCheckBo := map[string]interface{}{}

	columnModelData := templateManager.GetColumnModelDataForFormTemplate(sessionId, formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return ModelRenderVO{
		UserId:      userId,
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

func (c BankAccountInit) dealDelete(sessionId int, dataSource DataSource, formTemplate FormTemplate, queryData map[string]interface{}, nowIdLi []int) []map[string]interface{} {
	modelTemplateFactory := ModelTemplateFactory{}
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	_, db := global.GetConnection(sessionId)
	queryMap := map[string]interface{}{
		"_id": map[string]interface{}{
			"$nin": nowIdLi,
		},
		"A.accountType": 2,
	}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
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

		c.actionSupport.beforeDeleteData(sessionId, dataSource, formTemplate, &item)
		_, removeResult := txnManager.Remove(txnId, collectionName, item)
		if !removeResult {
			panic("删除失败")
		}
		c.actionSupport.afterDeleteData(sessionId, dataSource, formTemplate, &item)
	}
	return toDeleteLi
}

func (c BankAccountInit) getDataCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	userId, err := strconv.Atoi(c.Session["userId"])
	if err != nil {
		panic(err)
	}

	//	dataSourceModelId := c.Params.Get("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := c.Params.Get("formTemplateId")

	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 2,
	}
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
	}
	
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	c.actionSupport.beforeGetData(sessionId, dataSource, formTemplate)
	// 需要进行循环处理,再转回来,
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	result := querySupport.IndexWithSession(session, collectionName, queryMap, pageNo, pageSize, orderBy)
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
		c.actionSupport.afterGetData(sessionId, dataSource, formTemplate, &bo)
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

	columnModelData := templateManager.GetColumnModelDataForFormTemplate(sessionId, formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	//	modelTemplateFactory.ConvertDataType(dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	return ModelRenderVO{
		UserId:      userId,
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

func (c BankAccountInit) editDataCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	userId, err := strconv.Atoi(c.Session["userId"])
	if err != nil {
		panic(err)
	}

	//	dataSourceModelId := c.Params.Get("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := c.Params.Get("formTemplateId")
	queryDataStr := c.Params.Get("queryData")
	queryData := map[string]interface{}{}
	err = json.Unmarshal([]byte(queryDataStr), &queryData)
	if err != nil {
		panic(err)
	}
	strAccountId := fmt.Sprint(queryData["accountId"])
	//	strId := c.Params.Get("id")
	//	id, err := strconv.Atoi(strId)
	//	if err != nil {
	//		panic(err)
	//	}

	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 2,
	}
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
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
	// 需要进行循环处理,再转回来,
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	result := querySupport.IndexWithSession(session, collectionName, queryMap, pageNo, pageSize, orderBy)
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
		editMessage, isValid := c.actionSupport.editValidate(sessionId, dataSource, formTemplate, bo)
		if !isValid {
			panic(editMessage)
		}

		c.actionSupport.beforeEditData(sessionId, dataSource, formTemplate, &bo)
		c.actionSupport.afterEditData(sessionId, dataSource, formTemplate, &bo)
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

	columnModelData := templateManager.GetColumnModelDataForFormTemplate(sessionId, formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return ModelRenderVO{
		UserId:      userId,
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
		FormTemplate: formTemplate,
	}
}

func (c BankAccountInit) giveUpDataCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	userId, err := strconv.Atoi(c.Session["userId"])
	if err != nil {
		panic(err)
	}

	//	dataSourceModelId := c.Params.Get("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := c.Params.Get("formTemplateId")
	queryDataStr := c.Params.Get("queryData")
	queryData := map[string]interface{}{}
	err = json.Unmarshal([]byte(queryDataStr), &queryData)
	if err != nil {
		panic(err)
	}
	strAccountId := fmt.Sprint(queryData["accountId"])
	//	strId := c.Params.Get("id")
	//	id, err := strconv.Atoi(strId)
	//	if err != nil {
	//		panic(err)
	//	}

	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 2,
	}
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
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
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	result := querySupport.IndexWithSession(session, collectionName, queryMap, pageNo, pageSize, orderBy)
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
		c.actionSupport.beforeGiveUpData(sessionId, dataSource, formTemplate, &bo)
		c.actionSupport.afterGiveUpData(sessionId, dataSource, formTemplate, &bo)
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

	columnModelData := templateManager.GetColumnModelDataForFormTemplate(sessionId, formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return ModelRenderVO{
		UserId:      userId,
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

func (c BankAccountInit) refreshDataCommon() ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	userId, err := strconv.Atoi(c.Session["userId"])
	if err != nil {
		panic(err)
	}

	//	dataSourceModelId := c.Params.Get("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := c.Params.Get("formTemplateId")
	queryDataStr := c.Params.Get("queryData")
	queryData := map[string]interface{}{}
	err = json.Unmarshal([]byte(queryDataStr), &queryData)
	if err != nil {
		panic(err)
	}
	strAccountId := fmt.Sprint(queryData["accountId"])
	//	strId := c.Params.Get("id")
	//	id, err := strconv.Atoi(strId)
	//	if err != nil {
	//		panic(err)
	//	}

	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 2,
	}
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
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
	// 需要进行循环处理,再转回来,
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	result := querySupport.IndexWithSession(session, collectionName, queryMap, pageNo, pageSize, orderBy)
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
		c.actionSupport.beforeRefreshData(sessionId, dataSource, formTemplate, &bo)
		c.actionSupport.afterRefreshData(sessionId, dataSource, formTemplate, &bo)
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

	columnModelData := templateManager.GetColumnModelDataForFormTemplate(sessionId, formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return ModelRenderVO{
		UserId:      userId,
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}
