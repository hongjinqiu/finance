package console

import "github.com/robfig/revel"
import (
	. "com/papersns/common"
	. "com/papersns/component"
	"com/papersns/global"
	. "com/papersns/model"
	. "com/papersns/model/handler"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"html/template"
	"os"
	"strconv"
	"strings"
	. "com/papersns/error"
	. "com/papersns/mongo"
	"com/papersns/mongo"
	"log"
	"time"
)

func init() {

}

type Console struct {
	*revel.Controller
}

// 管理员查看页面,设置session.userId,以查看数据,
func (c Console) BbsPostAdminReplySchema() revel.Result {
	// 取一下bbsPostId

	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)
	defer c.RRollbackTxn(sessionId)

	schemaName := c.Params.Get("@name")

	templateManager := TemplateManager{}
	listTemplate := templateManager.GetListTemplate(schemaName)

	isFromList := true
	result := c.listSelectorCommon(&listTemplate, true, isFromList)
	bbsPostIdStr := c.Params.Get("bbsPostId")
	bbsPostId, err := strconv.Atoi(bbsPostIdStr)
	if err != nil {
		panic(err)
	}
	c.addOrUpdateBbsPostRead(sessionId, bbsPostId)
	c.CommitTxn(sessionId)

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		callback := c.Params.Get("callback")
		if callback == "" {
			dataBo := result["dataBo"]
			c.Response.ContentType = "application/json; charset=utf-8"
			return c.RenderJson(&dataBo)
		}
		dataBoText := result["dataBoText"].(string)
		c.Response.ContentType = "text/javascript; charset=utf-8"
		return c.RenderText(callback + "(" + dataBoText + ");")
	} else {
		//c.Response.ContentType = "text/html; charset=utf-8"
		c.RenderArgs["result"] = result
		return c.RenderTemplate(listTemplate.ViewTemplate.View)
		//		return c.Render(result)
	}
}

func (c Console) BbsPostReplySchema() revel.Result {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)
	defer c.RRollbackTxn(sessionId)

	schemaName := c.Params.Get("@name")

	templateManager := TemplateManager{}
	listTemplate := templateManager.GetListTemplate(schemaName)

	isFromList := true
	result := c.listSelectorCommon(&listTemplate, true, isFromList)
	bbsPostIdStr := c.Params.Get("bbsPostId")
	bbsPostId, err := strconv.Atoi(bbsPostIdStr)
	if err != nil {
		panic(err)
	}
	c.addOrUpdateBbsPostRead(sessionId, bbsPostId)
	c.CommitTxn(sessionId)

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		callback := c.Params.Get("callback")
		if callback == "" {
			dataBo := result["dataBo"]
			c.Response.ContentType = "application/json; charset=utf-8"
			return c.RenderJson(&dataBo)
		}
		dataBoText := result["dataBoText"].(string)
		c.Response.ContentType = "text/javascript; charset=utf-8"
		return c.RenderText(callback + "(" + dataBoText + ");")
	} else {
		//c.Response.ContentType = "text/html; charset=utf-8"
		c.RenderArgs["result"] = result
		return c.RenderTemplate(listTemplate.ViewTemplate.View)
		//		return c.Render(result)
	}
}

func (c Console) CommitTxn(sessionId int) {
	txnId := global.GetGlobalAttr(sessionId, "txnId")
	if txnId != nil {
		_, db := global.GetConnection(sessionId)
		txnManager := TxnManager{db}
		txnManager.Commit(txnId.(int))
	}
}

func (c Console) RRollbackTxn(sessionId int) {
	txnId := global.GetGlobalAttr(sessionId, "txnId")
	if txnId != nil {
		if x := recover(); x != nil {
			_, db := global.GetConnection(sessionId)
			txnManager := TxnManager{db}
			txnManager.Rollback(txnId.(int))
			panic(x)
		}
	}
}

func (c Console) addOrUpdateBbsPostRead(sessionId int, bbsPostId int) {
	session, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	modelTemplateFactory := ModelTemplateFactory{}
	bbsPostReadDS := modelTemplateFactory.GetDataSource("BbsPostRead")

	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	dateUtil := DateUtil{}
	qb := QuerySupport{}

	bbsPostRead, found := qb.FindByMapWithSession(session, "BbsPostRead", map[string]interface{}{
		"A.bbsPostId":  bbsPostId,
		"A.readBy":     userId,
	})
	if found {
		c.RSetModifyFixFieldValue(sessionId, bbsPostReadDS, &bbsPostRead)
		bbsPostReadA := bbsPostRead["A"].(map[string]interface{})
		bbsPostRead["A"] = bbsPostReadA

		bbsPostReadA["lastReadTime"] = dateUtil.GetCurrentYyyyMMddHHmmss()
		_, updateResult := txnManager.Update(txnId, "BbsPostRead", bbsPostRead)
		if !updateResult {
			panic(BusinessError{Message: "更新意见反馈阅读记录失败"})
		}
	} else {
		c.addBbsPostRead(sessionId, bbsPostId)
	}
}

func (c Console) addBbsPostRead(sessionId int, bbsPostId int) {
	_, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	modelTemplateFactory := ModelTemplateFactory{}
	bbsPostReadDS := modelTemplateFactory.GetDataSource("BbsPostRead")

	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	dateUtil := DateUtil{}
	sequenceNo := mongo.GetSequenceNo(db, "bbsPostReadId")
	bbsPostRead := map[string]interface{}{
		"_id": sequenceNo,
		"id":  sequenceNo,
		"A": map[string]interface{}{
			"id":           sequenceNo,
			"bbsPostId":    bbsPostId,
			"readBy":       userId,
			"lastReadTime": dateUtil.GetCurrentYyyyMMddHHmmss(),
		},
	}
	c.RSetCreateFixFieldValue(sessionId, bbsPostReadDS, &bbsPostRead)
	txnManager.Insert(txnId, "BbsPostRead", bbsPostRead)
}

