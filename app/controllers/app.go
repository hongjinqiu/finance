package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"bytes"
	. "com/papersns/common"
	. "com/papersns/component"
	"com/papersns/global"
	. "com/papersns/mongo"
	. "com/papersns/model"
	"compress/gzip"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"
	//	"time"
	"runtime/pprof"
)

var gzipRwlock sync.RWMutex = sync.RWMutex{}
var isRunTxnPeriod bool = false
var periodRwlock sync.RWMutex = sync.RWMutex{}

func getRunTxnPeriod() bool {
	periodRwlock.RLock()
	defer periodRwlock.RUnlock()

	return isRunTxnPeriod
}

func init() {
	revel.TemplateFuncs["inc"] = func(a int, b int) int {
		return a + b
	}
}

type App struct {
	*revel.Controller
}

func (c App) WriteHeapProfile() revel.Result {
	logFile, err := os.Create("/home/hongjinqiu/tmp/heap.out")
	if err != nil {
		panic(err)
	}
	err = pprof.WriteHeapProfile(logFile)
	if err != nil {
		panic(err)
	}
	return c.RenderText("write success")
}

func (c App) ModelTest() revel.Result {
	modelTemplateFactory := ModelTemplateFactory{}
	modelTemplateFactory.GetDataSource("GatheringBill")
	return c.RenderText("modelTest success")
}

func (c App) StartRunTxnPeriod() revel.Result {
	if !getRunTxnPeriod() {
		periodRwlock.Lock()
		defer periodRwlock.Unlock()

		if !isRunTxnPeriod {
			isRunTxnPeriod = true
			TxnPeriodTask{}.RunTxnPeriod()
		}

	}
	return c.RenderText("StartRunTxnPeriod")
}

func (c App) Step() revel.Result {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)

	userId, err := strconv.Atoi(c.Session["userId"])
	if err != nil {
		panic(err)
	}
	sysUser := map[string]interface{}{}
	_, db := global.GetConnection(sessionId)
	err = db.C("SysUser").Find(map[string]interface{}{
		"id": userId,
	}).One(&sysUser)
	sysUserMain := sysUser["A"].(map[string]interface{})

	maxStep := map[string]interface{}{}
	err = db.C("SysStep").Find(map[string]interface{}{
		"A.sysUnitId": sysUserMain["createUnit"],
	}).Sort("-A.type").Limit(1).One(&maxStep)
	if err != nil {
		panic(err)
	}

	currentStep := map[string]interface{}{}
	err = db.C("SysStep").Find(map[string]interface{}{
		"A.sysUnitId": sysUserMain["createUnit"],
		"A.status":    1,
	}).Sort("A.type").Limit(1).One(&currentStep)
	currentId := 0
	maxStepMain := maxStep["A"].(map[string]interface{})
	if err != nil {
		currentId = CommonUtil{}.GetIntFromMap(maxStepMain, "type") + 1
	} else {
		currentStepMain := currentStep["A"].(map[string]interface{})
		currentId = CommonUtil{}.GetIntFromMap(currentStepMain, "type")
	}
	response := map[string]interface{}{
		"maxId":     maxStepMain["type"],
		"currentId": currentId,
	}
	return c.RenderJson(response)
}

func (c App) IncomeTest() revel.Result {
	if c.Params.Query.Get("code") != "" {
		url := c.Request.URL.Path + "?" + c.Request.URL.RawQuery
		sessionId := global.GetSessionId()
		defer global.CloseSession(sessionId)
		defer global.RollbackTxn(sessionId)

		resStruct, userId, isStep := LoginService{}.DealLoginTest(sessionId, url)
		global.CommitTxn(sessionId)

		//		if true {
		//			return c.RenderText("income test")
		//		}
		c.Session["userId"] = fmt.Sprint(userId)

		loginService := LoginService{}
		if isStep {
			qb := QuerySupport{}
			session, db := global.GetConnection(sessionId)
			user := qb.FindByMapWithSessionExact(session, "SysUser", map[string]interface{}{
				"id": userId,
			})
			userMain := user["A"].(map[string]interface{})
			appKey := fmt.Sprint(resStruct["app_key"])
			stepTypeLi := loginService.GetStepTypeLi(appKey)
			stepLi := []interface{}{}
			err := db.C("SysStep").Find(map[string]interface{}{
				"A.sysUnitId": userMain["createUnit"],
				"A.type": map[string]interface{}{
					"$in": stepTypeLi,
				},
			}).Sort("A.type").All(&stepLi)
			if err != nil {
				panic(err)
			}
			c.RenderArgs["result"] = map[string]interface{}{
				"stepLi": stepLi,
			}
			//c.Response.ContentType = "text/html; charset=utf-8"
			return c.RenderTemplate("Step/Step.html")
		} else {
			return c.Redirect("/")
		}
	}

	if c.Session["userId"] == "" {
		if strings.Index(c.Request.Header.Get("HTTP_REFERER"), "taobao") > -1 {
			taobaoPath := revel.Config.StringDefault("TAOBAO_PATH", "")
			return c.Redirect(taobaoPath)
		} else {
			c.Response.ContentType = "text/plain; charset=utf-8"
			return c.RenderText("会话过期，请您从淘宝重新登录应用！")
		}
	}
	// 取得菜单等数据,
	userId, err := strconv.Atoi(c.Session["userId"])
	if err != nil {
		panic(err)
	}
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	qb := QuerySupport{}
	session, _ := global.GetConnection(sessionId)
	user := qb.FindByMapWithSessionExact(session, "SysUser", map[string]interface{}{
		"id": userId,
	})

	// 获取数据,
	result := map[string]interface{}{
		"user":            user,
		"unit":            c.getSysUnit(sessionId, user),
		"menuLi":          c.getMenuLi(sessionId),
		"gatheringBillLi": c.getGatheringBillLi(sessionId, user),
		"payBillLi":       c.getPayBillLi(sessionId, user),
	}
	return c.Render(result)
}

