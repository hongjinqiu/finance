function TemplateIterator() {}

TemplateIterator.prototype._iterateTemplateColumn = function(dataSetId, result, isContinue, iterateFunc) {
	for (var j = 0; j < g_formTemplateJsonData.FormElemLi.length; j++) {
		var formElem = g_formTemplateJsonData.FormElemLi[j];
		if (formElem.XMLName.Local == "column-model") {
			if (formElem.ColumnModel.DataSetId == dataSetId) {
				if (formElem.ColumnModel.ColumnLi) {
					for (var k = 0; k < formElem.ColumnModel.ColumnLi.length; k++) {
						var column = formElem.ColumnModel.ColumnLi[k];
						var iterateResult = iterateFunc(column);
						if (!isContinue && iterateResult) {
							return;
						}
					}
				}
			}
		}
	}
}

function IterateFunc(column, result) {
}

TemplateIterator.prototype.iterateAllTemplateColumn = function(dataSetId, result, iterateFunc) {
	var self = this;
	var isContinue = true;
	self._iterateTemplateColumn(dataSetId, result, isContinue, iterateFunc);
}

function IterateFunc(column, result) {
}

TemplateIterator.prototype.iterateAnyTemplateColumn = function(dataSetId, result, iterateFunc) {
	var self = this;
	var isContinue = false;
	self._iterateTemplateColumn(dataSetId, result, isContinue, iterateFunc);
}

TemplateIterator.prototype._iterateTemplateColumnModel = function(result, isContinue, iterateFunc) {
	for (var j = 0; j < g_formTemplateJsonData.FormElemLi.length; j++) {
		var formElem = g_formTemplateJsonData.FormElemLi[j];
		if (formElem.XMLName.Local == "column-model") {
			var iterateResult = iterateFunc(formElem.ColumnModel);
			if (!isContinue && iterateResult) {
				return;
			}
		}
	}
}

function IterateFunc(columnModel, result) {
}

TemplateIterator.prototype.iterateAllTemplateColumnModel = function(result, iterateFunc) {
	var self = this;
	var isContinue = true;
	self._iterateTemplateColumnModel(result, isContinue, iterateFunc);
}

function IterateFunc(columnModel, result) {
}

TemplateIterator.prototype.iterateAnyTemplateColumnModel = function(result, iterateFunc) {
	var self = this;
	var isContinue = false;
	self._iterateTemplateColumnModel(result, isContinue, iterateFunc);
}
