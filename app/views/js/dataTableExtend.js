YUI.add(
				'listtemplate-paginator',
				function(Y) {
					function ListTemplatePaginator() {
					}
					Y.mix(ListTemplatePaginator.prototype, [ Y.DataTable.Paginator ]);
					ListTemplatePaginator.prototype.processPageRequest = function(page_no, pag_state) {
						var rdata = this._mlistArray, pagv = this.get('paginator'), pagm = pagv.get('model'), rpp = pagm.get('itemsPerPage'), sortby = this.get('sortBy') || {}, istart, iend, url_obj, prop_istart, prop_ipp, prop_iend, prop_page, rqst_str;
						//
						//  Get paginator indices
						//
						if (pag_state) {
							istart = pag_state.itemIndexStart;
							iend = pag_state.itemIndexEnd || istart + rpp;
						} else {
							// usually here on first pass thru, when paginator initiates ...
							istart = (page_no - 1) * rpp;
							iend = istart + rpp - 1;
							iend = (rdata && iend > rdata.length) ? rdata.length : iend;
						}

						//
						//  Store the translated replacement object for the request converted
						//  from `serverPaginationMap` (or defaults if none) to a "normalized" format
						//

						url_obj = {}, prop_istart = this._srvPagMapObj('itemIndexStart'), prop_ipp = this._srvPagMapObj('itemsPerPage');
						prop_page = this._srvPagMapObj('page');
						prop_iend = this._srvPagMapObj('itemIndexEnd');

						url_obj[prop_page] = page_no; // page
						url_obj[prop_istart] = istart; // itemIndexStart
						url_obj[prop_iend] = iend; // itemIndexEnd
						url_obj[prop_ipp] = rpp; // itemsPerPage
						url_obj.sortBy = Y.JSON.stringify(sortby);

						// mix-in the model ATTRS with the url_obj
						url_obj = Y.merge(this.get('paginationState'), url_obj);

						//
						//  This is the main guts of retrieving the records,
						//    we already figured out if this was 'local' or 'server' based.
						//
						//   Now, process this page request thru either local data array slicing or
						//    simply firing off a remote server request ...
						//
						switch (this._pagDataSrc) {

						case 'ds':

							// fire off a request to DataSource, mixing in as the request string
							//  with ATTR `requestStringTemplate` with the "url_obj" map

							rqst_str = this.get('requestStringTemplate') || '';
							var queryString = getQueryString(Y);
							var pageQueryString = Y.Lang.sub(rqst_str, url_obj);
							if (queryString) {
								queryString += "&" + pageQueryString
							} else {
								queryString = pageQueryString
							}
							this.paginatorDSRequest(queryString);

							break;

						case 'mlist':

							// fire off a ModelSync.REST load "read" request, note that it mixes
							//   the ModelList ATTRS with 'url_obj' in creating the request

							this.paginatorMLRequest(url_obj);

							break;

						case 'local':

							//this.paginatorLocalRequest(page_no,istart,iend);
							this.paginatorLocalRequest(url_obj);

						}

						this.resizePaginator();
						this.fire('pageUpdate', {
							state : pag_state,
							view : pagv,
							urlObj : url_obj
						});
					}
					Y.DataTable.Paginator = ListTemplatePaginator;
					Y.Base.mix(Y.DataTable, [ Y.DataTable.Paginator ]);
				}, 'gallery-2013.01.16-21-05', {
					"requires" : [ "datatable-base", "base-build", "datatype", "json", "gallery-datatable-paginator" ]
				});

function DataTableManager() {
	this.param = null;
	this.dt = null;
}

