package bll

import (
    "magiccenter/util/modelhelper"
    "magiccenter/kernel/account/dal"
    "magiccenter/kernel/account/model"
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

func SaveGroup(id int, name string) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	group := model.Group{}
	group.Id = id
	group.Name = name
	
	return dal.SaveGroup(helper, group)	
}



