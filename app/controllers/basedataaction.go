package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/component"
	. "com/papersns/model"
	. "com/papersns/model/handler"
	"com/papersns/mongo"
	. "com/papersns/mongo"
	"com/papersns/global"
	"encoding/json"
	"strconv"
	"strings"
	"time"
	"fmt"
)

func init() {
}

type IActionSupport interface {
	beforeNewData(sessionId int, dataSource DataSource)
	afterNewData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	beforeCopyData(sessionId int, dataSource DataSource, srcBo map[string]interface{})
	afterCopyData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	editValidate(sessionId int, dataSource DataSource, bo map[string]interface{}) (string, bool)
	beforeEditData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	afterEditData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	beforeSaveData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	afterSaveData(sessionId int, dataSource DataSource, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow)
	beforeGiveUpData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	afterGiveUpData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	beforeDeleteData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	afterDeleteData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	beforeRefreshData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	afterRefreshData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	beforeCancelData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	afterCancelData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	beforeUnCancelData(sessionId int, dataSource DataSource, bo *map[string]interface{})
	afterUnCancelData(sessionId int, dataSource DataSource, bo *map[string]interface{})
}

type ActionSupport struct{}

func (o ActionSupport) beforeNewData(sessionId int, dataSource DataSource)                                          {}
func (o ActionSupport) afterNewData(sessionId int, dataSource DataSource, bo *map[string]interface{})               {}
func (o ActionSupport) beforeCopyData(sessionId int, dataSource DataSource, srcBo map[string]interface{})           {}
func (o ActionSupport) afterCopyData(sessionId int, dataSource DataSource, bo *map[string]interface{})              {}
func (o ActionSupport) editValidate(sessionId int, dataSource DataSource, bo map[string]interface{}) (string, bool) {
	return "", true
}
func (o ActionSupport) beforeEditData(sessionId int, dataSource DataSource, bo *map[string]interface{})             {}
func (o ActionSupport) afterEditData(sessionId int, dataSource DataSource, bo *map[string]interface{}) {
}
func (o ActionSupport) beforeSaveData(sessionId int, dataSource DataSource, bo *map[string]interface{}) {}
func (o ActionSupport) afterSaveData(sessionId int, dataSource DataSource, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
}
func (o ActionSupport) beforeGiveUpData(sessionId int, dataSource DataSource, bo *map[string]interface{})   {}
func (o ActionSupport) afterGiveUpData(sessionId int, dataSource DataSource, bo *map[string]interface{})    {}
func (o ActionSupport) beforeDeleteData(sessionId int, dataSource DataSource, bo *map[string]interface{})   {}
func (o ActionSupport) afterDeleteData(sessionId int, dataSource DataSource, bo *map[string]interface{})    {}
func (o ActionSupport) beforeRefreshData(sessionId int, dataSource DataSource, bo *map[string]interface{})  {}
func (o ActionSupport) afterRefreshData(sessionId int, dataSource DataSource, bo *map[string]interface{})   {}
func (o ActionSupport) beforeCancelData(sessionId int, dataSource DataSource, bo *map[string]interface{})   {}
func (o ActionSupport) afterCancelData(sessionId int, dataSource DataSource, bo *map[string]interface{})    {}
func (o ActionSupport) beforeUnCancelData(sessionId int, dataSource DataSource, bo *map[string]interface{}) {}
func (o ActionSupport) afterUnCancelData(sessionId int, dataSource DataSource, bo *map[string]interface{})  {}

type BaseDataAction struct {
	*revel.Controller
	actionSupport IActionSupport
}

func (c BaseDataAction) setCreateFixFieldValue(sessionId int, dataSource DataSource, bo *map[string]interface{}) {
	var result interface{} = ""
	userId, err := strconv.Atoi(c.Session["userId"])
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
	modelIterator := ModelIterator{}
	modelIterator.IterateDataBo(dataSource, bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}){
		(*data)["createBy"] = userId
		(*data)["createTime"] = createTime
		(*data)["createUnit"] = sysUser["createUnit"]
	})
}

