/**
 * @class PTriggerField
 * @extends PFormField
 * @param config {Object} Configuration object
 * @constructor
 * @description A form field which allows one or multiple values from a 
 * selection of choices
 */
Y.PTriggerField = Y.Base.create('p-trigger-field', Y.PFormField, [Y.WidgetParent, Y.WidgetChild], {

	FIELD_CLASS : 'table-layout-cell trigger_input',
	INPUT_TYPE: "text",
	SELECT_TEMPLATE: '<a></a>',
	SELECT_CLASS: 'trigger_select',
	VIEW_TEMPLATE: '<a></a>',
	VIEW_CLASS: 'trigger_view',
	DELETE_TEMPLATE: '<a></a>',
	DELETE_CLASS: 'trigger_delete',
	INPUT_NODE_CLASS: 'trigger_input',
	_selectNode: null,
	_viewNode: null,
	_deleteNode: null,
	
	_renderSelectNode : function () {
        this._selectNode = this._renderNode(this.SELECT_TEMPLATE, this.SELECT_CLASS);
    },
    
    _renderViewNode : function () {
        this._viewNode = this._renderNode(this.VIEW_TEMPLATE, this.VIEW_CLASS);
    },
    
    _renderDeleteNode : function () {
        this._deleteNode = this._renderNode(this.DELETE_TEMPLATE, this.DELETE_CLASS);
    },
	
	renderUI: function() {
		Y.PTriggerField.superclass.renderUI.apply(this, arguments);
		
		this._renderSelectNode();
		if (!this.get("multi")) {
			this._renderViewNode();
		}
		this._renderDeleteNode();
    },
    
    _getTempValueCommon: function(modelIteratorFunc, formTempFunc) {
    	var self = this;
    	var selectorName = null;
    	var modelIterator = new ModelIterator();
    	var result = "";
        modelIterator.iterateAllField(dataSourceJson, result, function(fieldGroup, result){
    		if (fieldGroup.Id == self.get("name") && fieldGroup.getDataSetId() == self.get("dataSetId")) {
    			selectorName = modelIteratorFunc(fieldGroup);
    		}
    	});
        if (!selectorName && formTemplateJsonData) {
        	var name = this.get("name");
    		var dataSetId = this.get("dataSetId");
    		for (var i = 0; i < formTemplateJsonData.FormElemLi.length; i++) {
    			var formElem = formTemplateJsonData.FormElemLi[i];
    			if (formElem.XMLName.Local == "column-model") {
    				if (formElem.ColumnModel.DataSetId == dataSetId) {
    					if (formElem.ColumnModel.ColumnLi) {
    						for (var j = 0; j < formElem.ColumnModel.ColumnLi.length; j++) {
    							var column = formElem.ColumnModel.ColumnLi[j];
    							if (column.Name == name) {
    								selectorName = formTempFunc(column);
    								break;
    							}
    						}
    					}
    				}
    			}
    		}
        }
        return selectorName;
    },
    
    _getSelectorId: function() {
    	var modelIteratorFunc = function(fieldGroup){
    		if (fieldGroup.jsConfig && fieldGroup.jsConfig.selectorName) {
				if (typeof(fieldGroup.jsConfig.selectorName) == "function") {
					return fieldGroup.jsConfig.selectorName();
				} else if (typeof(fieldGroup.jsConfig.selectorName) == "string") {
					return fieldGroup.jsConfig.selectorName;
				}
			}
    		return "";
    	};
    	var formTempFunc = function(column) {
    		return column.SelectorName;
    	};
    	return this._getTempValueCommon(modelIteratorFunc, formTempFunc);
    },
    
    _getSelectionMode: function() {
    	var modelIteratorFunc = function(fieldGroup){
    		if (fieldGroup.jsConfig && fieldGroup.jsConfig.selectionMode) {
				if (typeof(fieldGroup.jsConfig.selectionMode) == "function") {
					return fieldGroup.jsConfig.selectionMode();
				} else if (typeof(fieldGroup.jsConfig.selectionMode) == "string") {
					return fieldGroup.jsConfig.selectionMode;
				}
			}
    		return "";
    	};
    	var formTempFunc = function(column) {
    		return column.SelectionMode;
    	};
    	return this._getTempValueCommon(modelIteratorFunc, formTempFunc);
    },
    
    _getDisplayField: function() {
    	var modelIteratorFunc = function(fieldGroup){
    		if (fieldGroup.jsConfig && fieldGroup.jsConfig.displayField) {
				if (typeof(fieldGroup.jsConfig.displayField) == "function") {
					return fieldGroup.jsConfig.displayField();
				} else if (typeof(fieldGroup.jsConfig.displayField) == "string") {
					return fieldGroup.jsConfig.displayField;
				}
			}
    		return "";
    	};
    	var formTempFunc = function(column) {
    		return column.DisplayField;
    	};
    	return this._getTempValueCommon(modelIteratorFunc, formTempFunc);
    },
    
    _getText: function() {
    	var modelIteratorFunc = function(fieldGroup){
    		return "";
    	};
    	var formTempFunc = function(column) {
    		return column.Text;
    	};
    	return this._getTempValueCommon(modelIteratorFunc, formTempFunc);
    },
    
    bindUI: function() {
    	Y.PTriggerField.superclass.bindUI.apply(this, arguments);
    	
    	this._fieldNode.detach('change');
    	this._fieldNode.detach('blur');
    	
    	this.detach('valueChange');
    	
    	this.on('valueChange', Y.bind(function(e) {
            if (e.src != 'ui') {
            	if (!e.newVal) {
            		this._fieldNode.set('value', "");
            		return;
            	}
                //this._fieldNode.set('value', e.newVal + "_测试值");
                var selectorId = this._getSelectorId();
                var relationManager = new RelationManager();
                var relationItem = relationManager.getRelationBo(selectorId, e.newVal);
            	var displayField = this._getDisplayField();
            	var value = "";
            	if (displayField.indexOf("{") > -1) {
            		value = Y.Lang.sub(displayField, relationItem);
            	} else {
            		var keyLi = displayField.split(',');
            		for (var i = 0; i < keyLi.length; i++) {
            			if (relationItem[keyLi[i]]) {
            				value += relationItem[keyLi[i]] + ",";
            			}
            		}
            		if (value) {
            			value = value.substr(0, value.length - 1);
            		}
            	}
            	this._fieldNode.set('value', value);
            }
        },
        this));
    	
    	this._selectNode.on("click", Y.bind(function(e) {
    		var self = this;
    		var modelIterator = new ModelIterator();
        	var result = "";
        	window.s_selection = null;
            modelIterator.iterateAllField(dataSourceJson, result, function(fieldGroup, result){
        		if (fieldGroup.Id == self.get("name") && fieldGroup.getDataSetId() == self.get("dataSetId")) {
        			if (fieldGroup.jsConfig && fieldGroup.jsConfig.selection) {
        				window.s_selection = function(selectValueLi){
        					fieldGroup.jsConfig.selection(selectValueLi, self);
        				};
        			}
        			if (fieldGroup.jsConfig && fieldGroup.jsConfig.queryFunc) {
        				window.s_queryFunc = fieldGroup.jsConfig.queryFunc;
        			}
        		}
        	});
            if (!window.s_selection) {
            	window.s_selection = function(selectValueLi) {
            		self.set("value", selectValueLi.join(","));
//            		console.log(self);
//            		showAlert("选回值:" + selectValueLi.join(","));
            	}
            }
            var url = "/console/selectorschema?@name={NAME_VALUE}&@id={ID_VALUE}&@multi={MULTI_VALUE}&@displayField={DISPLAY_FIELD_VALUE}@entrance=true";
            url = url.replace("{NAME_VALUE}", this._getSelectorId());
            url = url.replace("{ID_VALUE}", this.get('value'));
            url = url.replace("{MULTI_VALUE}", this.get('multi'));
            url = url.replace("{DISPLAY_FIELD_VALUE}", this._getDisplayField());
    		var dialog = showModalDialog({
    			"title": this._getText(),
    			"url": url
    		});
    		window.s_closeDialog = function() {
    			if (window.s_dialog) {
    				window.s_dialog.hide();
    			}
    		}
    	}, this));
    	
    	if (!this.get("multi")) {
    		this._viewNode.on("click", Y.bind(function(e) {
    			showAlert("查看单击");
    		}, this));
    	}
    	
    	this._deleteNode.on("click", Y.bind(function(e) {
    		showAlert("删除单击");
    	}, this));
    },
    
    _syncSelectNode: function() {
    	if (this._selectNode) {
    		this._selectNode.setAttrs({
    			href: "javascript:void(0);",
    			title: "多选",
    			id: this.get('id') + Y.PFormField.FIELD_ID_SUFFIX
    		});
    	}
    },
    
    _syncViewNode: function() {
    	if (this._viewNode) {
    		this._viewNode.setAttrs({
    			href: "javascript:void(0);",
    			title: "查看",
    			id: this.get('id') + Y.PFormField.FIELD_ID_SUFFIX
    		});
    	}
    },
    
    _syncDeleteNode: function() {
    	if (this._deleteNode) {
    		this._deleteNode.setAttrs({
    			href: "javascript:void(0);",
    			title: "删除",
    			id: this.get('id') + Y.PFormField.FIELD_ID_SUFFIX
    		});
    	}
    },
    
    syncUI: function() {
    	Y.PTriggerField.superclass.syncUI.apply(this, arguments);
    	
    	this._syncSelectNode();
    	this._syncViewNode();
    	this._syncDeleteNode();
    },
    
    initializer: function() {
    	Y.PTriggerField.superclass.initializer.apply(this, arguments);
    	
    	var selectionMode = this._getSelectionMode();
    	if (selectonMode == "multi") {
    		this.set("multi", true);
    	} else {
    		this.set("multi", false);
    	}
    },

},
{
    ATTRS: {
        /** 
         * @attribute multi
         * @type Boolean
         * @default false
         * @description Set to true to allow multiple values to be selected
         */
        multi: {
            validator: Y.Lang.isBoolean,
            value: false
        },
        value: {
            value: '',
            validator: function(val) {
                if (!(Y.Lang.isString(val) || Y.Lang.isNumber(val))){
                	return false;
                }
                if (!val) {
                	return true;
                }
                var selectorId = this._getSelectorId();
                var relationManager = new RelationManager();
                var relationBo = relationManager.getRelationBo(selectorId, val);
                if (!relationBo) {
                	return false;
                }
                return true;
            }
        }
    }
});
