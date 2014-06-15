package app

import "github.com/robfig/revel"

import (
	"reflect"
	. "com/papersns/error"
//	"runtime"
)

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

func BusinessPanicFilter(c *revel.Controller, fc []revel.Filter) {
	defer func() {
		if x := recover(); x != nil {
			if reflect.TypeOf(x).Name() == "BusinessError" {
				err := x.(BusinessError)
				c.Result = c.RenderJson(map[string]interface{}{
					"success": false,
					"code": err.Code,
					"message": err.Error(),
				})
			} else {
				panic(x)
			}
		}
	}()
	fc[0](c, fc[1:])
}

func init() {// 运行的顺序是从上往下
//	runtime.GOMAXPROCS(1)
//	runtime.GOMAXPROCS(runtime.NumCPU())
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		BusinessPanicFilter,
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
