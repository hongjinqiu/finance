package controllers

import (
	. "com/papersns/common"
	. "com/papersns/component"
	. "com/papersns/error"
	"com/papersns/global"
	. "com/papersns/lock"
	. "com/papersns/mongo"
	"com/papersns/mongo"
	. "com/papersns/taobao"
	"fmt"
	"log"
)

type LoginService struct{}

func (o LoginService) DealLoginTest(sessionId int, url string) (resStruct map[string]interface{}, userId int, isStep bool) {
	isStep = false

//	taobaoInterface := TaobaoInterface{}
//	resStruct = taobaoInterface.GetUserInfo(url)
	/*
	{
		"oAuthInfo":   rsp,
		"top_session": rsp["access_token"],
		"app_key":     paramDict["client_id"],
		"top_appkey":  paramDict["client_id"],
		"topParameter": map[string]interface{}{
			"visitor_nick": rsp["taobao_user_nick"],
			"visitor_id":   rsp["taobao_user_id"],
			"sub_taobao_user_id": rsp["sub_taobao_user_id"],
			"sub_taobao_user_nick": rsp["sub_taobao_user_nick"],
		},
	}
	*/
	resStruct = map[string]interface{}{
		"oAuthInfo":   map[string]interface{}{
			"access_token": "test_token",
			"taobao_user_nick": "测试帐户20昵称",
			"taobao_user_id": 123456,
		},
		"top_session": "test_token",
		"app_key":     "21210514",
		"top_appkey":  "21210514",
		"topParameter": map[string]interface{}{
			"visitor_nick": "测试帐户20昵称",
			"visitor_id":   123456,
//			"sub_taobao_user_id": rsp["sub_taobao_user_id"],
//			"sub_taobao_user_nick": rsp["sub_taobao_user_nick"],
		},
	}
	topParameter := resStruct["topParameter"].(map[string]interface{})

	username := fmt.Sprint(topParameter["visitor_id"])
	nick := fmt.Sprint(topParameter["visitor_nick"])
	appKey := fmt.Sprint(resStruct["app_key"])

	session, db := global.GetConnection(sessionId)
	qb := QuerySupport{}
	query := map[string]interface{}{
		"A.nick": nick,
	}
	sysUser, found := qb.FindByMapWithSession(session, "SysUser", query)
	lockKey := ""
	if found {
		sysUserMaster := sysUser["A"].(map[string]interface{})
		lockKey = fmt.Sprint(CommonUtil{}.GetIntFromMap(sysUserMaster, "createUnit"))
	} else {
		lockKey = fmt.Sprint(nick)
	}

	// 加锁
	lockService := LockService{}
	unitLock := lockService.GetUnitLock(lockKey)
	(*unitLock).Lock()
	defer (*unitLock).Unlock()

	//	isDeptFirst := false
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	commonUtil := CommonUtil{}
	if !found {
		//taobaoShop := taobaoInterface.TaobaoShopGet(resStruct)
		taobaoShop := map[string]interface{}{
			"shop_get_response": map[string]interface{}{
				"shop": map[string]interface{}{
					"sid": "sid20",
					"nick": "管理员20_nick",
				},
				
			},
		}
		shoGetResponse := taobaoShop["shop_get_response"].(map[string]interface{})
		shop := shoGetResponse["shop"].(map[string]interface{})
		sid := fmt.Sprint(shop["sid"])
		userNick := fmt.Sprint(shop["nick"])
		unit, unitFound := qb.FindByMapWithSession(session, "SysUnit", map[string]interface{}{
			"A.sid": sid,
		})
		if !unitFound {
			//			isDeptFirst = true
			cid := fmt.Sprint(shop["cid"])
			title := fmt.Sprint(shop["title"])
			id := mongo.GetSequenceNo(db, "sysUnitId")
			unit = map[string]interface{}{
				"_id": id,
				"id":  id,
				"A": map[string]interface{}{
					"id":          id,
					"code":        sid,
					"name":        title,
					"sid":         sid,
					"cid":         cid,
					"userNick":    userNick,
					"createBy":    0,
					"createTime":  DateUtil{}.GetCurrentYyyyMMddHHmmss(),
					"createUnit":  id,
					"modifyBy":    0,
					"modifyTime":  0,
					"modifyUnit":  0,
					"attachCount": 0,
					"remark":      "",
				},
			}
			txnManager.Insert(txnId, "SysUnit", unit)
		} else {
			//			isDeptFirst = false
		}
		id := mongo.GetSequenceNo(db, "sysUserId")
		sysUser = map[string]interface{}{
			"_id": id,
			"id":  id,
			"A": map[string]interface{}{
				"id":          id,
				"code":        fmt.Sprint(username),
				"name":        username,
				"type":        2,
				"status":      1,
				"nick":        nick,
				"createBy":    id,
				"createTime":  DateUtil{}.GetCurrentYyyyMMddHHmmss(),
				"createUnit":  unit["id"],
				"modifyBy":    0,
				"modifyTime":  0,
				"modifyUnit":  0,
				"attachCount": 0,
				"remark":      "",
			},
		}
		txnManager.Insert(txnId, "SysUser", sysUser)
		unitA := unit["A"].(map[string]interface{})
		unit["A"] = unitA
		unitA["createBy"] = id
		unitId := commonUtil.GetIntFromMap(unit, "id")
		isStep = o.InitStep(sessionId, unitId, id, appKey)
	} else {
		//		println(sysUser)
		sysUserMaster := sysUser["A"].(map[string]interface{})
		unitId := sysUserMaster["createUnit"].(int)
		id := sysUserMaster["id"].(int)
		isStep = o.InitStep(sessionId, unitId, id, appKey)
	}
	// 处理子账号
	if topParameter["sub_taobao_user_nick"] != nil {
		nick := fmt.Sprint(topParameter["sub_taobao_user_nick"])
		query := map[string]interface{}{
			"A.nick": nick,
		}
		subSysUser, found := qb.FindByMapWithSession(session, "SysUser", query)
		if !found {
			id := mongo.GetSequenceNo(db, "sysUserId")
			sysUserMaster := sysUser["A"].(map[string]interface{})
			subSysUser = map[string]interface{}{
				"_id": id,
				"id":  id,
				"A": map[string]interface{}{
					"id":          id,
					"code":        fmt.Sprint(username),
					"name":        username,
					"type":        2,
					"status":      1,
					"nick":        nick,
					"createBy":    id,
					"createTime":  DateUtil{}.GetCurrentYyyyMMddHHmmss(),
					"createUnit":  sysUserMaster["createUnit"],
					"modifyBy":    0,
					"modifyTime":  0,
					"modifyUnit":  0,
					"attachCount": 0,
					"remark":      "",
				},
			}
			txnManager.Insert(txnId, "SysUser", subSysUser)
		}
		sysUser = subSysUser
	}
	//	c.commitTxn(sessionId)
	userId = commonUtil.GetIntFromMap(sysUser, "id")
	if isStep {
		go StepService{}.Run(sysUser)
	} else {
	}
	// sync data thread,暂时不用处理,
	sysUserMaster := sysUser["A"].(map[string]interface{})
	sysUnitId := commonUtil.GetIntFromMap(sysUserMaster, "createUnit")
	o.saveOrUpdateLastSessionData(sessionId, resStruct, sysUnitId, userId)
	return resStruct, userId, isStep
}

