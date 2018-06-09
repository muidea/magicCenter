package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/module/modules/content/dal"
	"muidea.com/magicCommon/model"
)

type mediaActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *mediaActionHandler) getAllMedia() []model.Summary {
	return dal.QueryAllMedia(i.dbhelper)
}

func (i *mediaActionHandler) getMedias(ids []int) []model.Media {
	return dal.QueryMedias(i.dbhelper, ids)
}

func (i *mediaActionHandler) findMediaByID(id int) (model.MediaDetail, bool) {
	return dal.QueryMediaByID(i.dbhelper, id)
}

func (i *mediaActionHandler) findMediaByCatalog(catalog int) []model.Summary {
	return dal.QueryMediaByCatalog(i.dbhelper, catalog)
}

func (i *mediaActionHandler) createMedia(name, desc, url, createDate string, catalog []int, expiration, author int) (model.Summary, bool) {
	return dal.CreateMedia(i.dbhelper, name, desc, url, createDate, expiration, author, catalog)
}

func (i *mediaActionHandler) saveMedia(media model.MediaDetail) (model.Summary, bool) {
	return dal.SaveMedia(i.dbhelper, media)
}

func (i *mediaActionHandler) destroyMedia(id int) bool {
	return dal.DeleteMediaByID(i.dbhelper, id)
}
