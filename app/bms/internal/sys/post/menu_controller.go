package post

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(menuCtl.PathIndex(), menuCtl.Index)
	xhttp.POST(menuCtl.PathIndex(), menuCtl.Index)
	xhttp.GET(menuCtl.PathShow(), menuCtl.Show)
	xhttp.POST(menuCtl.PathAdd(), menuCtl.Add)
	xhttp.POST(menuCtl.PathEdit(), menuCtl.Edit)
	xhttp.POST(menuCtl.PathSave(), menuCtl.Save)
	xhttp.POST(menuCtl.PathDisable(), menuCtl.Disable)
}
