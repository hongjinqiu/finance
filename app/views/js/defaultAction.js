function selectRowBtnDefaultAction(dataSetId, toolbarOrColumnModel, button, selectValueLi) {
	var formManager = new FormManager();
	var templateIterator = new TemplateIterator();
	var result = "";

	// use default action
	var dataLi = [];
	var columnResult = "";
	var columnLi = [];
	templateIterator.iterateAllTemplateColumn(dataSetId, columnResult, function IterateFunc(column, result) {
		columnLi.push(column);
	});
	for (var i = 0; i < selectValueLi.length; i++) {
		var data = formManager.getDataSetNewData(dataSetId);
		if (button.CRelationDS && button.CRelationDS.CRelationItemLi) {
			var relationItem = button.CRelationDS.CRelationItemLi[0];
			if (relationItem.CCopyConfigLi) {
				for (var j = 0; j < relationItem.CCopyConfigLi.length; j++) {
					var columnName = relationItem.CCopyConfigLi[j].CopyColumnName;
					var columnValue = selectValueLi[i].CopyValueField;
					_recurionApplyCopyField(data, columnLi, columnName, columnValue)
				}
			}
		}
		dataLi.push(data);
	}
	// 允许重复的判断,
	var gridDataLi = g_gridPanelDict["B"].dt.get("data").toJSON();
	var notAllowDuplicateColumn = [];
	var modelIterator = new ModelIterator();
	modelIterator.iterateAllField(dataSource, result, function(fieldGroup, result){
		if (fieldGroup.getDataSetId() == dataSetId && fieldGroup.AllowDuplicate == "false") {
			notAllowDuplicateColumn.push(fieldGroup.Id);
		}
	});
	for (var i = 0; i < dataLi.length; i++) {
		var isIn = false;
		for (var j = 0; j < gridDataLi.length; j++) {
			var flag = true;
			for (var k = 0; k < notAllowDuplicateColumn.length; k++) {
				flag = flag && (dataLi[i][notAllowDuplicateColumn[k]] == gridDataLi[j][notAllowDuplicateColumn[k]]);
			}
			if (flag) {
				isIn = true;
				break
			}
		}
		if (!isIn) {
			gridDataLi.push(dataLi[i]);
		}
	}
	g_gridPanelDict["B"].dt.set("data", gridDataLi);
}