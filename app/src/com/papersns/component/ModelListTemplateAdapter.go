package component

import (
	. "com/papersns/model"
	"strings"
)

type ModelListTemplateAdapter struct{}

// TODO, bytest
func (o ModelListTemplateAdapter) ApplyAdapter(listTemplate *ListTemplate) {
	if listTemplate.DataSourceModelId != "" {
		modelTemplateFactory := ModelTemplateFactory{}
		dataSource := modelTemplateFactory.GetDataSource(listTemplate.DataSourceModelId)
		o.applyDataProvider(dataSource, listTemplate)
		o.applyColumnModel(dataSource, listTemplate)
		o.applyQueryParameter(dataSource, listTemplate)
	}
}

// TODO, bytest
func (o ModelListTemplateAdapter) ApplyQueryParameter(listTemplate *ListTemplate, queryParameter *QueryParameter) {
	if listTemplate.DataSourceModelId != "" {
		if listTemplate.QueryParameterGroup.DataSetId != "" {
			queryParameter.Name = listTemplate.QueryParameterGroup.DataSetId + "." + queryParameter.Name
		}
	}
}

// TODO, bytest
func (o ModelListTemplateAdapter) ApplyColumnName(listTemplate *ListTemplate, column *Column) {
	if listTemplate.DataSourceModelId != "" {
		if listTemplate.QueryParameterGroup.DataSetId != "" {
			column.Name = listTemplate.QueryParameterGroup.DataSetId + "." + column.Name
		}
	}
}

// TODO, bytest
func (o ModelListTemplateAdapter) applyDataProvider(dataSource DataSource, listTemplate *ListTemplate) {
	if listTemplate.DataProvider.Collection == "" {
		listTemplate.DataProvider.Collection = dataSource.Id
	}
}

// TODO, bytest
func (o ModelListTemplateAdapter) applyColumnModel(dataSource DataSource, listTemplate *ListTemplate) {
	var result interface{} = ""
	commonMethod := CommonMethod{}
	commonMethod.recursionApplyColumnModel(dataSource, &listTemplate.ColumnModel, &result)
}

// TODO, bytest
func (o ModelListTemplateAdapter) applyQueryParameter(dataSource DataSource, listTemplate *ListTemplate) {
	commonMethod := CommonMethod{}
	var result interface{} = ""
	modelIterator := ModelIterator{}
	for i, _ := range listTemplate.QueryParameterGroup.QueryParameterLi {
		queryParameter := listTemplate.QueryParameterGroup.QueryParameterLi[i]
		if queryParameter.Auto == "true" {
			modelIterator.IterateAllField(&dataSource, &result, func(fieldGroup *FieldGroup, result *interface{}){
				if fieldGroup.IsMasterField() {
					name := queryParameter.Name
					if queryParameter.ColumnName != "" {
						name = queryParameter.ColumnName
					}
					if name == fieldGroup.Id {
						if queryParameter.Text == "" {
							queryParameter.Text = fieldGroup.DisplayName
						}
						if fieldGroup.FixHide == "true" {
							if queryParameter.Editor == "" {
								queryParameter.Editor = "hidden"
							}
						}
//						if queryParameter.Hidden == "" {
//							queryParameter.Hidden = fieldGroup.FixHide
//						}
						xmlName := commonMethod.getColumnXMLName(*fieldGroup)
						if xmlName != "" {
							o.applyQueryParameterAttr(xmlName, &queryParameter)
							o.applyQueryParameterSubAttr(xmlName, *fieldGroup, &queryParameter)
						}
					}
				}
			})
		}
	}
}

