package article

import (
	"webcenter/common"
	"webcenter/modelhelper"
)


type QueryManageInfo struct {
	ArticleInfo []ArticleSummary
}

type QueryAllArticleParam struct {
	accessCode string	
}

type QueryAllArticleResult struct {
	common.Result
	ArticleInfo []ArticleSummary
}

type QueryArticleParam struct {
	accessCode string
	id int
}

type QueryArticleResult struct {
	common.Result
	Article ArticleDetail
}

type DeleteArticleParam struct {
	accessCode string
	id int
}

type DeleteArticleResult struct {
	common.Result
}

type SubmitArticleParam struct {
	id int
	title string
	content string
	catalog []int
	submitDate string
	author int	
}

type SubmitArticleResult struct {
	common.Result
}


type EditArticleParam struct {
	accessCode string
	id int
}

type EditArticleResult struct {
	common.Result
	Article ArticleDetail
}

type articleController struct {
}

func (this *articleController)queryManageInfoAction() QueryManageInfo {
	info := QueryManageInfo{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
			
	info.ArticleInfo = QueryAllArticle(model)

	return info
}

func (this *articleController)queryAllArticleAction(param QueryAllArticleParam) QueryAllArticleResult {
	result := QueryAllArticleResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
			
	result.ArticleInfo = QueryAllArticle(model)
	result.ErrCode = 0

	return result
}

func (this *articleController)queryArticleAction(param QueryArticleParam) QueryArticleResult {
	result := QueryArticleResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	article, found := QueryArticleById(model, param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Article = article
	}
	
	return result
}

func (this *articleController)deleteArticleAction(param DeleteArticleParam) DeleteArticleResult {
	result := DeleteArticleResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
		
	if DeleteArticleById(model, param.id) {
		result.ErrCode = 0
		result.Reason = "删除文章成功"
	} else {
		result.ErrCode = 1
		result.Reason = "删除文章出错"		
	}
	
	return result
}
 

func (this *articleController)submitArticleAction(param SubmitArticleParam) SubmitArticleResult {
	result := SubmitArticleResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	article := NewArticle()
	article.SetId(param.id)
	article.SetName(param.title)
	article.SetContent(param.content)
	article.SetCreateDate(param.submitDate)
	article.SetCatalog(param.catalog)
	article.SetAuthor(param.author)
	
	if !SaveArticle(model, article) {
		result.ErrCode = 1
		result.Reason = "保存文章失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存文章成功"
	}
	
	return result
}


func (this *articleController)editArticleAction(param EditArticleParam) EditArticleResult {
	result := EditArticleResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	article, found := QueryArticleById(model, param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Article = article
	}
	
	return result
}




