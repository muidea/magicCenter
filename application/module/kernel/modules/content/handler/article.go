package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/content/dal"
)

type articleActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *articleActionHandler) getAllArticleSummary() []model.ArticleSummary {
	return dal.QueryAllArticleSummary(i.dbhelper)
}

func (i *articleActionHandler) findArticleByID(id int) (model.Article, bool) {
	return dal.QueryArticleByID(i.dbhelper, id)
}

func (i *articleActionHandler) findArticleByCatalog(catalog int) []model.ArticleSummary {
	return dal.QueryArticleSummaryByCatalog(i.dbhelper, catalog)
}

func (i *articleActionHandler) createArticle(title, content, createDate string, catalog []int, author int) (model.ArticleSummary, bool) {
	return dal.CreateArticle(i.dbhelper, title, content, catalog, author, createDate)
}

func (i *articleActionHandler) saveArticle(article model.Article) (model.ArticleSummary, bool) {
	return dal.SaveArticle(i.dbhelper, article)
}

func (i *articleActionHandler) destroyArticle(id int) bool {
	return dal.DeleteArticle(i.dbhelper, id)
}