func (c BaseDataAction) setModifyFixFieldValue(sessionId int, dataSource DataSource, bo *map[string]interface{}) {
	var result interface{} = ""
	userId, err := strconv.Atoi(c.Session["userId"])
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
	
	srcBo := map[string]interface{}{}
	srcQuery := map[string]interface{}{
		"_id": (*bo)["_id"],
	}
	modelTemplateFactory := ModelTemplateFactory{}
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	db.C(collectionName).Find(srcQuery).One(&srcBo)
	modelIterator := ModelIterator{}
	modelIterator.IterateDiffBo(dataSource, bo, srcBo, &result, func(fieldGroupLi []FieldGroup, destData *map[string]interface{}, srcData map[string]interface{}, result *interface{}){
		if destData != nil && srcData == nil {
			(*destData)["createBy"] = userId
			(*destData)["createTime"] = modifyTime
			(*destData)["createUnit"] = sysUser["createUnit"]
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
				(*destData)["modifyUnit"] = sysUser["createUnit"]
			}
		}
	})
}

func (c BaseDataAction) rollbackTxn(sessionId int) {
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

func (c BaseDataAction) commitTxn(sessionId int) {
	txnId := global.GetGlobalAttr(sessionId, "txnId")
	if txnId != nil {
		_, db := global.GetConnection(sessionId)
		txnManager := TxnManager{db}
		txnManager.Commit(txnId.(int))
	}
}

func (c BaseDataAction) renderCommon(bo map[string]interface{}, dataSource DataSource) revel.Result {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateAllFieldBo(dataSource, &bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}){
		if (*data)[fieldGroup.Id] != nil {
			(*data)[fieldGroup.Id] = fmt.Sprint((*data)[fieldGroup.Id])
		}
	})
	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(map[string]interface{}{
			"bo":         bo,
			"dataSource": dataSource,
		})
	}
	return c.Render()
}

/**
 * 列表页
 */
//func (baseData BaseDataAction) ListData() revel.Result {
//
//}

/**
 * 新增
 */
func (c BaseDataAction) NewData() revel.Result {
	c.actionSupport = ActionSupport{}
	
	bo, dataSource := c.newDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BaseDataAction) newDataCommon() (map[string]interface{}, DataSource) {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)
	
	dataSourceModelId := c.Params.Get("dataSourceModelId")
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	c.actionSupport.beforeNewData(sessionId, dataSource)
	bo := modelTemplateFactory.GetInstanceByDS(dataSource)
	c.actionSupport.afterNewData(sessionId, dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)

	c.commitTxn(sessionId)
	return bo, dataSource
}

