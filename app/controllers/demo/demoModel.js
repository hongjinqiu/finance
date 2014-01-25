var demoModel = {
	"A" : {
		"nick" : {
			"jsConfig" : {

			}
		}
	},
	"B" : {
		"attachCount" : {
			"jsConfig" : {
				displayField : "",
				valueField : "",
				selectorName : "",
				selectionMode : "single",
				listeners : {},
				defaultValueExprForJs : function() {
				},
				calcValueExprForJs : function() {
				},
				triggerEditor : function() {
				},
				validator: function() {
					
				}
			}
		}
	}
};

function addRow(dataSetId) {
	gridPanelDict[dataSetId].createAddRowGrid();
}
