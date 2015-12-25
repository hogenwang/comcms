package models

//获取文章详情
func GetArticle(id int64) *Article {
	entity := &Article{Id: id}
	has, err := x.Get(entity)
	if has && err == nil {
		return entity
	} else {
		return nil
	}

}

//获取文章列表
func GetArticleList(where, order string, limit, start int) ([]*Article, int64) {
	list := make([]*Article, 0)

	var err error
	err = x.Where(where).OrderBy(order).Limit(limit, start).Find(&list)

	if err != nil {
		return nil, 0
	} else {
		article := new(Article)
		total, _ := x.Where(where).Count(article)
		return list, total
	}

}

//添加
func AddArticle(entity *Article) (int64, error) {
	if _, err := x.Insert(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//编辑文章
func EditArticle(entity *Article) (int64, error) {
	if _, err := x.Id(entity.Id).Update(entity); err != nil {
		return -1, err
	} else {
		return entity.Id, nil
	}
}

//删除文章
func DelArticle(entity *Article) error {
	if _, err := x.Id(entity.Id).Delete(entity); err != nil {
		return err
	} else {
		return nil
	}
}
