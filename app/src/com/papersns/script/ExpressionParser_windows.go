package script

import (
	"sync"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"log"
	"regexp"
	"strconv"
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
	if expression == "" {
		return true
	}
	if recordJson == "" {
		recordJson = "{}"
	}
	os.Chdir("/goworkspace/src/finance/app/pyscript")
	recordJson = strconv.QuoteToASCII(recordJson)
	recordJson = recordJson[1:len(recordJson) - 1]
	
	expression = strconv.QuoteToASCII(expression)
	expression = expression[1:len(expression) - 1]
	
	cmd := exec.Command("python", "expression.py", "trueOrFalse", recordJson, expression)
	out, err := cmd.Output()
	if err != nil {
		log.Println("Parse(recordJson, expression string) bool")
		log.Println("recordJson:" + recordJson)
		log.Println("expression:" + expression)
		panic(err)
	}
	result := strings.ToLower(string(out))
	regx := regexp.MustCompile(`\s+$`)
	result = regx.ReplaceAllString(result, "")
	return result == "true"
}

func (o ExpressionParser) Validate(text, expression string) bool {
	if text == "" || expression == "" {
		return true
	}
	os.Chdir("/goworkspace/src/finance/app/pyscript")
	text = strconv.QuoteToASCII(text)
	text = text[1:len(text) - 1]
	
	expression = strconv.QuoteToASCII(expression)
	expression = expression[1:len(expression) - 1]
	
	cmd := exec.Command("python", "expression.py", "validate", text, expression)
	out, err := cmd.Output()
	if err != nil {
		log.Println("Validate(text, expression string) bool")
		log.Println("text:" + text)
		log.Println("expression:" + expression)
		panic(err)
	}
	result := strings.ToLower(string(out))
	regx := regexp.MustCompile(`\s+$`)
	result = regx.ReplaceAllString(result, "")
	return result == "true"
}

func (o ExpressionParser) ParseString(recordJson, expression string) string {
	if expression == "" {
		return ""
	}
	if recordJson == "" {
		recordJson = "{}"
	}
	os.Chdir("/goworkspace/src/finance/app/pyscript")
	recordJson = strconv.QuoteToASCII(recordJson)
	recordJson = recordJson[1:len(recordJson) - 1]
	
	expression = strconv.QuoteToASCII(expression)
	expression = expression[1:len(expression) - 1]
	
	cmd := exec.Command("python", "expression.py", "parseString", recordJson, expression)
	out, err := cmd.Output()
	if err != nil {
		log.Println("ParseString(recordJson, expression string) string")
		log.Println("recordJson:" + recordJson)
		log.Println("expression:" + expression)
		panic(err)
	}
	result := string(out)
	regx := regexp.MustCompile(`\s+$`)
	result = regx.ReplaceAllString(result, "")
	return result
}

func (o ExpressionParser) ParseModel(boJson, dataJson, expression string) string {
	if expression == "" {
		return ""
	}
	if boJson == "" {
		boJson = "{}"
	}
	if dataJson == "" {
		dataJson = "{}"
	}
	os.Chdir("/goworkspace/src/finance/app/pyscript")
	boJson = strconv.QuoteToASCII(boJson)
	boJson = boJson[1:len(boJson) - 1]
	
	dataJson = strconv.QuoteToASCII(dataJson)
	dataJson = dataJson[1:len(dataJson) - 1]
	
	expression = strconv.QuoteToASCII(expression)
	expression = expression[1:len(expression) - 1]
	
	cmd := exec.Command("python", "expression.py", "parseModel", boJson, dataJson, expression)
	out, err := cmd.Output()
	if err != nil {
		log.Println("ParseString(recordJson, expression string) string")
		log.Println("boJson:" + boJson)
		log.Println("dataJson:" + dataJson)
		log.Println("expression:" + expression)
		panic(err)
	}
	result := string(out)
	regx := regexp.MustCompile(`\s+$`)
	result = regx.ReplaceAllString(result, "")
	return result
}

func (o ExpressionParser) ParseBeforeBuildQuery(classMethod string, paramMap map[string]string) map[string]string {
	return paramMap
}

func (o ExpressionParser) ParseAfterBuildQuery(classMethod string, queryLi []map[string]interface{}) []map[string]interface{} {
	return queryLi
}

/*
func (o ExpressionParser) ParseAfterQueryData(classMethod string, items []interface{}) []interface{} {
	return items
}
*/
