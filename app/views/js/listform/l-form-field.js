Y.LFormField = Y.Base.create('l-form-field', Y.RFormField, [Y.WidgetChild], {
    bindUI: function() {
    	Y.LFormField.superclass.bindUI.apply(this, arguments);
    	var self = this;
    	
    	new LFormManager().applyEventBehavior(self);
    }
},
{

    ATTRS: {
    	
    }
});
