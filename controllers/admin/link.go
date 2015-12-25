package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/hogenwang/comcms/models"
	"strconv"
	//"time"
)

type LinkController struct {
	beego.Controller
}

//显示列表
func (c *LinkController) List() {
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
		where += " And Title like '%" + key + "%' "
	}
	CheckAdminLogin(&c.Controller, 0)
	list, total := models.GetLinkList(where, "Id desc", size, start)
	//计算页数
	totalpage := total / int64(page)
	if total%int64(page) > 0 {
		totalpage++
	}
	p := pagination.NewPaginator(c.Ctx.Request, size, total)
	c.Data["Title"] = "友情链接列表"
	c.Data["List"] = list
	c.Data["Page"] = page
	c.Data["Total"] = total
	c.Data["paginator"] = p
	c.TplNames = "admin/link.tpl"
}

//添加友情链接页面
func (c *LinkController) Add() {
	CheckAdminLogin(&c.Controller, 0)
	//添加的时候，先初始化一个实体
	entity := &models.Link{}
	entity.Id = 0
	entity.Rank = 999

	c.Data["Title"] = "添加链接"
	c.Data["Action"] = "add"
	c.Data["Entity"] = entity

	c.TplNames = "admin/link_add.tpl"
}

//编辑友情链接页面
func (c *LinkController) Edit() {
	CheckAdminLogin(&c.Controller, 0)

	myid := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(myid, 10, 64)
	if err != nil {
		EchoErrorPage(&c.Controller, "错误参数传递！", "/admin/link")
	}
	entity := models.GetLink(id)
	if entity == nil {
		EchoErrorPage(&c.Controller, "系统找不到本记录！", "/admin/link")
	}
	c.Data["Title"] = "修改友情链接"
	c.Data["Action"] = "edit"
	c.Data["Entity"] = entity
	c.TplNames = "admin/link_add.tpl"
}

//执行添加友情链接
func (c *LinkController) DoAdd() {
	CheckAdminLogin(&c.Controller, 1)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	entity := &models.Link{}
	entity.Title = c.GetString("Title")
	if entity.Title == "" {
		tip.Message = "站点名称不能为空！"
		EchoTip(&c.Controller, tip)
	}

	entity.Logo = c.GetString("Logo")
	entity.Description = c.GetString("Description")

	if c.GetString("IsHide") == "1" {
		entity.IsHide = 1
	} else {
		entity.IsHide = 0
	}
	entity.Rank, _ = strconv.ParseInt(c.GetString("Rank"), 10, 64)
	entity.Url = c.GetString("Url")

	if id, err := models.AddLink(entity); id > 0 && err == nil {
		//添加成功
		tip.Status = models.TipSuccess
		tip.Id = id
		tip.ReturnUrl = "/admin/link"
		tip.Message = "添加新友情链接成功"
	} else {
		tip.Id = id
		tip.Message = "添加新友情链接失败：" + err.Error()
	}
	EchoTip(&c.Controller, tip)
}

//执行修改友情链接
func (c *LinkController) DoEdit() {
	CheckAdminLogin(&c.Controller, 1)
	tip := &models.TipJSON{}
	tip.Status = models.TipError

	id, err := c.GetInt64("Id")
	if err != nil {
		tip.Message = "错误参数传递！"
		EchoTip(&c.Controller, tip)

	}
	entity := models.GetLink(id)
	if entity == nil {
		tip.Message = "系统找不到本记录！"
		EchoTip(&c.Controller, tip)
	}
	entity.Title = c.GetString("Title")
	if entity.Title == "" {
		tip.Message = "站点名称不能为空！"
		EchoTip(&c.Controller, tip)
	}
	entity.Logo = c.GetString("Logo")
	entity.Description = c.GetString("Description")

	if c.GetString("IsHide") == "1" {
		entity.IsHide = 1
	} else {
		entity.IsHide = 0
	}
	entity.Rank, _ = strconv.ParseInt(c.GetString("Rank"), 10, 64)
	entity.Url = c.GetString("Url")

	if id, err := models.EditLink(entity); id > 0 && err == nil {
		//添加成功
		tip.Status = models.TipSuccess
		tip.Id = id
		tip.ReturnUrl = "/admin/link"
		tip.Message = "友情链接修改成功"
	} else {
		tip.Id = id
		tip.Message = "友情链接修改失败：" + err.Error()
	}
	EchoTip(&c.Controller, tip)

}

//删除友情链接
func (c *LinkController) DoDel() {
	CheckAdminLogin(&c.Controller, 1)
	id, _ := strconv.ParseInt(c.GetString("id"), 10, 64)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	if id <= 0 {
		tip.Message = "错误参数传递！"
		EchoTip(&c.Controller, tip)
	}
	entity := models.GetLink(id)
	if entity == nil {
		tip.Message = "系统找不到本记录！"
		EchoTip(&c.Controller, tip)
	} else {
		if err := models.DelLink(entity); err != nil {
			tip.Message = "删除出错：" + err.Error()
			EchoTip(&c.Controller, tip)
		} else {
			tip.Status = models.TipSuccess
			tip.Message = "删除成功！"
			tip.ReturnUrl = "/admin/link"
			EchoTip(&c.Controller, tip)
		}
	}
}
