var listTemplateExtraInfo = {
	"ColumnModel" : {

	},
	"QueryParameter" : {
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
	}
};

function main(Y) {
	
		listTemplateExtraInfo.QueryParameter.property.listeners.valueChange(null, g_masterFormFieldDict["property"]);
		listTemplateExtraInfo.QueryParameter.payerType.listeners.valueChange(null, g_masterFormFieldDict["payerType"]);
}

