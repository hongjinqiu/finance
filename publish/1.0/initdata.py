#!/bin/bash
#coding=utf8

def getLastSessionDataInitData():
    return {
        'name': 'LastSessionData',
        'items': [{
            "_id": 1,
            "id": 1,
            "A": {
                "id" : 1,
                "resStruct" : {
                    "top_session" : "40517371e526ed5386ba9dcc0f8886504e8261b17597bmzr6J2m82691",
                    "url" : "http://127.0.0.1:8000/",
                    "agreement" : "true",
                    "app_key" : "12130139",
                    "agreementsign" : "12130139-22287347-346B2278422F673ECB579A9E682E0621",
                    "top_sign" : "5tQpWMO3Ws+eE+ygdPbgew==",
                    "top_appkey" : "12130139",
                    "topParameter" : {
                        "visitor_nick" : "sandbox_c_1",
                        "visitor_id" : "175978269",
                        "ts" : "1337240880774",
                        "iframe" : "1"
                    },
                    "top_parameters" : "aWZyYW1lPTEmdHM9MTMzNzI0MDg4MDc3NCZ2aXNpdG9yX2lkPTE3NTk3ODI2OSZ2aXNpdG9yX25pY2s9c2FuZGJveF9jXzE="
                },
                "updateTime" : 20120517154801,
                "sysUserId" : 10,
                "sysUnitId" : 10,
                "createTime" : 20120517130801,
            },
        }, {
            "_id": 2,
            "id": 2,
            "A": {
                "id" : 2,
                "resStruct" : {
                    "top_session" : "40517371e526ed5386ba9dcc0f8886504e8261b17597bmzr6J2m82691",
                    "url" : "http://127.0.0.1:8000/",
                    "agreement" : "true",
                    "app_key" : "12130139",
                    "agreementsign" : "12130139-22287347-346B2278422F673ECB579A9E682E0621",
                    "top_sign" : "5tQpWMO3Ws+eE+ygdPbgew==",
                    "top_appkey" : "12130139",
                    "topParameter" : {
                        "visitor_nick" : "sandbox_c_1",
                        "visitor_id" : "175978269",
                        "ts" : "1337240880774",
                        "iframe" : "1"
                    },
                    "top_parameters" : "aWZyYW1lPTEmdHM9MTMzNzI0MDg4MDc3NCZ2aXNpdG9yX2lkPTE3NTk3ODI2OSZ2aXNpdG9yX25pY2s9c2FuZGJveF9jXzE="
                },
                "updateTime" : 20120517154801,
                "sysUserId" : 20,
                "sysUnitId" : 20,
                "createTime" : 20120517130801,
            },
        }],
    }

def getSysUserInitData():
    return {
        'name': 'SysUser',
        'items': [{
            '_id': 10,
            'id': 10,
            'A': {
                'id': 10,
                'code': 'test10',
                'name': '测试帐户10',
                'type': 2,
                'status': 1,
                'sellerId': 10,
                'sellerNick': '测试帐户10卖家昵称',
                'nick': '测试帐户10昵称',
                "createBy" : 10,
                "createTime" : 20140227165212L,
                "modifyBy" : 0,
                "modifyTime" : 0,
                "createUnit" : 10,
                "modifyUnit" : 0,
            },
        }, {
            '_id': 20,
            'id': 20,
            'A': {
                'id': 20,
                'code': 'test20',
                'name': '测试帐户20',
                'type': 2,
                'status': 1,
                'sellerId': 20,
                'sellerNick': '测试帐户20卖家昵称',
                'nick': '测试帐户20昵称',
                "createBy" : 20,
                "createTime" : 20140227165212L,
                "modifyBy" : 0,
                "modifyTime" : 0,
                "createUnit" : 20,
                "modifyUnit" : 0,
            },
        }, {
            '_id': 1,
            'id': 1,
            'A': {
                'id': 1,
                'code': 'admin1',
                'name': 'hjq',
                'password': 'b0399d2029f64d445bd131ffaa399a42d2f8e7dc',
                'type': 1,
                'status': 1,
                'sellerId': 1,
                'sellerNick': '管理员帐户1卖家昵称',
                'nick': '管理员帐户1昵称',
                "createBy" : 1,
                "createTime" : 20140227165212L,
                "modifyBy" : 0,
                "modifyTime" : 0,
                "createUnit" : 1,
                "modifyUnit" : 0,
            },
        }],
    }

