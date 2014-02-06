package component

import (
	. "com/papersns/model"
	"reflect"
	"strings"
	"sync"
)

var rwlock sync.RWMutex = sync.RWMutex{}
var adapterDict map[string]reflect.Type = map[string]reflect.Type{}

func init() {
	rwlock.Lock()
	defer rwlock.Unlock()
	adapterDict[reflect.TypeOf(ModelListTemplateAdapter{}).Name()] = reflect.TypeOf(ModelListTemplateAdapter{})
	adapterDict[reflect.TypeOf(ModelFormTemplateAdapter{}).Name()] = reflect.TypeOf(ModelFormTemplateAdapter{})
}

func GetAdapterDict() map[string]reflect.Type {
	rwlock.RLock()
	defer rwlock.RUnlock()
	return adapterDict
}

type CommonMethod struct{}

// TODO, byTest
func (o CommonMethod) Parse(classMethod string, param []interface{}) []reflect.Value {
	exprContent := classMethod
	scriptStruct := strings.Split(exprContent, ".")[0]
	scriptStructMethod := strings.Split(exprContent, ".")[1]
	scriptType := GetAdapterDict()[scriptStruct]
	if scriptType == nil {
		panic("adatpter " + scriptStruct + " not found")
	}
	inst := reflect.New(scriptType).Elem().Interface()
	instValue := reflect.ValueOf(inst)
	in := []reflect.Value{}
	for _, item := range param {
		in = append(in, reflect.ValueOf(item))
	}
	return instValue.MethodByName(scriptStructMethod).Call(in)
}

func (o CommonMethod) recursionApplyColumnModel(dataSource DataSource, columnModel *ColumnModel, result *interface{}) {
	modelIterator := ModelIterator{}
	for i, _ := range columnModel.ColumnLi {
		column := &columnModel.ColumnLi[i]
		if column.ColumnModel.ColumnLi == nil {
			if column.XMLName.Local == "auto-column" || column.Auto == "true" {
				modelIterator.IterateAllField(&dataSource, result, func(fieldGroup *FieldGroup, result *interface{}) {
					isApplyColumn := false
					columnModelDataSetId := columnModel.DataSetId
					if columnModelDataSetId == "" {
						columnModelDataSetId = "A"
					}
					isApplyColumn = isApplyColumn || (fieldGroup.GetDataSetId() == columnModelDataSetId && column.Name == fieldGroup.Id)
					if isApplyColumn {
						if column.Text == "" {
							column.Text = fieldGroup.DisplayName
						}
						if column.Hideable == "" {
							column.Hideable = fieldGroup.FixHide
						}
						if column.XMLName.Local == "auto-column" {
							o.applyAutoColumnXMLName(*fieldGroup, column)
						}

						o.applyColumnExtend(*fieldGroup, column)
					}
				})
			}
		} else {
			o.recursionApplyColumnModel(dataSource, &column.ColumnModel, result)
		}
	}
}

func (o CommonMethod) applyAutoColumnXMLName(fieldGroup FieldGroup, column *Column) {
	xmlName := o.getColumnXMLName(fieldGroup)
	if xmlName != "" {
		column.XMLName.Local = xmlName
	}
}

func (o CommonMethod) getColumnXMLName(fieldGroup FieldGroup) string {
	isIntField := false
	intArray := []string{"SMALLINT", "INT", "LONGINT"}
	for _, item := range intArray {
		if strings.ToLower(fieldGroup.FieldDataType) == strings.ToLower(item) {
			isIntField = true
			break
		}
	}
	isFloatField := false
	floatArray := []string{"FLOAT", "MONEY", "DECIMAL"}
	for _, item := range floatArray {
		if strings.ToLower(fieldGroup.FieldDataType) == strings.ToLower(item) {
			isFloatField = true
			break
		}
	}
	isDateType := false
	dateArray := []string{"YEAR", "YEARMONTH", "DATE", "TIME", "DATETIME"}
	for _, item := range dateArray {
		if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower(item) {
			isDateType = true
			break
		}
	}
	if fieldGroup.IsRelationField() {
		return "select-column"
	} else if strings.ToLower(fieldGroup.FieldDataType) == "string" || strings.ToLower(fieldGroup.FieldDataType) == "remark" {
		return "string-column"
	} else if (isIntField || isFloatField) && !isDateType && fieldGroup.Dictionary == "" {
		return "number-column"
	} else if (isIntField || isFloatField) && isDateType && fieldGroup.Dictionary == "" {
		return "date-column"
	} else if (isIntField || isFloatField) && !isDateType && fieldGroup.Dictionary != "" {
		return "dictionary-column"
	}
	return ""
}

// TODO, bytest
func (o CommonMethod) applyColumnExtend(fieldGroup FieldGroup, column *Column) {
	if column.XMLName.Local == "select-column" {
		relationItem := fieldGroup.RelationDS.RelationItemLi[0]
		if column.DisplayField == "" {
			column.DisplayField = relationItem.DisplayField
		}
		if column.ValueField == "" {
			column.ValueField = relationItem.ValueField
		}
		if column.SelectorName == "" {
			column.SelectorName = relationItem.Id
		}
		if column.SelectionMode == "" {
			column.SelectionMode = "single"
		}
	} else if column.XMLName.Local == "string-column" {
		// do nothing
	} else if column.XMLName.Local == "number-column" {
		if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("MONEY") {
			if column.IsMoney == "" {
				column.IsMoney = "true"
			}
		} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("PRICE") {
			if column.IsUnitPrice == "" {
				column.IsUnitPrice = "true"
			}
		} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("UNITCOST") {
			if column.IsCost == "" {
				column.IsCost = "true"
			}
		} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("PERCENT") {
			if column.IsPercent == "" {
				column.IsPercent = "true"
			}
		} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("QUANTITY") {
			if column.IsQuantity == "" {
				column.IsQuantity = "true"
			}
		}
	} else if column.XMLName.Local == "date-column" {
		if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("DATE") {
			if column.DisplayPattern == "" {
				column.DisplayPattern = "yyyyMMdd" //需要从业务中查找,是一个系统配置,TODO,
			}
			if column.DbPattern == "" {
				column.DbPattern = "yyyyMMdd"
			}
		} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("DATETIME") {
			if column.DisplayPattern == "" {
				column.DisplayPattern = "yyyy-MM-dd HH:mm:ss" //需要从业务中查找,是一个系统配置,TODO,
			}
			if column.DbPattern == "" {
				column.DbPattern = "yyyyMMddHHmmss"
			}
		} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("YEAR") {
			if column.DisplayPattern == "" {
				column.DisplayPattern = "yyyy"
			}
			if column.DbPattern == "" {
				column.DbPattern = "yyyy"
			}
		} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("YEARMONTH") {
			if column.DisplayPattern == "" {
				column.DisplayPattern = "yyyyMM" //需要从业务中查找,是一个系统配置,TODO,
			}
			if column.DbPattern == "" {
				column.DbPattern = "yyyyMM"
			}
		} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("TIME") {
			if column.DisplayPattern == "" {
				column.DisplayPattern = "HHmmss" //需要从业务中查找,是一个系统配置,TODO,
			}
			if column.DbPattern == "" {
				column.DbPattern = "HHmmss"
			}
		}
	} else if column.XMLName.Local == "dictionary-column" {
		if column.Dictionary == "" {
			column.Dictionary = fieldGroup.Dictionary
		}
	}
}
