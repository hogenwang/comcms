package models

import (
	"fmt"
	"github.com/hogenwang/comcms/common"
	"time"
)

func AdminLogin(username, password, ip string) (*Admin, bool) {
	var entity = new(Admin)
	flag := false
	if len(username) < 4 || len(password) < 5 {
		return nil, flag
	}
	md5password := common.GetMd5String(password)
	entity = &Admin{UserName: username, PassWord: md5password}

	var err error
	has, err := x.Get(entity)
	//查询不到，返回
	if !has || err != nil {
		fmt.Println("没找到管理员：", entity)
		return nil, flag
	}
	fmt.Println("现在的：", entity)
	//写入记录
	entity.LastLoginIP = entity.ThisLoginIP
	entity.LastLoginTime = entity.ThisLoginTime
	entity.ThisLoginTime = time.Now()
	entity.ThisLoginIP = ip
	entity.LoginCount += 1
	fmt.Println("准备写入：", entity)
	_, err = x.Id(entity.Id).Update(entity)

	if err != nil {
		//写入失败记录
		fmt.Println("准备失败 ：", err.Error())
		return nil, flag
	}
	flag = true
	return entity, flag
}

//根据用户名获取管理员
func GetAdmin(username string) *Admin {
	var entity = new(Admin)
	entity = &Admin{UserName: username}
	var err error
	has, err := x.Get(entity)
	//查询不到，返回
	if !has || err != nil {
		return nil
	} else {
		return entity
	}
}

//根据ID获取管理员
func GetAdminById(id int64) *Admin {
	var entity = new(Admin)
	entity = &Admin{Id: id}
	var err error
	has, err := x.Get(entity)
	//查询不到，返回
	if !has || err != nil {
		return nil
	} else {
		return entity
	}
}

//添加
func AddAdmin(entity *Admin) (int64, error) {
	if _, err := x.Insert(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//编辑管理员
func EditAdmin(entity *Admin) (int64, error) {
	if _, err := x.Id(entity.Id).Update(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//删除文章
func DelAdmin(entity *Admin) error {
	if _, err := x.Id(entity.Id).Delete(entity); err != nil {
		return err
	} else {
		return nil
	}
}

//管理员
func GetAdminList(where, order string, limit, start int) ([]*Admin, int64) {
	list := make([]*Admin, 0)

	var err error
	err = x.Where(where).OrderBy(order).Limit(limit, start).Find(&list)

	if err != nil {
		return nil, 0
	} else {
		entity := new(Admin)
		total, _ := x.Where(where).Count(entity)
		return list, total
	}

}
