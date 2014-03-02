var modelExtraInfo = {
		
};

function main() {
	YUI(g_financeModule).use("finance-module", function(YNotUse){// 不能直接在父函数用use finance-module,会报错,因为在js父函数直接加载,其会直接使用调用
		if (g_id) {
			ajaxRequest({
				url: "/" + dataSourceJson.Id + "/GetData?format=json"
				,params: {
					"dataSourceModelId": dataSourceJson.Id,
					"id": g_id
				},
				callback: function(o) {
					var formManager = new FormManager();
					formManager.loadData2Form(dataSourceJson, o.bo);
					formManager.setFormStatus(g_formStatus);
				}
			});
		} else {
			ajaxRequest({
				url: "/" + dataSourceJson.Id + "/NewData?format=json"
				,params: {
					"dataSourceModelId": dataSourceJson.Id
				},
				callback: function(o) {
					var formManager = new FormManager();
					formManager.loadData2Form(dataSourceJson, o.bo);
					formManager.setFormStatus(g_formStatus);
				}
			});
		}
	});
}

