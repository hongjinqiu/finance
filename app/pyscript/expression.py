#!/usr/bin/python
#encoding=utf8

import json,sys

def trueOrFalse(jsonString, action):
    record = json.loads(jsonString)
    return str(eval(action) == True)

def parseString(jsonString, action):
    record = json.loads(jsonString)
    return str(eval(action))

#trueOrFalse "\"{\\\"name\\\": \\\"test\\\"}\"" "\"record[\\\"name\\\"] == \\\"test\\\"\""
if __name__ == '__main__':
    if len(sys.argv) < 4:
        sys.exit('input invalidate, len(argv) must == 3')
        
    methodName = sys.argv[1]
    jsonString = sys.argv[2]
    action = sys.argv[3]
    
    print eval('%s(%s,%s)' % (methodName, jsonString, action))
    