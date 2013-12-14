package handler

import (
	. "com/papersns/model"
	"fmt"
	"strconv"
	"labix.org/v2/mgo"
)

type UsedCheck struct{}

func (o UsedCheck) Insert(db *mgo.Database, fieldGroupLi []FieldGroup, bo *map[string]interface{}, data *map[string]interface{}) {
	for _, fieldGroup := range fieldGroupLi {
		if fieldGroup.IsRelationField() {
			modelTemplateFactory := ModelTemplateFactory{}
			relationItem, found := modelTemplateFactory.ParseRelationExpr(fieldGroup, *bo, *data)
			if !found {
				panic("数据源:" + fieldGroup.GetDataSource().Id + ",数据集:" + fieldGroup.GetDataSetId() + ",字段:" + fieldGroup.Id + ",配置的关联模型列表,不存在返回true的记录")
			}
			referenceData := map[string]interface{}{
				"reference": o.GetSourceReferenceLi(db, fieldGroup, bo),
				"beReference": o.GetBeReferenceLi(db, fieldGroup, relationItem, data),
			}
			err := db.C("PubReferenceLog").Insert(referenceData)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (o UsedCheck) Update(db *mgo.Database, fieldGroupLi []FieldGroup, bo *map[string]interface{}, destData *map[string]interface{}, srcData map[string]interface{}) {
	if destData != nil && srcData == nil {
		o.Insert(db, fieldGroupLi, bo, destData)
	} else if destData == nil && srcData != nil {
		o.Delete(db, fieldGroupLi, srcData)
	} else if destData != nil && srcData != nil {
		// 分析字段,如果字段都相等,不过帐,
		modelTemplateFactory := ModelTemplateFactory{}
		if modelTemplateFactory.IsDataDifferent(fieldGroupLi, *destData, srcData) {
			o.Delete(db, fieldGroupLi, srcData)
			o.Insert(db, fieldGroupLi, bo, destData)
		}
	}
}

func (o UsedCheck) Delete(db *mgo.Database, fieldGroupLi []FieldGroup, data map[string]interface{}) {
	dataSource := fieldGroupLi[0].GetDataSource()
	id, err := strconv.Atoi(fmt.Sprint(data["id"]))
	if err != nil {
		panic(err)
	}
	referenceQuery := []interface{}{
		dataSource.Id,
		fieldGroupLi[0].GetDataSetId(),
		"id",
		id,
	}
	db.C("PubReferenceLog").Remove(map[string]interface{}{
		"reference": referenceQuery,
	})
}



//reference:[[dataSource, dataSet, fieldName, id], [dataSource, dataSet, fieldName, id]]
//beReference:[[dataSource, dataSet, fieldName, id], [dataSource, dataSet, fieldName, id]]
func (o UsedCheck) GetSourceReferenceLi(db *mgo.Database, fieldGroup FieldGroup, bo *map[string]interface{}) []interface{} {
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
func (o UsedCheck) GetBeReferenceLi(db *mgo.Database, fieldGroup FieldGroup, relationItem RelationItem, data *map[string]interface{}) []interface{} {
	sourceLi := []interface{}{}
	relationId, err := strconv.Atoi(fmt.Sprint((*data)[fieldGroup.Id]))
	if err != nil {
		panic(err)
	}
	if relationItem.RelationDataSetId == "A" {
		sourceLi = append(sourceLi, []interface{}{relationItem.RelationModelId, "A", "id", relationId})
		return sourceLi
	} else {
		refData := map[string]interface{}{}
		query := map[string]interface{}{
			relationItem.RelationDataSetId + ".id": relationId,
		}
		//{"B.id": 2}
		err := db.C(relationItem.RelationModelId).Find(query).One(&refData)
		if err != nil {
			panic(err)
		}
		masterData := refData["A"].(map[string]interface{})
		masterDataId, err := strconv.Atoi(fmt.Sprint(masterData["id"]))
		if err != nil {
			panic(err)
		}
		sourceLi = append(sourceLi, []interface{}{relationItem.RelationModelId, "A", "id", masterDataId})
	}
	
	sourceLi = append(sourceLi, []interface{}{relationItem.RelationModelId, relationItem.RelationDataSetId, "id", relationId})
	return sourceLi
}
