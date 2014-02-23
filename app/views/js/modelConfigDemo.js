var sysUserModel = {
	"A" : {
		"nick" : {
			listeners: {}
		},
		"jsConfig": {
			listeners: {}
		}
	},
	"B" : {
		"attachCount" : {
			displayField : "",// 可以为函数
			valueField : "",// 可以为函数
			selectorName : "",// 可以为函数
			selectionMode : "single",// single|multi
			selection: function(datas){},// 单多选回调
			unSelection: function(){},// 单多选回调
			queryFunc: function(){},// 单多选回调
			listeners : {// 会用yui.on调用,
				focus: function(e){},
				blur: function(e){},
				tabchange: function(e){},//暂时没有实现
				select: function(e){},
				reset: function(e){},//暂时没有实现
				confirm: function(e){},//暂时没有实现
				change: function(e){},
				dblclick: function(e){},
				keydown: function(e){},
				click: function(e){}
			},
			formatter: function(o){},// 数据集字段函数,接受o作为参数,
			defaultValueExprForJs : function(bo, data) {},// 整个业务对象,单行数据
			calcValueExprForJs : function(bo, data) {}// 整个业务对象,单行数据,
//			validate: function(bo, data) {}// 业务的validate,覆盖不了字段上的validate方法,放到带数据集上处理,
		},
		afterNewData: function(dataSource, bo){},// defaultValueExprForJs->calcValueExprForJs->afterNewData->calcValueExprForJs
		listeners: {},
		beforeedit: function(editor, e){},//数据集函数,表格控件函数,参数:(editor,e),editor里面有record,
		validateedit: function(editor, e){},//数据集函数,表格控件函数,参数:(editor,e),editor里面有record,
		edit: function(editor, e){},//数据集函数,表格控件函数,参数:(editor,e)
		canceledit: function(editor, e){},//数据集函数,表格控件函数,参数:(editor,e)
		celldblclick: function() {},// 暂不实现
		validate: function(bo, data) {}// 整条记录的validate
	}
};



