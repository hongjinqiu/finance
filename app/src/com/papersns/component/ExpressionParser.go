package component

import (
	"github.com/sbinet/go-python"
	"strings"
	"sync"
)

var rwlock sync.RWMutex = sync.RWMutex{}
var flag bool = false
var expressionMod *python.PyObject = nil

func getExpressionMod() *python.PyObject {
	rwlock.RLock()
	defer rwlock.RUnlock()
	
	return expressionMod
}

func isEnvInit() bool {
	rwlock.RLock()
	defer rwlock.RUnlock()
	
	return flag
}

func InitPythonEnv() {
	rwlock.Lock()
	defer rwlock.Unlock()

	if flag {
		return
	}	
	
	err := python.Initialize()
	if err != nil {
		panic(err)
	}

	sys_path := python.PySys_GetObject("path")
	if sys_path == nil {
		panic("get sys.path return nil")
	}

	path := python.PyString_FromString("/home/hongjinqiu/goworkspace/src/finance/app/pyscript")
	if path == nil {
		panic("get path return nil")
	}

	err = python.PyList_Append(sys_path, path)
	if err != nil {
		panic(err)
	}
	
	expressionMod = python.PyImport_ImportModule("expression")
	if expressionMod == nil {
		panic("get module return null")
	}
	
	flag = true
}

func exitEnv() {
	python.Finalize()
}

type ExpressionParser struct{}

func (o ExpressionParser) Parse(recordJson, expression string) bool {
	if strings.TrimSpace(expression) == "" {
		return true
	}
	if !isEnvInit() {
		InitPythonEnv()
	}
	/*
	o.InitPythonEnv()
	defer o.exitEnv()

	*/
	/*
	expressionMod := python.PyImport_ImportModule("expression")
	if expressionMod == nil {
		panic("get module return null")
	}
	*/

	strfunc := getExpressionMod().GetAttrString("trueOrFalse")
	if strfunc == nil {
		panic("get function return null")
	}

	args1 := python.PyString_FromString(recordJson)
	args2 := python.PyString_FromString(expression)

	strargs := python.PyTuple_New(2)
	if strargs == nil {
		panic("build argument return null")
	}

	python.PyTuple_SET_ITEM(strargs, 0, args1)
	python.PyTuple_SET_ITEM(strargs, 1, args2)

	strret := strfunc.CallObject(strargs)
	if strret == nil {
		panic("call object return null")
	}

	python.PyErr_Print()
	python.PyErr_Clear()

	execResult := python.PyString_AS_STRING(strret)
	return strings.ToLower(execResult) == "true"
}
