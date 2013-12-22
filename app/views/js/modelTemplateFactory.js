function ModelTemplateFactory() {
}

/**
 * 建立父子双向关联
 */
/*func (o ModelTemplateFactory) applyReverseRelation(dataSource *DataSource) {
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
}*/
ModelTemplateFactory.prototype._applyReverseRelation = function() {
	
}

