Y.PNumberField = Y.Base.create('p-number-field', Y.RNumberField, [Y.WidgetChild], {
    bindUI: function() {
    	Y.PNumberField.superclass.bindUI.apply(this, arguments);
    	
    	var self = this;
    	new FormManager().applyEventBehavior(self, Y);
    },

    _validateReadonly: function(val) {
    	var self = this;
    	return new FormManager().validateReadonly(self, val, Y);
    },
    
    initializer: function() {
    	Y.PNumberField.superclass.initializer.apply(this, arguments);
    	var self = this;
    	
    	new FormManager().initializeAttr(self, Y);
    }
},
{

    ATTRS: {
    	dataSetId: {
            validator: Y.Lang.isString,
            writeOnce: true
        }
    }
});
