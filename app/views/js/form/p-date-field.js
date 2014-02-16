/**
 * @class PDateField
 * @extends PFormField
 * @param config {Object} Configuration object
 * @constructor
 * @description A form field which allows one or multiple values from a 
 * selection of choices
 */
Y.PDateField = Y.Base.create('p-date-field', Y.PFormField, [Y.WidgetParent, Y.WidgetChild], {
	FIELD_CLASS : 'table-layout-cell trigger_input',
	INPUT_TYPE: "text",
	DATE_TEMPLATE: '<a></a>',
	DATE_CLASS: 'trigger_date',
	_dateNode: null,
	_olNode: null,
	_overlay: null,
	_calendar: null,
	
	_renderDateNode : function () {
        this._dateNode = this._renderNode(this.DATE_TEMPLATE, this.DATE_CLASS);
    },
	
	renderUI: function() {
		Y.PDateField.superclass.renderUI.apply(this, arguments);
		
		this._renderDateNode();
		this._olNode = Y.Node.create('<div></div>');
		this.get('contentBox').appendChild(this._olNode);
    },
    
    _destroyOverlay: function() {
    	if (this._calendar) {
    		this._calendar.destroy();
    		this._calendar = null;
    	}
    	if (this._overlay) {
    		this._overlay.destroy();
    		this._overlay = null;
    	}
    },
    
    _createAndShowOverlay: function() {
    	var self = this;
    	if (this._overlay) {
    		return;
    	}
    	if (this.get("readonly")) {
    		return;
    	}
    	//this._destroyOverlay();
    	
    	this._overlay = new Y.Overlay({
    		bodyContent:"<div></div>",
    		visible : false,
    		width : '250px',
    		zIndex : 1000
    		//,xy : [500, 100]
    	});
    	
    	var node = this._fieldNode;
    	var xy = node.getXY();
    	var x = xy[0];
    	var y = xy[1];
    	//var width = parseInt(node.getComputedStyle("width"));
    	var height = parseInt(node.getComputedStyle("height"));
    	this._overlay.setAttrs({
    		x: x,
    		y: y + height + 4
    	});
    	this._overlay.render(this._olNode);
    	
    	this._calendar = new Y.Calendar({
    		width:'250px',
    		showPrevMonth: true,
    		showNextMonth: true,
    		date: new Date()}).render(this._overlay.get("bodyContent").getDOMNodes()[0]);
    	
    	var dtdate = Y.DataType.Date;
		this._calendar.on("selectionChange", function (ev) {
			var newDate = ev.newSelection[0];
			self.set("value", dtdate.format(newDate));
			self._destroyOverlay();
		});
		this._overlay.show();
    },
    
    _isClickInBoundingBox: function(e) {
    	if (this._calendar) {
    		var isInCalendar = e.target.ancestor("#" + this._calendar.get("boundingBox").get("id"));
    		if (isInCalendar != null) {
    			return true;
    		}
    		var isInFieldBox = e.target.ancestor("#" + this.get("boundingBox").get("id"));
    		if (isInFieldBox != null) {
    			return true;
    		}
    	}
    	return false;
    },
    
    bindUI: function() {
    	var self = this;
    	Y.PDateField.superclass.bindUI.apply(this, arguments);

    	this._fieldNode.on('focus', Y.bind(function () {
    		this._createAndShowOverlay();
		}, this));
    	
    	this._dateNode.on("click", Y.bind(function(e) {
    		if (!this._overlay) {
    			this._createAndShowOverlay();
    		} else {
    			this._destroyOverlay();
    		}
    	}, this));
    	
    	Y.one("document").on("click", Y.bind(function(e){
    		if (!this._isClickInBoundingBox(e)) {
    			this._destroyOverlay();
    		}
    	}, this));
    },
    
    _syncReadonly: function(e) {
    	Y.PDateField.superclass._syncReadonly.apply(this, arguments);
    	
    	var value = this.get('readonly');
        if (value === true) {
        	this._dateNode.setStyle("display", "none");
        	this._destroyOverlay();
        } else {
        	this._dateNode.setStyle("display", "");
        }
    },
    
    _syncDateNode: function() {
    	if (this._dateNode) {
    		this._dateNode.setAttrs({
    			href: "javascript:void(0);",
    			title: "日期选择",
    			id: Y.guid() + Y.PFormField.FIELD_ID_SUFFIX
    		});
    	}
    },
    
    syncUI: function() {
    	Y.PDateField.superclass.syncUI.apply(this, arguments);
    	
    	this._syncDateNode();
    },
    
    initializer: function() {
    	Y.PDateField.superclass.initializer.apply(this, arguments);
    	
    },

},
{
    ATTRS: {
    }
});
