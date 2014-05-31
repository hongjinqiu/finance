var modelExtraInfo = {
	"A" : {
		"code" : {
			selectorName : "",// 可以为函数
			listeners: {
			}
		}
	}
};

function main() {
	YUI(g_financeModule).use("finance-module", function(YNotUse){// 不能直接在父函数用use finance-module,会报错,因为在js父函数直接加载,其会直接使用调用
		if (g_id) {
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
	});
}
