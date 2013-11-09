var yInst;
var dtInst;

//--------------测试函数区----------------
function doEdit(o) {
	console.log(o);
	console.log(o.toJSON());
}

function test() {
	console.log(getSelectRecordLi());
}
//------------------------------

function getCheckboxCssSelector() {
	return yInst.Lang.sub(".yui3-datatable-data .yui3-datatable-col-{select} input",{
		"select": listTemplate.ColumnModel.CheckboxColumn.Name
	});
}

function getSelectRecordLi() {
	var li = yInst.all(getCheckboxCssSelector());
	var result = [];
	li.each(function(item){
		if (item.get("checked")) {
			result.push(dtInst.getRecord(item));
		}
	});
	return result;
}

function doVirtualColumnBtnAction(elem, fn){
	var o = dtInst.getRecord(yInst.one(elem));
	fn(o);
}

function getQueryString(Y) {
	var form = Y.one('#queryForm'), query;
  
	query = Y.QueryString.stringify(Y.Array.reduce(Y.one(form).all('input[name],select[name],textarea[name]')._nodes, {}, function (init, el, index, array) {
		var isCheckable = (el.type == "checkbox" || el.type == "radio");
		if ((isCheckable && el.checked) || !isCheckable) {
			init[el.name] = el.value;
		}
		return init;
	}));
 
	return query;
}

YUI.add('listtemplate-paginator', function (Y) {
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
    switch(this._pagDataSrc) {

        case 'ds':

            // fire off a request to DataSource, mixing in as the request string
            //  with ATTR `requestStringTemplate` with the "url_obj" map

            rqst_str = this.get('requestStringTemplate') || '';
            var queryString = getQueryString(Y);
            var pageQueryString = Y.Lang.sub(rqst_str,url_obj);
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
    this.fire('pageUpdate',{ state:pag_state, view:pagv, urlObj: url_obj });
}
	Y.DataTable.Paginator = ListTemplatePaginator;
	Y.Base.mix(Y.DataTable, [Y.DataTable.Paginator]);
}, 'gallery-2013.01.16-21-05', {"requires": ["datatable-base", "base-build", "datatype", "json", "gallery-datatable-paginator"]});

function currencyFormatFunc(o) {
	var formatConfig = null;
	var currencyField = o.column.currencyField;
	if (currencyField != "") {
		var prefix = null;
		var decimalPlaces = null;
		if (o.column.isMoney == "true") {// 是否金额
			if (sysParam[currencyField]) {// 本位币
				prefix = sysParam[currencyField]["prefix"];
				decimalPlaces = sysParam[currencyField]["decimalPlaces"];
			}
			if (o.data[currencyField]) {// 本行记录中是否存在对应币别
				prefix = o.data[currencyField]["prefix"];
				decimalPlaces = o.data[currencyField]["decimalPlaces"];
			}
		} else if (o.column.isUnitPrice == "true") {// 单价
			if (sysParam[currencyField]) {// 本位币
				prefix = sysParam[currencyField]["prefix"];
				decimalPlaces = sysParam[currencyField]["unitPriceDecimalPlaces"];
			}
			if (o.data[currencyField]) {// 本行记录中是否存在对应币别
				prefix = o.data[currencyField]["prefix"];
				decimalPlaces = o.data[currencyField]["unitPriceDecimalPlaces"];
			}
		} else if (o.column.isCost == "true") {// 成本
			if (sysParam[currencyField]) {// 本位币
				prefix = sysParam[currencyField]["prefix"];
				decimalPlaces = sysParam["unitCostDecimalPlaces"];
			}
			if (o.data[currencyField]) {// 本行记录中是否存在对应币别
				prefix = o.data[currencyField]["prefix"];
				decimalPlaces = sysParam["unitCostDecimalPlaces"];
			}
		} else {// 是否金额
			if (sysParam[currencyField]) {// 本位币
				prefix = sysParam[currencyField]["prefix"];
				decimalPlaces = sysParam[currencyField]["decimalPlaces"];
			}
			if (o.data[currencyField]) {// 本行记录中是否存在对应币别
				prefix = o.data[currencyField]["prefix"];
				decimalPlaces = o.data[currencyField]["decimalPlaces"];
			}
		}
		
		if (prefix !== null) {
			return yInst.DataType.Number.format(o.value, {
				prefix: prefix,
				decimalPlaces: decimalPlaces,
				decimalSeparator: ".",
				thousandsSeparator: sysParam.thousandsSeparator,
				suffix: ""
			});
		} else {
			console.log(o);
			console.log("在系统参数和本行记录中,没有找到currencyField:" + currencyField);
		}
	} else if (o.column.isPercent == "true") {// 本位币
		return yInst.DataType.Number.format(o.value, {
			prefix: "",
			decimalPlaces: sysParam["percentDecimalPlaces"],
			decimalSeparator: ".",
			thousandsSeparator: sysParam.thousandsSeparator,
			suffix: "%"	
		});
	}
	return yInst.DataType.Number.format(o.value, {
//    	prefix            : o.column.prefix     || '￥',
//		decimalPlaces     : o.column.decimalPlaces      || 2,
//		decimalSeparator  : o.column.decimalSeparator   || '.',
//		thousandsSeparator: o.column.thousandsSeparator || ',',
//		suffix            : o.column.suffix || ''
		prefix            : o.column.prefix,
		decimalPlaces     : o.column.decimalPlaces,
		decimalSeparator  : o.column.decimalSeparator,
		thousandsSeparator: o.column.thousandsSeparator,
		suffix            : o.column.suffix
	});
}

