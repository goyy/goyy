package user

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(roleCtl.PathIndex(), roleCtl.Index)
	xhttp.POST(roleCtl.PathIndex(), roleCtl.Index)
	xhttp.GET(roleCtl.PathShow(), roleCtl.Show)
	xhttp.POST(roleCtl.PathAdd(), roleCtl.Add)
	xhttp.POST(roleCtl.PathEdit(), roleCtl.Edit)
	xhttp.POST(roleCtl.PathSave(), roleCtl.Save)
	xhttp.POST(roleCtl.PathDisable(), roleCtl.Disable)
}
