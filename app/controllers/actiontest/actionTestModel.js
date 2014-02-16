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
			/*selection: function(selectValueLi, formObj) {
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
			defaultValueExprForJs : function() {
			},
			calcValueExprForJs : function() {
			},
			triggerEditor : function() {
			},
			validator : function() {

			}
		}
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
		console.log("s_selection");
		console.log(selectValueLi);
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

