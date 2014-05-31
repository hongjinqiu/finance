var listTemplateExtraInfo = {
	"ColumnModel" : {

	},
	"QueryParameter" : {
		
	}
};

function main() {
	YUI(g_financeModule).use("finance-module", function(Y){
		if (g_formDataJson && g_defaultBo) {
			for (var key in g_formDataJson) {
				g_defaultBo[key] = g_formDataJson[key];
			}
		}
		//Y.one("#queryParameters").setStyle("display", "none");
	});
}
