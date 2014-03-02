function editData() {//修改
	var formManager = new FormManager();
	var bo = formManager.getBo();
	ajaxRequest({
		url: "/" + dataSourceJson.Id + "/EditData?format=json"
		,params: {
			"dataSourceModelId": dataSourceJson.Id,
			"id": bo["id"]
		},
		callback: function(o) {
			formManager.loadData2Form(dataSourceJson, o.bo);
			formManager.setFormStatus("edit");
		}
	});
}

function saveData() {//保存
	var formManager = new FormManager();
	var bo = formManager.getBo();
	var validateResult = formManager.dsFormValidator(dataSourceJson, bo);
	if (!validateResult.result) {
		showError(validateResult.message);
	} else {
		ajaxRequest({
			url: "/" + dataSourceJson.Id + "/SaveData?format=json"
			,params: {
				"dataSourceModelId": dataSourceJson.Id,
				"jsonData": bo
			},
			callback: function(o) {
				showSuccess("保存数据成功");
				formManager.setFormStatus("view");
				formManager.loadData2Form(dataSourceJson, o.bo);
			}
		});
	}
}

function newData() {
	var formManager = new FormManager();
	var bo = formManager.getBo();
	ajaxRequest({
		url: "/" + dataSourceJson.Id + "/NewData?format=json"
		,params: {
			"dataSourceModelId": dataSourceJson.Id
		},
		callback: function(o) {
			formManager.loadData2Form(dataSourceJson, o.bo);
			formManager.setFormStatus("edit");
		}
	});
}

function copyData() {
	var formManager = new FormManager();
	var bo = formManager.getBo();
	ajaxRequest({
		url: "/" + dataSourceJson.Id + "/CopyData?format=json"
		,params: {
			"dataSourceModelId": dataSourceJson.Id,
			"id": bo["id"]
		},
		callback: function(o) {
			formManager.loadData2Form(dataSourceJson, o.bo);
			formManager.setFormStatus("edit");
		}
	});
}

function giveUpData() {
	var formManager = new FormManager();
	var bo = formManager.getBo();
	ajaxRequest({
		url: "/" + dataSourceJson.Id + "/GiveUpData?format=json"
		,params: {
			"dataSourceModelId": dataSourceJson.Id,
			"id": bo["id"]
		},
		callback: function(o) {
			formManager.loadData2Form(dataSourceJson, o.bo);
			formManager.setFormStatus("view");
		}
	});
}

function deleteData() {
	showWarning("您确定要删除吗？", function(){
		var formManager = new FormManager();
		var bo = formManager.getBo();
		ajaxRequest({
			url: "/" + dataSourceJson.Id + "/DeleteData?format=json"
			,params: {
				"dataSourceModelId": dataSourceJson.Id,
				"id": bo["id"]
			},
			callback: function(o) {
				location.href = "/console/listschema?@name=" + dataSourceJson.Id;
			}
		});
	})
}

function refreshData() {
	var formManager = new FormManager();
	var bo = formManager.getBo();
	ajaxRequest({
		url: "/" + dataSourceJson.Id + "/RefreshData?format=json"
		,params: {
			"dataSourceModelId": dataSourceJson.Id,
			"id": bo["id"]
		},
		callback: function(o) {
			formManager.loadData2Form(dataSourceJson, o.bo);
			formManager.setFormStatus("view");
		}
	});
}

function logList() {
	var formManager = new FormManager();
	var bo = formManager.getBo();
	ajaxRequest({
		url: "/" + dataSourceJson.Id + "/LogList?format=json"
		,params: {
			"dataSourceModelId": dataSourceJson.Id,
			"id": bo["id"]
		},
		callback: function(o) {
			YUI(g_financeModule).use("finance-module", function(Y) {
				showAlert(Y.JSON.stringify(o));
			});
		}
	});
}

function cancelData() {
	var formManager = new FormManager();
	var bo = formManager.getBo();
	ajaxRequest({
		url: "/" + dataSourceJson.Id + "/CancelData?format=json"
		,params: {
			"dataSourceModelId": dataSourceJson.Id,
			"id": bo["id"]
		},
		callback: function(o) {
			showSuccess("作废数据成功");
			formManager.loadData2Form(dataSourceJson, o.bo);
			formManager.setFormStatus("view");
		}
	});
}

function unCancelData() {
	var formManager = new FormManager();
	var bo = formManager.getBo();
	ajaxRequest({
		url: "/" + dataSourceJson.Id + "/UnCancelData?format=json"
		,params: {
			"dataSourceModelId": dataSourceJson.Id,
			"id": bo["id"]
		},
		callback: function(o) {
			showSuccess("反作废数据成功");
			formManager.loadData2Form(dataSourceJson, o.bo);
			formManager.setFormStatus("view");
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

function test() {
	var relationManager = new RelationManager();
	relationManager.getRelationBo("SysUserSelector", 16);
	return;
}

function ToolbarManager(){}

function setBorderTmp(btn, status) {
	if (status == "disabled") {
		btn.style.border = "1px solid black";
	} else {
		btn.style.border = "1px solid red";
	}
}

ToolbarManager.prototype.enableDisableToolbarBtn = function() {
	if (g_formStatus == "view") {
		var viewEnableBtnLi = ["listBtn","newBtn","copyBtn","editBtn","delBtn","cancelBtn","unCancelBtn","refreshBtn","usedQueryBtn"];
		var viewDisableBtnLi = ["saveBtn","giveUpBtn"];
		for (var i = 0; i < viewEnableBtnLi.length; i++) {
			var btn = document.getElementById(viewEnableBtnLi[i]);
			if (btn) {
				btn.disabled = "";
				setBorderTmp(btn, "");
			}
		}
		var cancelBtn = document.getElementById("cancelBtn");
		if (cancelBtn && masterFormFieldDict["billStatus"]) {
			if (masterFormFieldDict["billStatus"].get("value") == "0") {
				cancelBtn.disabled = "";
				setBorderTmp(cancelBtn, "");
			} else {
				cancelBtn.disabled = "disabled";
				setBorderTmp(cancelBtn, "disabled");
			}
		}
		var unCancelBtn = document.getElementById("unCancelBtn");
		if (unCancelBtn && masterFormFieldDict["billStatus"]) {
			if (masterFormFieldDict["billStatus"].get("value") == "4") {
				unCancelBtn.disabled = "";
				setBorderTmp(unCancelBtn, "");
			} else {
				unCancelBtn.disabled = "disabled";
				setBorderTmp(unCancelBtn, "disabled");
			}
		}
		
		for (var i = 0; i < viewDisableBtnLi.length; i++) {
			var btn = document.getElementById(viewDisableBtnLi[i]);
			if (btn) {
				btn.disabled = "disabled";
				setBorderTmp(btn, "disabled");
			}
		}
	} else {
		var editEnableBtnLi = ["listBtn","saveBtn","giveUpBtn"];
		var editDisableBtnLi = ["newBtn","copyBtn","editBtn","delBtn","cancelBtn","unCancelBtn","refreshBtn","usedQueryBtn"];
		for (var i = 0; i < editEnableBtnLi.length; i++) {
			var btn = document.getElementById(editEnableBtnLi[i]);
			if (btn) {
				btn.disabled = "";
				setBorderTmp(btn, "");
			}
		}
		for (var i = 0; i < editDisableBtnLi.length; i++) {
			var btn = document.getElementById(editDisableBtnLi[i]);
			if (btn) {
				btn.disabled = "disabled";
				setBorderTmp(btn, "disabled");
			}
		}
	}
}





