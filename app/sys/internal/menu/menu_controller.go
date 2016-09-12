package menu

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(ctl.ApiIndex(), ctl.Index, ctl.PermitView())
	xhttp.POST(ctl.ApiIndex(), ctl.Index, ctl.PermitView())
	xhttp.GET(ctl.ApiShow(), ctl.Show, ctl.PermitView())
	xhttp.POST(ctl.ApiShow(), ctl.Show, ctl.PermitView())
	xhttp.POST(ctl.ApiAdd(), ctl.Add, ctl.PermitAdd())
	xhttp.POST(ctl.ApiEdit(), ctl.Edit, ctl.PermitEdit())
	xhttp.POST(ctl.ApiSave(), ctl.Save, ctl.PermitAdd(), ctl.PermitEdit())
	xhttp.POST(ctl.ApiDisable(), ctl.Disable, ctl.PermitDisable())
	xhttp.GET(ctl.ApiTree(), ctl.Tree, ctl.PermitView())
}
