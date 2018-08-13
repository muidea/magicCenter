package handler

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/content/dal"
	"muidea.com/magicCommon/model"
)

type articleActionHandler struct {
}

func (i *articleActionHandler) getAllArticleSummary() []model.Summary {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllArticleSummary(dbhelper)
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

func (i *articleActionHandler) findArticleByCatalog(catalog model.CatalogUnit) []model.Summary {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryArticleSummaryByCatalog(dbhelper, catalog)
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
