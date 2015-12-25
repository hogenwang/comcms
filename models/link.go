package models

//获取友情链接详情
func GetLink(id int64) *Link {
	entity := &Link{Id: id}
	has, err := x.Get(entity)
	if has && err == nil {
		return entity
	} else {
		return nil
	}

}

//获取友情链接列表
func GetLinkList(where, order string, limit, start int) ([]*Link, int64) {
	list := make([]*Link, 0)

	var err error
	err = x.Where(where).OrderBy(order).Limit(limit, start).Find(&list)

	if err != nil {
		return nil, 0
	} else {
		Link := new(Link)
		total, _ := x.Where(where).Count(Link)
		return list, total
	}

}

//添加
func AddLink(entity *Link) (int64, error) {
	if _, err := x.Insert(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//编辑友情链接
func EditLink(entity *Link) (int64, error) {
	if _, err := x.Id(entity.Id).Update(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//删除友情链接
func DelLink(entity *Link) error {
	if _, err := x.Id(entity.Id).Delete(entity); err != nil {
		return err
	} else {
		return nil
	}
}
