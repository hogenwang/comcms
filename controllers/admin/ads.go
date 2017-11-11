package admin

import (
	"github.com/astaxie/beego"
	//"github.com/hogenwang/comcms2/common"
	"encoding/json"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/hogenwang/comcms/models"
	"strconv"
	//"time"
)

type AdsController struct {
	beego.Controller
}

//显示列表
func (c *AdsController) List() {
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
	list, total := models.GetAdsList(where, "Id desc", size, start)
	//计算页数
	totalpage := total / int64(page)
	if total%int64(page) > 0 {
		totalpage++
	}
	p := pagination.NewPaginator(c.Ctx.Request, size, total)
	c.Data["Title"] = "广告列表"
	c.Data["List"] = list
	c.Data["Page"] = page
	c.Data["Total"] = total
	c.Data["paginator"] = p
	c.TplName = "admin/ads.tpl"
}

//添加广告页面
func (c *AdsController) Add() {
	CheckAdminLogin(&c.Controller, 0)
	//添加的时候，先初始化一个实体
	entity := &models.Ads{}
	entity.Id = 0
	entity.Rank = 999

	c.Data["Title"] = "添加广告"
	c.Data["Action"] = "add"
	c.Data["Entity"] = entity

	c.TplName = "admin/ads_add.tpl"
}

//编辑广告页面
func (c *AdsController) Edit() {
	CheckAdminLogin(&c.Controller, 0)

	myid := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(myid, 10, 64)
	if err != nil {
		EchoErrorPage(&c.Controller, "错误参数传递！", "/admin/ads")
	}
	entity := models.GetAds(id)
	if entity == nil {
		EchoErrorPage(&c.Controller, "系统找不到本记录！", "/admin/ads")
	}
	c.Data["Title"] = "修改广告"
	c.Data["Action"] = "edit"
	c.Data["Entity"] = entity
	c.TplName = "admin/ads_add.tpl"
}

//执行添加广告
func (c *AdsController) DoAdd() {
	CheckAdminLogin(&c.Controller, 1)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	entity := &models.Ads{}
	entity.Title = c.GetString("Title")
	if entity.Title == "" {
		tip.Message = "标题不能为空！"
		EchoTip(&c.Controller, tip)
	}
	entity.Rank, _ = strconv.ParseInt(c.GetString("Rank"), 10, 64)
	entity.Description = c.GetString("Description")
	if c.GetString("IsHide") == "1" {
		entity.IsHide = 1
	} else {
		entity.IsHide = 0
	}
	entity.Tid, _ = strconv.ParseInt(c.GetString("Tid"), 10, 64)
	content := GetAdsDetail(&c.Controller)
	entity.Content = content

	if id, err := models.AddAds(entity); id > 0 && err == nil {
		//添加成功
		tip.Status = models.TipSuccess
		tip.Id = id
		tip.ReturnUrl = "/admin/ads"
		tip.Message = "添加新广告成功"
	} else {
		tip.Id = id
		tip.Message = "添加新广告失败：" + err.Error()
	}
	EchoTip(&c.Controller, tip)
}

//执行修改广告
func (c *AdsController) DoEdit() {
	CheckAdminLogin(&c.Controller, 1)
	tip := &models.TipJSON{}
	tip.Status = models.TipError

	id, err := c.GetInt64("Id")
	if err != nil {
		tip.Message = "错误参数传递！"
		EchoTip(&c.Controller, tip)
	}
	entity := models.GetAds(id)
	if entity == nil {
		tip.Message = "系统找不到本记录！"
		EchoTip(&c.Controller, tip)
	}
	entity.Title = c.GetString("Title")
	if entity.Title == "" {
		tip.Message = "标题不能为空！"
		EchoTip(&c.Controller, tip)
	}
	entity.Rank, _ = strconv.ParseInt(c.GetString("Rank"), 10, 64)
	entity.Description = c.GetString("Description")
	if c.GetString("IsHide") == "1" {
		entity.IsHide = 1
	} else {
		entity.IsHide = 0
	}
	entity.Tid, _ = strconv.ParseInt(c.GetString("Tid"), 10, 64)
	content := GetAdsDetail(&c.Controller)
	entity.Content = content

	if id, err := models.EditAds(entity); id > 0 && err == nil {
		//修改成功
		tip.Status = models.TipSuccess
		tip.Id = id
		tip.ReturnUrl = "/admin/ads"
		tip.Message = "广告修改成功"
	} else {
		tip.Id = id
		tip.Message = "修改广告失败：" + err.Error()
	}
	EchoTip(&c.Controller, tip)
}

