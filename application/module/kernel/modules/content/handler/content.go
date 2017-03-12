package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
)

// CreateContentHandler 新建ContentHandler
func CreateContentHandler() common.ContentHandler {
	dbhelper, _ := dbhelper.NewHelper()
	i := impl{
		articleHandler: articleActionHandler{dbhelper: dbhelper},
		catalogHandler: catalogActionHandler{dbhelper: dbhelper},
		linkHandler:    linkActionHandler{dbhelper: dbhelper},
		mediaHandler:   mediaActionHandler{dbhelper: dbhelper}}

	return &i
}

type impl struct {
	articleHandler articleActionHandler
	catalogHandler catalogActionHandler
	linkHandler    linkActionHandler
	mediaHandler   mediaActionHandler
}

func (i *impl) GetAllArticle() []model.ArticleSummary {
	return i.articleHandler.getAllArticleSummary()
}

func (i *impl) GetArticleByID(id int) (model.Article, bool) {
	return i.articleHandler.findArticleByID(id)
}

func (i *impl) GetArticleByCatalog(catalog int) []model.ArticleSummary {
	return i.articleHandler.findArticleByCatalog(catalog)
}

func (i *impl) CreateArticle(title, content, createDate string, catalog []int, author int) (model.ArticleSummary, bool) {
	return i.articleHandler.createArticle(title, content, createDate, catalog, author)
}

func (i *impl) SaveArticle(article model.Article) (model.ArticleSummary, bool) {
	return i.articleHandler.saveArticle(article)
}

func (i *impl) DestroyArticle(id int) bool {
	return i.articleHandler.destroyArticle(id)
}

func (i *impl) GetAllCatalog() []model.Catalog {
	return i.catalogHandler.getAllCatalog()
}

func (i *impl) GetCatalogByID(id int) (model.CatalogDetail, bool) {
	return i.catalogHandler.findCatalogByID(id)
}

func (i *impl) GetCatalogByParent(id int) []model.Catalog {
	return i.catalogHandler.findCatalogByParent(id)
}

func (i *impl) CreateCatalog(name string, parent []int, author int) (model.Catalog, bool) {
	return i.catalogHandler.createCatalog(name, parent, author)
}

func (i *impl) SaveCatalog(catalog model.CatalogDetail) (model.Catalog, bool) {
	return i.catalogHandler.saveCatalog(catalog)
}

func (i *impl) DestroyCatalog(id int) bool {
	return i.catalogHandler.destroyCatalog(id)
}

func (i *impl) GetAllLink() []model.Link {
	return i.linkHandler.getAllLink()
}

func (i *impl) GetLinkByID(id int) (model.Link, bool) {
	return i.linkHandler.findLinkByID(id)
}

func (i *impl) GetLinkByCatalog(catalog int) []model.Link {
	return i.linkHandler.findLinkByCatalog(catalog)
}

func (i *impl) CreateLink(name, url, logo string, catalog []int, author int) (model.Link, bool) {
	return i.linkHandler.createLink(name, url, logo, catalog, author)
}

func (i *impl) SaveLink(link model.Link) (model.Link, bool) {
	return i.linkHandler.saveLink(link)
}

func (i *impl) DestroyLink(id int) bool {
	return i.linkHandler.destroyLink(id)
}

func (i *impl) GetAllMedia() []model.MediaDetail {
	return i.mediaHandler.getAllMedia()
}

func (i *impl) GetMediaByID(id int) (model.MediaDetail, bool) {
	return i.mediaHandler.findMediaByID(id)
}

func (i *impl) GetMediaByCatalog(catalog int) []model.MediaDetail {
	return i.mediaHandler.findMediaByCatalog(catalog)
}

func (i *impl) CreateMedia(name, url, desc string, catalog []int, author int) (model.MediaDetail, bool) {
	return i.mediaHandler.createMedia(name, url, desc, catalog, author)
}

func (i *impl) SaveMedia(media model.MediaDetail) (model.MediaDetail, bool) {
	return i.mediaHandler.saveMedia(media)
}

func (i *impl) DestroyMedia(id int) bool {
	return i.mediaHandler.destroyMedia(id)
}
