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
	if code == "ACCOUNTING_YEAR_START_TREE" {// 会计期年度开始
		return o.GetAccountingYearStartProgramDictionary(db, code)
	}
	if code == "ACCOUNTING_YEAR_END_TREE" {// 会计期年度结束
		return o.GetAccountingYearEndProgramDictionary(db, code)
	}
	if code == "ACCOUNTING_PERIOD_START_TREE" {// 会计期期间开始
		return o.GetAccountingPeriodStartProgramDictionary(db, code)
	}
	if code == "ACCOUNTING_PERIOD_END_TREE" {// 会计期期间结束
		return o.GetAccountingPeriodEndProgramDictionary(db, code)
	}
	
	return nil
}

func (o ProgramDictionaryManager) GetAccountingYearStartProgramDictionary(db *mgo.Database, code string) map[string]interface{} {
	collection := "AccountingPeriod"
	c := db.C(collection)
	
	queryMap := map[string]interface{}{}
	
	itemResult := []map[string]interface{}{}
	err := c.Find(queryMap).All(&itemResult)
	if err != nil {
		panic(err)
	}
	
	result := map[string]interface{}{}
	result["code"] = code
	items := []interface{}{}
	for _, item := range itemResult {
		master := item["A"].(map[string]interface{})
		items = append(items, map[string]interface{}{
			"code": master["accountingYear"],
			"name": master["accountingYear"],
			"order": master["accountingYear"],
		})
	}
	result["items"] = items
	
	// 排序
	o.sortProgramDictionary(&result)
	return result
}

func (o ProgramDictionaryManager) GetAccountingYearEndProgramDictionary(db *mgo.Database, code string) map[string]interface{} {
	return o.GetAccountingYearStartProgramDictionary(db, code)
}

func (o ProgramDictionaryManager) GetAccountingPeriodStartProgramDictionary(db *mgo.Database, code string) map[string]interface{} {
	collection := "AccountingPeriod"
	c := db.C(collection)
	
	queryMap := map[string]interface{}{}
	
	itemResult := []map[string]interface{}{}
	err := c.Find(queryMap).All(&itemResult)
	if err != nil {
		panic(err)
	}
	
	result := map[string]interface{}{}
	result["code"] = code
	items := []interface{}{}
	for _, item := range itemResult {
		detailLi := item["B"].([]interface{})
		for _, detail := range detailLi {
			detailMap := detail.(map[string]interface{})
			isIn := false
			for _, dictItem := range items {
				dictItemMap := dictItem.(map[string]interface{})
				if fmt.Sprint(dictItemMap["code"]) == fmt.Sprint(detailMap["sequenceNo"]) {
					isIn = true
					break
				}
			}
			if !isIn {
				items = append(items, map[string]interface{}{
					"code": detailMap["sequenceNo"],
					"name": detailMap["sequenceNo"],
					"order": detailMap["sequenceNo"],
				})
			}
		}
	}
	result["items"] = items
	
	// 排序
	o.sortProgramDictionary(&result)
	return result
}

func (o ProgramDictionaryManager) GetAccountingPeriodEndProgramDictionary(db *mgo.Database, code string) map[string]interface{} {
	return o.GetAccountingPeriodStartProgramDictionary(db, code)
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
