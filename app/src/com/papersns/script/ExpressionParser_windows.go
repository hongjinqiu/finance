package script

import (
	"sync"
	"fmt"
)

var pyrwlock sync.RWMutex = sync.RWMutex{}
var flag bool = false

type ExpressionParser struct{}

func (o ExpressionParser) ParseGolang(bo map[string]interface{}, data map[string]interface{}, expression string) string {
	scriptManager := ScriptManager{}
	//Parse(classMethod string, param []interface{}) []interface{} {
	paramLi := []interface{}{bo, data}
	values := scriptManager.Parse(expression, paramLi)
	return fmt.Sprint(values[0])
}

func (o ExpressionParser) Parse(recordJson, expression string) bool {
	return true
}

func (o ExpressionParser) Validate(text, expression string) bool {
	return true
}

func (o ExpressionParser) ParseString(recordJson, expression string) string {
	return ""
}

func (o ExpressionParser) ParseModel(boJson, dataJson, expression string) string {
	return ""
}

func (o ExpressionParser) ParseBeforeBuildQuery(classMethod string, paramMap map[string]string) map[string]string {
	return paramMap
}

func (o ExpressionParser) ParseAfterBuildQuery(classMethod string, queryLi []map[string]interface{}) []map[string]interface{} {
	return queryLi
}

func (o ExpressionParser) ParseAfterQueryData(classMethod string, items []interface{}) []interface{} {
	return items
}

