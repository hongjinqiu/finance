#!/bin/bash
#coding=utf8

def getBillTypeInitData():
    return {
        'name': 'BillType',
        'items': [{
            '_id': 1,
            'id': 1,
            'A': {
                'id': 1,
                'code': '001',
                'name': '收款单',
            },
        }, {
            '_id': 2,
            'id': 2,
            'A': {
                'id': 2,
                'code': '002',
                'name': '付款单',
            },
        }],
    }

def getBillTypeParameterInitData():
    return {
        'name': 'BillTypeParameter',
        'items': [{
            "A" : {
                "billTypeId" : 1,
                "createUnit" : 0,
                "modifyUnit" : 0,
                "attachCount" : 0,
                "remark" : "",
                "property" : 2,
                "id" : 1,
                "createBy" : 15,
                "createTime" : 20140227165212L,
                "modifyBy" : 0,
                "modifyTime" : 0
            },
            "_id" : 1,
            "id" : 1,
            "pendingTransactions" : [],
        }, {
            "A" : {
                "billTypeId" : 2,
                "createUnit" : 0,
                "modifyUnit" : 0,
                "attachCount" : 0,
                "remark" : "",
                "property" : 2,
                "id" : 2,
                "createBy" : 15,
                "createTime" : 20140227171435L,
                "modifyBy" : 0,
                "modifyTime" : 0
            },
            "_id" : 2,
            "id" : 2,
            "pendingTransactions" : [],
        }],
    }

def getCurrencyTypeInitData():
    return {
        'name': 'CurrencyType',
        'items': [{
            "A" : {
                "remark" : "",
                "createUnit" : 0L,
                "modifyUnit" : 0,
                "id" : 1L,
                "code" : "RMB",
                "roundingWay" : 2L,
                "upDecimals" : 3L,
                "createTime" : 20140410153316L,
                "modifyTime" : 20140429112340L,
                "currencyTypeSign" : "",
                "amtDecimals" : 3L,
                "attachCount" : 0L,
                "name" : "人民币",
                "createBy" : 15L,
                "modifyBy" : 15L,
                "billStatus" : 0L
            },
            "_id" : 1,
            "id" : 1,
            "pendingTransactions" : [ ]
        }],
    }

def getInitDataLi():
    li = []
    li.append(getBillTypeInitData())
    li.append(getBillTypeParameterInitData())
    li.append(getCurrencyTypeInitData())
    return li
