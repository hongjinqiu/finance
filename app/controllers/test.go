package controllers

import "github.com/robfig/revel"
import (
	"math/rand"
	"time"
	"fmt"
)

type Test struct {
	*revel.Controller
}

func (c Test) Index() revel.Result {
	time.Sleep(time.Millisecond * 500)
	return c.RenderText("index0")
}

func (c Test) Index1() revel.Result {
	for i := 0; i <= 21474836; i++ {
		if i == rand.Int() {
			println(i)
		}
	}
	return c.RenderText("index1")
}

func (c Test) Index2() revel.Result {
	time.Sleep(time.Second)
	return c.RenderText("index2")
}

func (c Test) StringTest() revel.Result {
	return c.RenderText(fmt.Sprint(nil))
}
