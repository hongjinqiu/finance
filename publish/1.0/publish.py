#!/bin/bash
#coding=utf8

import sys,os

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

if __name__ == '__main__':
    initDictionary()

