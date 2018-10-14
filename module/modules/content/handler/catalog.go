package handler

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/content/dal"
	"muidea.com/magicCommon/def"
	"muidea.com/magicCommon/model"
)

type catalogActionHandler struct {
}

func (i *catalogActionHandler) getAllCatalog(filter *def.Filter) ([]model.Summary, int) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllCatalog(dbhelper, filter)
}

func (i *catalogActionHandler) getCatalogs(ids []int) []model.Catalog {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryCatalogs(dbhelper, ids)
}

func (i *catalogActionHandler) findCatalogByID(id int) (model.CatalogDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryCatalogByID(dbhelper, id)
}

func (i *catalogActionHandler) findCatalogByCatalog(catalog model.CatalogUnit, filter *def.Filter) ([]model.Summary, int) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryCatalogByCatalog(dbhelper, catalog, filter)
}

func (i *catalogActionHandler) createCatalog(name, description, createDate string, parent []model.CatalogUnit, author int) (model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.CreateCatalog(dbhelper, name, description, createDate, parent, author, false)
}

func (i *catalogActionHandler) saveCatalog(catalog model.CatalogDetail) (model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.SaveCatalog(dbhelper, catalog, false)
}

func (i *catalogActionHandler) destroyCatalog(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.DeleteCatalog(dbhelper, id)
}

func (i *catalogActionHandler) updateCatalog(catalogs []model.Catalog, parentCatalog model.CatalogUnit, description, updateDate string, updater int) ([]model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.UpdateCatalog(dbhelper, catalogs, parentCatalog, description, updateDate, updater)
}

func (i *catalogActionHandler) queryCatalogByName(name string, parentCatalog model.CatalogUnit) (model.CatalogDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryCatalogByName(dbhelper, name, parentCatalog)
}
