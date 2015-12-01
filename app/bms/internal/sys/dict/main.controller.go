package dict

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(ctl.PathIndex(), ctl.Index)
	xhttp.POST(ctl.PathIndex(), ctl.Index)
	xhttp.GET(ctl.PathShow(), ctl.Show)
	xhttp.POST(ctl.PathAdd(), ctl.Add)
	xhttp.POST(ctl.PathEdit(), ctl.Edit)
	xhttp.POST(ctl.PathSave(), ctl.Save)
	xhttp.POST(ctl.PathDisable(), ctl.Disable)
}
