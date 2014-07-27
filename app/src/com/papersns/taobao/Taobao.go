package taobao

import "github.com/robfig/revel"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type NoSubscribeError struct {
	Message string
}

func (e NoSubscribeError) Error() string {
	return e.Message
}

type AppCallLimitedError struct {
	Message string
}

func (e AppCallLimitedError) Error() string {
	return e.Message
}

type TopInvalidSessionError struct {
	Message string
}

func (e TopInvalidSessionError) Error() string {
	return e.Message
}

type ServiceNotReadyError struct {
	Message string
}

func (e ServiceNotReadyError) Error() string {
	return e.Message
}

type TopInvalidMethodError struct {
	Message string
}

func (e TopInvalidMethodError) Error() string {
	return e.Message
}

type UserWithoutShopError struct {
	Message string
}

func (e UserWithoutShopError) Error() string {
	return e.Message
}

type TaobaoError struct {
	Message string
}

func (e TaobaoError) Error() string {
	return e.Message
}

type TaobaoInterface struct {
}

func (o TaobaoInterface) GetUserInfo(url string) map[string]interface{} {
	callbackDict := o.GetCallbackDict(url)
	return o.GetAccessToken(callbackDict)
}

func (o TaobaoInterface) GetCallbackDict(callbackUrl string) map[string]string {
	callbackDict := map[string]string{}
	li := strings.Split(callbackUrl, "?")
	callbackDict["url"] = li[0]

	values, err := url.ParseQuery(li[1])
	if err != nil {
		panic(err)
	}
	for k, v := range values {
		callbackDict[k] = strings.Join(v, ",")
	}
	return callbackDict
}

func (o TaobaoInterface) GetAccessToken(result map[string]string) map[string]interface{} {
	code := result["code"]
	clientId := revel.Config.StringDefault("TAOBAO_APP_KEY_FINANCE", "")
	clientSecret := revel.Config.StringDefault("TAOBAO_APP_SECRET_FINANCE", "")
	redirectUri := revel.Config.StringDefault("REDIRECT_URI", "")
	paramDict := map[string]string{
		"code":          code,
		"grant_type":    "authorization_code",
		"client_id":     clientId,
		"client_secret": clientSecret,
		"redirect_uri":  redirectUri,
	}
	return o.GetToken(paramDict)
}

func (o TaobaoInterface) GetToken(paramDict map[string]string) map[string]interface{} {
	values := url.Values{}
	for k, v := range paramDict {
		values.Add(k, v)
	}

	OAUTH_TOKEN_ENV_URL := revel.Config.StringDefault("OAUTH_TOKEN_ENV_URL", "")
	resp, err := http.PostForm(OAUTH_TOKEN_ENV_URL, values)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	rsp := map[string]interface{}{}
	err = json.Unmarshal(body, &rsp)
	if err != nil {
		panic(err)
	}
	log.Println("response content is:" + string(body))
	rspByte, err := json.MarshalIndent(&rsp, "", "\t")
	if err != nil {
		panic(err)
	}
	log.Println("unmarshal response content is:" + string(rspByte))

	oResult := map[string]interface{}{
		"oAuthInfo":   rsp,
		"top_session": rsp["access_token"],
		"app_key":     paramDict["client_id"],
		"top_appkey":  paramDict["client_id"],
		"topParameter": map[string]interface{}{
			"visitor_nick": rsp["taobao_user_nick"],
			"visitor_id":   rsp["taobao_user_id"],
		},
	}
	if rsp["sub_taobao_user_id"] != nil && fmt.Sprint(rsp["sub_taobao_user_id"]) != "" {
		topParameter := oResult["topParameter"].(map[string]interface{})
		oResult["topParameter"] = topParameter
		topParameter["sub_taobao_user_id"] = rsp["sub_taobao_user_id"]
		topParameter["sub_taobao_user_nick"] = rsp["sub_taobao_user_nick"]
	}
	return oResult
}

