package model

import "github.com/robfig/revel"

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	. "com/papersns/script"
)

var rwlock sync.RWMutex = sync.RWMutex{}
var gDataSourceDict map[string]DataSourceInfo = map[string]DataSourceInfo{}

type DataSourceInfo struct {
	Path       string
	DataSource DataSource
}

type ModelTemplateFactory struct {
}

// TODO, byTest
func (o ModelTemplateFactory) GetDataSourceInfoLi() []DataSourceInfo {
	dataSourceInfo := []DataSourceInfo{}
	if len(gDataSourceDict) == 0 {
		o.loadDataSource()
	}
	
	rwlock.RLock()
	defer rwlock.RUnlock()
	
	for _, item := range gDataSourceDict {
		dataSourceInfo = append(dataSourceInfo, item)
	}
	return dataSourceInfo
}

// TODO, byTest
func (o ModelTemplateFactory) RefretorDataSourceInfo() []DataSourceInfo {
	o.clearDataSource()
	o.loadDataSource()
	dataSourceInfo := []DataSourceInfo{}
	for _, item := range gDataSourceDict {
		dataSourceInfo = append(dataSourceInfo, item)
	}
	return dataSourceInfo
}

func (o ModelTemplateFactory) GetDataSource(id string) DataSource {
	return o.GetDataSourceInfo(id).DataSource
}

// TODO, byTest
func (o ModelTemplateFactory) GetDataSourceInfo(id string) DataSourceInfo {
	if revel.Config.StringDefault("mode.dev", "true") == "true" {
		dataSourceInfo, found := o.findDataSourceInfo(id)
		if found {
			dataSourceInfo, err := o.loadSingleDataSourceWithLock(dataSourceInfo.Path)
			if err != nil {
				panic(err)
			}
			if dataSourceInfo.DataSource.Id == id {
				return dataSourceInfo
			}
		}
		o.clearDataSource()
		o.loadDataSource()
		dataSourceInfo, found = o.findDataSourceInfo(id)
		if found {
			return dataSourceInfo
		}
		panic(id + " not exists in DataSource list")
	}

	if len(gDataSourceDict) == 0 {
		o.loadDataSource()
	}
	dataSourceInfo, found := o.findDataSourceInfo(id)
	if found {
		return dataSourceInfo
	}
	panic(id + " not exists in DataSource list")
}

// TODO bytest,
func (o ModelTemplateFactory) findDataSourceInfo(id string) (DataSourceInfo, bool) {
	rwlock.RLock()
	defer rwlock.RUnlock()

	for _, item := range gDataSourceDict {
		if item.DataSource.Id == id {
			return item, true
		}
	}
	return DataSourceInfo{}, false
}

// TODO, byTest
func (o ModelTemplateFactory) clearDataSource() {
	rwlock.Lock()
	defer rwlock.Unlock()

	gDataSourceDict = map[string]DataSourceInfo{}
}

// TODO, byTest
func (o ModelTemplateFactory) loadDataSource() {
	rwlock.Lock()
	defer rwlock.Unlock()

	path := revel.Config.StringDefault("DATA_SOURCE_PATH", "")
	if path != "" {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Index(path, "ds_") > -1 && strings.Index(path, ".xml") > -1 && !info.IsDir() {
				_, err = o.loadSingleDataSource(path)
				if err != nil {
					return err
				}
			}

			return nil
		})
	}
}

// TODO, byTest
func (o ModelTemplateFactory) loadSingleDataSourceWithLock(path string) (DataSourceInfo, error) {
	rwlock.Lock()
	defer rwlock.Unlock()

	return o.loadSingleDataSource(path)
}

// TODO, byTest
func (o ModelTemplateFactory) loadSingleDataSource(path string) (DataSourceInfo, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return DataSourceInfo{}, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return DataSourceInfo{}, err
	}

	dataSource := DataSource{}
	err = xml.Unmarshal(data, &dataSource)
	if err != nil {
		return DataSourceInfo{}, err
	}

	o.applyFieldExtend(&dataSource)
	o.applyReverseRelation(&dataSource)

	dataSourceInfo := DataSourceInfo{
		Path:       path,
		DataSource: dataSource,
	}
	gDataSourceDict[dataSource.Id] = dataSourceInfo
	return dataSourceInfo, nil
}

func (o ModelTemplateFactory) GetInstanceByDS(dataSource DataSource) map[string]interface{} {
	bo := o.getBo(dataSource)
	o.applyDefaultValueExpr(dataSource, &bo)
	o.applyCalcValueExpr(dataSource, &bo)
	o.applyRelationFieldValue(dataSource, &bo)

	return bo
}

