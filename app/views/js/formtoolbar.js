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

function saveData() {//保存
	ajaxRequest({
		url: "/console/listschema?@name=ActionTest&format=json"
		,params: {
			"name": "名称",
			"age": 20,
			"address": "test"
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
