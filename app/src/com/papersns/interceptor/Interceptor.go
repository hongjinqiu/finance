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
}

func GetInterceptorDict() map[string]reflect.Type {
	rwlock.RLock()
	defer rwlock.RUnlock()
	return interceptorDict
}

type InterceptorManager struct{}

func (o InterceptorManager) ParseBeforeBuildQuery(classMethod string, paramMap *map[string]string) {
	//paramLi := []*interface{}{paramMap}
	paramLi := []*interface{}{}
	var paramPointer interface{} = *paramMap
	paramLi = append(paramLi, &paramPointer)
	o.parse(classMethod, &paramLi)
}

func (o InterceptorManager) ParseAfterBuildQuery(classMethod string, queryLi *[]map[string]interface{}) {
	paramLi := []*interface{}{}
	var paramPointer interface{} = *queryLi
	paramLi = append(paramLi, &paramPointer)
	o.parse(classMethod, &paramLi)
}

func (o InterceptorManager) ParseAfterQueryData(classMethod string, items *[]interface{}) {
	paramLi := []*interface{}{}
	var paramPointer interface{} = *items
	paramLi = append(paramLi, &paramPointer)
	o.parse(classMethod, &paramLi)
}

func (o InterceptorManager) parse(classMethod string, param *[]*interface{}) {
	if classMethod != "" {
		exprContent := classMethod
		scriptStruct := strings.Split(exprContent, ".")[0]
		scriptStructMethod := strings.Split(exprContent, ".")[1]
		scriptType := GetInterceptorDict()[scriptStruct]
		inst := reflect.New(scriptType).Elem().Interface()
		instValue := reflect.ValueOf(inst)
		in := []reflect.Value{}
		for i, _ := range *param {
			in = append(in, reflect.ValueOf((*param)[i]))
		}
		instValue.MethodByName(scriptStructMethod).Call(in)
	}
}
