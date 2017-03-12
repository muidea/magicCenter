package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/content/dal"
)

type catalogActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *catalogActionHandler) getAllCatalog() []model.Catalog {
	return dal.QueryAllCatalog(i.dbhelper)
}

func (i *catalogActionHandler) findCatalogByID(id int) (model.CatalogDetail, bool) {
	return dal.QueryCatalogByID(i.dbhelper, id)
}

func (i *catalogActionHandler) findCatalogByParent(id int) []model.Catalog {
	return dal.QuerySubCatalog(i.dbhelper, id)
}

func (i *catalogActionHandler) createCatalog(name string, parent []int, author int) (model.Catalog, bool) {
	return dal.CreateCatalog(i.dbhelper, name, parent, author)
}

func (i *catalogActionHandler) saveCatalog(catalog model.CatalogDetail) (model.Catalog, bool) {
	return dal.SaveCatalog(i.dbhelper, catalog)
}

func (i *catalogActionHandler) destroyCatalog(id int) bool {
	return dal.DeleteCatalog(i.dbhelper, id)
}
