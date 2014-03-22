Y.LNumberField = Y.Base.create('l-number-field', Y.RNumberField, [Y.WidgetChild], {
    bindUI: function() {
    	Y.LNumberField.superclass.bindUI.apply(this, arguments);
    	var self = this;
    	
    	new LFormManager().applyEventBehavior(self);
    }
},
{

    ATTRS: {
    	
    }
});
