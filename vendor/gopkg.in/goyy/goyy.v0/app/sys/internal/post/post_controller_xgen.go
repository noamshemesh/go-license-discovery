// generated by xgen -- DO NOT EDIT
package post

import (
	"gopkg.in/goyy/goyy.v0/web/controller"
)

var ctl = &Controller{
	JSONTreeController: controller.JSONTreeController{
		JSONController: controller.JSONController{
			Settings: controller.Settings{
				Project: "sys",
				Module:  "post",
				Title:   "POST",
			},
			Mgr: Mgr,
		},
	},
}

type Controller struct {
	controller.JSONTreeController
}