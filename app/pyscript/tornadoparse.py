#!/usr/bin/python
#encoding=utf8

import tornado.ioloop
import tornado.web
import json,exceptions,traceback,datetime, sys, re, threading, time, urllib, os, hashlib

#def prepare(self):
#    if self.request.headers["Content-Type"].startswith("application/json"):
#        self.json_args = json.loads(self.request.body)
#    else:
#        self.json_args = None

#        RequestHandler.get_argument(name, default=, []strip=True)[source]
#        RequestHandler.get_arguments(name, strip=True)[source]
#        RequestHandler.get_query_argument(name, default=, []strip=True)[source]
#        RequestHandler.get_query_arguments(name, strip=True)[source]
#        RequestHandler.get_body_argument(name, default=, []strip=True)[source]
#        RequestHandler.get_body_arguments(name, strip=True)[source]
#        RequestHandler.decode_argument(value, name=None)[source]

class MainHandler(tornado.web.RequestHandler):
    def get(self):
        self.set_header("Content-Type", "application/json; charset=utf-8")
        self.write(json.dumps({'name': '测试'}))
        
    def post(self):
        self.set_header("Content-Type", "text/plain")
        self.write("You wrote " + self.get_body_argument("message"))

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

class PyScriptParseHandler(tornado.web.RequestHandler):
    def get(self):
        self.set_header("Content-Type", "text/plain")
        self.write("don't use get, use post")
        
    def post(self):
        self.set_header("Content-Type", "text/plain")
        
        try:
            method = self.get_argument('method', default='')
            jsonString = self.get_argument('jsonString', default='')
            action = self.get_argument('action', default='')
            bo = self.get_argument('bo', default='')
            data = self.get_argument('data', default='')
    
            if method == "trueOrFalse":
                self.write(trueOrFalse(jsonString, action))
            if method == "validate":
                self.write(validate(jsonString, action))
            if method == "parseString":
                self.write(parseString(jsonString, action))
            if method == "parseModel":
                self.write(parseModel(bo, data, action))
        except exceptions.Exception:
            print(traceback.format_exc())
            self.write("bad request")
        finally:
            pass

application = tornado.web.Application([
    (r"/", MainHandler),
    (r"/pyscript/parse", PyScriptParseHandler),
    (r"/pyscript/parse/", PyScriptParseHandler),
])

if __name__ == "__main__":
    application.listen(8000)
    tornado.ioloop.IOLoop.instance().start()
