package admin

import (
	"github.com/astaxie/beego"
	//"github.com/hogenwang/comcms2/common"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/hogenwang/comcms/models"
	"strconv"
	"time"
)

type ArticleController struct {
	beego.Controller
}

//显示列表
func (c *ArticleController) List() {
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
	list, total := models.GetArticleList(where, "Id desc", size, start)
	//计算页数
	totalpage := total / int64(page)
	if total%int64(page) > 0 {
		totalpage++
	}
	p := pagination.NewPaginator(c.Ctx.Request, size, total)
	c.Data["Title"] = "文章列表"
	c.Data["List"] = list
	c.Data["Key"] = key
	c.Data["Page"] = page
	c.Data["Total"] = total
	c.Data["paginator"] = p
	c.TplName = "admin/article.tpl"
}

//添加文章页面
func (c *ArticleController) Add() {
	CheckAdminLogin(&c.Controller, 0)
	//添加的时候，先初始化一个实体
	entity := &models.Article{}
	entity.Id = 0
	entity.Created = time.Now()
	entity.Rank = 999
	//获取所有栏目
	Categories := models.GetCategoryTree(0, -1, true, 0)
	c.Data["Title"] = "添加文章"
	c.Data["Action"] = "add"
	c.Data["Entity"] = entity
	c.Data["Categories"] = Categories
	c.TplName = "admin/article_add.tpl"
}

//编辑文章页面
func (c *ArticleController) Edit() {
	CheckAdminLogin(&c.Controller, 0)

	myid := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(myid, 10, 64)
	if err != nil {
		EchoErrorPage(&c.Controller, "错误参数传递！", "/admin/category")
	}
	entity := models.GetArticle(id)
	if entity == nil {
		EchoErrorPage(&c.Controller, "系统找不到本记录！", "/admin/category")
	}
	//获取所有栏目
	Categories := models.GetCategoryTree(0, -1, true, 0)
	c.Data["Title"] = "修改文章"
	c.Data["Action"] = "edit"
	c.Data["Entity"] = entity
	c.Data["Categories"] = Categories
	c.TplName = "admin/article_add.tpl"
}

//执行添加文章
func (c *ArticleController) DoAdd() {
	CheckAdminLogin(&c.Controller, 1)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	entity := &models.Article{}
	entity.Title = c.GetString("Title")
	if entity.Title == "" {
		tip.Message = "标题不能为空！"
		EchoTip(&c.Controller, tip)
	}
	kid, _ := strconv.ParseInt(c.GetString("Kid"), 10, 64)
	if kid == 0 {
		tip.Message = "请选择一个栏目！"
		EchoTip(&c.Controller, tip)
	}
	entity.Origin = c.GetString("Origin")
	entity.OriginUrl = c.GetString("OriginUrl")
	if c.GetString("IsNew") == "1" {
		entity.IsNew = 1
	} else {
		entity.IsNew = 0
	}
	if c.GetString("IsRecommend") == "1" {
		entity.IsRecommend = 1
	} else {
		entity.IsRecommend = 0
	}
	if c.GetString("IsHide") == "1" {
		entity.IsHide = 1
	} else {
		entity.IsHide = 0
	}
	views, _ := strconv.ParseInt(c.GetString("Views"), 10, 64)
	addtime, _ := time.Parse("2006-01-02 15:04:05", c.GetString("Created"))
	entity.Created = addtime
	entity.Views = views
	entity.Rank, _ = strconv.ParseInt(c.GetString("Rank"), 10, 64)
	entity.LinkUrl = c.GetString("LinkUrl")
	entity.Keyword = c.GetString("Keyword")
	entity.Description = c.GetString("Description")
	entity.Pic = c.GetString("Pic")
	entity.Content = c.GetString("Content")

	if id, err := models.AddArticle(entity); id > 0 && err == nil {
		//添加成功
		tip.Status = models.TipSuccess
		tip.Id = id
		tip.ReturnUrl = "/admin/article"
		tip.Message = "添加新文章成功"
	} else {
		tip.Id = id
		tip.Message = "添加新文章失败：" + err.Error()
	}
	EchoTip(&c.Controller, tip)

}

//执行修改文章
func (c *ArticleController) DoEdit() {
	CheckAdminLogin(&c.Controller, 1)
	tip := &models.TipJSON{}
	tip.Status = models.TipError

	id, err := c.GetInt64("Id")
	if err != nil {
		tip.Message = "错误参数传递！"
		EchoTip(&c.Controller, tip)
	}
	entity := models.GetArticle(id)
	if entity == nil {
		tip.Message = "系统找不到本记录！"
		EchoTip(&c.Controller, tip)
	}

	entity.Title = c.GetString("Title")
	if entity.Title == "" {
		tip.Message = "标题不能为空！"
		EchoTip(&c.Controller, tip)
	}
	kid, _ := strconv.ParseInt(c.GetString("Kid"), 10, 64)
	if kid == 0 {
		tip.Message = "请选择一个栏目！"
		EchoTip(&c.Controller, tip)
	}
	entity.Origin = c.GetString("Origin")
	entity.OriginUrl = c.GetString("OriginUrl")
	if c.GetString("IsNew") == "1" {
		entity.IsNew = 1
	} else {
		entity.IsNew = 0
	}
	if c.GetString("IsRecommend") == "1" {
		entity.IsRecommend = 1
	} else {
		entity.IsRecommend = 0
	}
	if c.GetString("IsHide") == "1" {
		entity.IsHide = 1
	} else {
		entity.IsHide = 0
	}
	views, _ := strconv.ParseInt(c.GetString("Views"), 10, 64)
	addtime, _ := time.Parse("2006-01-02 15:04:05", c.GetString("Created"))
	entity.Created = addtime
	entity.Views = views
	entity.Rank, _ = strconv.ParseInt(c.GetString("Rank"), 10, 64)
	entity.LinkUrl = c.GetString("LinkUrl")
	entity.Keyword = c.GetString("Keyword")
	entity.Description = c.GetString("Description")
	entity.Pic = c.GetString("Pic")
	entity.Content = c.GetString("Content")

	if id, err := models.EditArticle(entity); id > 0 && err == nil {
		//修改成功
		tip.Status = models.TipSuccess
		tip.Id = id
		tip.ReturnUrl = "/admin/article"
		tip.Message = "文章修改成功"
	} else {
		tip.Id = id
		tip.Message = "文章修改失败：" + err.Error()
	}
	EchoTip(&c.Controller, tip)

}

//删除文章
func (c *ArticleController) DoDel() {
	CheckAdminLogin(&c.Controller, 1)
	id, _ := strconv.ParseInt(c.GetString("id"), 10, 64)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	if id <= 0 {
		tip.Message = "错误参数传递！"
		EchoTip(&c.Controller, tip)
	}
	entity := models.GetArticle(id)
	if entity == nil {
		tip.Message = "系统找不到本记录！"
		EchoTip(&c.Controller, tip)
	} else {
		if err := models.DelArticle(entity); err != nil {
			tip.Message = "删除出错：" + err.Error()
			EchoTip(&c.Controller, tip)
		} else {
			tip.Status = models.TipSuccess
			tip.Message = "删除成功！"
			tip.ReturnUrl = "/admin/article"
			EchoTip(&c.Controller, tip)
		}
	}
}
