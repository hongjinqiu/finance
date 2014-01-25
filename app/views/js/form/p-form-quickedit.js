YUI.add('papersns-form-quickedit', function(Y, NAME) {
	function PFormQuickEdit(config) {
		PFormQuickEdit.superclass.constructor.call(this, config);
	}

	PFormQuickEdit._incId = 0;

	PFormQuickEdit.NAME = "PFormQuickEditPlugin";
	PFormQuickEdit.NS = "pqe";

	PFormQuickEdit.ATTRS = {}

	PFormQuickEdit.getInnerId = function() {
		return ++PFormQuickEdit._incId;
	}

	Y.extend(PFormQuickEdit, Y.Plugin.Base, {
		initializer : function(config) {
			var self = this;
			var host = this.get('host');

			var columns = host.get("columns");
			Y.each(columns, function(rec, i) {
				rec.allowHTML = true;
				rec.formatter = function(o) {
					//console.log(o);
					var id = PFormQuickEdit.getInnerId();
					return "<div id='cell_" + id + "' class='pformquickedit-container pformquickedit-key:" + o.column.key + "'></div>";
				}

			});
			host.set("columns", columns);
			var h = this.afterHostEvent('render', function() {
				var rows = host._tbodyNode.get('children');
				host.get('data').each(function(rec, i) {
					var list = rows.item(i).all('.pformquickedit-container');
					var field_count = list.size();
					for ( var j = 0; j < field_count; j++) {
						var fieldName = self._getColumnKey(list.item(j));
						var textField = new Y.PTextField({
							name : fieldName,
							type : 'text',
							required : true,
							validateInline : true,
							validator : function(value, formFieldObj) {
								if (value != "abc") {
									formFieldObj.set("error", "值必须为abc");
									return false;
								}
								return true;
							}
						});
						if (!host.getRecord(i).formFieldLi) {
							host.getRecord(i).formFieldLi = [];
						}
						if (rec.get(fieldName)) {
							textField.set("value", rec.get(fieldName) + "");
						} else {
							textField.set("value", "");
						}
						host.getRecord(i).formFieldLi.push(textField);
						textField.render("#" + list.item(j).get("id"));
					}
				});
			});
			//host.testAttr = "papersns-form-quickedit";
			// 1.加上 allowHTML,
			// 2.加上 formatter,
		},
		getRecords : function() {
			var self = this;
			var host = this.get('host');
			var result = [];
			host.get("data").each(function(rec, i) {
				var record = {};
				for ( var j = 0; j < host.getRecord(i).formFieldLi.length; j++) {
					var formField = host.getRecord(i).formFieldLi[j];
					record[formField.get("name")] = formField.get("value");
				}
				result.push(record);
			});
			return result;
		},
		_getColumnKey : function(e) {
			var quick_edit_re = /pformquickedit-key:([^\s]+)/;
			var m = quick_edit_re.exec(e.get('className'));
			return m[1];
		}
	});

	Y.namespace("Plugin");
	Y.Plugin.DataTablePFormQuickEdit = PFormQuickEdit;
}, '1.1.0', {
	"requires" : [ "datatable-base", "papersns-form" ]
});