func (c App) StepTest() revel.Result {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)

	_, db := global.GetConnection(sessionId)
	stepLi := []interface{}{}
	err := db.C("SysStep").Find(map[string]interface{}{
		"A.sysUnitId": 1,
		"A.status":    1,
	}).Sort("A.type").All(&stepLi)
	if err != nil {
		panic(err)
	}
	c.RenderArgs["result"] = map[string]interface{}{
		"stepLi": stepLi,
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.RenderTemplate("Step/Step.html")
}

/**
http://uhz001889.chinaw3.com/?myEnt=approval&code=BiWbZfjYoCIp4oRzo7rgRpUv243721&state=
*/
func (c App) Index() revel.Result {
	if c.Params.Query.Get("code") != "" {
		url := c.Request.URL.Path + "?" + c.Request.URL.RawQuery
		sessionId := global.GetSessionId()
		defer global.CloseSession(sessionId)
		defer global.RollbackTxn(sessionId)

		resStruct, userId, isStep := LoginService{}.DealLogin(sessionId, url)
		global.CommitTxn(sessionId)
		c.Session["userId"] = fmt.Sprint(userId)

		loginService := LoginService{}
		if isStep {
			qb := QuerySupport{}
			session, db := global.GetConnection(sessionId)
			user := qb.FindByMapWithSessionExact(session, "SysUser", map[string]interface{}{
				"id": userId,
			})
			userMain := user["A"].(map[string]interface{})
			appKey := fmt.Sprint(resStruct["app_key"])
			stepTypeLi := loginService.GetStepTypeLi(appKey)
			stepLi := []interface{}{}
			err := db.C("SysStep").Find(map[string]interface{}{
				"A.sysUnitId": userMain["createUnit"],
				"A.type": map[string]interface{}{
					"$in": stepTypeLi,
				},
			}).Sort("A.type").All(&stepLi)
			if err != nil {
				panic(err)
			}
			c.RenderArgs["result"] = map[string]interface{}{
				"stepLi": stepLi,
			}
			//c.Response.ContentType = "text/html; charset=utf-8"
			return c.RenderTemplate("Step/Step.html")
		} else {
			return c.Redirect("/")
		}
	}

	if c.Session["userId"] == "" {
		if strings.Index(c.Request.Header.Get("HTTP_REFERER"), "taobao") > -1 {
			taobaoPath := revel.Config.StringDefault("TAOBAO_PATH", "")
			return c.Redirect(taobaoPath)
		} else {
			c.Response.ContentType = "text/plain; charset=utf-8"
			return c.RenderText("会话过期，请您从淘宝重新登录应用！")
		}
	}
	// 取得菜单等数据,
	userId, err := strconv.Atoi(c.Session["userId"])
	if err != nil {
		panic(err)
	}
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	qb := QuerySupport{}
	session, _ := global.GetConnection(sessionId)
	user := qb.FindByMapWithSessionExact(session, "SysUser", map[string]interface{}{
		"id": userId,
	})

	// 获取数据,
	result := map[string]interface{}{
		"user":            user,
		"unit":            c.getSysUnit(sessionId, user),
		"menuLi":          c.getMenuLi(sessionId),
		"gatheringBillLi": c.getGatheringBillLi(sessionId, user),
		"payBillLi":       c.getPayBillLi(sessionId, user),
	}
	return c.Render(result)
}

func (c App) getSysUnit(sessionId int, user map[string]interface{}) map[string]interface{} {
	userMain := user["A"].(map[string]interface{})
	session, _ := global.GetConnection(sessionId)
	qb := QuerySupport{}
	return qb.FindByMapWithSessionExact(session, "SysUnit", map[string]interface{}{
		"id": userMain["createUnit"],
	})
}

