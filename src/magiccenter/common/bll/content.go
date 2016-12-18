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

//===========================Article====================================

// QueryContentArticleListRequest 查询文章列表请求
type QueryContentArticleListRequest struct {
}

// QueryContentArticleListResponse 查询文章列表响应
type QueryContentArticleListResponse struct {
	Articles []commonmodel.ArticleSummary
}

// QueryContentArticles 查询文章列表
func QueryContentArticles() ([]commonmodel.ArticleSummary, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryContentArticleListRequest{}

	response := QueryContentArticleListResponse{}
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

//===========================Catalog====================================

// QueryContentCatalogListRequest 查询分类列表请求
type QueryContentCatalogListRequest struct {
}

// QueryContentCatalogListResponse 查询分类列表响应
type QueryContentCatalogListResponse struct {
	Catalogs []commonmodel.Catalog
}

// QueryContentCatalogs 查询分类列表
func QueryContentCatalogs() ([]commonmodel.Catalog, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryContentCatalogListRequest{}

	response := QueryContentCatalogListResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Catalogs, result
}

// QuerySingleCatalogRequest 查询分类列表请求
type QuerySingleCatalogRequest struct {
	ID int
}

// QuerySingleCatalogResponse 查询分类列表响应
type QuerySingleCatalogResponse struct {
	Found   bool
	Catalog commonmodel.CatalogDetail
}

// QuerySingleCatalog 查询指定分类
func QuerySingleCatalog(id int) (commonmodel.CatalogDetail, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QuerySingleCatalogRequest{ID: id}
	response := QuerySingleCatalogResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Catalog, result && response.Found
}

// CreateCatalogRequest 新建分类请求
type CreateCatalogRequest struct {
	Name    string
	Parent  []int
	Creater int
}

// CreateCatalogResponse 新建分类响应
type CreateCatalogResponse struct {
	Result  bool
	Catalog commonmodel.CatalogDetail
}

// CreateCatalog 新建分类
func CreateCatalog(name string, parent []int, creater int) (commonmodel.CatalogDetail, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := CreateCatalogRequest{Name: name, Parent: parent, Creater: creater}
	response := CreateCatalogResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Catalog, result && response.Result
}

// UpdateCatalogRequest 更新分类请求
type UpdateCatalogRequest struct {
	Catalog commonmodel.CatalogDetail
}

// UpdateCatalogResponse 更新分类响应
type UpdateCatalogResponse struct {
	Result  bool
	Catalog commonmodel.CatalogDetail
}

// UpdateCatalog 新建分类
func UpdateCatalog(catalog commonmodel.CatalogDetail) (commonmodel.CatalogDetail, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := UpdateCatalogRequest{Catalog: catalog}
	response := UpdateCatalogResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Catalog, result && response.Result
}

// DeleteCatalogRequest 删除分类请求
type DeleteCatalogRequest struct {
	ID int
}

// DeleteCatalogResponse 删除分类响应
type DeleteCatalogResponse struct {
	Result bool
}

// DeleteCatalog 删除分类
func DeleteCatalog(id int) bool {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := DeleteCatalogRequest{ID: id}
	response := DeleteCatalogResponse{}
	result := contentModule.Invoke(&request, &response)

	return result && response.Result
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

// QuerySingleLinkRequest 查询单个link
type QuerySingleLinkRequest struct {
	ID int
}

// QuerySingleLinkResponse 查询单个响应
type QuerySingleLinkResponse struct {
	Found bool
	Link  commonmodel.Link
}

// QuerySingleLink 查询单个Link
func QuerySingleLink(id int) (commonmodel.Link, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QuerySingleLinkRequest{ID: id}

	response := QuerySingleLinkResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Link, result && response.Found
}

// CreateLinkRequest 新建link
type CreateLinkRequest struct {
	Name    string
	URL     string
	Logo    string
	Catalog []int
	Creater int
}

// CreateLinkResponse 新建响应
type CreateLinkResponse struct {
	Result bool
	Link   commonmodel.Link
}

// CreateLink 新建Link
func CreateLink(name, url, logo string, catalog []int, creater int) (commonmodel.Link, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := CreateLinkRequest{Name: name, URL: url, Logo: logo, Catalog: catalog, Creater: creater}

	response := CreateLinkResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Link, result && response.Result
}

// UpdateLinkRequest 更新link
type UpdateLinkRequest struct {
	Link commonmodel.Link
}

// UpdateLinkResponse 更新响应
type UpdateLinkResponse struct {
	Result bool
	Link   commonmodel.Link
}

// UpdateLink 更新Link
func UpdateLink(lnk commonmodel.Link) (commonmodel.Link, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := UpdateLinkRequest{Link: lnk}

	response := UpdateLinkResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Link, result && response.Result
}

// DeleteLinkRequest 删除link
type DeleteLinkRequest struct {
	ID int
}

// DeleteLinkResponse 删除响应
type DeleteLinkResponse struct {
	Result bool
}

// DeleteLink 删除Link
func DeleteLink(id int) bool {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := DeleteLinkRequest{ID: id}
	response := DeleteLinkResponse{}

	result := contentModule.Invoke(&request, &response)

	return result && response.Result
}

// QueryContentMediaRequest 查询链接列表请求
type QueryContentMediaRequest struct {
}

// QueryContentMediaResponse 查询链接列表响应
type QueryContentMediaResponse struct {
	Medias []commonmodel.MediaDetail
}

// QueryContentMedias 查询链接列表
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

// QuerySingleMediaRequest 查询单个link
type QuerySingleMediaRequest struct {
	ID int
}

// QuerySingleMediaResponse 查询单个响应
type QuerySingleMediaResponse struct {
	Found bool
	Media commonmodel.MediaDetail
}

// QuerySingleMedia 查询单个Media
func QuerySingleMedia(id int) (commonmodel.MediaDetail, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QuerySingleMediaRequest{ID: id}

	response := QuerySingleMediaResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Media, result && response.Found
}

// CreateMediaRequest 新建link
type CreateMediaRequest struct {
	Name    string
	URL     string
	Type    string
	Desc    string
	Catalog []int
	Creater int
}

// CreateMediaResponse 新建响应
type CreateMediaResponse struct {
	Result bool
	Media  commonmodel.MediaDetail
}

// CreateMedia 新建Media
func CreateMedia(name, url, typeName, desc string, catalog []int, creater int) (commonmodel.MediaDetail, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := CreateMediaRequest{Name: name, URL: url, Type: typeName, Desc: desc, Catalog: catalog, Creater: creater}

	response := CreateMediaResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Media, result && response.Result
}

// UpdateMediaRequest 更新link
type UpdateMediaRequest struct {
	Media commonmodel.MediaDetail
}

// UpdateMediaResponse 更新响应
type UpdateMediaResponse struct {
	Result bool
	Media  commonmodel.MediaDetail
}

// UpdateMedia 更新Media
func UpdateMedia(lnk commonmodel.MediaDetail) (commonmodel.MediaDetail, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := UpdateMediaRequest{Media: lnk}

	response := UpdateMediaResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Media, result && response.Result
}

// DeleteMediaRequest 删除link
type DeleteMediaRequest struct {
	ID int
}

// DeleteMediaResponse 删除响应
type DeleteMediaResponse struct {
	Result bool
}

// DeleteMedia 删除Media
func DeleteMedia(id int) bool {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := DeleteMediaRequest{ID: id}
	response := DeleteMediaResponse{}

	result := contentModule.Invoke(&request, &response)

	return result && response.Result
}
