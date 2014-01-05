function ColumnManager() {
}

ColumnManager.prototype.currencyFormatFunc = function(o) {
	var formatConfig = null;
	var currencyField = o.column.currencyField;
	if (currencyField != "") {
		var prefix = null;
		var decimalPlaces = null;
		if (o.column.isMoney == "true") {// 是否金额
			if (sysParam[currencyField]) {// 本位币
				prefix = sysParam[currencyField]["prefix"];
				decimalPlaces = sysParam[currencyField]["decimalPlaces"];
			}
			if (o.data[currencyField]) {// 本行记录中是否存在对应币别
				prefix = o.data[currencyField]["prefix"];
				decimalPlaces = o.data[currencyField]["decimalPlaces"];
			}
		} else if (o.column.isUnitPrice == "true") {// 单价
			if (sysParam[currencyField]) {// 本位币
				prefix = sysParam[currencyField]["prefix"];
				decimalPlaces = sysParam[currencyField]["unitPriceDecimalPlaces"];
			}
			if (o.data[currencyField]) {// 本行记录中是否存在对应币别
				prefix = o.data[currencyField]["prefix"];
				decimalPlaces = o.data[currencyField]["unitPriceDecimalPlaces"];
			}
		} else if (o.column.isCost == "true") {// 成本
			if (sysParam[currencyField]) {// 本位币
				prefix = sysParam[currencyField]["prefix"];
				decimalPlaces = sysParam["unitCostDecimalPlaces"];
			}
			if (o.data[currencyField]) {// 本行记录中是否存在对应币别
				prefix = o.data[currencyField]["prefix"];
				decimalPlaces = sysParam["unitCostDecimalPlaces"];
			}
		} else {// 是否金额
			if (sysParam[currencyField]) {// 本位币
				prefix = sysParam[currencyField]["prefix"];
				decimalPlaces = sysParam[currencyField]["decimalPlaces"];
			}
			if (o.data[currencyField]) {// 本行记录中是否存在对应币别
				prefix = o.data[currencyField]["prefix"];
				decimalPlaces = o.data[currencyField]["decimalPlaces"];
			}
		}

		if (prefix !== null) {
			return yInst.DataType.Number.format(o.value, {
				prefix : prefix,
				decimalPlaces : decimalPlaces,
				decimalSeparator : ".",
				thousandsSeparator : sysParam.thousandsSeparator,
				suffix : ""
			});
		} else {
			console.log(o);
			console.log("在系统参数和本行记录中,没有找到currencyField:" + currencyField);
		}
	} else if (o.column.isPercent == "true") {// 本位币
		return yInst.DataType.Number.format(o.value, {
			prefix : "",
			decimalPlaces : sysParam["percentDecimalPlaces"],
			decimalSeparator : ".",
			thousandsSeparator : sysParam.thousandsSeparator,
			suffix : "%"
		});
	}
	return yInst.DataType.Number.format(o.value, {
		//    	prefix            : o.column.prefix     || '￥',
		//		decimalPlaces     : o.column.decimalPlaces      || 2,
		//		decimalSeparator  : o.column.decimalSeparator   || '.',
		//		thousandsSeparator: o.column.thousandsSeparator || ',',
		//		suffix            : o.column.suffix || ''
		prefix : o.column.prefix,
		decimalPlaces : o.column.decimalPlaces,
		decimalSeparator : o.column.decimalSeparator,
		thousandsSeparator : o.column.thousandsSeparator,
		suffix : o.column.suffix
	});
}

ColumnManager.prototype.createIdColumn = function(columnModel) {
	if (columnModel.IdColumn.Hideable != "true") {
		return {
			key: columnModel.IdColumn.Name,
			label: columnModel.IdColumn.Text
		};
	}
	return null;
}

ColumnManager.prototype.createCheckboxColumn = function(columnModel) {
	if (columnModel.CheckboxColumn.Hideable != "true") {
		var key = columnModel.CheckboxColumn.Name;
		if (columnModel.SelectionMode == "radio") {
			return {
				key:        key,
				allowHTML:  true, // to avoid HTML escaping
				label:      '选择',
				//formatter:      '<input type="radio" name="' + key + '" />'
				formatter:function(o) {
					if (o.value) {
						return '<input type="radio" name="' + key + '" />';
					}
					return "";
				}
				//,emptyCellValue: '<input type="checkbox"/>'
			};
		} else {
			return {
				key:        key,
				allowHTML:  true, // to avoid HTML escaping
				label:      '<input type="checkbox" class="protocol-select-all" title="全部选中"/>',
				//formatter:      '<input type="checkbox" />'
				formatter:function(o) {
					if (o.value) {
						return '<input type="checkbox" />';
					}
					return "";
				}
				//,emptyCellValue: '<input type="checkbox"/>'
			};
		}
	}
	return null;
}

