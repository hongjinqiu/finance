function FormManager() {
}

FormManager.prototype.getBo = function() {
	return {};
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
