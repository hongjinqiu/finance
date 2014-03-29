function ColumnDataSourceManager() {
}


ColumnDataSourceManager.prototype.getColumns = function(columnModelName, columnModel, Y) {
	var columnManager = new ColumnManager();
	var columns = columnManager.getColumns(columnModelName, columnModel, Y);
	if (g_dataSourceJson) {
		var modelIterator = new ModelIterator();
		var result = "";
		for (var i = 0; i < columns.length; i++) {
			columns[i].allowHTML = true;
			modelIterator.iterateAllField(g_dataSourceJson, result, function(fieldGroup, result){
				if (fieldGroup.getDataSetId() == columnModel.DataSetId && fieldGroup.Id == columns[i].key) {
					if (fieldGroup.AllowEmpty == "false") {
						columns[i].label = '<font style="color:red">*</font>' + columns[i].label;
					}
				}
			});
		}
	}
	return columns
}
