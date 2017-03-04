package bll

import (
	"magiccenter/common/model"
	"magiccenter/kernel/api/dal"
	"magiccenter/system"
)

// GetModulePage 查询指定Module的指定页面
func GetModulePage(owner, url string) (model.Page, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryPage(helper, owner, url)
}

// SaveModulePage 保存Page信息
func SaveModulePage(page model.Page) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.SavePage(helper, page)
}

// DeleteModulePage 删除Page
func DeleteModulePage(owner, url string) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeletePage(helper, owner, url)
}