function createIdColumn(listTemplate) {
	if (listTemplate.ColumnModel.IdColumn.Hideable != "true") {
		return {
			key: listTemplate.ColumnModel.IdColumn.Name,
			label: listTemplate.ColumnModel.IdColumn.Text
		};
	}
	return null;
}

function createCheckboxColumn(listTemplate) {
	if (listTemplate.ColumnModel.CheckboxColumn.Hideable != "true") {
		var key = listTemplate.ColumnModel.CheckboxColumn.Name;
		if (listTemplate.ColumnModel.SelectionMode == "radio") {
			return {
				key:        key,
				allowHTML:  true, // to avoid HTML escaping
				label:      '选择',
				//formatter:      '<input type="radio" name="' + key + '" />'
				formatter:function(o) {
					if (o.value) {
						return '<input type="radio" name="' + key + '" />';
					}
					return "";
				}
				//,emptyCellValue: '<input type="checkbox"/>'
			};
		} else {
			return {
				key:        key,
				allowHTML:  true, // to avoid HTML escaping
				label:      '<input type="checkbox" class="protocol-select-all" title="全部选中"/>',
				//formatter:      '<input type="checkbox" />'
				formatter:function(o) {
					if (o.value) {
						return '<input type="checkbox" />';
					}
					return "";
				}
				//,emptyCellValue: '<input type="checkbox"/>'
			};
		}
	}
	return null;
}

