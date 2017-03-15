package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/content/dal"
)

type catalogActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *catalogActionHandler) getAllCatalog() []model.Summary {
	return dal.QueryAllCatalog(i.dbhelper)
}

func (i *catalogActionHandler) findCatalogByID(id int) (model.CatalogDetail, bool) {
	return dal.QueryCatalogByID(i.dbhelper, id)
}

func (i *catalogActionHandler) findCatalogByCatalog(id int) []model.Summary {
	return dal.QueryCatalogByCatalog(i.dbhelper, id)
}

func (i *catalogActionHandler) createCatalog(name, description, createdate string, parent []int, author int) (model.Summary, bool) {
	return dal.CreateCatalog(i.dbhelper, name, description, createdate, parent, author)
}

func (i *catalogActionHandler) saveCatalog(catalog model.CatalogDetail) (model.Summary, bool) {
	return dal.SaveCatalog(i.dbhelper, catalog)
}

func (i *catalogActionHandler) destroyCatalog(id int) bool {
	return dal.DeleteCatalog(i.dbhelper, id)
}
