package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/hogenwang/comcms/common"
	"github.com/hogenwang/comcms/models"
	"strconv"
)

type AdminController struct {
	beego.Controller
}

//后台首页
func (c *AdminController) Index() {
	CheckAdminLogin(&c.Controller, 0)
	c.Data["Website"] = "Admin Control Panel"
	c.Data["Title"] = "后台管理首页"
	//c.Layout = "admin/layout.tpl"
	c.TplName = "admin/index.tpl"
}

//登录页面
func (c *AdminController) Login() {
	c.Data["Tip"] = "请输入用户名和密码登录系统"
	c.TplName = "admin/login.tpl"
}

//执行登录
func (c *AdminController) DoLogin() {
	username := c.GetString("username")
	password := c.GetString("password")

	tip := &models.TipJSON{}
	tip.Status = models.TipError

	if len(username) < 4 {
		tip.Message = "请输入最少5个字符的用户名!"
		EchoTip(&c.Controller, tip)
	}
	if len(password) < 5 {
		tip.Message = "请输入最少5个字符的密码!"
		EchoTip(&c.Controller, tip)
	}

	if admin, check := models.AdminLogin(username, password, c.Ctx.Input.IP()); check && admin != nil {
		//成功，则写入Session 如果不是string类型，必须转换成int 之类的，否则，获取不到！！！坑！！
		c.SetSession("adminid", int(admin.Id))
		c.SetSession("adminname", admin.UserName)
		c.SetSession("adminrole", int(admin.RoleId))
		//c.Redirect("/admin", 302)
		tip.Status = models.TipSuccess
		tip.Message = "登录成功!"
		tip.ReturnUrl = "/admin/"
		EchoTip(&c.Controller, tip)
	} else {
		tip.Message = "用户名或者密码错误!"
		EchoTip(&c.Controller, tip)

	}
}

//退出登录
func (c *AdminController) Logout() {
	c.DelSession("adminid")
	c.DelSession("adminname")
	c.DelSession("adminrole")
	c.Redirect("/admin/login", 302)
}

//修改密码页面
func (c *AdminController) ModPwd() {
	CheckAdminLogin(&c.Controller, 0)
	username, _ := c.GetSession("adminname").(string)
	c.Data["Title"] = "修改密码"
	c.Data["UserName"] = username
	c.TplName = "admin/admin_modpwd.tpl"
}

//执行修改密码
func (c *AdminController) DoModPwd() {
	CheckAdminLogin(&c.Controller, 1)
	username, _ := c.GetSession("adminname").(string)
	uname := c.GetString("UserName")
	oldpwd := c.GetString("oldPwd")
	newpwd := c.GetString("newPwd")
	newpwd2 := c.GetString("newPwd2")
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	if len(uname) < 4 {
		tip.Message = "用户名不能小于5个字符！"
		EchoTip(&c.Controller, tip)
	}
	if len(oldpwd) < 5 {
		tip.Message = "旧密码不能小于5个字符"
		EchoTip(&c.Controller, tip)
	}
	if len(newpwd) < 5 {
		tip.Message = "新密码不能小于5个字符！"
		EchoTip(&c.Controller, tip)
	}
	if newpwd != newpwd2 {
		tip.Message = "两次输入密码不一致，请重新输入！"
		EchoTip(&c.Controller, tip)
	}
	if username == uname && oldpwd == newpwd {
		tip.Message = "新旧密码一致，无法修改！"
		EchoTip(&c.Controller, tip)
	}

	//获取管理员

	admin := models.GetAdmin(username)
	if admin != nil {
		newpwdMd5 := common.GetMd5String(oldpwd)
		if admin.PassWord != newpwdMd5 {
			tip.Message = "旧密码错误，请重新输入！"
			EchoTip(&c.Controller, tip)
		}
		admin.UserName = uname
		admin.PassWord = common.GetMd5String(newpwd)
		models.EditAdmin(admin)
		tip.Status = models.TipSuccess
		tip.Message = "修改密码成功"
		if username != uname {
			tip.ReturnUrl = "/admin/logout"
		} else {
			tip.ReturnUrl = "/admin/"
		}
		EchoTip(&c.Controller, tip)

	} else {
		tip.Message = "登录授权过期！"
		EchoTip(&c.Controller, tip)
	}
}

