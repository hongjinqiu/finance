function CommonUtil() {
}

CommonUtil.prototype.getFuncOrString = function(text) {
	if (/^[a-zA-Z\d_]*$/.test(text)) {
		if (eval("typeof(" + text + ")") == "function") {
			return eval(text);
		}
	}
	return text;
}

var panelZIndex = 6;

/**
 * config: title,url
 */
function showModalDialog(config) {
	YUI(g_financeModule).use("finance-module", function(Y) {
		var title = config["title"];
		var url = config["url"];
	//	node.getComputedStyle("width")
		var width = 700;
		var height = 500;
			var node = Y.one("window");
	//		width = parseInt(node.getComputedStyle("width"));
	//		height = parseInt(node.getComputedStyle("height"));
			width = parseInt(node.get("winWidth"));
			height = parseInt(node.get("winHeight"));
		var frameWidth = width - 40;
		if (frameWidth <= 0) {
			frameWidth = 100;
		}
		var frameHeight = height - 40;
		if (frameHeight <= 0) {
			frameHeight = 100;
		}
//	var bodyContent = null;
		var bodyContent = "<iframe src='{src}' frameborder='0' style='width:100%;height:100%;overflow: auto;'></iframe>";
		bodyContent = Y.Lang.sub(bodyContent, {
			src: url
//			,width: frameWidth
//			,height: frameHeight
		});
	    var dialog = new Y.Panel({
	        contentBox : Y.Node.create('<div id="dialog" />'),
	        headerContent: title,
	        bodyContent: bodyContent,
	        width      : frameWidth,
	        height: frameHeight,
	        zIndex     : (++panelZIndex),
	        centered   : true,
	        modal      : true, // modal behavior
	        render     : '.popupModelDialog',
	        visible    : false, // make visible explicitly with .show()
	        plugins      : [Y.Plugin.Drag],
	        buttons: [
	                  {
	                      value: "close",// string or html string
	                      action: function(e) {
	                          e.preventDefault();
	                          dialog.hide();
	                      },
	                      section: Y.WidgetStdMod.HEADER
	                  }
	              ]
	    });

	    dialog.hide = function() {
	    	window.s_dialog = null;
			return this.destroy();
		}
	    
	    dialog.dd.addHandle('.yui3-widget-hd');
	    dialog.show();
	    window.s_dialog = dialog;
	});
}

/**
 * infoType:info,error,question,warn
 */
function showDialog(config){
	var infoType = config["infoType"];
	var title = config["title"];
	var msg = config["msg"];
	var callback = config["callback"];
	var width = config["width"] || 410;
	var height = config["height"] || 150;
	var bodyHeight = height - 23 - 40 - 50;
	var bodyContent = null;
	var footer = [];
	if (infoType == "info") {
		bodyContent = '<div class="message icon-info overflowAuto" style="height:' + bodyHeight + 'px;">' + msg + '</div>';
		footer = [{
            name     : 'proceed',
            label    : '确定',
            action   : 'onOK'
        }];
	} else if (infoType == "success") {
		bodyContent = '<div class="message icon-success overflowAuto" style="height:' + bodyHeight + 'px;">' + msg + '</div>';
		footer = [{
            name     : 'proceed',
            label    : '确定',
            action   : 'onOK'
        }];
	} else if (infoType == "warn") {
		bodyContent = '<div class="message icon-warn overflowAuto" style="height:' + bodyHeight + 'px;">' + msg + '</div>';
		footer = [{
            name     : 'proceed',
            label    : '确定',
            action   : 'onOK'
        }];
	} else if (infoType == "question") {
		bodyContent = '<div class="message icon-question overflowAuto" style="height:' + bodyHeight + 'px;">' + msg + '</div>';
		footer = [{
            name  : 'cancel',
            label : '取消',
            action: 'onCancel'
        }, {
            name     : 'proceed',
            label    : '确定',
            action   : 'onOK'
        }];
	} else if (infoType == "error") {
		bodyContent = '<div class="message icon-error overflowAuto" style="height:' + bodyHeight + 'px;">' + msg + '</div>';
		footer = [{
            name     : 'proceed',
            label    : '确定',
            action   : 'onOK'
        }];
	}
	
	YUI(g_financeModule).use("finance-module", function(Y) {
	    var dialog = new Y.Panel({
	        contentBox : Y.Node.create('<div id="dialog" />'),
	        headerContent: title,
	        bodyContent: bodyContent,
	        width      : width,
	        height: height,
	        zIndex     : (++panelZIndex),
	        centered   : true,
	        modal      : true, // modal behavior
	        render     : '.popupDialog',
	        visible    : false, // make visible explicitly with .show()
	        plugins      : [Y.Plugin.Drag],
	        buttons    : {
	        	footer: footer
	        }
	    });

	    dialog.onCancel = function (e) {
	        e.preventDefault();
	        this.hide();
	    }

	    dialog.onOK = function (e) {
	        e.preventDefault();
	        this.hide();
	        if (callback) {
	        	callback();
	        }
	    }
	    
	    dialog.hide = function() {
			return this.destroy();
		}
	    
	    dialog.dd.addHandle('.yui3-widget-hd');
	    dialog.show();
	    if (infoType == "info" || infoType == "success" || infoType == "warn" || infoType == "error") {
	    	dialog.getButton("proceed").focus();
	    }
	});
}

