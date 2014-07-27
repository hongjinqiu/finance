var setGridFromServer = true;

function updateDetailB(accountingYearFormObj, numAccountingPeriodFormObj) {
	var accountingYearValue = accountingYearFormObj.get("value");
	var formManager = new FormManager();
	var modelTemplateFactory = new ModelTemplateFactory();
	var numAccountingPeriodIntValue = parseInt(g_masterFormFieldDict["numAccountingPeriod"].get("value"), 10);
	var datas = [];
	executeGYUI(function(Y){
		var date = new Date();
		for (var i = 0; i < numAccountingPeriodIntValue; i++) {
			date.setYear(parseInt(accountingYearValue, 10));
			date.setMonth(i + 1);
			date.setDate(0);
			
			var id = modelTemplateFactory.getSequenceNo();
			var data = formManager.getDataSetNewData("B");
			data["_id"] = id;
			data["id"] = id;
			data["sequenceNo"] = i + 1;
			var numStr = null;
			if ((i + 1) < 10) {
				numStr = accountingYearValue + "0" + (i + 1) + "01";
			} else {
				numStr = accountingYearValue + "" + (i + 1) + "01";
			}
			data["startDate"] = numStr;
			data["endDate"] = Y.DataType.Date.format(date, {
				format: "%Y%m%d"
			});
			datas.push(data);
		}
	});
	g_gridPanelDict["B"].dt.set("data", datas);
}

var modelExtraInfo = {
	"A": {
		"numAccountingPeriod": {
			listeners : {
				valueChange: function(e, formObj) {
					if (setGridFromServer) {
						setGridFromServer = false;
						return;
					}
					var accountingYear = g_masterFormFieldDict["accountingYear"];
					if (formObj.validateField() && accountingYear.validateField()) {
						updateDetailB(accountingYear, formObj);
					}
				}
			}
		},
		"accountingYear": {
			listeners : {
				valueChange: function(e, formObj) {
					if (setGridFromServer) {
						setGridFromServer = false;
						return;
					}
					var numAccountingPeriod = g_masterFormFieldDict["numAccountingPeriod"];
					if (formObj.validateField() && numAccountingPeriod.validateField()) {
						updateDetailB(formObj, numAccountingPeriod);
					}
				}
			}
		},
		validateEdit: function(masterBo) {
			var messageLi = [];
			return messageLi;
		}
	},
	"B": {
		"endDate": {
			listeners : {
				valueChange: function(e, formObj) {
					if (g_gridPanelDict["B_addrow"]) {// 初次设值时,也会触发 valueChange 事件,此时,record中还有没有formFieldDict,
						var recordLi = g_gridPanelDict["B_addrow"].dt.get("data");
						var currentIndex = 0;
						recordLi.each(function(rec, recordIndex) {
							if (rec.formFieldDict && rec.formFieldDict["endDate"] == formObj) {
								currentIndex = recordIndex;
							}
						});
						var startDateValue = formObj.get("value");
						if (startDateValue.length >= 8 && /^\d*$/g.test(startDateValue)) {
							recordLi.each(function(rec, recordIndex) {
								if (recordIndex == (currentIndex + 1) && rec.formFieldDict) {
									var date = new Date();
									date.setFullYear(parseInt(startDateValue.substring(0,4), 10));
									date.setMonth(parseInt(startDateValue.substring(4,6), 10) - 1);
									date.setDate(parseInt(startDateValue.substring(6,8), 10));
									executeGYUI(function(Y) {
										date = Y.DataType.Date.addDays(date, 1);
										var dateValue = Y.DataType.Date.format(date, {
											format: "%Y%m%d"
										})
										rec.formFieldDict["startDate"].set("value", dateValue);
									});
								}
							});
						}
					}
				}
			}
		},
		beforeEdit: function(recordLi, record, recordIndex){
			if (recordIndex != 0) {
				record.formFieldDict["startDate"].set("readonly", true);
			}
		},
		validateEdit: function(jsonDataLi) {
			var messageLi = [];
			for (var i = 0; i < jsonDataLi.length; i++) {
				if (jsonDataLi[i].endDate < jsonDataLi[i].startDate) {
					messageLi.push("序号为" + (i + 1) + "的分录，结束日期不能<开始日期！");
				}
			}
			return messageLi;
		}
	}
};

function editAccountingDetail(dataSetId) {
	var selectRecordLi = g_gridPanelDict[dataSetId].dt.get("data").toArray();
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

function main(Y) {
		if (g_id) {
			if (g_copyFlag == "true") {// 复制
				ajaxRequest({
					url: "/" + g_dataSourceJson.Id + "/CopyData?format=json"
					,params: {
						"dataSourceModelId": g_dataSourceJson.Id,
						"formTemplateId": g_formTemplateJsonData.Id,
						"id": g_id
					},
					callback: function(o) {
						var formManager = new FormManager();
						formManager.setDetailIncId(g_dataSourceJson, o.bo);
						formManager.applyGlobalParamFromAjaxData(o);
						formManager.loadData2Form(g_dataSourceJson, o.bo);
						formManager.setFormStatus("edit");
					}
				});
			} else {
				ajaxRequest({
					url: "/" + g_dataSourceJson.Id + "/GetData?format=json"
					,params: {
						"dataSourceModelId": g_dataSourceJson.Id,
						"formTemplateId": g_formTemplateJsonData.Id,
						"id": g_id
					},
					callback: function(o) {
						var formManager = new FormManager();
						formManager.applyGlobalParamFromAjaxData(o);
						formManager.loadData2Form(g_dataSourceJson, o.bo);
						formManager.setFormStatus(g_formStatus);
					}
				});
			}
		} else {
			ajaxRequest({
				url: "/" + g_dataSourceJson.Id + "/NewData?format=json"
				,params: {
					"dataSourceModelId": g_dataSourceJson.Id,
					"formTemplateId": g_formTemplateJsonData.Id
				},
				callback: function(o) {
					var formManager = new FormManager();
					formManager.applyGlobalParamFromAjaxData(o);
					formManager.loadData2Form(g_dataSourceJson, o.bo);
					formManager.setFormStatus(g_formStatus);
				}
			});
		}
	}
