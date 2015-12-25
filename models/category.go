package models

import (
	"github.com/hogenwang/comcms/common"
	"strings"
)

//获取分类详情
func GetCategory(id int64) *Category {
	category := &Category{Id: id}
	has, err := x.Get(category)
	if has && err == nil {
		return category
	} else {
		return nil
	}

}

//添加分类
func AddCategory(entity *Category) (int64, error) {
	if _, err := x.Insert(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//编辑分类
func EditCategory(entity *Category) (int64, error) {
	if _, err := x.Id(entity.Id).Update(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//删除文章
func DelCategory(entity *Category) error {
	if _, err := x.Id(entity.Id).Delete(entity); err != nil {
		return err
	} else {
		return nil
	}
}

//获取分类
func GetCategoryList(key string, ishide bool) []*Category {
	categorys := make([]*Category, 0)
	var where string
	if ishide {
		where = " ishide = 1 "
	} else {
		where = " 1 = 1 "
	}
	var err error
	if key != "" {
		where += " Title like '%?%'"
		err = x.Where(where, key).Find(&categorys)

	} else {
		err = x.Where(where).Find(&categorys)
	}
	if err != nil {
		return nil
	} else {
		return categorys
	}

}
func GetCategoryListByParentId(pid int64) []*Category {
	categorys := make([]*Category, 0)
	var err error
	err = x.Where("pid = ?", pid).OrderBy("Rank asc,Id asc").Find(&categorys)
	if err != nil {
		return nil
	} else {
		return categorys
	}
}

//获取本分类下属分类列表

//获取分类 pid:上级栏目ID，maxLevel 最大的级别，从1 开始，isIndentation：是否缩进；removeid 不显示此ID及其下属级别
func GetCategoryTree(pid, maxLevel int64, isIndentation bool, removeid int64) []*Category {
	listTree := make([]*Category, 0)
	list := GetCategoryListByParentId(pid)

	if len(list) > 0 {
		for _, v := range list {
			//如果不是移除的 则循环
			if removeid == 0 || v.Id != removeid {
				//添加进去
				listTree = append(listTree, v)
				//如果未达到最大的级别，则递归
				if maxLevel == -1 || v.Level < maxLevel {
					listTmp := GetCategoryTree(v.Id, maxLevel, isIndentation, removeid)
					if len(listTmp) > 0 {
						for _, item := range listTmp {
							if item.Level > 0 && isIndentation {
								temp := &Category{}
								if common.DeepCopy(temp, item) == nil {
									temp.Title = strings.Repeat("　", int(item.Level*2)) + item.Title
									listTree = append(listTree, temp)
								} else {
									listTree = append(listTree, item)
								}

							} else {
								listTree = append(listTree, item)
							}

						}
					}
				}
			}
		}
	}
	return listTree
}
