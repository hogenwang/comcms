package models

//获取广告详情
func GetAds(id int64) *Ads {
	entity := &Ads{Id: id}
	has, err := x.Get(entity)
	if has && err == nil {
		return entity
	} else {
		return nil
	}

}

//获取广告列表
func GetAdsList(where, order string, limit, start int) ([]*Ads, int64) {
	list := make([]*Ads, 0)

	var err error
	err = x.Where(where).OrderBy(order).Limit(limit, start).Find(&list)

	if err != nil {
		return nil, 0
	} else {
		Ads := new(Ads)
		total, _ := x.Where(where).Count(Ads)
		return list, total
	}

}

//添加
func AddAds(entity *Ads) (int64, error) {
	if _, err := x.Insert(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//编辑广告
func EditAds(entity *Ads) (int64, error) {
	if _, err := x.Id(entity.Id).Update(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//删除广告
func DelAds(entity *Ads) error {
	if _, err := x.Id(entity.Id).Delete(entity); err != nil {
		return err
	} else {
		return nil
	}
}
