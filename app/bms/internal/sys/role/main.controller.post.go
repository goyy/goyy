package role

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(postCtl.PathIndex(), postCtl.Index)
	xhttp.POST(postCtl.PathIndex(), postCtl.Index)
	xhttp.GET(postCtl.PathShow(), postCtl.Show)
	xhttp.POST(postCtl.PathAdd(), postCtl.Add)
	xhttp.POST(postCtl.PathEdit(), postCtl.Edit)
	xhttp.POST(postCtl.PathSave(), postCtl.Save)
	xhttp.POST(postCtl.PathDisable(), postCtl.Disable)
}
