package bll

import (
	"log"
    "magiccenter/util/modelhelper"
    "magiccenter/kernel/content/dal"
    "magiccenter/kernel/content/model"
)

func QueryAllLink() []model.LinkDetail {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.QueryAllLink(helper)
}

func QueryLinkById(id int) (model.LinkDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.QueryLinkById(helper, id)	
}

func DeleteLinkById(id int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.DeleteLinkById(helper, id)	
}

func QueryLinkByCatalog(id int) []model.LinkDetail {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.QueryLinkByCatalog(helper, id)
}

func QueryLinkByRang(begin, offset int) []model.LinkDetail {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.QueryLinkByRang(helper, begin, offset)
}

func SaveLink(id int, name, url, logo string, uId int, catalogs []int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	link := model.LinkDetail{}
	link.Id = id
	link.Name = name
	link.Url = url
	link.Logo = logo
	link.Creater.Id = uId
	
	for _, ca := range catalogs {
		catalog, found := dal.QueryCatalogById(helper, ca)
		if found {
			c := model.Catalog{}
			c.Id = catalog.Id
			c.Name = catalog.Name
			link.Catalog = append(link.Catalog, c)
		} else {
			log.Printf("illegal catalog id, id:%d", ca)
		}
	}		
	
	return dal.SaveLink(helper, link)
}






