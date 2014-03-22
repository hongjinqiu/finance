Y.PTriggerField = Y.Base.create('p-trigger-field', Y.RTriggerField, [Y.WidgetChild], {
    bindUI: function() {
    	Y.PTriggerField.superclass.bindUI.apply(this, arguments);
    	var self = this;
    	
    	// apply value change copy field value,// TODO,
    	this.after('valueChange', Y.bind(function(e) {
			var listTemplateIterator = new ListTemplateIterator();
			var result = "";
			listTemplateIterator.iterateAnyTemplateQueryParameter(result, function(queryParameter, result){
				if (queryParameter.Name == self.get("name")) {
					var queryParameterManager = new QueryParameterManager();
					var formData = queryParameterManager.getQueryFormData();
					
					var relationItem = self._relationFuncTemplate(queryParameter, formData);
					if (relationItem) {
						if (relationItem.CCopyConfigLi) {
							var selectorName = self.get("selectorName")();
							if (self.get("value")) {
								var selectorDict = g_relationManager.getRelationBo(selectorName, self.get("value"));
								if (selectorDict) {
									for (var i = 0; i < relationItem.CCopyConfigLi.length; i++) {
										var copyColumnName = relationItem.CCopyConfigLi[i].CopyColumnName;
										var copyValueField = relationItem.CCopyConfigLi[i].CopyValueField;
										if (g_masterFormFieldDict[copyColumnName]) {
											var valueFieldLi = copyValueField.split(",");
											var valueLi = [];
											for (var j = 0; j < valueFieldLi.length; j++) {
												if (selectorDict[valueFieldLi[j]]) {
													valueLi.push(selectorDict[valueFieldLi[j]]);
												}
											}
											g_masterFormFieldDict[copyColumnName].set("value", valueLi.join(","));
										}
									}
								}
							} else {
								for (var i = 0; i < relationItem.CCopyConfigLi.length; i++) {
									var copyColumnName = relationItem.CCopyConfigLi[i].CopyColumnName;
									if (g_masterFormFieldDict[copyColumnName]) {
										g_masterFormFieldDict[copyColumnName].set("value", "");
									}
								}
							}
						}
					}
					
					return true;
				}
				return false;
			});
    	},
        this));
    	new FormManager().applyEventBehavior(self, Y);
    },

    _validateReadonly: function(val) {
    	var self = this;
    	return new FormManager().validateReadonly(self, val, Y);
    },
    
    initializer: function() {
    	Y.PTriggerField.superclass.initializer.apply(this, arguments);
    	var self = this;
    	
    	new FormManager().initializeAttr(self, Y);
    	
    	// 需要配置在extraInfo里面,
		var selectFunc = function(selectValueLi, formObj){
			
		}
		var unSelectFunc = function(formObj){
			
		}
		var queryFunc = function() {
			return {};
		}
		var multi = false;
		var selectorName = "";
		var displayField = "";
		var valueField = "id";
		var selectorTitle;
		
		self._setDefaultSelectAction();
		
		var modelIterator = new ModelIterator();
    	var result = "";
    	modelIterator.iterateAllField(g_dataSourceJson, result, function(fieldGroup, result){
    		if (fieldGroup.Id == self.get("name") && fieldGroup.getDataSetId() == self.get("dataSetId")) {
				selectFunc = fieldGroup.jsConfig.selectFunc;
				unSelectFunc = fieldGroup.jsConfig.unSelectFunc;
				queryFunc = fieldGroup.jsConfig.queryFunc;
    		}
    	});
    	var templateIterator = new TemplateIterator();
		var result = "";
		var dataSetId = self.get("dataSetId");
    	templateIterator.iterateAnyTemplateColumn(dataSetId, result, function IterateFunc(column, result) {
    		if (column.Name == self.get("name")) {
				selectorName = function() {
					var relationItem = self._relationFuncTemplate(dataSetId, column);
					if (relationItem) {
						return relationItem.CRelationConfig.SelectorName;
					}
					return "";
				}
				displayField = function() {
					var relationItem = self._relationFuncTemplate(dataSetId, column);
					if (relationItem) {
						return relationItem.CRelationConfig.DisplayField;
					}
					return "";
				}
				multi = function() {
					var relationItem = self._relationFuncTemplate(dataSetId, column);
					if (relationItem) {
						return relationItem.CRelationConfig.SelectionMode == "multi";
					}
					return false;
				}
				valueField = function() {
					var relationItem = self._relationFuncTemplate(dataSetId, column);
					if (relationItem) {
						return relationItem.CRelationConfig.ValueField;
					}
					return "";
				}
				selectorTitle = function() {
					var name = selectorName();
					if (name) {
						return g_relationBo[name].Description;
					}
					return "";
				}
    			
    			return true;
    		}
    		return false;
    	});
		
		this.set("multi", multi);
		this.set("selectorName", selectorName);
		this.set("displayField", displayField);
		this.set("valueField", valueField);
		this.set("selectFunc", selectFunc);
		this.set("unSelectFunc", unSelectFunc);
		this.set("queryFunc", queryFunc);
		this.set("selectorTitle", selectorTitle);
    },
    
    _setDefaultSelectAction: function() {
    	var self = this;
		var modelIterator = new ModelIterator();
    	var result = "";
    	modelIterator.iterateAllField(g_dataSourceJson, result, function(fieldGroup, result){
    		if (fieldGroup.Id == self.get("name") && fieldGroup.getDataSetId() == self.get("dataSetId")) {
    			if (!fieldGroup.jsConfig) {
    				fieldGroup.jsConfig = {};
    			}
    			if (!fieldGroup.jsConfig.selectFunc) {
    				fieldGroup.jsConfig.selectFunc = function(selectValueLi, formObj) {
    					if (!selectValueLi || selectValueLi.length == 0) {
    						self._getUnSelectionAction()(self);
    					} else {
    						formObj.set("value", selectValueLi.join(","));
    					}
    				}
    			}
    			if (!fieldGroup.jsConfig.unSelectFunc) {
    				fieldGroup.jsConfig.unSelectFunc = function(formObj) {
    					formObj.set("value", "");
    				}
    			}
    			if (!fieldGroup.jsConfig.queryFunc) {
    				fieldGroup.jsConfig.queryFunc = function() {
						return {};
					}
				}
    		}
    	});
    },
    
    _relationFuncTemplate: function(dataSetId, column) {
    	if (dataSetId == "A") {
    		var formManager = new FormManager();
    		var bo = formManager.getBo();
    		var data = bo["A"];
    		
    		var commonUtil = new CommonUtil();
    		return commonUtil.getCRelationItem(column.CRelationDS, bo, data);
    	} else {
    		// TODO 分录的东东,
    		return null;
    	}
    }
},
{

    ATTRS: {
    	dataSetId: {
            validator: Y.Lang.isString,
            writeOnce: true
        }
    }
});
