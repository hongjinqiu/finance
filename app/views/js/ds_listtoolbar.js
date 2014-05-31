function g_deleteRecord(o) {
	showWarning("确认删除？", function(){
		var url = "/" + listTemplate.DataSourceModelId + "/DeleteData?format=json";
		ajaxRequest({
			url: url
			,params: {
				"id": o.get("id"),
				"dataSourceModelId": listTemplate.DataSourceModelId
			},
			callback: function(o) {
//			showSuccess("删除数据成功");
				g_gridPanelDict["columnModel_1"].dt.refreshPaginator();
			}
		});
	});
}

function g_deleteRecords() {
	var selectRecords = g_gridPanelDict["columnModel_1"].getSelectRecordLi();
	if (selectRecords.length > 0) {
		showWarning("确认删除？", function(){
			var errorMsgLi = [];
			
			var url = "/" + listTemplate.DataSourceModelId + "/DeleteData?format=json";
			for (var i = 0; i < selectRecords.length; i++) {
				ajaxRequest({
					url: url
					,params: {
						"id": selectRecords[i].get("id"),
						"dataSourceModelId": listTemplate.DataSourceModelId
					},
					callback: function(o) {
						//g_gridPanelDict["columnModel_1"].dt.refreshPaginator();
					},
					failCallback: function(o) {
						var message = "记录" + selectRecords[i].get("code");
						message += "：" + o.message+"；";
						errorMsgLi.push(message);
					}
				});
			}
			if (errorMsgLi.length > 0) {
				showError(errorMsgLi.join("<br />"));
			}
			g_gridPanelDict["columnModel_1"].dt.refreshPaginator();
		});
	} else {
		showAlert("请选择记录！");
	}
}

