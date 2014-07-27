var modelExtraInfo = {
	"A" : {
		"code" : {
			selectorName : "",// 可以为函数
			listeners: {
			}
		}
	}
};

function main(Y) {
		if (g_id) {
			if (g_copyFlag == "true") {// 复制
				ajaxRequest({
					url: "/" + g_dataSourceJson.Id + "/CopyData?format=json"
					,params: {
						"dataSourceModelId": g_dataSourceJson.Id,
						"formTemplateId": g_formTemplateJsonData.Id,
						"id": g_id
					},
					callback: function(o) {
						var formManager = new FormManager();
						formManager.setDetailIncId(g_dataSourceJson, o.bo);
						formManager.applyGlobalParamFromAjaxData(o);
						formManager.loadData2Form(g_dataSourceJson, o.bo);
						formManager.setFormStatus("edit");
					}
				});
			} else {
				ajaxRequest({
					url: "/" + g_dataSourceJson.Id + "/GetData?format=json"
					,params: {
						"dataSourceModelId": g_dataSourceJson.Id,
						"formTemplateId": g_formTemplateJsonData.Id,
						"id": g_id
					},
					callback: function(o) {
						var formManager = new FormManager();
						formManager.applyGlobalParamFromAjaxData(o);
						formManager.loadData2Form(g_dataSourceJson, o.bo);
						formManager.setFormStatus(g_formStatus);
					}
				});
			}
		} else {
			ajaxRequest({
				url: "/" + g_dataSourceJson.Id + "/NewData?format=json"
				,params: {
					"dataSourceModelId": g_dataSourceJson.Id,
					"formTemplateId": g_formTemplateJsonData.Id
				},
				callback: function(o) {
					var formManager = new FormManager();
					formManager.applyGlobalParamFromAjaxData(o);
					formManager.loadData2Form(g_dataSourceJson, o.bo);
					formManager.setFormStatus(g_formStatus);
				}
			});
		}
	}
