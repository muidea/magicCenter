package bll

import (
	"log"
    "magiccenter/util/modelhelper"
    "magiccenter/kernel/content/dal"
    "magiccenter/kernel/content/model"
)

func QueryAllCatalog() []model.Catalog {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()		
	
	return dal.QueryAllCatalog(helper)
}

func QueryAllCatalogDetail() []model.CatalogDetail {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()		
	
	return dal.QueryAllCatalogDetail(helper)
}

func QueryCatalogById(id int) (model.CatalogDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()		
	
	catalog, result := dal.QueryCatalogById(helper, id)
	return catalog, result
}

func QueryAvalibleParentCatalog(id int) []model.Catalog {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()		

	return dal.QueryAvalibleParentCatalog(helper, id)	
}

func QuerySubCatalog(id int) []model.Catalog {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()		

	return dal.QuerySubCatalog(helper, id)	
}

func DeleteCatalog(id int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()		

	return dal.DeleteCatalog(helper, id)	
}

func SaveCatalog(id int, name string, uId int, parents []int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	catalog := model.CatalogDetail{}
	catalog.Id = id
	catalog.Name = name
	catalog.Creater.Id = uId
	
	for _, p := range parents {
		parent, found := dal.QueryCatalogById(helper, p)
		if found {
			ca := model.Catalog{}
			ca.Id = parent.Id
			ca.Name = parent.Name
			catalog.Parent = append(catalog.Parent, ca)
		} else {
			log.Printf("illegal catalog id, id:%d", p)
		}
	}
	
	return dal.SaveCatalog(helper, catalog)
}






