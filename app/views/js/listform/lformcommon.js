function ChoiceFieldManager(){}

ChoiceFieldManager.prototype.getChoices = function(name) {
	var choices = [];
	var listTemplateIterator = new ListTemplateIterator();
	var result = "";
	listTemplateIterator.iterateAnyTemplateQueryParameter(result, function(queryParameter, result){
		if (queryParameter.Name == name) {
			for (var i = 0; i < queryParameter.ParameterAttributeLi.length; i++) {
				if (queryParameter.ParameterAttributeLi[i].Name == "dictionary") {
					var dictionaryCode = queryParameter.ParameterAttributeLi[i].Value;
					var dictValueLi = g_layerBoLi[dictionaryCode];
					for (var j = 0; j < dictValueLi.length; j++) {
						choices.push({
							"label": dictValueLi[j].name,
							"value": dictValueLi[j].code
						});
					}
					break;
				}
			}
			return true;
		}
		return false;
	});
	return choices;
}

function LFormManager(){}

LFormManager.prototype.applyEventBehavior = function(formObj) {
	var self = formObj;
	var listTemplateIterator = new ListTemplateIterator();
	var result = "";
	listTemplateIterator.iterateAnyTemplateQueryParameter(result, function(queryParameter, result){
		if (queryParameter.Name == self.get("name")) {
			if (queryParameter.jsConfig) {
				for (var key in queryParameter.jsConfig.listeners) {
					if (key == "valueChange") {
						self.after("valueChange", function(key) {
							return function(e) {
								queryParameter.jsConfig.listeners[key](e, self);
							}
						}(key));
					} else {
						self._fieldNode.on(key, function(key) {
							return function(e) {
								queryParameter.jsConfig.listeners[key](e, self);
							}
						}(key));
					}
				}
			}
			
			return true;
		}
		return false;
	});
}