//管理员列表
func (c *AdminController) List() {
	CheckAdminLogin(&c.Controller, 0)
	//显示列表
	page, _ := c.GetInt("p")
	size, _ := c.GetInt("s")
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}
	//计算开始
	start := (page - 1) * size
	where := "1=1 "
	key := c.GetString("key")
	if key != "" {
		where += " And UserName like '%" + key + "%' "
	}
	list, total := models.GetAdminList(where, "Id desc", size, start)
	//计算页数
	totalpage := total / int64(page)
	if total%int64(page) > 0 {
		totalpage++
	}
	p := pagination.NewPaginator(c.Ctx.Request, size, total)
	c.Data["Title"] = "管理员列表"
	c.Data["List"] = list
	c.Data["Page"] = page
	c.Data["Total"] = total
	c.Data["paginator"] = p
	c.TplName = "admin/admin_list.tpl"
}

//显示添加管理员
func (c *AdminController) Add() {
	CheckAdminLogin(&c.Controller, 0)
	//添加的时候，先初始化一个实体
	entity := &models.Admin{}

	c.Data["Title"] = "添加管理员"
	c.Data["Action"] = "add"
	c.Data["Entity"] = entity
	c.TplName = "admin/admin_add.tpl"
}

//执行添加管理员
func (c *AdminController) DoAdd() {
	CheckAdminLogin(&c.Controller, 1)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	entity := &models.Admin{}
	username := c.GetString("UserName")
	password := c.GetString("PassWord")
	nickname := c.GetString("NickName")
	if len(username) < 4 {
		tip.Message = "用户名不能小于5个字符！"
		EchoTip(&c.Controller, tip)
	}
	if len(password) < 5 {
		tip.Message = "密码不能小于5个字符！"
		EchoTip(&c.Controller, tip)
	}
	a := models.GetAdmin(username)
	if a != nil {
		tip.Message = "该用户名已经存在，请修改！"
		EchoTip(&c.Controller, tip)
	}
	entity.UserName = username
	entity.PassWord = common.GetMd5String(password)
	entity.NickName = nickname
	entity.RoleId = 1 //否则无法登录
	if id, err := models.AddAdmin(entity); id > 0 && err == nil {
		//添加成功
		tip.Status = models.TipSuccess
		tip.Id = id
		tip.ReturnUrl = "/admin/admin/list"
		tip.Message = "添加管理员成功"
	} else {
		tip.Id = id
		tip.Message = "添加管理员失败：" + err.Error()
	}
	EchoTip(&c.Controller, tip)
}

//显示修改管理员
func (c *AdminController) Edit() {
	CheckAdminLogin(&c.Controller, 0)

	myid := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(myid, 10, 64)
	if err != nil {
		EchoErrorPage(&c.Controller, "错误参数传递！", "/admin/admin/list")
	}
	entity := models.GetAdminById(id)
	if entity == nil {
		EchoErrorPage(&c.Controller, "系统找不到本记录！", "/admin/category")
	}
	c.Data["Title"] = "编辑/查看管理员详情"
	c.Data["Action"] = "edit"
	c.Data["Entity"] = entity
	c.TplName = "admin/admin_add.tpl"
}

