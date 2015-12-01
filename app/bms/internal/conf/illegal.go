package conf

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.Conf.Illegal.Enable = true
	xhttp.Conf.Illegal.Excludes = excludes
	xhttp.Conf.Illegal.Values = values
}

var (
	excludes = []string{
		"/",
		"/login",
		"/sys/user/info",
	}
	values = []string{
		"getWriter,FileOutputStream,getRuntime,getRequest,getProperty,onabort,onblur,onchange,onclick,ondblclick,onerror,onfocus,onkeydown,onkeypress,onkeyup,onload,onmousedown,onmousemove,onmouseout,onmouseover,onmouseup,onreset,onresize,onselect,onsubmit,onunload,script,frameset,&lt;object,document.,.cookie,.href,alert(,expression(,$.get,$.post,$.ajax,",
	}
)