func (c Console) RSetCreateFixFieldValue(sessionId int, dataSource DataSource, bo *map[string]interface{}) {
	var result interface{} = ""
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	createTime, err := strconv.ParseInt(time.Now().Format("20060102150405"), 10, 64)
	if err != nil {
		panic(err)
	}
	_, db := global.GetConnection(sessionId)
	sysUser := map[string]interface{}{}
	query := map[string]interface{}{
		"_id": userId,
	}
	err = db.C("SysUser").Find(query).One(&sysUser)
	if err != nil {
		panic(err)
	}
	sysUserMaster := sysUser["A"].(map[string]interface{})
	modelIterator := ModelIterator{}
	modelIterator.IterateDataBo(dataSource, bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		(*data)["createBy"] = userId
		(*data)["createTime"] = createTime
		(*data)["createUnit"] = sysUserMaster["createUnit"]
	})
}

func (c Console) RSetModifyFixFieldValue(sessionId int, dataSource DataSource, bo *map[string]interface{}) {
	var result interface{} = ""
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	modifyTime, err := strconv.ParseInt(time.Now().Format("20060102150405"), 10, 64)
	if err != nil {
		panic(err)
	}
	_, db := global.GetConnection(sessionId)
	sysUser := map[string]interface{}{}
	query := map[string]interface{}{
		"_id": userId,
	}
	err = db.C("SysUser").Find(query).One(&sysUser)
	if err != nil {
		panic(err)
	}
	sysUserMaster := sysUser["A"].(map[string]interface{})

	srcBo := map[string]interface{}{}
	srcQuery := map[string]interface{}{
		"_id": (*bo)["id"],
		"A.createUnit": sysUserMaster["createUnit"],
	}
	// log
	modelTemplateFactory := ModelTemplateFactory{}
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	srcQueryByte, err := json.Marshal(&srcQuery)
	if err != nil {
		panic(err)
	}
	log.Println("RSetModifyFixFieldValue,collectionName:" + collectionName + ", query:" + string(srcQueryByte))
	db.C(collectionName).Find(srcQuery).One(&srcBo)
	modelIterator := ModelIterator{}
	modelIterator.IterateDiffBo(dataSource, bo, srcBo, &result, func(fieldGroupLi []FieldGroup, destData *map[string]interface{}, srcData map[string]interface{}, result *interface{}) {
		if destData != nil && srcData == nil {
			(*destData)["createBy"] = userId
			(*destData)["createTime"] = modifyTime
			(*destData)["createUnit"] = sysUserMaster["createUnit"]
		} else if destData == nil && srcData != nil {
			// 删除,不处理
		} else if destData != nil && srcData != nil {
			isMasterData := fieldGroupLi[0].IsMasterField()
			isDetailDataDiff := (!fieldGroupLi[0].IsMasterField()) && modelTemplateFactory.IsDataDifferent(fieldGroupLi, *destData, srcData)
			if isMasterData || isDetailDataDiff {
				(*destData)["createBy"] = srcData["createBy"]
				(*destData)["createTime"] = srcData["createTime"]
				(*destData)["createUnit"] = srcData["createUnit"]

				(*destData)["modifyBy"] = userId
				(*destData)["modifyTime"] = modifyTime
				(*destData)["modifyUnit"] = sysUserMaster["createUnit"]
			}
		}
	})
}

func (c Console) Summary() revel.Result {
	println("session is:", c.Session["userId"])
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)

	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate("Console")

	//	if true {
	//		xmlDataArray, err := xml.Marshal(&formTemplate)
	//		if err != nil {
	//			panic(err)
	//		}
	//		return c.RenderXml(&formTemplate)
	//	}

	formTemplateJsonDataArray, err := json.Marshal(&formTemplate)
	if err != nil {
		panic(err)
	}

	toolbarBo := map[string]interface{}{}

	dataBo := map[string]interface{}{}
	{
		listTemplateInfoLi := templateManager.GetListTemplateInfoLi()
		dataBo["Component"] = getSummaryListTemplateInfoLi(listTemplateInfoLi)
	}
	{
		selectorTemplateInfoLi := templateManager.GetSelectorTemplateInfoLi()
		dataBo["Selector"] = getSummarySelectorTemplateInfoLi(selectorTemplateInfoLi)
	}
	{
		formTemplateInfoLi := templateManager.GetFormTemplateInfoLi()
		dataBo["Form"] = getSummaryFormTemplateInfoLi(formTemplateInfoLi)
	}
	{
		modelTemplateFactory := ModelTemplateFactory{}
		dataSourceInfoLi := modelTemplateFactory.GetDataSourceInfoLi()
		dataBo["DataSource"] = getSummaryDataSourceInfoLi(dataSourceInfoLi)
	}
	for _, item := range formTemplate.FormElemLi {
		if item.XMLName.Local == "column-model" {
			if dataBo[item.ColumnModel.Name] == nil {
				dataBo[item.ColumnModel.Name] = []interface{}{}
			}
			items := dataBo[item.ColumnModel.Name].([]interface{})
			itemsDict := templateManager.GetColumnModelDataForColumnModel(sessionId, item.ColumnModel, items)
			items = itemsDict["items"].([]interface{})
			dataBo[item.ColumnModel.Name] = items
		} else if item.XMLName.Local == "toolbar" {
			toolbarBo[item.Toolbar.Name] = templateManager.GetToolbarBo(item.Toolbar)
		}
	}

	dataBoByte, err := json.Marshal(dataBo)
	if err != nil {
		panic(err)
	}

	//	c.Response.Status = http.StatusOK
	//	c.Response.ContentType = "text/plain; charset=utf-8"
	result := map[string]interface{}{
		"formTemplate":         formTemplate,
		"toolbarBo":            toolbarBo,
		"dataBo":               dataBo,
		"formTemplateJsonData": template.JS(string(formTemplateJsonDataArray)),
		"dataBoJson":           template.JS(string(dataBoByte)),
	}
	// formTemplate.ViewTemplate.View
	//	RenderText(text string, objs ...interface{}) Result

	viewPath := revel.Config.StringDefault("REVEL_VIEW_PATH", "")
	file, err := os.Open(viewPath + "/" + formTemplate.ViewTemplate.View)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	//	c.Response.Out
	//	return c.RenderTemplate(string(fileContent))
	funcMap := map[string]interface{}{
		"eq": func(a, b interface{}) bool {
			return a == b
		},
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	tmpl, err := template.New("summary").Funcs(funcMap).Parse(string(fileContent))
	if err != nil {
		panic(err)
	}
	tmplResult := map[string]interface{}{
		"result": result,
	}
	//tmpl.Execute(c.Response.Out, result)
	tmpl.Execute(c.Response.Out, tmplResult)
	return nil
	//	return c.Render(string(fileContent), result)
}

