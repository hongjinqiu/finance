#!/bin/bash
#coding=utf8

import sys,os,random

# Add a custom Python path.
if '..' not in sys.path:
   sys.path.append('..')

# Switch to the directory of your project. (Optional.)
os.chdir(os.path.abspath(os.curdir))
os.environ['DJANGO_SETTINGS_MODULE'] = "settings"

from common.urlUtil import *

def initDictionaryId():
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    dictionaryId = mongoDB.counters.find_one({'_id': 'pubDictionaryId'})
    if not dictionaryId:
        mongoDB.counters.save({'_id': 'pubDictionaryId', 'c': 1})

def initDictionary():
    initDictionaryId()
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    if True:
        mongoDB.PubDictionary.remove({'code': 'D_DICTTEST'})
    if not mongoDB.PubDictionary.find_one({'code': 'D_DICTTEST'}):
        seq = mongoDB.counters.find_and_modify(query={'_id': 'pubDictionaryId'}, update={'$inc': {'c': 1}})
        mongoDB.PubDictionary.save({
            '_id': seq['c'],
            'id': seq['c'],
            'code': 'D_DICTTEST',
            'name': '字典测试',
            'items': [{
                'code': 0,
                'name': '测试项1',
                'order': 0,
            }, {
                'code': 1,
                'name': '测试项2',
                'order': 2,
                'items': [{
                    'code': 0,
                    'name': '测试项2_0',
                    'order': 1,
                },{
                    'code': 1,
                    'name': '测试项2_1',
                    'order': 0,
                }],
            }, {
                'code': 2,
                'name': '测试项3',
                'order': 1,
            }],
        })

def initTransactionsId():
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    dictionaryId = mongoDB.counters.find_one({'_id': 'transactionsId'})
    if not dictionaryId:
        mongoDB.counters.save({'_id': 'transactionsId', 'c': 1})

def initTest1Id():
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    dictionaryId = mongoDB.counters.find_one({'_id': 'test1Id'})
    if not dictionaryId:
        mongoDB.counters.save({'_id': 'test1Id', 'c': 1})

def initTest2Id():
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    dictionaryId = mongoDB.counters.find_one({'_id': 'test2Id'})
    if not dictionaryId:
        mongoDB.counters.save({'_id': 'test2Id', 'c': 1})

def initDemoId():
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    demoId = mongoDB.counters.find_one({'_id': 'demoId'})
    if not demoId:
        mongoDB.counters.save({'_id': 'demoId', 'c': 1})
    demoId = mongoDB.counters.find_one({'_id': 'demoBId'})
    if not demoId:
        mongoDB.counters.save({'_id': 'demoBId', 'c': 1})
    demoId = mongoDB.counters.find_one({'_id': 'demoCId'})
    if not demoId:
        mongoDB.counters.save({'_id': 'demoCId', 'c': 1})

