package dictionary

import (
	"com/papersns/mongo"
	"fmt"
	"strconv"
	"sort"
	"labix.org/v2/mgo"
)

func GetProgramDictionaryInstance() ProgramDictionaryManager {
	return ProgramDictionaryManager{}
}

type ProgramDictionarySort struct {
	objLi []map[string]interface{}
}

func (o ProgramDictionarySort) Len() int {
	return len(o.objLi)
}

func (o ProgramDictionarySort) Less(i, j int) bool {
	orderI := o.objLi[i]["order"]
	if orderI == nil {
		return false
	}
	orderJ := o.objLi[j]["order"]
	if orderJ == nil {
		return false
	}

	order1, err := strconv.Atoi(fmt.Sprint(orderI))
	if err != nil {
		panic(err)
	}
	
	order2, err := strconv.Atoi(fmt.Sprint(orderJ))
	if err != nil {
		panic(err)
	}
	
	return order1 <= order2
}

func (o ProgramDictionarySort) Swap(i, j int) {
	o.objLi[i], o.objLi[j] = o.objLi[j], o.objLi[i]
}


type ProgramDictionaryManager struct {
}

func (o ProgramDictionaryManager) GetProgramDictionary(code string) map[string]interface{} {
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()
	
	return o.GetProgramDictionaryBySession(db, code)
}

func (o ProgramDictionaryManager) GetProgramDictionaryBySession(db *mgo.Database, code string) map[string]interface{} {
	if code == "SYSUSER_TREE" {
		return o.GetSysUserProgramDictionary(db, code)
	}
	
	return nil
}

func (o ProgramDictionaryManager) GetSysUserProgramDictionary(db *mgo.Database, code string) map[string]interface{} {
	collection := "SysUser"
	c := db.C(collection)
	
	queryMap := map[string]interface{}{}
	
	sysUserResult := []map[string]interface{}{}
	err := c.Find(queryMap).Limit(10).All(&sysUserResult)
	if err != nil {
		panic(err)
	}
	
	result := map[string]interface{}{}
	result["code"] = code
	//items := []map[string]interface{}{}
	items := []interface{}{}
	for idx, item := range sysUserResult {
		items = append(items, map[string]interface{}{
			"code": item["_id"],
			"name": item["nick"],
			"order": idx,
		})
	}
	result["items"] = items
	
	// 排序
	o.sortProgramDictionary(&result)
	return result
}

func (o ProgramDictionaryManager) sortProgramDictionary(programDictionary *map[string]interface{}) {
	items := (*programDictionary)["items"]
	if items != nil {
		itemsLi := items.([]interface{})
		itemsMapLi := []map[string]interface{}{}
		for i,_ := range itemsLi {
			tmpObj := itemsLi[i].(map[string]interface{})
			itemsMapLi = append(itemsMapLi, tmpObj)
		}
		dSort := ProgramDictionarySort{objLi: itemsMapLi}
		sort.Sort(dSort)
		(*programDictionary)["items"] = itemsMapLi
		
		for i,_ := range itemsMapLi {
			o.sortProgramDictionary(&(itemsMapLi[i]))
		}
	}
}