func getSummaryListTemplateInfoLi(listTemplateInfoLi []ListTemplateInfo) []interface{} {
	componentItems := []interface{}{}
	for _, item := range listTemplateInfoLi {
		module := "组件模型"
		if item.ListTemplate.DataSourceModelId != "" && item.ListTemplate.Adapter.Name != "" {
			module = "数据源模型适配"
		}
		componentItems = append(componentItems, map[string]interface{}{
			"id":     item.ListTemplate.Id,
			"name":   item.ListTemplate.Description,
			"module": module,
			"path":   item.Path,
		})
	}
	return componentItems
}

func getSummarySelectorTemplateInfoLi(selectorTemplateInfoLi []SelectorTemplateInfo) []interface{} {
	componentItems := []interface{}{}
	for _, item := range selectorTemplateInfoLi {
		module := "组件模型选择器"
		if item.ListTemplate.DataSourceModelId != "" && item.ListTemplate.Adapter.Name != "" {
			module = "数据源模型选择器适配"
		}
		id := item.ListTemplate.SelectorId
		if id == "" {
			id = item.ListTemplate.Id
		}
		componentItems = append(componentItems, map[string]interface{}{
			"id":     id,
			"name":   item.ListTemplate.Description,
			"module": module,
			"path":   item.Path,
		})
	}
	return componentItems
}

func getSummaryFormTemplateInfoLi(formTemplateInfoLi []FormTemplateInfo) []interface{} {
	formItems := []interface{}{}
	for _, item := range formTemplateInfoLi {
		module := "form模型"
		if item.FormTemplate.DataSourceModelId != "" && item.FormTemplate.Adapter.Name != "" {
			module = "数据源模型适配"
		}
		formItems = append(formItems, map[string]interface{}{
			"id":     item.FormTemplate.Id,
			"name":   item.FormTemplate.Description,
			"module": module,
			"path":   item.Path,
		})
	}
	return formItems
}

func getSummaryDataSourceInfoLi(dataSourceInfoLi []DataSourceInfo) []interface{} {
	dataSourceItems := []interface{}{}
	for _, item := range dataSourceInfoLi {
		dataSourceItems = append(dataSourceItems, map[string]interface{}{
			"id":     item.DataSource.Id,
			"name":   item.DataSource.DisplayName,
			"module": "数据源模型",
			"path":   item.Path,
		})
	}
	return dataSourceItems
}

func (c Console) ListSchema() revel.Result {
	schemaName := c.Params.Get("@name")

	templateManager := TemplateManager{}
	listTemplate := templateManager.GetListTemplate(schemaName)

	isFromList := true
	result := c.listSelectorCommon(&listTemplate, true, isFromList)

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		callback := c.Params.Get("callback")
		if callback == "" {
			dataBo := result["dataBo"]
			c.Response.ContentType = "application/json; charset=utf-8"
			return c.RenderJson(&dataBo)
		}
		dataBoText := result["dataBoText"].(string)
		c.Response.ContentType = "text/javascript; charset=utf-8"
		return c.RenderText(callback + "(" + dataBoText + ");")
	} else {
		//c.Response.ContentType = "text/html; charset=utf-8"
		c.RenderArgs["result"] = result
		c.RenderArgs["flash"] = c.Flash.Out
		c.RenderArgs["session"] = c.Session
		return c.RenderTemplate(listTemplate.ViewTemplate.View)
		//		return c.Render(result)
	}
}

