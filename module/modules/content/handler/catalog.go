package handler

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/content/dal"
	"muidea.com/magicCommon/model"
)

type catalogActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *catalogActionHandler) getAllCatalog() []model.Summary {
	return dal.QueryAllCatalog(i.dbhelper)
}

func (i *catalogActionHandler) getCatalogs(ids []int) []model.Catalog {
	return dal.QueryCatalogs(i.dbhelper, ids)
}

func (i *catalogActionHandler) findCatalogByID(id int) (model.CatalogDetail, bool) {
	return dal.QueryCatalogByID(i.dbhelper, id)
}

func (i *catalogActionHandler) findCatalogByCatalog(id int) []model.Summary {
	return dal.QueryCatalogByCatalog(i.dbhelper, id)
}

func (i *catalogActionHandler) createCatalog(name, description, createDate string, parent []int, author int) (model.Summary, bool) {
	return dal.CreateCatalog(i.dbhelper, name, description, createDate, parent, author, false)
}

func (i *catalogActionHandler) saveCatalog(catalog model.CatalogDetail) (model.Summary, bool) {
	return dal.SaveCatalog(i.dbhelper, catalog, false)
}

func (i *catalogActionHandler) destroyCatalog(id int) bool {
	return dal.DeleteCatalog(i.dbhelper, id)
}

func (i *catalogActionHandler) updateCatalog(catalogs []model.Catalog, parentCatalog int, updateDate string, updater int) ([]model.Catalog, bool) {
	return dal.UpdateCatalog(i.dbhelper, catalogs, parentCatalog, updateDate, updater)
}

func (i *catalogActionHandler) queryCatalogByName(name string, parentCatalog int) (model.CatalogDetail, bool) {
	return dal.QueryCatalogByName(i.dbhelper, name, parentCatalog)
}
