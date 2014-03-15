#!/bin/bash
#coding=utf8

def getBillTypeInitData():
    return {
        'name': 'BillType',
        'items': [{
            '_id': 1,
            'id': 1,
            'code': '001',
            'name': '收款单',
        }, {
            '_id': 2,
            'id': 2,
            'code': '002',
            'name': '付款单',
        }],
    }
    

def getInitDataLi():
    li = []
    li.append(getBillTypeInitData())
    return li
