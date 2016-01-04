package err

import (
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func (me *Controller) err403(c xhttp.Context) {
	c.HTML(xhttp.StatusOK, "err/403", nil)
}

func (me *Controller) err404(c xhttp.Context) {
	c.HTML(xhttp.StatusOK, "err/404", nil)
}

func (me *Controller) err500(c xhttp.Context) {
	c.HTML(xhttp.StatusOK, "err/500", nil)
}

func init() {
	xhttp.GET("/err/403", ctl.err403)
	xhttp.GET("/err/404", ctl.err404)
	xhttp.GET("/err/500", ctl.err500)
}

type Controller struct {
	controller.ClientController
}

var ctl = &Controller{}
