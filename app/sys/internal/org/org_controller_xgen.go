// generated by xgen -- DO NOT EDIT
package org

import (
	"gopkg.in/goyy/goyy.v0/web/controller"
)

var ctl = &Controller{
	JSONTreeController: controller.JSONTreeController{
		JSONController: controller.JSONController{
			Settings: controller.Settings{
				Project: "sys",
				Module:  "org",
				Title:   "ORG",
			},
			Mgr: Mgr,
		},
	},
}

type Controller struct {
	controller.JSONTreeController
}