func (c App) getGatheringBillLi(sessionId int, user map[string]interface{}) []map[string]interface{} {
	userMain := user["A"].(map[string]interface{})
	_, db := global.GetConnection(sessionId)
	li := []map[string]interface{}{}
	err := db.C("GatheringBill").Find(map[string]interface{}{
		"A.createUnit": userMain["createUnit"],
	}).Limit(6).Sort("-A.billDate").All(&li)
	if err != nil {
		panic(err)
	}
	result := []map[string]interface{}{}
	cashAccountIdLi := []int{}
	bankAccountIdLi := []int{}

	customerChamberlainId := []int{}
	providerChamberlainId := []int{}
	sysUserChamberlainId := []int{}

	commonUtil := CommonUtil{}
	dateUtil := DateUtil{}
	for _, item := range li {
		master := item["A"].(map[string]interface{})
		if fmt.Sprint(master["property"]) == "1" { // 银行
			bankAccountIdLi = append(bankAccountIdLi, commonUtil.GetIntFromMap(master, "accountId"))
		} else if fmt.Sprint(master["property"]) == "2" { // 现金
			cashAccountIdLi = append(cashAccountIdLi, commonUtil.GetIntFromMap(master, "accountId"))
		}
		if fmt.Sprint(master["chamberlainType"]) == "1" { // customer
			customerChamberlainId = append(customerChamberlainId, commonUtil.GetIntFromMap(master, "chamberlainId"))
		} else if fmt.Sprint(master["chamberlainType"]) == "2" { // provider
			providerChamberlainId = append(providerChamberlainId, commonUtil.GetIntFromMap(master, "chamberlainId"))
		} else if fmt.Sprint(master["chamberlainType"]) == "3" { // sysUser
			sysUserChamberlainId = append(sysUserChamberlainId, commonUtil.GetIntFromMap(master, "chamberlainId"))
		}
		resultItem := map[string]interface{}{
			"id":              master["id"],
			"billDate":        dateUtil.ConvertDate2String(fmt.Sprint(master["billDate"]), "20060102", "2006-01-02"),
			"billNo":          master["billNo"],
			"property":        master["property"],
			"accountId":       master["accountId"],
			"account":         "",
			"chamberlainType": master["chamberlainType"],
			"chamberlainId":   master["chamberlainId"],
			"chamberlain":     "",
			"amtGathering":    commonUtil.TrimZero(master["amtGathering"].(string)),
		}
		result = append(result, resultItem)
	}
	c.mergeCashAccount(sessionId, cashAccountIdLi, &result)
	c.mergeBankAccount(sessionId, bankAccountIdLi, &result)

	c.mergeGatheringCustomer(sessionId, customerChamberlainId, &result)
	c.mergeGatheringProvider(sessionId, providerChamberlainId, &result)
	c.mergeGatheringSysUser(sessionId, sysUserChamberlainId, &result)

	return result
}

func (c App) getPayBillLi(sessionId int, user map[string]interface{}) []map[string]interface{} {
	userMain := user["A"].(map[string]interface{})
	_, db := global.GetConnection(sessionId)
	li := []map[string]interface{}{}
	err := db.C("PayBill").Find(map[string]interface{}{
		"A.createUnit": userMain["createUnit"],
	}).Limit(6).Sort("-A.billDate").All(&li)
	if err != nil {
		panic(err)
	}
	result := []map[string]interface{}{}
	cashAccountIdLi := []int{}
	bankAccountIdLi := []int{}

	customerChamberlainId := []int{}
	providerChamberlainId := []int{}
	sysUserChamberlainId := []int{}

	commonUtil := CommonUtil{}
	dateUtil := DateUtil{}
	for _, item := range li {
		master := item["A"].(map[string]interface{})
		if fmt.Sprint(master["property"]) == "1" { // 银行
			bankAccountIdLi = append(bankAccountIdLi, commonUtil.GetIntFromMap(master, "accountId"))
		} else if fmt.Sprint(master["property"]) == "2" { // 现金
			cashAccountIdLi = append(cashAccountIdLi, commonUtil.GetIntFromMap(master, "accountId"))
		}
		if fmt.Sprint(master["payerType"]) == "1" { // customer
			customerChamberlainId = append(customerChamberlainId, commonUtil.GetIntFromMap(master, "payerId"))
		} else if fmt.Sprint(master["payerType"]) == "2" { // provider
			providerChamberlainId = append(providerChamberlainId, commonUtil.GetIntFromMap(master, "payerId"))
		} else if fmt.Sprint(master["payerType"]) == "3" { // sysUser
			sysUserChamberlainId = append(sysUserChamberlainId, commonUtil.GetIntFromMap(master, "payerId"))
		}
		resultItem := map[string]interface{}{
			"id":        master["id"],
			"billDate":  dateUtil.ConvertDate2String(fmt.Sprint(master["billDate"]), "20060102", "2006-01-02"),
			"billNo":    master["billNo"],
			"property":  master["property"],
			"accountId": master["accountId"],
			"account":   "",
			"payerType": master["payerType"],
			"payerId":   master["payerId"],
			"payer":     "",
			"amtPay":    commonUtil.TrimZero(master["amtPay"].(string)),
		}
		result = append(result, resultItem)
	}
	c.mergeCashAccount(sessionId, cashAccountIdLi, &result)
	c.mergeBankAccount(sessionId, bankAccountIdLi, &result)

	c.mergePayCustomer(sessionId, customerChamberlainId, &result)
	c.mergePayProvider(sessionId, providerChamberlainId, &result)
	c.mergePaySysUser(sessionId, sysUserChamberlainId, &result)

	return result
}

