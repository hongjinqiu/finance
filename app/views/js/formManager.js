function FormManager() {
}

FormManager.prototype.validateReadonly = function(formObj, val, Y) {
	var self = formObj;
	if (!Y.Lang.isBoolean(val)) {
		return false;
	}
	var templateIterator = new TemplateIterator();
	var result = "";
	var dataSetId = self.get("dataSetId");
	var validateResult = true;
	templateIterator.iterateAnyTemplateColumn(dataSetId, result, function(column, result){
		if (column.Name == self.get("name")) {
			if (column.FixReadOnly == "true" && !val) {
				validateResult = false;
			}
			return true;
		}
		return false;
	});
	return validateResult;
}

FormManager.prototype.initializeAttr = function(formObj, Y) {
	var self = formObj;
	if (g_dataSourceJson) {
    	var modelIterator = new ModelIterator();
    	var result = "";
    	modelIterator.iterateAllField(g_dataSourceJson, result, function(fieldGroup, result){
    		if (fieldGroup.Id == self.get("name") && fieldGroup.getDataSetId() == self.get("dataSetId")) {
    			if (fieldGroup.AllowEmpty != "true") {
    				self.set("required", true);
    			}
    		}
    	});
    	
    	var templateIterator = new TemplateIterator();
    	var result = "";
    	var dataSetId = self.get("dataSetId");
    	templateIterator.iterateAnyTemplateColumn(dataSetId, result, function(column, result){
    		if (column.Name == self.get("name")) {
    			if (column.FixReadOnly == "true") {
    				self.set("readonly", true);
    			} else if (column.ReadOnly == "true") {
    				self.set("readonly", true);
    			}
    			return true;
    		}
    		return false;
    	});
    	
    	var formManager = new FormManager();
    	self.set("validator", formManager.dsFormFieldValidator);
    }
}

FormManager.prototype.applyEventBehavior = function(formObj, Y) {
	var self = formObj;
	
	var dataSetId = self.get("dataSetId");
	var name = self.get("name");
	// 应用上js相关的操作,
    var modelIterator = new ModelIterator();
	var result = "";
	modelIterator.iterateAllField(g_dataSourceJson, result, function(fieldGroup, result){
		if (fieldGroup.Id == self.get("name") && fieldGroup.getDataSetId() == self.get("dataSetId")) {
			if (fieldGroup.jsConfig && fieldGroup.jsConfig.listeners) {
				for (var key in fieldGroup.jsConfig.listeners) {
					if (key == "valueChange") {
						self.after("valueChange", function(key) {
							return function(e) {
								fieldGroup.jsConfig.listeners[key](e, self);
							}
						}(key));
					} else {
						self._fieldNode.on(key, function(key) {
							return function(e) {
								fieldGroup.jsConfig.listeners[key](e, self);
							}
						}(key));
					}
				}
			}
		}
	});
	// observe的添加,主要用于清值,如果是用tree需要联动呢?到时再添加呗
	var templateIterator = new TemplateIterator();
	var result = "";
	templateIterator.iterateAnyTemplateColumn(dataSetId, result, function(column, result){
		if (column.Name == name) {
			if (column.ColumnAttributeLi) {
				for (var i = 0; i < column.ColumnAttributeLi.length; i++) {
					if (column.ColumnAttributeLi[i].Name == "observe") {
						var observeFields = column.ColumnAttributeLi[i].Value.split(",");
						if (dataSetId == "A") {
							self.after("valueChange", function() {
								for (var j = 0; j < observeFields.length; j++) {
									var targetObj = g_masterFormFieldDict[observeFields[j]];
									if (targetObj) {
										targetObj.set("value", "");
									}
								}
							});
						} else {
							self.after("valueChange", function() {
								if (g_gridPanelDict[dataSetId + "_addrow"]) {
									var formFieldDict = g_gridPanelDict[dataSetId + "_addrow"].dt.getRecord(self._fieldNode).formFieldDict;
									for (var j = 0; j < observeFields.length; j++) {
										var targetObj = formFieldDict[observeFields[j]];
										if (targetObj) {
											targetObj.set("value", "");
										}
									}
								}
							});
						}
					}
				}
			}
			return true;
		}
		return false;
	});
}