func (o ModelTemplateFactory) GetInstance(dataSourceModelId string) (DataSource, map[string]interface{}) {
	dataSource := o.GetDataSource(dataSourceModelId)

	bo := o.getBo(dataSource)
	o.applyDefaultValueExpr(dataSource, &bo)
	o.applyCalcValueExpr(dataSource, &bo)
	o.applyRelationFieldValue(dataSource, &bo)

	return dataSource, bo
}

func (o ModelTemplateFactory) GetCopyInstance(dataSourceModelId string, srcBo map[string]interface{}) (DataSource, map[string]interface{}) {
	dataSource, bo := o.GetInstance(dataSourceModelId)
	o.applyCopy(dataSource, &bo, srcBo)
	o.applyCalcValueExpr(dataSource, &bo)
	o.applyRelationFieldValue(dataSource, &bo)
	return dataSource, bo
}

func (o ModelTemplateFactory) extendFieldPoolField(fieldGroup *FieldGroup, fieldGroupLi []FieldGroup) {
	outFieldGroup := fieldGroup
	if outFieldGroup.Extends != "" {
		outFieldGroupElem := reflect.ValueOf(outFieldGroup).Elem()
		for j, _ := range fieldGroupLi {
			innerFieldGroup := fieldGroupLi[j]
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
	fieldPoolPath := revel.Config.StringDefault("FIELD_POOL_PATH", "")
	file, err := os.Open(fieldPoolPath)
	defer file.Close()
	if err != nil {
		panic(err)
	}

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
		o.extendFieldPoolField(&fields.FieldLi[i].FieldGroup, fieldGroupLi)
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
		o.extendFieldPoolField(fieldGroup, fieldGroupLi)
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
	return bo
}

func (o ModelTemplateFactory) applyDefaultValueExpr(dataSource DataSource, bo *map[string]interface{}) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	expressionParser := ExpressionParser{}
	boJsonData, err := json.Marshal(bo)
	if err != nil {
		panic(err)
	}
	boJson := string(boJsonData)
	modelIterator.IterateAllFieldBo(dataSource, bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
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
				exprContent := fieldGroup.DefaultValueExpr.Content
				content = expressionParser.ParseGolang(*bo, *data, exprContent)
			}
		}
		o.applyFieldGroupValueByString(fieldGroup, data, content)
	})
}

func (o ModelTemplateFactory) applyCalcValueExpr(dataSource DataSource, bo *map[string]interface{}) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	expressionParser := ExpressionParser{}
	boJsonData, err := json.Marshal(bo)
	if err != nil {
		panic(err)
	}
	boJson := string(boJsonData)
	modelIterator.IterateAllFieldBo(dataSource, bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
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
				exprContent := fieldGroup.CalcValueExpr.Content
				content = expressionParser.ParseGolang(*bo, *data, exprContent)
			}
			o.applyFieldGroupValueByString(fieldGroup, data, content)
		}
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

/**
 * 清除父子双向关联,双向关联会引起json.Marshal死循环
 */
func (o ModelTemplateFactory) ClearReverseRelation(dataSource *DataSource) {
	dataSource.MasterData.Parent = nil
	for i, _ := range dataSource.DetailDataLi {
		dataSource.DetailDataLi[i].Parent = nil
	}
	dataSource.MasterData.FixField.Parent = nil
	dataSource.MasterData.BizField.Parent = nil

	modelIterator := ModelIterator{}
	masterFixFieldLi := modelIterator.GetFixFieldLi(&dataSource.MasterData.FixField)
	for i, _ := range *masterFixFieldLi {
		(*masterFixFieldLi)[i].Parent = nil
	}
	for i, _ := range dataSource.MasterData.BizField.FieldLi {
		dataSource.MasterData.BizField.FieldLi[i].Parent = nil
	}

	for i, _ := range dataSource.DetailDataLi {
		dataSource.DetailDataLi[i].FixField.Parent = nil
		dataSource.DetailDataLi[i].BizField.Parent = nil

		detailFixFieldLi := modelIterator.GetFixFieldLi(&dataSource.DetailDataLi[i].FixField)
		for j, _ := range *detailFixFieldLi {
			(*detailFixFieldLi)[j].Parent = nil
		}

		for j, _ := range dataSource.DetailDataLi[i].BizField.FieldLi {
			dataSource.DetailDataLi[i].BizField.FieldLi[j].Parent = nil
		}
	}
}