//执行修改管理员
func (c *AdminController) DoEdit() {
	CheckAdminLogin(&c.Controller, 1)
	myid := c.Ctx.Input.Param(":id")
	editid, err := strconv.ParseInt(myid, 10, 64)
	tip := &models.TipJSON{}
	tip.Status = models.TipError

	if err != nil {
		tip.Message = "错误提交！"
		EchoTip(&c.Controller, tip)

	}

	id, err := c.GetInt64("Id")
	if err != nil {
		tip.Message = "错误参数传递！"
		EchoTip(&c.Controller, tip)

	}
	if editid != id {
		tip.Message = "错误参数传递！"
		EchoTip(&c.Controller, tip)
	}
	entity := models.GetAdminById(id)
	if entity == nil {
		tip.Message = "系统找不到本记录！"
		EchoTip(&c.Controller, tip)
	}
	username := c.GetString("UserName")
	password := c.GetString("PassWord")
	nickname := c.GetString("NickName")
	if len(username) < 4 {
		tip.Message = "用户名不能小于5个字符！"
		EchoTip(&c.Controller, tip)
	}
	if password != "" && len(password) < 5 {
		tip.Message = "密码不能小于5个字符！"
		EchoTip(&c.Controller, tip)
	}
	//如果用户名修改了。需要判断
	if entity.UserName != username {
		checkadmin := models.GetAdmin(username)
		if checkadmin != nil {
			tip.Message = "修改的用户名已经存在，请修改其他的用户名！"
			EchoTip(&c.Controller, tip)
		} else {
			entity.UserName = username
		}
	}
	if password != "" {
		entity.PassWord = common.GetMd5String(password)
	}
	entity.NickName = nickname
	if id, err := models.EditAdmin(entity); id > 0 && err == nil {
		//添加成功
		tip.Status = models.TipSuccess
		tip.Id = id
		tip.ReturnUrl = "/admin/admin/list"
		tip.Message = "编辑管理员信息成功"
	} else {
		tip.Id = id
		tip.Message = "编辑管理员失败：" + err.Error()
	}
	EchoTip(&c.Controller, tip)
}

//执行删除管理员
func (c *AdminController) DoDel() {
	id, _ := strconv.ParseInt(c.GetString("id"), 10, 64)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	if id <= 0 {
		tip.Message = "错误参数传递！"
		EchoTip(&c.Controller, tip)
	}
	myid, _ := c.GetSession("adminid").(int)
	if id == int64(myid) {
		tip.Message = "您不能删除自己！"
		EchoTip(&c.Controller, tip)
	}
	entity := models.GetAdminById(id)
	if entity == nil {
		tip.Message = "系统找不到本记录！"
		EchoTip(&c.Controller, tip)
	} else {
		if err := models.DelAdmin(entity); err != nil {
			tip.Message = "删除出错：" + err.Error()
			EchoTip(&c.Controller, tip)
		} else {
			tip.Status = models.TipSuccess
			tip.Message = "删除成功！"
			tip.ReturnUrl = "/admin/admin/list"
			EchoTip(&c.Controller, tip)
		}
	}
}

//判断是否登陆
func CheckAdminLogin(c *beego.Controller, t int64) {
	id, _ := c.GetSession("adminid").(int)
	username, _ := c.GetSession("adminname").(string)
	adminrole, _ := c.GetSession("adminrole").(int)
	if id == 0 || username == "" || adminrole == 0 {
		switch t {
		case 0:
			c.Redirect("/admin/login", 301)
		case 1:
			//返回JSON
			json := &models.TipJSON{}
			json.Status = models.TipError
			json.Message = "请先登录，再执行此操作"
			json.ReturnUrl = "/admin/login"

			c.Data["json"] = json
			c.ServeJSON()
			c.StopRun()
		}

	}
	c.Data["AdminName"] = username
	c.Data["AdminId"] = id
	c.Layout = "admin/layout.tpl"
}

//提示错误页面
func EchoErrorPage(c *beego.Controller, message string, url string) {
	isGoback := false
	if url == "" {
		isGoback = true
	}
	c.Data["Message"] = message
	c.Data["Url"] = url
	c.Data["IsGoback"] = isGoback
	c.Layout = "admin/error.tpl"
	c.StopRun()
}

//信息提示，如果异步，返回json、如果同步，则页面提示并返回
func EchoTip(c *beego.Controller, json *models.TipJSON) {
	if c.IsAjax() {
		//异步提交
		c.Data["json"] = json
		c.ServeJSON()
		c.StopRun()
	} else {
		tpl := "admin/error.tpl"
		if json.Status == models.TipSuccess {
			tpl = "admin/success.tpl"
		}
		c.Data["Tip"] = json
		c.Layout = tpl
		//此处不能用 c.StopRun() 返回，否则会空白页面！
		//c.StopRun()
		return
	}
}
