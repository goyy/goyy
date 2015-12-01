package user

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(roleCtl.ApiIndex(), roleCtl.Index)
	xhttp.POST(roleCtl.ApiIndex(), roleCtl.Index)
	xhttp.GET(roleCtl.ApiShow(), roleCtl.Show)
	xhttp.POST(roleCtl.ApiAdd(), roleCtl.Add)
	xhttp.POST(roleCtl.ApiEdit(), roleCtl.Edit)
	xhttp.POST(roleCtl.ApiSave(), roleCtl.Save)
	xhttp.POST(roleCtl.ApiDisable(), roleCtl.Disable)
}
