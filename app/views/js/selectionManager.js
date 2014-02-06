function SelectionManager() {
}

/**
 * 选中后,添加到selectionBo里面去
 */
SelectionManager.prototype.addSelectionBo = function(obj) {
	if (!selectionBo) {
		selectionBo = {};
	}
	selectionBo[obj["id"]] = obj;
}

SelectionManager.prototype.getSelectionBo = function(id) {
	if (!id || id === "" || id === 0 || id === "0") {
		return null;
	}
	if (selectionBo[id]) {
		return selectionBo[id];
	}
	return null;
}