function createVirtualColumn(listTemplate, columnIndex) {
	var i = columnIndex;
	if (listTemplate.ColumnModel.ColumnLi[i].XMLName.Local == "virtual-column" && listTemplate.ColumnModel.ColumnLi[i].Hideable != "true") {
		var virtualColumn = listTemplate.ColumnModel.ColumnLi[i];
		return {
			key: listTemplate.ColumnModel.ColumnLi[i].Name,
			label: listTemplate.ColumnModel.ColumnLi[i].Text,
			allowHTML:  true, // to avoid HTML escaping
			formatter:      function(virtualColumn){
				return function(o){
					var htmlLi = [];
					var buttonBoLi = o.value[virtualColumn.Buttons.XMLName.Local];
					for (var j = 0; j < virtualColumn.Buttons.ButtonLi.length; j++) {
						var btnTemplate = null;
						if (virtualColumn.Buttons.ButtonLi[j].Mode == "fn") {
							btnTemplate = "<input type='button' value='{value}' onclick='doVirtualColumnBtnAction(this, {handler})' class='{class}' />";
						} else if (virtualColumn.Buttons.ButtonLi[j].Mode == "url") {
							btnTemplate = "<input type='button' value='{value}' onclick='location.href=\"{href}\"' class='{class}' />";
						} else {
							btnTemplate = "<input type='button' value='{value}' onclick='showModalDialog(\"{href}\")' class='{class}' />";
						}
						if (buttonBoLi[j]["isShow"]) {
							// handler进行值的预替换,
							var handler = virtualColumn.Buttons.ButtonLi[j].Handler;
							handler = yInst.Lang.sub(handler, o.data);
							htmlLi.push(yInst.Lang.sub(btnTemplate, {
								value: virtualColumn.Buttons.ButtonLi[j].Text,
								handler: handler,
								"class": virtualColumn.Buttons.ButtonLi[j].IconCls,
								href: handler
							}));
						}
					}
					return htmlLi.join("");
				}
			}(virtualColumn)
		};
	}
	return null;
}

function createNumberColumn(listTemplate, columnIndex) {
	var i = columnIndex;
	var decimalPlaces = 2;
	if (listTemplate.ColumnModel.ColumnLi[i].DecimalPlaces) {
		decimalPlaces = parseInt(listTemplate.ColumnModel.ColumnLi[i].DecimalPlaces);
	}
	var isFormatter = listTemplate.ColumnModel.ColumnLi[i].Prefix != "";
	isFormatter = isFormatter || listTemplate.ColumnModel.ColumnLi[i].DecimalPlaces != "";
	isFormatter = isFormatter || listTemplate.ColumnModel.ColumnLi[i].DecimalSeparator != "";
	isFormatter = isFormatter || listTemplate.ColumnModel.ColumnLi[i].ThousandsSeparator != "";
	isFormatter = isFormatter || listTemplate.ColumnModel.ColumnLi[i].Suffix != "";
	
	// 财务相关字段的判断,以决定是否用 formatter 函数,
	isFormatter = isFormatter || listTemplate.ColumnModel.ColumnLi[i].CurrencyField != "";
	isFormatter = isFormatter || listTemplate.ColumnModel.ColumnLi[i].IsPercent != "";
	
	if (isFormatter) {
		return {
			key: listTemplate.ColumnModel.ColumnLi[i].Name,
			label: listTemplate.ColumnModel.ColumnLi[i].Text,
			formatter: currencyFormatFunc,
			
			prefix: listTemplate.ColumnModel.ColumnLi[i].Prefix,
			decimalPlaces: decimalPlaces,
			decimalSeparator: listTemplate.ColumnModel.ColumnLi[i].DecimalSeparator,
			thousandsSeparator: listTemplate.ColumnModel.ColumnLi[i].ThousandsSeparator,
			suffix: listTemplate.ColumnModel.ColumnLi[i].Suffix,
			
			currencyField: listTemplate.ColumnModel.ColumnLi[i].CurrencyField,
			isPercent: listTemplate.ColumnModel.ColumnLi[i].IsPercent,
			isMoney: listTemplate.ColumnModel.ColumnLi[i].IsMoney,
			isUnitPrice: listTemplate.ColumnModel.ColumnLi[i].IsUnitPrice,
			isCost: listTemplate.ColumnModel.ColumnLi[i].IsCost
		};
	}
	return {
		key: listTemplate.ColumnModel.ColumnLi[i].Name,
		label: listTemplate.ColumnModel.ColumnLi[i].Text
	};
}

function convertDate2DisplayPattern(displayPattern) {
	displayPattern = displayPattern.replace("yyyy", "%G");
	displayPattern = displayPattern.replace("MM", "%m");
	displayPattern = displayPattern.replace("dd", "%d");
	displayPattern = displayPattern.replace("HH", "%H");
	displayPattern = displayPattern.replace("mm", "%M");
	displayPattern = displayPattern.replace("ss", "%S");
	return displayPattern;
}

