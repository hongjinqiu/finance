function RelationManager() {
}

/**
 * 开窗选回后,加入到relationBo中
 */
RelationManager.prototype.addRelationBo = function(selectorId, obj) {
	if (!relationBo[selectorId]) {
		relationBo[selectorId] = {};
	}
	relationBo[selectorId][obj["id"]] = obj;
}

RelationManager.prototype.getRelationBo = function(selectorId, id) {
	if (!selectorId || selectorId === "" || selectorId === 0 || selectorId === "0") {
		return null;
	}
	if (!id || id === "" || id === 0 || id === "0") {
		return null;
	}
	if (relationBo[selectorId] && relationBo[selectorId][id]) {
		return relationBo[selectorId][id];
	}
	var self = this;
	var result = null;
	ajaxRequest({
		url : "/console/relation?selectorId=" + selectorId + "&id=" + id + "&date=" + new Date(),
		method: "GET",
		callback : function(o) {
			YUI().use("node", "event", "json", "io-base", function(Y){
				var data = Y.JSON.parse(o.responseText);
				result = data["result"];
				if (result) {
					self.addRelationBo(selectorId, result);
				}
			});
		}
	});
	return result;
}
