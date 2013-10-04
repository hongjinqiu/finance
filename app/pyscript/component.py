#!/usr/bin/python
#encoding=utf8
import json,sys

class SysUser():
#    def beforeBuildQuery(self, jsonString):
    def beforeBuildQuery(self):
        return 'before build query'
#        query = json.loads(jsonString)
#        query["nick"] = u'测试'
#        return json.dumps(query)

    def afterQueryData(self, jsonString):
        items = json.loads(jsonString)
        for item in items:
            item["nick"] = u'测试'
            #item["nick"] = u'\u6d4b\u8bd5'
        return json.dumps(items)    

if __name__ == '__main__':
    pass
