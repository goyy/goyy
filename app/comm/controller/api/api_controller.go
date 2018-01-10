package home

import (
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET("/token", ctl.token)
}

func (me *Controller) token(c xhttp.Context) {
	sessionId := c.Session().Id()
	r := result.Result{Success: true, Token: sessionId}
	c.JSON(xhttp.StatusOK, r)
}

type Controller struct {
	controller.JSONController
}

var ctl = &Controller{}
