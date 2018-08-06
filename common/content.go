package common

import (
	"muidea.com/magicCommon/model"
)

// ContentHandler 内容处理Handler
type ContentHandler interface {
	GetAllArticle() []model.Summary
	GetArticles(ids []int) []model.Article
	GetArticleByID(id int) (model.ArticleDetail, bool)
	GetArticleByCatalog(catalog model.CatalogUnit) []model.Summary
	CreateArticle(title, content, createDate string, catalog []model.CatalogUnit, author int) (model.Summary, bool)
	SaveArticle(article model.ArticleDetail) (model.Summary, bool)
	DestroyArticle(id int) bool

	GetAllCatalog() []model.Summary
	GetCatalogs(ids []int) []model.Catalog
	GetCatalogByID(id int) (model.CatalogDetail, bool)
	GetCatalogByCatalog(catalog model.CatalogUnit) []model.Summary
	CreateCatalog(name, description, createDate string, catalog []model.CatalogUnit, creater int) (model.Summary, bool)
	SaveCatalog(catalog model.CatalogDetail) (model.Summary, bool)
	DestroyCatalog(id int) bool

	// 更新Catalog，如果不存在，则新建一个
	UpdateCatalog(catalogs []model.Catalog, parentCatalog model.CatalogUnit, description, updateDate string, updater int) ([]model.Summary, bool)
	// 查询指定名称的Catalog
	QueryCatalogByName(name string, parentCatalog model.CatalogUnit) (model.CatalogDetail, bool)

	GetAllLink() []model.Summary
	GetLinks(ids []int) []model.Link
	GetLinkByID(id int) (model.LinkDetail, bool)
	GetLinkByCatalog(catalog model.CatalogUnit) []model.Summary
	CreateLink(name, desc, url, logo, createDate string, catalog []model.CatalogUnit, creater int) (model.Summary, bool)
	SaveLink(link model.LinkDetail) (model.Summary, bool)
	DestroyLink(id int) bool

	GetAllMedia() []model.Summary
	GetMedias(ids []int) []model.Media
	GetMediaByID(id int) (model.MediaDetail, bool)
	GetMediaByCatalog(catalog model.CatalogUnit) []model.Summary
	CreateMedia(name, desc, fileToken, createDate string, catalog []model.CatalogUnit, expiration, creater int) (model.Summary, bool)
	BatchCreateMedia(medias []model.MediaItem, createDate string, creater int) ([]model.Summary, bool)
	SaveMedia(media model.MediaDetail) (model.Summary, bool)
	DestroyMedia(id int) bool

	GetCommentByCatalog(catalog model.CatalogUnit) []model.CommentDetail
	CreateComment(subject, content, createDate string, catalog []model.CatalogUnit, creater int) (model.Summary, bool)
	SaveComment(comment model.CommentDetail) (model.Summary, bool)
	DestroyComment(id int) bool

	GetSummaryByIDs(ids []model.CatalogUnit) []model.Summary
	QuerySummaryByName(name, summaryType string) (model.Summary, bool)
	QuerySummaryContent(id int, summaryType string) []model.Summary
	GetSummaryByUser(uids []int) []model.Summary

	GetContentSummary() model.ContentSummary
	GetLastContent(count int) []model.ContentUnit
}