DataTableManager.prototype.showLoadingImg = function() {
	var gridRenderId = this.param.render.replace("#", "");
	YUI().use("node", function(Y) {
		var node = Y.one(gridRender + " tbody.yui3-datatable-data");
		var xy = node.getXY();
		x = xy[0];
		y = xy[1];
		var width = parseInt(node.getComputedStyle("width"));
		var height = parseInt(node.getComputedStyle("height"));

		var loadingNode = Y.one("#" + gridRenderId + "_loading");
		if (!loadingNode) {
			var loadingStyleLi = [];
			loadingStyleLi.push('position: absolute;');
			loadingStyleLi.push('z-index: 999;');
			loadingStyleLi.push('background-color: white;');
			loadingStyleLi.push('opacity: 0.5;');
			loadingStyleLi.push('filter:alpha(opacity=50);');
			loadingStyleLi.push('width: ' + width + 'px;');
			loadingStyleLi.push('height: ' + height + 'px;');
			loadingStyleLi.push('left: ' + x + 'px;');
			loadingStyleLi.push('top: ' + y + 'px;');
			loadingStyleLi.push('display: none;');

			var loadingImgLi = [];
			loadingImgLi.push('width: 100%;');
			loadingImgLi.push('height: 100%;');
			loadingImgLi.push('text-align: center;');
			var marginTop = parseInt((height - 16) / 2);
			if (marginTop < 0) {
				maringTop = 0;
			}
			loadingImgLi.push('margin-top: ' + marginTop + 'px;');

			var htmlLi = [];
			htmlLi.push('<div id="' + gridRenderId + '_loading" style="' + loadingStyleLi.join("") + '">');
			htmlLi.push('<div id="' + gridRenderId + '_loadingImg" style="' + loadingImgLi.join("") + '">');
			htmlLi.push('<img src="/public/galleryimages/loading_indicator.gif" title="加载中..." border="0" width="16" height="16"/>');
			htmlLi.push('</div>');
			htmlLi.push('</div>');
			Y.one("body").append(htmlLi.join(""));
			loadingNode = Y.one("#loading");
		}
		loadingNode.setStyle("width", width + "px");
		loadingNode.setStyle("height", height + "px");
		loadingNode.setStyle("left", x + "px");
		loadingNode.setStyle("top", y + "px");

		Y.one("#loadingImg").setStyle("marginTop", parseInt((height - 16) / 2) + 'px');

		loadingNode.setStyle("display", "");
	});
}

DataTableManager.prototype.hideLoadingImg = function() {
	var gridRenderId = this.param.render.replace("#", "");
	YUI().use("node", function(Y) {
		var loadingNode = Y.one("#" + gridRenderId + "_loading");
		loadingNode.setStyle("display", "none");
	});
}

