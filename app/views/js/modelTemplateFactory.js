function ModelTemplateFactory() {
}

/**
 * 建立父子双向关联
 */
ModelTemplateFactory.prototype._applyReverseRelation = function(dataSource) {
	dataSource.MasterData.Parent = dataSource;
	for (var i = 0; i < dataSource.DetailDataLi.length; i++) {
		dataSource.DetailDataLi[i].Parent = dataSource;
	}
	dataSource.MasterData.FixField.Parent = dataSource.MasterData;
	dataSource.MasterData.BizField.Parent = dataSource.MasterData;
	var modelIterator = new ModelIterator();
	var masterFixFieldLi = modelIterator.getFixFieldLi(dataSource.MasterData.FixField);
	for (var i = 0; i < masterFixFieldLi.length; i++) {
		masterFixFieldLi[i].Parent = dataSource.MasterData.FixField;
	}
	for (var i = 0; i < dataSource.MasterData.BizField.FieldLi.length; i++) {
		dataSource.MasterData.BizField.FieldLi[i].Parent = dataSource.MasterData.BizField;
	}
	for (var i = 0; i < dataSource.DetailDataLi.length; i++) {
		dataSource.DetailDataLi[i].FixField.Parent = dataSource.DetailDataLi[i];
		dataSource.DetailDataLi[i].BizField.Parent = dataSource.DetailDataLi[i];
		
		var detailFixFieldLi = modelIterator.getFixFieldLi(dataSource.DetailDataLi[i].FixField);
		for (var j = 0; j < detailFixFieldLi.length; j++) {
			detailFixFieldLi[j].Parent = dataSource.DetailDataLi[i].FixField;
		}
		
		for (var j = 0; j < dataSource.DetailDataLi[i].BizField.FieldLi.length; j++) {
			dataSource.DetailDataLi[i].BizField.FieldLi[j].Parent = dataSource.DetailDataLi[i].BizField;
		}
	}
}

/**
 * 为字段加入是否主数据集字段的方法
 */
ModelTemplateFactory.prototype._applyIsMasterField = function(dataSource) {
	var modelIterator = new ModelIterator();
	var masterFixFieldLi = modelIterator.getFixFieldLi(dataSource.MasterData.FixField);
	for (var i = 0; i < masterFixFieldLi.length; i++) {
		masterFixFieldLi[i].isMasterField = function() {
			return true;
		}
	}
	for (var i = 0; i < dataSource.MasterData.BizField.FieldLi.length; i++) {
		dataSource.MasterData.BizField.FieldLi[i].isMasterField = function() {
			return true;
		}
	}
	for (var i = 0; i < dataSource.DetailDataLi.length; i++) {
		var detailFixFieldLi = modelIterator.getFixFieldLi(dataSource.DetailDataLi[i].FixField);
		for (var j = 0; j < detailFixFieldLi.length; j++) {
			detailFixFieldLi[j].isMasterField = function() {
				return false;
			}
		}
		
		for (var j = 0; j < dataSource.DetailDataLi[i].BizField.FieldLi.length; j++) {
			dataSource.DetailDataLi[i].BizField.FieldLi[j].isMasterField = function() {
				return false;
			}
		}
	}
}

/**
 * 为字段加入是否关联字段的方法
 */
ModelTemplateFactory.prototype._applyIsRelationField = function(dataSource) {
	var modelIterator = new ModelIterator();
	var result = {};
	modelIterator.iterateAllField(dataSource, result, function(fieldGroup, result){
		fieldGroup.isRelationField = function(){
			if (fieldGroup.RelationDS && fieldGroup.RelationDS.RelationItemLi && fieldGroup.RelationDS.RelationItemLi.length > 0) {
				return true;
			}
			return false;
		}
	});
}

/**
 * 默认用第一个关联字段生成关联配置
 */