func (o ModelTemplateFactory) applyRelationFieldValue(dataSource DataSource, bo *map[string]interface{}) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateAllFieldBo(dataSource, bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		if fieldGroup.IsRelationField() {
			relationItem, found := o.ParseRelationExpr(fieldGroup, *bo, *data)
			if found {
				(*data)[fieldGroup.Id+"_ref"] = map[string]interface{}{
					"Id":                relationItem.Id,
					"RelationExpr":      true,
					"RelationModelId":   relationItem.RelationModelId,
					"RelationDataSetId": relationItem.RelationDataSetId,
					"DisplayField":      relationItem.DisplayField,
				}
			}
		}
	})
}

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
					exprContent := item.RelationExpr.Content
					content = expressionParser.ParseGolang(bo, data, exprContent)
				}
				if strings.ToLower(content) == "true" {
					return item, true
				}
			}
		}
	}

	return RelationItem{}, false
}

func (o ModelTemplateFactory) applyCopy(dataSource DataSource, destBo *map[string]interface{}, srcBo map[string]interface{}) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateDataBo(dataSource, &srcBo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		if !fieldGroupLi[0].IsMasterField() {
			if (*destBo)[fieldGroupLi[0].GetDataSetId()] == nil {
				(*destBo)[fieldGroupLi[0].GetDataSetId()] = []interface{}{}
			}
			dataSetLi := (*destBo)[fieldGroupLi[0].GetDataSetId()].([]interface{})
			copyData := map[string]interface{}{}
			dataSetLi = append(dataSetLi, copyData)
			(*destBo)[fieldGroupLi[0].GetDataSetId()] = dataSetLi
		}
	})
	o.applyDefaultValueExpr(dataSource, destBo)
	result = ""
	modelIterator.IterateAllFieldTwoBo(&dataSource, destBo, srcBo, &result, func(fieldGroup *FieldGroup, destData *map[string]interface{}, srcData map[string]interface{}, result *interface{}) {
		if fieldGroup.AllowCopy == "" || fieldGroup.AllowCopy == "true" {
			if srcData[fieldGroup.Id] != nil {
				(*destData)[fieldGroup.Id] = srcData[fieldGroup.Id]
			}
		}
	})
}

func (o ModelTemplateFactory) IsDataDifferent(fieldGroupLi []FieldGroup, destData map[string]interface{}, srcData map[string]interface{}) bool {
	for _, item := range fieldGroupLi {
		destStrData := fmt.Sprint(destData[item.Id])
		srcStrData := fmt.Sprint(srcData[item.Id])
		if destStrData != srcStrData {
			if item.Id != "modifyTime" {
				return true
			}
		}
	}
	return false
}

func (o ModelTemplateFactory) ConvertDataType(dataSource DataSource, bo *map[string]interface{}) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateAllFieldBo(dataSource, bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		content := ""
		if (*data)[fieldGroup.Id] != nil {
			content = fmt.Sprint((*data)[fieldGroup.Id])
		}
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
				value, err := strconv.ParseInt(content, 10, 64)
				//value, err := strconv.Atoi(content)
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

func (o ModelTemplateFactory) GetRelationLi(sId int, dataSource DataSource, bo map[string]interface{}) []map[string]interface{} {
	if bo["_id"] != nil {
		id := fmt.Sprint(bo["_id"])
		if id == "" || id == "0" {
			return []map[string]interface{}{}
		}
	}
	li := []map[string]interface{}{}
	
	modelIterator := ModelIterator{}
	var result interface{} = ""

	modelIterator.IterateAllFieldBo(dataSource, &bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}){
		relationItem, found := fieldGroup.GetRelationItem(bo, *data)
		if found {
			if (*data)[fieldGroup.Id] != nil {
				strRelationId := fmt.Sprint((*data)[fieldGroup.Id])
				if strRelationId != "" && strRelationId != "0" {
					relationId, err := strconv.Atoi(strRelationId)
					if err != nil {
						panic(err)
					}
					isContain := false
					for _, item := range li {
						tmpRelationId := fmt.Sprint(item["relationId"])
						tmpSelectorId := fmt.Sprint(item["selectorId"])
						if tmpRelationId == strRelationId && tmpSelectorId == relationItem.Id {
							isContain = true
							break
						}
					}
					if !isContain {
						li = append(li, map[string]interface{}{
							"relationId": relationId,
							"selectorId": relationItem.Id,
						})
					}
				}
			}
			
		}
	})
	
	return li
}