def getSysUnitInitData():
    return {
        'name': 'SysUnit',
        'items': [{
            '_id': 10,
            'id': 10,
            'A': {
                'id': 10,
                'code': 'unit10',
                'name': '店铺10',
                'sysUserId': 10,
                'sysUserNick': '管理员10_nick',
                'sid': 'sid10',
                'cid': 'cid10',
                'evaluationCount': '10',
                'eServicePermission': '1',
                "createBy" : 10,
                "createTime" : 20140227165212L,
                "modifyBy" : 0,
                "modifyTime" : 0,
                "createUnit" : 10,
                "modifyUnit" : 0,
            },
        }, {
            '_id': 20,
            'id': 20,
            'A': {
                'id': 20,
                'code': 'unit20',
                'name': '店铺20',
                'sysUserId': 20,
                'sysUserNick': '管理员20_nick',
                'sid': 'sid20',
                'cid': 'cid20',
                'evaluationCount': '20',
                'eServicePermission': '1',
                "createBy" : 20,
                "createTime" : 20140227165212L,
                "modifyBy" : 0,
                "modifyTime" : 0,
                "createUnit" : 20,
                "modifyUnit" : 0,
            },
        }, {
            '_id': 1,
            'id': 1,
            'A': {
                'id': 1,
                'code': 'unit1',
                'name': '店铺1',
                'sysUserId': 1,
                'sysUserNick': '管理员1_nick',
                'sid': 'sid1',
                'cid': 'cid1',
                'evaluationCount': '1',
                'eServicePermission': '1',
                "createBy" : 1,
                "createTime" : 20140227165212L,
                "modifyBy" : 0,
                "modifyTime" : 0,
                "createUnit" : 1,
                "modifyUnit" : 0,
            },
        }],
    }

def getSystemParameterInitData():
    return {
        'name': 'SystemParameter',
        'items': [{
            "A" : {
                "createBy" : 1,
                "modifyBy" : 0,
                "billStatus" : 0,
                "percentDecimals" : 3,
                "costDecimals" : 3,
                "attachCount" : 0,
                "taxTypeId" : 0,
                "id" : 1,
                "createTime" : 20140721105938L,
                "modifyTime" : 0,
                "createUnit" : 1,
                "modifyUnit" : 0,
                "remark" : "",
                "percentRoundingWay" : 2,
                "thousandDecimals" : 2,
                "currencyTypeId" : 3,
            },
            "_id" : 1,
            "id" : 1,
            "pendingTransactions" : [ ]
        }, {
            "A" : {
                "createBy" : 10,
                "modifyBy" : 0,
                "billStatus" : 0,
                "percentDecimals" : 3,
                "costDecimals" : 3,
                "attachCount" : 0,
                "taxTypeId" : 0,
                "id" : 10,
                "createTime" : 20140721105938L,
                "modifyTime" : 0,
                "createUnit" : 10,
                "modifyUnit" : 0,
                "remark" : "",
                "percentRoundingWay" : 2,
                "thousandDecimals" : 2,
                "currencyTypeId" : 1,
            },
            "_id" : 10,
            "id" : 10,
            "pendingTransactions" : [ ]
        }, {
            "A" : {
                "createBy" : 20,
                "modifyBy" : 0,
                "billStatus" : 0,
                "percentDecimals" : 3,
                "costDecimals" : 3,
                "attachCount" : 0,
                "taxTypeId" : 0,
                "id" : 20,
                "createTime" : 20140721105938L,
                "modifyTime" : 0,
                "createUnit" : 20,
                "modifyUnit" : 0,
                "remark" : "",
                "percentRoundingWay" : 2,
                "thousandDecimals" : 2,
                "currencyTypeId" : 2,
            },
            "_id" : 20,
            "id" : 20,
            "pendingTransactions" : [ ]
        }],
    }

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
                "createUnit" : 10,
                "modifyUnit" : 0,
                "attachCount" : 0,
                "remark" : "",
                "property" : 2,
                "id" : 1,
                "createBy" : 10,
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
                "createUnit" : 10,
                "modifyUnit" : 0,
                "attachCount" : 0,
                "remark" : "",
                "property" : 2,
                "id" : 2,
                "createBy" : 10,
                "createTime" : 20140227171435L,
                "modifyBy" : 0,
                "modifyTime" : 0
            },
            "_id" : 2,
            "id" : 2,
            "pendingTransactions" : [],
        }, {
            "A" : {
                "billTypeId" : 1,
                "createUnit" : 20,
                "modifyUnit" : 0,
                "attachCount" : 0,
                "remark" : "",
                "property" : 2,
                "id" : 3,
                "createBy" : 20,
                "createTime" : 20140227165212L,
                "modifyBy" : 0,
                "modifyTime" : 0
            },
            "_id" : 3,
            "id" : 3,
            "pendingTransactions" : [],
        }, {
            "A" : {
                "billTypeId" : 2,
                "createUnit" : 20,
                "modifyUnit" : 0,
                "attachCount" : 0,
                "remark" : "",
                "property" : 2,
                "id" : 4,
                "createBy" : 20,
                "createTime" : 20140227171435L,
                "modifyBy" : 0,
                "modifyTime" : 0
            },
            "_id" : 4,
            "id" : 4,
            "pendingTransactions" : [],
        }, {
            "A" : {
                "billTypeId" : 1,
                "createUnit" : 1,
                "modifyUnit" : 0,
                "attachCount" : 0,
                "remark" : "",
                "property" : 2,
                "id" : 5,
                "createBy" : 1,
                "createTime" : 20140227165212L,
                "modifyBy" : 0,
                "modifyTime" : 0
            },
            "_id" : 5,
            "id" : 5,
            "pendingTransactions" : [],
        }, {
            "A" : {
                "billTypeId" : 2,
                "createUnit" : 1,
                "modifyUnit" : 0,
                "attachCount" : 0,
                "remark" : "",
                "property" : 2,
                "id" : 6,
                "createBy" : 1,
                "createTime" : 20140227171435L,
                "modifyBy" : 0,
                "modifyTime" : 0
            },
            "_id" : 6,
            "id" : 6,
            "pendingTransactions" : [],
        }],
    }