ColumnManager.prototype.createVirtualColumn = function(render, columnModel, columnIndex) {
	var i = columnIndex;
	if (columnModel.ColumnLi[i].XMLName.Local == "virtual-column" && columnModel.ColumnLi[i].Hideable != "true") {
		var virtualColumn = columnModel.ColumnLi[i];
		return {
			key: columnModel.ColumnLi[i].Name,
			label: columnModel.ColumnLi[i].Text,
			allowHTML:  true, // to avoid HTML escaping
			formatter:      function(virtualColumn){
				return function(o){
					var htmlLi = [];
					var buttonBoLi = o.value[virtualColumn.Buttons.XMLName.Local];
					for (var j = 0; j < virtualColumn.Buttons.ButtonLi.length; j++) {
						var btnTemplate = null;
						if (virtualColumn.Buttons.ButtonLi[j].Mode == "fn") {
							var gridRenderId = render.replace("#", "");
							btnTemplate = "<input type='button' value='{value}' onclick='doVirtualColumnBtnAction(\"{gridRenderId}\", this, {handler})' class='{class}' />";
						} else if (virtualColumn.Buttons.ButtonLi[j].Mode == "url") {
							btnTemplate = "<input type='button' value='{value}' onclick='location.href=\"{href}\"' class='{class}' />";
						} else {
							btnTemplate = "<input type='button' value='{value}' onclick='window.open(\"{href}\")' class='{class}' />";
						}
						if (buttonBoLi[j]["isShow"]) {
							YUI().use("node", function(Y){
								// handler进行值的预替换,
								var handler = virtualColumn.Buttons.ButtonLi[j].Handler;
								handler = Y.Lang.sub(handler, o.data);
								htmlLi.push(Y.Lang.sub(btnTemplate, {
									value: virtualColumn.Buttons.ButtonLi[j].Text,
									handler: handler,
									"class": virtualColumn.Buttons.ButtonLi[j].IconCls,
									href: handler,
									gridRenderId: gridRenderId
								}));
							});
						}
					}
					return htmlLi.join("");
				}
			}(virtualColumn)
		};
	}
	return null;
}

ColumnManager.prototype.createNumberColumn = function(columnConfig) {
	var decimalPlaces = 2;
	if (columnConfig.DecimalPlaces) {
		decimalPlaces = parseInt(columnConfig.DecimalPlaces);
	}
	var isFormatter = columnConfig.Prefix != "";
	isFormatter = isFormatter || columnConfig.DecimalPlaces != "";
	isFormatter = isFormatter || columnConfig.DecimalSeparator != "";
	isFormatter = isFormatter || columnConfig.ThousandsSeparator != "";
	isFormatter = isFormatter || columnConfig.Suffix != "";
	
	// 财务相关字段的判断,以决定是否用 formatter 函数,
	isFormatter = isFormatter || columnConfig.CurrencyField != "";
	isFormatter = isFormatter || columnConfig.IsPercent != "";
	
	if (isFormatter) {
		var columnManager = new ColumnManager();
		return {
			key: columnConfig.Name,
			label: columnConfig.Text,
			formatter: columnManager.currencyFormatFunc,
			
			prefix: columnConfig.Prefix,
			decimalPlaces: decimalPlaces,
			decimalSeparator: columnConfig.DecimalSeparator,
			thousandsSeparator: columnConfig.ThousandsSeparator,
			suffix: columnConfig.Suffix,
			
			currencyField: columnConfig.CurrencyField,
			isPercent: columnConfig.IsPercent,
			isMoney: columnConfig.IsMoney,
			isUnitPrice: columnConfig.IsUnitPrice,
			isCost: columnConfig.IsCost
		};
	}
	return {
		key: columnConfig.Name,
		label: columnConfig.Text
	};
}

ColumnManager.prototype.convertDate2DisplayPattern = function(displayPattern) {
	displayPattern = displayPattern.replace("yyyy", "%G");
	displayPattern = displayPattern.replace("MM", "%m");
	displayPattern = displayPattern.replace("dd", "%d");
	displayPattern = displayPattern.replace("HH", "%H");
	displayPattern = displayPattern.replace("mm", "%M");
	displayPattern = displayPattern.replace("ss", "%S");
	return displayPattern; 
}

