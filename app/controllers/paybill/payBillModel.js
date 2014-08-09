var modelExtraInfo = {
		"A" : {
			"property" : {
				listeners : {
					valueChange : function(e, formObj) {
						if (formObj.get("value") == "" || formObj.get("value") == "3") {// 空(请选择),4:其他
							g_masterFormFieldDict["accountId"].set("readonly", true);
						} else {
							g_masterFormFieldDict["accountId"].set("readonly", false);
						}
					}
				}
			},
			"payerType" : {
				listeners : {
					valueChange : function(e, formObj) {
						if (formObj.get("value") == "" || formObj.get("value") == "4") {// 空(请选择),4:其他
							g_masterFormFieldDict["payerId"].set("readonly", true);
						} else {
							g_masterFormFieldDict["payerId"].set("readonly", false);
						}
					}
				}
			}
		},
		"B" : {
			"accountType" : {
				listeners : {
					valueChange : function(e, formObj) {
						var formManager = new FormManager();
						var fieldDict = formManager.getFieldDict(formObj);
						if (fieldDict["accountId"]) {
							if (formObj.get("value") == "") {// 空(请选择)
								fieldDict["accountId"].set("readonly", true);
							} else {
								fieldDict["accountId"].set("readonly", false);
							}
						}
					}
				}
			}
		},
		validate : function(bo, masterMessageLi, detailMessageDict) {
			// 验证联动,
			// 业务属性
			if (bo.A.property == "1" || bo.A.property == "2") {
				var accountIdValue = g_masterFormFieldDict["accountId"].get("value");
				if (!accountIdValue || accountIdValue == "0") {
					masterMessageLi.push("付款账户不允许为空");
					g_masterFormFieldDict["accountId"].set("error", "不允许为空");
				}
			}
			// 收款对象类型
			if (bo.A.payerType == "1" || bo.A.payerType == "2" || bo.A.payerType == "3") {
				var payerIdValue = g_masterFormFieldDict["payerId"].get("value");
				if (!payerIdValue || payerIdValue == "0") {
					masterMessageLi.push("付款对象不允许为空");
					g_masterFormFieldDict["payerId"].set("error", "不允许为空");
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
		if (g_formStatus != "view") {
			modelExtraInfo.A.property.listeners.valueChange(null, g_masterFormFieldDict["property"]);
			modelExtraInfo.A.payerType.listeners.valueChange(null, g_masterFormFieldDict["payerType"]);
		}
	}
