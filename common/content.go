package common

import (
	"muidea.com/magicCommon/def"
	"muidea.com/magicCommon/model"
)

// ContentHandler 内容处理Handler
type ContentHandler interface {
	GetAllArticle(filter *def.Filter) ([]model.Summary, int)
	GetArticles(ids []int) []model.Article
	GetArticleByID(id int) (model.ArticleDetail, bool)
	GetArticleByCatalog(catalog model.CatalogUnit, filter *def.Filter) ([]model.Summary, int)
	CreateArticle(title, content, createDate string, catalog []model.CatalogUnit, author int) (model.Summary, bool)
	SaveArticle(article model.ArticleDetail) (model.Summary, bool)
	DestroyArticle(id int) bool

	GetAllCatalog(filter *def.Filter) ([]model.Summary, int)
	GetCatalogs(ids []int) []model.Catalog
	GetCatalogByID(id int) (model.CatalogDetail, bool)
	GetCatalogByCatalog(catalog model.CatalogUnit, filter *def.Filter) ([]model.Summary, int)
	CreateCatalog(name, description, createDate string, catalog []model.CatalogUnit, creater int) (model.Summary, bool)
	SaveCatalog(catalog model.CatalogDetail) (model.Summary, bool)
	DestroyCatalog(id int) bool

	// 更新Catalog，如果不存在，则新建一个
	UpdateCatalog(catalogs []model.Catalog, parentCatalog model.CatalogUnit, description, updateDate string, updater int) ([]model.Summary, bool)
	// 查询指定名称的Catalog
	QueryCatalogByName(name string, parentCatalog model.CatalogUnit) (model.CatalogDetail, bool)

	GetAllLink(filter *def.Filter) ([]model.Summary, int)
	GetLinks(ids []int) []model.Link
	GetLinkByID(id int) (model.LinkDetail, bool)
	GetLinkByCatalog(catalog model.CatalogUnit, filter *def.Filter) ([]model.Summary, int)
	CreateLink(name, desc, url, logo, createDate string, catalog []model.CatalogUnit, creater int) (model.Summary, bool)
	SaveLink(link model.LinkDetail) (model.Summary, bool)
	DestroyLink(id int) bool

	GetAllMedia(filter *def.Filter) ([]model.Summary, int)
	GetMedias(ids []int) []model.Media
	GetMediaByID(id int) (model.MediaDetail, bool)
	GetMediaByCatalog(catalog model.CatalogUnit, filter *def.Filter) ([]model.Summary, int)
	CreateMedia(name, desc, fileToken, createDate string, catalog []model.CatalogUnit, expiration, creater int) (model.Summary, bool)
	BatchCreateMedia(medias []model.MediaItem, createDate string, creater int) ([]model.Summary, bool)
	SaveMedia(media model.MediaDetail) (model.Summary, bool)
	DestroyMedia(id int) bool

	GetCommentByCatalog(catalog model.CatalogUnit, filter *def.PageFilter) ([]model.CommentDetail, int)
	CreateComment(subject, content, createDate string, catalog []model.CatalogUnit, creater int) (model.Summary, bool)
	SaveComment(comment model.CommentDetail) (model.Summary, bool)
	DestroyComment(id int) bool

	GetSummaryByIDs(ids []model.CatalogUnit) []model.Summary
	QuerySummaryByName(summaryName, summaryType string, catalog model.CatalogUnit) (model.Summary, bool)
	QuerySummaryContent(summary model.CatalogUnit, specialType string, filter *def.Filter) ([]model.Summary, int)
	GetSummaryByUser(uids []int, filter *def.Filter) ([]model.Summary, int)

	GetContentSummary() model.ContentSummary
	GetLastContent(count int, filter *def.PageFilter) ([]model.ContentUnit, int)
}
