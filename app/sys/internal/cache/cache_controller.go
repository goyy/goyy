package cache

import (
	"gopkg.in/goyy/goyy.v0/data/cache"
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.POST(ctl.ApiShow(), ctl.Show, ctl.PermitView())
}

// Get cached values
func (me *Controller) Show(c xhttp.Context) {
	val, err := cache.Get(c.Param("key"))
	if err != nil {
		val = ""
	}
	err = c.JSON(xhttp.StatusOK, me.Success(c, val))
	if err != nil {
		me.Error(c, err)
		return
	}
}

var ctl = &Controller{
	JSONController: controller.JSONController{
		Settings: controller.Settings{
			Project: "sys",
			Module:  "cache",
			Title:   "CACHE",
		},
	},
}

type Controller struct {
	controller.JSONController
}