FormManager.prototype.setChoices = function(formObj) {
	var self = formObj;
	var choices = [];
	if (g_layerBoLi) {
		var templateIterator = new TemplateIterator();
    	var result = "";
    	var dataSetId = self.get("dataSetId");
    	templateIterator.iterateAnyTemplateColumn(dataSetId, result, function(column, result){
    		if (column.Name == self.get("name")) {
    			if (g_layerBoLi[column.Dictionary]) {
					for (var k = 0; k < g_layerBoLi[column.Dictionary].length; k++) {
						var dictionaryItem = g_layerBoLi[column.Dictionary][k];
						choices.push({
							value: dictionaryItem["code"],
							label: dictionaryItem["name"]
						});
					}
				}
    			return true;
    		}
    		return false;
    	});
	}
	self.set("choices", choices);
}

FormManager.prototype.getBo = function() {
	var modelIterator = new ModelIterator();
	var dataSource = g_dataSourceJson;
	var bo = {"A": {}};
	var result = "";
	for (var key in g_masterFormFieldDict) {
		var formFieldObj = g_masterFormFieldDict[key];
		if (formFieldObj) {
			bo["A"][key] = formFieldObj.get("value");
		}
	}
	modelIterator.iterateAllDataSet(dataSource, result, function(dataSet, result){
		var dataSetId = dataSet.Id;
		if (dataSetId != "A") {
			var gridObj = g_gridPanelDict[dataSetId];
			if (gridObj) {
				bo[dataSetId] = gridObj.dt.get("data").toJSON();
			}
		}
	});
	if (bo["A"] && bo["A"]["id"]) {
		bo["_id"] = bo["A"]["id"];
		bo["id"] = bo["A"]["id"];
	} else {
		bo["_id"] = "";
		bo["id"] = "";
	}
	return bo;
}

FormManager.prototype.getDataSetNewData = function(dataSetId) {
	var self = this;
	var dataSource = g_dataSourceJson;
	var modelTemplateFactory = new ModelTemplateFactory();
	var bo = self.getBo();
	var data = {};
	modelTemplateFactory.applyDataSetDefaultValue(dataSource, dataSetId, bo, data);
	modelTemplateFactory.applyDataSetCalcValue(dataSource, dataSetId, bo, data);
	
	var result = "";
	var modelIterator = new ModelIterator();
	modelIterator.iterateAllDataSet(dataSource, result, function(dataSet, result){
		if (dataSet.Id == dataSetId) {
			if (dataSet.jsConfig && dataSet.jsConfig.afterNewData) {
				dataSet.jsConfig.afterNewData(bo, data);
			}
		}
	});
	data["id"] = modelTemplateFactory.getSequenceNo();
	return data;
}

FormManager.prototype.getDataSetCopyData = function(dataSetId, srcData) {
	var self = this;
	var dataSource = g_dataSourceJson;
	var modelTemplateFactory = new ModelTemplateFactory();
	var bo = self.getBo();
	var destData = {};
	modelTemplateFactory.applyDataSetDefaultValue(dataSource, dataSetId, bo, destData);
	modelTemplateFactory.applyDataSetCopyValue(dataSource, dataSetId, srcData, destData);
	modelTemplateFactory.applyDataSetCalcValue(dataSource, dataSetId, bo, destData);
	
	var result = "";
	var modelIterator = new ModelIterator();
	modelIterator.iterateAllDataSet(dataSource, result, function(dataSet, result){
		if (dataSet.Id == dataSetId) {
			if (dataSet.jsConfig && dataSet.jsConfig.afterNewData) {
				dataSet.jsConfig.afterNewData(bo, destData);
			}
		}
	});
	destData["id"] = modelTemplateFactory.getSequenceNo();
	return destData;
}

