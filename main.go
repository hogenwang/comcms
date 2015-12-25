package main

import (
	"github.com/astaxie/beego"
	"github.com/hogenwang/comcms/controllers"
	_ "github.com/hogenwang/comcms/routers"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
