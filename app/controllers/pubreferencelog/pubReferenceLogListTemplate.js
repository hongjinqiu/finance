var listTemplateExtraInfo = {
	"ColumnModel" : {

	},
	"QueryParameter" : {
		
	}
};

function main(Y) {
		if (g_formDataJson && g_defaultBo) {
			for (var key in g_formDataJson) {
				g_defaultBo[key] = g_formDataJson[key];
			}
		}
		//Y.one("#queryParameters").setStyle("display", "none");
}
