package sysuser

import "github.com/robfig/revel"
import (
	. "com/papersns/component"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
//	"net/http"
	"os"
)

func init() {
}

type SysUser struct {
	*revel.Controller
}

func (c SysUser) Schema() revel.Result {
	file, err := os.Open("/home/hongjinqiu/goworkspace/src/finance/app/controllers/sysuser/SysUserForm.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	formTemplate := FormTemplate{}
	err = xml.Unmarshal(data, &formTemplate)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}
	{
		for i, _ := range formTemplate.FormElemLi {
			formElem := &formTemplate.FormElemLi[i]
			if formElem.XMLName.Local == "html" {
				formElemXmlData, err := xml.Marshal(formElem)
				if err != nil {
					panic(err)
				}
				err = xml.Unmarshal(formElemXmlData, &(formElem.Html))
				if err != nil {
					panic(err)
				}
			} else if formElem.XMLName.Local == "toolbar" {
				formElemXmlData, err := xml.Marshal(formElem)
				if err != nil {
					panic(err)
				}
				err = xml.Unmarshal(formElemXmlData, &(formElem.Toolbar))
				if err != nil {
					panic(err)
				}
			} else if formElem.XMLName.Local == "column-model" {
				formElemXmlData, err := xml.Marshal(formElem)
				if err != nil {
					panic(err)
				}
				err = xml.Unmarshal(formElemXmlData, &(formElem.ColumnModel))
				if err != nil {
					panic(err)
				}
			}
		}
	}

	// 1.query data,
	// from data-provider
	// from query-parameters

	xmlDataArray, err := xml.MarshalIndent(formTemplate, "", "\t")
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

	jsonDataArray, err := json.MarshalIndent(formTemplate, "", "\t")
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.Render(err)
	}

//	c.Response.Status = http.StatusOK
//	c.Response.ContentType = "text/plain; charset=utf-8"
	result := map[string]interface{}{
		"formTemplate": formTemplate,
	}
	return c.Render(result)
	return c.RenderText(string(jsonDataArray))
	return c.RenderText(string(xmlDataArray))
}
