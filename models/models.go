package models

import (
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

//XORM 的引擎
var x *xorm.Engine

const (
	//数据库类型 可选 sqlite mysql
	dbtype        = "mysql"
	TipSuccess    = "success"
	TipError      = "error"
	ComCMSVersion = "0.1.0" //版本
)

func init() {
	var err error
	switch {
	case dbtype == "mysql":
		x, err = xorm.NewEngine("mysql", "root:root@/comcms?charset=utf8")
	case dbtype == "sqlite":
		x, err = xorm.NewEngine("sqlite3", "./data/sqlite.db")
	}
	//x, err = xorm.NewEngine("mysql", "root:root@/comcms?charset=utf8")

	if err != nil {
		log.Fatalf("fail to create engine: %v", err)
	}
	err = x.Sync2(new(Ads), new(Link), new(Guestbook), new(Config), new(Article), new(Admin), new(AdminRole), new(Category))
	x.ShowSQL = true
	x.ShowDebug = true
	x.ShowErr = true
	//设置缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 10000)
	x.SetDefaultCacher(cacher)

}

/*---这里开始是所有模型---*/
//系统提示JSON。
type TipJSON struct {
	Id        int64
	Status    string
	Message   string
	ReturnUrl string
	Other     string
}

//管理员
type Admin struct {
	Id            int64
	UserName      string `xorm:"varchar(20) unique"`
	PassWord      string `xorm:"varchar(50)"`
	NickName      string `xorm:"varchar(20)"`
	RoleId        int64  `xorm:"index default(0)"`
	LastLoginTime time.Time
	LastLoginIP   string    `xorm:"varchar(20)"`
	ThisLoginTime time.Time `xorm:"updated"`
	ThisLoginIP   string    `xorm:"varchar(20)"`
	IsLock        int64     `xorm:"index default(0)"`
	EditorId      int64     `xorm:"index default(0)"`
	Notes         string    `xorm:"varchar(250)"`
	LoginCount    int64     `xorm:"index default(0)"`
}

//管理组
type AdminRole struct {
	Id              int64
	RoleName        string `xorm:"unique"`
	RoleDescription string `xorm:"varchar(250)"`
	IsSuperAdmin    int64
	Stars           int64
	Color           string `xorm:"varchar(20)"`
	NotAllowDel     int64
	Rank            int64
}

//系统配置
type Config struct {
	Id             int64
	SiteName       string    `xorm:"varchar(100)"`
	SiteURL        string    `xorm:"varchar(100)"`
	SiteLogo       string    `xorm:"varchar(200)"`
	ICP            string    `xorm:"varchar(200)"`
	SiteEmail      string    `xorm:"varchar(50)"`
	Copyright      string    `xorm:"text"`
	IsCloseSite    int64     `xorm:"index default(0)"`
	CloseReason    string    `xorm:"text"`
	Keyword        string    `xorm:"varchar(250)"`
	Description    string    `xorm:"varchar(250)"`
	SiteTitle      string    `xorm:"varchar(250)"`
	SearchMinTime  int64     `xorm:"default(30)"`
	OnlineQQ       string    `xorm:"varchar(250)"`
	OnlineSkype    string    `xorm:"varchar(250)"`
	OnlineWangWang string    `xorm:"varchar(250)"`
	Skin           string    `xorm:"varchar(50)"`
	LastUpdateTime time.Time `xorm:"updated" `
}

//分类
type Category struct {
	Id                 int64
	Pid                int64     `xorm:"index"`
	Ctype              int64     `xorm:"index"`
	Title              string    `xorm:"index varchar(250)"`
	Content            string    `xorm:"text index"`
	PageTitle          string    `xorm:"index"`
	Rank               int64     `xorm:"index"`
	Level              int64     `xorm:"index"`
	Keyword            string    `xorm:"varchar(250)"`
	Description        string    `xorm:"varchar(250)"`
	LinkUrl            string    `xorm:"varchar(250)"`
	TitleColor         string    `xorm:"varchar(20)"`
	TemplateFile       string    `xorm:"varchar(250)"`
	DetailTemplateFile string    `xorm:"varchar(250)"`
	IsList             int64     `xorm:"index default(1)"`
	PageSize           int64     `xorm:"index default(15)"`
	IsLock             int64     `xorm:"index"`
	IsDel              int64     `xorm:"index"`
	IsHide             int64     `xorm:"index"`
	IsDisabled         int64     `xorm:"index"`
	IsComment          int64     `xorm:"index"`
	IsHeaderNav        int64     `xorm:"index"`
	IsFooterNav        int64     `xorm:"index"`
	Counts             int64     `xorm:"index"`
	Created            time.Time `xorm:"created index"`
	Updated            time.Time `xorm:"updated index"`
	CatalogId          int64     `xorm:"index"`
	Pic                string    `xorm:"varchar(250) index"`
	AdsId              int64     `xorm:"index"`
}

//文章
type Article struct {
	Id           int64
	Kid          int64     `xorm:"index"`
	Title        string    `xorm:"index varchar(250)"`
	Content      string    `xorm:"text index"`
	Created      time.Time `xorm:"created index"`
	Updated      time.Time `xorm:"updated index"`
	AuthorId     int64     `xorm:"index"`
	Origin       string    `xorm:"varchar(250) index"`
	OriginUrl    string    `xorm:"varchar(250) index"`
	Rank         int64     `xorm:"index"`
	Keyword      string    `xorm:"varchar(250)"`
	Description  string    `xorm:"varchar(250)"`
	LinkUrl      string    `xorm:"varchar(250)"`
	TitleColor   string    `xorm:"varchar(20)"`
	Pic          string    `xorm:"varchar(250) index"`
	Tag          string    `xorm:"varchar(250) index"`
	TemplateFile string    `xorm:"varchar(250) index"`
	FileName     string    `xorm:"varchar(250) index"`
	Views        int64     `xorm:"index"`
	IsPass       int64     `xorm:"index default(1)"`
	IsRecommend  int64     `xorm:"index"`
	IsTop        int64     `xorm:"index"`
	IsBest       int64     `xorm:"index"`
	IsNew        int64     `xorm:"index"`
	IsDel        int64     `xorm:"index"`
	IsMember     int64     `xorm:"index"`
	IsHide       int64     `xorm:"index"`
	CommentCount int64     `xorm:"index"`
}

//友情链接分类
type LinkKind struct {
	Id          int64
	Title       string `xorm:"index varchar(250)"`
	Rank        int64  `xorm:"index default(0)"`
	Description string `xorm:"varchar(250)"`
}

//友情链接
type Link struct {
	Id          int64
	Kid         int64  `xorm:"index default(0)"`
	Title       string `xorm:"index varchar(250)"`
	Rank        int64  `xorm:"index default(0)"`
	Url         string `xorm:"varchar(250)"`
	Description string `xorm:"varchar(250)"`
	IsHide      int64  `xorm:"index default(0)"`
	Logo        string `xorm:"varchar(250)"`
}

//留言板分类
type GuestbookKind struct {
	Id          int64
	Title       string `xorm:"index varchar(250)"`
	Rank        int64  `xorm:"index default(0)"`
	Description string `xorm:"varchar(250)"`
	IsHide      int64  `xorm:"index default(0)"`
}

//留言板
type Guestbook struct {
	Id       int64
	Kid      int64     `xorm:"index default(0)"`
	Title    string    `xorm:"index varchar(250)"`
	Content  int64     `xorm:"text index"`
	Uid      int64     `xorm:"index default(0)"`
	UserName string    `xorm:"index varchar(50)"`
	UserImg  string    `xorm:"index varchar(250)"`
	Created  time.Time `xorm:"created index"`
	IsVerify int64     `xorm:"index default(0)"`
	IsRead   int64     `xorm:"index default(0)"`
	IsDel    int64     `xorm:"index default(0)"`
	IP       string    `xorm:"varchar(20)"`
	Email    string    `xorm:"index varchar(100)"`
	Tel      string    `xorm:"index varchar(200)"`
	QQ       string    `xorm:"index varchar(20)"`
	Skype    string    `xorm:"index varchar(200)"`
}

//广告分类
type AdsKind struct {
	Id          int64
	Title       string `xorm:"index varchar(250)"`
	Rank        int64  `xorm:"index default(0)"`
	Description string `xorm:"varchar(250)"`
}

//广告
type Ads struct {
	Id          int64
	Kid         int64  `xorm:"index default(0)"`
	Tid         int64  `xorm:"index default(0)"` //广告代码类型：0代码、1文字广告、2图片广告、3Flash广告、4幻灯片广告
	Title       string `xorm:"index varchar(250)"`
	Rank        int64  `xorm:"index default(0)"`
	Description string `xorm:"varchar(250)"`
	Content     string `xorm:"text index"`
	IsHide      int64  `xorm:"index default(0)"`
}

//广告详情
//代码广告
type ScriptAds struct {
	Content string
}

//文字广告
type TextAds struct {
	Txt   string
	Link  string
	Style string
}

//图片广告
type ImgAds struct {
	Img    string
	Link   string
	Width  int64
	Height int64
	Alt    string
}

//Flash 广告
type FlashAds struct {
	Swf    string
	Width  int64
	Height int64
}
