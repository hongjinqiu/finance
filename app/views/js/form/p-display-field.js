Y.PDisplayField = Y.Base.create('p-display-field', Y.RDisplayField, [Y.WidgetChild], {
    bindUI: function() {
    	Y.PDisplayField.superclass.bindUI.apply(this, arguments);
    	
    	var self = this;
    	new FormManager().applyEventBehavior(self, Y);
    },

    _validateReadonly: function(val) {
    	var self = this;
    	return new FormManager().validateReadonly(self, val, Y);
    },
    
    initializer: function() {
    	Y.PDisplayField.superclass.initializer.apply(this, arguments);
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
