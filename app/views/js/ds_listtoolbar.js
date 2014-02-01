function doDeleteRecord(o) {
	var url = "/" + listTemplate.DataSourceModelId + "/DeleteData?format=json";
	ajaxRequest({
		url: url
		,params: {
			"id": o.get("id"),
			"dataSourceModelId": listTemplate.DataSourceModelId
		},
		callback: function(o) {
//			showSuccess("删除数据成功");
			gridPanelDict["columnModel_1"].dt.refreshPaginator();
		}
	});
}
