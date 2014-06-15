#!/usr/bin/python
#encoding=utf8

import json,sys,exceptions,traceback,datetime

def trueOrFalse(jsonString, action):
    data = json.loads(jsonString)
    return str(eval(action) == True)
    
def validate(text, action):
    return str(eval(action) == True)

def parseString(jsonString, action):
    data = json.loads(jsonString)
    return str(eval(action))
    
def parseModel(bo, data, action):
    bo = json.loads(bo)
    data = json.loads(data)
    return str(eval(action))

#trueOrFalse "\"{\\\"name\\\": \\\"test\\\"}\"" "\"record[\\\"name\\\"] == \\\"test\\\"\""
if __name__ == '__main__':
#    if len(sys.argv) < 6:
#        sys.exit('input invalidate, len(argv) must == 3')

    command = sys.argv[1] + '(' + ','.join(['"' + item + '"' for item in sys.argv[2:]]) + ')'
    try:
        print eval(command)
    except exceptions.Exception:
        fOut = open('error.txt', 'w')
        fOut.write(command)
        fOut.write('\n')
        fOut.write(traceback.format_exc())
        fOut.close()
        raise exceptions.Exception(command)
