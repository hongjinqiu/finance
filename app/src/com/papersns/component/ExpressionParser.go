package component

import (
	"github.com/sbinet/go-python"
	"strings"
)

type ExpressionParser struct{}

func (o ExpressionParser) initEnv() {
	err := python.Initialize()
	if err != nil {
		panic(err)
	}

	sys_path := python.PySys_GetObject("path")
	if sys_path == nil {
		panic("get sys.path return nil")
	}

	path := python.PyString_FromString("/home/hongjinqiu/goworkspace/src/finance")
	if path == nil {
		panic("get path return nil")
	}

	err = python.PyList_Append(sys_path, path)
	if err != nil {
		panic(err)
	}
	
}

func (o ExpressionParser) exitEnv() {
	python.Finalize()
}

func (o ExpressionParser) Parse(recordJson, expression string) bool {
	if strings.TrimSpace(expression) == "" {
		return true
	}
	
	o.initEnv()
	defer o.exitEnv()

	mymod := python.PyImport_ImportModule("expression")
	if mymod == nil {
		panic("get module return null")
	}

	strfunc := mymod.GetAttrString("trueOrFalse")
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
