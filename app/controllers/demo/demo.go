package demo

import "github.com/robfig/revel"
import (
)

func init() {
}

type Demo struct {
	*revel.Controller
}

func (c Demo) Schema() revel.Result {
	//c.Response.ContentType = "text/html; charset=utf-8"
	return c.Render()
}
