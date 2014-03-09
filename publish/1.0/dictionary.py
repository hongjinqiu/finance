#!/bin/bash
#coding=utf8

def getDictionaryLi():
    li = []
    li.append({
        'code': 'D_DECIMALS_TYPE',
        'name': '小数位数',
        'items': [
            {'code': 1,'name': '整数','order': 1},
            {'code': 2,'name': '1位小数','order': 2},
            {'code': 3,'name': '2位小数','order': 3},
            {'code': 4,'name': '3位小数','order': 4},
            {'code': 5,'name': '4位小数','order': 5},
            {'code': 6,'name': '5位小数','order': 6},
            {'code': 7,'name': '6位小数','order': 7},
        ],
    })
    li.append({
        'code': 'D_ROUNDING_WAY_TYPE',
        'name': '舍入方式',
        'items': [
            {'code': 1,'name': '全舍','order': 1},
            {'code': 2,'name': '四舍五入','order': 2},
            {'code': 3,'name': '全入','order': 3},
        ],
    })
    li.append({
        'code': 'D_QFWF',
        'name': '千分位符',
        'items': [
            {'code': 1,'name': '不启用','order': 1},
            {'code': 2,'name': ',','order': 2},
        ],
    })
    li.append({
        'code': 'D_YESNO',
        'name': '是否',
        'items': [
            {'code': 1,'name': '否','order': 1},
            {'code': 2,'name': '是','order': 2},
        ],
    })
    li.append({
        'code': 'D_FIN_BUSI_PROPERTY',
        'name': '业务属性',
        'items': [
            {'code': 1,'name': '银行存款','order': 1},
            {'code': 2,'name': '现金','order': 2},
            {'code': 3,'name': '其它','order': 3},
        ],
    })
    li.append({
        'code': 'D_BILL_TYPE_PARAMETER',
        'name': '单据类型参数类型',
        'items': [
            {'code': 1,'name': '收款单类型参数','order': 1},
            {'code': 2,'name': '付款单类型参数','order': 2},
        ],
    })
    li.append({
        'code': 'D_CURRENT_OBJECT_TYPE',
        'name': '往来对象类型',
        'items': [
            {'code': 1,'name': '个人','order': 1},
            {'code': 2,'name': '公司','order': 2},
        ],
    })
    li.append({
        'code': 'D_CUSTOMER_STATUS',
        'name': '客户状态',
        'items': [
            {'code': 1,'name': '正常','order': 1},
            {'code': 2,'name': '冻结','order': 2},
        ],
    })
    li.append({
        'code': 'D_FIN_ACCOUNT_PROPERTY',
        'name': '账户属性',
        'items': [
            {'code': 1,'name': '基本户','order': 1},
            {'code': 2,'name': '个人账户','order': 2},
            {'code': 3,'name': '保证金户','order': 3},
            {'code': 4,'name': '在途资金','order': 4},
            {'code': 5,'name': '虚拟账户','order': 5},
        ],
    })
    li.append({
        'code': 'D_FIN_LIMITS_CONTROL',
        'name': '赤字控制',
        'items': [
            {'code': 1,'name': '禁止','order': 1},
            {'code': 2,'name': '警告','order': 2},
        ],
    })
    li.append({
        'code': 'D_FIN_BALANCE_ATTR',
        'name': '属性',
        'items': [
            {'code': 1,'name': '银行存款','order': 1},
            {'code': 2,'name': '现金','order': 2},
            {'code': 3,'name': '支票','order': 3},
            {'code': 4,'name': '汇票','order': 4},
            {'code': 5,'name': '会员卡','order': 5},
            {'code': 6,'name': '其他','order': 6},
        ],
    })
    li.append({
        'code': 'D_FIN_PAY_PACT_TYPE',
        'name': '类别',
        'items': [
            {'code': 1,'name': '销售','order': 1},
            {'code': 2,'name': '采购','order': 2},
        ],
    })
    li.append({
        'code': 'D_FIN_PAY_PACT_START_DATE',
        'name': '起算日',
        'items': [
            {'code': 1,'name': '开票日','order': 1},
            {'code': 2,'name': '订单日','order': 2},
            {'code': 3,'name': '收货/交货日','order': 3},
        ],
    })
    li.append({
        'code': 'D_FIN_BALANCE_AMEND',
        'name': '起算日修正',
        'items': [
            {'code': 1,'name': '次月初','order': 1},
            {'code': 2,'name': '起算日','order': 2},
        ],
    })
    li.append({
        'code': 'D_FIN_RECKONING_AMEND',
        'name': '结帐日修正',
        'items': [
            {'code': 1,'name': '加日数','order': 1},
            {'code': 2,'name': '加月数','order': 2},
        ],
    })
    li.append({
        'code': 'D_FIN_PAY_DATE',
        'name': '预计收/付款日',
        'items': [
            {'code': 1,'name': '结算日','order': 1},
            {'code': 2,'name': '结帐日','order': 2},
        ],
    })
    li.append({
        'code': 'D_FIN_OBJECT_TYPE',
        'name': '对象类型',
        'items': [
            {'code': 1,'name': '客户','order': 1},
            {'code': 2,'name': '供应商','order': 2},
            {'code': 3,'name': '人员','order': 3},
            {'code': 4,'name': '其他','order': 4},
        ],
    })
    li.append({
        'code': 'D_FEE_ACCOUNT_TYPE',
        'name': '费用账户类型',
        'items': [
            {'code': 1,'name': '银行存款','order': 1},
            {'code': 2,'name': '现金','order': 2},
        ],
    })
    li.append({
        'code': 'D_ACCOUNT_TYPE',
        'name': '账户类型',
        'items': [
            {'code': 1,'name': '现金账户','order': 1},
            {'code': 2,'name': '银行账户','order': 2},
        ],
    })


    return li
