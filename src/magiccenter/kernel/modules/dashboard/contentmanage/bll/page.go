package bll

import (
	"magiccenter/common/model"
	"magiccenter/kernel/modules/dashboard/contentmanage/dal"
	"magiccenter/util/dbhelper"
)

// GetModulePage 查询指定Module的指定页面
func GetModulePage(owner, url string) (model.Page, bool) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryPage(helper, owner, url)
}

// SaveModulePage 保存Page信息
func SaveModulePage(page model.Page) bool {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.SavePage(helper, page)
}

// DeleteModulePage 删除Page
func DeleteModulePage(owner, url string) bool {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeletePage(helper, owner, url)
}
