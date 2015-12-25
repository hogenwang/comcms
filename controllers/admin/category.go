package admin

//分类处理
import (
	"github.com/astaxie/beego"
	"github.com/hogenwang/comcms/common"
	"github.com/hogenwang/comcms/models"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

//显示列表
func (c *CategoryController) List() {
	//显示栏目列表
	CheckAdminLogin(&c.Controller, 0)
	categories := models.GetCategoryTree(0, -1, false, 0)
	c.Data["List"] = categories
	c.Data["Title"] = "栏目列表"
	c.TplNames = "admin/category.tpl"
}

//显示添加文章栏目
func (c *CategoryController) Add() {
	CheckAdminLogin(&c.Controller, 0)
	//添加的时候，先初始化一个实体
	entity := &models.Category{}
	entity.Id = 0
	entity.PageSize = 15
	entity.Rank = 0
	entity.DetailTemplateFile = "article.tpl"
	entity.TemplateFile = "list_article.tpl"
	//获取所有上级栏目
	PCategories := models.GetCategoryTree(0, -1, true, 0)
	//获取模板目录
	skin := "default"
	templatesList, err := common.GetFolderFiles("views/"+skin, "")
	if err != nil {
		templatesList = nil
	}
	c.Data["Title"] = "添加栏目"
	c.Data["Action"] = "add"
	c.Data["Entity"] = entity
	c.Data["TemplatesList"] = templatesList
	c.Data["PCategories"] = PCategories
	c.TplNames = "admin/category_add.tpl"
}

//修改文章栏目页面
func (c *CategoryController) Edit() {
	CheckAdminLogin(&c.Controller, 0)
	myid := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(myid, 10, 64)
	if err != nil {
		EchoErrorPage(&c.Controller, "错误参数传递！", "/admin/category")
	}
	entity := models.GetCategory(id)
	if entity == nil {
		EchoErrorPage(&c.Controller, "系统找不到本记录！", "/admin/category")
	}
	//获取所有上级栏目
	PCategories := models.GetCategoryTree(0, -1, true, entity.Id)
	//获取模板目录
	skin := "default"
	templatesList, err := common.GetFolderFiles("views/"+skin, "")
	if err != nil {
		templatesList = nil
	}
	c.Data["Title"] = "修改栏目"
	c.Data["Entity"] = entity
	c.Data["Action"] = "edit"
	c.Data["TemplatesList"] = templatesList
	c.Data["PCategories"] = PCategories
	c.TplNames = "admin/category_add.tpl"
}

//执行添加
func (c *CategoryController) DoAdd() {
	CheckAdminLogin(&c.Controller, 1)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	entity := &models.Category{}
	entity.Title = c.GetString("Title")
	entity.Ctype = 0
	entity.Content = c.GetString("Content")
	newPid, _ := strconv.ParseInt(c.GetString("Pid"), 10, 64)
	//如果PID大于0 则获取分类，本身级别增加1
	if newPid != entity.Pid {
		if newPid > 0 {
			pc := models.GetCategory(newPid)
			if pc != nil {
				entity.Level = pc.Level + 1
			} else {
				entity.Pid = 0
				entity.Level = 1
			}
		} else {
			entity.Pid = 0
			entity.Level = 1
		}
	}

	if entity.Title == "" {
		tip.Message = "分类名称不能为空"
		c.Data["json"] = tip
		c.ServeJson()
		c.StopRun()
	}

	if c.GetString("IsList") == "1" {
		entity.IsList = 1
	} else {
		entity.IsList = 0
	}
	if c.GetString("IsHide") == "1" {
		entity.IsHide = 1
	} else {
		entity.IsHide = 0
	}
	entity.PageSize, _ = strconv.ParseInt(c.GetString("PageSize"), 10, 64)
	entity.Rank, _ = strconv.ParseInt(c.GetString("Rank"), 10, 64)
	entity.Keyword = c.GetString("Keyword")
	entity.Description = c.GetString("Description")
	entity.Pic = c.GetString("Pic")
	entity.DetailTemplateFile = c.GetString("DetailTemplateFile")
	entity.TemplateFile = c.GetString("TemplateFile")
	entity.LinkUrl = c.GetString("LinkUrl")
	if id, err := models.AddCategory(entity); id > 0 && err == nil {
		//添加成功
		tip.Status = models.TipSuccess
		tip.Id = id
		tip.ReturnUrl = "/admin/category"
		tip.Message = "添加新分类成功"
	} else {
		tip.Id = id
		tip.Message = "添加新分类失败：" + err.Error()
	}
	c.Data["json"] = tip
	c.ServeJson()
	c.StopRun()

}

//执行修改
func (c *CategoryController) DoEdit() {
	CheckAdminLogin(&c.Controller, 1)
	id, err := c.GetInt64("Id")
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	if err != nil {
		tip.Message = "错误参数传递！"
		c.Data["json"] = tip
		c.ServeJson()
		c.StopRun()
	}
	entity := models.GetCategory(id)
	if entity == nil {
		tip.Message = "系统找不到本记录！"
		c.Data["json"] = tip
		c.ServeJson()
		c.StopRun()
	}
	entity.Title = c.GetString("Title")
	entity.Ctype = 0
	entity.Content = c.GetString("Content")
	entity.Pid, _ = strconv.ParseInt(c.GetString("Pid"), 10, 64)
	//如果PID大于0 则获取分类，本身级别增加1
	if entity.Pid > 0 {
		pc := models.GetCategory(entity.Pid)
		if pc != nil {
			entity.Level = pc.Level + 1
		} else {
			entity.Pid = 0
			entity.Level = 1
		}
	} else {
		entity.Level = 1
	}

	if entity.Title == "" {
		tip.Message = "分类名称不能为空"
		c.Data["json"] = tip
		c.ServeJson()
		c.StopRun()
	}

	if c.GetString("IsList") == "1" {
		entity.IsList = 1
	} else {
		entity.IsList = 0
	}
	if c.GetString("IsHide") == "1" {
		entity.IsHide = 1
	} else {
		entity.IsHide = 0
	}
	entity.PageSize, _ = c.GetInt64("PageSize") // strconv.ParseInt(c.GetString("PageSize"), 10, 64)
	entity.Rank, _ = strconv.ParseInt(c.GetString("Rank"), 10, 64)
	entity.Keyword = c.GetString("Keyword")
	entity.Description = c.GetString("Description")
	entity.Pic = c.GetString("Pic")
	entity.DetailTemplateFile = c.GetString("DetailTemplateFile")
	entity.TemplateFile = c.GetString("TemplateFile")
	entity.LinkUrl = c.GetString("LinkUrl")
	if id, err := models.EditCategory(entity); id > 0 && err == nil {
		//添加成功
		tip.Status = models.TipSuccess
		tip.Id = id
		tip.ReturnUrl = "/admin/category"
		tip.Message = "编辑分类成功"
	} else {
		tip.Id = id
		tip.Message = "编辑分类失败：" + err.Error()
	}
	c.Data["json"] = tip
	c.ServeJson()
	c.StopRun()
}

//删除文章
func (c *CategoryController) DoDel() {
	CheckAdminLogin(&c.Controller, 1)
	id, _ := strconv.ParseInt(c.GetString("id"), 10, 64)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	if id <= 0 {
		tip.Message = "错误参数传递！"
		EchoTip(&c.Controller, tip)
	}
	entity := models.GetCategory(id)
	if entity == nil {
		tip.Message = "系统找不到本记录！"
		EchoTip(&c.Controller, tip)
	} else {
		if err := models.DelCategory(entity); err != nil {
			tip.Message = "删除出错：" + err.Error()
			EchoTip(&c.Controller, tip)
		} else {
			tip.Status = models.TipSuccess
			tip.Message = "删除成功！"
			tip.ReturnUrl = "/admin/category"
			EchoTip(&c.Controller, tip)
		}
	}
}
