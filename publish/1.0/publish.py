#!/bin/bash
#coding=utf8

import sys,os,random

# Add a custom Python path.
if '..' not in sys.path:
   sys.path.append('..')

# Switch to the directory of your project. (Optional.)
#os.chdir(os.path.abspath(os.curdir))
#os.environ['DJANGO_SETTINGS_MODULE'] = "settings"

#from common.urlUtil import *
import pymongo,threading
from dictionary import *
from sequence import *
from initdata import *
from menu import *

threadLocal = threading.local()
MONGODB_HOST = 'localhost'
MONGODB_PORT = 27017
MONGODB_ADDRESS = 'localhost:27017'
MONGODB_DATABASE_NAME = 'aftermarket2'
MONGODB_USER = None
MONGODB_PASSWORD = None

def getThreadLocalMongoDB():
    if not hasattr(threadLocal, 'mongoDict'):
        #con = pymongo.Connection(MONGODB_HOST, MONGODB_PORT)
        con = pymongo.Connection(MONGODB_ADDRESS)
        if MONGODB_USER:
            con[MONGODB_DATABASE_NAME].authenticate(MONGODB_USER, MONGODB_PASSWORD)
        threadLocal.mongoDict = {
            'mongoDB': con[MONGODB_DATABASE_NAME],
            'con': con
        }
    return threadLocal.mongoDict
        
def closeThreadLocalMongoDB():
    if hasattr(threadLocal, 'mongoDict'):
        threadLocal.mongoDict['con'].disconnect()
        del threadLocal.mongoDict

def initDictionaryId():
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    dictionaryId = mongoDB.counters.find_one({'_id': 'pubDictionaryId'})
    if not dictionaryId:
        mongoDB.counters.save({'_id': 'pubDictionaryId', 'c': 1})

def initDictionary():
    initDictionaryId()
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    li = getDictionaryLi()
    codeLi = [item["code"] for item in li]
    mongoDB.PubDictionary.remove({'_id': {'$nin': codeLi}})
    for item in li:
        if not mongoDB.PubDictionary.find_one({'code': item["code"]}):
            seq = mongoDB.counters.find_and_modify(query={'_id': 'pubDictionaryId'}, update={'$inc': {'c': 1}})
            item["_id"] = seq["c"]
            item["id"] = seq["c"]
            mongoDB.PubDictionary.save(item)
        else:
            dictionaryItem = mongoDB.PubDictionary.find_one({'code': item["code"]})
            for subItem in item:
                if subItem != "id" and subItem != "_id":
                    dictionaryItem[subItem] = item[subItem]
            mongoDB.PubDictionary.save(dictionaryItem)
            

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

def initSequence():
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    li = getSequenceLi()
    idLi = [item["_id"] for item in li]
    mongoDB.counters.remove({'_id': {'$nin': idLi}})
    for item in li:
        if not mongoDB.counters.find_one({'_id': item["_id"]}):
            mongoDB.counters.save(item)

def initInitData():
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    li = getInitDataLi()
    for item in li:
        for subItem in item['items']:
            data = mongoDB[item['name']].find_one({'_id': subItem['_id']})
            mongoDB[item['name']].save(subItem)

def initDemo():
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    codeLi = []
    for i in range(25):
        codeLi.append('code' + str(i + 1))
    for item in codeLi:
        if not mongoDB.Demo.find_one({'A.code': item}):
            bLi = []
            cLi = []
            for i in range(25):
                seq = mongoDB.counters.find_and_modify(query={'_id': 'demoId'}, update={'$inc': {'c': 1}})
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
                seq = mongoDB.counters.find_and_modify(query={'_id': 'demoId'}, update={'$inc': {'c': 1}})
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

def initSysUser():
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    for item in mongoDB.SysUser.find():
        item['A'] = {
            '_id': item['_id'],
            'id': item['id'],
            'code': 'code' + str(item['_id']) + '_' + str(random.randint(0,100)),
            'name': 'name' + str(item['_id']) + '_' + str(random.randint(0,100)),
            'type': 2,
            'status': 1,
            'sellerId': random.randint(0,10000),
            'sellerNick': 'sellerNick' + str(item['_id']) + '_' + str(random.randint(0,100)),
            'nick': 'nick' + str(item['_id']) + '_' + str(random.randint(0,100)),
            'createUnit': 1,
        }
        if item["_id"] == 15:
            item["A"]["type"] = 1
        mongoDB.SysUser.save(item)

def initMenu():
    mongoDB = getThreadLocalMongoDB()['mongoDB']
    li = dealMenuLi()
    levelLi = [item["level"] for item in li]
    mongoDB.Menu.remove({'level': {'$nin': levelLi}})
    for item in li:
        if not mongoDB.Menu.find_one({'level': item["level"]}):
            seq = mongoDB.counters.find_and_modify(query={'_id': 'menuId'}, update={'$inc': {'c': 1}})
            item["_id"] = seq["c"]
            item["id"] = seq["c"]
            mongoDB.Menu.save(item)
        else:
            menuItem = mongoDB.Menu.find_one({'level': item["level"]})
            for subItem in item:
                if subItem != "id" and subItem != "_id":
                    menuItem[subItem] = item[subItem]
            mongoDB.Menu.save(menuItem)

if __name__ == '__main__':
    initSequence()
    initDictionary()
    initInitData()
    initDemo()
    initMenu()
    #initSysUser()
