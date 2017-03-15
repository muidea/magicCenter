package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/content/dal"
)

type mediaActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *mediaActionHandler) getAllMedia() []model.Summary {
	return dal.QueryAllMedia(i.dbhelper)
}

func (i *mediaActionHandler) findMediaByID(id int) (model.MediaDetail, bool) {
	return dal.QueryMediaByID(i.dbhelper, id)
}

func (i *mediaActionHandler) findMediaByCatalog(catalog int) []model.Summary {
	return dal.QueryMediaByCatalog(i.dbhelper, catalog)
}

func (i *mediaActionHandler) createMedia(name, url, desc, createdate string, catalog []int, author int) (model.Summary, bool) {
	return dal.CreateMedia(i.dbhelper, name, url, desc, createdate, author, catalog)
}

func (i *mediaActionHandler) saveMedia(media model.MediaDetail) (model.Summary, bool) {
	return dal.SaveMedia(i.dbhelper, media)
}

func (i *mediaActionHandler) destroyMedia(id int) bool {
	return dal.DeleteMediaByID(i.dbhelper, id)
}