func (c App) mergeCashAccount(sessionId int, cashAccountIdLi []int, result *[]map[string]interface{}) {
	_, db := global.GetConnection(sessionId)
	cashAccountLi := []map[string]interface{}{}
	err := db.C("CashAccount").Find(map[string]interface{}{
		"id": map[string]interface{}{
			"$in": cashAccountIdLi,
		},
	}).All(&cashAccountLi)
	if err != nil {
		panic(err)
	}
	for i, item := range *result {
		(*result)[i] = item

		for _, cashAccount := range cashAccountLi {
			cashAccountMaster := cashAccount["A"].(map[string]interface{})
			if fmt.Sprint(item["property"]) == "2" && fmt.Sprint(item["accountId"]) == fmt.Sprint(cashAccountMaster["id"]) {
				item["account"] = fmt.Sprint(cashAccountMaster["code"]) + "," + fmt.Sprint(cashAccountMaster["name"])
				break
			}
		}
	}
}

func (c App) mergeBankAccount(sessionId int, bankAccountIdLi []int, result *[]map[string]interface{}) {
	_, db := global.GetConnection(sessionId)
	bankAccountLi := []map[string]interface{}{}
	err := db.C("BankAccount").Find(map[string]interface{}{
		"id": map[string]interface{}{
			"$in": bankAccountIdLi,
		},
	}).All(&bankAccountLi)
	if err != nil {
		panic(err)
	}
	for i, item := range *result {
		(*result)[i] = item

		for _, bankAccount := range bankAccountLi {
			bankAccountMaster := bankAccount["A"].(map[string]interface{})
			if fmt.Sprint(item["property"]) == "1" && fmt.Sprint(item["accountId"]) == fmt.Sprint(bankAccountMaster["id"]) {
				item["account"] = fmt.Sprint(bankAccountMaster["code"]) + "," + fmt.Sprint(bankAccountMaster["name"])
				break
			}
		}
	}
}

func (c App) mergeGatheringCustomer(sessionId int, customerIdLi []int, result *[]map[string]interface{}) {
	_, db := global.GetConnection(sessionId)
	customerLi := []map[string]interface{}{}
	err := db.C("Customer").Find(map[string]interface{}{
		"id": map[string]interface{}{
			"$in": customerIdLi,
		},
	}).All(&customerLi)
	if err != nil {
		panic(err)
	}
	for i, item := range *result {
		(*result)[i] = item

		for _, customer := range customerLi {
			customerMaster := customer["A"].(map[string]interface{})
			if fmt.Sprint(item["chamberlainType"]) == "1" && fmt.Sprint(item["chamberlainId"]) == fmt.Sprint(customerMaster["id"]) {
				item["chamberlain"] = fmt.Sprint(customerMaster["code"]) + "," + fmt.Sprint(customerMaster["name"])
				break
			}
		}
	}
}

func (c App) mergeGatheringProvider(sessionId int, providerIdLi []int, result *[]map[string]interface{}) {
	_, db := global.GetConnection(sessionId)
	providerLi := []map[string]interface{}{}
	err := db.C("Provider").Find(map[string]interface{}{
		"id": map[string]interface{}{
			"$in": providerIdLi,
		},
	}).All(&providerLi)
	if err != nil {
		panic(err)
	}
	for i, item := range *result {
		(*result)[i] = item

		for _, provider := range providerLi {
			providerMaster := provider["A"].(map[string]interface{})
			if fmt.Sprint(item["chamberlainType"]) == "2" && fmt.Sprint(item["chamberlainId"]) == fmt.Sprint(providerMaster["id"]) {
				item["chamberlain"] = fmt.Sprint(providerMaster["code"]) + "," + fmt.Sprint(providerMaster["name"])
				break
			}
		}
	}
}

func (c App) mergeGatheringSysUser(sessionId int, sysUserIdLi []int, result *[]map[string]interface{}) {
	_, db := global.GetConnection(sessionId)
	sysUserLi := []map[string]interface{}{}
	err := db.C("SysUser").Find(map[string]interface{}{
		"id": map[string]interface{}{
			"$in": sysUserIdLi,
		},
	}).All(&sysUserLi)
	if err != nil {
		panic(err)
	}
	for i, item := range *result {
		(*result)[i] = item

		for _, sysUser := range sysUserLi {
			sysUserMaster := sysUser["A"].(map[string]interface{})
			if fmt.Sprint(item["chamberlainType"]) == "3" && fmt.Sprint(item["chamberlainId"]) == fmt.Sprint(sysUserMaster["id"]) {
				item["chamberlain"] = fmt.Sprint(sysUserMaster["code"]) + "," + fmt.Sprint(sysUserMaster["name"])
				break
			}
		}
	}
}

