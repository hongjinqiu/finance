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
		"chamberlainType" : {
			listeners : {
				valueChange : function(e, formObj) {
					if (formObj.get("value") == "" || formObj.get("value") == "4") {// 空(请选择),4:其他
						g_masterFormFieldDict["chamberlainId"].set("readonly", true);
					} else {
						g_masterFormFieldDict["chamberlainId"].set("readonly", false);
					}
				}
			}
		}
	}
};

function main() {
	YUI(g_financeModule).use("finance-module", function(YNotUse) {// 不能直接在父函数用use finance-module,会报错,因为在js父函数直接加载,其会直接使用调用
		listTemplateExtraInfo.QueryParameter.property.listeners.valueChange(null, g_masterFormFieldDict["property"]);
		listTemplateExtraInfo.QueryParameter.chamberlainType.listeners.valueChange(null, g_masterFormFieldDict["chamberlainType"]);
	});
}

