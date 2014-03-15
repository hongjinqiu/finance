Y.LChoiceField = Y.Base.create('l-choice-field', Y.RChoiceField, [Y.WidgetChild], {
	initializer : function () {
		Y.LChoiceField.superclass.initializer.apply(this, arguments);
		var self = this;
		
		var choiceFieldManager = new ChoiceFieldManager();
		this.set("choices", choiceFieldManager.getChoices(self.get("name")));
    }
},
{

    ATTRS: {
    	
    }
});
