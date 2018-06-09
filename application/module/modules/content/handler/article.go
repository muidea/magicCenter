package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/module/modules/content/dal"
	"muidea.com/magicCommon/model"
)

type articleActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *articleActionHandler) getAllArticleSummary() []model.Summary {
	return dal.QueryAllArticleSummary(i.dbhelper)
}

func (i *articleActionHandler) getArticles(ids []int) []model.Article {
	return dal.QueryArticles(i.dbhelper, ids)
}

func (i *articleActionHandler) findArticleByID(id int) (model.ArticleDetail, bool) {
	return dal.QueryArticleByID(i.dbhelper, id)
}

func (i *articleActionHandler) findArticleByCatalog(catalog int) []model.Summary {
	return dal.QueryArticleSummaryByCatalog(i.dbhelper, catalog)
}

func (i *articleActionHandler) createArticle(title, content, createDate string, catalog []int, author int) (model.Summary, bool) {
	return dal.CreateArticle(i.dbhelper, title, content, catalog, author, createDate)
}

func (i *articleActionHandler) saveArticle(article model.ArticleDetail) (model.Summary, bool) {
	return dal.SaveArticle(i.dbhelper, article)
}

func (i *articleActionHandler) destroyArticle(id int) bool {
	return dal.DeleteArticle(i.dbhelper, id)
}
