package handler

import (
	"log"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/common/daemon"
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/common/resource"
	common_const "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

// CreateContentHandler 新建ContentHandler
func CreateContentHandler() common.ContentHandler {
	dbhelper, _ := dbhelper.NewHelper()
	i := &impl{
		dbhelper:       dbhelper,
		articleHandler: articleActionHandler{dbhelper: dbhelper},
		catalogHandler: catalogActionHandler{dbhelper: dbhelper},
		linkHandler:    linkActionHandler{dbhelper: dbhelper},
		mediaHandler:   mediaActionHandler{dbhelper: dbhelper},
		commentHandler: commentActionHandler{dbhelper: dbhelper}}

	daemon.RegisterTimerHandler(i)

	return i
}

type impl struct {
	dbhelper       dbhelper.DBHelper
	articleHandler articleActionHandler
	catalogHandler catalogActionHandler
	linkHandler    linkActionHandler
	mediaHandler   mediaActionHandler
	commentHandler commentActionHandler
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

func (i *impl) GetArticleByCatalog(catalog model.CatalogUnit) []model.Summary {
	return i.articleHandler.findArticleByCatalog(catalog)
}

func (i *impl) CreateArticle(title, content, createDate string, catalog []model.CatalogUnit, author int) (model.Summary, bool) {
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

func (i *impl) GetCatalogByCatalog(id model.CatalogUnit) []model.Summary {
	return i.catalogHandler.findCatalogByCatalog(id)
}

func (i *impl) CreateCatalog(name, description, createDate string, parent []model.CatalogUnit, author int) (model.Summary, bool) {
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

func (i *impl) UpdateCatalog(catalogs []model.Catalog, parentCatalog model.CatalogUnit, description, updateDate string, updater int) ([]model.Summary, bool) {
	log.Print(parentCatalog)

	return i.catalogHandler.updateCatalog(catalogs, parentCatalog, description, updateDate, updater)
}

func (i *impl) QueryCatalogByName(name string, parentCatalog model.CatalogUnit) (model.CatalogDetail, bool) {
	return i.catalogHandler.queryCatalogByName(name, parentCatalog)
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

func (i *impl) GetLinkByCatalog(catalog model.CatalogUnit) []model.Summary {
	return i.linkHandler.findLinkByCatalog(catalog)
}

func (i *impl) CreateLink(name, desc, url, logo, createDate string, catalog []model.CatalogUnit, author int) (model.Summary, bool) {
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

func (i *impl) GetMediaByCatalog(catalog model.CatalogUnit) []model.Summary {
	return i.mediaHandler.findMediaByCatalog(catalog)
}

func (i *impl) CreateMedia(name, desc, fileToken, createDate string, catalog []model.CatalogUnit, expiration, author int) (model.Summary, bool) {
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

func (i *impl) GetCommentByCatalog(catalog model.CatalogUnit) []model.CommentDetail {
	return i.commentHandler.findCommentByCatalog(catalog)
}

func (i *impl) CreateComment(subject, content, createDate string, catalog []model.CatalogUnit, author int) (model.Summary, bool) {
	return i.commentHandler.createComment(subject, content, createDate, catalog, author)
}

func (i *impl) SaveComment(comment model.CommentDetail) (model.Summary, bool) {
	return i.commentHandler.saveComment(comment)
}

func (i *impl) DisableComment(id int) bool {
	return i.commentHandler.disableComment(id)
}

func (i *impl) DestroyComment(id int) bool {
	referenceRes := resource.QueryReferenceResource(i.dbhelper, id, model.COMMENT, "")
	if len(referenceRes) > 0 {
		return false
	}

	return i.commentHandler.destroyComment(id)
}

func (i *impl) GetSummaryByIDs(ids []model.CatalogUnit) []model.Summary {
	summaryList := []model.Summary{}
	articleIds := []int{}
	catalogIds := []int{}
	linkIds := []int{}
	mediaIds := []int{}
	for _, val := range ids {
		switch val.Type {
		case model.ARTICLE:
			articleIds = append(articleIds, val.ID)
		case model.CATALOG:
			catalogIds = append(catalogIds, val.ID)
		case model.LINK:
			linkIds = append(linkIds, val.ID)
		case model.MEDIA:
			mediaIds = append(mediaIds, val.ID)
		}
	}
	articles := resource.QueryResourceByIDs(i.dbhelper, articleIds, model.ARTICLE)
	for _, r := range articles {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}
	catalogs := resource.QueryResourceByIDs(i.dbhelper, catalogIds, model.CATALOG)
	for _, r := range catalogs {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}
	if util.ExistIntArray(common_const.SystemContentCatalog.ID, catalogIds) {
		summaryList = append(summaryList, *common_const.SystemContentCatalog.Summary())
	}

	links := resource.QueryResourceByIDs(i.dbhelper, linkIds, model.LINK)
	for _, r := range links {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}
	medias := resource.QueryResourceByIDs(i.dbhelper, mediaIds, model.MEDIA)
	for _, r := range medias {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(i.dbhelper, value.ID, value.Type)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, *r.CatalogUnit())
		}
	}

	return summaryList
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
		summary.Catalog = append(summary.Catalog, *r.CatalogUnit())
	}

	return summary, ok
}

func (i *impl) QuerySummaryContent(id int, summaryType string) []model.Summary {
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
			summary.Catalog = append(summary.Catalog, *r.CatalogUnit())
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
			summary.Catalog = append(summary.Catalog, *r.CatalogUnit())
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
		if v.RType() == model.COMMENT {
			continue
		}

		item := model.ContentUnit{Title: v.RName(), Type: v.RType(), CreateDate: v.RCreateDate()}

		resultList = append(resultList, item)
	}
	return resultList
}

func (i *impl) Handle() {
	i.mediaHandler.expirationCheck()
}
