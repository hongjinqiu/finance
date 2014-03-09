function ListTemplateIterator() {}

function IterateFunc(column, result) {
}

ListTemplateIterator.prototype._iterateTemplateColumn = function(result, isContinue, iterateFunc) {
	for (var i = 0; i < listTemplate.ColumnModel.ColumnLi.length; i++) {
		var column = listTemplate.ColumnModel.ColumnLi[i];
		var iterateResult = iterateFunc(column);
		if (!isContinue && iterateResult) {
			return;
		}
	}
}

ListTemplateIterator.prototype.iterateAllTemplateColumn = function(result, iterateFunc) {
	var self = this;
	var isContinue = true;
	self._iterateTemplateColumn(result, isContinue, iterateFunc);
}

ListTemplateIterator.prototype.iterateAnyTemplateColumn = function(dataSetId, result, iterateFunc) {
	var self = this;
	var isContinue = false;
	self._iterateTemplateColumn(result, isContinue, iterateFunc);
}

function IterateFunc(queryParameter, result) {
}

ListTemplateIterator.prototype._iterateTemplateQueryParameter = function(result, isContinue, iterateFunc) {
	for (var i = 0; i < listTemplate.QueryParameterGroup.QueryParameterLi.length; i++) {
		var queryParameter = listTemplate.QueryParameterGroup.QueryParameterLi[i];
		var iterateResult = iterateFunc(column);
		if (!isContinue && iterateResult) {
			return;
		}
	}
}

ListTemplateIterator.prototype.iterateAllTemplateQueryParameter = function(result, iterateFunc) {
	var self = this;
	var isContinue = true;
	self._iterateTemplateQueryParameter(result, isContinue, iterateFunc);
}

ListTemplateIterator.prototype.iterateAnyTemplateQueryParameter = function(dataSetId, result, iterateFunc) {
	var self = this;
	var isContinue = false;
	self._iterateTemplateQueryParameter(result, isContinue, iterateFunc);
}