/**
 * 数据源字段 fieldGroup 的验证器,返回messageLi
 * 其中,日期控件传的是 input 框里面的值,而不是value,日期控件,get("value")时,其取回的是yyyyMMdd,
 * @param value
 * @param fieldGroup
 */
FormManager.prototype.dsFieldGroupValidator = function(value, dateSeperator, fieldGroup) {
	var messageLi = [];
	if (fieldGroup.AllowEmpty != "true") {
		if (value === "" || value === null || value === undefined) {
			messageLi.push("不允许为空");
			return messageLi;
		}
	}
	
	var isDataTypeNumber = false;
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "DECIMAL";
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "FLOAT";
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "INT";
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "LONGINT";
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "MONEY";
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "SMALLINT";
	var isUnLimit = fieldGroup.LimitOption == undefined || fieldGroup.LimitOption == "" || fieldGroup.LimitOption == "unLimit";
	var dateEnumLi = ["YEAR","YEARMONTH","DATE","TIME","DATETIME"];
	var isDate = false;
	for (var i = 0; i < dateEnumLi.length; i++) {
		if (dateEnumLi[i] == fieldGroup.FieldNumberType) {
			isDate = true;
			break;
		}
	}
	if (isDataTypeNumber && isDate) {
		var isAllowEmptyAndZero = fieldGroup.AllowEmpty == "true" && value == "0";
		if (fieldGroup.FieldNumberType == "YEAR") {
			if (!/^\d{4}$/.test(value) && !isAllowEmptyAndZero) {
				messageLi.push("格式错误，正确格式类似于：1970");
				return messageLi;
			}
		} else if (fieldGroup.FieldNumberType == "YEARMONTH") {
			var message = "";
			if (dateSeperator == "-") {
				message = "格式错误，正确格式类似于：1970-01";
			} else {
				message = "格式错误，正确格式类似于：1970/01";
			}
			if (!/^\d{4}\d{2}$/.test(value) && !isAllowEmptyAndZero) {
				messageLi.push(message);
				return messageLi;
			}
		} else if (fieldGroup.FieldNumberType == "DATE") {
			var message = "";
			if (dateSeperator == "-") {
				message = "格式错误，正确格式类似于：1970-01-02";
			} else {
				message = "格式错误，正确格式类似于：1970/01/02";
			}
			if (!/^\d{4}\d{2}\d{2}$/.test(value) && !isAllowEmptyAndZero) {
				messageLi.push(message);
				return messageLi;
			}
		} else if (fieldGroup.FieldNumberType == "TIME") {
			if (!/^\d{2}\d{2}\d{2}$/.test(value) && !isAllowEmptyAndZero) {
				messageLi.push("格式错误，正确格式类似于：03:04:05");
				return messageLi;
			}
		} else if (fieldGroup.FieldNumberType == "DATETIME") {
			var message = "";
			if (dateSeperator == "-") {
				message = "格式错误，正确格式类似于：1970-01-02 03:04:05";
			} else {
				message = "格式错误，正确格式类似于：1970/01/02 03:04:05";
			}
			if (!/^\d{4}\d{2}\d{2}\d{2}\d{2}\d{2}$/.test(value) && !isAllowEmptyAndZero) {
				messageLi.push(message);
				return messageLi;
			}
		}
	} else if (isDataTypeNumber) {
		if (fieldGroup.Id != "id" && !/^-?\d*(\.\d*)?$/.test(value)) {
			messageLi.push("必须由数字小数点组成");
			return messageLi;
		}
		if (!isUnLimit) {
			var fieldValueFloat = parseFloat(value);
			if (fieldGroup.LimitOption == "limitMax") {
				var maxValue = parseFloat(fieldGroup.LimitMax);
				if (maxValue < fieldValueFloat) {
					messageLi.push("超出最大值" + fieldGroup.LimitMax);
				}
			} else if (fieldGroup.LimitOption == "limitMin") {
				var minValue = parseFloat(fieldGroup.LimitMin);
				if (fieldValueFloat < minValue) {
					messageLi.push("小于最小值" + fieldGroup.LimitMin);
				}
			} else if (fieldGroup.LimitOption == "limitRange") {
				var minValue = parseFloat(fieldGroup.LimitMin);
				var maxValue = parseFloat(fieldGroup.LimitMax);
				if (fieldValueFloat < minValue || maxValue < fieldValueFloat) {
					messageLi.push("超出范围("+fieldGroup.LimitMin+"~"+fieldGroup.LimitMax+")");
				}
			}
		}
	} else {
		var isDataTypeString = false;
		isDataTypeString = isDataTypeString || fieldGroup.FieldDataType == "STRING";
		isDataTypeString = isDataTypeString || fieldGroup.FieldDataType == "REMARK";
		var isFieldLengthLimit = fieldGroup.FieldLength != "";
		if (isDataTypeString && isFieldLengthLimit) {
			var limit = parseFloat(fieldGroup.FieldLength);
			if (value.length > limit) {
				messageLi.push("长度超出最大值"+fieldGroup.FieldLength);
			}
		}
	}
	return messageLi;
}

