package content

import (
    "webcenter/session"
    "webcenter/common"
)

type GetAllContentParam struct {
	session *session.Session
	accessCode string	
}

type GetAllContentResult struct {
	common.Result
	ArticleInfo []ArticleInfo
	Catalog []Catalog
}

type GetAllArticleParam struct {
	session *session.Session
	accessCode string	
}

type GetAllArticleResult struct {
	common.Result
	ArticleInfo []ArticleInfo
}

type GetArticleParam struct {
	session *session.Session
	accessCode string
	id int
}

type GetArticleReault struct {
	common.Result
	Article Article
}

type DeleteArticleParam struct {
	session *session.Session
	accessCode string
	id int
}

type DeleteArticleReault struct {
	common.Result
}

type GetCatalogParam struct {
	session *session.Session
	accessCode string
	id int
}

type GetCatalogReault struct {
	common.Result
	Catalog Catalog
}


type DeleteCatalogParam struct {
	session *session.Session
	accessCode string
	id int
}

type DeleteCatalogReault struct {
	common.Result
}

type GetAllCatalogParam struct {
	session *session.Session
	accessCode string	
}

type GetAllCatalogResult struct {
	common.Result
	Catalog []Catalog
}

type SubmitArticleParam struct {
	session *session.Session
	accessCode string
	id int
	title string
	content string
	catalog int
	author int
	submitDate string	
}

type SubmitArticleResult struct {
	common.Result
}


type SubmitCatalogParam struct {
	session *session.Session
	accessCode string
	id int
	name string
	author int
	submitDate string	
}

type SubmitCatalogResult struct {
	common.Result
}

type contentController struct {
}
 
func (this *contentController)getAllContentAction(param GetAllContentParam) GetAllContentResult {
	result := GetAllContentResult{}
	
	model, err := NewModel()
	if err != nil {
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	} 
	
	result.ArticleInfo = model.GetAllArticleInfo()
	result.Catalog = model.GetAllCatalog()
	result.ErrCode = 0

	model.Release()
	
	return result
}

func (this *contentController)getAllArticleAction(param GetAllArticleParam) GetAllArticleResult {
	result := GetAllArticleResult{}
	
	model, err := NewModel()
	if err != nil {
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	} 
	
	result.ArticleInfo = model.GetAllArticleInfo()
	result.ErrCode = 0

	model.Release()
	
	return result
}

func (this *contentController)getArticleAction(param GetArticleParam) GetArticleReault {
	result := GetArticleReault{}
	
	model, err := NewModel()
	if err != nil {
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	} 
	
	article, found := model.GetArticle(param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Article = article
	}
	
	model.Release()
	
	return result
}

func (this *contentController)deleteArticleAction(param DeleteArticleParam) DeleteArticleReault {
	result := DeleteArticleReault{}
	
	model, err := NewModel()
	if err != nil {
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	} 
	
	model.DeleteArticle(param.id)
	result.ErrCode = 0
	
	model.Release()
	
	return result
}
 
func (this *contentController)getAllCatalogAction(param GetAllCatalogParam) GetAllCatalogResult {
	result := GetAllCatalogResult{}
	
	model, err := NewModel()
	if err != nil {
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	} 
	
	result.Catalog = model.GetAllCatalog()
	result.ErrCode = 0
	model.Release()
	
	return result
}

 
func (this *contentController)getCatalogAction(param GetCatalogParam) GetCatalogReault {
	result := GetCatalogReault{}
	
	model, err := NewModel()
	if err != nil {
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	} 
	
	catalog, found := model.GetCatalog(param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Catalog = catalog
	}

	model.Release()

	return result
}

func (this *contentController)deleteCatalogAction(param DeleteCatalogParam) DeleteCatalogReault {
	result := DeleteCatalogReault{}
	
	model, err := NewModel()
	if err != nil {
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	
	articleInfoList := model.QueryArticleByCatalog(param.id)
	if (len(articleInfoList) >0) {
		result.ErrCode = 1
		result.Reason = "该分类被引用，无法立即删除"
		return result
	}
	
	model.DeleteCatalog(param.id)
	result.ErrCode = 0
	
	model.Release()
	
	return result
}

func (this *contentController)submitArticleAction(param SubmitArticleParam) SubmitArticleResult {
	result := SubmitArticleResult{}
	
	model, err := NewModel()
	if err != nil {
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	} 

	article := newArticle()
	article.Id = param.id
	article.Title = param.title
	article.Content = param.content
	article.Author.Id = param.author
	article.Catalog.Id = param.catalog
	article.CreateDate = param.submitDate	
	
	if !model.SaveArticle(article) {
		result.ErrCode = 1
		result.Reason = "保存文章失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存文章成功"
	}
	
	model.Release()

	return result
}

func (this *contentController)submitCatalogAction(param SubmitCatalogParam) SubmitCatalogResult {
	result := SubmitCatalogResult{}
	
	model, err := NewModel()
	if err != nil {
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	} 

	catalog := newCatalog()
	catalog.Id = param.id
	catalog.Name = param.name
	catalog.Creater.Id = param.author
	
	if !model.SaveCatalog(catalog) {
		result.ErrCode = 1
		result.Reason = "保存分类失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存分类成功"
	}
	
	model.Release()

	return result
}


