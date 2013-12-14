package model

import (
	"encoding/xml"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	. "com/papersns/component"
	"fmt"
	"strconv"
)

type ModelTemplateFactory struct {
}

func (o ModelTemplateFactory) GetDataSource(dataSourceModelId string) DataSource {
	dataSource, _ := o.getInstance(dataSourceModelId)
	return dataSource
}

func (o ModelTemplateFactory) GetInstanceByDS(dataSource DataSource) map[string]interface{} {
	_, bo := o.getInstance(dataSource.Id)
	return bo
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
	bo := o.getBo(dataSource)
	o.applyDefaultValueExpr(&dataSource, &bo)
	o.applyCalcValueExpr(&dataSource, &bo)
	o.applyRelationFieldValue(&dataSource, &bo)

	return dataSource, bo
}

// TODO, byTest
func (o ModelTemplateFactory) GetCopyInstance(dataSourceModelId string, srcBo map[string]interface{}) (DataSource, map[string]interface{}) {
	dataSource, bo := o.getInstance(dataSourceModelId)
	o.applyCopy(dataSource, &bo, srcBo)
	o.applyCalcValueExpr(&dataSource, &bo)
	o.applyRelationFieldValue(&dataSource, &bo)
	return dataSource, bo
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
			content := ""
			o.applyFieldGroupValueByString(*fieldGroup, &item, content)
		}
	})
	data, err := json.MarshalIndent(bo, "", "\t")
	if err != nil {
		panic(err)
	}
	println(string(data))
	return bo
}

// TODO,byTest
func (o ModelTemplateFactory) applyDefaultValueExpr(dataSource *DataSource, bo *map[string]interface{}) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	expressionParser := ExpressionParser{}
	boJsonData, err := json.Marshal(bo)
	if err != nil {
		panic(err)
	}
	boJson := string(boJsonData)
	modelIterator.IterateAllFieldBo(*dataSource, bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, result *interface{}){
		var content string = ""
		if fieldGroup.DefaultValueExpr.Content != "" {
			if fieldGroup.DefaultValueExpr.Mode == "" || fieldGroup.DefaultValueExpr.Mode == "text" {
				content = fieldGroup.DefaultValueExpr.Content
			} else if fieldGroup.DefaultValueExpr.Mode == "python" {
				dataJsonData, err := json.Marshal(data)
				if err != nil {
					panic(err)
				}
				dataJson := string(dataJsonData)
				content = expressionParser.ParseModel(boJson, dataJson, fieldGroup.DefaultValueExpr.Content)
			} else if fieldGroup.DefaultValueExpr.Mode == "golang" {
				content = fieldGroup.DefaultValueExpr.Content// TODO
			}
		}
		o.applyFieldGroupValueByString(fieldGroup, data, content)
	})
	data, err := json.MarshalIndent(bo, "", "\t")
	if err != nil {
		panic(err)
	}
	println(string(data))
}

// TODO,byTest
func (o ModelTemplateFactory) applyCalcValueExpr(dataSource *DataSource, bo *map[string]interface{}) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	expressionParser := ExpressionParser{}
	boJsonData, err := json.Marshal(bo)
	if err != nil {
		panic(err)
	}
	boJson := string(boJsonData)
	modelIterator.IterateAllFieldBo(*dataSource, bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, result *interface{}){
		var content string = ""
		if fieldGroup.CalcValueExpr.Content != "" {
			if fieldGroup.CalcValueExpr.Mode == "" || fieldGroup.CalcValueExpr.Mode == "text" {
				content = fieldGroup.CalcValueExpr.Content
			} else if fieldGroup.CalcValueExpr.Mode == "python" {
				dataJsonData, err := json.Marshal(data)
				if err != nil {
					panic(err)
				}
				dataJson := string(dataJsonData)
				content = expressionParser.ParseModel(boJson, dataJson, fieldGroup.CalcValueExpr.Content)
			} else if fieldGroup.CalcValueExpr.Mode == "golang" {
				content = fieldGroup.CalcValueExpr.Content// TODO
			}
		}
		o.applyFieldGroupValueByString(fieldGroup, data, content)
	})
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
	for i, _ := range *masterFixFieldLi {
		(*masterFixFieldLi)[i].Parent = dataSource.MasterData.FixField
	}
	for i, _ := range dataSource.MasterData.BizField.FieldLi {
		dataSource.MasterData.BizField.FieldLi[i].Parent = dataSource.MasterData.BizField
	}
	
	for i, _ := range dataSource.DetailDataLi {
		dataSource.DetailDataLi[i].FixField.Parent = dataSource.DetailDataLi[i]
		dataSource.DetailDataLi[i].BizField.Parent = dataSource.DetailDataLi[i]
	
		detailFixFieldLi := modelIterator.GetFixFieldLi(&dataSource.DetailDataLi[i].FixField)
		for j, _ := range *detailFixFieldLi {
			(*detailFixFieldLi)[j].Parent = dataSource.DetailDataLi[i].FixField
		}
		
		for j, _ := range dataSource.DetailDataLi[i].BizField.FieldLi {
			dataSource.DetailDataLi[i].BizField.FieldLi[j].Parent = dataSource.DetailDataLi[i].BizField
		}
	}
}

