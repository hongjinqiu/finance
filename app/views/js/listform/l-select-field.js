Y.LSelectField = Y.Base.create('l-select-field', Y.RSelectField, [Y.WidgetChild], {
	initializer : function () {
		Y.LSelectField.superclass.initializer.apply(this, arguments);
		var self = this;
		
		var choiceFieldManager = new ChoiceFieldManager();
		this.set("choices", choiceFieldManager.getChoices(self.get("name")));
    }
},
{

    ATTRS: {
    	
    }
});