/*
	DisplayPattern string `xml:"displayPattern,attr"`
	DbPattern      string `xml:"dbPattern,attr"`
 */
function createDateColumn(listTemplate, columnIndex) {
	var i = columnIndex;
	var dbPattern = listTemplate.ColumnModel.ColumnLi[i].DbPattern;
	var displayPattern = listTemplate.ColumnModel.ColumnLi[i].DisplayPattern;
	if (dbPattern && displayPattern) {
		return {
			key: listTemplate.ColumnModel.ColumnLi[i].Name,
			label: listTemplate.ColumnModel.ColumnLi[i].Text,
			dbPattern: dbPattern,
			displayPattern: displayPattern,
			formatter: function(o) {
				if (o.value !== undefined && o.value !== null) {
					var date = new Date();
					var value = o.value + "";
					if (o.column.dbPattern.indexOf("yyyy") > -1) {
						var start = o.column.dbPattern.indexOf("yyyy");
						var end = o.column.dbPattern.indexOf("yyyy") + "yyyy".length;
						var yyyy = value.substring(start, end);
						date.setYear(parseInt(yyyy));
					}
					if (o.column.dbPattern.indexOf("MM") > -1) {
						var start = o.column.dbPattern.indexOf("MM");
						var end = o.column.dbPattern.indexOf("MM") + "MM".length;
						var mm = value.substring(start, end);
						date.setMonth(parseInt(mm) - 1);
					}
					if (o.column.dbPattern.indexOf("dd") > -1) {
						var start = o.column.dbPattern.indexOf("dd");
						var end = o.column.dbPattern.indexOf("dd") + "dd".length;
						var dd = value.substring(start, end);
						date.setDate(parseInt(dd));
					}
					if (o.column.dbPattern.indexOf("HH") > -1) {
						var start = o.column.dbPattern.indexOf("HH");
						var end = o.column.dbPattern.indexOf("HH") + "HH".length;
						var hh = value.substring(start, end);
						date.setHours(parseInt(hh));
					}
					if (o.column.dbPattern.indexOf("mm") > -1) {
						var start = o.column.dbPattern.indexOf("mm");
						var end = o.column.dbPattern.indexOf("mm") + "mm".length;
						var mm = value.substring(start, end);
						date.setMinutes(mm);
					}
					if (o.column.dbPattern.indexOf("ss") > -1) {
						var start = o.column.dbPattern.indexOf("ss");
						var end = o.column.dbPattern.indexOf("ss") + "ss".length;
						var ss = value.substring(start, end);
						date.setSeconds(ss);
					}
					// js格式参考 http://yuilibrary.com/yui/docs/api/classes/Date.html#method_format
					var displayPattern = convertDate2DisplayPattern(o.column.displayPattern);
					return yInst.DataType.Date.format(date, {
						format: displayPattern
					});
				}
				return o.value;
			}
		};
	} else {
		console.log(listTemplate.ColumnModel.ColumnLi[i]);
		console.log("日期字段未同时配置dbPattern和displayPattern");
	}
	return {
		key: listTemplate.ColumnModel.ColumnLi[i].Name,
		label: listTemplate.ColumnModel.ColumnLi[i].Text
	};
}

function createBooleanColumn(listTemplate, columnIndex) {
	var i = columnIndex;
	return {
		key: listTemplate.ColumnModel.ColumnLi[i].Name,
		label: listTemplate.ColumnModel.ColumnLi[i].Text,
		formatter: function(o) {
			if (o.value + "" == "true") {
				return "是";
			} else if (o.value + "" == "false") {
				return "否";
			}
			return o.value;
		}
	};
}