func (c BaseDataAction) GetData() revel.Result {
	bo, dataSource := c.getDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BaseDataAction) getDataCommon() (map[string]interface{}, DataSource) {
	dataSourceModelId := c.Params.Get("dataSourceModelId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}

	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	bo, found := querySupport.FindByMap(collectionName, queryMap)
	if !found {
		panic("GetData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	
	modelTemplateFactory.ConvertDataType(dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	return bo, dataSource
}

/**
 * 复制
 */
func (c BaseDataAction) CopyData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, dataSource := c.copyDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BaseDataAction) copyDataCommon() (map[string]interface{}, DataSource) {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	dataSourceModelId := c.Params.Get("dataSourceModelId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}

	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	srcBo, found := querySupport.FindByMap(collectionName, queryMap)
	if !found {
		panic("CopyData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	
	modelTemplateFactory.ConvertDataType(dataSource, &srcBo)
	c.actionSupport.beforeCopyData(sessionId, dataSource, srcBo)
	dataSource, bo := modelTemplateFactory.GetCopyInstance(dataSourceModelId, srcBo)
	c.actionSupport.afterCopyData(sessionId, dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return bo, dataSource
}

/**
 * 修改,只取数据,没涉及到数据库保存,涉及到数据库保存的方法是SaveData,
 */
func (c BaseDataAction) EditData() revel.Result {
	c.actionSupport = ActionSupport{}
	
	bo, dataSource := c.editDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BaseDataAction) editDataCommon() (map[string]interface{}, DataSource) {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	dataSourceModelId := c.Params.Get("dataSourceModelId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	bo, found := querySupport.FindByMap(collectionName, queryMap)
	if !found {
		panic("RefreshData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	editMessage, isValid := c.actionSupport.editValidate(sessionId, dataSource, bo)
	if !isValid {
		panic(editMessage)
	}

	c.actionSupport.beforeEditData(sessionId, dataSource, &bo)
	c.actionSupport.afterEditData(sessionId, dataSource, &bo)
	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return bo, dataSource
}

/**
 * 保存
 */
func (c BaseDataAction) SaveData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, dataSource := c.saveCommon()

	return c.renderCommon(bo, dataSource)
}

func (c BaseDataAction) saveCommon() (map[string]interface{}, DataSource) {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	dataSourceModelId := c.Params.Form.Get("dataSourceModelId")
	jsonBo := c.Params.Form.Get("jsonData")

	bo := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonBo), &bo)
	if err != nil {
		panic(err)
	}
	
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	strId := modelTemplateFactory.GetStrId(bo)
	if strId == "" || strId == "0" {
		c.setCreateFixFieldValue(sessionId, dataSource, &bo)
	} else {
		c.setModifyFixFieldValue(sessionId, dataSource, &bo)
		editMessage, isValid := c.actionSupport.editValidate(sessionId, dataSource, bo)
		if !isValid {
			panic(editMessage)
		}
	}
	
	c.actionSupport.beforeSaveData(sessionId, dataSource, &bo)
	financeService := FinanceService{}

	diffDataRowLi := financeService.SaveData(sessionId, dataSource, &bo)

	c.actionSupport.afterSaveData(sessionId, dataSource, &bo, diffDataRowLi)
	modelTemplateFactory.ClearReverseRelation(&dataSource)
	
	c.commitTxn(sessionId)
	
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": bo["_id"],
	}
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	bo, _ = querySupport.FindByMap(collectionName, queryMap)
	return bo, dataSource
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BaseDataAction) GiveUpData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, dataSource := c.giveUpDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BaseDataAction) giveUpDataCommon() (map[string]interface{}, DataSource) {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	dataSourceModelId := c.Params.Get("dataSourceModelId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}

	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	bo, found := querySupport.FindByMap(collectionName, queryMap)
	if !found {
		panic("giveUpData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	c.actionSupport.beforeGiveUpData(sessionId, dataSource, &bo)
	c.actionSupport.afterGiveUpData(sessionId, dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return bo, dataSource
}

/**
 * 删除
 */
func (c BaseDataAction) DeleteData() revel.Result {
	c.actionSupport = ActionSupport{}
	
	bo, dataSource := c.deleteDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BaseDataAction) deleteDataCommon() (map[string]interface{}, DataSource) {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	dataSourceModelId := c.Params.Get("dataSourceModelId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}

	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	bo, found := querySupport.FindByMap(collectionName, queryMap)
	if !found {
		panic("DeleteData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	c.actionSupport.beforeDeleteData(sessionId, dataSource, &bo)

	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()

	usedCheck := UsedCheck{}
	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateDataBo(dataSource, &bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		if fieldGroupLi[0].IsMasterField() {
			usedCheck.DeleteAll(sessionId, fieldGroupLi, *data)
		}
	})

	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	_, removeResult := txnManager.Remove(txnId, dataSourceModelId, bo)
	if !removeResult {
		panic("删除失败")
	}
	
	c.actionSupport.afterDeleteData(sessionId, dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return bo, dataSource
}

/**
 * 刷新
 */
func (c BaseDataAction) RefreshData() revel.Result {
	c.actionSupport = ActionSupport{}
	bo, dataSource := c.refreshDataCommon()
	
	return c.renderCommon(bo, dataSource)
}

func (c BaseDataAction) refreshDataCommon() (map[string]interface{}, DataSource) {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

	dataSourceModelId := c.Params.Get("dataSourceModelId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}

	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"_id": id,
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	bo, found := querySupport.FindByMap(collectionName, queryMap)
	if !found {
		panic("RefreshData, dataSouceModelId=" + dataSourceModelId + ", id=" + strId + " not found")
	}
	
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	c.actionSupport.beforeRefreshData(sessionId, dataSource, &bo)
	c.actionSupport.afterRefreshData(sessionId, dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.commitTxn(sessionId)
	return bo, dataSource
}

/**
 * 被用查询
 */
func (c BaseDataAction) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}

func (c BaseDataAction) logListCommon() map[string]interface{} {
	dataSourceModelId := c.Params.Get("dataSourceModelId")
	strId := c.Params.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}

	collectionName := "PubReferenceLog"
	// reference,beReference
	querySupport := QuerySupport{}
	query := map[string]interface{}{
		"beReference": []interface{}{dataSourceModelId, "A", "id", id},
	}
	pageNo := 1
	pageSize := 10
	orderBy := ""
	return querySupport.Index(collectionName, query, pageNo, pageSize, orderBy)
}
