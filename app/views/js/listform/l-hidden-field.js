Y.LHiddenField = Y.Base.create('l-hidden-field', Y.RHiddenField, [Y.WidgetChild], {
    bindUI: function() {
    	Y.LHiddenField.superclass.bindUI.apply(this, arguments);
    	var self = this;
    	
    	new LFormManager().applyEventBehavior(self);
    }
},
{

    ATTRS: {
    	
    }
});