/**
@param url /asdfas/zasdfasdf/?param1=value1&param2=value2
*/
func (o LoginService) DealLogin(sessionId int, url string) (resStruct map[string]interface{}, userId int, isStep bool) {
	log.Print("receive top login request, url is:", url)
	isStep = false

	taobaoInterface := TaobaoInterface{}
	resStruct = taobaoInterface.GetUserInfo(url)
	topParameter := resStruct["topParameter"].(map[string]interface{})

	username := fmt.Sprint(topParameter["visitor_id"])
	nick := fmt.Sprint(topParameter["visitor_nick"])
	appKey := fmt.Sprint(resStruct["app_key"])

	session, db := global.GetConnection(sessionId)
	qb := QuerySupport{}
	query := map[string]interface{}{
		"A.nick": nick,
	}
	sysUser, found := qb.FindByMapWithSession(session, "SysUser", query)
	lockKey := ""
	if found {
		sysUserMaster := sysUser["A"].(map[string]interface{})
		lockKey = fmt.Sprint(CommonUtil{}.GetIntFromMap(sysUserMaster, "createUnit"))
	} else {
		lockKey = fmt.Sprint(nick)
	}

	// 加锁
	lockService := LockService{}
	unitLock := lockService.GetUnitLock(lockKey)
	(*unitLock).Lock()
	defer (*unitLock).Unlock()

	//	isDeptFirst := false
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	commonUtil := CommonUtil{}
	if !found {
		taobaoShop := taobaoInterface.TaobaoShopGet(resStruct)
		shoGetResponse := taobaoShop["shop_get_response"].(map[string]interface{})
		shop := shoGetResponse["shop"].(map[string]interface{})
		sid := fmt.Sprint(shop["sid"])
		userNick := fmt.Sprint(shop["nick"])
		unit, unitFound := qb.FindByMapWithSession(session, "SysUnit", map[string]interface{}{
			"A.sid": sid,
		})
		if !unitFound {
			//			isDeptFirst = true
			cid := fmt.Sprint(shop["cid"])
			title := fmt.Sprint(shop["title"])
			id := mongo.GetSequenceNo(db, "sysUnitId")
			unit = map[string]interface{}{
				"_id": id,
				"id":  id,
				"A": map[string]interface{}{
					"id":          id,
					"code":        sid,
					"name":        title,
					"sid":         sid,
					"cid":         cid,
					"userNick":    userNick,
					"createBy":    0,
					"createTime":  DateUtil{}.GetCurrentYyyyMMddHHmmss(),
					"createUnit":  id,
					"modifyBy":    0,
					"modifyTime":  0,
					"modifyUnit":  0,
					"attachCount": 0,
					"remark":      "",
				},
			}
			txnManager.Insert(txnId, "SysUnit", unit)
		} else {
			//			isDeptFirst = false
		}
		id := mongo.GetSequenceNo(db, "sysUserId")
		sysUser = map[string]interface{}{
			"_id": id,
			"id":  id,
			"A": map[string]interface{}{
				"id":          id,
				"code":        fmt.Sprint(username),
				"name":        username,
				"type":        2,
				"status":      1,
				"nick":        nick,
				"createBy":    id,
				"createTime":  DateUtil{}.GetCurrentYyyyMMddHHmmss(),
				"createUnit":  unit["id"],
				"modifyBy":    0,
				"modifyTime":  0,
				"modifyUnit":  0,
				"attachCount": 0,
				"remark":      "",
			},
		}
		txnManager.Insert(txnId, "SysUser", sysUser)
		unitA := unit["A"].(map[string]interface{})
		unit["A"] = unitA
		unitA["createBy"] = id
		unitId := commonUtil.GetIntFromMap(unit, "id")
		isStep = o.InitStep(sessionId, unitId, id, appKey)
	} else {
		//		println(sysUser)
		sysUserMaster := sysUser["A"].(map[string]interface{})
		unitId := sysUserMaster["createUnit"].(int)
		id := sysUserMaster["id"].(int)
		isStep = o.InitStep(sessionId, unitId, id, appKey)
	}
	// 处理子账号
	if topParameter["sub_taobao_user_nick"] != nil {
		nick := fmt.Sprint(topParameter["sub_taobao_user_nick"])
		query := map[string]interface{}{
			"A.nick": nick,
		}
		subSysUser, found := qb.FindByMapWithSession(session, "SysUser", query)
		if !found {
			id := mongo.GetSequenceNo(db, "sysUserId")
			sysUserMaster := sysUser["A"].(map[string]interface{})
			subSysUser = map[string]interface{}{
				"_id": id,
				"id":  id,
				"A": map[string]interface{}{
					"id":          id,
					"code":        fmt.Sprint(username),
					"name":        username,
					"type":        2,
					"status":      1,
					"nick":        nick,
					"createBy":    id,
					"createTime":  DateUtil{}.GetCurrentYyyyMMddHHmmss(),
					"createUnit":  sysUserMaster["createUnit"],
					"modifyBy":    0,
					"modifyTime":  0,
					"modifyUnit":  0,
					"attachCount": 0,
					"remark":      "",
				},
			}
			txnManager.Insert(txnId, "SysUser", subSysUser)
		}
		sysUser = subSysUser
	}
	//	c.commitTxn(sessionId)
	userId = commonUtil.GetIntFromMap(sysUser, "id")
	if isStep {
		go StepService{}.Run(sysUser)
	}
	// sync data thread,暂时不用处理,
	sysUserMaster := sysUser["A"].(map[string]interface{})
	sysUnitId := commonUtil.GetIntFromMap(sysUserMaster, "createUnit")
	o.saveOrUpdateLastSessionData(sessionId, resStruct, sysUnitId, userId)
	return resStruct, userId, isStep
}