/*
DisplayPattern string `xml:"displayPattern,attr"`
DbPattern      string `xml:"dbPattern,attr"`
*/
ColumnManager.prototype.createDateColumn = function(columnConfig) {
	var dbPattern = columnConfig.DbPattern;
	var displayPattern = columnConfig.DisplayPattern;
	if (dbPattern && displayPattern) {
		return {
			key: columnConfig.Name,
			label: columnConfig.Text,
			dbPattern: dbPattern,
			displayPattern: displayPattern,
			formatter: function(o) {
				if (o.value !== undefined && o.value !== null) {
					var date = new Date();
					var value = o.value + "";
					if (o.column.dbPattern.indexOf("yyyy") > -1) {
						var start = o.column.dbPattern.indexOf("yyyy");
						var end = o.column.dbPattern.indexOf("yyyy") + "yyyy".length;
						var yyyy = value.substring(start, end);
						date.setYear(parseInt(yyyy));
					}
					if (o.column.dbPattern.indexOf("MM") > -1) {
						var start = o.column.dbPattern.indexOf("MM");
						var end = o.column.dbPattern.indexOf("MM") + "MM".length;
						var mm = value.substring(start, end);
						date.setMonth(parseInt(mm) - 1);
					}
					if (o.column.dbPattern.indexOf("dd") > -1) {
						var start = o.column.dbPattern.indexOf("dd");
						var end = o.column.dbPattern.indexOf("dd") + "dd".length;
						var dd = value.substring(start, end);
						date.setDate(parseInt(dd));
					}
					if (o.column.dbPattern.indexOf("HH") > -1) {
						var start = o.column.dbPattern.indexOf("HH");
						var end = o.column.dbPattern.indexOf("HH") + "HH".length;
						var hh = value.substring(start, end);
						date.setHours(parseInt(hh));
					}
					if (o.column.dbPattern.indexOf("mm") > -1) {
						var start = o.column.dbPattern.indexOf("mm");
						var end = o.column.dbPattern.indexOf("mm") + "mm".length;
						var mm = value.substring(start, end);
						date.setMinutes(mm);
					}
					if (o.column.dbPattern.indexOf("ss") > -1) {
						var start = o.column.dbPattern.indexOf("ss");
						var end = o.column.dbPattern.indexOf("ss") + "ss".length;
						var ss = value.substring(start, end);
						date.setSeconds(ss);
					}
					// js格式参考 http://yuilibrary.com/yui/docs/api/classes/Date.html#method_format
					var columnManager = new ColumnManager();
					var displayPattern = columnManager.convertDate2DisplayPattern(o.column.displayPattern);
					return yInst.DataType.Date.format(date, {
						format: displayPattern
					});
				}
				return o.value;
			}
		};
	} else {
		console.log(columnConfig);
		console.log("日期字段未同时配置dbPattern和displayPattern");
	}
	return {
		key: columnConfig.Name,
		label: columnConfig.Text
	};
}

ColumnManager.prototype.createBooleanColumn = function(columnConfig) {
	return {
		key: columnConfig.Name,
		label: columnConfig.Text,
		formatter: function(o) {
			if (o.value + "" == "true") {
				return "是";
			} else if (o.value + "" == "false") {
				return "否";
			}
			return o.value;
		}
	};
}

ColumnManager.prototype.createDictionaryColumn = function(columnConfig) {
	return {
		key: columnConfig.Name,
		label: columnConfig.Text,
		formatter: function(o) {
			var dictionaryValue = o.data[columnConfig.Name + "_DICTIONARY_NAME"];
			if (dictionaryValue !== undefined && dictionaryValue !== null) {
				return dictionaryValue;
			}
			console.log(columnConfig);
			console.log("字典字段没找到_DICTIONARY_NAME,code:" + o.value);
			return o.value;
		}
	};
}

ColumnManager.prototype.createRowIndexColumn = function(columnModel) {
	if (columnModel.Rownumber == "true") {
		return {
			key: "",
			label: "序号",
			formatter: function(o) {
				return o.rowIndex + 1;
			}
		};
	}
	return null;
}

ColumnManager.prototype.createColumn = function(columnConfig) {
	var self = this;
	if (columnConfig.XMLName.Local != "virtual-column" && columnConfig.Hideable != "true") {
		if (columnConfig.ColumnModel.ColumnLi) {
			var result = {
				label: columnConfig.Text,
				"children": []
			};
			for (var i = 0; i < columnConfig.ColumnModel.ColumnLi.length; i++) {
				var childColumn = self.createColumn(columnConfig.ColumnModel.ColumnLi[i]);
				if (childColumn) {
					result.children.push(childColumn);
				}
			}
			return result;
		}
		
		if (columnConfig.XMLName.Local == "number-column") {
			return self.createNumberColumn(columnConfig);
		} else if (columnConfig.XMLName.Local == "date-column") {
			return self.createDateColumn(columnConfig);
		} else if (columnConfig.XMLName.Local == "boolean-column") {
			return self.createBooleanColumn(columnConfig);
		} else if (columnConfig.XMLName.Local == "dictionary-column") {
			return self.createDictionaryColumn(columnConfig);
		}
		return {
			key: columnConfig.Name,
			label: columnConfig.Text
		};
	}
	return null;
}

ColumnManager.prototype.getColumns = function(render, columnModel, Y) {
	var self = this;
	var columns = [];
	var checkboxColumn = self.createCheckboxColumn(columnModel);
	if (checkboxColumn) {
		columns.push(checkboxColumn);
	}
	var idColumn = self.createIdColumn(columnModel);
	if (idColumn) {
		columns.push(idColumn);
	}
	var rowIndexColumn = self.createRowIndexColumn(columnModel);
	if (rowIndexColumn) {
		columns.push(rowIndexColumn);
	}
	
	for (var i = 0; i < columnModel.ColumnLi.length; i++) {
		var column = self.createColumn(columnModel.ColumnLi[i]);
		if (column) {
			columns.push(column);
		} else {
			var virtualColumn = self.createVirtualColumn(render, columnModel, i);
			if (virtualColumn) {
				columns.push(virtualColumn);
			}
		}
	}
	return columns;
}



