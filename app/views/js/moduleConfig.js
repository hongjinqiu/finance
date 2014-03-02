var g_financeModule = {
	modules: {
        "papersns-form": {
            fullpath: '/app/FormJS',
            requires: ['node', 'widget-base', 'widget-htmlparser', 'io-form', 'widget-parent', 'widget-child', 'base-build', 'substitute', 'io-upload-iframe', 'collection']
        },
        "papersns-form-quickedit": {
        	fullpath: '/app/comboview?js/form/p-form-quickedit.js',
            requires: ["datatable-base", "papersns-form"]
        },
        "finance-module": {
        	fullpath: '/app/comboview?js/financeModule.js',
        	requires: ["papersns-form", "papersns-form-quickedit", "node","widget-base","widget-htmlparser","io-form",
        	           "widget-parent","widget-child","base-build","substitute","io-upload-iframe","collection","overlay",
        	           "calendar","datatype-date","event","json","datatable","datasource-get","datasource-jsonschema",
        	           "datatable-datasource","datatable-sort","datatable-scroll","cssbutton",
        	           "gallery-datatable-paginator","listtemplate-paginator","datatype-date-format","io-base","anim",
        	           "array-extras","querystring-stringify","cssfonts","dataschema-json","datasource-io",
        	           "model-sync-rest","gallery-paginator-view","tabview","panel","dd-plugin","resize-plugin",
        	           "gallery-layout","gallery-quickedit","gallery-paginator","datatable-base"]
        }
    }
}
