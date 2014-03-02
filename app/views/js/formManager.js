function FormManager() {
}

FormManager.prototype.getBo = function() {
	var modelIterator = new ModelIterator();
	var dataSource = dataSourceJson;
	var bo = {};
	var result = "";
	modelIterator.iterateAllField(dataSource, result, function(fieldGroup, result){
		if (fieldGroup.isMasterField()) {
			var formFieldObj = masterFormFieldDict[fieldGroup.Id];
			if (formFieldObj) {
				if (!bo[fieldGroup.getDataSetId()]) {
					bo[fieldGroup.getDataSetId()] = {};
				}
				bo[fieldGroup.getDataSetId()][fieldGroup.Id] = formFieldObj.get("value");
			}
		}
	});
	modelIterator.iterateAllDataSet(dataSource, result, function(dataSet, result){
		var dataSetId = dataSet.Id;
		if (dataSetId != "A") {
			var gridObj = gridPanelDict[dataSetId];
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
	var dataSource = dataSourceJson;
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
	var dataSource = dataSourceJson;
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
	var isUnLimit = fieldGroup.LimitOption == "" || fieldGroup.LimitOption == "unLimit";
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
	} else if (isDataTypeNumber && !isUnLimit) {
		if (!/^-?\d*(\.\d*)?$/.test(value)) {
			messageLi.push("必须由数字小数点组成");
			return messageLi;
		}
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
	modelIterator.iterateAllField(dataSourceJson, result, function(fieldGroup, result){
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
			var formFieldObj = masterFormFieldDict[fieldGroup.Id];
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

FormManager.prototype.setFormStatus = function(status) {
	var self = this;
	g_formStatus = status;
	self._setMasterFormFieldStatus(status);
	self._setDetailGridStatus(status);
	var toolbarManager = new ToolbarManager();
	toolbarManager.enableDisableToolbarBtn();
}

FormManager.prototype._setMasterFormFieldStatus = function(status) {
	for (var key in masterFormFieldDict) {
		masterFormFieldDict[key].set("readonly", status == "view");
	}
}

FormManager.prototype._setDetailGridStatus = function(status) {
	var modelIterator = new ModelIterator();
	var result = "";
	var dataSource = dataSourceJson;
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
			var detailGrid = gridPanelDict[dataSet.Id];
			if (detailGrid) {
				var templateIterator = new TemplateIterator();
				templateIterator.iterateAnyTemplateColumn(dataSet.Id, result, function(column, result){
					if (column.XMLName.Local == "virtual-column") {
						if (status == "view") {
							var virtualColumn = gridPanelDict[dataSet.Id].dt.getColumn(column.Name);
							if (virtualColumn) {
								gridPanelDict[dataSet.Id].virtualColumn = virtualColumn;
								gridPanelDict[dataSet.Id].dt.removeColumn(column.Name);
							}
						} else {
							var virtualColumn = gridPanelDict[dataSet.Id].dt.getColumn(column.Name);
							if (!virtualColumn) {
								virtualColumn = gridPanelDict[dataSet.Id].virtualColumn;
								if (virtualColumn) {
									gridPanelDict[dataSet.Id].dt.addColumn(virtualColumn, 1);
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
	modelIterator.iterateAllFieldBo(dataSource, bo, result, function(fieldGroup, data, rowIndex, result){
		if (fieldGroup.isMasterField()) {
			if (masterFormFieldDict[fieldGroup.Id]) {
				if (data[fieldGroup.Id] !== undefined) {
					masterFormFieldDict[fieldGroup.Id].set("value", data[fieldGroup.Id]);
				} else {
					masterFormFieldDict[fieldGroup.Id].set("value", "");
				}
			}
		}
	});
	modelIterator.iterateAllDataSet(dataSource, result, function(dataSet, result){
		if (dataSet.Id != "A") {
			if (gridPanelDict[dataSet.Id]) {
				if (bo[dataSet.Id] !== undefined) {
					gridPanelDict[dataSet.Id].dt.set("data", bo[dataSet.Id]);
				} else {
					gridPanelDict[dataSet.Id].dt.set("data", []);
				}
			}
		}
	});
}

