package bll

import (
	commonmodel "magiccenter/common/model"
	"magiccenter/system"
)

// ContentModuleID ID
const ContentModuleID = "3a7123ec-63f0-5e46-1234-e6ca1af6fe4e"

// QueryContentMetaRequest 查询内容元数据请求
type QueryContentMetaRequest struct {
}

// QueryContentMetaResponse 查询内容元数据响应
type QueryContentMetaResponse struct {
	ContentMetas []commonmodel.ContentMeta
}

// QueryContentMetas 查询内容元数据
func QueryContentMetas() ([]commonmodel.ContentMeta, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryContentMetaRequest{}

	response := QueryContentMetaResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.ContentMetas, result
}

// QueryContentArticleRequest 查询文章列表请求
type QueryContentArticleRequest struct {
}

// QueryContentArticleResponse 查询文章列表响应
type QueryContentArticleResponse struct {
	Articles []commonmodel.ArticleSummary
}

// QueryContentArticles 查询文章列表
func QueryContentArticles() ([]commonmodel.ArticleSummary, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryContentArticleRequest{}

	response := QueryContentArticleResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Articles, result
}

// QuerySingleArticleRequest 查询文章列表请求
type QuerySingleArticleRequest struct {
	ID int
}

// QuerySingleArticleResponse 查询文章列表响应
type QuerySingleArticleResponse struct {
	Found   bool
	Article commonmodel.Article
}

// QuerySingleArticle 查询指定文章
func QuerySingleArticle(id int) (commonmodel.Article, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QuerySingleArticleRequest{ID: id}

	response := QuerySingleArticleResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Article, response.Found && result
}

// CreateArticleRequest 新建文章请求
type CreateArticleRequest struct {
	Title   string
	Content string
	Catalog []int
	Creater int
}

// CreateArticleResponse 新建文章响应
type CreateArticleResponse struct {
	Result  bool
	Article commonmodel.ArticleSummary
}

// CreateArticle 新建文章
func CreateArticle(title, content string, catalogs []int, creater int) (commonmodel.ArticleSummary, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := CreateArticleRequest{Title: title, Content: content, Catalog: catalogs, Creater: creater}

	response := CreateArticleResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Article, response.Result && result
}

// UpdateArticleRequest 更新文章请求
type UpdateArticleRequest struct {
	Article commonmodel.Article
}

// UpdateArticleResponse 更新文章响应
type UpdateArticleResponse struct {
	Result  bool
	Article commonmodel.ArticleSummary
}

// UpdateArticle 更新文章
func UpdateArticle(article commonmodel.Article) (commonmodel.ArticleSummary, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := UpdateArticleRequest{}
	request.Article = article

	response := UpdateArticleResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Article, response.Result && result
}

// DeleteArticleRequest 删除文章请求
type DeleteArticleRequest struct {
	ID int
}

// DeleteArticleResponse 删除文章响应
type DeleteArticleResponse struct {
	Result bool
}

// DeleteArticle 删除文章
func DeleteArticle(id int) bool {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := DeleteArticleRequest{ID: id}

	response := DeleteArticleResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Result && result
}

// QueryContentCatalogRequest 查询分类列表请求
type QueryContentCatalogRequest struct {
}

// QueryContentCatalogResponse 查询分类列表响应
type QueryContentCatalogResponse struct {
	Catalogs []commonmodel.Catalog
}

// QueryContentCatalogs 查询分类列表
func QueryContentCatalogs() ([]commonmodel.Catalog, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryContentCatalogRequest{}

	response := QueryContentCatalogResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Catalogs, result
}

// QueryContentLinkRequest 查询链接列表请求
type QueryContentLinkRequest struct {
}

// QueryContentLinkResponse 查询链接列表响应
type QueryContentLinkResponse struct {
	Links []commonmodel.Link
}

// QueryContentLinks 查询链接列表
func QueryContentLinks() ([]commonmodel.Link, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryContentLinkRequest{}

	response := QueryContentLinkResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Links, result
}

// QueryContentMediaRequest 查询文件列表请求
type QueryContentMediaRequest struct {
}

// QueryContentMediaResponse 查询文件列表响应
type QueryContentMediaResponse struct {
	Medias []commonmodel.MediaDetail
}

// QueryContentMedias 查询文件列表
func QueryContentMedias() ([]commonmodel.MediaDetail, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryContentMediaRequest{}

	response := QueryContentMediaResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Medias, result
}
