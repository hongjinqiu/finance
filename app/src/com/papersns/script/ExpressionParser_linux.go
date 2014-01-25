package script

import (
	"encoding/json"
	"fmt"
	"github.com/sbinet/go-python"
	"strings"
	"sync"
)

var pyrwlock sync.RWMutex = sync.RWMutex{}
var flag bool = false
var expressionMod *python.PyObject = nil
var componentMod *python.PyObject = nil

func getExpressionMod() *python.PyObject {
	pyrwlock.RLock()
	defer pyrwlock.RUnlock()

	return expressionMod
}

func getComponentMod() *python.PyObject {
	pyrwlock.RLock()
	defer pyrwlock.RUnlock()

	return componentMod
}

func isEnvInit() bool {
	pyrwlock.RLock()
	defer pyrwlock.RUnlock()

	return flag
}

func InitPythonEnv() {
	pyrwlock.Lock()
	defer pyrwlock.Unlock()

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

func (o ExpressionParser) ParseGolang(bo map[string]interface{}, data map[string]interface{}, expression string) string {
	scriptManager := ScriptManager{}
	//Parse(classMethod string, param []interface{}) []interface{} {
	paramLi := []interface{}{bo, data}
	values := scriptManager.Parse(expression, paramLi)
	return fmt.Sprint(values[0])
}

func (o ExpressionParser) Parse(recordJson, expression string) bool {
	if recordJson == "" || expression == "" {
		return true
	}
	methodName := "trueOrFalse"
	execResult := o.parseExpression(methodName, []string{recordJson, expression})
	return strings.ToLower(execResult) == "true"
}

func (o ExpressionParser) Validate(text, expression string) bool {
	if text == "" || expression == "" {
		return true
	}
	methodName := "validate"
	execResult := o.parseExpression(methodName, []string{text, expression})
	return strings.ToLower(execResult) == "true"
}

func (o ExpressionParser) ParseString(recordJson, expression string) string {
	methodName := "parseString"
	execResult := o.parseExpression(methodName, []string{recordJson, expression})
	return execResult
}

func (o ExpressionParser) ParseModel(boJson, dataJson, expression string) string {
	methodName := "parseModel"
	return o.parseExpression(methodName, []string{boJson, dataJson, expression})
}

//func (o ExpressionParser) parseExpression(methodName, recordJson, expression string) string {
func (o ExpressionParser) parseExpression(methodName string, param []string) string {
	expression := param[len(param)-1 : len(param)][0]
	if strings.TrimSpace(expression) == "" {
		return ""
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

	strfunc := getExpressionMod().GetAttrString(methodName)
	if strfunc == nil {
		panic("get function return null")
	}

	strargs := python.PyTuple_New(len(param))
	if strargs == nil {
		panic("build argument return null")
	}

	for i := 0; i < len(param); i++ {
		args := python.PyString_FromString(strings.TrimSpace(param[i]))
		python.PyTuple_SET_ITEM(strargs, i, args)
	}

	strret := strfunc.CallObject(strargs)
	if strret == nil {
		panic("call object return null")
	}

	python.PyErr_Print()
	python.PyErr_Clear()

	execResult := python.PyString_AS_STRING(strret)
	return execResult
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

/*
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
*/

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
