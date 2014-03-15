Y.LDateField = Y.Base.create('l-date-field', Y.RDateField, [Y.WidgetChild], {
	initializer : function () {
		Y.LDateField.superclass.initializer.apply(this, arguments);
		var self = this;
		
		var dbPattern = "";
		var displayPattern = "";
		var listTemplateIterator = new ListTemplateIterator();
		var result = "";
		listTemplateIterator.iterateAnyTemplateQueryParameter(result, function(queryParameter, result){
			if (queryParameter.Name == self.get("name")) {
				for (var i = 0; i < queryParameter.ParameterAttributeLi.length; i++) {
					if (queryParameter.ParameterAttributeLi[i].Name == "dbPattern") {
						dbPattern = queryParameter.ParameterAttributeLi[i].Value;
					} else if (queryParameter.ParameterAttributeLi[i].Name == "displayPattern") {
						displayPattern = queryParameter.ParameterAttributeLi[i].Value;
					}
				}
				return true;
			}
			return false;
		});
		this.set("dbPattern", dbPattern);
		this.set("displayPattern", displayPattern);
    }
},
{

    ATTRS: {
    	
    }
});
