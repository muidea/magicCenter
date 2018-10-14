package handler

import (
	"log"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/common/daemon"
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/common/resource"
	common_const "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/def"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

// CreateContentHandler 新建ContentHandler
func CreateContentHandler() common.ContentHandler {
	i := &impl{
		articleHandler: articleActionHandler{},
		catalogHandler: catalogActionHandler{},
		linkHandler:    linkActionHandler{},
		mediaHandler:   mediaActionHandler{},
		commentHandler: commentActionHandler{}}

	daemon.RegisterTimerHandler(i)

	return i
}

type impl struct {
	articleHandler articleActionHandler
	catalogHandler catalogActionHandler
	linkHandler    linkActionHandler
	mediaHandler   mediaActionHandler
	commentHandler commentActionHandler
}

func (i *impl) GetAllArticle(filter *def.Filter) ([]model.Summary, int) {
	return i.articleHandler.getAllArticleSummary(filter)
}

func (i *impl) GetArticles(ids []int) []model.Article {
	return i.articleHandler.getArticles(ids)
}

func (i *impl) GetArticleByID(id int) (model.ArticleDetail, bool) {
	return i.articleHandler.findArticleByID(id)
}

func (i *impl) GetArticleByCatalog(catalog model.CatalogUnit, filter *def.Filter) ([]model.Summary, int) {
	return i.articleHandler.findArticleByCatalog(catalog, filter)
}

func (i *impl) CreateArticle(title, content, createDate string, catalog []model.CatalogUnit, author int) (model.Summary, bool) {
	return i.articleHandler.createArticle(title, content, createDate, catalog, author)
}

func (i *impl) SaveArticle(article model.ArticleDetail) (model.Summary, bool) {
	return i.articleHandler.saveArticle(article)
}

func (i *impl) DestroyArticle(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	_, resCount := resource.QueryReferenceResource(dbhelper, id, model.ARTICLE, "", nil)
	if resCount > 0 {
		return false
	}

	return i.articleHandler.destroyArticle(id)
}

func (i *impl) GetAllCatalog(pageFilter *def.PageFilter) ([]model.Summary, int) {
	return i.catalogHandler.getAllCatalog(pageFilter)
}

func (i *impl) GetCatalogs(ids []int) []model.Catalog {
	return i.catalogHandler.getCatalogs(ids)
}

func (i *impl) GetCatalogByID(id int) (model.CatalogDetail, bool) {
	return i.catalogHandler.findCatalogByID(id)
}

func (i *impl) GetCatalogByCatalog(id model.CatalogUnit, pageFilter *def.PageFilter) ([]model.Summary, int) {
	return i.catalogHandler.findCatalogByCatalog(id, pageFilter)
}

func (i *impl) CreateCatalog(name, description, createDate string, parent []model.CatalogUnit, author int) (model.Summary, bool) {
	return i.catalogHandler.createCatalog(name, description, createDate, parent, author)
}

func (i *impl) SaveCatalog(catalog model.CatalogDetail) (model.Summary, bool) {
	return i.catalogHandler.saveCatalog(catalog)
}

func (i *impl) DestroyCatalog(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	_, resCount := resource.QueryReferenceResource(dbhelper, id, model.CATALOG, "", nil)
	if resCount > 0 {
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

func (i *impl) GetAllLink(pageFilter *def.PageFilter) ([]model.Summary, int) {
	return i.linkHandler.getAllLink(pageFilter)
}

func (i *impl) GetLinks(ids []int) []model.Link {
	return i.linkHandler.getLinks(ids)
}

func (i *impl) GetLinkByID(id int) (model.LinkDetail, bool) {
	return i.linkHandler.findLinkByID(id)
}

func (i *impl) GetLinkByCatalog(catalog model.CatalogUnit, pageFilter *def.PageFilter) ([]model.Summary, int) {
	return i.linkHandler.findLinkByCatalog(catalog, pageFilter)
}

func (i *impl) CreateLink(name, desc, url, logo, createDate string, catalog []model.CatalogUnit, author int) (model.Summary, bool) {
	return i.linkHandler.createLink(name, desc, url, logo, createDate, catalog, author)
}

func (i *impl) SaveLink(link model.LinkDetail) (model.Summary, bool) {
	return i.linkHandler.saveLink(link)
}

func (i *impl) DestroyLink(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	_, resCount := resource.QueryReferenceResource(dbhelper, id, model.LINK, "", nil)
	if resCount > 0 {
		return false
	}

	return i.linkHandler.destroyLink(id)
}

func (i *impl) GetAllMedia(pageFilter *def.PageFilter) ([]model.Summary, int) {
	return i.mediaHandler.getAllMedia(pageFilter)
}

func (i *impl) GetMedias(ids []int) []model.Media {
	return i.mediaHandler.getMedias(ids)
}

func (i *impl) GetMediaByID(id int) (model.MediaDetail, bool) {
	return i.mediaHandler.findMediaByID(id)
}

func (i *impl) GetMediaByCatalog(catalog model.CatalogUnit, pageFilter *def.PageFilter) ([]model.Summary, int) {
	return i.mediaHandler.findMediaByCatalog(catalog, pageFilter)
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
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	_, resCount := resource.QueryReferenceResource(dbhelper, id, model.MEDIA, "", nil)
	if resCount > 0 {
		return false
	}

	return i.mediaHandler.destroyMedia(id)
}

func (i *impl) GetCommentByCatalog(catalog model.CatalogUnit, pageFilter *def.PageFilter) ([]model.CommentDetail, int) {
	return i.commentHandler.findCommentByCatalog(catalog, pageFilter)
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
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	_, resCount := resource.QueryReferenceResource(dbhelper, id, model.COMMENT, "", nil)
	if resCount > 0 {
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

	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	articles, _ := resource.QueryResourceByIDs(dbhelper, articleIds, model.ARTICLE, nil)
	for _, r := range articles {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}
	catalogs, _ := resource.QueryResourceByIDs(dbhelper, catalogIds, model.CATALOG, nil)
	for _, r := range catalogs {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}
	if util.ExistIntArray(common_const.SystemContentCatalog.ID, catalogIds) {
		summaryList = append(summaryList, *common_const.SystemContentCatalog.Summary())
	}

	links, _ := resource.QueryResourceByIDs(dbhelper, linkIds, model.LINK, nil)
	for _, r := range links {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}
	medias, _ := resource.QueryResourceByIDs(dbhelper, mediaIds, model.MEDIA, nil)
	for _, r := range medias {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress, _ := resource.QueryRelativeResource(dbhelper, value.ID, value.Type, nil)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, *r.CatalogUnit())
		}
	}

	return summaryList
}

func (i *impl) QuerySummaryByName(summaryName, summaryType string, catalog model.CatalogUnit) (model.Summary, bool) {
	summary := model.Summary{}

	var res resource.Resource
	found := false
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	ress, _ := resource.QueryResourceByName(dbhelper, summaryName, summaryType, nil)
	for _, val := range ress {
		subRes := val.Relative()
		for _, sv := range subRes {
			if sv.RId() == catalog.ID && sv.RType() == catalog.Type {
				res = val
				found = true
				break
			}
		}
		if len(subRes) == 0 && common_const.IsSystemContentCatalog(catalog) {
			res = val
			found = true
		}

		if found {
			break
		}
	}

	if !found {
		return summary, found
	}

	summary.ID = res.RId()
	summary.Name = res.RName()
	summary.Description = res.RDescription()
	summary.Type = res.RType()
	summary.CreateDate = res.RCreateDate()
	summary.Creater = res.ROwner()
	for _, r := range res.Relative() {
		summary.Catalog = append(summary.Catalog, *r.CatalogUnit())
	}

	return summary, found
}

func (i *impl) QuerySummaryContent(summary model.CatalogUnit, filter *def.Filter) ([]model.Summary, int) {
	summaryList := []model.Summary{}
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	resList, resCount := resource.QueryReferenceResource(dbhelper, summary.ID, summary.Type, "", filter)
	for _, r := range resList {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		for _, v := range r.Relative() {
			summary.Catalog = append(summary.Catalog, *v.CatalogUnit())
		}

		summaryList = append(summaryList, summary)
	}

	return summaryList, resCount
}

func (i *impl) GetSummaryByUser(uids []int, pageFilter *def.PageFilter) ([]model.Summary, int) {
	summaryList := []model.Summary{}
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	resList, resCount := resource.QueryResourceByUser(dbhelper, uids, pageFilter)
	for _, r := range resList {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		for _, v := range r.Relative() {
			summary.Catalog = append(summary.Catalog, *v.CatalogUnit())
		}

		summaryList = append(summaryList, summary)
	}

	return summaryList, resCount
}

func (i *impl) GetContentSummary() model.ContentSummary {
	result := model.ContentSummary{}

	_, articleCount := i.articleHandler.getAllArticleSummary(nil)
	articleItem := model.UnitSummary{Name: "文章", Type: "article", Count: articleCount}
	result = append(result, articleItem)

	_, catalogCount := i.catalogHandler.getAllCatalog(nil)
	catalogItem := model.UnitSummary{Name: "分类", Type: "catalog", Count: catalogCount}
	result = append(result, catalogItem)

	_, linkCount := i.linkHandler.getAllLink(nil)
	linkItem := model.UnitSummary{Name: "链接", Type: "link", Count: linkCount}
	result = append(result, linkItem)

	_, mediaCount := i.mediaHandler.getAllMedia(nil)
	mediaItem := model.UnitSummary{Name: "文件", Type: "media", Count: mediaCount}
	result = append(result, mediaItem)

	return result
}

func (i *impl) GetLastContent(count int, pageFilter *def.PageFilter) ([]model.ContentUnit, int) {
	resultList := []model.ContentUnit{}
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	res, resCount := resource.GetLastResource(dbhelper, count, pageFilter)
	for _, v := range res {
		if v.RType() == model.COMMENT {
			continue
		}

		item := model.ContentUnit{Title: v.RName(), Type: v.RType(), CreateDate: v.RCreateDate()}

		resultList = append(resultList, item)
	}
	return resultList, resCount
}

func (i *impl) Handle() {
	i.mediaHandler.expirationCheck()
}
