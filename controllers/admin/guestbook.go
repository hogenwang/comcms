package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/hogenwang/comcms/models"
	"strconv"
	//"time"
)

type GuestbookController struct {
	beego.Controller
}

//显示列表
func (c *GuestbookController) List() {
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
	list, total := models.GetGuestbookList(where, "Id desc", size, start)
	//计算页数
	totalpage := total / int64(page)
	if total%int64(page) > 0 {
		totalpage++
	}
	p := pagination.NewPaginator(c.Ctx.Request, size, total)
	c.Data["Title"] = "留言板列表"
	c.Data["List"] = list
	c.Data["Page"] = page
	c.Data["Total"] = total
	c.Data["paginator"] = p
	c.TplNames = "admin/guestbook.tpl"
}

//查看留言板页面
func (c *GuestbookController) View() {
	CheckAdminLogin(&c.Controller, 0)

	myid := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(myid, 10, 64)
	if err != nil {
		EchoErrorPage(&c.Controller, "错误参数传递！", "/admin/guestbook")
	}
	entity := models.GetGuestbook(id)
	if entity == nil {
		EchoErrorPage(&c.Controller, "系统找不到本记录！", "/admin/guestbook")
	}
	c.Data["Title"] = "查看留言板详情"
	c.Data["Action"] = "view"
	c.Data["Entity"] = entity
	c.TplNames = "admin/guestbook_view.tpl"
}

//删除留言板
func (c *GuestbookController) DoDel() {
	CheckAdminLogin(&c.Controller, 1)
	id, _ := strconv.ParseInt(c.GetString("id"), 10, 64)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	if id <= 0 {
		tip.Message = "错误参数传递！"
		EchoTip(&c.Controller, tip)
	}
	entity := models.GetGuestbook(id)
	if entity == nil {
		tip.Message = "系统找不到本记录！"
		EchoTip(&c.Controller, tip)
	} else {
		if err := models.DelGuestbook(entity); err != nil {
			tip.Message = "删除出错：" + err.Error()
			EchoTip(&c.Controller, tip)
		} else {
			tip.Status = models.TipSuccess
			tip.Message = "删除成功！"
			tip.ReturnUrl = "/admin/guestbook"
			EchoTip(&c.Controller, tip)
		}
	}
}