def getCurrencyTypeInitData():
    return {
        'name': 'CurrencyType',
        'items': [{
            "A" : {
                "remark" : "",
                "createUnit" : 10L,
                "modifyUnit" : 10,
                "id" : 1L,
                "code" : "RMB",
                "roundingWay" : 2L,
                "upDecimals" : 3L,
                "createTime" : 20140415153316L,
                "modifyTime" : 20140429112340L,
                "currencyTypeSign" : "",
                "amtDecimals" : 3L,
                "attachCount" : 0L,
                "name" : "人民币",
                "createBy" : 10L,
                "modifyBy" : 10L,
                "billStatus" : 0L
            },
            "_id" : 1,
            "id" : 1,
            "pendingTransactions" : [ ]
        }, {
            "A" : {
                "remark" : "",
                "createUnit" : 20L,
                "modifyUnit" : 20,
                "id" : 2L,
                "code" : "RMB",
                "roundingWay" : 2L,
                "upDecimals" : 3L,
                "createTime" : 20140415153316L,
                "modifyTime" : 20140429112340L,
                "currencyTypeSign" : "",
                "amtDecimals" : 3L,
                "attachCount" : 0L,
                "name" : "人民币2",
                "createBy" : 20,
                "modifyBy" : 20,
                "billStatus" : 0L
            },
            "_id" : 2,
            "id" : 2,
            "pendingTransactions" : [ ]
        }, {
            "A" : {
                "remark" : "",
                "createUnit" : 1L,
                "modifyUnit" : 1,
                "id" : 3L,
                "code" : "RMB",
                "roundingWay" : 2L,
                "upDecimals" : 3L,
                "createTime" : 20140415153316L,
                "modifyTime" : 20140429112340L,
                "currencyTypeSign" : "",
                "amtDecimals" : 3L,
                "attachCount" : 0L,
                "name" : "人民币2",
                "createBy" : 1,
                "modifyBy" : 1,
                "billStatus" : 0L
            },
            "_id" : 3,
            "id" : 3,
            "pendingTransactions" : [ ]
        }],
    }

def getInitDataLi():
    li = []
    li.append(getSysUserInitData())
    li.append(getSysUnitInitData())
    li.append(getSystemParameterInitData())
    li.append(getBillTypeInitData())
    li.append(getBillTypeParameterInitData())
    li.append(getCurrencyTypeInitData())
    li.append(getLastSessionDataInitData())
    return li
