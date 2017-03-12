package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/content/dal"
)

type linkActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *linkActionHandler) getAllLink() []model.Link {
	return dal.QueryAllLink(i.dbhelper)
}

func (i *linkActionHandler) findLinkByID(id int) (model.Link, bool) {
	return dal.QueryLinkByID(i.dbhelper, id)
}

func (i *linkActionHandler) findLinkByCatalog(catalog int) []model.Link {
	return dal.QueryLinkByCatalog(i.dbhelper, catalog)
}

func (i *linkActionHandler) createLink(name, url, logo string, catalog []int, author int) (model.Link, bool) {
	return dal.CreateLink(i.dbhelper, name, url, logo, author, catalog)
}

func (i *linkActionHandler) saveLink(link model.Link) (model.Link, bool) {
	return dal.SaveLink(i.dbhelper, link)
}

func (i *linkActionHandler) destroyLink(id int) bool {
	return dal.DeleteLinkByID(i.dbhelper, id)
}
