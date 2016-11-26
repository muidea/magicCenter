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

// QueryContentArticleRequest 查询文章列表请求
type QueryContentArticleRequest struct {
}

// QueryContentArticleResponse 查询文章列表响应
type QueryContentArticleResponse struct {
	Articles []commonmodel.ArticleSummary
}

// QueryContentCatalogRequest 查询分类列表请求
type QueryContentCatalogRequest struct {
}

// QueryContentCatalogResponse 查询分类列表响应
type QueryContentCatalogResponse struct {
	Catalogs []commonmodel.Catalog
}

// QueryContentLinkRequest 查询链接列表请求
type QueryContentLinkRequest struct {
}

// QueryContentLinkResponse 查询链接列表响应
type QueryContentLinkResponse struct {
	Links []commonmodel.Link
}

// QueryContentMediaRequest 查询文件列表请求
type QueryContentMediaRequest struct {
}

// QueryContentMediaResponse 查询文件列表响应
type QueryContentMediaResponse struct {
	Medias []commonmodel.MediaDetail
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
