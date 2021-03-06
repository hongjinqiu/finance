/**
 * @class PChoiceField
 * @extends PFormField
 * @param config {Object} Configuration object
 * @constructor
 * @description A form field which allows one or multiple values from a 
 * selection of choices
 */
Y.PChoiceField = Y.Base.create('p-choice-field', Y.PFormField, [Y.WidgetParent, Y.WidgetChild], {

    LABEL_TEMPLATE: '<span></span>',
    SINGLE_CHOICE: Y.PRadioField,
    MULTI_CHOICE: Y.PCheckboxField,

    /**
     * @method _validateChoices
     * @protected
     * @param {Object} val
     * @description Validates the value passe to the choices attribute
     */
    _validateChoices: function(val) {
        if (!Y.Lang.isArray(val)) {
            Y.log('Choice values must be in an array', 'warn');
            return false;
        }

        var i = 0,
        len = val.length;

        for (; i < len; i++) {
            if (!Y.Lang.isObject(val[i])) {
                Y.log('Choice that is not an object cannot be used', 'warn');
                delete val[i];
                continue;
            }
            if (!val[i].label ||
            (!Y.Lang.isString(val[i].label) && !Y.Lang.isNumber(val[i].value)) ||
            //!val[i].value ||
            (val[i].value === "") ||
            (!Y.Lang.isString(val[i].value) && !Y.Lang.isNumber(val[i].value))) {
                Y.log('Choice without label and value cannot be used', 'warn');
                delete val[i];
                continue;
            }
        }

        return true;
    },

    _renderFieldNode: function() {
        var contentBox = this.get('contentBox'),
            parent = contentBox.one("." + this.FIELD_CLASS),
            choices = this.get('choices'),
            multiple = this.get('multi'),
            fieldType = (multiple === true ? this.MULTI_CHOICE: this.SINGLE_CHOICE);

        if (!parent) {
            parent = contentBox;
        }
        Y.Array.each(choices,
        function(c, i, a) {
            var cfg = {
                checked : c.checked,
                value: c.value,
                id: (this.get('id') + '_choice' + i),
                name: this.get('name'),
                label: c.label
            },
            field = new fieldType(cfg);

            field.render(parent);
        }, this);
        this._fieldNode = parent.all('input');
    },

    _syncFieldNode: function() {
        //var choices = this.get('value').split(',');
    	var choices = (this.get('value') + "").split(',');

        if (choices && choices.length > 0) {
            choices = Y.Array.map(choices, function(choice) {
                return Y.Lang.trim(choice);
            });

            this._fieldNode.each(function(node, index, list) {
                var nodeValue = Y.Lang.trim(node.get('value'));
                if (!!~Y.Array.indexOf(choices, nodeValue)) {
                    node.set('checked', true);
                } else {
                    node.set('checked', false);
                }
            });
        }
    },
    
    _syncReadonly: function(e) {
        var dis = this.get('readonly');
        if (dis === true) {
            this._fieldNode.setAttribute('disabled', 'disabled');
        } else {
            this._fieldNode.removeAttribute('disabled');
        }
    },

    /**
     * @method _afterChoicesChange
     * @description When the available choices for the choice field change,
     *     the old ones are removed and the new ones are rendered.
     */
    _afterChoicesChange: function(event) {
        var contentBox = this.get("contentBox");
        contentBox.all(".yui3-form-field").remove();
        this._renderFieldNode();
    },
    
    initializer: function() {
    	Y.PChoiceField.superclass.initializer.apply(this, arguments);
    	
    	if (g_layerBo) {
    		var name = this.get("name");
    		var dataSetId = this.get("dataSetId");
    		for (var i = 0; i < g_formTemplateJsonData.FormElemLi.length; i++) {
    			var formElem = g_formTemplateJsonData.FormElemLi[i];
    			if (formElem.XMLName.Local == "column-model") {
    				if (formElem.ColumnModel.DataSetId == dataSetId) {
    					if (formElem.ColumnModel.ColumnLi) {
    						for (var j = 0; j < formElem.ColumnModel.ColumnLi.length; j++) {
    							var column = formElem.ColumnModel.ColumnLi[j];
    							if (column.Name == name) {
    								if (g_layerBo[column.Dictionary]) {
    									var choices = [];
    									for (var k = 0; k < g_layerBo[column.Dictionary].length; k++) {
    										var dictionaryItem = g_layerBo[column.Dictionary][k];
    										choices.push({
    											value: dictionaryItem["code"],
    											label: dictionaryItem["name"]
    										});
    									}
    									this.set("choices", choices);
    								}
    							}
    						}
    					}
    				}
    			}
    		}
    	}
    },

    clear: function() {
        this._fieldNode.each(function(node, index, list) {
            node.set('checked', false);
        },
        this);

        this.set('value', '');
    },

    bindUI: function() {
        this._fieldNode.on('change', Y.bind(function(e) {
            var value = '',
                type = this.get('multi') ? 'checkbox' : 'radio';

            this._fieldNode.each(function(node, index, list) {
                if (node.get('type') == type && node.get('checked') === true) {
                    if (value.length > 0) {
                        value += ',';
                    }
                    value += node.get('value');
                }
            }, this);
            this.set('value', value, {fromUI : true});
        },
        this));
        this.after('choicesChange', this._afterChoicesChange);

        this.after('valueChange', function (e) {
            if (!e.fromUI) {
                this._syncFieldNode();
            }
        });
        
        this.after('readonlyChange', Y.bind(function(e) {
        	this._syncReadonly();
        },
        this));
    }

},
{
    ATTRS: {
        /** 
         * @attribute choices
         * @type Array
         * @description The choices to render into this field
         */
        choices: {
            valueFn : function () {
                return [];
            },
            validator: function(val) {
                return this._validateChoices(val);
            }
        },

        /** 
         * @attribute multi
         * @type Boolean
         * @default false
         * @description Set to true to allow multiple values to be selected
         */
        multi: {
            validator: Y.Lang.isBoolean,
            value: false
        },
        value: {
            value: '',
            validator: function(val) {
                return Y.Lang.isString(val) || Y.Lang.isNumber(val);
            }
        }
    }
});
