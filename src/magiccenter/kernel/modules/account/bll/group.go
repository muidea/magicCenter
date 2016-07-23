package bll

import (
	"magiccenter/kernel/modules/account/dal"
	"magiccenter/kernel/modules/account/model"
	"magiccenter/util/modelhelper"
)

func QueryAllGroupInfo() []model.GroupInfo {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllGroupInfo(helper)
}

func QueryGroupById(id int) (model.Group, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryGroupById(helper, id)
}

func DeleteGroup(id int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteGroup(helper, id)
}

func SaveGroup(id int, name string, creater int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	group := model.Group{}
	group.Id = id
	group.Name = name
	group.Creater.Id = creater

	return dal.SaveGroup(helper, group)
}
