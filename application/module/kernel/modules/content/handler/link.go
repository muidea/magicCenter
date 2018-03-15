package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/content/dal"
)

type linkActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *linkActionHandler) getAllLink() []model.Summary {
	return dal.QueryAllLink(i.dbhelper)
}

func (i *linkActionHandler) getLinks(ids []int) []model.Link {
	return dal.QueryLinks(i.dbhelper, ids)
}

func (i *linkActionHandler) findLinkByID(id int) (model.LinkDetail, bool) {
	return dal.QueryLinkByID(i.dbhelper, id)
}

func (i *linkActionHandler) findLinkByCatalog(catalog int) []model.Summary {
	return dal.QueryLinkByCatalog(i.dbhelper, catalog)
}

func (i *linkActionHandler) createLink(name, url, logo, createDate string, catalog []int, author int) (model.Summary, bool) {
	return dal.CreateLink(i.dbhelper, name, url, logo, createDate, author, catalog)
}

func (i *linkActionHandler) saveLink(link model.LinkDetail) (model.Summary, bool) {
	return dal.SaveLink(i.dbhelper, link)
}

func (i *linkActionHandler) destroyLink(id int) bool {
	return dal.DeleteLinkByID(i.dbhelper, id)
}
