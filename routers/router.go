package routers

import (
	"github.com/astaxie/beego"
	"github.com/hogenwang/comcms/controllers"
	"github.com/hogenwang/comcms/controllers/admin"
)

func init() {
	//注册静态文件
	beego.SetStaticPath("/attach", "attach")
	beego.Router("/", &controllers.MainController{})

	//管理后台
	beego.Router("/admin", &admin.AdminController{}, "*:Index")
	beego.Router("/admin/login", &admin.AdminController{}, "get:Login;post:DoLogin")
	beego.Router("/admin/logout", &admin.AdminController{}, "get:Logout")
	beego.Router("/admin/admin/modpwd", &admin.AdminController{}, "get:ModPwd;post:DoModPwd")
	beego.Router("/admin/admin/list", &admin.AdminController{}, "get:List")
	beego.Router("/admin/admin/add", &admin.AdminController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/admin/edit/:id([0-9]+)", &admin.AdminController{}, "get:Edit;post:DoEdit")
	beego.Router("/admin/admin/del", &admin.AdminController{}, "post:DoDel")

	//UEditor
	beego.Router("/ueditor/go/controller", &admin.UEController{}, "*:UEditor")
	//beego.AutoRouter(&admin.UEController{})
	//Webuploader
	beego.Router("/admin/webupload", &admin.WebuploadController{}, "*:WebUploader")

	//系统配置
	beego.Router("/admin/config", &admin.ConfigController{})

	//后台文章
	beego.Router("/admin/article", &admin.ArticleController{}, "*:List")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/article/edit/:id([0-9]+)", &admin.ArticleController{}, "get:Edit;post:DoEdit")
	beego.Router("/admin/article/del", &admin.ArticleController{}, "post:DoDel")
	//后台文章栏目
	beego.Router("/admin/category", &admin.CategoryController{}, "*:List")
	beego.Router("/admin/category/add", &admin.CategoryController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/category/edit/:id([0-9]+)", &admin.CategoryController{}, "get:Edit;post:DoEdit")
	beego.Router("/admin/category/del", &admin.CategoryController{}, "post:DoDel")

	//后台友情链接
	beego.Router("/admin/link", &admin.LinkController{}, "*:List")
	beego.Router("/admin/link/add", &admin.LinkController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/link/edit/:id([0-9]+)", &admin.LinkController{}, "get:Edit;post:DoEdit")
	beego.Router("/admin/link/del", &admin.LinkController{}, "post:DoDel")
	//后台留言板
	beego.Router("/admin/guestbook", &admin.GuestbookController{}, "*:List")
	beego.Router("/admin/guestbook/view/:id([0-9]+)", &admin.GuestbookController{}, "get:View")
	beego.Router("/admin/guestbook/del", &admin.GuestbookController{}, "post:DoDel")
	//后台广告管理
	beego.Router("/admin/ads", &admin.AdsController{}, "*:List")
	beego.Router("/admin/ads/add", &admin.AdsController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/ads/edit/:id([0-9]+)", &admin.AdsController{}, "get:Edit;post:DoEdit")
	beego.Router("/admin/ads/del", &admin.AdsController{}, "post:DoDel")
}
