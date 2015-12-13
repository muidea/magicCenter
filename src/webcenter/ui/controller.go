package ui

import (
	"log"
	"webcenter/common"
	"webcenter/modelhelper"
)

type ArticleView struct {
	common.Result
	Article ArticleContent
	ArticleCatalog []CatalogContent
	SiteLink []LinkContent
}


type IndexView struct {
	common.Result
	ArticleSummary []ArticleContent
	ArticleCatalog []CatalogContent
	SiteLink []LinkContent
}

type uiController struct {
}

func (this *uiController)ViewArticleAction(id int) ArticleView {
	log.Print("ViewAction");
	
	view := ArticleView{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("new model failed")
	}
	defer model.Release()
	
	articleview, found := GetArticleContent(model, id)
	if !found {
		log.Printf("can't find article,id:%d", id)
		view.ErrCode = 1
		view.Reason = "找不到指定内容"
		return view
	}
	
	view.Article = articleview;
	view.ArticleCatalog = GetCatalog(model)
	view.SiteLink = GetLink(model)
	view.ErrCode = 0;
	
	return view
}

func (this *uiController)IndexAction() IndexView {
	log.Print("IndexAction");
	
	view := IndexView{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("new model failed")
	}
	defer model.Release()	
	
	view.ArticleSummary = GetArticleSummary(model, 0,4)
	view.ArticleCatalog = GetCatalog(model)
	view.SiteLink = GetLink(model)
	view.ErrCode = 0

	return view
}