def initDemo():
    initDemoId()
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    codeLi = []
    for i in range(25):
        codeLi.append('code' + str(i + 1))
    for item in codeLi:
        if not mongoDB.Demo.find_one({'A.code': item}):
            bLi = []
            cLi = []
            for i in range(25):
                seq = mongoDB.counters.find_and_modify(query={'_id': 'demoBId'}, update={'$inc': {'c': 1}})
                bLi.append({
                    '_id': seq['c'],
                    'id': seq['c'],
                    "createBy": 1,
                    "createTime": 20130102030405 + random.randint(0,50),
                    "createUnit": 1,
                    "modifyBy": 1,
                    "modifyUnit": 1,
                    "modifyTime": 20130102030405 + random.randint(0,50),
                    "attachCount": 2,
                    "remark": "备注",
                    "code": item,
                    "name": item + "_名称",
                    "stringColumn": "字符串列",
                    "moneyNumberColumn": 3.1 + random.randint(0,100),
                    "priceNumberColumn": 3.2 + random.randint(0,100),
                    "unitCostNumberColumn": 3.3 + random.randint(0,100),
                    "percentNumberColumn": 3.4 + random.randint(0,100),
                    "quantityNumberColumn": 3.5 + random.randint(0,100),
                    "dateTest": 20130102 + random.randint(0,3),
                    "timeTest": 10203 + random.randint(0,7) * 100 + random.randint(0,50),
                    "dateTimeTest": 2013010203150405 + random.randint(0,6) * 100 + random.randint(0,50),
                    "yearTest": 2013 + random.randint(0,3),
                    "yearMonthTest": 201301 + random.randint(0,8),
                    "dictionaryTest": random.randint(0,2),
                    "selectTest": 1,
                })
                seq = mongoDB.counters.find_and_modify(query={'_id': 'demoCId'}, update={'$inc': {'c': 1}})
                cLi.append({
                    '_id': seq['c'],
                    'id': seq['c'],
                    "createBy": 1,
                    "createTime": 20130102030405 + random.randint(0,50),
                    "createUnit": 1,
                    "modifyBy": 1,
                    "modifyUnit": 1,
                    "modifyTime": 20130102030405 + random.randint(0,50),
                    "attachCount": 2,
                    "remark": "备注",
                    "code": item,
                    "name": item + "_名称",
                    "stringColumn": "字符串列",
                    "moneyNumberColumn": 3.1 + random.randint(0,100),
                    "priceNumberColumn": 3.2 + random.randint(0,100),
                    "unitCostNumberColumn": 3.3 + random.randint(0,100),
                    "percentNumberColumn": 3.4 + random.randint(0,100),
                    "quantityNumberColumn": 3.5 + random.randint(0,100),
                    "dateTest": 20130102 + random.randint(0,3),
                    "timeTest": 10203 + random.randint(0,7) * 100 + random.randint(0,50),
                    "dateTimeTest": 2013010203150405 + random.randint(0,6) * 100 + random.randint(0,50),
                    "yearTest": 2013 + random.randint(0,3),
                    "yearMonthTest": 201301 + random.randint(0,8),
                    "dictionaryTest": random.randint(0,2),
                    "selectTest": 1,
                })
            
            seq = mongoDB.counters.find_and_modify(query={'_id': 'demoId'}, update={'$inc': {'c': 1}})
            mongoDB.Demo.save({
                '_id': seq['c'],
                'id': seq['c'],
                'A': {
                    '_id': seq['c'],
                    'id': seq['c'],
                    "createBy": 1,
                    "createTime": 20130102030405,
                    "createUnit": 1,
                    "modifyBy": 1,
                    "modifyUnit": 1,
                    "modifyTime": 20140102030405,
                    "attachCount": 2,
                    "remark": "备注",
                    "code": item,
                    "name": item + "_名称",
                    "stringColumn": "字符串列",
                    "moneyNumberColumn": 3.1 + random.randint(0,100),
                    "priceNumberColumn": 3.2 + random.randint(0,100),
                    "unitCostNumberColumn": 3.3 + random.randint(0,100),
                    "percentNumberColumn": 3.4 + random.randint(0,100),
                    "quantityNumberColumn": 3.5 + random.randint(0,100),
                    "dateTest": 20130102 + random.randint(0,3),
                    "timeTest": 10203 + random.randint(0,7) * 100 + random.randint(0,50),
                    "dateTimeTest": 2013010203150405 + random.randint(0,6) * 100 + random.randint(0,50),
                    "yearTest": 2013 + random.randint(0,3),
                    "yearMonthTest": 201301 + random.randint(0,8),
                    "dictionaryTest": random.randint(0,2),
                    "selectTest": 1,
                },
                'B': bLi,
                'C': cLi,
            })


if __name__ == '__main__':
    initDictionary()
    initTransactionsId()
    initTest1Id()
    initTest2Id()
    initDemoId()
    initDemo()


