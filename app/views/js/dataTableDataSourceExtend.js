DataTableManager.prototype.createAddRowGrid = function(inputDataLi) {
	var self = this;
	YUI(g_financeModule).use("finance-module",
			 function(Y) {
				var pluginDataTableManager = new DataTableManager();
				var doPopupConfirm = function() {
					var li = pluginDataTableManager.dt.pqe.getRecords();
					// 输入中有,输出中没有,删除
					for (var i = 0; i < inputDataLi.length; i++) {
						var isIn = false;
						for (var j = 0; j < li.length; j++) {
							if (inputDataLi[i].id == li[j].id) {
								isIn = true;
								break;
							}
						}
						if (!isIn) {
							self.dt.removeRow(inputDataLi[i].id);
						}
					}
					
					for (var i = 0; i < li.length; i++) {
						var record = self.dt.getRecord(li[i].id);
						if (record) {
							for (var key in li[i]) {
								record.set(key, li[i][key]);
							}
						}
					}
					self.dt.addRows(li);
				};
				var bodyHtmlLi = [];
				bodyHtmlLi.push("<div class='alignLeft'>");
				bodyHtmlLi.push("<input type='button' value='新增' class='' onclick='g_pluginAddRow(\"" + self.param.columnModelName + "\")'/>");
				bodyHtmlLi.push("<input type='button' value='删除' class='' onclick='g_pluginRemoveRow(\"" + self.param.columnModelName + "\")'/>");
				bodyHtmlLi.push("</div>");
				bodyHtmlLi.push('<div style="overflow: auto" id="' + self.param.columnModelName + "_addrow" + '"></div>');

				var dialog = new Y.Panel({
					contentBox : Y.Node.create('<div id="detail-grid-addrow-dialog" />'),
					headerContent : "新增" + self.param.columnModel.Text,
					bodyContent : bodyHtmlLi.join(""),
					width : 1000,
					zIndex : 6,
					centered : true,
					modal : true, // modal behavior
					render : '.example',
					visible : false, // make visible explicitly with .show()
					plugins : [ Y.Plugin.Drag ],
					buttons : {
						footer : [ {
							name : 'cancel',
							label : '取消',
							action : 'onCancel'
						}, {
							name : 'proceed',
							label : '确定',
							action : 'onOK'
						} ]
					}
				});

				dialog.onCancel = function(e) {
					e.preventDefault();
					this.hide();
					// the callback is not executed, and is
					// callback reference removed, so it won't persist
					this.callback = false;
				}

				dialog.onOK = function(e) {
					e.preventDefault();
					this.hide();
					// code that executes the user confirmed action goes here
					if (this.callback) {
						this.callback();
					}
					// callback reference removed, so it won't persist
					this.callback = false;
				}

				dialog.hide = function() {
					g_gridPanelDict[self.param.columnModelName + "_addrow"] = null;
					return this.destroy();
				}

				dialog.dd.addHandle('.yui3-widget-hd');
				//		Y.one('#dialog .message').setHTML('mnopq Are you sure you want to [take some action]?');
				//		Y.one('#dialog .message').set('className', 'message icon-bubble');
				dialog.callback = doPopupConfirm;
				var data = [{}];
				if (inputDataLi) {
					data = inputDataLi;
				}
				var param = {
					data : data,
					columnModel : self.param.columnModel,
					columnModelName : self.param.columnModelName + "_addrow",// 用于virtualColumn的btn,onclick,回找grid的,暂时没用,
					render : "#" + self.param.columnModelName + "_addrow",// 用panel里面的东东,
					url : "",
					totalResults : g_dataBo.totalResults || 50,
					pageSize : 10000,
					paginatorContainer : null,
					paginatorTemplate : null,
					//columnManager : new ColumnDataSourceManager(self),
					plugin : Y.Plugin.DataTablePFormQuickEdit
				};

				pluginDataTableManager.createDataGrid(Y, param);
				g_gridPanelDict[self.param.columnModelName + "_addrow"] = pluginDataTableManager;
				dialog.show();
			});
}