func (c Console) SelectorSchema() revel.Result {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)

	schemaName := c.Params.Get("@name")

	templateManager := TemplateManager{}
	listTemplate := templateManager.GetSelectorTemplate(schemaName)
	c.setSelectionMode(&listTemplate)
	c.setDisplayField(&listTemplate)
	isGetBo := false
	if c.Params.Get("format") != "" {
		isGetBo = true
	}
	isFromList := false
	result := c.listSelectorCommon(&listTemplate, isGetBo, isFromList)

	selectionBo := map[string]interface{}{
		"url":         templateManager.GetViewUrl(listTemplate),
		"Description": listTemplate.Description,
	}
	ids := c.Params.Get("@id")
	if ids != "" {
		relationLi := []map[string]interface{}{}
		strIdLi := strings.Split(ids, ",")
		selectorId := listTemplate.SelectorId
		if selectorId == "" {
			selectorId = listTemplate.Id
		}
		for _, item := range strIdLi {
			if item != "" {
				id, err := strconv.Atoi(item)
				if err != nil {
					panic(err)
				}
				relationLi = append(relationLi, map[string]interface{}{
					"relationId": id,
					"selectorId": selectorId,
				})
			}
		}
		templateManager := TemplateManager{}
		relationBo := templateManager.GetRelationBo(sessionId, relationLi)
		if relationBo[selectorId] != nil {
			selectionBo = relationBo[selectorId].(map[string]interface{})
		}
	}
	selectionBoByte, err := json.Marshal(&selectionBo)
	if err != nil {
		panic(err)
	}

	commonUtil := CommonUtil{}
	selectionBoJson := string(selectionBoByte)
	selectionBoJson = commonUtil.FilterJsonEmptyAttr(selectionBoJson)
	result["selectionBoJson"] = template.JS(selectionBoJson)

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		callback := c.Params.Get("callback")
		if callback == "" {
			dataBo := result["dataBo"]
			c.Response.ContentType = "application/json; charset=utf-8"
			return c.RenderJson(&dataBo)
		}
		dataBoText := result["dataBoText"].(string)
		c.Response.ContentType = "text/javascript; charset=utf-8"
		return c.RenderText(callback + "(" + dataBoText + ");")
	} else {
		//		return c.Render(result)
		//c.Response.ContentType = "text/html; charset=utf-8"
		c.RenderArgs["result"] = result
		return c.RenderTemplate(listTemplate.ViewTemplate.SelectorView)
	}
}

func (c Console) setSelectionMode(listTemplate *ListTemplate) {
	multi := c.Params.Get("@multi")
	if multi != "" {
		if multi == "true" {
			listTemplate.ColumnModel.SelectionMode = "checkbox"
		} else {
			listTemplate.ColumnModel.SelectionMode = "radio"
		}
	}
}

func (c Console) setDisplayField(listTemplate *ListTemplate) {
	displayField := c.Params.Get("@displayField")
	if displayField != "" {
		if strings.Contains(displayField, "{") {
			listTemplate.ColumnModel.SelectionTemplate = displayField
		} else {
			strFieldLi := strings.Split(displayField, ",")
			fieldLi := []string{}
			for _, item := range strFieldLi {
				fieldLi = append(fieldLi, "{"+item+"}")
			}
			listTemplate.ColumnModel.SelectionTemplate = strings.Join(fieldLi, ",")
		}
	}
}

func (c Console) listSelectorCommon(listTemplate *ListTemplate, isGetBo bool, isFromList bool) map[string]interface{} {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)

	// 1.toolbar bo
	templateManager := TemplateManager{}
	//templateManager.ApplyDictionaryForQueryParameter(listTemplate)
	//templateManager.ApplyTreeForQueryParameter(listTemplate)
	toolbarBo := templateManager.GetToolbarForListTemplate(*listTemplate)
	paramMap := map[string]string{}

	defaultBo := templateManager.GetQueryDefaultValue(*listTemplate)
	defaultBoByte, err := json.Marshal(&defaultBo)
	if err != nil {
		panic(err)
	}
	for key, value := range defaultBo {
		paramMap[key] = value
	}
	paramMap, _ = c.getCookieDataAndParamMap(sessionId, *listTemplate, isFromList, paramMap)

	formDataByte, err := json.Marshal(&paramMap)
	if err != nil {
		panic(err)
	}

	//	}
	pageNo := 1
	pageSize := 10
	if listTemplate.DataProvider.Size != "" {
		pageSizeInt, err := strconv.Atoi(listTemplate.DataProvider.Size)
		if err != nil {
			panic(err)
		}
		pageSize = pageSizeInt
	}
	if c.Params.Get("pageNo") != "" {
		pageNoInt, _ := strconv.ParseInt(c.Params.Get("pageNo"), 10, 0)
		if pageNoInt > 1 {
			pageNo = int(pageNoInt)
		}
	}
	if c.Params.Get("pageSize") != "" {
		pageSizeInt, _ := strconv.ParseInt(c.Params.Get("pageSize"), 10, 0)
		if pageSizeInt >= 10 {
			pageSize = int(pageSizeInt)
		}
	}
	dataBo := map[string]interface{}{
		"totalResults": 0,
		"items":        []interface{}{},
	}
	relationBo := map[string]interface{}{}
	usedCheckBo := map[string]interface{}{}
	//if c.Params.Get("@entrance") != "true" {
	if isGetBo {
		dataBo = templateManager.GetBoForListTemplate(sessionId, listTemplate, paramMap, pageNo, pageSize)
		relationBo = dataBo["relationBo"].(map[string]interface{})
		//delete(dataBo, "relationBo")

		// usedCheck的修改,
		if listTemplate.DataSourceModelId != "" {
			modelTemplateFactory := ModelTemplateFactory{}
			dataSource := modelTemplateFactory.GetDataSource(listTemplate.DataSourceModelId)
			items := dataBo["items"].([]interface{})
			usedCheck := UsedCheck{}

			usedCheckBo = usedCheck.GetListUsedCheck(sessionId, dataSource, items, listTemplate.ColumnModel.DataSetId)
		}
	}
	dataBo["usedCheckBo"] = usedCheckBo

	dataBoByte, err := json.Marshal(&dataBo)
	if err != nil {
		panic(err)
	}

	relationBoByte, err := json.Marshal(&relationBo)
	if err != nil {
		panic(err)
	}

	listTemplateByte, err := json.Marshal(listTemplate)
	if err != nil {
		panic(err)
	}

	usedCheckByte, err := json.Marshal(&usedCheckBo)
	if err != nil {
		panic(err)
	}
	
	// 系统参数的获取
	commonUtil := CommonUtil{}
	userId := commonUtil.GetIntFromString(c.Session["userId"])
	sysParam := c.getSysParam(sessionId, userId)
	sysParamJson, err := json.Marshal(&sysParam)
	if err != nil {
		panic(err)
	}

	queryParameterRenderLi := c.getQueryParameterRenderLi(*listTemplate)

	//showParameterLi := templateManager.GetShowParameterLiForListTemplate(listTemplate)
	showParameterLi := []QueryParameter{}
	hiddenParameterLi := templateManager.GetHiddenParameterLiForListTemplate(listTemplate)

	layerBo := templateManager.GetLayerForListTemplate(sessionId, *listTemplate)
	iLayerBo := layerBo["layerBo"]
	layerBoByte, err := json.Marshal(&iLayerBo)
	if err != nil {
		panic(err)
	}
	iLayerBoLi := layerBo["layerBoLi"]
	layerBoLiByte, err := json.Marshal(&iLayerBoLi)
	if err != nil {
		panic(err)
	}
	layerBoJson := string(layerBoByte)
	layerBoJson = commonUtil.FilterJsonEmptyAttr(layerBoJson)
	layerBoLiJson := string(layerBoLiByte)
	layerBoLiJson = commonUtil.FilterJsonEmptyAttr(layerBoLiJson)

	result := map[string]interface{}{
		"pageSize":               pageSize,
		"listTemplate":           listTemplate,
		"toolbarBo":              toolbarBo,
		"showParameterLi":        showParameterLi,
		"hiddenParameterLi":      hiddenParameterLi,
		"queryParameterRenderLi": queryParameterRenderLi,
		"dataBo":                 dataBo,
		//		"columns":       columns,
		"dataBoText":       string(dataBoByte),
		"dataBoJson":       template.JS(string(dataBoByte)),
		"relationBoJson":   template.JS(string(relationBoByte)),
		"listTemplateJson": template.JS(string(listTemplateByte)),
		"layerBoJson":      template.JS(layerBoJson),
		"layerBoLiJson":    template.JS(layerBoLiJson),
		"defaultBoJson":    template.JS(string(defaultBoByte)),
		"formDataJson":     template.JS(string(formDataByte)),
		"usedCheckJson":    template.JS(string(usedCheckByte)),
		"sysParamJson":    template.JS(string(sysParamJson)),
		//		"columnsJson":   string(columnsByte),
	}
	return result
}

