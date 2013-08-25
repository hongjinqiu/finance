import json

def trueOrFalse(jsonString, action):
    record = json.loads(jsonString)
    return str(eval(action) == True)
