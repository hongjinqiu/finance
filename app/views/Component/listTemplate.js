YUI.add('listtemplate-datatable-paginator', function (Y) {
	function ListTemplatePaginator() {}
	Y.mix(ListTemplatePaginator.prototype, [Y.DataTable.Paginator]);
	ListTemplatePaginator.prototype.processPageRequest = function(page_no, pag_state) {
        var rdata = this._mlistArray,
        pagv  = this.get('paginator'),
        pagm  = pagv.get('model'),
        rpp   = pagm.get('itemsPerPage'),
        sortby= this.get('sortBy') || {},
        istart, iend, url_obj, prop_istart, prop_ipp, prop_iend, prop_page, rqst_str;
    //
    //  Get paginator indices
    //
    if ( pag_state ) {
        istart = pag_state.itemIndexStart;
        iend   = pag_state.itemIndexEnd || istart + rpp;
    } else {
        // usually here on first pass thru, when paginator initiates ...
        istart = ( page_no - 1 ) * rpp;
        iend = istart + rpp - 1;
        iend = ( rdata && iend > rdata.length ) ? rdata.length : iend;
    }

    //
    //  Store the translated replacement object for the request converted
    //  from `serverPaginationMap` (or defaults if none) to a "normalized" format
    //

    url_obj     = {},
    prop_istart = this._srvPagMapObj('itemIndexStart'),
    prop_ipp    = this._srvPagMapObj('itemsPerPage');
    prop_page   = this._srvPagMapObj('page');
    prop_iend   = this._srvPagMapObj('itemIndexEnd');

    url_obj[prop_page]   = page_no;      // page
    url_obj[prop_istart] = istart;      // itemIndexStart
    url_obj[prop_iend]   = iend;        // itemIndexEnd
    url_obj[prop_ipp]    = rpp;         // itemsPerPage
    url_obj.sortBy       = Y.JSON.stringify( sortby );

    // mix-in the model ATTRS with the url_obj
    url_obj = Y.merge(this.get('paginationState'), url_obj);

    //
    //  This is the main guts of retrieving the records,
    //    we already figured out if this was 'local' or 'server' based.
    //
    //   Now, process this page request thru either local data array slicing or
    //    simply firing off a remote server request ...
    //
    console.log(this._pagDataSrc);
    switch(this._pagDataSrc) {

        case 'ds':

            // fire off a request to DataSource, mixing in as the request string
            //  with ATTR `requestStringTemplate` with the "url_obj" map

            rqst_str = this.get('requestStringTemplate') || '';
            console.log(rqst_str);
            console.log(url_obj);
            this.paginatorDSRequest( Y.Lang.sub(rqst_str,url_obj) + "&page=1" );

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
    this.fire('pageUpdate',{ state:pag_state, view:pagv, urlObj: url_obj });
}
	Y.DataTable.Paginator = ListTemplatePaginator;
	Y.Base.mix(Y.DataTable, [Y.DataTable.Paginator]);
}, 'gallery-2013.01.16-21-05', {"requires": ["datatable-base", "base-build", "datatype", "json", "gallery-datatable-paginator"]});

function getColumns(listTemplate) {
	var columns = [];
	columns.push({
		key:        'select',
		allowHTML:  true, // to avoid HTML escaping
		label:      '<input type="checkbox" class="protocol-select-all" title="全部选中"/>',
		formatter:      '<input type="checkbox" />'
		//,emptyCellValue: '<input type="checkbox"/>'
	});
	if (listTemplate.ColumnModel.IdColumn.Hideable != "true") {
		columns.push({
			key: listTemplate.ColumnModel.IdColumn.Name,
			label: listTemplate.ColumnModel.IdColumn.Text
		});
	}
	for (var i = 0; i < listTemplate.ColumnModel.ColumnLi.length; i++) {
		if (listTemplate.ColumnModel.ColumnLi[i].Hideable != "true") {
			columns.push({
				key: listTemplate.ColumnModel.ColumnLi[i].Name,
				label: listTemplate.ColumnModel.ColumnLi[i].Text
			});
		}
	}
	return columns;
}
YUI().use("node", "event", "json", "datatable", "datasource-get", "datasource-jsonschema", "datatable-datasource", "datatable-sort", "datatable-scroll", "cssbutton", 'cssfonts', 'dataschema-json','datasource-io','model-sync-rest',  "gallery-datatable-paginator", 'gallery-paginator-view', "listtemplate-datatable-paginator", function(Y) {
	Y.on("domready", function(e) {
		console.log(Y.DataTable.Paginator.processPageRequest);
		var dataBo = Y.JSON.parse(dataBoJson);
		var listTemplate = Y.JSON.parse(listTemplateJson);
		var columns = getColumns(listTemplate);
		var data = dataBo.items;

		var url = "/component/listtemplate?format=json";
		//var dataSource = new Y.DataSource.Get({ source: url });
		var dataSource = new Y.DataSource.IO({ 
			source: url,
			ioConfig: {
		        method: "POST"
			}
		});
		//**{page}**, **{totalItems}**, **{itemsPerPage}**, **{lastPage}**, **{totalPages}**, **{itemIndexStart}**, **{itemIndexEnd}**
		dataSource.plug(Y.Plugin.DataSourceJSONSchema, {
		  schema: {
		      resultListLocator: "items"
	    	  ,metaFields: {
	    		  	page:   'pageNo',
	    		  	itemsPerPage:     'pageSize',
	    		  	totalItems: 'totalResults'
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
		
		var dt = new Y.DataTable({
			columns: columns
			//,data: data
//			,datasource: dataSource
			,paginationSource: "server"
			,requestStringTemplate: "pageNo={page}"
			,paginator: new Y.PaginatorView({
                model:              new Y.PaginatorModel({itemsPerPage:10}),
                container:          '#pagContC',
                paginatorTemplate:  '#tmpl-bar',
                pageOptions:        [ 10, 20, 50 ]
            }),
            serverPaginationMap: {
            	//totalItems:     'totalItems',
                itemsPerPage:   { toServer:'pageSize', fromServer:'pageSize' },
                page: 'pageNo'
            },

            paginatorResize: true   // this is now a DT attribute (no longer a PaginatorView attribute)
		});
		dt.plug(Y.Plugin.DataTableDataSource, { datasource: dataSource });
		dt.render("#columnModel");
		dt.datasource.load({ request: "pageNo=1" });
		dt.detach('*:change');
		
		dt.delegate("click", function(e){
			var checked = e.target.get('checked') || undefined;
			Y.all(".yui3-datatable-data .yui3-datatable-col-select input").set("checked", checked ? "checked" : "");
		}, ".protocol-select-all", dt);
		
		dt.delegate("click", function(e){
			var checkLi = Y.all(".yui3-datatable-data .yui3-datatable-col-select input").get("checked");
			console.log(checkLi);
			var isAllSelect = true;
			var i = 0;
			for(; i < checkLi.length; i++) {
				if (!checkLi[i]) {
					isAllSelect = false;
					break;
				}
			}
			Y.one(".protocol-select-all").set("checked", isAllSelect ? "checked" : "");
		}, ".yui3-datatable-data .yui3-datatable-col-select input", dt);
		
		Y.one("#testBtn").on("click", function(e){
			//dt.datasource.load({ request: "&1=1" });
			var pagModel = dt.get('paginator').get('model');
			//pagModel.set("page", 2);
			//dt.get("paginator").sayHello();
			dt.sayHello();
//			dt.processPageRequest(2);
//			pagModel.fire("pageChange", pagModel);
		});
	});
});
