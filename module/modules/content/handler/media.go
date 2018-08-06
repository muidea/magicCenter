package handler

import (
	"sync"
	"time"

	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/content/dal"
	"muidea.com/magicCommon/model"
)

type mediaActionHandler struct {
	dbhelper dbhelper.DBHelper

	routesLock         sync.RWMutex
	preCheckTimeStamp  *time.Time
	mediaExpirationMap map[int]time.Time
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

func (i *mediaActionHandler) findMediaByCatalog(catalog model.CatalogUnit) []model.Summary {
	return dal.QueryMediaByCatalog(i.dbhelper, catalog)
}

func (i *mediaActionHandler) createMedia(name, desc, fileToken, createDate string, catalog []model.CatalogUnit, expiration, author int) (model.Summary, bool) {
	result, ok := dal.CreateMedia(i.dbhelper, name, desc, fileToken, createDate, expiration, author, catalog)

	i.loadAllMediaExpiration()

	return result, ok
}

func (i *mediaActionHandler) batchCreateMedia(medias []model.MediaItem, createDate string, creater int) ([]model.Summary, bool) {
	result, ok := dal.BatchCreateMedia(i.dbhelper, medias, createDate, creater)

	i.loadAllMediaExpiration()

	return result, ok
}

func (i *mediaActionHandler) saveMedia(media model.MediaDetail) (model.Summary, bool) {
	result, ok := dal.SaveMedia(i.dbhelper, media)

	i.loadAllMediaExpiration()

	return result, ok
}

func (i *mediaActionHandler) destroyMedia(id int) bool {
	result := dal.DeleteMediaByID(i.dbhelper, id)

	i.routesLock.Lock()
	defer i.routesLock.Unlock()
	delete(i.mediaExpirationMap, id)

	return result
}

func (i *mediaActionHandler) expirationCheck() {

	current := time.Now()
	if i.preCheckTimeStamp == nil {
		i.preCheckTimeStamp = &current
		i.loadAllMediaExpiration()

		return
	}

	if current.Sub(*(i.preCheckTimeStamp)).Hours() < 1 {
		return
	}

	i.preCheckTimeStamp = &current

	allMediaExpirationMap := map[int]time.Time{}
	{
		i.routesLock.RLock()
		defer i.routesLock.RUnlock()

		for k, v := range i.mediaExpirationMap {
			allMediaExpirationMap[k] = v
		}
	}

	for k, v := range allMediaExpirationMap {
		if current.Sub(v) >= 0 {
			dal.DeleteMediaByID(i.dbhelper, k)
		}
	}
}

func (i *mediaActionHandler) loadAllMediaExpiration() {
	i.routesLock.Lock()
	defer i.routesLock.Unlock()

	i.mediaExpirationMap = dal.LoadMediaExpiration(i.dbhelper)
}
