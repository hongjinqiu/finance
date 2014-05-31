var listTemplateExtraInfo = {
	"ColumnModel" : {

	},
	"QueryParameter" : {
		
	}
};

function main() {
	
}

function editJump(record) {
	if (record.get("billTypeId") == 1) {// 收款单类型参数
		location.href = "/console/formschema/?@name=BillReceiveTypeParameter&id=" + record.get("id");
	} else {// 付款单类型参数
		location.href = "/console/formschema/?@name=BillPaymentTypeParameter&id=" + record.get("id");
	}
}

function viewJump(record) {
	if (record.get("billTypeId") == 1) {// 收款单类型参数
		location.href = "/console/formschema/?@name=BillReceiveTypeParameter&id=" + record.get("id") + "&formStatus=view";
	} else {// 付款单类型参数
		location.href = "/console/formschema/?@name=BillPaymentTypeParameter&id=" + record.get("id") + "&formStatus=view";
	}
}
