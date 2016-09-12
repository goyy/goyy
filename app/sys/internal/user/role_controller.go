package user

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(roleCtl.ApiIndex(), roleCtl.Index, roleCtl.PermitView())
	xhttp.POST(roleCtl.ApiIndex(), roleCtl.Index, roleCtl.PermitView())
	xhttp.GET(roleCtl.ApiShow(), roleCtl.Show, roleCtl.PermitView())
	xhttp.POST(roleCtl.ApiShow(), roleCtl.Show, roleCtl.PermitView())
	xhttp.POST(roleCtl.ApiAdd(), roleCtl.Add, roleCtl.PermitAdd())
	xhttp.POST(roleCtl.ApiEdit(), roleCtl.Edit, roleCtl.PermitEdit())
	xhttp.POST(roleCtl.ApiSave(), roleCtl.Save, roleCtl.PermitAdd(), roleCtl.PermitEdit())
	xhttp.POST(roleCtl.ApiDisable(), roleCtl.Disable, roleCtl.PermitDisable())
}
