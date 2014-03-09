/**
 * @class RNumberField
 * @extends RTextField
 * @param config {Object} Configuration object
 * @constructor
 * @description A hidden field node
 */
Y.RNumberField = Y.Base.create('r-number-field', Y.RTextField, [Y.WidgetChild], {
    INPUT_TYPE: "text"
},
{
    /**
	 * @property RNumberField.ATTRS
	 * @type Object
	 * @static
	 */
    ATTRS: {
    }

});
