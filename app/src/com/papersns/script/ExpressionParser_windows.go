package script

import "github.com/robfig/revel"
import (
	"sync"
	"fmt"
//	"os"
//	"os/exec"
	"strings"
	"log"
	"regexp"
//	"strconv"
	"io/ioutil"
	"net/http"
	"net/url"
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
	
	values := url.Values{}
	values.Add("method", "parse")
	values.Add("jsonString", recordJson)
	values.Add("action", expression)

	PYTHON_PARSE_URL := revel.Config.StringDefault("PYTHON_PARSE_URL", "")
	resp, err := http.PostForm(PYTHON_PARSE_URL, values)
	if err != nil {
		log.Print("parse request url ", PYTHON_PARSE_URL, " return error, values is:", values)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	result := strings.ToLower(string(body))
	regx := regexp.MustCompile(`^\s+|\s+$`)
	result = regx.ReplaceAllString(result, "")
	if result == "bad request" {
		log.Print("Parse(recordJson, expression string) bool")
		log.Print("recordJson:" + recordJson)
		log.Print("expression:" + expression)
		panic("parse error")
	}
	return result == "true"
	
	/*
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
	regx := regexp.MustCompile(`^\s+|\s+$`)
	result = regx.ReplaceAllString(result, "")
	return result == "true"
	*/
}

func (o ExpressionParser) Validate(text, expression string) bool {
	if text == "" || expression == "" {
		return true
	}
	
	values := url.Values{}
	values.Add("method", "validate")
	values.Add("jsonString", text)
	values.Add("action", expression)

	PYTHON_PARSE_URL := revel.Config.StringDefault("PYTHON_PARSE_URL", "")
	resp, err := http.PostForm(PYTHON_PARSE_URL, values)
	if err != nil {
		log.Print("validate request url ", PYTHON_PARSE_URL, " return error, values is:", values)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	result := strings.ToLower(string(body))
	regx := regexp.MustCompile(`^\s+|\s+$`)
	result = regx.ReplaceAllString(result, "")
	if result == "bad request" {
		log.Print("Validate(text, expression string) bool")
		log.Print("text:" + text)
		log.Print("expression:" + expression)
		panic("parse error")
	}
	return result == "true"
	
	/*
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
	regx := regexp.MustCompile(`^\s+|\s+$`)
	result = regx.ReplaceAllString(result, "")
	return result == "true"
	*/
}

func (o ExpressionParser) ParseString(recordJson, expression string) string {
	if expression == "" {
		return ""
	}
	if recordJson == "" {
		recordJson = "{}"
	}
	
	values := url.Values{}
	values.Add("method", "parseString")
	values.Add("jsonString", recordJson)
	values.Add("action", expression)

	PYTHON_PARSE_URL := revel.Config.StringDefault("PYTHON_PARSE_URL", "")
	resp, err := http.PostForm(PYTHON_PARSE_URL, values)
	if err != nil {
		log.Print("ParseString request url ", PYTHON_PARSE_URL, " return error, values is:", values)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	result := strings.ToLower(string(body))
	regx := regexp.MustCompile(`^\s+|\s+$`)
	result = regx.ReplaceAllString(result, "")
	if result == "bad request" {
		log.Print("ParseString(recordJson, expression string) string")
		log.Print("recordJson:" + recordJson)
		log.Print("expression:" + expression)
		panic("parse error")
	}
	return result
	
	/*
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
	regx := regexp.MustCompile(`^\s+|\s+$`)
	result = regx.ReplaceAllString(result, "")
	return result
	*/
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
	
	values := url.Values{}
	values.Add("method", "parseModel")
	values.Add("bo", boJson)
	values.Add("data", dataJson)
	values.Add("action", expression)

	PYTHON_PARSE_URL := revel.Config.StringDefault("PYTHON_PARSE_URL", "")
	resp, err := http.PostForm(PYTHON_PARSE_URL, values)
	if err != nil {
		log.Print("ParseModel request url ", PYTHON_PARSE_URL, " return error, values is:", values)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	result := strings.ToLower(string(body))
	regx := regexp.MustCompile(`^\s+|\s+$`)
	result = regx.ReplaceAllString(result, "")
	if result == "bad request" {
		log.Print("ParseString(recordJson, expression string) string")
		log.Print("boJson:" + boJson)
		log.Print("dataJson:" + dataJson)
		log.Print("expression:" + expression)
		panic("parse error")
	}
	return result
	
	/*
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
	regx := regexp.MustCompile(`^\s+|\s+$`)
	result = regx.ReplaceAllString(result, "")
	return result
	*/
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