function showAlert(msg, callback, width, height){
	showDialog({
		"infoType": "info",
		"title": "提示信息",
		"msg": msg,
		"callback": callback,
		"width": width,
		"height": height
	});
}

function showSuccess(msg, callback, width, height){
	showDialog({
		"infoType": "success",
		"title": "成功信息",
		"msg": msg,
		"callback": callback,
		"width": width,
		"height": height
	});
}

function showError(msg, callback, width, height){
	showDialog({
		"infoType": "error",
		"title": "错误信息",
		"msg": msg,
		"callback": callback,
		"width": width,
		"height": height
	});
}

function showWarning(msg, callback, width, height){
	showDialog({
		"infoType": "warn",
		"title": "警告信息",
		"msg": msg,
		"callback": callback,
		"width": width,
		"height": height
	});
}

function showConfirm(msg, callback, width, height){
	showDialog({
		"infoType": "question",
		"title": "确认信息",
		"msg": msg,
		"callback": callback,
		"width": width,
		"height": height
	});
}

/**
 * 配置demo:
 * {
 * 	sync: true | false,
 * 	method: GET | POST,
 * 	params: post data,
 * 	callback: success callback function,
 * }
 */
function ajaxRequest(option){
	// 有用的配置为 doCallback, 自己对failure,error进行提示即可,
	// url,params,async,scope,
	YUI(g_financeModule).use("finance-module", function(Y){
//		var paramData = Y.JSON.stringify(option["params"]);
		var paramData = {};
		if (option.params) {
			for (var k in option.params) {
				if (typeof(option.params[k]) == "object") {
					paramData[k] = Y.JSON.stringify(option.params[k]);
				} else {
					paramData[k] = option.params[k];
				}
			}
		}
		var cfg = {
			sync: option.sync !== undefined ? option.sync : true,
			method: option.method || 'POST',
			data: Y.QueryString.stringify(paramData),
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded',
			},
			on: {
				start: function(){
//								console.log("start");
				},
				complete: function(){
//								console.log("complete");
				},
				success: function(id, o, args){
					if (option.callback) {
						try {
							if (o.responseText) {
								option.callback(Y.JSON.parse(o.responseText));
							} else {
								option.callback({});
							}
						} catch (e) {
							console.error(option["url"]);
							console.error(e);
						}
					}
				},
				failure: function(id, o, args){// failure调用在complete之前,
					var text = o.responseText;
					var reg = /panic\(&#34;(.*?)&#34;\)/.test(text);
					var msg = RegExp.$1;
					if (msg) {
						showError(msg);
					} else {
						showError(text, null, 600, 400);
					}
				},
				end: function(){
//								console.log("end");
				}
			},
		};
//		console.log(Y.QueryString.stringify(paramData));
		Y.io(option["url"], cfg);
//		function complete(id, o, args) {
//			var id = id; // Transaction ID.
//			var data = Y.JSON.parse(o.responseText);
//			
//		};
//		io:complete
//		io:end
//		io:failure
//		io:progress
//		io:start
//		io:success
//		io:xdrReady
//		Y.on('io:complete', complete, Y, []);
//		var request = Y.io(uri);
//		Y.QueryString.stringify
		// import json
	});
}

function g_setMasterFormFieldStatus(status) {
	if (status == "view") {
		for (var key in masterFormFieldDict) {
			
		}
	}
}

function isNumber(value) {
	return /^-?\d*(\.\d*)?$/.test(value);
}





