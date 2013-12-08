package model

import (
)

type ModelIterator struct {}

type IterateFunc func(fieldGroup *FieldGroup, data *map[string]interface{}, result *interface{})

func (o ModelIterator) IterateAllFieldBo(dataSource *DataSource, bo *map[string]interface{}, result *interface{}, iterateFunc IterateFunc) {
	data := (*bo)["A"].(map[string]interface{})
	o.iterateFixFieldBo(&dataSource.MasterData.FixField, &data, result, iterateFunc)
	o.iterateBizFieldBo(&dataSource.MasterData.BizField, &data, result, iterateFunc)
	
	for i, _ := range dataSource.DetailDataLi {
		dataLi := (*bo)[dataSource.DetailDataLi[i].Id].([]interface{})
		for subIndex, _ := range dataLi {
			detailData := dataLi[subIndex].(map[string]interface{})
			o.iterateFixFieldBo(&dataSource.DetailDataLi[i].FixField, &detailData, result, iterateFunc)
			o.iterateBizFieldBo(&dataSource.DetailDataLi[i].BizField, &detailData, result, iterateFunc)
		}
	}
}

func (o ModelIterator) GetFixFieldLi(fixField *FixField) []*FieldGroup {
	fixFieldLi := []*FieldGroup{}
	fixFieldLi = append(fixFieldLi, &fixField.PrimaryKey.FieldGroup)
	fixFieldLi = append(fixFieldLi, &fixField.CreateBy.FieldGroup)
	fixFieldLi = append(fixFieldLi, &fixField.CreateTime.FieldGroup)
	fixFieldLi = append(fixFieldLi, &fixField.ModifyBy.FieldGroup)
	fixFieldLi = append(fixFieldLi, &fixField.ModifyTime.FieldGroup)
	fixFieldLi = append(fixFieldLi, &fixField.BillStatus.FieldGroup)
	fixFieldLi = append(fixFieldLi, &fixField.AttachCount.FieldGroup)
	fixFieldLi = append(fixFieldLi, &fixField.Remark.FieldGroup)
	return fixFieldLi
}

func (o ModelIterator) iterateFixFieldBo(fixField *FixField, data *map[string]interface{}, result *interface{}, iterateFunc IterateFunc) {
	fixFieldLi := o.GetFixFieldLi(fixField)
	
	for i, _ := range fixFieldLi {
		iterateFunc(fixFieldLi[i], data, result)
	}
}

func (o ModelIterator) iterateBizFieldBo(bizField *BizField, data *map[string]interface{}, result *interface{}, iterateFunc IterateFunc) {
	for i, _ := range bizField.FieldLi {
		iterateFunc(&bizField.FieldLi[i].FieldGroup, data, result)
	}
}

type IterateFieldFunc func(fieldGroup *FieldGroup, result *interface{})

func (o ModelIterator) IterateAllField(dataSource *DataSource, result *interface{}, iterateFunc IterateFieldFunc) {
	o.iterateFixField(&dataSource.MasterData.FixField, result, iterateFunc)
	o.iterateBizField(&dataSource.MasterData.BizField, result, iterateFunc)
	
	for i, _ := range dataSource.DetailDataLi {
		o.iterateFixField(&dataSource.DetailDataLi[i].FixField, result, iterateFunc)
		o.iterateBizField(&dataSource.DetailDataLi[i].BizField, result, iterateFunc)
	}
}

func (o ModelIterator) iterateFixField(fixField *FixField, result *interface{}, iterateFunc IterateFieldFunc) {
	fixFieldLi := o.GetFixFieldLi(fixField)
	
	for i, _ := range fixFieldLi {
		iterateFunc(fixFieldLi[i], result)
	}
}

func (o ModelIterator) iterateBizField(bizField *BizField, result *interface{}, iterateFunc IterateFieldFunc) {
	for i, _ := range bizField.FieldLi {
		iterateFunc(&bizField.FieldLi[i].FieldGroup, result)
	}
}

