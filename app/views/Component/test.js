var listTemplate = {
	"XMLName" : {
		"Space" : "http://www.papersns.com/template",
		"Local" : "list-template"
	},
	"ListId" : "SysUser",
	"SelectorId" : "SysUserSelector",
	"Adapter" : {
		"XMLName" : {
			"Space" : "http://www.papersns.com/template",
			"Local" : "adapter"
		},
		"Name" : "name_adapter"
	},
	"Description" : "员工外出申请",
	"Scripts" : "/hr/archives/empsalarycheck/empsalarychecklist.js",
	"ViewTemplate" : {
		"XMLName" : {
			"Space" : "http://www.papersns.com/template",
			"Local" : "view-template"
		},
		"View" : "component/template.ui.ftl"
	},
	"Toolbar" : {
		"XMLName" : {
			"Space" : "http://www.papersns.com/template",
			"Local" : "toolbar"
		},
		"ButtonGroup" : {
			"XMLName" : {
				"Space" : "",
				"Local" : ""
			},
			"ButtonLi" : null
		},
		"ButtonLi" : [
				{
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "button"
					},
					"Expression" : "test_expression",
					"ButtonAttribute" : {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "button-attribute"
						},
						"Name" : "functionId",
						"Value" : "50010"
					},
					"Xtype" : "",
					"Name" : "",
					"Text" : "新增",
					"IconCls" : "test",
					"IconAlign" : "",
					"Disabled" : "",
					"Hidden" : "",
					"ArrowAlign" : "",
					"Scale" : "",
					"Rowspan" : "",
					"Handler" : "doSave",
					"Mode" : "fn"
				},
				{
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "button"
					},
					"Expression" : "",
					"ButtonAttribute" : {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "button-attribute"
						},
						"Name" : "functionId",
						"Value" : "50010"
					},
					"Xtype" : "",
					"Name" : "",
					"Text" : "删除",
					"IconCls" : "",
					"IconAlign" : "",
					"Disabled" : "",
					"Hidden" : "",
					"ArrowAlign" : "",
					"Scale" : "",
					"Rowspan" : "",
					"Handler" : "/oa/selfapply/egress/getEgress.go?EMPEGRESS_ID={EMPEGRESS_ID}",
					"Mode" : "url^"
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "button"
					},
					"Expression" : "",
					"ButtonAttribute" : {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "button-attribute"
						},
						"Name" : "functionId",
						"Value" : "10018"
					},
					"Xtype" : "",
					"Name" : "",
					"Text" : "导出",
					"IconCls" : "",
					"IconAlign" : "",
					"Disabled" : "",
					"Hidden" : "",
					"ArrowAlign" : "",
					"Scale" : "",
					"Rowspan" : "",
					"Handler" : "doExport",
					"Mode" : "url"
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "button"
					},
					"Expression" : "",
					"ButtonAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Xtype" : "",
					"Name" : "",
					"Text" : "获取记录测试",
					"IconCls" : "",
					"IconAlign" : "",
					"Disabled" : "",
					"Hidden" : "",
					"ArrowAlign" : "",
					"Scale" : "",
					"Rowspan" : "",
					"Handler" : "test",
					"Mode" : "fn"
				} ],
		"Export" : "true",
		"Exporter" : "",
		"ExportParam" : "",
		"FreezedHeader" : "",
		"ExportChart" : "",
		"ExcelChart" : "",
		"ExcelChartType" : "",
		"ExportTitle" : "",
		"ExportSuffix" : ""
	},
	"Security" : {
		"XMLName" : {
			"Space" : "http://www.papersns.com/template",
			"Local" : "security"
		},
		"FunctionId" : "10018",
		"Override" : "",
		"DEFAULT_RESOURCE_CODE" : ""
	},
	"BeforeBuildQuery" : "",
	"AfterBuildQuery" : "",
	"AfterQueryData" : "SysUser.afterQueryData",
	"DataProvider" : {
		"XMLName" : {
			"Space" : "http://www.papersns.com/template",
			"Local" : "data-provider"
		},
		"Collection" : "SysUser",
		"FixBsonQuery" : "{\"_id\": {\"$gte\": 15}}",
		"Map" : "",
		"Reduce" : "",
		"Size" : "10",
		"BsonIntercept" : ""
	},
	"ColumnModel" : {
		"XMLName" : {
			"Space" : "http://www.papersns.com/template",
			"Local" : "column-model"
		},
		"CheckboxColumn" : {
			"XMLName" : {
				"Space" : "http://www.papersns.com/template",
				"Local" : "checkbox-column"
			},
			"ColumnAttributeLi" : null,
			"Expression" : "1 == 1",
			"Hideable" : "false",
			"Name" : "checkboxSelect"
		},
		"IdColumn" : {
			"XMLName" : {
				"Space" : "http://www.papersns.com/template",
				"Local" : "id-column"
			},
			"Name" : "_id",
			"Bson" : "",
			"Text" : "编号",
			"Align" : "",
			"Graggable" : "",
			"Groupable" : "",
			"Hideable" : "true",
			"Editable" : "",
			"MenuDisabled" : "",
			"Sortable" : "",
			"Comparable" : "",
			"Locked" : "",
			"Width" : "",
			"ExcelWidth" : "",
			"Renderer" : "",
			"RendererTemplate" : "",
			"SummaryText" : "",
			"SummaryType" : "",
			"Cycle" : "",
			"Exported" : ""
		},
		"ColumnLi" : [
				{
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "virtual-column"
					},
					"Name" : "FUN_C",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "操作",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "40",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "buttons"
						},
						"ButtonLi" : [
								{
									"XMLName" : {
										"Space" : "http://www.papersns.com/template",
										"Local" : "button"
									},
									"Expression" : "1 == 1",
									"ButtonAttribute" : {
										"XMLName" : {
											"Space" : "http://www.papersns.com/template",
											"Local" : "button-attribute"
										},
										"Name" : "code",
										"Value" : "U"
									},
									"Xtype" : "",
									"Name" : "",
									"Text" : "编辑",
									"IconCls" : "bj_btn",
									"IconAlign" : "",
									"Disabled" : "",
									"Hidden" : "",
									"ArrowAlign" : "",
									"Scale" : "",
									"Rowspan" : "",
									"Handler" : "doEdit",
									"Mode" : "fn"
								},
								{
									"XMLName" : {
										"Space" : "http://www.papersns.com/template",
										"Local" : "button"
									},
									"Expression" : "1 == 1",
									"ButtonAttribute" : {
										"XMLName" : {
											"Space" : "",
											"Local" : ""
										},
										"Name" : "",
										"Value" : ""
									},
									"Xtype" : "",
									"Name" : "",
									"Text" : "查看",
									"IconCls" : "ck_btn",
									"IconAlign" : "",
									"Disabled" : "",
									"Hidden" : "",
									"ArrowAlign" : "",
									"Scale" : "",
									"Rowspan" : "",
									"Handler" : "/component/schema.go?@name=DEMO_VIEW&EMPEGRESS_ID={id}",
									"Mode" : "url^"
								} ]
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "string-column"
					},
					"Name" : "id",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "ID",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "string-column"
					},
					"Name" : "nick",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "昵称",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "number-column"
					},
					"Name" : "numTest",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "金额",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "￥",
					"DecimalPlaces" : "2",
					"DecimalSeparator" : ".",
					"ThousandsSeparator" : ",",
					"Suffix" : "",
					"CurrencyField" : "",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "number-column"
					},
					"Name" : "numTest1",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "金额1",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "localCurrency",
					"IsMoney" : "true",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "number-column"
					},
					"Name" : "numTest2",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "金额2",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "currency",
					"IsMoney" : "true",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "number-column"
					},
					"Name" : "numTest3",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "金额3",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "localCurrency",
					"IsMoney" : "",
					"IsUnitPrice" : "true",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "number-column"
					},
					"Name" : "numTest4",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "金额4",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "currency",
					"IsMoney" : "",
					"IsUnitPrice" : "true",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "number-column"
					},
					"Name" : "numTest5",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "金额5",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "localCurrency",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "true",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "number-column"
					},
					"Name" : "numTest6",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "金额6",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "currency",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "true",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "number-column"
					},
					"Name" : "numTest7",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "金额7",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : "true"
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "number-column"
					},
					"Name" : "numTest8",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "金额8",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "￥",
					"DecimalPlaces" : "2",
					"DecimalSeparator" : ".",
					"ThousandsSeparator" : ",",
					"Suffix" : "",
					"CurrencyField" : "",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "number-column"
					},
					"Name" : "currency",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "币别测试",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "true",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "string-column"
					},
					"Name" : "dateTest",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "日期测试",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "string-column"
					},
					"Name" : "APP_NAME",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "申请人",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "true",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "date-column"
					},
					"Name" : "APP_TIME",
					"ColumnAttributeLi" : null,
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "",
					"Text" : "申请日期",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				}, {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "string-column"
					},
					"Name" : "AUDIT_STATUS",
					"ColumnAttributeLi" : [ {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "column-attribute"
						},
						"Name" : "paraName",
						"Value" : "AUDIT_STATUS"
					} ],
					"EditorAttribute" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"Name" : "",
						"Value" : ""
					},
					"Bson" : "'{value}'",
					"Text" : "审核状态",
					"Align" : "",
					"Graggable" : "",
					"Groupable" : "",
					"Hideable" : "",
					"Editable" : "",
					"MenuDisabled" : "",
					"Sortable" : "",
					"Comparable" : "",
					"Locked" : "",
					"Width" : "",
					"ExcelWidth" : "",
					"Renderer" : "",
					"RendererTemplate" : "",
					"SummaryText" : "",
					"SummaryType" : "",
					"Cycle" : "",
					"Exported" : "",
					"Format" : "",
					"DisplayPattern" : "",
					"DbPattern" : "",
					"TrueText" : "",
					"FalseText" : "",
					"UndefinedText" : "",
					"Dictionary" : "",
					"Complex" : "",
					"Buttons" : {
						"XMLName" : {
							"Space" : "",
							"Local" : ""
						},
						"ButtonLi" : null
					},
					"Prefix" : "",
					"DecimalPlaces" : "",
					"DecimalSeparator" : "",
					"ThousandsSeparator" : "",
					"Suffix" : "",
					"CurrencyField" : "",
					"IsMoney" : "",
					"IsUnitPrice" : "",
					"IsCost" : "",
					"IsPercent" : ""
				} ],
		"AutoLoad" : "",
		"SummaryLoad" : "",
		"SummaryStat" : "",
		"Summation" : "",
		"GroupSummation" : "",
		"GroupMerge" : "",
		"ShowGroupFilter" : "",
		"ShowAggregationFilter" : "",
		"AutoRowHeight" : "",
		"Nowrap" : "",
		"Rownumber" : "true",
		"SelectionMode" : "radio",
		"ShowClearBtn" : "",
		"SelectionSupport" : "",
		"GroupField" : "",
		"BsonOrderBy" : "",
		"SaveUrl" : "",
		"DeleteUrl" : "",
		"StoreIntercept" : "",
		"RecordIntercept" : "",
		"SelectionTemplate" : "",
		"SelectionTitle" : ""
	},
	"QueryParameterGroup" : {
		"XMLName" : {
			"Space" : "http://www.papersns.com/template",
			"Local" : "query-parameters"
		},
		"FixedParameterLi" : null,
		"QueryParameterLi" : [ {
			"XMLName" : {
				"Space" : "http://www.papersns.com/template",
				"Local" : "query-parameter"
			},
			"ParameterAttributeLi" : null,
			"Name" : "nick",
			"Text" : "昵称",
			"ColumnName" : "",
			"EnterParam" : "",
			"Editor" : "textfield",
			"Restriction" : "like",
			"ColSpan" : "",
			"RowSpan" : "",
			"Value" : "",
			"OtherName" : "",
			"Having" : "",
			"Required" : "",
			"UseIn" : ""
		}, {
			"XMLName" : {
				"Space" : "http://www.papersns.com/template",
				"Local" : "query-parameter"
			},
			"ParameterAttributeLi" : null,
			"Name" : "dept_id",
			"Text" : "部门",
			"ColumnName" : "",
			"EnterParam" : "",
			"Editor" : "numberfield",
			"Restriction" : "eq",
			"ColSpan" : "",
			"RowSpan" : "",
			"Value" : "",
			"OtherName" : "",
			"Having" : "",
			"Required" : "",
			"UseIn" : ""
		}, {
			"XMLName" : {
				"Space" : "http://www.papersns.com/template",
				"Local" : "query-parameter"
			},
			"ParameterAttributeLi" : null,
			"Name" : "type",
			"Text" : "类型",
			"ColumnName" : "",
			"EnterParam" : "",
			"Editor" : "combotree",
			"Restriction" : "in",
			"ColSpan" : "",
			"RowSpan" : "",
			"Value" : "",
			"OtherName" : "",
			"Having" : "",
			"Required" : "",
			"UseIn" : ""
		}, {
			"XMLName" : {
				"Space" : "http://www.papersns.com/template",
				"Local" : "query-parameter"
			},
			"ParameterAttributeLi" : [ {
				"XMLName" : {
					"Space" : "http://www.papersns.com/template",
					"Local" : "parameter-attribute"
				},
				"Name" : "inFormat",
				"Value" : "yyyy-MM-dd"
			}, {
				"XMLName" : {
					"Space" : "http://www.papersns.com/template",
					"Local" : "parameter-attribute"
				},
				"Name" : "queryFormat",
				"Value" : "yyyyMMdd"
			} ],
			"Name" : "createTimeBegin",
			"Text" : "申请日期从",
			"ColumnName" : "createTime",
			"EnterParam" : "",
			"Editor" : "datefield",
			"Restriction" : "ge",
			"ColSpan" : "",
			"RowSpan" : "",
			"Value" : "",
			"OtherName" : "",
			"Having" : "",
			"Required" : "",
			"UseIn" : ""
		}, {
			"XMLName" : {
				"Space" : "http://www.papersns.com/template",
				"Local" : "query-parameter"
			},
			"ParameterAttributeLi" : [ {
				"XMLName" : {
					"Space" : "http://www.papersns.com/template",
					"Local" : "parameter-attribute"
				},
				"Name" : "inFormat",
				"Value" : "yyyyMMdd"
			}, {
				"XMLName" : {
					"Space" : "http://www.papersns.com/template",
					"Local" : "parameter-attribute"
				},
				"Name" : "queryFormat",
				"Value" : "yyyyMMddHHmmss"
			} ],
			"Name" : "createTimeEnd",
			"Text" : "申请日期从",
			"ColumnName" : "createTime",
			"EnterParam" : "",
			"Editor" : "datefield",
			"Restriction" : "le",
			"ColSpan" : "",
			"RowSpan" : "",
			"Value" : "",
			"OtherName" : "",
			"Having" : "",
			"Required" : "",
			"UseIn" : ""
		}, {
			"XMLName" : {
				"Space" : "http://www.papersns.com/template",
				"Local" : "query-parameter"
			},
			"ParameterAttributeLi" : null,
			"Name" : "_id",
			"Text" : "主键",
			"ColumnName" : "",
			"EnterParam" : "",
			"Editor" : "hidden",
			"Restriction" : "eq",
			"ColSpan" : "",
			"RowSpan" : "",
			"Value" : "",
			"OtherName" : "",
			"Having" : "",
			"Required" : "",
			"UseIn" : ""
		} ],
		"FormColumns" : "",
		"EnableEnterParam" : ""
	}
};