//删除广告
func (c *AdsController) DoDel() {
	CheckAdminLogin(&c.Controller, 1)
	id, _ := strconv.ParseInt(c.GetString("id"), 10, 64)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	if id <= 0 {
		tip.Message = "错误参数传递！"
		EchoTip(&c.Controller, tip)
	}
	entity := models.GetAds(id)
	if entity == nil {
		tip.Message = "系统找不到本记录！"
		EchoTip(&c.Controller, tip)
	} else {
		if err := models.DelAds(entity); err != nil {
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

//获取广告内容
func GetAdsDetail(c *beego.Controller) string {
	str := ""
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	tid, _ := strconv.ParseInt(c.GetString("Tid"), 10, 64)
	switch tid {
	case 0: //代码
		script := &models.ScriptAds{}
		script.Content = c.GetString("txtScript")
		if script.Content == "" {
			tip.Message = "代码不能为空！"
			EchoTip(c, tip)
		}
		arrstr, _ := json.Marshal(script)
		str = string(arrstr)
	case 1: //文字
		text := &models.TextAds{}
		text.Txt = c.GetString("txt_Txt")
		text.Link = c.GetString("txt_Txt")
		text.Style = c.GetString("txt_Style")
		if text.Txt == "" {
			tip.Message = "文字内容不能为空！"
			EchoTip(c, tip)
		}
		if text.Link == "" {
			tip.Message = "文字链接不能为空！"
			EchoTip(c, tip)
		}
		arrstr, _ := json.Marshal(text)
		str = string(arrstr)
	case 2: //图片类
		img := &models.ImgAds{}
		img.Img = c.GetString("img_Img")
		img.Alt = c.GetString("img_Alt")
		img.Link = c.GetString("img_Link")
		img.Height, _ = strconv.ParseInt(c.GetString("img_Height"), 10, 64)
		img.Width, _ = strconv.ParseInt(c.GetString("img_Width"), 10, 64)
		if img.Img == "" {
			tip.Message = "图片地址不能为空！"
			EchoTip(c, tip)
		}
		if img.Link == "" {
			tip.Message = "图片链接不能为空！"
			EchoTip(c, tip)
		}
		arrstr, _ := json.Marshal(img)
		str = string(arrstr)
	case 3: //Flash
		flash := &models.FlashAds{}
		flash.Swf = c.GetString("flash_Swf")
		flash.Height, _ = strconv.ParseInt(c.GetString("flash_Height"), 10, 64)
		flash.Width, _ = strconv.ParseInt(c.GetString("flash_Width"), 10, 64)
		if flash.Swf == "" {
			tip.Message = "Flash 地址不能为空！"
			EchoTip(c, tip)
		}
		arrstr, _ := json.Marshal(flash)
		str = string(arrstr)
	case 4: //幻灯片
		sw, _ := strconv.ParseInt(c.GetString("slide_Width"), 10, 64)
		sh, _ := strconv.ParseInt(c.GetString("slide_Height"), 10, 64)
		SImg := c.GetStrings("slide_Img")
		SLink := c.GetStrings("slide_Link")
		SAlt := c.GetStrings("slide_Alt")
		if len(SImg) <= 0 {
			tip.Message = "幻灯片图片不能为空！"
			EchoTip(c, tip)
		}
		if len(SLink) <= 0 {
			tip.Message = "幻灯片图片链接不能为空！"
			EchoTip(c, tip)
		}
		length := len(SImg)
		listImg := []*models.ImgAds{}
		for i := 0; i < length; i++ {
			tmp := models.ImgAds{}
			if SImg[i] != "" && SLink[i] != "" {
				tmp.Img = SImg[i]
				tmp.Link = SLink[i]
				tmp.Alt = SAlt[i]
				tmp.Height = sh
				tmp.Width = sw
				listImg = append(listImg, &tmp)
			}
		}
		if listImg == nil || len(listImg) <= 0 {
			tip.Message = "幻灯片请最少设置一个图片！"
			EchoTip(c, tip)
		}
		arrstr, _ := json.Marshal(listImg)
		str = string(arrstr)
	}
	return str
}