ModelTemplateFactory.prototype._applyRelationFieldValue = function(dataSource) {
	var modelIterator = new ModelIterator();
	var result = {};
	var commonUtil = new CommonUtil();
	modelIterator.iterateAllField(dataSource, result, function(fieldGroup, result){
		if (fieldGroup.isRelationField()) {
			if (!fieldGroup.jsConfig) {
				fieldGroup.jsConfig = {};
			}
			var relationItem = fieldGroup.RelationDS.RelationItemLi[0];
			var triggerConfig = {
				displayField: commonUtil.getFuncOrString(relationItem.DisplayField),
				valueField: commonUtil.getFuncOrString(relationItem.ValueField),
				selectorName: commonUtil.getFuncOrString(relationItem.Id),
				selectionMode: "single"
			};
			for (var key in triggerConfig) {
				fieldGroup.jsConfig[key] = triggerConfig[key];
			}
		}
	});
}

/**
 * 添加获取主数据集方法
 */
ModelTemplateFactory.prototype._applyGetMasterData = function(dataSource) {
	var modelIterator = new ModelIterator();
	var result = {};
	modelIterator.iterateAllField(dataSource, result, function(fieldGroup, result){
		fieldGroup.getMasterData = function() {
			if (this.isMasterField()) {
				return this.Parent.Parent;
			}
			return null;
		}
	});
}

/**
 * 添加获取分录数据集方法
 */
ModelTemplateFactory.prototype._applyGetDetailData = function(dataSource) {
	var modelIterator = new ModelIterator();
	var result = {};
	modelIterator.iterateAllField(dataSource, result, function(fieldGroup, result){
		fieldGroup.getDetailData = function() {
			if (this.isMasterField()) {
				return null;
			}
			return this.Parent.Parent;
		}
	});
}

/**
 * 添加获取数据源方法
 */
ModelTemplateFactory.prototype._applyGetDataSource = function(dataSource) {
	var modelIterator = new ModelIterator();
	var result = {};
	modelIterator.iterateAllField(dataSource, result, function(fieldGroup, result){
		fieldGroup.getDataSource = function() {
			if (this.isMasterField()) {
				return this.getMasterData().Parent;
			}
			return this.getDetailData().Parent;
		}
	});
}

/**
 * 添加获取数据集Id方法
 */
ModelTemplateFactory.prototype._applyGetDataSetId = function(dataSource) {
	var modelIterator = new ModelIterator();
	var result = {};
	modelIterator.iterateAllField(dataSource, result, function(fieldGroup, result){
		fieldGroup.getDataSetId = function() {
			if (this.isMasterField()) {
				return this.getMasterData().Id;
			}
			return this.getDetailData().Id;
		}
	});
}

/**
 * 扩展dataSource,当前只扩展FieldGroup的内容,
 * 客户端多了属性:
 * jsConfig: {
 * 		defaultValueExprForJs:function(){}
 * 		calcValueExprForJs:function(){}
 * 		triggerEditor:function(){}
 * }
 */
ModelTemplateFactory.prototype.extendDataSource = function(dataSource, modelExtraInfo) {
	var modelIterator = new ModelIterator();
	var result = {};
	modelIterator.iterateAllField(dataSource, result, function(fieldGroup, result){
		var dataSetConfig = modelExtraInfo[fieldGroup.getDataSetId()];
		if (dataSetConfig && dataSetConfig[fieldGroup.Id]) {
			for (var key in dataSetConfig[fieldGroup.Id]) {
				if (!fieldGroup.jsConfig) {
					fieldGroup.jsConfig = {};
				}
				fieldGroup.jsConfig[key] = dataSetConfig[fieldGroup.Id][key];
			}
		}
	});
}

/**
 * 扩展dataSource,扩展字段的关联关系等等,
 */
ModelTemplateFactory.prototype.enhanceDataSource = function(dataSource) {
	this._applyReverseRelation(dataSource);
	this._applyIsMasterField(dataSource);
	this._applyIsRelationField(dataSource);
	this._applyRelationFieldValue(dataSource);
	this._applyGetMasterData(dataSource);
	this._applyGetDetailData(dataSource);
	this._applyGetDataSource(dataSource);
	this._applyGetDataSetId(dataSource);
}


