package script

import (
	"sync"
	"reflect"
	"strconv"
	"fmt"
)

var rwlock sync.RWMutex = sync.RWMutex{}
var scriptDict map[string]reflect.Type = map[string]reflect.Type{}

func init() {
	rwlock.Lock()
	defer rwlock.Unlock()
	scriptDict[reflect.TypeOf(SysUserScript{}).Name()] = reflect.TypeOf(SysUserScript{})
}

func GetScriptDict() map[string]reflect.Type {
	rwlock.RLock()
	defer rwlock.RUnlock()
	return scriptDict
}

type SysUserScript struct {}

func (o SysUserScript) GetIntTest(bo map[string]interface{}, data map[string]interface{}) string {
	masterData := bo["A"].(map[string]interface{})
	attachCount, err := strconv.Atoi(fmt.Sprint(masterData["attachCount"]))
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(attachCount * 20 + 30)
}
