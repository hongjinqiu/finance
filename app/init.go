package app

import "github.com/robfig/revel"

import (
	. "com/papersns/common"
	. "com/papersns/error"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
	"sync"
)

var adminURLLi []string = []string{"/hjq/becomeshopuser"}
var noSessionURLLi []string = []string{"", "/", "/hjq/login", "/app/startruntxnperiod"}
var noCacheURLLi []string = []string{"/app/combo", "/public", "/app/comboview", "/app/FormJS"}
var dateRwLock sync.RWMutex = sync.RWMutex{}
var dateFlag string = ""

func getDateFlag() string {
	dateRwLock.RLock()
	defer dateRwLock.RUnlock()
	
	return dateFlag
}

func setDateFlag() {
	dateRwLock.Lock()
	defer dateRwLock.Unlock()
	
	dateFlag = fmt.Sprint(DateUtil{}.GetCurrentYyyyMMddHHmmss())
}

func SessionAdapter(c *revel.Controller, fc []revel.Filter) {
	userAgent := strings.Join(c.Request.Header["User-Agent"], ",")
	if strings.Index(userAgent, "Firefox") > -1 {
		c.Session["userId"] = "10"
	} else {
		c.Session["userId"] = "20"
	}
	c.Session["adminUserId"] = "1"
	fc[0](c, fc[1:])
}

/*
type URL struct {
    Scheme   string
    Opaque   string    // encoded opaque data
    User     *Userinfo // username and password information
    Host     string    // host or host:port
    Path     string
    RawQuery string // encoded query values, without '?'
    Fragment string // fragment for references, without '#'
}
*/
func SessionEffectFilter(c *revel.Controller, fc []revel.Filter) {
	isAdminURL := false
	for _, item := range adminURLLi {
		if strings.Index(c.Request.URL.Path, item) > -1 {
			isAdminURL = true
			break
		}
	}
	if isAdminURL {
		if c.Session["adminUserId"] == "" {
			panic(BusinessError{Message: "管理员会话过期，请您重新登录"})
		}
	} else {
		isNoSession := false
		for _, item := range noSessionURLLi {
			if item == c.Request.URL.Path {
				isNoSession = true
				break
			}
		}
		if !isNoSession {
			if c.Session["userId"] == "" {
				panic(BusinessError{Message: "用户会话过期，请您重新登录"})
			}
		}
	}
	fc[0](c, fc[1:])
}

func UTF8Filter(c *revel.Controller, fc []revel.Filter) {
	c.Response.ContentType = "text/html; charset=utf-8"
	fc[0](c, fc[1:])
}

func DateFlagFilter(c *revel.Controller, fc []revel.Filter) {
	if revel.Config.StringDefault("mode.dev", "true") == "true" {
		dateUtil := DateUtil{}
		c.Flash.Out["dateFlag"] = fmt.Sprint(dateUtil.GetCurrentYyyyMMddHHmmss())
	} else {
		c.Flash.Out["dateFlag"] = getDateFlag()
	}
	fc[0](c, fc[1:])
}

func CacheControlFilter(c *revel.Controller, fc []revel.Filter) {
	if revel.Config.StringDefault("mode.dev", "true") != "true" {
		for _, item := range noCacheURLLi {
			if strings.Index(c.Request.URL.Path, item) == 0 {
				currentTime := time.Now()
				currentTime = currentTime.Add(time.Hour * 24 * 365 * 10)
				expiresFormat := currentTime.Format("Mon, 02 Jan 2006 15:04:05 GMT")
				c.Response.Out.Header().Set("Cache-Control", "max-age=315360000")
				c.Response.Out.Header().Set("Expires", expiresFormat)
				break
			}
		}
	}
	fc[0](c, fc[1:])
}

func TimeStatFilter(c *revel.Controller, fc []revel.Filter) {
	begin := time.Now()
	resp, err := http.Get("http://localhost:28017/top?text=1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyBegin, err := ioutil.ReadAll(resp.Body)

	defer func() {
		resp, err := http.Get("http://localhost:28017/top?text=1")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		bodyEnd, err := ioutil.ReadAll(resp.Body)

		beginDict := map[string]interface{}{}
		endDict := map[string]interface{}{}
		err = json.Unmarshal(bodyBegin, &beginDict)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bodyEnd, &endDict)
		if err != nil {
			panic(err)
		}

		totalsBegin := beginDict["totals"].(map[string]interface{})
		totalsEnd := endDict["totals"].(map[string]interface{})
		var mongoSpend float64 = 0
		for k1, v1 := range totalsBegin {
			for k2, v2 := range totalsEnd {
				if strings.Index(k1, "aftermarket") > -1 && k1 == k2 {
					v1Dict := v1.(map[string]interface{})
					v2Dict := v2.(map[string]interface{})
					v1Total := v1Dict["total"].(map[string]interface{})
					v2Total := v2Dict["total"].(map[string]interface{})

					v1Time, err := strconv.ParseInt(fmt.Sprint(v1Total["time"]), 10, 64)
					if err != nil {
						panic(err)
					}
					v2Time, err := strconv.ParseInt(fmt.Sprint(v2Total["time"]), 10, 64)
					if err != nil {
						panic(err)
					}
					minus, _ := strconv.ParseFloat(fmt.Sprint(v2Time-v1Time), 64)
					mongoSpend += minus / 1000.0
				}
			}
		}

		end := time.Now()
		println("url:", c.Request.URL.Path)
		totalTime, _ := strconv.ParseFloat(fmt.Sprint(end.UnixNano()-begin.UnixNano()), 64)
		println("total time spend is:", totalTime/1000000000.0)
		println("mongo time spend is:", mongoSpend)
	}()
	fc[0](c, fc[1:])
}

type ProfileWriter struct {
	LogFile *os.File
}

func (o ProfileWriter) Write(p []byte) (n int, err error) {
	n, err = o.LogFile.Write(p)
	if err != nil {
		panic(err)
	}
	return n, err
}

func ProfileFilter(c *revel.Controller, fc []revel.Filter) {
	logFile, err := os.Create("d:/profile.txt")
	if err != nil {
		panic(err)
	}
	profileWriter := ProfileWriter{logFile}
	err = pprof.StartCPUProfile(profileWriter)
	if err != nil {
		panic(err)
	}
	defer func() {
		pprof.StopCPUProfile()
		profileWriter.LogFile.Close()
		println("close file")
	}()
	fc[0](c, fc[1:])
}

func BusinessPanicFilter(c *revel.Controller, fc []revel.Filter) {
	defer func() {
		if x := recover(); x != nil {
			if reflect.TypeOf(x).Name() == "BusinessError" {
				err := x.(BusinessError)
				c.Response.ContentType = "application/json; charset=utf-8"
				c.Result = c.RenderJson(map[string]interface{}{
					"success": false,
					"code":    err.Code,
					"message": err.Error(),
				})
			} else {
				if revel.Config.StringDefault("mode.dev", "true") != "true" {
					log.Print(x, "\n", string(debug.Stack()))
				}
				panic(x)
			}
		}
	}()
	fc[0](c, fc[1:])
}

func init() { // 运行的顺序是从上往下
	//	runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Filters is the default set of global filters.
	setDateFlag()
	revel.Filters = []revel.Filter{
		revel.PanicFilter, // Recover from panics and display an error page instead.
		BusinessPanicFilter,
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		//		SessionAdapter,
		SessionEffectFilter,
		UTF8Filter,
		DateFlagFilter,
		CacheControlFilter,
		//		TimeStatFilter,
		//		ProfileFilter,
		revel.ActionInvoker, // Invoke the action.
	}
}
