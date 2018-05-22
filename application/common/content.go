package common

import (
	"muidea.com/magicCommon/model"
)

// ContentHandler 内容处理Handler
type ContentHandler interface {
	GetAllArticle() []model.Summary
	GetArticles(ids []int) []model.Article
	GetArticleByID(id int) (model.ArticleDetail, bool)
	GetArticleByCatalog(catalog int) []model.Summary
	CreateArticle(title, content, createDate string, catalog []int, author int) (model.Summary, bool)
	SaveArticle(article model.ArticleDetail) (model.Summary, bool)
	DestroyArticle(id int) bool

	GetAllCatalog() []model.Summary
	GetCatalogs(ids []int) []model.Catalog
	GetCatalogByID(id int) (model.CatalogDetail, bool)
	GetCatalogByCatalog(id int) []model.Summary
	CreateCatalog(name, description, createDate string, catalog []int, creater int) (model.Summary, bool)
	SaveCatalog(catalog model.CatalogDetail) (model.Summary, bool)
	DestroyCatalog(id int) bool
	// 更新Catalog，如果不存在，则新建一个
	UpdateCatalog(catalogs []model.Catalog, updateDate string, updater int) ([]model.Catalog, bool)
	// 查询指定名称的Catalog
	QueryCatalogByName(name string) (model.CatalogDetail, bool)

	GetAllLink() []model.Summary
	GetLinks(ids []int) []model.Link
	GetLinkByID(id int) (model.LinkDetail, bool)
	GetLinkByCatalog(catalog int) []model.Summary
	CreateLink(name, desc, url, logo, createDate string, catalog []int, creater int) (model.Summary, bool)
	SaveLink(link model.LinkDetail) (model.Summary, bool)
	DestroyLink(id int) bool

	GetAllMedia() []model.Summary
	GetMedias(ids []int) []model.Media
	GetMediaByID(id int) (model.MediaDetail, bool)
	GetMediaByCatalog(catalog int) []model.Summary
	CreateMedia(name, desc, url, createDate string, catalog []int, expiration, creater int) (model.Summary, bool)
	SaveMedia(media model.MediaDetail) (model.Summary, bool)
	DestroyMedia(id int) bool

	GetSummaryByCatalog(id int) []model.Summary

	GetContentSummary() model.ContentSummary
	GetLastContent(count int) []model.ContentUnit
}
