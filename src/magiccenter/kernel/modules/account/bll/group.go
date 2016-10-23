package bll

import (
	"magiccenter/common/model"
	"magiccenter/kernel/modules/account/dal"
	"magiccenter/system"
)

// QueryAllGroup 查询所有分组
func QueryAllGroup() []model.Group {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllGroup(helper)
}

// QueryGroupByID 查询指定分组
func QueryGroupByID(id int) (model.Group, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryGroupByID(helper, id)
}

// QueryGroupByName 查询指定分组
func QueryGroupByName(name string) (model.Group, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryGroupByName(helper, name)
}

// DeleteGroup 删除分组
func DeleteGroup(id int) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteGroup(helper, id)
}

// SaveGroup 保存分组信息
func SaveGroup(id int, name, desc string) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	group := model.Group{}
	group.ID = id
	group.Name = name
	group.Description = desc

	return dal.SaveGroup(helper, group)
}
