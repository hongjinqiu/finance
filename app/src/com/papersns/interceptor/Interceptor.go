package interceptor

import (
	"sync"
	"reflect"
	"strings"
)

var rwlock sync.RWMutex = sync.RWMutex{}
var interceptorDict map[string]reflect.Type = map[string]reflect.Type{}

func init() {
	rwlock.Lock()
	defer rwlock.Unlock()
	interceptorDict[reflect.TypeOf(SysUserInterceptor{}).Name()] = reflect.TypeOf(SysUserInterceptor{})
	interceptorDict[reflect.TypeOf(ModelListTemplateInterceptor{}).Name()] = reflect.TypeOf(ModelListTemplateInterceptor{})
	interceptorDict[reflect.TypeOf(PubReferenceLogInterceptor{}).Name()] = reflect.TypeOf(PubReferenceLogInterceptor{})
}

func GetInterceptorDict() map[string]reflect.Type {
	rwlock.RLock()
	defer rwlock.RUnlock()
	return interceptorDict
}

type InterceptorManager struct{}

func (o InterceptorManager) ParseBeforeBuildQuery(classMethod string, paramMap map[string]string) map[string]string {
	if classMethod == "" {
		return paramMap
	}

	paramLi := []interface{}{}
	paramLi = append(paramLi, paramMap)
	values := o.parse(classMethod, paramLi)
	if values != nil {
		return values[0].(map[string]string)
	}
	return paramMap
}

func (o InterceptorManager) ParseAfterBuildQuery(classMethod string, queryLi []map[string]interface{}) []map[string]interface{} {
	if classMethod == "" {
		return queryLi
	}
	
	paramLi := []interface{}{}
	paramLi = append(paramLi, queryLi)
	values := o.parse(classMethod, paramLi)
	if values != nil {
		return values[0].([]map[string]interface{})
	}
	return queryLi
}

func (o InterceptorManager) ParseAfterQueryData(classMethod string, dataSetId string, items []interface{}) []interface{} {
	if classMethod == "" {
		return items
	}

	paramLi := []interface{}{}
	paramLi = append(paramLi, dataSetId)
	paramLi = append(paramLi, items)
	values := o.parse(classMethod, paramLi)
	if values != nil {
		return values[0].([]interface{})
	}
	return items
}

func (o InterceptorManager) parse(classMethod string, param []interface{}) []interface{} {
	if classMethod != "" {
		exprContent := classMethod
		scriptStruct := strings.Split(exprContent, ".")[0]
		scriptStructMethod := strings.Split(exprContent, ".")[1]
		scriptType := GetInterceptorDict()[scriptStruct]
		if scriptType == nil {
			panic(scriptStruct + " is not exist")
		}
		inst := reflect.New(scriptType).Elem().Interface()
		instValue := reflect.ValueOf(inst)
		in := []reflect.Value{}
		for i, _ := range param {
			in = append(in, reflect.ValueOf(param[i]))
		}
		values := instValue.MethodByName(scriptStructMethod).Call(in)
		result := []interface{}{}
		for i, _ := range values {
			result = append(result, values[i].Interface())
		}
		return result
	}
	return nil
}
