package handler

import (
	. "com/papersns/model"
	"fmt"
	"strconv"
	"labix.org/v2/mgo"
)

type UsedCheck struct{}

func (o UsedCheck) Execute(db *mgo.Database, fieldGroup *FieldGroup, bo *map[string]interface{}, dataType int) {
//	if data[fieldGroup.Id] == nil {
//		return
//	}
	if fieldGroup.IsRelationField() {
		relationItem, found := fieldGroup.GetRelationItem()
		fmt.Println(relationItem)
		if found {
			if dataType == ADD || dataType == AFTER_UPDATE {// add
				
			} else {// delete
				
			}
		}
	}
}

//reference:[[dataSource, dataSet, fieldName, id], [dataSource, dataSet, fieldName, id]]
//beReference:[[dataSource, dataSet, fieldName, id], [dataSource, dataSet, fieldName, id]]
func (o UsedCheck) GetSourceReferenceLi(db *mgo.Database, fieldGroup *FieldGroup, bo *map[string]interface{}) []interface{} {
	masterData := (*bo)["A"].(map[string]interface{})
	sourceLi := []interface{}{}
	
	srcDataSourceId := fieldGroup.GetDataSource().Id
	srcDataSetId := fieldGroup.GetDataSetId()
	srcFieldName := "id"
	iId := fmt.Sprint(masterData["id"])
	id, err := strconv.Atoi(iId)
	if err != nil {
		panic(err)
	}
	refLi := []interface{}{srcDataSourceId, srcDataSetId, srcFieldName, id}
	sourceLi = append(sourceLi, refLi)
	if fieldGroup.IsMasterField() {
		srcDataSourceId = fieldGroup.GetDataSource().Id
		srcDataSetId = fieldGroup.GetDataSetId()
		srcFieldName = fieldGroup.Id
		iId := fmt.Sprint(masterData["id"])
		id, err := strconv.Atoi(iId)
		if err != nil {
			panic(err)
		}
		refLi2 := []interface{}{srcDataSourceId, srcDataSetId, srcFieldName, id}
		sourceLi = append(sourceLi, refLi2)
	} else {
		srcDataSourceId = fieldGroup.GetDataSource().Id
		srcDataSetId = fieldGroup.GetDataSetId()
		dataSetData := (*bo)[srcDataSetId].(map[string]interface{})
		srcFieldName = "id"
		iId := fmt.Sprint(dataSetData["id"])
		id, err := strconv.Atoi(iId)
		if err != nil {
			panic(err)
		}
		refLi2 := []interface{}{srcDataSourceId, srcDataSetId, srcFieldName, id}
		sourceLi = append(sourceLi, refLi2)
		
		srcDataSourceId = fieldGroup.GetDataSource().Id
		srcDataSetId = fieldGroup.GetDataSetId()
		dataSetData = (*bo)[srcDataSetId].(map[string]interface{})
		srcFieldName = fieldGroup.Id
		iId = fmt.Sprint(dataSetData["id"])
		id, err = strconv.Atoi(iId)
		if err != nil {
			panic(err)
		}
		refLi3 := []interface{}{srcDataSourceId, srcDataSetId, srcFieldName, id}
		sourceLi = append(sourceLi, refLi3)
	}
	return sourceLi
}

//reference:[[dataSource, dataSet, fieldName, id], [dataSource, dataSet, fieldName, id]]
//beReference:[[dataSource, dataSet, fieldName, id], [dataSource, dataSet, fieldName, id]]
func (o UsedCheck) GetBeReferenceLi(db *mgo.Database, fieldGroup *FieldGroup, bo *map[string]interface{}) []interface{} {
	_, found := fieldGroup.GetRelationItem()
	if found {
		
	}
	return []interface{} {}
}
