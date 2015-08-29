package content

import (
    "webcenter/session"
    "webcenter/common"
)

type AllArticleInfoResult struct {
	common.Result
	ArticleInfo []ArticleInfo
}

type AllCatalogResult struct {
	common.Result
	Catalog []Catalog
}

type contentController struct {
}
 
func (this *contentController)getAllArticleInfoAction(session *session.Session) AllArticleInfoResult {
	result := AllArticleInfoResult{}
	
	model, err := NewModel()
	if err != nil {
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	} 
	
	result.ArticleInfo = model.GetAllArticleInfo()
		
	model.Release()
	
	return result
}

 
func (this *contentController)getAllCatalogAction(session *session.Session) AllCatalogResult {
	result := AllCatalogResult{}
	
	model, err := NewModel()
	if err != nil {
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	} 
	
	result.Catalog = model.GetAllCatalog()
		
	model.Release()
	
	return result
}
