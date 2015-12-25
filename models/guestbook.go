package models

//获取留言板详情
func GetGuestbook(id int64) *Guestbook {
	entity := &Guestbook{Id: id}
	has, err := x.Get(entity)
	if has && err == nil {
		return entity
	} else {
		return nil
	}

}

//获取留言板列表
func GetGuestbookList(where, order string, limit, start int) ([]*Guestbook, int64) {
	list := make([]*Guestbook, 0)

	var err error
	err = x.Where(where).OrderBy(order).Limit(limit, start).Find(&list)

	if err != nil {
		return nil, 0
	} else {
		Guestbook := new(Guestbook)
		total, _ := x.Where(where).Count(Guestbook)
		return list, total
	}

}

//添加
func AddGuestbook(entity *Guestbook) (int64, error) {
	if _, err := x.Insert(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//编辑留言板
func EditGuestbook(entity *Guestbook) (int64, error) {
	if _, err := x.Id(entity.Id).Update(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//删除留言板
func DelGuestbook(entity *Guestbook) error {
	if _, err := x.Id(entity.Id).Delete(entity); err != nil {
		return err
	} else {
		return nil
	}
}
