package accountinout

import (

)

type AccountInOutItemParam struct {
	AccountType int	// 1:现金,2:银行
	AccountId int	// 账户ID
	CurrencyTypeId int	// 币别ID
	ExchangeRateShow string	// 汇率
	ExchangeRate string	// 汇率隐藏
	AmtIncrease string // 收款金额,传正数,
	AmtReduce string // 付款金额,传正数
	BillTypeId int // 单据类型ID
	BillDataSourceName string // 单据数据源模型ID
	BillCollectionName string // 单据mongoDB collection 名称
	BillDetailName string // 单据数据源模型分录名
	BillId int	// 单据主键ID
	BillDetailId int // 单据分录主键ID
	BillNo string // 单据编号
	BillDate int // 单据日期
	BalanceDate int // 结算日期
	BalanceTypeId int // 结算方式
	BalanceNo string // 结算号
	ChamberlainType int // 收款对象类型
	ChamberlainId int // 收款对象
	CreateBy int
	CreateTime int64
	CreateUnit int
	ModifyBy int
	ModifyUnit int
	ModifyTime int64
//	AccountInOutId int // 月档ID
}

func (o AccountInOutItemParam) ToMap() map[string]interface{} {
	result := map[string]interface{}{}

	result["accountType"] = o.AccountType
	result["accountId"] = o.AccountId
	result["currencyTypeId"] = o.CurrencyTypeId
	result["exchangeRateShow"] = o.ExchangeRateShow
	result["exchangeRate"] = o.ExchangeRate
	result["amtIncrease"] = o.AmtIncrease
	result["amtReduce"] = o.AmtReduce
	result["billTypeId"] = o.BillTypeId
	result["billDataSourceName"] = o.BillDataSourceName
	result["billCollectionName"] = o.BillCollectionName
	result["billDetailName"] = o.BillDetailName
	result["billId"] = o.BillId
	result["billDetailId"] = o.BillDetailId
	result["billNo"] = o.BillNo
	result["billDate"] = o.BillDate
	result["balanceDate"] = o.BalanceDate
	result["balanceTypeId"] = o.BalanceTypeId
	result["balanceNo"] = o.BalanceNo
	result["chamberlainType"] = o.ChamberlainType
	result["chamberlainId"] = o.ChamberlainId
	result["createBy"] = o.CreateBy
	result["createTime"] = o.CreateTime
	result["createUnit"] = o.CreateUnit
	result["modifyBy"] = o.ModifyBy
	result["modifyUnit"] = o.ModifyUnit
	result["modifyTime"] = o.ModifyTime
//	result["accountInOutId"] = o.AccountInOutId

	return result
}

