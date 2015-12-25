package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "COMCMS"
	c.Data["Email"] = "hogenwang@vip.qq.com"
	t := "2015-10-15 18:40:30"
	t2, _ := time.Parse("2006-01-02 15:04:05", t)
	c.Data["T2"] = t2
	c.TplNames = "index.tpl"
}
