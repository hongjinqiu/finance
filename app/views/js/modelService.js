function ModelIterator() {
}

function IterateFunc(fieldGroup, data, result) {
}

ModelIterator.prototype.iterateAllFieldBo = function(dataSource, bo, result, iterateFunc) {
	var self = this;
	self.iterateDataBo(dataSource, bo, result, function(fieldGroupLi, data, result){
		for (var i = 0; i < fieldGroupLi.length; i++) {
			iterateFunc(fieldGroupLi[i], data, result);
		}
	})
}

function IterateFieldFunc(fieldGroup, result){}

ModelIterator.prototype.iterateAllField = function(dataSource, result, iterateFunc) {
	var self = this;
	var fieldGroupLi = self._getDataSetFieldGroupLi(dataSource.MasterData.FixField, dataSource.MasterData.BizField)
	for (var i = 0; i < fieldGroupLi.length; i++) {
		iterateFunc(fieldGroupLi[i], result);
	}
	for (var i = 0; i < dataSource.DetailDataLi; i++) {
		var fieldGroupLi = self._getDataSetFieldGroupLi(dataSource.DetailDataLi[i].FixField, dataSource.DetailDataLi[i].BizField);
		for (var j = 0; j < fieldGroupLi.length; j++) {
			iterateFunc(fieldGroupLi[j], result);
		}
	}
}


ModelIterator.prototype.getFixFieldLi = function(fixField) {
	var fixFieldLi = [];
	fixFieldLi.push(fixField.PrimaryKey.FieldGroup);
	fixFieldLi.push(fixField.CreateBy.FieldGroup);
	fixFieldLi.push(fixField.CreateTime.FieldGroup);
	fixFieldLi.push(fixField.CreateUnit.FieldGroup);
	fixFieldLi.push(fixField.ModifyBy.FieldGroup);
	fixFieldLi.push(fixField.ModifyTime.FieldGroup);
	fixFieldLi.push(fixField.ModifyUnit.FieldGroup);
	fixFieldLi.push(fixField.BillStatus.FieldGroup);
	fixFieldLi.push(fixField.AttachCount.FieldGroup);
	fixFieldLi.push(fixField.Remark.FieldGroup);
	return fixFieldLi;
}

ModelIterator.prototype._getDataSetFieldGroupLi = function(fixField, bizField) {
	var self = this;
	var fieldGroupLi = self.getFixFieldLi(fixField);
	for (var i = 0; i < bizField.FieldLi) {
		fieldGroupLi.push(bizField.FieldLi[i].FieldGroup);
	}
	return fieldGroupLi;
}

function IterateDataFunc(fieldGroupLi, data, result) {}

ModelIterator.prototype.iterateDataBo = function(dataSource, bo, result, iterateFunc) {
	var self = this;
	self._iterateMasterDataBo(dataSource, bo, result, iterateFunc);
	self._iterateDetailDataBo(dataSource, bo, result, iterateFunc)
}

ModelIterator.prototype._iterateMasterDataBo = function(dataSource, bo, result, iterateFunc) {
	var self = this;
	var data = bo["A"];
	var fieldGroupLi = self._getDataSetFieldGroupLi(dataSource.MasterData.FixField, dataSource.MasterData.BizField)
	iterateFunc(fieldGroupLi, data, result)
}

ModelIterator.prototype._iterateDetailDataBo = function(dataSource, bo, result, iterateFunc) {
	var self = this;
	for (var i = 0; i < dataSource.DetailDataLi.length; i++) {
		var item = dataSource.DetailDataLi[i];
		var fieldGroupLi = self._getDataSetFieldGroupLi(item.FixField, item.BizField);
		var dataLi = bo[item.Id];
		for (var j = 0; j < dataLi.length; j++) {
			var data = dataLi[j];
			iterateFunc(fieldGroupLi, data, result);
		}
	}
}