func (c App) mergePayCustomer(sessionId int, customerIdLi []int, result *[]map[string]interface{}) {
	_, db := global.GetConnection(sessionId)
	customerLi := []map[string]interface{}{}
	err := db.C("Customer").Find(map[string]interface{}{
		"id": map[string]interface{}{
			"$in": customerIdLi,
		},
	}).All(&customerLi)
	if err != nil {
		panic(err)
	}
	for i, item := range *result {
		(*result)[i] = item

		for _, customer := range customerLi {
			customerMaster := customer["A"].(map[string]interface{})
			if fmt.Sprint(item["payerType"]) == "1" && fmt.Sprint(item["payerId"]) == fmt.Sprint(customerMaster["id"]) {
				item["payer"] = fmt.Sprint(customerMaster["code"]) + "," + fmt.Sprint(customerMaster["name"])
				break
			}
		}
	}
}

func (c App) mergePayProvider(sessionId int, providerIdLi []int, result *[]map[string]interface{}) {
	_, db := global.GetConnection(sessionId)
	providerLi := []map[string]interface{}{}
	err := db.C("Provider").Find(map[string]interface{}{
		"id": map[string]interface{}{
			"$in": providerIdLi,
		},
	}).All(&providerLi)
	if err != nil {
		panic(err)
	}
	for i, item := range *result {
		(*result)[i] = item

		for _, provider := range providerLi {
			providerMaster := provider["A"].(map[string]interface{})
			if fmt.Sprint(item["payerType"]) == "2" && fmt.Sprint(item["payerId"]) == fmt.Sprint(providerMaster["id"]) {
				item["payer"] = fmt.Sprint(providerMaster["code"]) + "," + fmt.Sprint(providerMaster["name"])
				break
			}
		}
	}
}

func (c App) mergePaySysUser(sessionId int, sysUserIdLi []int, result *[]map[string]interface{}) {
	_, db := global.GetConnection(sessionId)
	sysUserLi := []map[string]interface{}{}
	err := db.C("SysUser").Find(map[string]interface{}{
		"id": map[string]interface{}{
			"$in": sysUserIdLi,
		},
	}).All(&sysUserLi)
	if err != nil {
		panic(err)
	}
	for i, item := range *result {
		(*result)[i] = item

		for _, sysUser := range sysUserLi {
			sysUserMaster := sysUser["A"].(map[string]interface{})
			if fmt.Sprint(item["payerType"]) == "3" && fmt.Sprint(item["payerId"]) == fmt.Sprint(sysUserMaster["id"]) {
				item["payer"] = fmt.Sprint(sysUserMaster["code"]) + "," + fmt.Sprint(sysUserMaster["name"])
				break
			}
		}
	}
}

func (c App) getMenuLi(sessionId int) []interface{} {
	line1 := []map[string]interface{}{
		map[string]interface{}{"name": "收款单", "image": "40174.gif"},
		map[string]interface{}{"name": "付款单", "image": "40171.gif"},
		map[string]interface{}{"name": "资金汇总表", "image": "40153.gif"},
		map[string]interface{}{"name": "现金账户初始化", "image": "40170.gif"},
	}
	line2 := []map[string]interface{}{
		map[string]interface{}{"name": "银行账户初始化", "image": "40138.gif"},
		map[string]interface{}{"name": "系统参数", "image": "6.gif"},
		map[string]interface{}{"name": "单据类型参数", "image": "11.gif"},
		map[string]interface{}{"name": "会计期", "image": "69.gif"},
	}
	line3 := []map[string]interface{}{
		map[string]interface{}{"name": "银行账户", "image": "53.gif"},
		map[string]interface{}{"name": "现金账户", "image": "42.gif"},
		map[string]interface{}{"name": "客户", "image": "7.gif"},
		map[string]interface{}{"name": "银行资料", "image": "68.gif"},
	}

	nameLi := []string{}
	for _, item := range line1 {
		nameLi = append(nameLi, item["name"].(string))
	}
	for _, item := range line2 {
		nameLi = append(nameLi, item["name"].(string))
	}
	for _, item := range line3 {
		nameLi = append(nameLi, item["name"].(string))
	}

	_, db := global.GetConnection(sessionId)
	menuLi := []map[string]interface{}{}
	err := db.C("Menu").Find(map[string]interface{}{
		"name": map[string]interface{}{
			"$in": nameLi,
		},
		"isLeaf": 1,
	}).All(&menuLi)
	if err != nil {
		panic(err)
	}

	for i, item := range line1 {
		line1[i] = item
		for _, menu := range menuLi {
			if item["name"].(string) == fmt.Sprint(menu["name"]) {
				item["url"] = template.JSStr(fmt.Sprint(menu["url"]))
				break
			}
		}
	}
	for i, item := range line2 {
		line2[i] = item
		for _, menu := range menuLi {
			if item["name"].(string) == fmt.Sprint(menu["name"]) {
				item["url"] = template.JSStr(fmt.Sprint(menu["url"]))
				break
			}
		}
	}
	for i, item := range line3 {
		line3[i] = item
		for _, menu := range menuLi {
			if item["name"].(string) == fmt.Sprint(menu["name"]) {
				item["url"] = template.JSStr(fmt.Sprint(menu["url"]))
				break
			}
		}
	}
	return []interface{}{
		line1,
		line2,
		line3,
	}
}

