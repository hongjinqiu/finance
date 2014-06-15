package accountinout

import (

)

type AccountInOutParam struct {
	AccountType int	// 1:现金,2:银行
	AccountId int	// 账户ID
	CurrencyTypeId int	// 币别ID
	ExchangeRateShow string	// 汇率
	ExchangeRate string	// 汇率隐藏
	AccountingPeriodYear int	// 年度
	AccountingPeriodMonth int	// 会计期序号
	AmtIncrease string	// 本期增加,传入正数,	收款单为本期增加,利息的正数为本期增加,费用的负数为本期增加
	AmtReduce string	// 本期减少,传入正数,		付款为本期减少,利息负数为本期减少,费用正数为本期减少
	DiffDataType int	// 参考:DiffDataType.go
	CreateBy int
	CreateTime int64
	CreateUnit int
	ModifyBy int
	ModifyUnit int
	ModifyTime int64

	AccountInOutItemParam AccountInOutItemParam	// 日记账明细
}
