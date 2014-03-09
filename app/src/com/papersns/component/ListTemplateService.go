package component

import (
	
)

type ListTemplateIterator struct {}

type IterateTemplateColumnFunc func(column Column, result *interface{})

func (o ListTemplateIterator) IterateTemplateColumn(listTemplate ListTemplate, result *interface{}, iterateFunc IterateTemplateColumnFunc) {
	for _, item := range listTemplate.ColumnModel.ColumnLi {
		iterateFunc(item, result)
	}
}

type IterateTemplateQueryParameterFunc func(queryParameter QueryParameter, result *interface{})

func (o ListTemplateIterator) IterateTemplateQueryParameter(listTemplate ListTemplate, result *interface{}, iterateFunc IterateTemplateQueryParameterFunc) {
	for _, item := range listTemplate.QueryParameterGroup.QueryParameterLi {
		iterateFunc(item, result)
	}
}
