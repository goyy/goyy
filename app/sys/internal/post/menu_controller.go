package post

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(menuCtl.ApiIndex(), menuCtl.Index)
	xhttp.POST(menuCtl.ApiIndex(), menuCtl.Index)
	xhttp.GET(menuCtl.ApiShow(), menuCtl.Show)
	xhttp.POST(menuCtl.ApiAdd(), menuCtl.Add)
	xhttp.POST(menuCtl.ApiEdit(), menuCtl.Edit)
	xhttp.POST(menuCtl.ApiSave(), menuCtl.Save)
	xhttp.POST(menuCtl.ApiDisable(), menuCtl.Disable)
}
