package article

import (
	"log"
	"webcenter/common"
	"webcenter/modelhelper"
)

type QueryArticleParam struct {
	accessCode string
	id int
}

type QueryArticleResult struct {
	common.Result
	Article Article
}

type DeleteArticleParam struct {
	accessCode string
	id int
}

type DeleteArticleResult struct {
	common.Result
}

type QueryAllArticleInfoParam struct {
	accessCode string	
}

type QueryAllArticleInfoResult struct {
	common.Result
	ArticleInfo []ArticleInfo
}

type SubmitArticleParam struct {
	id int
	title string
	content string
	catalog int
	submitDate string
	author int	
}

type SubmitArticleResult struct {
	common.Result
}

type articleController struct {
}


func (this *articleController)queryAllArticleInfoAction(param QueryAllArticleInfoParam) QueryAllArticleInfoResult {
	result := QueryAllArticleInfoResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
			
	result.ArticleInfo = GetAllArticleInfo(model)
	result.ErrCode = 0

	return result
}

func (this *articleController)queryArticleAction(param QueryArticleParam) QueryArticleResult {
	result := QueryArticleResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
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
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
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
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
	
	article := newArticle()
	article.Id = param.id
	article.Title = param.title
	article.Content = param.content
	article.Author = param.author
	article.Catalog = param.catalog
	article.CreateDate = param.submitDate	
	
	if !SaveArticle(model, article) {
		result.ErrCode = 1
		result.Reason = "保存文章失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存文章成功"
	}
	
	return result
}





