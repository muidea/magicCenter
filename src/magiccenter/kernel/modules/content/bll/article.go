package bll

import (
	"magiccenter/common/model"
	"magiccenter/kernel/modules/content/dal"
	"magiccenter/system"
	"time"
)

// QueryAllArticleSummary 查询全部文章摘要
func QueryAllArticleSummary() []model.ArticleSummary {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllArticleSummary(helper)
}

// QueryArticleByID 查询指定文章
func QueryArticleByID(id int) (model.Article, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	ar, result := dal.QueryArticleByID(helper, id)

	return ar, result
}

// DeleteArticle 删除文章
func DeleteArticle(id int) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteArticle(helper, id)
}

// CreateArticle 新建文章
func CreateArticle(title, content string, uID int, catalogs []int) (model.ArticleSummary, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	createDate := time.Now().Format("2006-01-02 15:04:05")
	return dal.CreateArticle(helper, title, content, catalogs, uID, createDate)
}

// SaveArticle 保存文章
func SaveArticle(article model.Article) (model.ArticleSummary, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	article.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	return dal.SaveArticle(helper, article)
}

// QueryArticleByCatalog 查询指定分类文章
func QueryArticleByCatalog(id int) []model.ArticleSummary {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryArticleByCatalog(helper, id)
}

// QueryArticleByRang 查询指定范围文章
func QueryArticleByRang(begin int, offset int) []model.ArticleSummary {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryArticleByRang(helper, begin, offset)
}