function createDictionaryColumn(listTemplate, columnIndex) {
	var i = columnIndex;
	return {
		key: listTemplate.ColumnModel.ColumnLi[i].Name,
		label: listTemplate.ColumnModel.ColumnLi[i].Text,
		formatter: function(o) {
			var dictionaryValue = o.data[listTemplate.ColumnModel.ColumnLi[i].Name + "_DICTIONARY_NAME"];
			if (dictionaryValue !== undefined && dictionaryValue !== null) {
				return dictionaryValue;
			}
			console.log(listTemplate.ColumnModel.ColumnLi[i]);
			console.log("字典字段没找到_DICTIONARY_NAME,code:" + o.value);
			return o.value;
		}
	};
}

function createRowIndexColumn(listTemplate) {
	if (listTemplate.ColumnModel.Rownumber == "true") {
		return {
			key: "",
			label: "序号",
			formatter: function(o) {
				return o.rowIndex + 1;
			}
		};
	}
	return null;
}

function createColumn(listTemplate, columnIndex) {
	var i = columnIndex;
	if (listTemplate.ColumnModel.ColumnLi[i].XMLName.Local != "virtual-column" && listTemplate.ColumnModel.ColumnLi[i].Hideable != "true") {
		if (listTemplate.ColumnModel.ColumnLi[i].XMLName.Local == "number-column") {
			return createNumberColumn(listTemplate, columnIndex);
		} else if (listTemplate.ColumnModel.ColumnLi[i].XMLName.Local == "date-column") {
			return createDateColumn(listTemplate, columnIndex);
		} else if (listTemplate.ColumnModel.ColumnLi[i].XMLName.Local == "boolean-column") {
			return createBooleanColumn(listTemplate, columnIndex);
		} else if (listTemplate.ColumnModel.ColumnLi[i].XMLName.Local == "dictionary-column") {
			return createDictionaryColumn(listTemplate, columnIndex);
		}
		return {
			key: listTemplate.ColumnModel.ColumnLi[i].Name,
			label: listTemplate.ColumnModel.ColumnLi[i].Text
		};
	}
	return null;
}

function getColumns(listTemplate, Y) {
	var columns = [];
	var checkboxColumn = createCheckboxColumn(listTemplate);
	if (checkboxColumn) {
		columns.push(checkboxColumn);
	}
	var idColumn = createIdColumn(listTemplate);
	if (idColumn) {
		columns.push(idColumn);
	}
	var rowIndexColumn = createRowIndexColumn(listTemplate);
	if (rowIndexColumn) {
		columns.push(rowIndexColumn);
	}
	
	for (var i = 0; i < listTemplate.ColumnModel.ColumnLi.length; i++) {
		var column = createColumn(listTemplate, i);
		if (column) {
			columns.push(column);
		} else {
			var virtualColumn = createVirtualColumn(listTemplate, i);
			if (virtualColumn) {
				columns.push(virtualColumn);
			}
		}
	}
	return columns;
}

