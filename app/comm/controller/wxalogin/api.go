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
	Uid:       "uid",
	Code:      "code",
	Genre:     "genre",
}

type user struct {
	TableName string
	Id        string
	Uid       string
	Code      string
	Genre     string
}