func (c App) MenuList() revel.Result {
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)

	_, db := global.GetConnection(sessionId)
	menuResultLi := []map[string]interface{}{}
	err := db.C("Menu").Find(nil).Sort("level").All(&menuResultLi)
	if err != nil {
		panic(err)
	}

	menuLi := []map[string]interface{}{}
	for _, item := range menuResultLi {
		level := fmt.Sprint(item["level"])
		if len(level) == 3 {
			menuLi = append(menuLi, item)

			subMenuLi := []map[string]interface{}{}
			for _, subItem := range menuResultLi {
				subLevel := fmt.Sprint(subItem["level"])
				if len(subLevel) == 6 && subLevel[0:3] == level {
					subMenuLi = append(subMenuLi, subItem)
				}
			}
			item["subMenuLi"] = subMenuLi
		}
	}

	result := map[string]interface{}{
		"menuLi": menuLi,
	}
	return c.Render(result)
}

func (c App) Instructions() revel.Result {
	return c.Render()
}

func (c App) Logout() revel.Result {
	c.Session["userId"] = ""
	c.Session["adminUserId"] = ""
	return c.Redirect(revel.Config.StringDefault("OAUTH_LOGOUT_PATH", ""))
}

/*
Cache-Control:max-age=315360000
Connection:keep-alive
Date:Thu, 19 Sep 2013 08:25:26 GMT
Expires:Sat, 05 Sep 2026 00:00:00 GMT
Server:ATS/3.2.0
Vary:Accept-Encoding
Via:http/1.1 l4.ycs.swp.yahoo.com (ApacheTrafficServer/3.2.0)

js
Connection:Keep-Alive
Content-Encoding:gzip
Content-Type:text/javascript;charset=UTF-8
Date:Thu, 19 Sep 2013 08:27:48 GMT
Server:Apache
Transfer-Encoding:chunked
*/
func (c App) Combo() revel.Result {
	//	url := c.Request.URL.Path + "?" + c.Request.URL.RawQuery
	//	start := time.Now()

	nameConcat := c.getFileNameConcatFromQuery()

	acceptEncoding := c.Request.Header.Get("Accept-Encoding")
	if strings.Index(acceptEncoding, "gzip") > -1 {
		text := ""
		if revel.Config.StringDefault("mode.dev", "true") == "true" {
			content := c.getComboFileContent()
			data := bytes.Buffer{}
			w := gzip.NewWriter(&data)
			w.Write([]byte(content))
			w.Close()
			text = data.String()
		} else {
			if c.isFileExist(nameConcat) {
				text = string(c.getGzipContent(nameConcat))
			} else {
				content := c.getComboFileContent()
				text = string(c.gzipAndSave(nameConcat, content))
			}
		}

		c.Response.Status = http.StatusOK
		if strings.Index(c.Params.Query.Encode(), ".css") <= -1 {
			c.Response.ContentType = "text/javascript;charset=UTF-8"
		} else {
			c.Response.ContentType = "text/css;charset=UTF-8"
		}
		c.Response.Out.Header().Set("Content-Encoding", "gzip")
		//		end := time.Now()
		//		println("^^^^^^^^^^^^^^^^url is:" + url + " time spend is:" + fmt.Sprint((end.UnixNano() - start.UnixNano())))
		return c.RenderText(text)
	}

	content := c.getComboFileContent()
	c.Response.Status = http.StatusOK
	if strings.Index(c.Params.Query.Encode(), ".css") <= -1 {
		c.Response.ContentType = "text/javascript;charset=UTF-8"
	} else {
		c.Response.ContentType = "text/css;charset=UTF-8"
	}
	return c.RenderText(content)
}

type StringArraySort struct {
	objLi []string
}

func (o StringArraySort) Len() int {
	return len(o.objLi)
}

func (o StringArraySort) Less(i, j int) bool {
	return o.objLi[i] <= o.objLi[j]
}

func (o StringArraySort) Swap(i, j int) {
	o.objLi[i], o.objLi[j] = o.objLi[j], o.objLi[i]
}

