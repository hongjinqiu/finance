package controllers

import "github.com/robfig/revel"
import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

/*
Cache-Control:max-age=315360000
Connection:keep-alive
Date:Thu, 19 Sep 2013 08:25:26 GMT
Expires:Sat, 05 Sep 2026 00:00:00 GMT
Server:ATS/3.2.0
Vary:Accept-Encoding
Via:http/1.1 l4.ycs.swp.yahoo.com (ApacheTrafficServer/3.2.0)

js
Connection:Keep-Alive
Content-Encoding:gzip
Content-Type:text/javascript;charset=UTF-8
Date:Thu, 19 Sep 2013 08:27:48 GMT
Server:Apache
Transfer-Encoding:chunked
*/
func (c App) Combo() revel.Result {
	jsPath := revel.Config.StringDefault("JS_PATH", "")
	content := ""
	for k := range c.Params.Query {
		file, err := os.Open(path.Join(jsPath, k))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		
		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		content += string(data) + "\n"
	}

	acceptEncoding := c.Request.Header.Get("Accept-Encoding")
	if strings.Index(acceptEncoding, "gzip") > -1 {
		data := bytes.Buffer{}
		w := gzip.NewWriter(&data)
		w.Write([]byte(content))
		w.Close()

		c.Response.Status = http.StatusOK
		if strings.Index(c.Params.Query.Encode(), ".css") <= -1 {
			c.Response.ContentType = "text/javascript;charset=UTF-8"
		} else {
			c.Response.ContentType = "text/css;charset=UTF-8"
		}
		c.Response.Out.Header().Set("Content-Encoding", "gzip")
		return c.RenderText(data.String())
	}

	c.Response.Status = http.StatusOK
	if strings.Index(c.Params.Query.Encode(), ".css") <= -1 {
		c.Response.ContentType = "text/javascript;charset=UTF-8"
	} else {
		c.Response.ContentType = "text/css;charset=UTF-8"
	}
	return c.RenderText(content)
}
func (c App) ComboView() revel.Result {
	jsPath := revel.Config.StringDefault("COMBO_VIEW_PATH", "")
	content := ""
	for k := range c.Params.Query {
		file, err := os.Open(path.Join(jsPath, k))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		
		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		content += string(data) + "\n"
	}

	acceptEncoding := c.Request.Header.Get("Accept-Encoding")
	if strings.Index(acceptEncoding, "gzip") > -1 {
		data := bytes.Buffer{}
		w := gzip.NewWriter(&data)
		w.Write([]byte(content))
		w.Close()

		c.Response.Status = http.StatusOK
		if strings.Index(c.Params.Query.Encode(), ".css") <= -1 {
			c.Response.ContentType = "text/javascript;charset=UTF-8"
		} else {
			c.Response.ContentType = "text/css;charset=UTF-8"
		}
		c.Response.Out.Header().Set("Content-Encoding", "gzip")
		return c.RenderText(data.String())
	}

	c.Response.Status = http.StatusOK
	if strings.Index(c.Params.Query.Encode(), ".css") <= -1 {
		c.Response.ContentType = "text/javascript;charset=UTF-8"
	} else {
		c.Response.ContentType = "text/css;charset=UTF-8"
	}
	return c.RenderText(content)
}
