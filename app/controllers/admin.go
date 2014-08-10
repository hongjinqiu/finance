package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/revel"
	"strings"
	"com/papersns/global"
	. "com/papersns/component"
	"fmt"
	"io"
	"crypto/sha1"
)

func init() {
}

type Hjq struct {
	BaseDataAction
}

func (c Hjq) Login() revel.Result {
	if strings.ToLower(c.Request.Method) == "get" {
		//c.Response.ContentType = "text/html; charset=utf-8"
		return c.Render()
	}
	username := c.Params.Get("username")
	password := c.Params.Get("password")
	
	hash := sha1.New()
	_, err := io.WriteString(hash, password)
	if err != nil {
		panic(err)
	}
	encryPassword := fmt.Sprintf("%x", hash.Sum(nil))
	
	sessionId := global.GetSessionId()
	defer global.CloseSession(sessionId)
	
	session, _ := global.GetConnection(sessionId)
	qb := QuerySupport{}
	user, found := qb.FindByMapWithSession(session, "SysUser", map[string]interface{}{
		"A.type": 1,
		"A.name": username,
		"A.password": encryPassword,
	})
	if !found {
		c.Response.ContentType = "text/plain; charset=utf-8"
		return c.RenderText("用户名密码错误")
	}
	c.Session["adminUserId"] = fmt.Sprint(user["id"])
	c.Session["userId"] = fmt.Sprint(user["id"])
	
	return c.Redirect("/console/listschema?@name=LastSessionData&cookie=false")
}

func (c Hjq) BecomeShopUser() revel.Result {
	id := c.Params.Get("id")
	c.Session["userId"] = id
	return c.Redirect("/")
}
