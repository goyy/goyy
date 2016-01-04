package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.Conf.Secure.ForbidUrl = "/err/403.html"
	xhttp.Conf.Secure.SuccessUrl = "/home.html"
	xhttp.Conf.Secure.Filters = []xtype.Map{
		{"/apis/sys/**/(save|saved)", "forbid"},
		{"/apis/**/(disable|disabled)", "forbid"},
		{"/apis/sys/menu/**", "forbid"},
		{"/apis/sys/post/**", "forbid"},
		{"/apis/sys/user/role/**", "forbid"},
		{"/apis/sys/user/(index|show|add|edit|repwd)", "forbid"},
		{"/**", "anon"},
	}
}
