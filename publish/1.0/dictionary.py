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
        'name': '帐户属性',
        'items': [
            {'code': 1,'name': '基本户','order': 1},
            {'code': 2,'name': '个人帐户','order': 2},
            {'code': 3,'name': '保证金户','order': 3},
            {'code': 4,'name': '在途资金','order': 4},
            {'code': 5,'name': '虚拟帐户','order': 5},
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
    return li
