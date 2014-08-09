package console

import "github.com/robfig/revel"
import (
	. "com/papersns/common"
	. "com/papersns/component"
	. "com/papersns/error"
	"com/papersns/global"
	. "com/papersns/model"
	. "com/papersns/mongo"
	"com/papersns/mongo"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// 管理员查看页面,设置session.userId,以查看数据,
func (c Console) BbsPostAdminReplySchema() revel.Result {
	// 取一下bbsPostId

	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", c.Session["userId"])
	global.SetGlobalAttr(sessionId, "adminUserId", c.Session["adminUserId"])
	defer global.CloseSession(sessionId)
	defer c.rollbackTxn(sessionId)

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
	c.commitTxn(sessionId)

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
	defer c.rollbackTxn(sessionId)

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
	c.commitTxn(sessionId)

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

func (c Console) commitTxn(sessionId int) {
	txnId := global.GetGlobalAttr(sessionId, "txnId")
	if txnId != nil {
		_, db := global.GetConnection(sessionId)
		txnManager := TxnManager{db}
		txnManager.Commit(txnId.(int))
	}
}

func (c Console) rollbackTxn(sessionId int) {
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
		c.setModifyFixFieldValue(sessionId, bbsPostReadDS, &bbsPostRead)
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
	c.setCreateFixFieldValue(sessionId, bbsPostReadDS, &bbsPostRead)
	txnManager.Insert(txnId, "BbsPostRead", bbsPostRead)
}

func (c Console) setCreateFixFieldValue(sessionId int, dataSource DataSource, bo *map[string]interface{}) {
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

func (c Console) setModifyFixFieldValue(sessionId int, dataSource DataSource, bo *map[string]interface{}) {
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
	log.Println("setModifyFixFieldValue,collectionName:" + collectionName + ", query:" + string(srcQueryByte))
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