func (c Console) getCookieDataAndParamMap(sessionId int, listTemplate ListTemplate, isFromList bool, paramMap map[string]string) (map[string]string, map[string]string) {
	isHasCookie := false
	if c.Params.Query.Get("cookie") != "false" {
		isHasCookie = true
	}
	isConfigCookie := false
	if listTemplate.Cookie.Name != "" {
		isConfigCookie = true
	}
	cookieData := map[string]string{}
	if isFromList && isHasCookie && isConfigCookie {
		cookieStr := c.Session[listTemplate.Cookie.Name]
		if cookieStr != "" {
			err := json.Unmarshal([]byte(cookieStr), &cookieData)
			if err != nil {
				panic(err)
			}
			for k, v := range cookieData {
				paramMap[k] = v
			}
		}
	}
	formQueryData := map[string]string{}
	//	c.Request.URL
	for k, v := range c.Params.Form {
		value := strings.Join(v, ",")
		paramMap[k] = value
		formQueryData[k] = value
	}
	for k, v := range c.Params.Query {
		value := strings.Join(v, ",")
		paramMap[k] = value
		formQueryData[k] = value
	}
	
	if isFromList && isConfigCookie && !isHasCookie {
		c.Session[listTemplate.Cookie.Name] = ""
	} else if isFromList && isConfigCookie && isHasCookie {
		cookieFormQueryData := map[string]string{}
		for k, v := range cookieData {
			cookieFormQueryData[k] = v
		}
		for k, v := range formQueryData {
			cookieFormQueryData[k] = v
		}
		cookieStr, err := json.Marshal(&cookieFormQueryData)
		if err != nil {
			panic(err)
		}
		c.Session[listTemplate.Cookie.Name] = string(cookieStr)
	}
	cookieData = map[string]string{}
	cookieStr := c.Session[listTemplate.Cookie.Name]
	if cookieStr != "" {
		err := json.Unmarshal([]byte(cookieStr), &cookieData)
		if err != nil {
			panic(err)
		}
	}
	return paramMap, cookieData
}

func (c Console) getSysParam(sessionId int, userId int) map[string]interface{} {
	commonUtil := CommonUtil{}
	systemParameter := c.getSystemParameter(sessionId, userId)
	systemParameterMain := systemParameter["A"].(map[string]interface{})
	currencyTypeId := commonUtil.GetIntFromMap(systemParameterMain, "currencyTypeId")
	currencyType := c.getCurrencyType(sessionId, currencyTypeId)
	currencyTypeMain := currencyType["A"].(map[string]interface{})
	thousandsSeparator := ","
	if fmt.Sprint(systemParameterMain["thousandDecimals"]) == "1" {
		thousandsSeparator = ""
	}
	amtDecimals := commonUtil.GetIntFromMap(currencyTypeMain, "amtDecimals")
	upDecimals := commonUtil.GetIntFromMap(currencyTypeMain, "upDecimals")
	costDecimals := commonUtil.GetIntFromMap(systemParameterMain, "costDecimals")
	percentDecimals := commonUtil.GetIntFromMap(systemParameterMain, "percentDecimals")
	return map[string]interface{}{
		"localCurrency": map[string]interface{}{
			"prefix": currencyTypeMain["currencyTypeSign"],
			"decimalPlaces": amtDecimals - 1,
			"unitPriceDecimalPlaces": upDecimals - 1,
		},
		"unitCostDecimalPlaces": costDecimals - 1,
		"percentDecimalPlaces": percentDecimals - 1,
		"thousandsSeparator": thousandsSeparator,
	}
}

