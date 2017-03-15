package common

import (
	"muidea.com/magicCenter/application/common/model"
)

// ContentHandler 内容处理Handler
type ContentHandler interface {
	GetAllArticle() []model.Summary
	GetArticleByID(id int) (model.ArticleDetail, bool)
	GetArticleByCatalog(catalog int) []model.Summary
	CreateArticle(title, content, createDate string, catalog []int, author int) (model.Summary, bool)
	SaveArticle(article model.ArticleDetail) (model.Summary, bool)
	DestroyArticle(id int) bool

	GetAllCatalog() []model.Summary
	GetCatalogByID(id int) (model.CatalogDetail, bool)
	GetCatalogByCatalog(id int) []model.Summary
	CreateCatalog(name, description, createdate string, catalog []int, author int) (model.Summary, bool)
	SaveCatalog(catalog model.CatalogDetail) (model.Summary, bool)
	DestroyCatalog(id int) bool

	GetAllLink() []model.Summary
	GetLinkByID(id int) (model.LinkDetail, bool)
	GetLinkByCatalog(catalog int) []model.Summary
	CreateLink(name, url, logo, createdate string, catalog []int, author int) (model.Summary, bool)
	SaveLink(link model.LinkDetail) (model.Summary, bool)
	DestroyLink(id int) bool

	GetAllMedia() []model.Summary
	GetMediaByID(id int) (model.MediaDetail, bool)
	GetMediaByCatalog(catalog int) []model.Summary
	CreateMedia(name, url, desc, createdate string, catalog []int, author int) (model.Summary, bool)
	SaveMedia(media model.MediaDetail) (model.Summary, bool)
	DestroyMedia(id int) bool
}
