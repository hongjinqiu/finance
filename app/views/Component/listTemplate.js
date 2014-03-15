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

function getSelectRecordLi() {
	return dtInst.getSelectRecordLi();
}

// 用 dataTableExtend 里面的函数
//function doVirtualColumnBtnAction(gridPanelId, elem, fn){
//	var inst = gridPanelDict[gridPanelId];
//	return inst.doVirtualColumnBtnAction(elem, fn);
//}

function getQueryString(Y) {
	var form = Y.one('#queryForm'), query;
  
	query = Y.QueryString.stringify(Y.Array.reduce(Y.one(form).all('input[name],select[name],textarea[name]')._nodes, {}, function (init, el, index, array) {
		var isCheckable = (el.type == "checkbox" || el.type == "radio");
		if ((isCheckable && el.checked) || !isCheckable) {
			if (isCheckable && el.checked) {
				if (!init[el.name]) {
					init[el.name] = el.value;
				} else {
					init[el.name] += "," + el.value;
				}
			} else {
				init[el.name] = el.value;
			}
		}
		return init;
	}));
 
	return query;
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

function getDsUrl(listTemplate) {
	return "/console/listschema?@name=" + listTemplate.Id + "&format=json";
}

function listMain() {
	YUI(g_financeModule).use("finance-module", function(Y){
			//applyDateLocale(Y);
			yInst = Y;
			var dataTableManager = new DataTableManager();
			/*
//	paginatorContainer : '#pagContC',
//	paginatorTemplate : '#tmpl-bar',
			 */
			var renderName = "#columnModel_1";
			var columnModelName = renderName.replace("#", "");
			var param = {
					data:dataBo.items,
					columnModel:listTemplate.ColumnModel,
					columnModelName:columnModelName,
					render:renderName,
					url:getDsUrl(listTemplate),
					totalResults: dataBo.totalResults || 1,
					pageSize: DATA_PROVIDER_SIZE,
					paginatorContainer : '#pagContC',
					paginatorTemplate : '#tmpl-bar'
			};
			dtInst = dataTableManager.createDataGrid(yInst, param);
			gridPanelDict[columnModelName] = dtInst;
			var queryParameterManager = new QueryParameterManager();
			queryParameterManager.applyQueryParameter();
			queryParameterManager.applyQueryDefaultValue();
			applyQueryBtnBehavior();
	});
}

function applyQueryBtnBehavior() {
	YUI(g_financeModule).use("finance-module", function(Y){
		Y.one("#queryBtn").on("click", function(e){
			var pagModel = dtInst.dt.get('paginator').get('model');
			var page = pagModel.get("page");
			if (page == 1) {
				dtInst.dt.refreshPaginator();
			} else {
				pagModel.set("page", 1);
			}
		});
		/*
$("#btn_more").click(function(){
	$("#btn_more").css("display","none");	  
	$("#btn_up").css("display","block");	
	$("#search1").slideDown();							  
  });
	$("#btn_up").click(function(){
	$("#btn_more").css("display","block");	  
	$("#btn_up").css("display","none");	
	$("#search1").slideUp();
		 */
		var duration = 0.4;
		Y.one("#btnMore").on("click", function(e){
			var trCount = Y.all("#queryMain .queryLine").size();
			if (trCount > 1) {
				var myAnim = new Y.Anim({
					node: '#queryContent',
					to: {
						height: 26 * trCount
					},
					duration: duration
				});
				myAnim.run();
			}
			Y.one("#btnMore").setStyle("display", "none");
			Y.one("#btnUp").setStyle("display", "");
		});
		Y.one("#btnUp").on("click", function(e){
			var myAnim = new Y.Anim({
				node: '#queryContent',
				to: {
					height: 22
				},
				duration: duration
			});
			myAnim.run();
			Y.one("#btnMore").setStyle("display", "");
			Y.one("#btnUp").setStyle("display", "none");
		});
	});
}