func (c Console) getSystemParameter(sessionId int, userId int) map[string]interface{} {
	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
	user := querySupport.FindByMapWithSessionExact(session, "SysUser", map[string]interface{}{
		"_id": userId,
	})
	userMain := user["A"].(map[string]interface{})
	systemParameter := querySupport.FindByMapWithSessionExact(session, "SystemParameter", map[string]interface{}{
		"A.createUnit": userMain["createUnit"],
	})
	return systemParameter
}

func (c Console) getCurrencyType(sessionId int, currencyTypeId int) map[string]interface{} {
	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
	currencyType := querySupport.FindByMapWithSessionExact(session, "CurrencyType", map[string]interface{}{
		"_id": currencyTypeId,
	})
	return currencyType
}

// TODO,by test
func (c Console) FormSchema() revel.Result {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)

	schemaName := c.Params.Get("@name")
	strId := c.Params.Get("id")
	formStatus := c.Params.Get("formStatus")
	copyFlag := c.Params.Get("copyFlag")

	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(schemaName)

	result := map[string]interface{}{
		"formTemplate": formTemplate,
		"id":           strId,
		"formStatus":   formStatus,
		"copyFlag":     copyFlag,
	}
	if formTemplate.DataSourceModelId != "" {
		// 光有formTemplate不行,还要有model的内容,才可以渲染数据
		modelTemplateFactory := ModelTemplateFactory{}
		dataSource := modelTemplateFactory.GetDataSource(formTemplate.DataSourceModelId)
		modelTemplateFactory.ClearReverseRelation(&dataSource)
		dataSourceByte, err := json.Marshal(&dataSource)
		if err != nil {
			panic(err)
		}
		result["dataSource"] = dataSource
		commonUtil := CommonUtil{}
		dataSourceJson := string(dataSourceByte)
		dataSourceJson = commonUtil.FilterJsonEmptyAttr(dataSourceJson)
		result["dataSourceJson"] = template.JS(dataSourceJson)
	}
	//toolbarBo
	toolbarBo := map[string]interface{}{}
	for i, item := range formTemplate.FormElemLi {
		if item.XMLName.Local == "toolbar" {
			toolbarBo[item.Toolbar.Name] = templateManager.GetToolbarBo(item.Toolbar)
		}
		// 加入主数据集tag,页面渲染用
		if item.XMLName.Local == "column-model" && item.ColumnModel.DataSetId == "A" {
			formTemplate.FormElemLi[i].RenderTag = item.ColumnModel.DataSetId + "_" + fmt.Sprint(i)
		}
	}
	result["toolbarBo"] = toolbarBo
	dataBo := map[string]interface{}{}
	relationBo := map[string]interface{}{}
	result["dataBo"] = dataBo
	result["relationBo"] = relationBo

	relationBoByte, err := json.Marshal(&relationBo)
	if err != nil {
		panic(err)
	}

	// 主数据集的后台渲染
	result["masterRenderLi"] = c.getMasterRenderLi(formTemplate)

	formTemplateJsonDataArray, err := json.Marshal(&formTemplate)
	if err != nil {
		panic(err)
	}

	dataBoByte, err := json.Marshal(&dataBo)
	if err != nil {
		panic(err)
	}

	layerBo := templateManager.GetLayerForFormTemplate(sessionId, formTemplate)
	iLayerBo := layerBo["layerBo"]
	layerBoByte, err := json.Marshal(&iLayerBo)
	if err != nil {
		panic(err)
	}
	iLayerBoLi := layerBo["layerBoLi"]
	layerBoLiByte, err := json.Marshal(&iLayerBoLi)
	if err != nil {
		panic(err)
	}

	commonUtil := CommonUtil{}
	userId := commonUtil.GetIntFromString(c.Session["userId"])
	sysParam := c.getSysParam(sessionId, userId)
	sysParamJson, err := json.Marshal(&sysParam)
	if err != nil {
		panic(err)
	}
	result["sysParamJson"] = template.JS(string(sysParamJson))
	
	formTemplateJsonData := string(formTemplateJsonDataArray)
	formTemplateJsonData = commonUtil.FilterJsonEmptyAttr(formTemplateJsonData)
	result["formTemplateJsonData"] = template.JS(formTemplateJsonData)
	dataBoJson := string(dataBoByte)
	dataBoJson = commonUtil.FilterJsonEmptyAttr(dataBoJson)
	result["dataBoJson"] = template.JS(dataBoJson)
	layerBoJson := string(layerBoByte)
	layerBoJson = commonUtil.FilterJsonEmptyAttr(layerBoJson)
	result["layerBoJson"] = template.JS(layerBoJson)
	layerBoLiJson := string(layerBoLiByte)
	layerBoLiJson = commonUtil.FilterJsonEmptyAttr(layerBoLiJson)
	result["layerBoLiJson"] = template.JS(layerBoLiJson)
	relationBoJson := string(relationBoByte)
	relationBoJson = commonUtil.FilterJsonEmptyAttr(relationBoJson)
	result["relationBoJson"] = template.JS(relationBoJson)

	viewPath := revel.Config.StringDefault("REVEL_VIEW_PATH", "")
	file, err := os.Open(viewPath + "/" + formTemplate.ViewTemplate.View)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	funcMap := map[string]interface{}{
		"eq": func(a, b interface{}) bool {
			return a == b
		},
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	tmpl, err := template.New("formSchema").Funcs(funcMap).Parse(string(fileContent))
	if err != nil {
		panic(err)
	}
	tmplResult := map[string]interface{}{
		"result": result,
		"flash": c.Flash.Out,
		"session": c.Session,
	}
	tmpl.Execute(c.Response.Out, tmplResult)
	return nil
}

