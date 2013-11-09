package tree

import (
	"com/papersns/mongo"
	"fmt"
	"strconv"
	"sort"
	"labix.org/v2/mgo"
)

func GetInstance() TreeManager {
	return TreeManager{}
}

type TreeSort struct {
	objLi []map[string]interface{}
}

func (o TreeSort) Len() int {
	return len(o.objLi)
}

func (o TreeSort) Less(i, j int) bool {
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

func (o TreeSort) Swap(i, j int) {
	o.objLi[i], o.objLi[j] = o.objLi[j], o.objLi[i]
}


type TreeManager struct {
}

func (o TreeManager) GetTree(code string) map[string]interface{} {
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()
	
	return o.GetTreeBySession(db, code)
}

func (o TreeManager) GetTreeBySession(db *mgo.Database, code string) map[string]interface{} {
	if code == "SYSUSER_TREE" {
		return o.GetSysUserTree(db, code)
	}
	
	return map[string]interface{}{}
}

func (o TreeManager) GetSysUserTree(db *mgo.Database, code string) map[string]interface{} {
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
	o.sortTree(&result)
	return result
}

func (o TreeManager) sortTree(tree *map[string]interface{}) {
	items := (*tree)["items"]
	if items != nil {
		itemsLi := items.([]interface{})
		itemsMapLi := []map[string]interface{}{}
		for i,_ := range itemsLi {
			tmpObj := itemsLi[i].(map[string]interface{})
			itemsMapLi = append(itemsMapLi, tmpObj)
		}
		dSort := TreeSort{objLi: itemsMapLi}
		sort.Sort(dSort)
		(*tree)["items"] = itemsMapLi
		
		for i,_ := range itemsMapLi {
			o.sortTree(&(itemsMapLi[i]))
		}
	}
}
