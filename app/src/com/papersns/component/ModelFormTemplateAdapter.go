package component

import (
	. "com/papersns/model"
)

type ModelFormTemplateAdapter struct{}

// TODO, bytest
func (o ModelFormTemplateAdapter) ApplyAdapter(iFormTemplate interface{}) {
	formTemplate := (iFormTemplate).(FormTemplate)
	if formTemplate.DataSourceModelId != "" {
		modelTemplateFactory := ModelTemplateFactory{}
		dataSource := modelTemplateFactory.GetDataSource(formTemplate.DataSourceModelId)
		o.applyDetailDataSet(dataSource, &formTemplate)
	}
}

// TODO, bytest
func (o ModelFormTemplateAdapter) applyDetailDataSet(dataSource DataSource, formTemplate *FormTemplate) {
	commonMethod := CommonMethod{}
	var result interface{} = ""
	for i, _ := range formTemplate.FormElemLi {
		if formTemplate.FormElemLi[i].XMLName.Local == "column-model" {
			if formTemplate.FormElemLi[i].ColumnModel.DataSetId != "A" {
				commonMethod.recursionApplyColumnModel(dataSource, &formTemplate.FormElemLi[i].ColumnModel, &result)
			}
		}
	}
}
