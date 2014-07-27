var modelExtraInfo = {
	
};

function viewItem(record) {
	var queryMode = getCheckboxValue("queryMode");
	var url = "/console/listschema?@name=AccountInOutItem&accountType={accountType}&accountId={accountId}&currencyTypeId={currencyTypeId}&queryMode={queryMode}";
	url = url.replace("{accountType}", record.get("accountType"));
	url = url.replace("{accountId}", record.get("accountId"));
	url = url.replace("{currencyTypeId}", record.get("currencyTypeId"));
	url = url.replace("{queryMode}", queryMode);
	if (queryMode == "1") {// 按日期查询
		var dateParam = "billDateBegin={billDateBegin}&billDateEnd={billDateEnd}";
		dateParam = dateParam.replace("{billDateBegin}", g_masterFormFieldDict["billDateBegin"].get("value"));
		dateParam = dateParam.replace("{billDateEnd}", g_masterFormFieldDict["billDateEnd"].get("value"));
		url += "&" + dateParam;
	} else if (queryMode == "2") {// 按期间查询
		var periodParam = "accountingYearStart={accountingYearStart}&accountingPeriodStart={accountingPeriodStart}&accountingYearEnd={accountingYearEnd}&accountingPeriodEnd={accountingPeriodEnd}";
		periodParam = periodParam.replace("{accountingYearStart}", g_masterFormFieldDict["accountingYearStart"].get("value"));
		periodParam = periodParam.replace("{accountingPeriodStart}", g_masterFormFieldDict["accountingPeriodStart"].get("value"));
		periodParam = periodParam.replace("{accountingYearEnd}", g_masterFormFieldDict["accountingYearEnd"].get("value"));
		periodParam = periodParam.replace("{accountingPeriodEnd}", g_masterFormFieldDict["accountingPeriodEnd"].get("value"));
		url += "&" + periodParam;
	}
	showModalDialog({
		"title": record.get("accountName"),
		"url": url
	});
}

function updateQueryModeWidget() {
	if (document.getElementById("queryByDate").checked) {
		g_masterFormFieldDict["billDateBegin"].set("readonly", false);
		g_masterFormFieldDict["billDateEnd"].set("readonly", false);
		
		g_masterFormFieldDict["accountingYearStart"].set("readonly", true);
		g_masterFormFieldDict["accountingPeriodStart"].set("readonly", true);
		g_masterFormFieldDict["accountingYearEnd"].set("readonly", true);
		g_masterFormFieldDict["accountingPeriodEnd"].set("readonly", true);
	} else {
		g_masterFormFieldDict["accountingYearStart"].set("readonly", false);
		g_masterFormFieldDict["accountingPeriodStart"].set("readonly", false);
		g_masterFormFieldDict["accountingYearEnd"].set("readonly", false);
		g_masterFormFieldDict["accountingPeriodEnd"].set("readonly", false);
		
		g_masterFormFieldDict["billDateBegin"].set("readonly", true);
		g_masterFormFieldDict["billDateEnd"].set("readonly", true);
	}
}

function queryForm() {
	var formManager = new FormManager();
	var bo = formManager.getBo();
	var validateResult = formManager.dsFormValidator(g_dataSourceJson, bo);
	if (!validateResult.result) {
		showError(validateResult.message);
	} else {
		bo["A"]["queryMode"] = getCheckboxValue("queryMode");
		bo["A"]["displayMode"] = getCheckboxValue("displayMode");
		bo["B"] = [];
		ajaxRequest({
			url: "/" + g_dataSourceJson.Id + "/GetData?format=json"
			,params: {
				"dataSourceModelId": g_dataSourceJson.Id,
				"formTemplateId": g_formTemplateJsonData.Id,
				"jsonData": bo
//			,"id": g_id
			},
			callback: function(o) {
				var formManager = new FormManager();
				formManager.applyGlobalParamFromAjaxData(o);
//			formManager.loadData2Form(g_dataSourceJson, o.bo);
//			formManager.setFormStatus(g_formStatus);
				
				var bo = o.bo;
				var modelIterator = new ModelIterator();
				var result = "";
				modelIterator.iterateAllDataSet(g_dataSourceJson, result, function(dataSet, result){
					if (dataSet.Id != "A") {
						if (g_gridPanelDict[dataSet.Id]) {
							if (bo[dataSet.Id] !== undefined) {
								g_gridPanelDict[dataSet.Id].dt.set("data", bo[dataSet.Id]);
							} else {
								g_gridPanelDict[dataSet.Id].dt.set("data", []);
							}
						}
					}
				});
			}
		});
	}
}

function main(Y) {
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
		updateQueryModeWidget();
		/*if (g_id) {
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
		}*/
	}
