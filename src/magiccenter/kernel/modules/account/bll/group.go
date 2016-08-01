package bll

import (
	"magiccenter/kernel/modules/account/dal"
	"magiccenter/kernel/modules/account/model"
	"magiccenter/util/modelhelper"
)

// QueryAllGroup 查询所有分组
func QueryAllGroup() []model.Group {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllGroup(helper)
}

// QueryGroupByID 查询指定分组
func QueryGroupByID(id int) (model.Group, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryGroupByID(helper, id)
}

// QueryGroupByName 查询指定分组
func QueryGroupByName(name string) (model.Group, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryGroupByName(helper, name)
}

// DeleteGroup 删除分组
func DeleteGroup(id int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteGroup(helper, id)
}

// SaveGroup 保存分组信息
func SaveGroup(id int, name string) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	group := model.Group{}
	group.ID = id
	group.Name = name

	return dal.SaveGroup(helper, group)
}
