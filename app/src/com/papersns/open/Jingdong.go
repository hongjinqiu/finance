package open

import "github.com/robfig/revel"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	. "com/papersns/common"
	. "com/papersns/error"
	"sort"
	"crypto/md5"
	"code.google.com/p/go.text/encoding/simplifiedchinese"
)

type JingdongError struct {
	Message string
}

func (e JingdongError) Error() string {
	return e.Message
}

type JingdongInterface struct {
}

func (o JingdongInterface) GetUserInfo(url string) map[string]interface{} {
	callbackDict := o.GetCallbackDict(url)
	return o.GetAccessToken(callbackDict)
}

func (o JingdongInterface) GetCallbackDict(callbackUrl string) map[string]string {
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

func (o JingdongInterface) GetAccessToken(result map[string]string) map[string]interface{} {
	code := result["code"]
	clientId := revel.Config.StringDefault("JD_APP_KEY_FINANCE", "")
	clientSecret := revel.Config.StringDefault("JD_APP_SECRET_FINANCE", "")
	redirectUri := revel.Config.StringDefault("REDIRECT_URI", "")
	paramDict := map[string]string{
		"code":          code,
		"grant_type":    "authorization_code",
		"client_id":     clientId,
		"client_secret": clientSecret,
		"redirect_uri":  redirectUri,
		"state":         result["state"],
	}
	return o.GetToken(paramDict)
}

func (o JingdongInterface) GetToken(paramDict map[string]string) map[string]interface{} {
	values := url.Values{}
	for k, v := range paramDict {
		values.Add(k, v)
	}
	OAUTH_TOKEN_ENV_URL := revel.Config.StringDefault("JD_OAUTH_TOKEN_ENV_URL", "")
	resp, err := http.PostForm(OAUTH_TOKEN_ENV_URL, values)
	if err != nil {
		log.Print("GetToken request ", OAUTH_TOKEN_ENV_URL, " fail, request parameter is:", values)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	// 对body的内容为gbk,转成 utf8
	gbk := simplifiedchinese.GBK
	transform := gbk.NewDecoder()
	utf8Byte := make([]byte, len(body) * 3)
	n, _, _ := transform.Transform(utf8Byte, body, true)
	
	rsp := map[string]interface{}{}
//	err = json.Unmarshal(body, &rsp)
	err = json.Unmarshal(utf8Byte[:n], &rsp)
	if err != nil {
		log.Print("GetToken convert response to json fail, body is:", string(utf8Byte[:n]))
		panic(err)
	}
	log.Println("response content is:" + string(utf8Byte[:n]))
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
			"visitor_nick": rsp["user_nick"],
			"visitor_id":   rsp["uid"],
		},
	}
	return oResult
}

/**
{
    "jingdong_vender_shop_query_response": {
        "shop_jos_result": {
            "open_time": "",
            "shop_id": "",
            "category_main_name": "",
            "category_main": "",
            "vender_id": "",
            "brief": "",
            "logo_url": "",
            "shop_name": ""
        }
    }
}
*/
func (o JingdongInterface) GetShopInfo(taobaoSysDict map[string]interface{}) map[string]interface{} {
	params := map[string]string{
		"method":  "jingdong.vender.shop.query",
		"access_token": fmt.Sprint(taobaoSysDict["top_session"]),
		"app_key": fmt.Sprint(taobaoSysDict["app_key"]),
	}
	// 格式转化,
	jingDongResult := o.RepeatGetJingdongDataTemplate(params)
	jingdong_vender_shop_query_response := map[string]interface{}{}
	if jingDongResult["jingdong_vender_shop_query_responce"] != nil {
		jingdong_vender_shop_query_response = jingDongResult["jingdong_vender_shop_query_responce"].(map[string]interface{})
	} else if jingDongResult["jingdong_vender_shop_query_response"].(map[string]interface{}) != nil {
		jingdong_vender_shop_query_response = jingDongResult["jingdong_vender_shop_query_response"].(map[string]interface{})
	}
	
	if jingdong_vender_shop_query_response["shop_jos_result"] == nil {
		panic(BusinessError{Message:"店铺不存在"})
	}
	
	shop_jos_result := jingdong_vender_shop_query_response["shop_jos_result"].(map[string]interface{})
	return map[string]interface{}{
		"shop_get_response": map[string]interface{}{
			"shop": map[string]interface{}{
				"sid": shop_jos_result["shop_id"],
				"nick": shop_jos_result["shop_name"],
			},
		},
	}
}

