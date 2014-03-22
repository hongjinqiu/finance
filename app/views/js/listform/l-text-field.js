Y.LTextField = Y.Base.create('l-text-field', Y.RTextField, [Y.WidgetChild], {
    bindUI: function() {
    	Y.LTextField.superclass.bindUI.apply(this, arguments);
    	var self = this;
    	
    	new LFormManager().applyEventBehavior(self);
    }
},
{

    ATTRS: {
    	
    }
});
