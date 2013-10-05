package component

import (
	"github.com/sbinet/go-python"
	"strings"
	"sync"
	"encoding/json"
)

var rwlock sync.RWMutex = sync.RWMutex{}
var flag bool = false
var expressionMod *python.PyObject = nil
var componentMod *python.PyObject = nil

func getExpressionMod() *python.PyObject {
	rwlock.RLock()
	defer rwlock.RUnlock()
	
	return expressionMod
}

func getComponentMod() *python.PyObject {
	rwlock.RLock()
	defer rwlock.RUnlock()
	
	return componentMod
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
		panic("get module expression return null")
	}
	
	componentMod = python.PyImport_ImportModule("component")
	if componentMod == nil {
		panic("get module component return null")
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
	args2 := python.PyString_FromString(strings.TrimSpace(expression))

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

func (o ExpressionParser) ParseBeforeBuildQuery(classMethod string, paramMap map[string]string) map[string]string {
	if strings.TrimSpace(classMethod) == "" {
		return paramMap
	}
	
	execResult := o.parseClassMethod(classMethod, paramMap)
	
	result := map[string]string{}
	err := json.Unmarshal([]byte(execResult), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func (o ExpressionParser) ParseAfterBuildQuery(classMethod string, queryLi []map[string]interface{}) []map[string]interface{} {
	if strings.TrimSpace(classMethod) == "" {
		return queryLi
	}
	
	execResult := o.parseClassMethod(classMethod, queryLi)
	
	result := []map[string]interface{}{}
	err := json.Unmarshal([]byte(execResult), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func (o ExpressionParser) ParseAfterQueryData(classMethod string, items []interface{}) []interface{} {
	if strings.TrimSpace(classMethod) == "" {
		return items
	}

	execResult := o.parseClassMethod(classMethod, items)
	
	result := []interface{}{}
	err := json.Unmarshal([]byte(execResult), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func (o ExpressionParser) parseClassMethod(classMethod string, obj interface{}) string {
	if !isEnvInit() {
		InitPythonEnv()
	}
	
	className := strings.Split(classMethod, ".")[0]
	methodName := strings.Split(classMethod, ".")[1]
	
	class := getComponentMod().GetAttrString(className)
	if class == nil {
		panic("get class:" + className + " return nil")
	}
	
	classArgs := python.PyTuple_New(0)
	object := class.CallObject(classArgs)
	
	method := object.GetAttrString(methodName)
	if method == nil {
		panic("get method:" + methodName + " return nil")
	}

	jsonStringByte, err := json.Marshal(&obj)
	if err != nil {
		panic(err)
	}
	
	strargs := python.PyTuple_New(1)
	args1 := python.PyString_FromString(string(jsonStringByte))
	python.PyTuple_SET_ITEM(strargs, 0, args1)
	
	strret := method.CallObject(strargs)
	if strret == nil {
		panic("call object return null")
	}

	python.PyErr_Print()
	python.PyErr_Clear()

	return python.PyString_AS_STRING(strret)
}

