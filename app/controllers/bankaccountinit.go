package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/accountinout"
	. "com/papersns/error"
	. "com/papersns/model"
	. "com/papersns/model/handler"
	"fmt"
	"strconv"
	"strings"
)

func init() {
}

type BankAccountInitSupport struct {
	ActionSupport
}

func (c BankAccountInitSupport) afterSaveData(sessionId int, dataSource DataSource, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	for _, item := range *diffDateRowLi {
		if item.SrcData != nil && item.DestData != nil { // 修改
			// 旧数据反过账,新数据正过账
			c.logBankAccount(sessionId, dataSource, item.SrcData, BEFORE_UPDATE)
			c.logBankAccount(sessionId, dataSource, *(item.DestData), AFTER_UPDATE)
		} else if item.SrcData == nil && item.DestData != nil { // 新增
			// 新数据正过账
			c.logBankAccount(sessionId, dataSource, *(item.DestData), ADD)
		}
	}
}

func (c BankAccountInitSupport) afterDeleteData(sessionId int, dataSource DataSource, bo *map[string]interface{}) {
	// 反过账
	data := (*bo)["A"].(map[string]interface{})
	c.logBankAccount(sessionId, dataSource, data, DELETE)
}

func (c BankAccountInitSupport) checkLimitsControl(sessionId int, diffDateRowAllLi []DiffDataRow, continueAnyAll string) {
	accountInOutService := AccountInOutService{}
	forbidLi, warnLi := accountInOutService.CheckBankAccountDiffDataRowLimitControl(sessionId, diffDateRowAllLi)

	if len(forbidLi) > 0 {
		panic(BusinessError{
			Code:    LIMIT_CONTROL_FORBID,
			Message: strings.Join(forbidLi, "<br />"),
		})
	}
	if len(warnLi) > 0 && continueAnyAll == "false" {
		panic(BusinessError{
			Code:    LIMIT_CONTROL_WARN,
			Message: strings.Join(warnLi, "<br />"),
		})
	}
}

func (c BankAccountInitSupport) logBankAccount(sessionId int, dataSource DataSource, data map[string]interface{}, diffDataType int) {
	accountId, err := strconv.Atoi(fmt.Sprint(data["accountId"]))
	if err != nil {
		panic(err)
	}
	currencyTypeId, err := strconv.Atoi(fmt.Sprint(data["currencyTypeId"]))
	if err != nil {
		panic(err)
	}
	accountInOutParam := AccountInOutParam{
		AccountId:      accountId,
		CurrencyTypeId: currencyTypeId,
		AmtIncrease:    fmt.Sprint(data["amtEarly"]),
		DiffDataType:   diffDataType,
	}

	accountInOutService := AccountInOutService{}
	accountInOutService.LogBankAccountInOut(sessionId, accountInOutParam)
}

type BankAccountInit struct {
	BaseDataAction
}

func (c BankAccountInit) SaveData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	modelRenderVO := c.saveCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankAccountInit) DeleteData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}

	modelRenderVO := c.deleteDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankAccountInit) EditData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	modelRenderVO := c.editDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankAccountInit) NewData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	modelRenderVO := c.newDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankAccountInit) GetData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	modelRenderVO := c.getDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 复制
 */
func (c BankAccountInit) CopyData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	modelRenderVO := c.copyDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BankAccountInit) GiveUpData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	modelRenderVO := c.giveUpDataCommon()
	return c.renderCommon(modelRenderVO)
}

/**
 * 刷新
 */
func (c BankAccountInit) RefreshData() revel.Result {
	c.actionSupport = BankAccountInitSupport{}
	modelRenderVO := c.refreshDataCommon()
	return c.renderCommon(modelRenderVO)
}

func (c BankAccountInit) LogList() revel.Result {
	result := c.logListCommon()

	format := c.Params.Get("format")
	if strings.ToLower(format) == "json" {
		c.Response.ContentType = "application/json; charset=utf-8"
		return c.RenderJson(result)
	}
	return c.Render()
}
