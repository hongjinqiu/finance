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
            item['numTest'] = 1000 * 1000 + 0.12345678901
            item['numTest1'] = 1000 * 1000 + 0.12345678901
            item['numTest2'] = 1000 * 1000 + 0.12345678901
            item['numTest3'] = 1000 * 1000 + 0.12345678901
            item['numTest4'] = 1000 * 1000 + 0.12345678901
            item['numTest5'] = 1000 * 1000 + 0.12345678901
            item['numTest6'] = 1000 * 1000 + 0.12345678901
            item['numTest7'] = 1000 * 1000 + 0.12345678901
            item['numTest8'] = 1000 * 1000 + 0.12345678901

        for i,item in enumerate(items):
            if i == 0:
                item['currency'] = {
                    'prefix': "$",
                    'decimalPlaces': 3,
                    'unitPriceDecimalPlaces': 6,
                    'decimalSeparator': "^",
                    'thousandsSeparator': "_",
                    'suffix': "&"
                }
            else:
                item['currency'] = {
                    'prefix': "*",
                    'decimalPlaces': 4,
                    'unitPriceDecimalPlaces': 8,
                    'decimalSeparator': "@",
                    'thousandsSeparator': "=",
                    'suffix': "!"
                }
                
        for item in items:
            item['dateTest'] = 20131020
            item['dateTimeTest'] = 20131020114825
            
        for item in items:
            item['boolTest'] = True
            item['boolTest2'] = False

        return json.dumps(items)    

if __name__ == '__main__':
    pass
