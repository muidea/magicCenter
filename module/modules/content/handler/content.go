package handler

import (
	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/common/resource"
	"muidea.com/magicCommon/model"
)

// CreateContentHandler 新建ContentHandler
func CreateContentHandler() common.ContentHandler {
	dbhelper, _ := dbhelper.NewHelper()
	i := impl{
		dbhelper:       dbhelper,
		articleHandler: articleActionHandler{dbhelper: dbhelper},
		catalogHandler: catalogActionHandler{dbhelper: dbhelper},
		linkHandler:    linkActionHandler{dbhelper: dbhelper},
		mediaHandler:   mediaActionHandler{dbhelper: dbhelper}}

	return &i
}

type impl struct {
	dbhelper       dbhelper.DBHelper
	articleHandler articleActionHandler
	catalogHandler catalogActionHandler
	linkHandler    linkActionHandler
	mediaHandler   mediaActionHandler
}

func (i *impl) GetAllArticle() []model.Summary {
	return i.articleHandler.getAllArticleSummary()
}

func (i *impl) GetArticles(ids []int) []model.Article {
	return i.articleHandler.getArticles(ids)
}

func (i *impl) GetArticleByID(id int) (model.ArticleDetail, bool) {
	return i.articleHandler.findArticleByID(id)
}

func (i *impl) GetArticleByCatalog(catalog int) []model.Summary {
	return i.articleHandler.findArticleByCatalog(catalog)
}

func (i *impl) CreateArticle(title, content, createDate string, catalog []int, author int) (model.Summary, bool) {
	return i.articleHandler.createArticle(title, content, createDate, catalog, author)
}

func (i *impl) SaveArticle(article model.ArticleDetail) (model.Summary, bool) {
	return i.articleHandler.saveArticle(article)
}

func (i *impl) DestroyArticle(id int) bool {
	referenceRes := resource.QueryReferenceResource(i.dbhelper, id, model.ARTICLE, "")
	if len(referenceRes) > 0 {
		return false
	}

	return i.articleHandler.destroyArticle(id)
}

func (i *impl) GetAllCatalog() []model.Summary {
	return i.catalogHandler.getAllCatalog()
}

func (i *impl) GetCatalogs(ids []int) []model.Catalog {
	return i.catalogHandler.getCatalogs(ids)
}

func (i *impl) GetCatalogByID(id int) (model.CatalogDetail, bool) {
	return i.catalogHandler.findCatalogByID(id)
}

func (i *impl) GetCatalogByCatalog(id int) []model.Summary {
	return i.catalogHandler.findCatalogByCatalog(id)
}

func (i *impl) CreateCatalog(name, description, createDate string, parent []int, author int) (model.Summary, bool) {
	return i.catalogHandler.createCatalog(name, description, createDate, parent, author)
}

func (i *impl) SaveCatalog(catalog model.CatalogDetail) (model.Summary, bool) {
	return i.catalogHandler.saveCatalog(catalog)
}

func (i *impl) DestroyCatalog(id int) bool {
	referenceRes := resource.QueryReferenceResource(i.dbhelper, id, model.CATALOG, "")
	if len(referenceRes) > 0 {
		return false
	}

	return i.catalogHandler.destroyCatalog(id)
}

func (i *impl) UpdateCatalog(catalogs []model.Catalog, updateDate string, updater int) ([]model.Catalog, bool) {
	return i.catalogHandler.updateCatalog(catalogs, updateDate, updater)
}

func (i *impl) QueryCatalogByName(name string) (model.CatalogDetail, bool) {
	return i.catalogHandler.queryCatalogByName(name)
}

func (i *impl) GetAllLink() []model.Summary {
	return i.linkHandler.getAllLink()
}

func (i *impl) GetLinks(ids []int) []model.Link {
	return i.linkHandler.getLinks(ids)
}

func (i *impl) GetLinkByID(id int) (model.LinkDetail, bool) {
	return i.linkHandler.findLinkByID(id)
}

func (i *impl) GetLinkByCatalog(catalog int) []model.Summary {
	return i.linkHandler.findLinkByCatalog(catalog)
}

func (i *impl) CreateLink(name, desc, url, logo, createDate string, catalog []int, author int) (model.Summary, bool) {
	return i.linkHandler.createLink(name, desc, url, logo, createDate, catalog, author)
}

