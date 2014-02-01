package app

import "github.com/robfig/revel"

func SessionAdapter(c *revel.Controller, fc []revel.Filter) {
	c.Session["userId"] = "15"
	fc[0](c, fc[1:])
}

func SessionEffectFilter(c *revel.Controller, fc []revel.Filter) {
	if c.Session["userId"] == "" {
		panic("会话过期，请您重新登录")
	}
	fc[0](c, fc[1:])
}

func init() {// 运行的顺序是从上往下
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		SessionAdapter,
		SessionEffectFilter,
		revel.ActionInvoker,           // Invoke the action.
	}
}