func (o LoginService) saveOrUpdateLastSessionData(sessionId int, resStruct map[string]interface{}, sysUnitId int, sysUserId int) {
	session, db := global.GetConnection(sessionId)
	qb := QuerySupport{}
	lastSessionData, found := qb.FindByMapWithSession(session, "LastSessionData", map[string]interface{}{
		"A.sysUserId": sysUserId,
		"A.sysUnitId": sysUnitId,
	})
	txnManager := TxnManager{db}
	if !found {
		id := mongo.GetSequenceNo(db, "lastSessionDataId")
		txnId := global.GetTxnId(sessionId)
		lastSessionData := map[string]interface{}{
			"_id": id,
			"id":  id,
			"A": map[string]interface{}{
				"id":          id,
				"sysUserId":   sysUserId,
				"sysUnitId":   sysUnitId,
				"resStruct":   resStruct,
				"createBy":    sysUserId,
				"createTime":  DateUtil{}.GetCurrentYyyyMMddHHmmss(),
				"createUnit":  sysUnitId,
				"modifyBy":    0,
				"modifyTime":  0,
				"modifyUnit":  0,
				"attachCount": 0,
				"remark":      "",
			},
		}
		txnManager.Insert(txnId, "LastSessionData", lastSessionData)
	} else {
		txnId := global.GetTxnId(sessionId)
		lastSessionDataMaster := lastSessionData["A"].(map[string]interface{})
		lastSessionData["A"] = lastSessionDataMaster

		lastSessionDataMaster["modifyBy"] = sysUserId
		lastSessionDataMaster["modifyTime"] = DateUtil{}.GetCurrentYyyyMMddHHmmss()
		lastSessionDataMaster["modifyBy"] = sysUnitId

		_, updateResult := txnManager.Update(txnId, "LastSessionData", lastSessionData)
		if !updateResult {
			panic(BusinessError{Message: "更新LastSessionData失败"})
		}
	}
}