func (c Console) getMasterRenderLi(formTemplate FormTemplate) map[string]interface{} {
	if formTemplate.DataSourceModelId == "" {
		return nil
	}
	result := map[string]interface{}{}

	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(formTemplate.DataSourceModelId)

	modelIterator := ModelIterator{}
	var message interface{} = ""
	for i, item := range formTemplate.FormElemLi {
		if item.XMLName.Local == "column-model" && item.ColumnModel.DataSetId == "A" {
			lineColSpan, err := strconv.Atoi(item.ColumnModel.ColSpan)
			if err != nil {
				lineColSpan = 1
			}
			container := [][]map[string]interface{}{}
			containerItem := []map[string]interface{}{}
			lineColSpanSum := 0
			for _, column := range item.ColumnModel.ColumnLi {
				if column.XMLName.Local == "html" {
					columnColSpan, err := strconv.Atoi(column.ColSpan)
					if err != nil {
						columnColSpan = 1
					}
					containerItem = append(containerItem, map[string]interface{}{
						"isHtml": "true",
						"html":   column.Html,
					})
					lineColSpanSum += columnColSpan
					if lineColSpanSum >= lineColSpan {
						container = append(container, containerItem)
						containerItem = []map[string]interface{}{}
						lineColSpanSum = lineColSpanSum - lineColSpan
					}
				} else {
					isModelField := false
					modelIterator.IterateAllField(&dataSource, &message, func(fieldGroup *FieldGroup, result *interface{}) {
						if fieldGroup.IsMasterField() && fieldGroup.Id == column.Name {
							isModelField = true
							if column.Hideable != "true" && column.ManualRender != "true" {
								columnColSpan, err := strconv.Atoi(column.ColSpan)
								if err != nil {
									columnColSpan = 1
								}
								containerItem = append(containerItem, map[string]interface{}{
									"isHtml":      "false",
									"required":    fmt.Sprint(fieldGroup.AllowEmpty == "false"),
									"label":       column.Text,
									"name":        column.Name,
									"columnWidth": column.ColumnWidth,
									"columnSpan":  columnColSpan - 1,
									"labelWidth":  column.LabelWidth,
								})
								lineColSpanSum += columnColSpan
								if lineColSpanSum >= lineColSpan {
									container = append(container, containerItem)
									containerItem = []map[string]interface{}{}
									lineColSpanSum = lineColSpanSum - lineColSpan
								}
							}
						}
					})
					if !isModelField {
						if column.Hideable != "true" && column.ManualRender != "true" {
							columnColSpan, err := strconv.Atoi(column.ColSpan)
							if err != nil {
								columnColSpan = 1
							}
							containerItem = append(containerItem, map[string]interface{}{
								"isHtml":      "false",
								"required":    false,
								"label":       column.Text,
								"name":        column.Name,
								"columnWidth": column.ColumnWidth,
								"columnSpan":  columnColSpan - 1,
								"labelWidth":  column.LabelWidth,
							})
							lineColSpanSum += columnColSpan
							if lineColSpanSum >= lineColSpan {
								container = append(container, containerItem)
								containerItem = []map[string]interface{}{}
								lineColSpanSum = lineColSpanSum - lineColSpan
							}
						}
					}
				}
			}
			if 0 < lineColSpanSum && lineColSpanSum < lineColSpan {
				container = append(container, containerItem)
			}
			result[item.DataSetId+"_"+fmt.Sprint(i)] = container
		}
	}

	return result
}

func (c Console) getQueryParameterRenderLi(listTemplate ListTemplate) [][]map[string]interface{} {
	lineColSpan := 6
	container := [][]map[string]interface{}{}
	containerItem := []map[string]interface{}{}
	lineColSpanSum := 0
	listTemplateIterator := ListTemplateIterator{}
	var result interface{} = ""
	listTemplateIterator.IterateTemplateQueryParameter(listTemplate, &result, func(queryParameter QueryParameter, result *interface{}) {
		if queryParameter.Editor != "hiddenfield" {
			columnColSpan := 2
			containerItem = append(containerItem, map[string]interface{}{
				"label": queryParameter.Text,
				"name":  queryParameter.Name,
			})
			lineColSpanSum += columnColSpan
			if lineColSpanSum >= lineColSpan {
				container = append(container, containerItem)
				containerItem = []map[string]interface{}{}
				lineColSpanSum = lineColSpanSum - lineColSpan
			}
		}
	})
	if 0 < lineColSpanSum && lineColSpanSum < lineColSpan {
		container = append(container, containerItem)
	}
	return container
}

func (c Console) Relation() revel.Result {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)

	selectorId := c.Params.Get("selectorId")
	id := c.Params.Get("id")

	templateManager := TemplateManager{}
	relationLi := []map[string]interface{}{
		map[string]interface{}{
			"selectorId": selectorId,
			"relationId": id,
		},
	}
	relationBo := templateManager.GetRelationBo(sessionId, relationLi)
	var result interface{} = nil
	var url interface{} = nil
	if relationBo[selectorId] != nil {
		selRelationBo := relationBo[selectorId].(map[string]interface{})
		if selRelationBo[id] != nil {
			result = selRelationBo[id]
			url = selRelationBo["url"]
		}
	}
	c.Response.ContentType = "application/json; charset=utf-8"
	return c.RenderJson(map[string]interface{}{
		"result": result,
		"url":    url,
	})
}

