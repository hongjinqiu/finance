/**
 * 没用到,应该会废弃掉
 */
function ColumnDataSourceManager() {
}

ColumnDataSourceManager.prototype.createColumn = function(columnConfig) {
	var self = this;
	if (true) {
		return {
			key: columnConfig.Name,
			allowHTML:  true, // to avoid HTML escaping
			label: columnConfig.Text,
			formatter: function(o) {
				if (!columnConfig.Name) {
					return "";
				}
				return "test";
				console.log(o);
				var renderId = o.column.id + "_" + columnConfig.Name;
//				YUI(formJsConfig).use("node", "event", "papersns-form", function(Y) {
//					var field = new Y.PTextField({name : 'input1', type : 'text', required: true, validateInline: true, validator : function(value, formFieldObj){
//						if (value != "abc") {
//							formFieldObj.set("error", "值必须为abc");
//							return false;
//						}
//						return true;
//					}});
//					field.set("renderId", renderId);
//					if (!o.record.formFieldDict) {
//						o.record.formFieldDict = {};
//					}
//					o.record.formFieldDict[columnConfig.Name] = field;
//				});
				return "<div id='" + renderId + "'></div>";
			}
		}
	}
	var self = this;
	if (columnConfig.XMLName.Local != "virtual-column" && columnConfig.Hideable != "true") {
		if (columnConfig.ColumnModel.ColumnLi) {
			var result = {
				label: columnConfig.Text,
				"children": []
			};
			for (var i = 0; i < columnConfig.ColumnModel.ColumnLi.length; i++) {
				var childColumn = self.createColumn(columnConfig.ColumnModel.ColumnLi[i]);
				if (childColumn) {
					result.children.push(childColumn);
				}
			}
			return result;
		}
		
		if (columnConfig.XMLName.Local == "number-column") {
			return self.createNumberColumn(columnConfig);
		} else if (columnConfig.XMLName.Local == "date-column") {
			return self.createDateColumn(columnConfig);
		} else if (columnConfig.XMLName.Local == "boolean-column") {
			return self.createBooleanColumn(columnConfig);
		} else if (columnConfig.XMLName.Local == "dictionary-column") {
			return self.createDictionaryColumn(columnConfig);
		}
		return {
			key: columnConfig.Name,
			label: columnConfig.Text
		};
	}
	return null;
}

ColumnDataSourceManager.prototype.getColumns = function(columnModelName, columnModel, Y) {
	var self = this;
	self.yInst = Y;
	var columnManager = new ColumnManager();
	columnManager.yInst = Y;
	
	var columns = [];
	var checkboxColumn = columnManager.createCheckboxColumn(columnModel);
	if (checkboxColumn) {
		columns.push(checkboxColumn);
	}
	var idColumn = columnManager.createIdColumn(columnModel);
	if (idColumn) {
		columns.push(idColumn);
	}
	var rowIndexColumn = columnManager.createRowIndexColumn(columnModel);
	if (rowIndexColumn) {
		columns.push(rowIndexColumn);
	}
	
	for (var i = 0; i < columnModel.ColumnLi.length; i++) {
		var column = self.createColumn(columnModel.ColumnLi[i]);
		if (column) {
			columns.push(column);
		} else {
			var virtualColumn = columnManager.createVirtualColumn(columnModelName, columnModel, i);
			if (virtualColumn) {
				columns.push(virtualColumn);
			}
		}
	}
	return columns;
}
