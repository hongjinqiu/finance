var modelExtraInfo = {
	"A" : {
		"code" : {
			displayField : "",
			valueField : "",
			selectorName : "",
			selectionMode : "single",
			listeners : {
				focus : function(e, formFieldObj) {
					console.log("A focus");
				}
				,blur : function(e, formFieldObj) {
					console.log("A blur");
				}
				,change : function(e, formFieldObj) {
					console.log("A change");
				}
				,dblclick : function(e, formFieldObj) {
					console.log("A dblclick");
				}
				,keydown : function(e, formFieldObj) {
					console.log("A keydown");
				}
				,
				click : function(e, formFieldObj) {
					console.log("A click");
				}
			},
			defaultValueExprForJs : function() {
			},
			calcValueExprForJs : function() {
			},
			triggerEditor : function() {
			},
			validator : function() {

			}
		},
		"selectTest" : {
			queryFunc: function() {
				return {
					code: "0",
					name: "0"
				};
			}
			/*selectFunc: function(selectValueLi, formObj) {
				console.log("in selection");
				console.log(selectValueLi);
				console.log(formObj);
				console.log("end selection");
			}*/
		}
	},
	"B" : {
		"code" : {
			displayField : "",
			valueField : "",
			selectorName : "",
			selectionMode : "single",
			listeners : {
				focus : function(e, formFieldObj) {
					console.log("focus");
				},
				blur : function(e, formFieldObj) {
					console.log("blur");
				},
				change : function(e, formFieldObj) {
					console.log("change");
				},
				dblclick : function(e, formFieldObj) {
					console.log("dblclick");
				},
				keydown : function(e, formFieldObj) {
					console.log("keydown");
				},
				click : function(e, formFieldObj) {
					console.log("click");
				}
			},
			/*defaultValueExprForJs : function(bo, data) {
				return "this is code in js";
			},
			calcValueExprForJs : function(bo, data) {
				return data["name"] + " re calc test";
				return "this is code in calc value";
			},*/
			triggerEditor : function() {
			},
			validator : function() {

			}
		},
		afterNewData: function(bo, data){
			console.log("after new data");
		} 
		/*,
		defaultValueExprForJs : function(bo, data) {
			return {};
		},// 整个业务对象,单行数据
		calcValueExprForJs : function(bo, data) {
			return {};
		}// 整个业务对象,单行数据,
		*/
	}
};

/**
 * 选择按钮,弹窗选择
 */
function actionTestBSelect(dataSetId) {
//	gridPanelDict[dataSetId].createAddRowGrid();
//	console.log("g_popupSelect");

//	var modelIterator = new ModelIterator();
//	var result = "";
	window.s_selection = function(selectValueLi) {
		var relationManager = new RelationManager();
		var li = [];
		for (var i = 0; i < selectValueLi.length; i++) {
			var data = relationManager.getRelationBo("SysUserSelector", selectValueLi[i]);
			li.push({
				"code": data.code,
				"name": data.name
			});
		}
		gridPanelDict["B"].dt.addRows(li);
	};
	/*
	window.s_queryFunc = function() {
		return {};
	};
	*/
	
    var url = "/console/selectorschema?@name={NAME_VALUE}&@multi={MULTI_VALUE}";
    url = url.replace("{NAME_VALUE}", "SysUserSelector");
    url = url.replace("{MULTI_VALUE}", "true");
    var title = "";
    for (var i = 0; i < dataSourceJson.DetailDataLi.length; i++) {
    	if (dataSourceJson.DetailDataLi[i].Id == dataSetId) {
    		title = dataSourceJson.DetailDataLi[i].DisplayName;
    		break;
    	}
    }
	var dialog = showModalDialog({
		"title": title,
		"url": url
	});
	window.s_closeDialog = function() {
		if (window.s_dialog) {
			window.s_dialog.hide();
		}
		window.s_dialog = null;
		window.s_selection = null;
		window.s_queryFunc = null;
	}
}

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