/**
 * datasource field validator
 */
FormManager.prototype.dsFormFieldValidator = function(value, formFieldObj) {
	var self = this;
	var modelIterator = new ModelIterator();
	var messageLi = [];
	var result = "";
	var formManager = new FormManager();
	modelIterator.iterateAllField(g_dataSourceJson, result, function(fieldGroup, result){
		if (fieldGroup.Id == formFieldObj.get("name") && fieldGroup.getDataSetId() == formFieldObj.get("dataSetId")) {
			var dateSeperator = formManager._getDateSeperator(fieldGroup.getDataSetId(), fieldGroup.Id);
			messageLi = formManager.dsFieldGroupValidator(value, dateSeperator, fieldGroup);
		}
	});
	
	if (messageLi.length > 0) {
		formFieldObj.set("error", messageLi.join("<br />"));
		return false;
	}
	
	return true;
}

FormManager.prototype._getDateSeperator = function(dataSetId, name) {
	var dateSeperator = null;
	var templateIterator = new TemplateIterator();
	var result = "";
	templateIterator.iterateAnyTemplateColumn(dataSetId, result, function(column, result){
		if (column.Name == name) {
			if (column.XMLName.Local == "date-column") {
				if (column.DisplayPattern.indexOf("-") > -1) {
					dateSeperator = "-";
				} else if (column.DisplayPattern.indexOf("/") > -1) {
					dateSeperator = "/";
				}
			}
			return true;
		}
		return false;
	});
	return dateSeperator;
}

FormManager.prototype.dsFormValidator = function(dataSource, bo) {
	var modelIterator = new ModelIterator();
	var messageLi = [];
	var result = "";
	var formManager = new FormManager();
	modelIterator.iterateAllFieldBo(dataSource, bo, result, function(fieldGroup, data, rowIndex, result){
		if (fieldGroup.isMasterField()) {
			var formFieldObj = g_masterFormFieldDict[fieldGroup.Id];
			var value = data[fieldGroup.Id];
			if (value !== undefined && formFieldObj) {
				if(!formManager.dsFormFieldValidator(value, formFieldObj)) {
					messageLi.push(fieldGroup.DisplayName + formFieldObj.get("error"));
				}
			}
		} else {
			var value = data[fieldGroup.Id];
			if (value !== undefined) {
				var dateSeperator = formManager._getDateSeperator(fieldGroup.getDataSetId(), fieldGroup.Id);
				var lineMessageLi = formManager.dsFieldGroupValidator(value, dateSeperator, fieldGroup);
				if (lineMessageLi.length > 0) {
					messageLi.push("序号为" + (rowIndex + 1) + "的分录，" + fieldGroup.DisplayName + lineMessageLi.join("，"));
				}
			}
		}
	});
	if (messageLi.length > 0) {
		return {
			"result": false,
			"message": messageLi.join("<br />")
		};
	}
	return {
		"result": true
	};
}

