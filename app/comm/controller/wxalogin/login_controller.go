package wxalogin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET("/wxalogin", ctl.login)
}

func (me *Controller) login(c xhttp.Context) {
	code := c.Param("code")
	r := getWxInfo(code)
	if r.Success {
		var p *xtype.Principal
		if Wx2Principal != nil {
			p = Wx2Principal(r.OpenId, r.UnionId, r.SessionKey)
		} else {
			p = getPrincipal(r)
		}
		err := c.Session().SetPrincipal(*p)
		if err != nil {
			logger.Errorln("login.SetPrincipal err:", err)
		}
	}
	sid := c.Session().Id()
	err := c.JSON(xhttp.StatusOK, res{Success: true, Message: "", Token: sid})
	if err != nil {
		logger.Errorln("response err:", err)
	}
}

func getPrincipal(r *result) *xtype.Principal {
	id, ok := getUserId(r.UnionId)
	if !ok {
		id = createUser(r.UnionId, r.OpenId)
	}
	p := &xtype.Principal{
		Id:        id,
		Code:      r.OpenId,
		Key:       r.SessionKey,
		LoginName: r.UnionId,
		LoginTime: times.NowUnixStr(),
	}
	return p
}

func getWxInfo(code string) *result {
	// http request
	url := "https://api.weixin.qq.com/sns/jscode2session"
	query := fmt.Sprintf("?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code", AppId, AppSecret, code)
	fmt.Println(query)
	resp, err := http.Get(url + query)
	if err != nil {
		logger.Errorln("getWxInfo.Get err:", err)
		return &result{Success: false}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorln("getWxInfo.ReadAll err:", err)
		return &result{Success: false}
	}
	// json unmarshal
	r := &result{Success: true}
	err = json.Unmarshal(body, r)
	if err != nil {
		logger.Errorln("getWxInfo.Unmarshal err:", err)
		return &result{Success: false}
	}
	if strings.IsBlank(r.OpenId) {
		r.Success = false
	} else {
		if strings.IsBlank(r.UnionId) {
			r.UnionId = r.OpenId
		}
	}
	return r
}

type res struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	Message string `json:"message"`
}

type result struct {
	Success    bool   `json:"success"`
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type Controller struct {
	controller.JSONController
}

var ctl = &Controller{}
