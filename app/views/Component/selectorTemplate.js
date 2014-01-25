DataTableManager.prototype.doAfterResponse = function() {
	// this is a total HACK, should figure a better way than Y.later ...
	YUI().use("node", "event", function(Y) {
		Y.later( 25, self, function(){
			syncCheckboxWhenChangeSelection(Y, dtInst.dt);
		} );
	});
}

function syncSelectionWhenChangeCheckbox(Y, dataGrid, nodeLi) {
	nodeLi.each(function(node, index, nodeLi) {
		if (node.get("checked")) {
			var id = dataGrid.getRecord(node).get("id");
			// 是否添加
			var idInputItem = Y.one("#selectionResult .selectionItem input[value='" + id + "']");
			if (!idInputItem) {
				var tempLi = ['<div class="selectionItem">'];
				tempLi.push('<div class="left">{display}</div>');
				tempLi.push('<div class="right" onclick="removeSelection(this)"><input type="hidden" name="selectionId" value="{id}" /></div>');
				tempLi.push('</div>');
				
				var display = Y.Lang.sub(listTemplate.ColumnModel.SelectionTemplate, dataGrid.getRecord(node).getAttrs());
				
				var tempContent = Y.Lang.sub(tempLi.join(""), {
					"display": display,
					"id": id
				});
				
				Y.one("#selectionResult").setHTML(Y.one("#selectionResult").getHTML() + tempContent);
			}
		} else {
			// 是否删除
			var id = dataGrid.getRecord(node).get("id");
			var idInputItem = Y.one("#selectionResult .selectionItem input[value='" + id + "']");
			if (idInputItem) {
				var selectionItemLi = Y.all("#selectionResult .selectionItem");
				selectionItemLi.each(function(selectionItem, selectionItemIndex, selectionItemLi){
					if (selectionItem.one("input[value='" + id + "']")) {
						Y.one("#selectionResult").removeChild(selectionItem);
					}
				});
			}
		}
	});
}

function removeSelection(elem) {
	YUI().use("node", "event", function(Y) {
		Y.one(elem).ancestor(".selectionItem").remove();
		syncCheckboxWhenChangeSelection(Y, dtInst.dt);
	});
}

function syncCheckboxWhenChangeSelection(Y, dataGrid) {
	var selectionInputValueLi = Y.all("#selectionResult .selectionItem input").get("value");
	var checkboxItemCssSelector = dtInst.getCheckboxCssSelector();
	var checkboxItemLi = yInst.all(checkboxItemCssSelector);
	checkboxItemLi.each(function(checkboxItem, index){
		var id = dataGrid.getRecord(checkboxItem).get("id");
		var isSelected = selectionInputValueLi.some(function(value){return value == id});
		if (isSelected) {
			if (!checkboxItem.get("checked")) {
				checkboxItem.set("checked", isSelected);
			}
		} else {
			if (checkboxItem.get("checked")) {
				checkboxItem.set("checked", isSelected);
			}
		}
	});
	var itemCheckLi = checkboxItemLi.get("checked");
	var checkboxAllCssSelector = dtInst.getCheckboxAllCssSelector();
	var checkAllNode = yInst.one(checkboxAllCssSelector);
	var isAllChecked = itemCheckLi.every(function(value){return value});
	if (isAllChecked) {
		if (!checkAllNode.get("checked")) {
			checkAllNode.set("checked", isAllChecked);
		}
	} else {
		if (checkAllNode.get("checked")) {
			checkAllNode.set("checked", isAllChecked);
		}
	}
}

YUI().use("node", "event", function(Y) {
	Y.on("domready", function(e) {
		var dataGrid = dtInst.dt;
		var checkboxItemInnerCssSelector = dtInst.getCheckboxInnerCssSelector();
		var checkboxCssSelector = dtInst.getCheckboxCssSelector();
		dataGrid.delegate("click", function(e) {
			var nodeLi = yInst.all(checkboxCssSelector);
			syncSelectionWhenChangeCheckbox(Y, dataGrid, nodeLi);
		}, checkboxItemInnerCssSelector, dataGrid);

		var checkboxAllInnerCssSelector = dtInst.getCheckboxAllInnerCssSelector();
		dataGrid.delegate("click", function(e) {
			var nodeLi = yInst.all(checkboxCssSelector);
			syncSelectionWhenChangeCheckbox(Y, dataGrid, nodeLi);
		}, checkboxAllInnerCssSelector, dataGrid);
		
		Y.one("#confirmBtn").on("click", function(e){
			syncCheckboxWhenChangeSelection(Y, dataGrid);
		});
	});
});