FormManager.prototype.dsDetailValidator = function(dataSource, dataSetId, detailDataLi) {
	var bo = {};
	bo[dataSetId] = detailDataLi;
	var modelIterator = new ModelIterator();
	var result = "";
	modelIterator.iterateAllDataSet(dataSource, result, function(dataSet, result){
		if (dataSet == "A") {
			bo["A"] = {};
		} else if (!bo[dataSet]) {
			bo[dataSet] = [];
		}
	});
	
	var messageLi = [];
	var formManager = new FormManager();
	modelIterator.iterateAllFieldBo(dataSource, bo, result, function(fieldGroup, data, rowIndex, result){
		if (!fieldGroup.isMasterField() && fieldGroup.getDataSetId() == dataSetId) {
			var formFieldDict = g_gridPanelDict[dataSetId + "_addrow"].dt.getRecord(rowIndex).formFieldDict;
			var formFieldObj = formFieldDict[fieldGroup.Id];
			var value = data[fieldGroup.Id];
			if (value !== undefined && formFieldObj) {
				if(!formManager.dsFormFieldValidator(value, formFieldObj)) {
					//messageLi.push(fieldGroup.DisplayName + formFieldObj.get("error"));
					messageLi.push("序号为" + (rowIndex + 1) + "的分录，" + fieldGroup.DisplayName + formFieldObj.get("error"));
				}
			}
		}
	});
	if (messageLi.length > 0) {
		return {
			"result": false,
			"message": messageLi.join("<br />")
		};
	}
	return {
		"result": true
	};
}

FormManager.prototype.setFormStatus = function(status) {
	var self = this;
	g_formStatus = status;
	self._setMasterFormFieldStatus(status);
	self._setDetailGridStatus(status);
	var toolbarManager = new ToolbarManager();
	toolbarManager.enableDisableToolbarBtn();
}

FormManager.prototype._setMasterFormFieldStatus = function(status) {
	for (var key in g_masterFormFieldDict) {
		g_masterFormFieldDict[key].set("readonly", status == "view");
	}
}

FormManager.prototype._setDetailGridStatus = function(status) {
	var modelIterator = new ModelIterator();
	var result = "";
	var dataSource = g_dataSourceJson;
	modelIterator.iterateAllDataSet(dataSource, result, function(dataSet, result){
		if (dataSet.Id != "A") {
			var tbar = document.getElementById(dataSet.Id + "_tbar");
			if (tbar) {
				if (status == "view") {
					tbar.style.display = "none";
				} else {
					tbar.style.display = "";
				}
			}
			var detailGrid = g_gridPanelDict[dataSet.Id];
			if (detailGrid) {
				var templateIterator = new TemplateIterator();
				templateIterator.iterateAnyTemplateColumn(dataSet.Id, result, function(column, result){
					if (column.XMLName.Local == "virtual-column") {
						if (status == "view") {
							var virtualColumn = g_gridPanelDict[dataSet.Id].dt.getColumn(column.Name);
							if (virtualColumn) {
								g_gridPanelDict[dataSet.Id].virtualColumn = virtualColumn;
								g_gridPanelDict[dataSet.Id].dt.removeColumn(column.Name);
							}
						} else {
							var virtualColumn = g_gridPanelDict[dataSet.Id].dt.getColumn(column.Name);
							if (!virtualColumn) {
								virtualColumn = g_gridPanelDict[dataSet.Id].virtualColumn;
								if (virtualColumn) {
									g_gridPanelDict[dataSet.Id].dt.addColumn(virtualColumn, 1);
								}
							}
						}
						return true;
					}
					return false;
				});
			}
		}
	});
}

FormManager.prototype.loadData2Form = function(dataSource, bo) {
	var modelIterator = new ModelIterator();
	var result = "";
	if (bo["A"]) {
		for (var key in bo["A"]) {
			if (g_masterFormFieldDict[key]) {
				g_masterFormFieldDict[key].set("value", bo["A"][key] || "");
			}
		}
	}
	modelIterator.iterateAllDataSet(dataSource, result, function(dataSet, result){
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

