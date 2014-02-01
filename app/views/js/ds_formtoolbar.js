function masterValidate() {
	var message = "";
	
	
	
	return {
		"result": true,
		"message": message
	};
}

function detailValidate() {
	var message = "";
//	var store = gridPanel.getStore();
//	store.each(function(record, index) {
//		// 预算科目维度为金额时，若金额不为0，单价和数量是均空或者均有值，否则报错
//		if (record.get('BUDGET_DIMENSION') == "01") {// 01:金额,02:数量
//			if (parseFloat(record.get('AMT')) != 0) {
//				var allEmpty = (record.get('UP') == "" || parseFloat(record.get('UP')) == 0) && (record.get('QTY') == "" || parseFloat(record.get('QTY')) == 0);
//				var allNotNull = (record.get('UP') != "" && parseFloat(record.get('UP')) != 0) && (record.get('QTY') != "" || parseFloat(record.get('QTY')) != 0);
//				if (!allEmpty && !allNotNull) {
//					message += "请先完善序号为" + (index + 1) + "的数据，预算科目维度为金额，金额不为0，单价和数量必须是均空或者均有值;<br />";
//				}
//			}
//		}
//		var amtOccurs = parseFloat(record.get("AMT_OCCUR") || "0");
//		if (amtOccurs == 0) {
//			message += "请先完善序号为" + (index + 1) + "的数据，发生额不能为0;<br />";
//		}
//	});
	return {
		"result": true,
		"message": message
	};
}

function editData() {//保存
	ajaxRequest({
		url: "/ActionTest/EditData?format=json"
		,params: {
			"age": 20,
			"dataSourceModelId": "ActionTest",
			"jsonData": {
				"_id": 26,
				"id": 26,
				"A": {
					"_id": 26,
					"id": 26,
					"code": "testCode主数据集测试11",
					"name": "testName主数据集测试11",
				},
				"B": [{
					"_id": 27,
					"id": 27,
					"code": "testCodeB分录B修改测试",
					"name": "testNameB分录B修改测试",
				}],
				"C": [{
					"code": "testCodeC分录C修改测试",
					"name": "testNameC分录C修改测试",
				}]
			}
		},
		callback: function(o) {
			showSuccess("保存数据成功");
		}
	});
}

function saveData() {//保存
	ajaxRequest({
		url: "/ActionTest/SaveData?format=json"
		,params: {
			"age": 20,
			"dataSourceModelId": "ActionTest",
			"jsonData": {
				"A": {
					"code": "testCode主数据集",
					"name": "testName主数据集",
				},
				"B": [{
					"code": "testCodeB分录B",
					"name": "testNameB分录B",
				}],
				"C": [{
					"code": "testCodeC分录C",
					"name": "testNameC分录C",
				}]
			}
		},
		callback: function(o) {
			showSuccess("保存数据成功");
		}
	});
	masterValidResult = masterValidate();
	if (masterValidResult.result) {
		var detailValidateResult = detailValidate();
		if (detailValidateResult.result) {
//			Ext.MessageBox.alert('提示信息', message);
//			return;
		}
		
		// showMask
		// ajaxRequest save
		// hideMask
	}
}

function newData() {//新增
	ajaxRequest({
		url: "/ActionTest/NewData?format=json"
		,params: {
			"dataSourceModelId": "ActionTest"
		},
		callback: function(o) {
			console.log(o);
			showSuccess(o.responseText);
		}
	});
}

function copyData() {
	ajaxRequest({
		url: "/ActionTest/CopyData?format=json"
		,params: {
			"id": 26,
			"dataSourceModelId": "ActionTest"
		},
		callback: function(o) {
			console.log(o);
			showSuccess(o.responseText);
		}
	});
}

function giveUpData() {
	ajaxRequest({
		url: "/ActionTest/GiveUpData?format=json"
		,params: {
			"id": 26,
			"dataSourceModelId": "ActionTest"
		},
		callback: function(o) {
			console.log(o);
			showSuccess(o.responseText);
		}
	});
}

function refreshData() {
	ajaxRequest({
		url: "/ActionTest/RefreshData?format=json"
		,params: {
			"id": 26,
			"dataSourceModelId": "ActionTest"
		},
		callback: function(o) {
			console.log(o);
			showSuccess(o.responseText);
		}
	});
}

function logListData() {
	ajaxRequest({
		url: "/ActionTest/LogList?format=json"
		,params: {
			"id": 26,
			"dataSourceModelId": "ActionTest"
		},
		callback: function(o) {
			console.log(o);
			showSuccess(o.responseText);
		}
	});
}

function cancelData() {
	ajaxRequest({
		url: "/ActionTest/CancelData?format=json"
		,params: {
			"id": 26,
			"dataSourceModelId": "ActionTest"
		},
		callback: function(o) {
			console.log(o);
			showSuccess(o.responseText);
		}
	});
}

function unCancelData() {
	ajaxRequest({
		url: "/ActionTest/UnCancelData?format=json"
		,params: {
			"id": 26,
			"dataSourceModelId": "ActionTest"
		},
		callback: function(o) {
			console.log(o);
			showSuccess(o.responseText);
		}
	});
}

function getData() {
	ajaxRequest({
		url: "/ActionTest/GetData?format=json"
		,params: {
			"id": 26,
			"dataSourceModelId": "ActionTest"
		},
		callback: function(o) {
			console.log(o);
			showSuccess(o.responseText);
		}
	});
}