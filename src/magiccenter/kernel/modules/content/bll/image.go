package bll

import (
	"log"
    "magiccenter/util/modelhelper"
    "magiccenter/kernel/content/dal"
    "magiccenter/kernel/content/model"
)

func QueryAllImage() []model.ImageDetail {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.QueryAllImage(helper)
}

func QueryImageById(id int) (model.ImageDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.QueryImageById(helper, id)	
}

func DeleteImageById(id int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.DeleteImageById(helper, id)	
}

func QueryImageByCatalog(id int) []model.ImageDetail {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.QueryImageByCatalog(helper, id)
}

func QueryImageByRang(begin, offset int) []model.ImageDetail {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.QueryImageByRang(helper, begin, offset)
}

func SaveImage(id int, name, url, desc string, uId int, catalogs []int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	image := model.ImageDetail{}
	image.Id = id
	image.Name = name
	image.Url = url
	image.Desc = desc
	image.Creater.Id = uId
	
	for _, ca := range catalogs {
		catalog, found := dal.QueryCatalogById(helper, ca)
		if found {
			c := model.Catalog{}
			c.Id = catalog.Id
			c.Name = catalog.Name
			image.Catalog = append(image.Catalog, c)
		} else {
			log.Printf("illegal catalog id, id:%d", ca)
		}
	}		
	
	return dal.SaveImage(helper, image)
}





