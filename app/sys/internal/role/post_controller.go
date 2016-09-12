package role

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(postCtl.ApiIndex(), postCtl.Index, postCtl.PermitView())
	xhttp.POST(postCtl.ApiIndex(), postCtl.Index, postCtl.PermitView())
	xhttp.GET(postCtl.ApiShow(), postCtl.Show, postCtl.PermitView())
	xhttp.POST(postCtl.ApiShow(), postCtl.Show, postCtl.PermitView())
	xhttp.POST(postCtl.ApiAdd(), postCtl.Add, postCtl.PermitAdd())
	xhttp.POST(postCtl.ApiEdit(), postCtl.Edit, postCtl.PermitEdit())
	xhttp.POST(postCtl.ApiSave(), postCtl.Save, postCtl.PermitAdd(), postCtl.PermitEdit())
	xhttp.POST(postCtl.ApiDisable(), postCtl.Disable, postCtl.PermitDisable())
}