func (o JingdongInterface) RepeatGetJingdongDataTemplate(params map[string]string) map[string]interface{} {
	repeatCount := 5
	var outerErr error = nil
	for runCount := 0; runCount < repeatCount; runCount++ {
		taobaoData, err := o.GetJingdongDataTemplate(params)
		if err == nil {
			return taobaoData
		} else {
			outerErr = err
		}
	}
	panic(outerErr)
}

func (o JingdongInterface) GetJingdongDataTemplate(params map[string]string) (map[string]interface{}, error) {
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

type StringSort struct {
	objLi []string
}

func (o StringSort) Len() int {
	return len(o.objLi)
}

func (o StringSort) Less(i, j int) bool {
	return o.objLi[i] <= o.objLi[j]
}

func (o StringSort) Swap(i, j int) {
	o.objLi[i], o.objLi[j] = o.objLi[j], o.objLi[i]
}

func (o JingdongInterface) getSign(params map[string]string) string {
	appSecret := revel.Config.StringDefault("JD_APP_SECRET_FINANCE", "")
	keyLi := []string{}
	for k, _ := range params {
		keyLi = append(keyLi, k)
	}
	stringSort := StringSort{keyLi}
	sort.Sort(stringSort)
	result := appSecret
	for i := 0; i < len(keyLi); i++ {
		result += keyLi[i] + params[keyLi[i]]
	}
	result += appSecret
	sign := fmt.Sprintf("%x", md5.Sum([]byte(result)))
	return strings.ToUpper(sign)
}

func (o JingdongInterface) ReadURL(params map[string]string) (map[string]interface{}, error) {
	// 添加签名处理,
	paramDict := map[string]string{
		"format":       "json",
		"v":            "2.0",
		"timestamp":	DateUtil{}.GetDateByFormat("2006-01-02 15:04:05"),
	}
	for k, v := range params {
		paramDict[k] = v
	}
	
	paramDict["sign"] = o.getSign(paramDict)
	
	paramDictByte, err := json.MarshalIndent(&paramDict, "", "\t")
	if err != nil {
		return nil, err
	}
	log.Println("readurl, param is:" + string(paramDictByte))

	values := url.Values{}
	for k, v := range paramDict {
		values.Add(k, v)
	}

	OAUTH_ENV_URL := revel.Config.StringDefault("JD_OAUTH_ENV_URL", "")
	resp, err := http.PostForm(OAUTH_ENV_URL, values)
	if err != nil {
		log.Println("readurl ", OAUTH_ENV_URL, " fail, parameter is:", values)
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

func (o JingdongInterface) ParseRsp(paramDict map[string]string, rsp string) (map[string]interface{}, error) {
	rspStruct := map[string]interface{}{}
	err := json.Unmarshal([]byte(rsp), &rspStruct)
	if err != nil {
		log.Println("ParseRsp fail, rsp is:", rsp)
		return nil, err
	}
	rspStructByte, err := json.MarshalIndent(&rspStruct, "", "\t")
	if err != nil {
		return nil, err
	}
	log.Println("ParseRsp response content convert to json is:" + string(rspStructByte))
	if rspStruct["error_response"] != nil {
		//		errorResponse := rspStruct["error_response"].(map[string]interface{})
		//		code := fmt.Sprint(errorResponse["code"])
		return nil, JingdongError{Message: "jingdong data template return error response" + string(rspStructByte)}
	}
	return rspStruct, nil
}
