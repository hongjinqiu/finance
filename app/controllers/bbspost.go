package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	. "com/papersns/common"
	. "com/papersns/component"
	. "com/papersns/error"
	"com/papersns/global"
	. "com/papersns/model"
	. "com/papersns/model/handler"
	. "com/papersns/mongo"
	"com/papersns/mongo"
	"fmt"
	"strconv"
	"strings"
)

func init() {
}

type BbsPostSupport struct {
	ActionSupport
}

func (c BbsPostSupport) RAfterSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	master := (*bo)["A"].(map[string]interface{})
	if fmt.Sprint(master["type"]) == "1" { // 主题帖
		c.bbsPostAfterSaveData(sessionId, dataSource, bo, diffDateRowLi)
	} else if fmt.Sprint(master["type"]) == "2" { // 主题帖回复
		c.bbsPostReplyAfterSaveData(sessionId, dataSource, bo, diffDateRowLi)
	}
}

func (c BbsPostSupport) bbsPostReplyAfterSaveData(sessionId int, dataSource DataSource, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	session, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)

	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	// 更新bbsPostId对应的主题帖的 lastReplyBy = self,lastReplyTime = currentTime,
	commonUtil := CommonUtil{}
	dateUtil := DateUtil{}
	master := (*bo)["A"].(map[string]interface{})
	qb := QuerySupport{}
	mainBbsPostQuery := map[string]interface{}{
		"_id": commonUtil.GetIntFromMap(master, "bbsPostId"),
	}
	bbsPostCollectionName := "BbsPost"
	mainBbsPost, found := qb.FindByMapWithSession(session, bbsPostCollectionName, mainBbsPostQuery)
	if !found {
		panic(BusinessError{Message: "主题帖未找到"})
	}
	mainBbsPostMaster := mainBbsPost["A"].(map[string]interface{})
	mainBbsPost["A"] = mainBbsPostMaster
	mainBbsPostMaster["lastReplyBy"] = userId
	mainBbsPostMaster["lastReplyTime"] = dateUtil.GetCurrentYyyyMMddHHmmss()
	if _, updateResult := txnManager.Update(txnId, bbsPostCollectionName, mainBbsPost); !updateResult {
		panic(BusinessError{Message: "主题帖更新失败"})
	}

	for _, item := range *diffDateRowLi {
		isNewOrUpdate := (item.SrcData != nil && item.DestData != nil) || (item.SrcData == nil && item.DestData != nil)
		if isNewOrUpdate {
			// 旧数据反过账,新数据正过账,修改后,更新 bbsPostId,readBy,bbsPostId
			bbsPostId := commonUtil.GetIntFromMap(*item.DestData, "bbsPostId")
			c.addOrUpdateBbsPostRead(sessionId, bbsPostId)
		}
	}
}

func (c BbsPostSupport) bbsPostAfterSaveData(sessionId int, dataSource DataSource, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	commonUtil := CommonUtil{}
	for _, item := range *diffDateRowLi {
		if item.SrcData != nil && item.DestData != nil { // 修改
			// 旧数据反过账,新数据正过账,修改后,更新 bbsPostId,readBy,bbsPostId
			bbsPostId := commonUtil.GetIntFromMap(*item.DestData, "id")
			c.addOrUpdateBbsPostRead(sessionId, bbsPostId)
		} else if item.SrcData == nil && item.DestData != nil { // 新增
			// 新数据正过账,新增记阅读记录,
			bbsPostId := commonUtil.GetIntFromMap(*item.DestData, "id")
			c.addBbsPostRead(sessionId, bbsPostId)
		}
	}
}

func (c BbsPostSupport) addOrUpdateBbsPostRead(sessionId int, bbsPostId int) {
	session, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	bbsPost := BbsPost{}
	modelTemplateFactory := ModelTemplateFactory{}
	bbsPostReadDS := modelTemplateFactory.GetDataSource("BbsPostRead")

	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	dateUtil := DateUtil{}
	qb := QuerySupport{}
	
	bbsPostRead, found := qb.FindByMapWithSession(session, "BbsPostRead", map[string]interface{}{
		"A.bbsPostId": bbsPostId,
		"A.readBy":    userId,
	})
	if found {
		bbsPost.RSetModifyFixFieldValue(sessionId, bbsPostReadDS, &bbsPostRead)
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

func (c BbsPostSupport) addBbsPostRead(sessionId int, bbsPostId int) {
	_, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	bbsPost := BbsPost{}
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
	bbsPost.RSetCreateFixFieldValue(sessionId, bbsPostReadDS, &bbsPostRead)
	txnManager.Insert(txnId, "BbsPostRead", bbsPostRead)
}

func (o BbsPostSupport) RBeforeDeleteData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{})   {
	session, _ := global.GetConnection(sessionId)
	// 已回复过的帖子不可删除
	qb := QuerySupport{}
	bbsPostCollectionName := "BbsPost"
	bbsPostReplyQuery := map[string]interface{}{
		"A.bbsPostId": (*bo)["id"],
	}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		bbsPostReplyQuery[k] = v
	}
	
	_, found := qb.FindByMapWithSession(session, bbsPostCollectionName, bbsPostReplyQuery)
	if found {
		panic(BusinessError{Message: "存在回复的主题帖不可删除"})
	}
}

func (c BbsPostSupport) RAfterDeleteData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	// 反过账
	_, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}

	query := map[string]interface{}{
		"A.bbsPostId": (*bo)["id"],
		"A.readBy":    userId,
	}
	txnManager.RemoveAll(txnId, "BbsPostRead", query)
}

type BbsPost struct {
	BaseDataAction
}

func (c BbsPost) SaveData() revel.Result {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.RSaveCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BbsPost) DeleteData() revel.Result {
	c.RActionSupport = BbsPostSupport{}

	modelRenderVO := c.RDeleteDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BbsPost) EditData() revel.Result {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.REditDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BbsPost) NewData() revel.Result {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.RNewDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BbsPost) GetData() revel.Result {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.RGetDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BbsPost) CopyData() revel.Result {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.RCopyDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BbsPost) GiveUpData() revel.Result {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.RGiveUpDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BbsPost) RefreshData() revel.Result {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.RRefreshDataCommon()
	return c.RRenderCommon(modelRenderVO)
}

func (c BbsPost) LogList() revel.Result {
	result := c.RLogListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