func (c App) getFileNameConcatFromQuery() string {
	queryLi := []string{}
	name := ""
	for k := range c.Params.Query {
		//		name += k
		queryLi = append(queryLi, k)
	}
	stringArraySort := StringArraySort{queryLi}
	sort.Sort(stringArraySort)
	name = strings.Join(stringArraySort.objLi, "")
	return name
}

func (c App) getComboFileContent() string {
	jsPath := revel.Config.StringDefault("JS_PATH", "")
	content := ""
	commonUtil := CommonUtil{}
	for k := range c.Params.Query {
		if !commonUtil.IsNumber(k) && k != "" {
			file, err := os.Open(path.Join(jsPath, k))
			defer file.Close()
			if err != nil {
				panic(err)
			}

			data, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			content += string(data) + "\n"
		}
	}
	return content
}

func (c App) ComboView() revel.Result {
	//	url := c.Request.URL.Path + "?" + c.Request.URL.RawQuery
	//	start := time.Now()

	nameConcat := c.getFileNameConcatFromQuery()

	acceptEncoding := c.Request.Header.Get("Accept-Encoding")
	if strings.Index(acceptEncoding, "gzip") > -1 {
		text := ""
		if revel.Config.StringDefault("mode.dev", "true") == "true" {
			content := c.getComboViewFileContent()
			data := bytes.Buffer{}
			w := gzip.NewWriter(&data)
			w.Write([]byte(content))
			w.Close()
			text = data.String()
		} else {
			if c.isFileExist(nameConcat) {
				text = string(c.getGzipContent(nameConcat))
			} else {
				content := c.getComboViewFileContent()
				text = string(c.gzipAndSave(nameConcat, content))
			}
		}

		c.Response.Status = http.StatusOK
		if strings.Index(c.Params.Query.Encode(), ".css") <= -1 {
			c.Response.ContentType = "text/javascript;charset=UTF-8"
		} else {
			c.Response.ContentType = "text/css;charset=UTF-8"
		}
		//		end := time.Now()
		//		println("^^^^^^^^^^^^^^^^ comboview url is:" + url + " time spend is:" + fmt.Sprint((end.UnixNano() - start.UnixNano())))
		c.Response.Out.Header().Set("Content-Encoding", "gzip")
		return c.RenderText(text)
	}

	content := c.getComboViewFileContent()
	c.Response.Status = http.StatusOK
	if strings.Index(c.Params.Query.Encode(), ".css") <= -1 {
		c.Response.ContentType = "text/javascript;charset=UTF-8"
	} else {
		c.Response.ContentType = "text/css;charset=UTF-8"
	}
	return c.RenderText(content)
}

func (c App) getComboViewFileContent() string {
	jsPath := revel.Config.StringDefault("COMBO_VIEW_PATH", "")
	content := ""
	commonUtil := CommonUtil{}
	for k := range c.Params.Query {
		if !commonUtil.IsNumber(k) && k != "" {
			if strings.Index(k, ".js") == -1 && strings.Index(k, ".css") == -1 {
				panic("fileName is:" + k + ", expect ends with .js or .css")
			}
			isFileExist := false
			for _, filePath := range strings.Split(jsPath, ":") {
				if _, err := os.Stat(path.Join(filePath, k)); err != nil {
					if os.IsNotExist(err) {
						continue
					}
				}
				isFileExist = true
				file, err := os.Open(path.Join(filePath, k))
				defer file.Close()
				if err != nil {
					panic(err)
				}

				data, err := ioutil.ReadAll(file)
				if err != nil {
					panic(err)
				}
				content += string(data) + "\n"
				break
			}
			if !isFileExist {
				panic(k + " is not exists")
			}
		}
	}
	return content
}

func (c App) FormJS() revel.Result {
	//	start := time.Now()

	acceptEncoding := c.Request.Header.Get("Accept-Encoding")
	if strings.Index(acceptEncoding, "gzip") > -1 {
		text := ""
		if revel.Config.StringDefault("mode.dev", "true") == "true" {
			content := c.getFormJsContent()
			data := bytes.Buffer{}
			w := gzip.NewWriter(&data)
			w.Write([]byte(content))
			w.Close()
			text = data.String()
		} else {
			formJsNameLi := c.getFormJsLi()
			nameConcat := strings.Join(formJsNameLi, "")
			if c.isFileExist(nameConcat) {
				text = string(c.getGzipContent(nameConcat))
			} else {
				content := c.getFormJsContent()
				text = string(c.gzipAndSave(nameConcat, content))
			}
		}

		c.Response.Status = http.StatusOK
		c.Response.ContentType = "text/javascript;charset=UTF-8"
		c.Response.Out.Header().Set("Content-Encoding", "gzip")

		//		end := time.Now()
		//		println("^^^^^^^^^^^^^^^^ formjs url time spend is:" + fmt.Sprint((end.UnixNano() - start.UnixNano())))
		return c.RenderText(text)
	}

	content := c.getFormJsContent()
	c.Response.Status = http.StatusOK
	c.Response.ContentType = "text/javascript;charset=UTF-8"
	return c.RenderText(content)
}