DataTableManager.prototype.createDataGrid = function(Y, param, config) {
	var self = this;
	this.param = param;
	
	var data = param.data;
	var columnModel = param.columnModel;
	var render = param.render;
	var url = param.url;
	var totalResults = param.totalResults;
	var pageSize = param.pageSize;
	var paginatorContainer = param.paginatorContainer;
	var paginatorTemplate = param.paginatorTemplate;
	
	var columnManager = new ColumnManager();
	var columns = columnManager.getColumns(render, columnModel, Y);

	//var dataSource = new Y.DataSource.Get({ source: url });
	var dataSource = new Y.DataSource.IO({
		source : url,
		ioConfig : {
			method : "POST"
		},
		on : {
			request : function(e) {
				self.showLoadingImg();
			},
			response : function(e) {
				self.hideLoadingImg();
			}
		}
	});
	if (paginatorContainer) {
		//**{page}**, **{totalItems}**, **{itemsPerPage}**, **{lastPage}**, **{totalPages}**, **{itemIndexStart}**, **{itemIndexEnd}**
		dataSource.plug(Y.Plugin.DataSourceJSONSchema, {
			schema : {
				resultListLocator : "items",
				metaFields : {
					page : 'pageNo',
					itemsPerPage : 'pageSize',
					totalItems : 'totalResults'
				}
		/* ,resultFields: [
		    "Title",
		    "Phone",
		    {
		        key: "Rating",
		        locator: "Rating.AverageRating",
		        parser: function (val) {
		            // YQL is returning "NaN" for unrated restaurants
		            return isNaN(val) ? -1 : +val;
		        }
		    }
		] */
			}
		});
	}
	var gridConfig = {
		columns : columns,
		data : data
		//		,datasource: dataSource
	};
	if (paginatorContainer) {
		var paginatorConfig = {
			paginationSource : "server",
			requestStringTemplate : "pageNo={page}&pageSize={itemsPerPage}",
			paginator : new Y.PaginatorView({
				model : new Y.PaginatorModel({
					itemsPerPage : pageSize
				}),
				container : paginatorContainer,
				paginatorTemplate : paginatorTemplate,
				pageOptions : [ 10, 20, 50 ]
			}),
			serverPaginationMap : {
				//totalItems:     'totalItems',
				itemsPerPage : {
					toServer : 'pageSize',
					fromServer : 'pageSize'
				},
				page : 'pageNo'
			},

			paginatorResize : true
			// this is now a DT attribute (no longer a PaginatorView attribute)
		};
		for(var key in paginatorConfig) {
			gridConfig[key] = paginatorConfig[key];
		}
	}
	var dt = new Y.DataTable(gridConfig);
	dt.plug(Y.Plugin.DataTableDataSource, {
		datasource : dataSource
	});
	if (paginatorContainer) {
		dt.get('paginator').get('model').set('totalItems', totalResults);
	}
	//		dt.resizePaginator();
	dt.render(render);
	//dt.datasource.load({ request: "pageNo=1" });
	//	dt.processPageRequest(1);
	dt.detach('*:change');

	var checkboxCssSelector = self.getCheckboxCssSelector();
	dt.delegate("click", function(e) {
		var checked = e.target.get('checked') || undefined;
		Y.all(checkboxCssSelector).set("checked", checked ? "checked" : "");
	}, ".protocol-select-all", dt);

	dt.delegate("click", function(e) {
		var checkLi = Y.all(checkboxCssSelector).get("checked");
		var isAllSelect = true;
		var i = 0;
		for (; i < checkLi.length; i++) {
			if (!checkLi[i]) {
				isAllSelect = false;
				break;
			}
		}
		// 单选没有全部选中的按钮
		if (Y.one(".protocol-select-all")) {
			Y.one(".protocol-select-all").set("checked", isAllSelect ? "checked" : "");
		}
	}, checkboxCssSelector, dt);
	this.dt = dt;
	return this;
//	return dt;
}

DataTableManager.prototype.getCheckboxCssSelector = function() {
	var renderName = this.param.render;
	var columnModel = this.param.columnModel;
	var result;
	YUI().use("lang", "node", function(Y){
		result = Y.Lang.sub(renderName + " .yui3-datatable-data .yui3-datatable-col-{select} input",{
			"select": columnModel.CheckboxColumn.Name
		});
	});
	return result;
}

DataTableManager.prototype.getSelectRecordLi = function() {
	var self = this;
	var renderName = this.param.render;
	var columnModel = this.param.columnModel;
	var dt = this.dt;
	
	var checkboxCssSelector = self.getCheckboxCssSelector();
	
	var result = [];
	YUI().use("lang", "node", function(Y){
		var li = Y.all(checkboxCssSelector);
		li.each(function(item){
			if (item.get("checked")) {
				result.push(dt.getRecord(item));
			}
		});
	});
	return result;
}

DataTableManager.prototype.doVirtualColumnBtnAction = function(elem, fn){
	var self = this;
	var dt = self.dt;
	YUI().use("lang", "node", function(Y){
		var o = dt.getRecord(Y.one(elem));
		fn(o);
	});
}

/**
 * 外部一般不会调用这个方法,这个方法主要用于做模型控制台的重构用,其它的一般都是ajax table,自动会有loadingImg动画,
 */
DataTableManager.prototype.syncData = function(data){
	var self = this;
	var dt = self.dt;
	self.showLoadingImg()
	dt.set("data", data)
	self.hideLoadingImg()
}
