function doRefretorComponent() {
	doRefretor("Component");
}

function doRefretorSelector() {
	doRefretor("Selector");
}

function doRefretorForm() {
	doRefretor("Form");
}

function doRefretorDataSource() {
	doRefretor("DataSource");
}

function doRefretor(name) {
	var dtManager = gridPanelDict[name];
	var uri = "/console/refretor?type=" + name;
	Y.on('io:complete', function(id, o, args) {
		var id = id; // Transaction ID.
		var data = Y.JSON.parse(o.responseText);
		dtManager.syncData(data.items);
	}, Y, []);
	var request = Y.io(uri);
}
