package common

import (
	"muidea.com/magicCenter/application/common/model"
)

// ContentHandler 内容处理Handler
type ContentHandler interface {
	GetAllArticle() []model.ArticleSummary
	GetArticleByID(id int) (model.Article, bool)
	GetArticleByCatalog(catalog int) []model.ArticleSummary
	CreateArticle(title, content, createDate string, catalog []int, author int) (model.ArticleSummary, bool)
	SaveArticle(article model.Article) (model.ArticleSummary, bool)
	DestroyArticle(id int) bool

	GetAllCatalog() []model.Catalog
	GetCatalogByID(id int) (model.CatalogDetail, bool)
	GetCatalogByParent(id int) []model.Catalog
	CreateCatalog(name string, parent []int, author int) (model.Catalog, bool)
	SaveCatalog(catalog model.CatalogDetail) (model.Catalog, bool)
	DestroyCatalog(id int) bool

	GetAllLink() []model.Link
	GetLinkByID(id int) (model.Link, bool)
	GetLinkByCatalog(catalog int) []model.Link
	CreateLink(name, url, logo string, catalog []int, author int) (model.Link, bool)
	SaveLink(link model.Link) (model.Link, bool)
	DestroyLink(id int) bool

	GetAllMedia() []model.MediaDetail
	GetMediaByID(id int) (model.MediaDetail, bool)
	GetMediaByCatalog(catalog int) []model.MediaDetail
	CreateMedia(name, url, desc string, catalog []int, author int) (model.MediaDetail, bool)
	SaveMedia(media model.MediaDetail) (model.MediaDetail, bool)
	DestroyMedia(id int) bool
}
