DataTableManager.prototype.createAddRowGrid = function() {
	var self = this;
	YUI(formQuickEditJsConfig).use("node", "event", 'array-extras', 'querystring-stringify', "json", "datatable", "datasource-get", "datasource-jsonschema",
			"datatable-datasource", "datatable-sort", "datatable-scroll", "cssbutton", 'cssfonts', 'dataschema-json', 'datasource-io', 'model-sync-rest',
			"gallery-datatable-paginator", 'gallery-paginator-view', "listtemplate-paginator", "datatype-date-format", "io-base", "anim", "panel", "dd-plugin",
			"papersns-form-quickedit", function(Y) {
				var dataTableManager = new DataTableManager();
				var doPopupConfirm = function() {
					var li = dataTableManager.dt.pqe.getRecords();
					{
						for (var i = 0; i < li.length; i++) {
							delete li[i].id
						}
					}
					self.dt.addRows(li);
				};

				var dialog = new Y.Panel({
					contentBox : Y.Node.create('<div id="detail-grid-addrow-dialog" />'),
					headerContent : "新增" + self.param.columnModel.Text,
					bodyContent : '<div style="overflow: auto" id="' + self.param.columnModelName + "_addrow" + '"></div>',
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
							label : 'Cancel',
							action : 'onCancel'
						}, {
							name : 'proceed',
							label : 'OK',
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
					return this.destroy();
				}

				dialog.dd.addHandle('.yui3-widget-hd');
				//		Y.one('#dialog .message').setHTML('mnopq Are you sure you want to [take some action]?');
				//		Y.one('#dialog .message').set('className', 'message icon-bubble');
				dialog.callback = doPopupConfirm;
				var param = {
					data : [ {} ],
					columnModel : self.param.columnModel,
					columnModelName : self.param.columnModelName + "_addrow",// 用于virtualColumn的btn,onclick,回找grid的,暂时没用,
					render : "#" + self.param.columnModelName + "_addrow",// 用panel里面的东东,
					url : "",
					totalResults : dataBo.totalResults || 50,
					pageSize : 10000,
					paginatorContainer : null,
					paginatorTemplate : null,
					columnManager : new ColumnDataSourceManager(self),
					plugin : Y.Plugin.DataTablePFormQuickEdit
				};

				dataTableManager.createDataGrid(Y, param);

				dialog.show();
			});
}

/**
 * 点击新增,新增一行
 */
function g_addRow(dataSetId) {
	gridPanelDict[dataSetId].createAddRowGrid();
}


