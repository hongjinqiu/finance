function RelationManager() {
}

/**
 * 开窗选回后,加入到relationBo中
 */
RelationManager.prototype.addRelationBo = function(selectorId, url, obj) {
	if (!relationBo[selectorId]) {
		relationBo[selectorId] = {};
	}
	relationBo[selectorId][obj["id"]] = obj;
	if (url) {
		relationBo[selectorId]["url"] = url;
	}
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
			result = o["result"];
			if (result) {
				var url = o["url"];
				self.addRelationBo(selectorId, url, result);
			}
		}
	});
	return result;
}
