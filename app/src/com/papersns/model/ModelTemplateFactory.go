package model

import (
	"encoding/xml"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
)

type ModelTemplateFactory struct {
}

func (o ModelTemplateFactory) GetInstance(dataSourceModelId string) (DataSource, map[string]interface{}) {
	return o.getInstance(dataSourceModelId)
}

func (o ModelTemplateFactory) getInstance(dataSourceModelId string) (DataSource, map[string]interface{}) {
	file, err := os.Open("/home/hongjinqiu/goworkspace/src/finance/app/src/com/papersns/model/xml/pc_ds_sysuser.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	dataSource := DataSource{}
	err = xml.Unmarshal(data, &dataSource)
	if err != nil {
		panic(err)
	}

	o.applyFieldExtend(&dataSource)
	o.applyReverseRelation(&dataSource)

	o.applyDefaultValueExpr(&dataSource)
	//	o.applyCalcValueExpr(&dataSource)
	//	o.applyRelationFieldValue(&dataSource)
	bo := o.getBO(&dataSource)

	return dataSource, bo
}

func (o ModelTemplateFactory) GetCopyInstance(dataSourceModelId string, srcBo map[string]interface{}) (DataSource, map[string]interface{}) {
	dataSource, bo := o.getInstance(dataSourceModelId)
	o.applyCopy(dataSource, &bo, &srcBo)
	return dataSource, srcBo
}

func (o ModelTemplateFactory) extendFieldPoolField(fieldGroup *FieldGroup, fieldGroupLi *[]FieldGroup) {
	outFieldGroup := fieldGroup
	if outFieldGroup.Extends != "" {
		outFieldGroupElem := reflect.ValueOf(outFieldGroup).Elem()
		for j, _ := range *fieldGroupLi {
			innerFieldGroup := (*fieldGroupLi)[j]
			innerFieldGroupReflect := reflect.ValueOf(innerFieldGroup)
			if outFieldGroup.Extends == innerFieldGroup.Id {
				for k := 0; k < outFieldGroupElem.Type().NumField(); k++ {
					if outFieldGroupElem.Field(k).Kind() == reflect.String {
						outValue := outFieldGroupElem.Field(k).Interface().(string)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(string)
						if outValue == "" && innerValue != "" {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Bool {
						// 不处理
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Int {
						outValue := outFieldGroupElem.Field(k).Interface().(int)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(int)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Int8 {
						outValue := outFieldGroupElem.Field(k).Interface().(int8)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(int8)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Int16 {
						outValue := outFieldGroupElem.Field(k).Interface().(int16)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(int16)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Int32 {
						outValue := outFieldGroupElem.Field(k).Interface().(int32)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(int32)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Int64 {
						outValue := outFieldGroupElem.Field(k).Interface().(int64)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(int64)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Uint {
						outValue := outFieldGroupElem.Field(k).Interface().(uint)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(uint)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Uint8 {
						outValue := outFieldGroupElem.Field(k).Interface().(uint8)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(uint8)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Uint16 {
						outValue := outFieldGroupElem.Field(k).Interface().(uint16)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(uint16)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Uint32 {
						outValue := outFieldGroupElem.Field(k).Interface().(uint32)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(uint32)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Uint64 {
						outValue := outFieldGroupElem.Field(k).Interface().(uint64)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(uint64)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Uintptr {
						outValue := outFieldGroupElem.Field(k).Interface().(uintptr)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(uintptr)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Float32 {
						outValue := outFieldGroupElem.Field(k).Interface().(float32)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(float32)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Float64 {
						outValue := outFieldGroupElem.Field(k).Interface().(float64)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(float64)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Complex64 {
						outValue := outFieldGroupElem.Field(k).Interface().(complex64)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(complex64)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					} else if outFieldGroupElem.Field(k).Kind() == reflect.Complex128 {
						outValue := outFieldGroupElem.Field(k).Interface().(complex128)
						innerValue := innerFieldGroupReflect.Field(k).Interface().(complex128)
						if outValue == 0 && innerValue != 0 {
							outFieldGroupElem.Field(k).Set(innerFieldGroupReflect.Field(k))
						}
					}

				}
			}
		}
	}
}

func (o ModelTemplateFactory) getPoolFields() Fields {
	file, err := os.Open("/home/hongjinqiu/goworkspace/src/finance/app/src/com/papersns/model/xml/fieldpool.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fields := Fields{}
	err = xml.Unmarshal(data, &fields)

	fieldGroupLi := []FieldGroup{}
	for i, _ := range fields.FieldLi {
		fieldGroupLi = append(fieldGroupLi, fields.FieldLi[i].FieldGroup)
	}
	for i, _ := range fields.FieldLi {
		o.extendFieldPoolField(&fields.FieldLi[i].FieldGroup, &fieldGroupLi)
	}

	for i, _ := range fields.FieldLi {
		outFieldRelationDS := &fields.FieldLi[i].FieldGroup.RelationDS
		if outFieldRelationDS.RelationItemLi == nil {
			for j, _ := range fields.FieldLi {
				if fields.FieldLi[i].FieldGroup.Extends == fields.FieldLi[j].FieldGroup.Id {
					innerFieldRelationDS := &fields.FieldLi[j].FieldGroup.RelationDS
					if innerFieldRelationDS.RelationItemLi != nil {
						outFieldRelationDS.RelationItemLi = innerFieldRelationDS.RelationItemLi
					}
				}
			}
		}
	}
	return fields
}

func (o ModelTemplateFactory) applyFieldExtend(dataSource *DataSource) {
	modelIterator := ModelIterator{}
	var result interface{} = ""

	fields := o.getPoolFields()
	fieldGroupLi := []FieldGroup{}
	for i, _ := range fields.FieldLi {
		fieldGroupLi = append(fieldGroupLi, fields.FieldLi[i].FieldGroup)
	}

	modelIterator.IterateAllField(dataSource, &result, func(fieldGroup *FieldGroup, result *interface{}) {
		o.extendFieldPoolField(fieldGroup, &fieldGroupLi)
	})
}

func (o ModelTemplateFactory) getBo(dataSource DataSource) map[string]interface{} {
	bo := map[string]interface{}{
		"A": map[string]interface{}{},
	}
	for _, item := range dataSource.DetailDataLi {
		bo[item.Id] = []interface{}{}
	}
	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateAllField(&dataSource, &result, func(fieldGroup *FieldGroup, result *interface{}) {
		if fieldGroup.IsMasterField() {
			item := bo["A"].(map[string]interface{})
			stringArray := []string{"STRING", "REMARK"}
			for _, stringItem := range stringArray {
				if stringItem == fieldGroup.FieldDataType {
					item[fieldGroup.Id] = ""
					break
				}
			}
			intArray := []string{"SMALLINT", "INT", "LONGINT", "FLOAT", "MONEY", "DECIMAL"}
			for _, intItem := range intArray {
				if intItem == fieldGroup.FieldDataType {
					item[fieldGroup.Id] = 0
					break
				}
			}
			boolArray := []string{"BOOLEAN"}
			for _, boolItem := range boolArray {
				if boolItem == fieldGroup.FieldDataType {
					item[fieldGroup.Id] = false
					break
				}
			}
		}
	})
	data, err := json.MarshalIndent(bo, "", "\t")
	if err != nil {
		panic(err)
	}
	println(string(data))
	return bo
}

// TODO
func (o ModelTemplateFactory) applyDefaultValueExpr(dataSource *DataSource) {
	bo := o.getBo(*dataSource)
	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateAllFieldBo(dataSource, &bo, &result, func(fieldGroup *FieldGroup, data *map[string]interface{}, result *interface{}){
		
	})
	data, err := json.MarshalIndent(bo, "", "\t")
	if err != nil {
		panic(err)
	}
	println(string(data))
}

// TODO
func (o ModelTemplateFactory) applyCalcValueExpr(dataSource *DataSource) {

}

/**
 * 建立父子双向关联
 */
func (o ModelTemplateFactory) applyReverseRelation(dataSource *DataSource) {
	dataSource.MasterData.Parent = (*dataSource)
	for i, _ := range dataSource.DetailDataLi {
		dataSource.DetailDataLi[i].Parent = (*dataSource)
	}
	dataSource.MasterData.FixField.Parent = dataSource.MasterData
	dataSource.MasterData.BizField.Parent = dataSource.MasterData

	modelIterator := ModelIterator{}
	masterFixFieldLi := modelIterator.GetFixFieldLi(&dataSource.MasterData.FixField)
	for i, _ := range masterFixFieldLi {
		masterFixFieldLi[i].Parent = dataSource.MasterData.FixField
	}
	for i, _ := range dataSource.MasterData.BizField.FieldLi {
		dataSource.MasterData.BizField.FieldLi[i].Parent = dataSource.MasterData.BizField
	}
	
	for i, _ := range dataSource.DetailDataLi {
		dataSource.DetailDataLi[i].FixField.Parent = dataSource.DetailDataLi[i]
		dataSource.DetailDataLi[i].BizField.Parent = dataSource.DetailDataLi[i]
	
		detailFixFieldLi := modelIterator.GetFixFieldLi(&dataSource.DetailDataLi[i].FixField)
		for j, _ := range detailFixFieldLi {
			detailFixFieldLi[j].Parent = dataSource.DetailDataLi[i].FixField
		}
		
		for j, _ := range dataSource.DetailDataLi[i].BizField.FieldLi {
			dataSource.DetailDataLi[i].BizField.FieldLi[j].Parent = dataSource.DetailDataLi[i].BizField
		}
	}
}

// TODO
func (o ModelTemplateFactory) applyRelationFieldValue(dataSource *DataSource) {

}

// TODO
func (o ModelTemplateFactory) getBO(dataSource *DataSource) map[string]interface{} {
	return map[string]interface{}{}
}

// TODO
func (o ModelTemplateFactory) applyCopy(dataSource DataSource, bo *map[string]interface{}, srcBo *map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{}
}
