package bll

import (
	"magiccenter/kernel/modules/content/dal"
	"magiccenter/kernel/modules/content/model"
	"magiccenter/util/modelhelper"
)

// QueryAllCatalog 查询全部分类
func QueryAllCatalog() []model.Catalog {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllCatalog(helper)
}

// QueryAllCatalogDetail 查询全部分类详情
func QueryAllCatalogDetail() []model.CatalogDetail {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllCatalogDetail(helper)
}

// QueryCatalogByID 查询指定分类
func QueryCatalogByID(id int) (model.CatalogDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	catalog, result := dal.QueryCatalogByID(helper, id)
	return catalog, result
}

// QueryAvalibleParentCatalog 查询可用父类
func QueryAvalibleParentCatalog(id int) []model.Catalog {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAvalibleParentCatalog(helper, id)
}

// QuerySubCatalog 查询子类
func QuerySubCatalog(id int) []model.Catalog {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QuerySubCatalog(helper, id)
}

// DeleteCatalog 删除分类
func DeleteCatalog(id int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteCatalog(helper, id)
}

// SaveCatalog 保存分类
func SaveCatalog(id int, name string, uID int, parents []int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	catalog := model.CatalogDetail{}
	catalog.Id = id
	catalog.Name = name
	catalog.Creater = uID
	catalog.Parent = parents

	return dal.SaveCatalog(helper, catalog)
}
