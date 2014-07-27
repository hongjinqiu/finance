function showReply() {
	document.getElementById("replyDiv").style.display = "";
	document.getElementById("content").focus();
}

function reply() {
	var bo = {
		"A": {
			"type": document.getElementById("type").value,
			"content": document.getElementById("content").value,
			"bbsPostId": g_masterFormFieldDict["bbsPostId"].get("value")
		}
	};
	ajaxRequest({
		url: "/BbsPost/SaveData?format=json"
		,params: {
			"dataSourceModelId": "BbsPost",
			"formTemplateId": "BbsPost",
			"jsonData": bo
		},
		callback: function(o) {
			showSuccess("保存数据成功");
//			formManager.setFormStatus("view");
//			formManager.applyGlobalParamFromAjaxData(o);
//			formManager.loadData2Form(g_dataSourceJson, o.bo);
			g_relationManager.mergeRelationBo(o.relationBo);
//			datalist
			var trHtmlLi = [];
			trHtmlLi.push("<tr>");
			trHtmlLi.push("	<td>");
			trHtmlLi.push("		{userDisplay}");
			trHtmlLi.push("	</td>");
			trHtmlLi.push("	<td>");
			trHtmlLi.push("		<div>创建时间：{createTime}<br />{content}</div>");
			trHtmlLi.push("	</td>");
			trHtmlLi.push("</tr>");
			var trHtml = trHtmlLi.join("");
			var userDisplay = getUserDisplay(o.bo.A.createBy);
			trHtml = trHtml.replace("{userDisplay}", userDisplay);
			var createTime = o.bo.A.createTime;
			var createTimeDisplay = createTime.substring(0,4) + "-" + createTime.substring(4,6) + "-" + createTime.substring(6,8) + " " + createTime.substring(8,10) + ":" + createTime.substring(10,12) + ":" + createTime.substring(12,14);
			trHtml = trHtml.replace("{createTime}", createTimeDisplay);
			trHtml = trHtml.replace("{content}", o.bo.A.content);
			executeGYUI(function(Y){
				Y.one("#datalist").append(trHtml);
			});
			document.getElementById("replyDiv").style.display = "none";
			document.getElementById("content").value = "";
		}
	});
}

function getUserDisplay(id) {
	var user = g_relationBo["SysUserSelector"][id];
	if (user) {
		return user.code + "," + user.name;
	}
	return "";
}

function main(Y) {
		//applyDateLocale(Y);
		var queryParameterManager = new QueryParameterManager();
		queryParameterManager.applyQueryDefaultValue(Y);
		queryParameterManager.applyFormData(Y);
		queryParameterManager.applyObserveEventBehavior();
}