func (c App) getFormJsContent() string {
	jsPath := revel.Config.StringDefault("COMBO_VIEW_PATH", "")
	content := ""
	formJsLi := c.getFormJsLi()
	// 加入日期标记,gzip到目标文件时,有用,
	//	commonUtil := CommonUtil{}
	//	for k := range c.Params.Query {
	//		if commonUtil.IsNumber(k) && k != "" {
	//
	//		}
	//	}
	for _, k := range formJsLi {
		if strings.Index(k, ".js") == -1 && strings.Index(k, ".css") == -1 {
			panic("fileName is:" + k + ", expect ends with .js or .css")
		}
		isFileExist := false
		for _, filePath := range strings.Split(jsPath, ":") {
			if _, err := os.Stat(path.Join(filePath, k)); err != nil {
				if os.IsNotExist(err) {
					continue
				}
			}
			isFileExist = true
			file, err := os.Open(path.Join(filePath, k))
			defer file.Close()
			if err != nil {
				panic(err)
			}

			data, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			content += string(data) + "\n"
			break
		}
		if !isFileExist {
			panic(k + " is not exists")
		}
	}
	prefix := "YUI.add('papersns-form', function(Y) {\n"
	suffix := "}, '1.1.0' ,{requires:['node', 'widget-base', 'widget-htmlparser', 'io-form', 'widget-parent', 'widget-child', 'base-build', 'substitute', 'io-upload-iframe', 'collection', 'overlay', 'calendar', 'datatype-date']});\n"
	return prefix + content + suffix
}

func (c App) getFormJsLi() []string {
	formJsLi := []string{"js/rootform/r-form-field.js", "js/rootform/r-text-field.js", "js/rootform/r-hidden-field.js", "js/rootform/r-checkbox-field.js", "js/rootform/r-radio-field.js", "js/rootform/r-choice-field.js", "js/rootform/r-select-field.js", "js/rootform/r-trigger-field.js", "js/rootform/r-number-field.js", "js/rootform/r-display-field.js", "js/rootform/r-textarea-field.js", "js/rootform/r-date-field.js"}
	lFormJsLi := []string{"js/listform/lformcommon.js", "js/listform/l-form-field.js", "js/listform/l-text-field.js", "js/listform/l-hidden-field.js", "js/listform/l-checkbox-field.js", "js/listform/l-radio-field.js", "js/listform/l-choice-field.js", "js/listform/l-select-field.js", "js/listform/l-trigger-field.js", "js/listform/l-number-field.js", "js/listform/l-display-field.js", "js/listform/l-textarea-field.js", "js/listform/l-date-field.js"}
	pFormJsLi := []string{"js/form/p-form-field.js", "js/form/p-text-field.js", "js/form/p-hidden-field.js", "js/form/p-checkbox-field.js", "js/form/p-radio-field.js", "js/form/p-choice-field.js", "js/form/p-select-field.js", "js/form/p-trigger-field.js", "js/form/p-number-field.js", "js/form/p-display-field.js", "js/form/p-textarea-field.js", "js/form/p-date-field.js"}
	for _, k := range pFormJsLi {
		formJsLi = append(formJsLi, k)
	}
	for _, k := range lFormJsLi {
		formJsLi = append(formJsLi, k)
	}
	return formJsLi
}

func (c App) isFileExist(name string) bool {
	h := md5.New()
	io.WriteString(h, name)
	gzipFileName := fmt.Sprintf("%x", h.Sum(nil))
	gzipPath := revel.Config.StringDefault("GZIP_PATH", "")
	if _, err := os.Stat(path.Join(gzipPath, gzipFileName)); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

func (c App) getGzipContent(name string) []byte {
	h := md5.New()
	io.WriteString(h, name)
	gzipFileName := fmt.Sprintf("%x", h.Sum(nil))
	gzipPath := revel.Config.StringDefault("GZIP_PATH", "")
	filePath := path.Join(gzipPath, gzipFileName)

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (c App) gzipAndSave(name string, content string) []byte {
	gzipRwlock.Lock()
	defer gzipRwlock.Unlock()

	h := md5.New()
	io.WriteString(h, name)
	gzipFileName := fmt.Sprintf("%x", h.Sum(nil))
	gzipPath := revel.Config.StringDefault("GZIP_PATH", "")
	filePath := path.Join(gzipPath, gzipFileName)

	if !c.isFileExist(name) {
		data := bytes.Buffer{}
		w := gzip.NewWriter(&data)
		w.Write([]byte(content))
		w.Close()

		bytes := data.Bytes()
		err := ioutil.WriteFile(filePath, bytes, os.ModeDevice|0666)
		if err != nil {
			panic(err)
		}
		return bytes
	} else {
		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		return bytes
	}
}
