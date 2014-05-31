var d = {
	"XMLName" : {
		"Space" : "http://www.papersns.com/template",
		"Local" : "form-template"
	},
	"Id" : "CashAccountInit",
	"DataSourceModelId" : "CashAccountInit",
	"Adapter" : {
		"XMLName" : {
			"Space" : "http://www.papersns.com/template",
			"Local" : "adapter"
		},
		"Name" : "ModelFormTemplateAdapter"
	},
	"Description" : "现金账户初始化表单",
	"Scripts" : "cashaccountinit/cashAccountInitModel.js",
	"ViewTemplate" : {
		"XMLName" : {
			"Space" : "http://www.papersns.com/template",
			"Local" : "view-template"
		},
		"View" : "Console/FormSchema.html"
	},
	"FormElemLi" : [
			{
				"XMLName" : {
					"Space" : "http://www.papersns.com/template",
					"Local" : "toolbar"
				},
				"InnerHTML" : "\r\n\t\t<!-- <button name=\"listBtn\" text=\"列表页\" mode=\"url\" handler=\"/console/listschema?@name=CashAccountInit\" iconCls=\"test\"></button> -->\r\n\t\t<!-- <button name=\"newBtn\" text=\"新增\" mode=\"fn\" handler=\"newData\" iconCls=\"\"></button> -->\r\n\t\t<!-- <button name=\"copyBtn\" text=\"复制\" mode=\"fn\" handler=\"copyData\"></button> -->\r\n\t\t<button name=\"editBtn\" text=\"修改\" mode=\"fn\" handler=\"editData\" iconCls=\"\"></button>\r\n\t\t<button name=\"saveBtn\" text=\"保存\" mode=\"fn\" handler=\"saveData\" iconCls=\"\"></button>\r\n\t\t<button name=\"giveUpBtn\" text=\"放弃\" mode=\"fn\" handler=\"giveUpData\" iconCls=\"\"></button>\r\n\t\t<button name=\"delBtn\" text=\"删除\" mode=\"fn\" handler=\"deleteData\" iconCls=\"\"></button>\r\n\t\t<button name=\"refreshBtn\" text=\"刷新\" mode=\"fn\" handler=\"refreshData\" iconCls=\"\"></button>\r\n\t\t<button name=\"usedQueryBtn\" text=\"被用查询\" mode=\"fn\" handler=\"logList\" iconCls=\"\"></button>\r\n\t",
				"Html" : {
					"XMLName" : {}
				},
				"Toolbar" : {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "toolbar"
					},
					"ButtonGroup" : {
						"XMLName" : {}
					},
					"ButtonLi" : [ {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "button"
						},
						"ButtonAttribute" : {
							"XMLName" : {}
						},
						"Name" : "editBtn",
						"Text" : "修改",
						"Handler" : "editData",
						"Mode" : "fn"
					}, {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "button"
						},
						"ButtonAttribute" : {
							"XMLName" : {}
						},
						"Name" : "saveBtn",
						"Text" : "保存",
						"Handler" : "saveData",
						"Mode" : "fn"
					}, {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "button"
						},
						"ButtonAttribute" : {
							"XMLName" : {}
						},
						"Name" : "giveUpBtn",
						"Text" : "放弃",
						"Handler" : "giveUpData",
						"Mode" : "fn"
					}, {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "button"
						},
						"ButtonAttribute" : {
							"XMLName" : {}
						},
						"Name" : "delBtn",
						"Text" : "删除",
						"Handler" : "deleteData",
						"Mode" : "fn"
					}, {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "button"
						},
						"ButtonAttribute" : {
							"XMLName" : {}
						},
						"Name" : "refreshBtn",
						"Text" : "刷新",
						"Handler" : "refreshData",
						"Mode" : "fn"
					}, {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "button"
						},
						"ButtonAttribute" : {
							"XMLName" : {}
						},
						"Name" : "usedQueryBtn",
						"Text" : "被用查询",
						"Handler" : "logList",
						"Mode" : "fn"
					} ]
				},
				"ColumnModel" : {
					"XMLName" : {},
					"CheckboxColumn" : {
						"XMLName" : {}
					},
					"IdColumn" : {
						"XMLName" : {}
					},
					"Toolbar" : {
						"XMLName" : {},
						"ButtonGroup" : {
							"XMLName" : {}
						}
					},
					"EditorToolbar" : {
						"XMLName" : {},
						"ButtonGroup" : {
							"XMLName" : {}
						}
					}
				}
			},
			{
				"XMLName" : {
					"Local" : "column-model"
				},
				"InnerHTML" : "<checkbox-column><expression></expression></checkbox-column><id-column name=\"id\" text=\"编号\" hideable=\"true\"></id-column><toolbar><button-group></button-group></toolbar><editor-toolbar><button-group></button-group></editor-toolbar><select-column xmlns=\"http://www.papersns.com/template\" name=\"accountId\" text=\"账户名称\" hideable=\"false\" auto=\"true\" colSpan=\"2\" labelWidth=\"15%\" fixReadOnly=\"false\" zeroShowEmpty=\"false\">&#xA;&#x9;&#x9;&#x9;&#xA;&#x9;&#x9;<editor></editor><listeners></listeners><relationDS><relationItem name=\"CashAccount\"><relationExpr mode=\"\">true</relationExpr><jsRelationExpr mode=\"\">true</jsRelationExpr><relationConfig selectorName=\"CashAccountSelector\" displayField=\"code,name\" valueField=\"id\" selectionMode=\"multi\"></relationConfig></relationItem></relationDS><column-model><checkbox-column><expression></expression></checkbox-column><id-column></id-column><toolbar><button-group></button-group></toolbar><editor-toolbar><button-group></button-group></editor-toolbar></column-model><buttons></buttons></select-column>",
				"Html" : {
					"XMLName" : {}
				},
				"Toolbar" : {
					"XMLName" : {},
					"ButtonGroup" : {
						"XMLName" : {}
					}
				},
				"ColumnModel" : {
					"XMLName" : {
						"Space" : "http://www.papersns.com/template",
						"Local" : "column-model"
					},
					"CheckboxColumn" : {
						"XMLName" : {}
					},
					"IdColumn" : {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "id-column"
						},
						"Name" : "id",
						"Text" : "编号",
						"Hideable" : "true"
					},
					"Toolbar" : {
						"XMLName" : {},
						"ButtonGroup" : {
							"XMLName" : {}
						}
					},
					"EditorToolbar" : {
						"XMLName" : {},
						"ButtonGroup" : {
							"XMLName" : {}
						}
					},
					"ColumnLi" : [ {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "select-column"
						},
						"Html" : "\n\t\t\t\n\t\t",
						"Name" : "accountId",
						"Editor" : {
							"XMLName" : {}
						},
						"Listeners" : {
							"XMLName" : {}
						},
						"CRelationDS" : {
							"XMLName" : {},
							"CRelationItemLi" : [ {
								"XMLName" : {},
								"Name" : "CashAccount",
								"CRelationExpr" : {
									"XMLName" : {},
									"Content" : "true"
								},
								"CJsRelationExpr" : {
									"XMLName" : {},
									"Content" : "true"
								},
								"CRelationConfig" : {
									"XMLName" : {},
									"SelectorName" : "CashAccountSelector",
									"DisplayField" : "code,name",
									"ValueField" : "id",
									"SelectionMode" : "multi"
								}
							} ]
						},
						"Text" : "账户名称",
						"Hideable" : "false",
						"Auto" : "true",
						"ColSpan" : "2",
						"LabelWidth" : "15%",
						"FixReadOnly" : "false",
						"ZeroShowEmpty" : "false",
						"ColumnModel" : {
							"XMLName" : {},
							"CheckboxColumn" : {
								"XMLName" : {}
							},
							"IdColumn" : {
								"XMLName" : {}
							},
							"Toolbar" : {
								"XMLName" : {},
								"ButtonGroup" : {
									"XMLName" : {}
								}
							},
							"EditorToolbar" : {
								"XMLName" : {},
								"ButtonGroup" : {
									"XMLName" : {}
								}
							}
						},
						"Buttons" : {
							"XMLName" : {}
						}
					} ],
					"DataSetId" : "A",
					"ColSpan" : "2"
				},
				"RenderTag" : "A_1",
				"DataSetId" : "A",
				"ColSpan" : "2"
			},
			{
				"XMLName" : {
					"Local" : "column-model"
				},
				"InnerHTML" : "<checkbox-column hideable=\"false\" name=\"checkboxSelect\"><expression></expression></checkbox-column><id-column name=\"id\" text=\"编号\" hideable=\"true\"></id-column><toolbar><button-group></button-group><button xmlns=\"http://www.papersns.com/template\" text=\"新增\" iconCls=\"yui3-button btn-show\" handler=\"g_addRow\" mode=\"fn\"><expression></expression><button-attribute></button-attribute></button><button xmlns=\"http://www.papersns.com/template\" text=\"编辑\" iconCls=\"test\" handler=\"g_editRow\" mode=\"fn\"><expression></expression><button-attribute></button-attribute></button><button xmlns=\"http://www.papersns.com/template\" text=\"删除\" iconCls=\"test\" handler=\"g_removeRow\" mode=\"fn\"><expression></expression><button-attribute></button-attribute></button></toolbar><editor-toolbar><button-group></button-group></editor-toolbar><virtual-column xmlns=\"http://www.papersns.com/template\" name=\"FUN_C\" text=\"操作\">&#xA;&#x9;&#x9;&#x9;&#xA;&#x9;&#x9;<editor></editor><listeners></listeners><relationDS></relationDS><column-model><checkbox-column><expression></expression></checkbox-column><id-column></id-column><toolbar><button-group></button-group></toolbar><editor-toolbar><button-group></button-group></editor-toolbar></column-model><buttons><button xmlns=\"http://www.papersns.com/template\" text=\"编辑\" iconCls=\"test\" handler=\"g_editSingleRow\" mode=\"fn\"><expression></expression><button-attribute></button-attribute></button><button xmlns=\"http://www.papersns.com/template\" text=\"复制\" iconCls=\"test\" handler=\"g_copyRow\" mode=\"fn\"><expression></expression><button-attribute></button-attribute></button><button xmlns=\"http://www.papersns.com/template\" text=\"删除\" iconCls=\"test\" handler=\"g_removeSingleRow\" mode=\"fn\"><expression></expression><button-attribute></button-attribute></button></buttons></virtual-column><dictionary-column xmlns=\"http://www.papersns.com/template\" name=\"accountType\" text=\"账户类型\" hideable=\"true\" fixReadOnly=\"true\" zeroShowEmpty=\"false\" dictionary=\"D_ACCOUNT_TYPE\"><editor></editor><listeners></listeners><relationDS></relationDS><column-model><checkbox-column><expression></expression></checkbox-column><id-column></id-column><toolbar><button-group></button-group></toolbar><editor-toolbar><button-group></button-group></editor-toolbar></column-model><buttons></buttons></dictionary-column><select-column xmlns=\"http://www.papersns.com/template\" name=\"accountId\" text=\"账户代码\" hideable=\"false\" fixReadOnly=\"false\" zeroShowEmpty=\"false\">&#xA;&#x9;&#x9;&#x9;&#xA;&#x9;&#x9;<editor></editor><listeners></listeners><relationDS><relationItem name=\"CashAccount\"><relationExpr mode=\"\">true</relationExpr><jsRelationExpr mode=\"\">true</jsRelationExpr><relationConfig selectorName=\"CashAccountSelector\" displayField=\"code\" valueField=\"id\" selectionMode=\"single\"></relationConfig><copyConfig copyColumnName=\"name\" copyValueField=\"name\"></copyConfig><copyConfig copyColumnName=\"currencyTypeId\" copyValueField=\"currencyTypeId\"></copyConfig></relationItem></relationDS><column-model><checkbox-column><expression></expression></checkbox-column><id-column></id-column><toolbar><button-group></button-group></toolbar><editor-toolbar><button-group></button-group></editor-toolbar></column-model><buttons></buttons></select-column><string-column xmlns=\"http://www.papersns.com/template\" name=\"name\" text=\"账户名称\" hideable=\"false\" dsFieldMap=\"CashAccount.A.name\" fixReadOnly=\"true\" zeroShowEmpty=\"false\"><editor></editor><listeners></listeners><relationDS></relationDS><column-model><checkbox-column><expression></expression></checkbox-column><id-column></id-column><toolbar><button-group></button-group></toolbar><editor-toolbar><button-group></button-group></editor-toolbar></column-model><buttons></buttons></string-column><select-column xmlns=\"http://www.papersns.com/template\" name=\"currencyTypeId\" text=\"币别\" hideable=\"false\" fixReadOnly=\"true\" zeroShowEmpty=\"false\"><editor></editor><listeners></listeners><relationDS><relationItem name=\"CurrencyType\"><relationExpr mode=\"\">true</relationExpr><jsRelationExpr mode=\"\">true</jsRelationExpr><relationConfig selectorName=\"CurrencyTypeSelector\" displayField=\"code,name\" valueField=\"id\" selectionMode=\"single\"></relationConfig></relationItem></relationDS><column-model><checkbox-column><expression></expression></checkbox-column><id-column></id-column><toolbar><button-group></button-group></toolbar><editor-toolbar><button-group></button-group></editor-toolbar></column-model><buttons></buttons></select-column><string-column xmlns=\"http://www.papersns.com/template\" name=\"exchangeRateShow\" text=\"汇率\" hideable=\"false\" fixReadOnly=\"true\" zeroShowEmpty=\"false\"><editor></editor><listeners></listeners><relationDS></relationDS><column-model><checkbox-column><expression></expression></checkbox-column><id-column></id-column><toolbar><button-group></button-group></toolbar><editor-toolbar><button-group></button-group></editor-toolbar></column-model><buttons></buttons></string-column><number-column xmlns=\"http://www.papersns.com/template\" name=\"exchangeRate\" text=\"汇率隐藏\" hideable=\"true\" fixReadOnly=\"false\" zeroShowEmpty=\"false\"><editor></editor><listeners></listeners><relationDS></relationDS><column-model><checkbox-column><expression></expression></checkbox-column><id-column></id-column><toolbar><button-group></button-group></toolbar><editor-toolbar><button-group></button-group></editor-toolbar></column-model><buttons></buttons></number-column><number-column xmlns=\"http://www.papersns.com/template\" name=\"amtEarly\" text=\"期初金额\" hideable=\"false\" fixReadOnly=\"false\" zeroShowEmpty=\"false\"><editor></editor><listeners></listeners><relationDS></relationDS><column-model><checkbox-column><expression></expression></checkbox-column><id-column></id-column><toolbar><button-group></button-group></toolbar><editor-toolbar><button-group></button-group></editor-toolbar></column-model><buttons></buttons></number-column>",
				"Html" : {
					"XMLName" : {}
				},
				"Toolbar" : {
					"XMLName" : {},
					"ButtonGroup" : {
						"XMLName" : {}
					}
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
						"Hideable" : "false",
						"Name" : "checkboxSelect"
					},
					"IdColumn" : {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "id-column"
						},
						"Name" : "id",
						"Text" : "编号",
						"Hideable" : "true"
					},
					"Toolbar" : {
						"XMLName" : {
							"Space" : "http://www.papersns.com/template",
							"Local" : "toolbar"
						},
						"ButtonGroup" : {
							"XMLName" : {}
						},
						"ButtonLi" : [ {
							"XMLName" : {
								"Space" : "http://www.papersns.com/template",
								"Local" : "button"
							},
							"ButtonAttribute" : {
								"XMLName" : {}
							},
							"Text" : "新增",
							"IconCls" : "yui3-button btn-show",
							"Handler" : "g_addRow",
							"Mode" : "fn"
						}, {
							"XMLName" : {
								"Space" : "http://www.papersns.com/template",
								"Local" : "button"
							},
							"ButtonAttribute" : {
								"XMLName" : {}
							},
							"Text" : "编辑",
							"IconCls" : "test",
							"Handler" : "g_editRow",
							"Mode" : "fn"
						}, {
							"XMLName" : {
								"Space" : "http://www.papersns.com/template",
								"Local" : "button"
							},
							"ButtonAttribute" : {
								"XMLName" : {}
							},
							"Text" : "删除",
							"IconCls" : "test",
							"Handler" : "g_removeRow",
							"Mode" : "fn"
						} ]
					},
					"EditorToolbar" : {
						"XMLName" : {},
						"ButtonGroup" : {
							"XMLName" : {}
						}
					},
					"ColumnLi" : [
							{
								"XMLName" : {
									"Space" : "http://www.papersns.com/template",
									"Local" : "virtual-column"
								},
								"Html" : "\n\t\t\t\n\t\t",
								"Name" : "FUN_C",
								"Editor" : {
									"XMLName" : {}
								},
								"Listeners" : {
									"XMLName" : {}
								},
								"CRelationDS" : {
									"XMLName" : {}
								},
								"Text" : "操作",
								"ColumnModel" : {
									"XMLName" : {},
									"CheckboxColumn" : {
										"XMLName" : {}
									},
									"IdColumn" : {
										"XMLName" : {}
									},
									"Toolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									},
									"EditorToolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									}
								},
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
												"ButtonAttribute" : {
													"XMLName" : {}
												},
												"Text" : "编辑",
												"IconCls" : "test",
												"Handler" : "g_editSingleRow",
												"Mode" : "fn"
											},
											{
												"XMLName" : {
													"Space" : "http://www.papersns.com/template",
													"Local" : "button"
												},
												"ButtonAttribute" : {
													"XMLName" : {}
												},
												"Text" : "复制",
												"IconCls" : "test",
												"Handler" : "g_copyRow",
												"Mode" : "fn"
											},
											{
												"XMLName" : {
													"Space" : "http://www.papersns.com/template",
													"Local" : "button"
												},
												"ButtonAttribute" : {
													"XMLName" : {}
												},
												"Text" : "删除",
												"IconCls" : "test",
												"Handler" : "g_removeSingleRow",
												"Mode" : "fn"
											} ]
								}
							},
							{
								"XMLName" : {
									"Space" : "http://www.papersns.com/template",
									"Local" : "dictionary-column"
								},
								"Name" : "accountType",
								"Editor" : {
									"XMLName" : {}
								},
								"Listeners" : {
									"XMLName" : {}
								},
								"CRelationDS" : {
									"XMLName" : {}
								},
								"Text" : "账户类型",
								"Hideable" : "true",
								"FixReadOnly" : "true",
								"ZeroShowEmpty" : "false",
								"ColumnModel" : {
									"XMLName" : {},
									"CheckboxColumn" : {
										"XMLName" : {}
									},
									"IdColumn" : {
										"XMLName" : {}
									},
									"Toolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									},
									"EditorToolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									}
								},
								"Dictionary" : "D_ACCOUNT_TYPE",
								"Buttons" : {
									"XMLName" : {}
								}
							},
							{
								"XMLName" : {
									"Space" : "http://www.papersns.com/template",
									"Local" : "select-column"
								},
								"Html" : "\n\t\t\t\n\t\t",
								"Name" : "accountId",
								"Editor" : {
									"XMLName" : {}
								},
								"Listeners" : {
									"XMLName" : {}
								},
								"CRelationDS" : {
									"XMLName" : {},
									"CRelationItemLi" : [ {
										"XMLName" : {},
										"Name" : "CashAccount",
										"CRelationExpr" : {
											"XMLName" : {},
											"Content" : "true"
										},
										"CJsRelationExpr" : {
											"XMLName" : {},
											"Content" : "true"
										},
										"CRelationConfig" : {
											"XMLName" : {},
											"SelectorName" : "CashAccountSelector",
											"DisplayField" : "code",
											"ValueField" : "id",
											"SelectionMode" : "single"
										},
										"CCopyConfigLi" : [
												{
													"XMLName" : {
														"Space" : "http://www.papersns.com/template",
														"Local" : "copyConfig"
													},
													"CopyColumnName" : "name",
													"CopyValueField" : "name"
												},
												{
													"XMLName" : {
														"Space" : "http://www.papersns.com/template",
														"Local" : "copyConfig"
													},
													"CopyColumnName" : "currencyTypeId",
													"CopyValueField" : "currencyTypeId"
												} ]
									} ]
								},
								"Text" : "账户代码",
								"Hideable" : "false",
								"FixReadOnly" : "false",
								"ZeroShowEmpty" : "false",
								"ColumnModel" : {
									"XMLName" : {},
									"CheckboxColumn" : {
										"XMLName" : {}
									},
									"IdColumn" : {
										"XMLName" : {}
									},
									"Toolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									},
									"EditorToolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									}
								},
								"Buttons" : {
									"XMLName" : {}
								}
							},
							{
								"XMLName" : {
									"Space" : "http://www.papersns.com/template",
									"Local" : "string-column"
								},
								"Name" : "name",
								"Editor" : {
									"XMLName" : {}
								},
								"Listeners" : {
									"XMLName" : {}
								},
								"CRelationDS" : {
									"XMLName" : {}
								},
								"Text" : "账户名称",
								"Hideable" : "false",
								"DsFieldMap" : "CashAccount.A.name",
								"FixReadOnly" : "true",
								"ZeroShowEmpty" : "false",
								"ColumnModel" : {
									"XMLName" : {},
									"CheckboxColumn" : {
										"XMLName" : {}
									},
									"IdColumn" : {
										"XMLName" : {}
									},
									"Toolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									},
									"EditorToolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									}
								},
								"Buttons" : {
									"XMLName" : {}
								}
							},
							{
								"XMLName" : {
									"Space" : "http://www.papersns.com/template",
									"Local" : "select-column"
								},
								"Name" : "currencyTypeId",
								"Editor" : {
									"XMLName" : {}
								},
								"Listeners" : {
									"XMLName" : {}
								},
								"CRelationDS" : {
									"XMLName" : {},
									"CRelationItemLi" : [ {
										"XMLName" : {},
										"Name" : "CurrencyType",
										"CRelationExpr" : {
											"XMLName" : {},
											"Content" : "true"
										},
										"CJsRelationExpr" : {
											"XMLName" : {},
											"Content" : "true"
										},
										"CRelationConfig" : {
											"XMLName" : {},
											"SelectorName" : "CurrencyTypeSelector",
											"DisplayField" : "code,name",
											"ValueField" : "id",
											"SelectionMode" : "single"
										}
									} ]
								},
								"Text" : "币别",
								"Hideable" : "false",
								"FixReadOnly" : "true",
								"ZeroShowEmpty" : "false",
								"ColumnModel" : {
									"XMLName" : {},
									"CheckboxColumn" : {
										"XMLName" : {}
									},
									"IdColumn" : {
										"XMLName" : {}
									},
									"Toolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									},
									"EditorToolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									}
								},
								"Buttons" : {
									"XMLName" : {}
								}
							},
							{
								"XMLName" : {
									"Space" : "http://www.papersns.com/template",
									"Local" : "string-column"
								},
								"Name" : "exchangeRateShow",
								"Editor" : {
									"XMLName" : {}
								},
								"Listeners" : {
									"XMLName" : {}
								},
								"CRelationDS" : {
									"XMLName" : {}
								},
								"Text" : "汇率",
								"Hideable" : "false",
								"FixReadOnly" : "true",
								"ZeroShowEmpty" : "false",
								"ColumnModel" : {
									"XMLName" : {},
									"CheckboxColumn" : {
										"XMLName" : {}
									},
									"IdColumn" : {
										"XMLName" : {}
									},
									"Toolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									},
									"EditorToolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									}
								},
								"Buttons" : {
									"XMLName" : {}
								}
							},
							{
								"XMLName" : {
									"Space" : "http://www.papersns.com/template",
									"Local" : "number-column"
								},
								"Name" : "exchangeRate",
								"Editor" : {
									"XMLName" : {}
								},
								"Listeners" : {
									"XMLName" : {}
								},
								"CRelationDS" : {
									"XMLName" : {}
								},
								"Text" : "汇率隐藏",
								"Hideable" : "true",
								"FixReadOnly" : "false",
								"ZeroShowEmpty" : "false",
								"ColumnModel" : {
									"XMLName" : {},
									"CheckboxColumn" : {
										"XMLName" : {}
									},
									"IdColumn" : {
										"XMLName" : {}
									},
									"Toolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									},
									"EditorToolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									}
								},
								"Buttons" : {
									"XMLName" : {}
								}
							},
							{
								"XMLName" : {
									"Space" : "http://www.papersns.com/template",
									"Local" : "number-column"
								},
								"Name" : "amtEarly",
								"Editor" : {
									"XMLName" : {}
								},
								"Listeners" : {
									"XMLName" : {}
								},
								"CRelationDS" : {
									"XMLName" : {}
								},
								"Text" : "期初金额",
								"Hideable" : "false",
								"FixReadOnly" : "false",
								"ZeroShowEmpty" : "false",
								"ColumnModel" : {
									"XMLName" : {},
									"CheckboxColumn" : {
										"XMLName" : {}
									},
									"IdColumn" : {
										"XMLName" : {}
									},
									"Toolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									},
									"EditorToolbar" : {
										"XMLName" : {},
										"ButtonGroup" : {
											"XMLName" : {}
										}
									}
								},
								"Buttons" : {
									"XMLName" : {}
								}
							} ],
					"DataSetId" : "B",
					"Text" : "现金帐户初始数据"
				},
				"DataSetId" : "B",
				"Text" : "现金帐户初始数据"
			} ]
};