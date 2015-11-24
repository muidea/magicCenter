package ui

import (
	"log"
	"webcenter/common"
)

type ArticleView struct {
	common.Result
	Article ArticleContent
	ArticleCatalog []ArticleCatalog
	SiteLink []SiteLink
}


type IndexView struct {
	common.Result
	ArticleSummary []ArticleSummary
	ArticleCatalog []ArticleCatalog
	SiteLink []SiteLink
}

type uiController struct {
}

func (this *uiController)ViewArticleAction(id int) ArticleView {
	log.Print("ViewAction");
	
	view := ArticleView{}
	
	model, err := newModel()
	if err != nil {
		log.Print("create uiModel failed")
		view.ErrCode = 1
		view.Reason = "创建Model失败"
		return view
	}
	defer model.Release()
	
	articleview, found := model.GetArticleView(id)
	if !found {
		log.Printf("can't find article,id:%d", id)
		view.ErrCode = 1
		view.Reason = "找不到指定内容"
		return view
	}
	
	view.Article = articleview;
	view.ArticleCatalog = model.GetArticleCatalog()
	view.SiteLink = model.GetSiteLink()
	view.ErrCode = 0;
	
	return view
}

func (this *uiController)IndexAction() IndexView {
	log.Print("IndexAction");
	
	view := IndexView{}
	
	model, err := newModel()
	if err != nil {
		log.Print("create uiModel failed")
		view.ErrCode = 1
		view.Reason = "创建Model失败"
		return view
	}
	defer model.Release()
	
	view.ArticleSummary = model.GetArticleSummary(0,4)
	view.ArticleCatalog = model.GetArticleCatalog()
	view.SiteLink = model.GetSiteLink()
	view.ErrCode = 0

	return view
}