/**
appKey参数为为了应付有可能存在的多的项目而添加的
*/
func (o LoginService) GetStepTypeLi(appKey string) []int {
	return []int{3, 5, 6, 7, 9, 12, 14, 15, 16, 18, 19, 20}
}

/**
返回所有的初始化项目
*/
func (o LoginService) GetStepLi() []interface{} {
	li := []interface{}{}
	li = append(li, []interface{}{"初始化供应商类别", 3})
	li = append(li, []interface{}{"初始化币别", 5})
	li = append(li, []interface{}{"初始化银行资料", 6})
	li = append(li, []interface{}{"初始化计量单位", 7})
	li = append(li, []interface{}{"初始化客户类别", 9})
	li = append(li, []interface{}{"初始化税率类别", 12})
	li = append(li, []interface{}{"初始化收入费用类别", 14})
	li = append(li, []interface{}{"初始化收入费用项目", 15})
	li = append(li, []interface{}{"初始化会计期", 16})
	li = append(li, []interface{}{"初始化收款单类型参数", 18})
	li = append(li, []interface{}{"初始化付款单类型参数", 19})
	li = append(li, []interface{}{"初始化系统参数", 20})

	return li
}

func (o LoginService) InitStep(sessionId int, unitId int, userId int, appKey string) bool {
	session, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)

	isStep := false
	li := o.GetStepLi()
	stepTypeLi := o.GetStepTypeLi(appKey)
	stepCount, err := db.C("SysStep").Find(map[string]interface{}{
		"A.sysUnitId": unitId,
		"A.type": map[string]interface{}{
			"$in": stepTypeLi,
		},
	}).Count()
	if err != nil {
		panic(err)
	}
	if stepCount != len(stepTypeLi) {
		isStep = true
		stepLi := []interface{}{}
		for _, item := range li {
			itemLi := item.([]interface{})
			itemType := itemLi[1].(int)
			isIn := false
			for _, stepTypeItem := range stepTypeLi {
				if itemType == stepTypeItem {
					isIn = true
					break
				}
			}
			if isIn {
				stepLi = append(stepLi, item)
			}
		}
		qb := QuerySupport{}
		for _, item := range stepLi {
			itemLi := item.([]interface{})
			itemType := itemLi[1].(int)
			_, found := qb.FindByMapWithSession(session, "SysStep", map[string]interface{}{
				"A.sysUnitId": unitId,
				"A.type":      itemType,
			})
			if !found {
				id := mongo.GetSequenceNo(db, "sysStepId")
				txnManager.Insert(txnId, "SysStep", map[string]interface{}{
					"_id": id,
					"id":  id,
					"A": map[string]interface{}{
						"id":         id,
						"name":       itemLi[0].(string),
						"type":       itemType,
						"status":     1,
						"sysUserId":  userId,
						"sysUnitId":  unitId,
						"createBy":   userId,
						"createTime": DateUtil{}.GetCurrentYyyyMMddHHmmss(),
						"createUnit": unitId,
					},
				})
			}
		}
	}
	return isStep
}
