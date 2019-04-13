package handler

import (
	"github.com/muidea/magicCenter/common/dbhelper"
	"github.com/muidea/magicCenter/module/modules/content/dal"
	"github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/model"
)

type articleActionHandler struct {
}

func (i *articleActionHandler) getAllArticleSummary(filter *def.Filter) ([]model.Summary, int) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllArticleSummary(dbhelper, filter)
}

func (i *articleActionHandler) getArticles(ids []int) []model.Article {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryArticles(dbhelper, ids)
}

func (i *articleActionHandler) findArticleByID(id int) (model.ArticleDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryArticleByID(dbhelper, id)
}

func (i *articleActionHandler) findArticleByCatalog(catalog model.CatalogUnit, filter *def.Filter) ([]model.Summary, int) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryArticleSummaryByCatalog(dbhelper, catalog, filter)
}

func (i *articleActionHandler) createArticle(title, content, createDate string, catalog []model.CatalogUnit, author int) (model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.CreateArticle(dbhelper, title, content, catalog, author, createDate)
}

func (i *articleActionHandler) saveArticle(article model.ArticleDetail) (model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.SaveArticle(dbhelper, article)
}

func (i *articleActionHandler) destroyArticle(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.DeleteArticle(dbhelper, id)
}
