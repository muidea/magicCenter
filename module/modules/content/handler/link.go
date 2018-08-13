package handler

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/content/dal"
	"muidea.com/magicCommon/model"
)

type linkActionHandler struct {
}

func (i *linkActionHandler) getAllLink() []model.Summary {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllLink(dbhelper)
}

func (i *linkActionHandler) getLinks(ids []int) []model.Link {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryLinks(dbhelper, ids)
}

func (i *linkActionHandler) findLinkByID(id int) (model.LinkDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryLinkByID(dbhelper, id)
}

func (i *linkActionHandler) findLinkByCatalog(catalog model.CatalogUnit) []model.Summary {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryLinkByCatalog(dbhelper, catalog)
}

func (i *linkActionHandler) createLink(name, desc, url, logo, createDate string, catalog []model.CatalogUnit, author int) (model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.CreateLink(dbhelper, name, desc, url, logo, createDate, author, catalog)
}

func (i *linkActionHandler) saveLink(link model.LinkDetail) (model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.SaveLink(dbhelper, link)
}

func (i *linkActionHandler) destroyLink(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.DeleteLinkByID(dbhelper, id)
}
