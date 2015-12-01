package role

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(postCtl.ApiIndex(), postCtl.Index)
	xhttp.POST(postCtl.ApiIndex(), postCtl.Index)
	xhttp.GET(postCtl.ApiShow(), postCtl.Show)
	xhttp.POST(postCtl.ApiAdd(), postCtl.Add)
	xhttp.POST(postCtl.ApiEdit(), postCtl.Edit)
	xhttp.POST(postCtl.ApiSave(), postCtl.Save)
	xhttp.POST(postCtl.ApiDisable(), postCtl.Disable)
}
