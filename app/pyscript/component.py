#!/usr/bin/python
#encoding=utf8
import json, sys

class SysUser():
#    def beforeBuildQuery(self, jsonString):
    def beforeBuildQuery(self, jsonString):
        query = json.loads(jsonString)
#        query["nick"] = u''
        return json.dumps(query)

    def afterBuildQuery(self, jsonString):
        query = json.loads(jsonString)
#        for item in query:
#            if item.keys()[0] == "nick":
#                item['nick'] = 20
#        query.append({"age": 20})
        return json.dumps(query)

    def afterQueryData(self, jsonString):
        items = json.loads(jsonString)
        for item in items:
            item["nick"] = u'测试 by python'
            item['UNIT_NAME'] = u'单位名称aaa'
        return json.dumps(items)    

if __name__ == '__main__':
    pass
