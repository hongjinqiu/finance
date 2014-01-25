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
	return c.Render()
}