func (o ModelListTemplateAdapter) applyQueryParameterAttr(xmlName string, queryParameter *QueryParameter) {
	if xmlName == "select-column" {
		if queryParameter.Editor == "" {
			queryParameter.Editor = "trigger"
		}
		if queryParameter.Restriction == "" {
			queryParameter.Restriction = "in"
		}
	} else if xmlName == "string-column" {
		if queryParameter.Editor == "" {
			queryParameter.Editor = "textfield"
		}
		if queryParameter.Restriction == "" {
			queryParameter.Restriction = "like"
		}
	} else if xmlName == "number-column" {
		if queryParameter.Editor == "" {
			queryParameter.Editor = "numberfield"
		}
		if queryParameter.Restriction == "" {
			queryParameter.Restriction = "eq"
		}
	} else if xmlName == "date-column" {
		if queryParameter.Editor == "" {
			queryParameter.Editor = "datefield"
		}
		if queryParameter.Restriction == "" {
			queryParameter.Restriction = "eq"
		}
	} else if xmlName == "dictionary-column" {
		if queryParameter.Editor == "" {
			queryParameter.Editor = "combo"
		}
		if queryParameter.Restriction == "" {
			queryParameter.Restriction = "eq"
		}
	}
}

func (o ModelListTemplateAdapter) applyQueryParameterSubAttr(xmlName string, fieldGroup FieldGroup, queryParameter *QueryParameter) {
	if xmlName == "select-column" {
		// do nothing
	} else if xmlName == "string-column" {
		// do nothing
	} else if xmlName == "number-column" {
		// do nothing
	} else if xmlName == "date-column" {
		hasInFormat := false
		hasQueryFormat := false
		if queryParameter.ParameterAttributeLi != nil {
			for _, attrItem := range queryParameter.ParameterAttributeLi {
				if attrItem.Name == "inFormat" {
					hasInFormat = true
					break
				}
			}
			for _, attrItem := range queryParameter.ParameterAttributeLi {
				if attrItem.Name == "queryFormat" {
					hasQueryFormat = true
					break
				}
			}
		}
		if !hasInFormat || !hasQueryFormat {
			if queryParameter.ParameterAttributeLi == nil {
				queryParameter.ParameterAttributeLi = []ParameterAttribute{}
			}
			if !hasInFormat {
				parameterAttribute := ParameterAttribute{}
				parameterAttribute.Name = "inFormat"
				if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("DATE") {
					parameterAttribute.Value = "yyyyMMdd"//需要从业务中查找,是一个系统配置,TODO,
				} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("DATETIME") {
					parameterAttribute.Value = "yyyy-MM-dd HH:mm:ss"//需要从业务中查找,是一个系统配置,TODO,
				} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("YEAR") {
					parameterAttribute.Value = "yyyy"
				} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("YEARMONTH") {
					parameterAttribute.Value = "yyyyMM"//需要从业务中查找,是一个系统配置,TODO,
				} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("TIME") {
					parameterAttribute.Value = "HHmmss"//需要从业务中查找,是一个系统配置,TODO,
				}
				queryParameter.ParameterAttributeLi = append(queryParameter.ParameterAttributeLi, parameterAttribute)
			}
			if !hasQueryFormat {
				parameterAttribute := ParameterAttribute{}
				parameterAttribute.Name = "queryFormat"
				if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("DATE") {
					parameterAttribute.Value = "yyyyMMdd"
				} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("DATETIME") {
					parameterAttribute.Value = "yyyyMMddHHmmss"
				} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("YEAR") {
					parameterAttribute.Value = "yyyy"
				} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("YEARMONTH") {
					parameterAttribute.Value = "yyyyMM"
				} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("TIME") {
					parameterAttribute.Value = "HHmmss"
				}
				queryParameter.ParameterAttributeLi = append(queryParameter.ParameterAttributeLi, parameterAttribute)
			}
		}
	} else if xmlName == "dictionary-column" {
		hasDictionary := false
		if queryParameter.ParameterAttributeLi != nil {
			for _, attrItem := range queryParameter.ParameterAttributeLi {
				if attrItem.Name == "dictionary" {
					hasDictionary = true
					break
				}
			}
		}
		if !hasDictionary {
			if queryParameter.ParameterAttributeLi == nil {
				queryParameter.ParameterAttributeLi = []ParameterAttribute{}
			}
			parameterAttribute := ParameterAttribute{}
			parameterAttribute.Name = "dictionary"
			parameterAttribute.Value = fieldGroup.Dictionary
			queryParameter.ParameterAttributeLi = append(queryParameter.ParameterAttributeLi, parameterAttribute)
		}
	}
}

