function QueryParameterManager() {}

QueryParameterManager.prototype.applyQueryDefaultValue = function() {
	YUI(g_financeModule).use("finance-module", function(Y){
		if (g_defaultBo) {
			for (var key in g_defaultBo) {
				var field = g_masterFormFieldDict[key];
				if (field) {
					field.set("value", g_defaultBo[key]);
				}
			}
		}
	});
}

QueryParameterManager.prototype.applyQueryParameter = function() {
	if (true) {
		return;
	}
	YUI(g_financeModule).use("finance-module", function(Y){
		var columnManager = new ColumnManager();
		for (var i = 0; i < listTemplate.QueryParameterGroup.QueryParameterLi.length; i++) {
			var queryParameterManager = new QueryParameterManager();
			var queryParameter = listTemplate.QueryParameterGroup.QueryParameterLi[i];
			if (queryParameter.Editor == "numberfield") {
				Y.one("#" + queryParameter.Name).on("keypress", function(e) {
					if (!(e.keyCode == 9 || e.keyCode == 13 || e.keyCode == 46 || e.keyCode == 116 || e.keyCode == 118 || (e.keyCode >= 48 && e.keyCode <=57))) {// 0-9.,118:ctrl + v, 116:Ctrl + F5,13: enter,9: tab
						e.preventDefault();
					}
				});
			} else if (queryParameter.Editor == "datefield") {
				var dateFormat = null;
				for (var j = 0; j < queryParameter.ParameterAttributeLi.length; j++) {
					if (queryParameter.ParameterAttributeLi[j].Name == "displayPattern") {
						dateFormat = queryParameter.ParameterAttributeLi[j].Value;
						break;
					}
				}
				if (dateFormat) {
					dateFormat = columnManager.convertDate2DisplayPattern(dateFormat);
				}
			}
			queryParameterManager.applyQueryParameterObserve(queryParameter);
		}
	});
}

QueryParameterManager.prototype.applyQueryParameterObserve = function(queryParameter) {
	YUI(g_financeModule).use("finance-module", function(Y){
		if (queryParameter.ParameterAttributeLi) {
			for (var j = 0; j < queryParameter.ParameterAttributeLi.length; j++) {
				if (queryParameter.ParameterAttributeLi[j].Name == "observe") {
					Y.one("#" + queryParameter.Name).on("change", function(queryParameter, observeAttr){
						return function(e){
							var queryParameterManager = new QueryParameterManager();
							var targetQueryParameter = queryParameterManager.findQueryParameter(listTemplate, observeAttr.Value);
							if (document.getElementById(queryParameter.Name).value) {
								var treeUrlAttr = queryParameterManager.findQueryParameterAttr(targetQueryParameter, "treeUrl");
								// ajax requeset,
								var uri = "/tree/" + treeUrlAttr.Value;
								if (uri.indexOf("?") > -1) {
									uri += "&parentId=" + document.getElementById(queryParameter.Name).value;
								} else {
									uri += "?parentId=" + document.getElementById(queryParameter.Name).value;
								}
								function complete(id, o, args) {
									var id = id; // Transaction ID.
									var data = Y.JSON.parse(o.responseText);
									var htmlLi = ['<option value="">请选择</option>'];
									for (var k = 0; k < data.length; k++) {
										htmlLi.push(Y.Lang.sub('<option value="{code}">{name}</option>', data[k]));
									}
									Y.one("#" + targetQueryParameter.Name).setHTML(htmlLi.join(""));
								};
								Y.on('io:complete', complete, Y, []);
								var request = Y.io(uri);
							} else {
								Y.one("#" + targetQueryParameter.Name).setHTML('<option value="">请选择</option>');
							}
						}
					}(queryParameter, queryParameter.ParameterAttributeLi[j]));
					break;
				}
			}
		}
	});
}

QueryParameterManager.prototype.findQueryParameter = function(listTemplate, name) {
	for (var i = 0; i < listTemplate.QueryParameterGroup.QueryParameterLi.length; i++) {
		if (listTemplate.QueryParameterGroup.QueryParameterLi[i].Name == name) {
			return listTemplate.QueryParameterGroup.QueryParameterLi[i];
		}
	}
	return null;
}

QueryParameterManager.prototype.findQueryParameterAttr = function(queryParameter, name) {
	for (var i = 0; i < queryParameter.ParameterAttributeLi.length; i++) {
		if (queryParameter.ParameterAttributeLi[i].Name == name) {
			return queryParameter.ParameterAttributeLi[i];
		}
	}
	return null;
}

QueryParameterManager.prototype.getQueryField = function(Y, name) {
	var listTemplateIterator = new ListTemplateIterator();
	var result = "";
	var field = null;
	listTemplateIterator.iterateAnyTemplateQueryParameter(result, function(queryParameter, result){
		if (queryParameter.Name == name) {
			if (queryParameter.Editor == "hiddenfield") {
				field = new Y.LHiddenField({
					name : name,
					validateInline: true
				});
			} else if (queryParameter.Editor == "textfield") {
				field = new Y.LTextField({
					name : name,
					validateInline: true
				});
			} else if (queryParameter.Editor == "textareafield") {
				field = new Y.LTextareaField({
					name : name,
					validateInline: true
				});
			} else if (queryParameter.Editor == "numberfield") {
				field = new Y.LNumberField({
					name : name,
					validateInline: true
				});
			} else if (queryParameter.Editor == "datefield") {
				field = new Y.LDateField({
					name : name,
					validateInline: true
				});
			} else if (queryParameter.Editor == "combofield") {
				field = new Y.LSelectField({
					name : name,
					validateInline: true
				});
			} else if (queryParameter.Editor == "displayfield") {
				field = new Y.LDisplayField({
					name : name,
					validateInline: true
				});
			} else if (queryParameter.Editor == "checkboxfield") {
				field = new Y.LChoiceField({
					name : name,
					validateInline: true,
					multi: true
				});
			} else if (queryParameter.Editor == "radiofield") {
				field = new Y.LChoiceField({
					name : name,
					validateInline: true
				});
			} else if (queryParameter.Editor == "triggerfield") {
				field = new Y.LTriggerField({
					name : name,
					validateInline: true
				});
			}
			return true;
		}
		return false;
	});
	return field;
}

QueryParameterManager.prototype.getQueryFormData = function() {
	var result = {};
	for (var key in g_masterFormFieldDict) {
		result[key] = g_masterFormFieldDict[key].get("value");
	}
	return result;
}

