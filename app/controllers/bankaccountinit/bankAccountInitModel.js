var modelExtraInfo = {

};

function main() {
	YUI(g_financeModule).use("finance-module", function(YNotUse) {// 不能直接在父函数用use finance-module,会报错,因为在js父函数直接加载,其会直接使用调用
		ajaxRequest({
			url : "/BankAccountInit/GetData?format=json",
			params : {
				"dataSourceModelId" : g_dataSourceJson.Id,
				"formTemplateId" : g_formTemplateJsonData.Id
			},
			callback : function(o) {
				var formManager = new FormManager();
				formManager.applyGlobalParamFromAjaxData(o);
				formManager.loadData2Form(g_dataSourceJson, o.bo);
				formManager.setFormStatus(g_formStatus);
				if (g_formStatus == "view") {
					enableQueryParameters();
				} else {
					disableQueryParameters();
				}
			}
		});
		/*
		if (g_id) {
			ajaxRequest({
				url: "/BankAccountInit/GetData?format=json"
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
				url: "/BankAccountInit/NewData?format=json"
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
		*/
	});
}

function enableQueryParameters() {
	g_masterFormFieldDict["accountId"].set("readonly", false);
	document.getElementById("queryBtn").disabled = "";
	document.getElementById("resetBtn").disabled = "";
	
	document.getElementById("queryBtn").style.border = "1px solid red";
	document.getElementById("resetBtn").style.border = "1px solid red";
}

function disableQueryParameters() {
	g_masterFormFieldDict["accountId"].set("readonly", true);
	document.getElementById("queryBtn").disabled = "true";
	document.getElementById("resetBtn").disabled = "true";
	
	document.getElementById("queryBtn").style.border = "1px solid blue";
	document.getElementById("resetBtn").style.border = "1px solid blue";
}

function bankAccountInitGiveUpData() {
	var formManager = new FormManager();
	var bo = formManager.getBo();
	showConfirm("您确定要放弃吗？", function() {
		ajaxRequest({
			url : "/" + g_dataSourceJson.Id + "/GiveUpData?format=json",
			params : {
				"dataSourceModelId" : g_dataSourceJson.Id,
				"formTemplateId" : g_formTemplateJsonData.Id,
				"queryData": bo["A"]
			},
			callback : function(o) {
				formManager.applyGlobalParamFromAjaxData(o);
				formManager.loadData2Form(g_dataSourceJson, o.bo);
				formManager.setFormStatus("view");
				enableQueryParameters();
			}
		});
	});
}

function bankAccountInitQueryData() {
	var formManager = new FormManager();
	var bo = formManager.getBo();
	ajaxRequest({
		url: "/" + g_dataSourceJson.Id + "/RefreshData?format=json"
		,params: {
			"dataSourceModelId": g_dataSourceJson.Id,
			"formTemplateId": g_formTemplateJsonData.Id,
			"queryData": bo["A"]
		},
		callback: function(o) {
			formManager.setFormStatus("view");
			formManager.applyGlobalParamFromAjaxData(o);
			formManager.loadData2Form(g_dataSourceJson, o.bo);
			enableQueryParameters();
		}
	});
}

function bankAccountInitEditData() {
	var formManager = new FormManager();
	var bo = formManager.getBo();
	ajaxRequest({
		url: "/" + g_dataSourceJson.Id + "/EditData?format=json"
		,params: {
			"dataSourceModelId": g_dataSourceJson.Id,
			"formTemplateId": g_formTemplateJsonData.Id,
			"queryData": bo["A"]
		},
		callback: function(o) {
			formManager.applyGlobalParamFromAjaxData(o);
			formManager.loadData2Form(g_dataSourceJson, o.bo);
			formManager.setFormStatus("edit");
			disableQueryParameters();
		}
	});
}

function resetQueryParameter() {
	for ( var key in g_masterFormFieldDict) {
		g_masterFormFieldDict[key].set("value", "");
	}
}
