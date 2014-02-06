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

function addRow(dataSetId) {
	gridPanelDict[dataSetId].createAddRowGrid();
}
