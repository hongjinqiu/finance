package component

import (

)

type FormTemplateIterator struct {}

type IterateFormTemplateColumnFunc func(column Column, result *interface{})

func (o FormTemplateIterator) IterateTemplateColumn(formTemplate FormTemplate, result *interface{}, iterateFunc IterateFormTemplateColumnFunc) {
	if formTemplate.FormElemLi != nil {
		listTemplateIterator := ListTemplateIterator{}
		for _, item := range formTemplate.FormElemLi {
			if item.XMLName.Local == "column-model" {
				columnLi := []Column{}
				listTemplateIterator.recursionGetColumnItem(item.ColumnModel, &columnLi)
				for _, item := range columnLi {
					iterateFunc(item, result)
				}
			}
		}
	}
}
