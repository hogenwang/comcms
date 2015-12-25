package admin

import (
	"github.com/astaxie/beego"
	"github.com/hogenwang/comcms/models"
)

type ConfigController struct {
	beego.Controller
}

//获取配置
func (c *ConfigController) Get() {
	CheckAdminLogin(&c.Controller, 0)
	cfg := &models.Config{}
	cfg = models.GetConfig()
	c.Data["Title"] = "修改系统设置"
	c.Data["Cfg"] = cfg
	c.TplNames = "admin/config.tpl"
}

//执行更新
func (c *ConfigController) Post() {
	CheckAdminLogin(&c.Controller, 1)
	tip := &models.TipJSON{}
	tip.Status = models.TipError
	cfg := &models.Config{}
	cfg = models.GetConfig()
	if cfg == nil {
		tip.Message = "系统找不到本记录"
		c.Data["json"] = tip
		c.ServeJson()
		c.StopRun()
	}
	cfg.SiteName = c.GetString("SiteName")
	if cfg.SiteName == "" {
		tip.Message = "请输入站点名称"
		c.Data["json"] = tip
		c.ServeJson()
		c.StopRun()
	}
	cfg.SiteURL = c.GetString("SiteURL")
	cfg.SiteLogo = c.GetString("SiteLogo")
	cfg.SiteEmail = c.GetString("SiteEmail")
	cfg.ICP = c.GetString("ICP")
	cfg.Copyright = c.GetString("Copyright")
	cfg.OnlineQQ = c.GetString("OnlineQQ")
	cfg.SiteTitle = c.GetString("SiteTitle")
	cfg.Keyword = c.GetString("Keyword")
	cfg.Description = c.GetString("Description")
	err := models.EditConfig(cfg)
	if err == nil {
		tip.Status = models.TipSuccess
		tip.Message = "修改站点配置成功"
		c.Data["json"] = tip
		c.ServeJson()
		c.StopRun()
	} else {
		tip.Message = "修改站点配置失败：" + err.Error()
		c.Data["json"] = tip
		c.ServeJson()
		c.StopRun()
	}
}