// TODO,byTest
func (o ModelTemplateFactory) applyRelationFieldValue(dataSource *DataSource, bo *map[string]interface{}) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateAllFieldBo(*dataSource, bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, result *interface{}){
		relationItem, found := o.ParseRelationExpr(fieldGroup, *bo, *data)
		if !found {
			panic("数据源:" + dataSource.Id + ",数据集:" + fieldGroup.GetDataSetId() + ",字段:" + fieldGroup.Id + ",配置的关联模型列表,不存在返回true的记录")
		}
		(*data)[fieldGroup.Id + "_ref"] = map[string]interface{}{
			"RelationExpr": true,
			"RelationModelId": relationItem.RelationModelId,
			"RelationDataSetId": relationItem.RelationDataSetId,
			"DisplayField": relationItem.DisplayField,
		}
	})
}

// TODO,byTest,
func (o ModelTemplateFactory) ParseRelationExpr(fieldGroup FieldGroup, bo map[string]interface{}, data map[string]interface{}) (RelationItem, bool) {
	fieldValue := fmt.Sprint(data[fieldGroup.Id])
	if fieldValue != "" {
		boJsonByte, err := json.Marshal(bo)
		if err != nil {
			panic(err)
		}
		boJson := string(boJsonByte)
		expressionParser := ExpressionParser{}
		for _, item := range fieldGroup.RelationDS.RelationItemLi {
			if item.RelationExpr.Content != "" {
				var content string
				if item.RelationExpr.Mode == "" || item.RelationExpr.Mode == "text" {
					content = item.RelationExpr.Content
				} else if item.RelationExpr.Mode == "python" {
					dataJsonData, err := json.Marshal(&data)
					if err != nil {
						panic(err)
					}
					dataJson := string(dataJsonData)
					content = expressionParser.ParseModel(boJson, dataJson, item.RelationExpr.Content)
				} else if item.RelationExpr.Mode == "golang" {
					content = item.RelationExpr.Content// TODO
				}
				if strings.ToLower(content) == "true" {
					return item, true
				}
			}
		}
	}

	return RelationItem{}, false
}

// TODO,byTest,
func (o ModelTemplateFactory) applyCopy(dataSource DataSource, destBo *map[string]interface{}, srcBo map[string]interface{}) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	//o.IterateDataBo(dataSource, bo, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, result *interface{}){
	modelIterator.IterateDataBo(dataSource, &srcBo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, result *interface{}){
		if !fieldGroupLi[0].IsMasterField() {
			if (*destBo)[fieldGroupLi[0].GetDataSetId()] == nil {
				(*destBo)[fieldGroupLi[0].GetDataSetId()] = []interface{}{}
			}
			dataSetLi := (*destBo)[fieldGroupLi[0].GetDataSetId()].([]interface{})
			copyData := map[string]interface{}{}
			for _, item := range fieldGroupLi {
				content := ""
				o.applyFieldGroupValueByString(item, &copyData, content)
			}
			dataSetLi = append(dataSetLi, copyData)
			(*destBo)[fieldGroupLi[0].GetDataSetId()] = dataSetLi
		}
	})
	result = ""
	modelIterator.IterateAllFieldTwoBo(&dataSource, destBo, srcBo, &result, func(fieldGroup *FieldGroup, destData *map[string]interface{}, srcData map[string]interface{}, result *interface{}) {
		if fieldGroup.AllowCopy == "" || fieldGroup.AllowCopy == "true" {
			(*destData)[fieldGroup.Id] = srcData[fieldGroup.Id]
		}
	})
}

// TODO,byTest,
func (o ModelTemplateFactory) IsDataDifferent(fieldGroupLi []FieldGroup, destData map[string]interface{}, srcData map[string]interface{}) bool {
	for _, item := range fieldGroupLi {
		if destData[item.Id] != srcData[item.Id] {
			return true
		}
	}
	return false
} 

// TODO,byTest,
func (o ModelTemplateFactory) ConvertDataType(dataSource DataSource, bo *map[string]interface{}) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateAllFieldBo(dataSource, bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, result *interface{}){
		content := fmt.Sprint((*data)[fieldGroup.Id])
		o.applyFieldGroupValueByString(fieldGroup, data, content)
	})
}

// TODO applyFieldGroupValue by default,
func (o ModelTemplateFactory) applyFieldGroupValueByString(fieldGroup FieldGroup, data *map[string]interface{}, content string) {
	stringArray := []string{"STRING", "REMARK"}
	for _, stringItem := range stringArray {
		if stringItem == fieldGroup.FieldDataType {
			(*data)[fieldGroup.Id] = content
			return
		}
	}
	intArray := []string{"SMALLINT", "INT", "LONGINT"}
	for _, intItem := range intArray {
		if intItem == fieldGroup.FieldDataType {
			if content == "" {
				(*data)[fieldGroup.Id] = 0
			} else {
				value, err := strconv.Atoi(content)
				if err != nil {
					panic(err)
				}
				(*data)[fieldGroup.Id] = value
			}
			return
		}
	}
	floatArray := []string{"FLOAT", "MONEY", "DECIMAL"}
	for _, floatItem := range floatArray {
		if floatItem == fieldGroup.FieldDataType {
			if content == "" {
				(*data)[fieldGroup.Id] = 0
			} else {
				value, err := strconv.ParseFloat(content, 32)
				if err != nil {
					panic(err)
				}
				(*data)[fieldGroup.Id] = float32(value)
			}
			return
		}
	}
	boolArray := []string{"BOOLEAN"}
	for _, boolItem := range boolArray {
		if boolItem == fieldGroup.FieldDataType {
			if content == "" {
				(*data)[fieldGroup.Id] = false
			} else {
				value, err := strconv.ParseBool(content)
				if err != nil {
					panic(err)
				}
				(*data)[fieldGroup.Id] = value
			}
			return
		}
	}
}
