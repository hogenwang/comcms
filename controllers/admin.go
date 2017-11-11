package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.Data["Website"] = "Admin Control Panel"
	c.TplName = "admin/index.tpl"
}
