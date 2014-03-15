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
		defer file.Close()
		if err != nil {
			panic(err)
		}

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
		if strings.Index(k, ".js") == -1 && strings.Index(k, ".css") == -1 {
			panic("fileName is:" + k + ", expect ends with .js or .css")
		}
		isFileExist := false
		for _, filePath := range strings.Split(jsPath, ":") {
			if _, err := os.Stat(path.Join(filePath, k)); err != nil {
				if os.IsNotExist(err) {
					continue
				}
			}
			isFileExist = true
			file, err := os.Open(path.Join(filePath, k))
			defer file.Close()
			if err != nil {
				panic(err)
			}

			data, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			content += string(data) + "\n"
			break
		}
		if !isFileExist {
			panic(k + " is not exists")
		}
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

func (c App) FormJS() revel.Result {
	jsPath := revel.Config.StringDefault("COMBO_VIEW_PATH", "")
	content := ""
	formJsLi := []string{"js/rootform/r-form-field.js","js/rootform/r-text-field.js","js/rootform/r-hidden-field.js","js/rootform/r-checkbox-field.js","js/rootform/r-radio-field.js","js/rootform/r-choice-field.js","js/rootform/r-select-field.js","js/rootform/r-trigger-field.js","js/rootform/r-number-field.js","js/rootform/r-display-field.js","js/rootform/r-textarea-field.js","js/rootform/r-date-field.js"}
	lFormJsLi := []string{"js/listform/lformcommon.js", "js/listform/l-form-field.js","js/listform/l-text-field.js","js/listform/l-hidden-field.js","js/listform/l-checkbox-field.js","js/listform/l-radio-field.js","js/listform/l-choice-field.js","js/listform/l-select-field.js","js/listform/l-trigger-field.js","js/listform/l-number-field.js","js/listform/l-display-field.js","js/listform/l-textarea-field.js","js/listform/l-date-field.js"}
	pFormJsLi := []string{"js/form/p-form-field.js","js/form/p-text-field.js","js/form/p-hidden-field.js","js/form/p-checkbox-field.js","js/form/p-radio-field.js","js/form/p-choice-field.js","js/form/p-select-field.js","js/form/p-trigger-field.js","js/form/p-number-field.js","js/form/p-display-field.js","js/form/p-textarea-field.js","js/form/p-date-field.js"}
	for _, k := range pFormJsLi {
		formJsLi = append(formJsLi, k)
	}
	for _, k := range lFormJsLi {
		formJsLi = append(formJsLi, k)
	}
	for _, k := range formJsLi {
		if strings.Index(k, ".js") == -1 && strings.Index(k, ".css") == -1 {
			panic("fileName is:" + k + ", expect ends with .js or .css")
		}
		isFileExist := false
		for _, filePath := range strings.Split(jsPath, ":") {
			if _, err := os.Stat(path.Join(filePath, k)); err != nil {
				if os.IsNotExist(err) {
					continue
				}
			}
			isFileExist = true
			file, err := os.Open(path.Join(filePath, k))
			defer file.Close()
			if err != nil {
				panic(err)
			}

			data, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			content += string(data) + "\n"
			break
		}
		if !isFileExist {
			panic(k + " is not exists")
		}
	}
	prefix := "YUI.add('papersns-form', function(Y) {\n"
	suffix := "}, '1.1.0' ,{requires:['node', 'widget-base', 'widget-htmlparser', 'io-form', 'widget-parent', 'widget-child', 'base-build', 'substitute', 'io-upload-iframe', 'collection', 'overlay', 'calendar', 'datatype-date']});\n"
	content = prefix + content + suffix

	acceptEncoding := c.Request.Header.Get("Accept-Encoding")
	if strings.Index(acceptEncoding, "gzip") > -1 {
		data := bytes.Buffer{}
		w := gzip.NewWriter(&data)
		w.Write([]byte(content))
		w.Close()

		c.Response.Status = http.StatusOK
		c.Response.ContentType = "text/javascript;charset=UTF-8"
		c.Response.Out.Header().Set("Content-Encoding", "gzip")
		return c.RenderText(data.String())
	}

	c.Response.Status = http.StatusOK
	c.Response.ContentType = "text/javascript;charset=UTF-8"
	return c.RenderText(content)
}
