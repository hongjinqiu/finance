Y.LTextareaField = Y.Base.create('l-textarea-field', Y.RTextareaField, [Y.WidgetChild], {
    bindUI: function() {
    	Y.LTextareaField.superclass.bindUI.apply(this, arguments);
    	var self = this;
    	
    	new LFormManager().applyEventBehavior(self);
    }
},
{

    ATTRS: {
    	
    }
});