func (i *impl) SaveLink(link model.LinkDetail) (model.Summary, bool) {
	return i.linkHandler.saveLink(link)
}

func (i *impl) DestroyLink(id int) bool {
	referenceRes := resource.QueryReferenceResource(i.dbhelper, id, model.LINK, "")
	if len(referenceRes) > 0 {
		return false
	}

	return i.linkHandler.destroyLink(id)
}

func (i *impl) GetAllMedia() []model.Summary {
	return i.mediaHandler.getAllMedia()
}

func (i *impl) GetMedias(ids []int) []model.Media {
	return i.mediaHandler.getMedias(ids)
}

func (i *impl) GetMediaByID(id int) (model.MediaDetail, bool) {
	return i.mediaHandler.findMediaByID(id)
}

func (i *impl) GetMediaByCatalog(catalog int) []model.Summary {
	return i.mediaHandler.findMediaByCatalog(catalog)
}

func (i *impl) CreateMedia(name, desc, fileToken, createDate string, catalog []int, expiration, author int) (model.Summary, bool) {
	return i.mediaHandler.createMedia(name, desc, fileToken, createDate, catalog, expiration, author)
}

func (i *impl) BatchCreateMedia(medias []model.MediaItem, createDate string, creater int) ([]model.Summary, bool) {
	return i.mediaHandler.batchCreateMedia(medias, createDate, creater)
}

func (i *impl) SaveMedia(media model.MediaDetail) (model.Summary, bool) {
	return i.mediaHandler.saveMedia(media)
}

func (i *impl) DestroyMedia(id int) bool {
	referenceRes := resource.QueryReferenceResource(i.dbhelper, id, model.MEDIA, "")
	if len(referenceRes) > 0 {
		return false
	}

	return i.mediaHandler.destroyMedia(id)
}

func (i *impl) QuerySummaryByName(name, summaryType string) (model.Summary, bool) {
	summary := model.Summary{}
	res, ok := resource.QueryResourceByName(i.dbhelper, name, summaryType)
	if !ok {
		return summary, ok
	}

	summary.ID = res.RId()
	summary.Name = res.RName()
	summary.Description = res.RDescription()
	summary.Type = res.RType()
	summary.CreateDate = res.RCreateDate()
	summary.Creater = res.ROwner()

	ress := resource.QueryRelativeResource(i.dbhelper, res.RId(), res.RType())
	for _, r := range ress {
		summary.Catalog = append(summary.Catalog, r.RId())
	}

	return summary, ok
}

func (i *impl) GetSummaryContent(id int, summaryType string) []model.Summary {
	summaryList := []model.Summary{}
	resList := resource.QueryReferenceResource(i.dbhelper, id, summaryType, "")
	for _, r := range resList {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(i.dbhelper, value.ID, value.Type)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return summaryList
}

func (i *impl) GetSummaryByUser(uids []int) []model.Summary {
	summaryList := []model.Summary{}
	resList := resource.QueryResourceByUser(i.dbhelper, uids)
	for _, r := range resList {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(i.dbhelper, value.ID, value.Type)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return summaryList
}

func (i *impl) GetContentSummary() model.ContentSummary {
	result := model.ContentSummary{}

	articleCount := len(i.articleHandler.getAllArticleSummary())
	articleItem := model.UnitSummary{Name: "文章", Type: "article", Count: articleCount}
	result = append(result, articleItem)

	catalogCount := len(i.catalogHandler.getAllCatalog())
	catalogItem := model.UnitSummary{Name: "分类", Type: "catalog", Count: catalogCount}
	result = append(result, catalogItem)

	linkCount := len(i.linkHandler.getAllLink())
	linkItem := model.UnitSummary{Name: "链接", Type: "link", Count: linkCount}
	result = append(result, linkItem)

	mediaCount := len(i.mediaHandler.getAllMedia())
	mediaItem := model.UnitSummary{Name: "文件", Type: "media", Count: mediaCount}
	result = append(result, mediaItem)

	return result
}

func (i *impl) GetLastContent(count int) []model.ContentUnit {
	resultList := []model.ContentUnit{}
	res := resource.GetLastResource(i.dbhelper, count)
	for _, v := range res {
		item := model.ContentUnit{Title: v.RName(), Type: v.RType(), CreateDate: v.RCreateDate()}

		resultList = append(resultList, item)
	}
	return resultList
}
