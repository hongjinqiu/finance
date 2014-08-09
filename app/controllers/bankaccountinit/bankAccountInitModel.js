var modelExtraInfo = {

};

function main(Y) {
	
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
}

function enableQueryParameters() {
	g_masterFormFieldDict["accountId"].set("readonly", false);
	var qBtn = document.getElementById("queryBtn");
	var rBtn = document.getElementById("resetBtn");
	qBtn.disabled = "";
	rBtn.disabled = "";
	
	if (qBtn.className == "disable_but_box") {
		qBtn.className = qBtn.origClassName;
	}
	
	if (rBtn.className == "disable_but_box") {
		rBtn.className = rBtn.origClassName;
	}
}

function disableQueryParameters() {
	g_masterFormFieldDict["accountId"].set("readonly", true);
	var qBtn = document.getElementById("queryBtn");
	var rBtn = document.getElementById("resetBtn");
	qBtn.disabled = "true";
	rBtn.disabled = "true";
	
	if (!qBtn.origClassName) {
		qBtn.origClassName = qBtn.className;
	}
	qBtn.className = "disable_but_box";
	
	if (!rBtn.origClassName) {
		rBtn.origClassName = rBtn.className;
	}
	rBtn.className = "disable_but_box";
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
