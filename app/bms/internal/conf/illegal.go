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
		"/home.html",
	}
	values = []string{
		"getWriter,FileOutputStream,getRuntime,getRequest,getProperty,script,frameset,iframe,<marquee,<object,document.,.cookie,.href,alert(,confirm(,prompt(,expression(,$.get,$.post,$.ajax,touchstart,touchmove,touchend,touchcancel,gesturestart,gesturechange,gestureend,onorientationchange,orientationchange,",
		"onabort,onafterprint,onbeforeonload,onbeforeprint,onbeforeunload,onblur,oncanplay,oncanplaythrough,onchange,onclick,onconte,oncontextmenu,ondblclick,ondrag,ondragend,ondragenter,ondragleave,ondragover,ondragstart,ondrop,ondurationchange,onemptied,onended,onerror,onfocus,onformchange,onforminput,onhaschange,oninvalid,oninput,onkeydown,onkeypress,onkeyup,onload,onloadeddata,onloadedmetadata,",
		"onloadstart,onmessage,onmousedown,onmousemove,onmouseout,onmouseover,onmouseup,onmousewheel,onoffline,ononline,onpagehide,onpageshow,onpause,onplay,onplaying,onpopstate,onpropertychange,onprogress,onratechange,onreadystatechange,onredo,onreset,onresize,onscroll,onseeked,onseeking,onselect,onstalled,onstart,onstorage,onsubmit,onsuspend,ontimeupdate,onundo,onunload,onvolumechange,onwaiting,",
	}
)
