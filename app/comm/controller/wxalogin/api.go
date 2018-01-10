package wxalogin

import (
	"gopkg.in/goyy/goyy.v0/comm/xtype"
)

var AppId string
var AppSecret string

var Wx2Principal func(openId, unionId, sessionKey string) *xtype.Principal

var User = &user{
	TableName: "sys_user",
	Id:        "id",
	Code:      "code",
	Passwd:    "passwd",
	LoginName: "login_name",
}

type user struct {
	TableName string
	Id        string
	Code      string
	Passwd    string
	LoginName string
}
