// generated by xgen -- DO NOT EDIT
package menu

import (
	"gopkg.in/goyy/goyy.v0/web/controller"
)

var ctl = &Controller{
	ClientTreeController: controller.ClientTreeController{
		ClientController: controller.ClientController{
			Settings: controller.Settings{
				Project: "sys",
				Module:  "menu",
				Title:   "MENU",
			},
			DTO: func() interface{} {
				return &DTO{}
			},
			DTOs: func() interface{} {
				return &DTOs{}
			},
		},
	},
}

type Controller struct {
	controller.ClientTreeController
}
