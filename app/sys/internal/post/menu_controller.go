package post

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(menuCtl.ApiIndex(), menuCtl.Index, menuCtl.PermitView())
	xhttp.POST(menuCtl.ApiIndex(), menuCtl.Index, menuCtl.PermitView())
	xhttp.GET(menuCtl.ApiShow(), menuCtl.Show, menuCtl.PermitView())
	xhttp.POST(menuCtl.ApiAdd(), menuCtl.Add, menuCtl.PermitAdd())
	xhttp.POST(menuCtl.ApiEdit(), menuCtl.Edit, menuCtl.PermitEdit())
	xhttp.POST(menuCtl.ApiSave(), menuCtl.SaveAndTx, menuCtl.PermitAdd(), menuCtl.PermitEdit())
	xhttp.POST(menuCtl.ApiDisable(), menuCtl.DisableAndTx, menuCtl.PermitDisable())
}
