package bll

import (
	"magiccenter/kernel/modules/content/dal"
	"magiccenter/kernel/modules/content/model"
	"magiccenter/util/dbhelper"
)

// QueryAllLink 查询全部Link
func QueryAllLink() []model.Link {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllLink(helper)
}

// QueryLinkByID 查询指定Link
func QueryLinkByID(id int) (model.Link, bool) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryLinkByID(helper, id)
}

// DeleteLinkByID 删除Link
func DeleteLinkByID(id int) bool {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteLinkByID(helper, id)
}

// QueryLinkByCatalog 查询指定分类的Link
func QueryLinkByCatalog(id int) []model.Link {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryLinkByCatalog(helper, id)
}

// QueryLinkByRang 查询指定范围的Link
func QueryLinkByRang(begin, offset int) []model.Link {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryLinkByRang(helper, begin, offset)
}

// SaveLink 保存Link
func SaveLink(id int, name, url, logo string, uID int, catalogs []int) bool {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	link := model.Link{}
	link.ID = id
	link.Name = name
	link.URL = url
	link.Logo = logo
	link.Creater = uID
	link.Catalog = catalogs

	return dal.SaveLink(helper, link)
}
