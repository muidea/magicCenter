package handler

import (
	"sync"
	"time"

	"github.com/muidea/magicCenter/common/dbhelper"
	"github.com/muidea/magicCenter/module/modules/content/dal"
	"github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/model"
)

type mediaActionHandler struct {
	routesLock         sync.RWMutex
	preCheckTimeStamp  *time.Time
	mediaExpirationMap map[int]time.Time
}

func (i *mediaActionHandler) getAllMedia(filter *def.Filter) ([]model.Summary, int) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllMedia(dbhelper, filter)
}

func (i *mediaActionHandler) getMedias(ids []int) []model.Media {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryMedias(dbhelper, ids)
}

func (i *mediaActionHandler) findMediaByID(id int) (model.MediaDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryMediaByID(dbhelper, id)
}

func (i *mediaActionHandler) findMediaByCatalog(catalog model.CatalogUnit, filter *def.Filter) ([]model.Summary, int) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryMediaByCatalog(dbhelper, catalog, filter)
}

func (i *mediaActionHandler) createMedia(name, desc, fileToken, createDate string, catalog []model.CatalogUnit, expiration, author int) (model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	result, ok := dal.CreateMedia(dbhelper, name, desc, fileToken, createDate, expiration, author, catalog)

	i.loadAllMediaExpiration()

	return result, ok
}

func (i *mediaActionHandler) batchCreateMedia(medias []model.MediaItem, createDate string, creater int) ([]model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	result, ok := dal.BatchCreateMedia(dbhelper, medias, createDate, creater)

	i.loadAllMediaExpiration()

	return result, ok
}

func (i *mediaActionHandler) saveMedia(media model.MediaDetail) (model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	result, ok := dal.SaveMedia(dbhelper, media)

	i.loadAllMediaExpiration()

	return result, ok
}

func (i *mediaActionHandler) destroyMedia(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	result := dal.DeleteMediaByID(dbhelper, id)

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

	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	for k, v := range allMediaExpirationMap {
		if current.Sub(v) >= 0 {
			dal.DeleteMediaByID(dbhelper, k)
		}
	}
}

func (i *mediaActionHandler) loadAllMediaExpiration() {
	i.routesLock.Lock()
	defer i.routesLock.Unlock()
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	i.mediaExpirationMap = dal.LoadMediaExpiration(dbhelper)
}