func (c Console) Refretor() revel.Result {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)

	refretorType := c.Params.Get("type")
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate("Console")

	if refretorType == "Component" {
		listTemplateInfoLi := templateManager.RefretorListTemplateInfo()
		items := getSummaryListTemplateInfoLi(listTemplateInfoLi)
		for _, item := range formTemplate.FormElemLi {
			if item.XMLName.Local == "column-model" && item.ColumnModel.Name == "Component" {
				itemsDict := templateManager.GetColumnModelDataForColumnModel(sessionId, item.ColumnModel, items)
				items = itemsDict["items"].([]interface{})
				break
			}
		}

		dataBo := map[string]interface{}{
			"items": items,
		}

		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataBo)
	}
	if refretorType == "Selector" {
		selectorTemplateInfoLi := templateManager.RefretorSelectorTemplateInfo()
		items := getSummarySelectorTemplateInfoLi(selectorTemplateInfoLi)
		for _, item := range formTemplate.FormElemLi {
			if item.XMLName.Local == "column-model" && item.ColumnModel.Name == "Selector" {
				itemsDict := templateManager.GetColumnModelDataForColumnModel(sessionId, item.ColumnModel, items)
				items = itemsDict["items"].([]interface{})
				break
			}
		}

		dataBo := map[string]interface{}{
			"items": items,
		}
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataBo)
	}
	if refretorType == "Form" {
		formTemplateInfoLi := templateManager.RefretorFormTemplateInfo()
		items := getSummaryFormTemplateInfoLi(formTemplateInfoLi)
		for _, item := range formTemplate.FormElemLi {
			if item.XMLName.Local == "column-model" && item.ColumnModel.Name == "Form" {
				itemsDict := templateManager.GetColumnModelDataForColumnModel(sessionId, item.ColumnModel, items)
				items = itemsDict["items"].([]interface{})
				break
			}
		}

		dataBo := map[string]interface{}{
			"items": items,
		}
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataBo)
	}
	if refretorType == "DataSource" {
		modelTemplateFactory := ModelTemplateFactory{}
		dataSourceTemplateInfoLi := modelTemplateFactory.RefretorDataSourceInfo()
		items := getSummaryDataSourceInfoLi(dataSourceTemplateInfoLi)
		for _, item := range formTemplate.FormElemLi {
			if item.XMLName.Local == "column-model" && item.ColumnModel.Name == "DataSource" {
				itemsDict := templateManager.GetColumnModelDataForColumnModel(sessionId, item.ColumnModel, items)
				items = itemsDict["items"].([]interface{})
				break
			}
		}

		dataBo := map[string]interface{}{
			"items": items,
		}
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(&dataBo)
	}
	c.Response.ContentType = "application/json; charset=utf-8"
	return c.RenderJson(map[string]interface{}{
		"message": "可能传入了错误的refretorType:" + refretorType,
	})
}

func (c Console) Xml() revel.Result {
	refretorType := c.Params.Get("type")
	id := c.Params.Get("@name")
	templateManager := TemplateManager{}

	if refretorType == "Component" {
		listTemplate := templateManager.GetListTemplate(id)
		c.Response.ContentType = "application/xml; charset=utf-8"
		return c.RenderXml(&listTemplate)
	}
	if refretorType == "Selector" {
		selectorTemplate := templateManager.GetSelectorTemplate(id)
		c.Response.ContentType = "application/xml; charset=utf-8"
		return c.RenderXml(&selectorTemplate)
	}
	if refretorType == "Form" {
		formTemplate := templateManager.GetFormTemplate(id)
		c.Response.ContentType = "application/xml; charset=utf-8"
		return c.RenderXml(&formTemplate)
	}
	if refretorType == "DataSource" {
		modelTemplateFactory := ModelTemplateFactory{}
		dataSourceTemplate := modelTemplateFactory.GetDataSource(id)
		c.Response.ContentType = "application/xml; charset=utf-8"
		return c.RenderXml(&dataSourceTemplate)
	}
	c.Response.ContentType = "application/json; charset=utf-8"
	return c.RenderJson(map[string]interface{}{
		"message": "可能传入了错误的refretorType:" + refretorType,
	})
}

func (c Console) RawXml() revel.Result {
	refretorType := c.Params.Get("type")
	id := c.Params.Get("@name")
	templateManager := TemplateManager{}

	if refretorType == "Component" {
		listTemplateInfo := templateManager.GetListTemplateInfo(id)
		listTemplate := ListTemplate{}
		file, err := os.Open(listTemplateInfo.Path)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = xml.Unmarshal(data, &listTemplate)
		if err != nil {
			panic(err)
		}

		c.Response.ContentType = "application/xml; charset=utf-8"
		return c.RenderXml(&listTemplate)
	}
	if refretorType == "Selector" {
		selectorTemplateInfo := templateManager.GetSelectorTemplateInfo(id)
		selectorTemplate := ListTemplate{}
		file, err := os.Open(selectorTemplateInfo.Path)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = xml.Unmarshal(data, &selectorTemplate)
		if err != nil {
			panic(err)
		}

		c.Response.ContentType = "application/xml; charset=utf-8"
		return c.RenderXml(&selectorTemplate)
	}
	if refretorType == "Form" {
		formTemplateInfo := templateManager.GetFormTemplateInfo(id)
		formTemplate := FormTemplate{}
		file, err := os.Open(formTemplateInfo.Path)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = xml.Unmarshal(data, &formTemplate)
		if err != nil {
			panic(err)
		}

		c.Response.ContentType = "application/xml; charset=utf-8"
		return c.RenderXml(&formTemplate)
	}
	c.Response.ContentType = "application/json; charset=utf-8"
	return c.RenderJson(map[string]interface{}{
		"message": "可能传入了错误的refretorType:" + refretorType,
	})
}
