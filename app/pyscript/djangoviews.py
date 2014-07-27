# Create your views here.
#! /usr/bin/env python
# coding=utf-8

# Create your views here.
from django.shortcuts import render_to_response
from django.shortcuts import get_object_or_404
from django.http import HttpResponse, HttpResponseRedirect, Http404, HttpResponseBadRequest
from django.template import RequestContext
from django.contrib.sessions.models import Session
from models import *
import datetime, sys, settings, MySQLdb, re, threading, time, urllib, os, hashlib, cStringIO, Image, random, traceback, exceptions, json

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

def parse(request):
    try:
        #pass
        #request.GET.get('traceId')
        method = request.POST.get('method')
        jsonString = request.POST.get('jsonString')
        action = request.POST.get('action')
        bo = request.POST.get('bo')
        data = request.POST.get('data')

        if method == "trueOrFalse":
            return HttpResponse(trueOrFalse(jsonString, action))
        if method == "validate":
            return HttpResponse(validate(jsonString, action))
        if method == "parseString":
            return HttpResponse(parseString(jsonString, action))
        if method == "parseModel":
            return HttpResponse(parseModel(bo, data, action))
        return HttpResponse("bad request")
    except exceptions.Exception:
        print(traceback.format_exc())
        return HttpResponse("bad request")
    finally:
        pass