func (o TaobaoInterface) TaobaoShopGet(taobaoSysDict map[string]interface{}) map[string]interface{} {
	topParameter := taobaoSysDict["topParameter"].(map[string]interface{})
	params := map[string]string{
		"method":  "taobao.shop.get",
		"session": fmt.Sprint(taobaoSysDict["top_session"]),
		"app_key": fmt.Sprint(taobaoSysDict["app_key"]),
		"fields":  "sid,cid,title,nick",
		"nick":    fmt.Sprint(topParameter["visitor_nick"]),
	}
	return o.RepeatGetTaobaoDataTemplate(params)
}

func (o TaobaoInterface) RepeatGetTaobaoDataTemplate(params map[string]string) map[string]interface{} {
	repeatCount := 5
	for runCount := 0; runCount < repeatCount; runCount++ {
		taobaoData, err := o.GetTaobaoDataTemplate(params)
		if err == nil {
			return taobaoData
		}
	}
	return nil
}

func (o TaobaoInterface) GetTaobaoDataTemplate(params map[string]string) (map[string]interface{}, error) {
	urlResult, err := o.ReadURL(params)
	if err != nil {
		return nil, err
	}
	paramDict := urlResult["paramDict"].(map[string]string)
	rsp := urlResult["rsp"].(string)
	parseResult, err := o.ParseRsp(paramDict, rsp)
	if err != nil {
		return nil, err
	}
	return parseResult, nil
}

func (o TaobaoInterface) ReadURL(params map[string]string) (map[string]interface{}, error) {
	paramDict := map[string]string{
		"format":       "json",
		"access_token": params["session"],
		"v":            "2.0",
	}
	for k, v := range params {
		paramDict[k] = v
	}
	paramDictByte, err := json.MarshalIndent(&paramDict, "", "\t")
	if err != nil {
		return nil, err
	}
	log.Println("readurl, param is:" + string(paramDictByte))

	values := url.Values{}
	for k, v := range paramDict {
		values.Add(k, v)
	}

	OAUTH_ENV_URL := revel.Config.StringDefault("OAUTH_ENV_URL", "")
	resp, err := http.PostForm(OAUTH_ENV_URL, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println("ReadURL response content is:" + string(body))
	//	rsp := map[string]interface{}{}
	//	err = json.Unmarshal(body, &rsp)
	//	if err != nil {
	//		panic(err)
	//	}
	//	rspByte, err := json.MarshalIndent(&rsp, "", "\t")
	//	if err != nil {
	//		panic(err)
	//	}
	//	log.Println("ReadURL response content convert to json is:" + string(rspByte))
	return map[string]interface{}{
		"paramDict": paramDict,
		"rsp":       string(body),
	}, nil
}

func (o TaobaoInterface) ParseRsp(paramDict map[string]string, rsp string) (map[string]interface{}, error) {
	rspStruct := map[string]interface{}{}
	err := json.Unmarshal([]byte(rsp), &rspStruct)
	if err != nil {
		return nil, err
	}
	rspStructByte, err := json.MarshalIndent(&rspStruct, "", "\t")
	if err != nil {
		return nil, err
	}
	log.Println("ParseRsp response content convert to json is:" + string(rspStructByte))
	if rspStruct["error_response"] != nil {
		errorResponse := rspStruct["error_response"].(map[string]interface{})
		code := fmt.Sprint(errorResponse["code"])
		if code == "27" { // TOP_INVALID_SESSION_CODE
			return nil, TopInvalidSessionError{Message: "Invalid TOP_INVALID_SESSION_CODE"}
		}
		if code == "isp.receivenum-service-notready" {
			return nil, ServiceNotReadyError{Message: "Top service not ready, will try next login"}
		}
		if code == "7" {
			return nil, AppCallLimitedError{Message: "app call limited exception"}
		}
		if code == "22" {
			return nil, TopInvalidMethodError{Message: "Invalid method"}
		}
		if code == "560" {
			return nil, UserWithoutShopError{Message: "isv.invalid-parameter:user-without-shop"}
		}
		return nil, TaobaoError{Message: "taobao data template return error response" + string(rspStructByte)}
	}
	return rspStruct, nil
}
