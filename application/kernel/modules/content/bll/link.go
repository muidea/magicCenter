package bll

import (
	"magiccenter/common/model"
	"magiccenter/kernel/modules/content/dal"
	"magiccenter/system"
)

// QueryAllLink 查询全部Link
func QueryAllLink() []model.Link {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllLink(helper)
}

// QueryLinkByID 查询指定Link
func QueryLinkByID(id int) (model.Link, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryLinkByID(helper, id)
}

// DeleteLinkByID 删除Link
func DeleteLinkByID(id int) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteLinkByID(helper, id)
}

// QueryLinkByCatalog 查询指定分类的Link
func QueryLinkByCatalog(id int) []model.Link {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryLinkByCatalog(helper, id)
}

// QueryLinkByRang 查询指定范围的Link
func QueryLinkByRang(begin, offset int) []model.Link {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryLinkByRang(helper, begin, offset)
}

// CreateLink 新建Link
func CreateLink(name, url, logo string, uID int, catalogs []int) (model.Link, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.CreateLink(helper, name, url, logo, uID, catalogs)
}

// SaveLink 保存Link
func SaveLink(lnk model.Link) (model.Link, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.SaveLink(helper, lnk)
}
