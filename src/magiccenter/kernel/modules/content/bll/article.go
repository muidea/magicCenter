package bll

import (
	"magiccenter/common/model"
	"magiccenter/kernel/modules/content/dal"
	"magiccenter/util/dbhelper"
	"time"
)

// QueryAllArticleSummary 查询全部文章摘要
func QueryAllArticleSummary() []model.ArticleSummary {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllArticleSummary(helper)
}

// QueryArticleByID 查询指定文章
func QueryArticleByID(id int) (model.Article, bool) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	ar, result := dal.QueryArticleByID(helper, id)

	return ar, result
}

// DeleteArticle 删除文章
func DeleteArticle(id int) bool {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteArticle(helper, id)
}

// SaveArticle 保存文章
func SaveArticle(id int, title, content string, uID int, catalogs []int) bool {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	article := model.Article{}
	article.ID = id
	article.Title = title
	article.Content = content
	article.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	article.Author = uID
	article.Catalog = catalogs

	return dal.SaveArticle(helper, article)
}

// QueryArticleByCatalog 查询指定分类文章
func QueryArticleByCatalog(id int) []model.ArticleSummary {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryArticleByCatalog(helper, id)
}

// QueryArticleByRang 查询指定范围文章
func QueryArticleByRang(begin int, offset int) []model.ArticleSummary {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryArticleByRang(helper, begin, offset)
}
