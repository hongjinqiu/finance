Y.LDisplayField = Y.Base.create('l-display-field', Y.RDisplayField, [Y.WidgetChild], {
    bindUI: function() {
    	Y.LDisplayField.superclass.bindUI.apply(this, arguments);
    	var self = this;
    	
    	new LFormManager().applyEventBehavior(self);
    }
},
{

    ATTRS: {
    	
    }
});
