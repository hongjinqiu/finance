function QueryParameterManager() {}

QueryParameterManager.prototype.applyQueryParameter = function() {
	YUI().use("node", "event", function(Y){
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
					if (queryParameter.ParameterAttributeLi[j].Name == "inFormat") {
						dateFormat = queryParameter.ParameterAttributeLi[j].Value;
						break;
					}
				}
				if (dateFormat) {
					dateFormat = columnManager.convertDate2DisplayPattern(dateFormat);
					/*
				var calendar = new Y.Calendar({
					trigger: "#" + queryParameter.Name,
					//dates: ['09/14/2009', '09/15/2009'],
					//dateFormat: '%d/%m/%y %A',
					dateFormat: dateFormat,
					setValue: true,
					selectMultipleDates: false
				}).render();
					 */
				}
			}
			queryParameterManager.applyQueryParameterObserve(queryParameter);
		}
	});
}

QueryParameterManager.prototype.applyQueryParameterObserve = function(queryParameter) {
	YUI().use("node", "event", "json", "io-base", function(Y){
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