function createDataGrid() {
	var Y = yInst;
	var columns = getColumns(listTemplate, Y);
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
		,data: data
//		,datasource: dataSource
		,paginationSource: "server"
		,requestStringTemplate: "pageNo={page}&pageSize={itemsPerPage}"
		,paginator: new Y.PaginatorView({
			model:              new Y.PaginatorModel({itemsPerPage:DATA_PROVIDER_SIZE}),
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
	dtInst = dt;
	dt.plug(Y.Plugin.DataTableDataSource, { datasource: dataSource });
	dt.get('paginator').get('model').set('totalItems', dataBo.totalResults);;
//		dt.resizePaginator();
	dt.render("#columnModel");
	//dt.datasource.load({ request: "pageNo=1" });
//	dt.processPageRequest(1);
	dt.detach('*:change');
	
	dt.delegate("click", function(e){
		var checked = e.target.get('checked') || undefined;
		Y.all(getCheckboxCssSelector()).set("checked", checked ? "checked" : "");
	}, ".protocol-select-all", dt);
	
	dt.delegate("click", function(e){
		var checkLi = Y.all(getCheckboxCssSelector()).get("checked");
		var isAllSelect = true;
		var i = 0;
		for(; i < checkLi.length; i++) {
			if (!checkLi[i]) {
				isAllSelect = false;
				break;
			}
		}
		// 单选没有全部选中的按钮
		if (Y.one(".protocol-select-all")) {
			Y.one(".protocol-select-all").set("checked", isAllSelect ? "checked" : "");
		}
	}, getCheckboxCssSelector(), dt);
}

function applyQueryParameter() {
	var Y = yInst;
	for (var i = 0; i < listTemplate.QueryParameterGroup.QueryParameterLi.length; i++) {
		var queryParameter = listTemplate.QueryParameterGroup.QueryParameterLi[i];
		if (queryParameter.Editor == "numberfield") {
			Y.one("#" + queryParameter.Name).on("keypress", function(e) {
				if (!(e.keyCode == 9 || e.keyCode == 13 || e.keyCode == 46 || e.keyCode == 116 || e.keyCode == 118 || (e.keyCode >= 48 && e.keyCode <=57))) {// 0-9.,118:ctrl + v, 116:Ctrl + F5,13: enter,9: tab
					e.preventDefault();
				}
			});
		} else if (queryParameter.Editor == "datefield") {
			var dateFormat = null;
			for (var j = 0; j < queryParameter.ParameterAttributeLi.length; j++) {
				if (queryParameter.ParameterAttributeLi[j].Name == "inFormat") {
					dateFormat = queryParameter.ParameterAttributeLi[j].Value;
					break;
				}
			}
			if (dateFormat) {
				dateFormat = convertDate2DisplayPattern(dateFormat);
				var calendar = new Y.Calendar({
					trigger: "#" + queryParameter.Name,
					//dates: ['09/14/2009', '09/15/2009'],
					//dateFormat: '%d/%m/%y %A',
					dateFormat: dateFormat,
					setValue: true,
					selectMultipleDates: false
				}).render();
			}
		}
		// TODO observe 的相关操作,在change的时候,如果被观察方是combo,刷新option
		
	}
}

function applyDateLocale(Y) {
	if (!Y.DataType.Date.Locale) {
		Y.DataType.Date.Locale = {};
		var YDateEn = {
				a: ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"],
				A: ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"],
				b: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
				B: ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"],
				c: "%a %d %b %Y %T %Z",
				p: ["AM", "PM"],
				P: ["am", "pm"],
				r: "%I:%M:%S %p",
				x: "%d/%m/%y",
				X: "%T"
		};
		Y.DataType.Date.Locale["en"] = YDateEn;
		
		Y.DataType.Date.Locale["en-US"] = Y.merge(YDateEn, {
			c: "%a %d %b %Y %I:%M:%S %p %Z",
			x: "%m/%d/%Y",
			X: "%I:%M:%S %p"
		});
		
		Y.DataType.Date.Locale["en-GB"] = Y.merge(YDateEn, {
			r: "%l:%M:%S %P %Z"
		});
		Y.DataType.Date.Locale["en-AU"] = Y.merge(YDateEn);
	}
}

YUI().use("node", "event", 'array-extras', 'querystring-stringify', "json", "datatable", "datasource-get", "datasource-jsonschema", "datatable-datasource", "datatable-sort", "datatable-scroll", "cssbutton", 'cssfonts', 'dataschema-json','datasource-io','model-sync-rest',"gallery-datatable-paginator",'gallery-paginator-view',"listtemplate-paginator","datatype-date-format","gallery-aui-calendar-datepicker-select", function(Y) {
	Y.on("domready", function(e) {
		applyDateLocale(Y);
		yInst = Y;
		createDataGrid();
		applyQueryParameter();
		
		Y.one("#queryBtn").on("click", function(e){
			var pagModel = dtInst.get('paginator').get('model');
			var page = pagModel.get("page");
			if (page == 1) {
				dtInst.refreshPaginator();
			} else {
				pagModel.set("page", 1);
			}
		});
	});
});