function doPluginVirtualColumnBtnAction(columnModelName, elem, fn){
	var self = g_gridPanelDict[columnModelName + "_addrow"];
	var dt = self.dt;
	var yInst = self.yInst;
	var o = dt.getRecord(yInst.one(elem));
	fn(o, columnModelName);
}

/**
 * 插件表头新增,新增一行
 */
function g_pluginAddRow(dataSetId) {
	var formManager = new FormManager();
	var data = formManager.getDataSetNewData(dataSetId);
	g_gridPanelDict[dataSetId + "_addrow"].dt.addRow(data);
}

/**
 * 插件表头删除,删除多行
 */
function g_pluginRemoveRow(dataSetId) {
	var selectRecordLi = g_gridPanelDict[dataSetId + "_addrow"].getSelectRecordLi();
	if (selectRecordLi.length == 0) {
		showAlert("请先选择");
	} else {
		for (var i = 0; i < selectRecordLi.length; i++) {
			g_gridPanelDict[dataSetId + "_addrow"].dt.removeRow(selectRecordLi[i]);
		}
	}
}

/**
 * 点击删除,删除一行
 */
function g_pluginRemoveSingleRow(o, dataSetId) {
	g_gridPanelDict[dataSetId + "_addrow"].dt.removeRow(o);
}

/**
 * 点击行项复制,复制一行
 */
function g_pluginCopyRow(o, dataSetId) {
//	var inputDataLi = [];
	var formManager = new FormManager();
	var id = o.get("id");
	var li = g_gridPanelDict[dataSetId + "_addrow"].dt.pqe.getRecords();
	var data = {};
	for (var i = 0; i < li.length; i++) {
		if (li[i].id == id) {
			data = li[i];
			break;
		}
	}
	var data = formManager.getDataSetCopyData(dataSetId, data);
//	inputDataLi.push(data);
	g_gridPanelDict[dataSetId + "_addrow"].dt.addRow(data);
}

/**
 * 点击新增,新增一行
 */
function g_addRow(dataSetId) {
	var inputDataLi = [];
	var formManager = new FormManager();
	var data = formManager.getDataSetNewData(dataSetId);
	inputDataLi.push(data);
	g_gridPanelDict[dataSetId].createAddRowGrid(inputDataLi);
}

/**
 * 点击删除,删除多行
 */
function g_removeRow(dataSetId) {
	var selectRecordLi = g_gridPanelDict[dataSetId].getSelectRecordLi();
	if (selectRecordLi.length == 0) {
		showAlert("请先选择");
	} else {
		for (var i = 0; i < selectRecordLi.length; i++) {
			g_gridPanelDict[dataSetId].dt.removeRow(selectRecordLi[i]);
		}
	}
}

/**
 * 点击删除,删除一行
 */
function g_removeSingleRow(o, dataSetId) {
	g_gridPanelDict[dataSetId].dt.removeRow(o);
}

/**
 * 点击行项编辑,编辑一行
 */
function g_editSingleRow(o, dataSetId) {
	var inputDataLi = [];
	inputDataLi.push(o.toJSON());
	g_gridPanelDict[dataSetId].createAddRowGrid(inputDataLi);
}

/**
 * 点击行项复制,复制一行
 */
function g_copyRow(o, dataSetId) {
	var inputDataLi = [];
	var formManager = new FormManager();
	var data = formManager.getDataSetCopyData(dataSetId, o.toJSON());
	inputDataLi.push(data);
	g_gridPanelDict[dataSetId].createAddRowGrid(inputDataLi);
}

/**
 * 点击表格头编辑,编辑多行
 */
function g_editRow(dataSetId) {
	var selectRecordLi = g_gridPanelDict[dataSetId].getSelectRecordLi();
	if (selectRecordLi.length == 0) {
		showAlert("请先选择");
	} else {
		var inputDataLi = [];
		for (var i = 0; i < selectRecordLi.length; i++) {
			inputDataLi.push(selectRecordLi[i].toJSON());
		}
		g_gridPanelDict[dataSetId].createAddRowGrid(inputDataLi);
	}
}



