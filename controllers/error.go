package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	url := c.Ctx.Input.URL()
	c.Data["Title"] = "404 Page not found"
	c.Data["URL"] = url
	c.TplName = "404.html"
}
