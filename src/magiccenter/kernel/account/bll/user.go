package bll

import (
    "magiccenter/util/modelhelper"
    "magiccenter/kernel/account/dal"
    "magiccenter/kernel/account/model"
)

func QueryAllUser() []model.UserDetailView {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllUser(helper)
}

func QueryUserByAccount(account string) (model.UserDetailView,bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
		
	return dal.QueryUserByAccount(helper, account)
}

func VerifyUserByAccount(account,password string) (model.UserDetailView,bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
		
	return dal.VerifyUserByAccount(helper, account,password)
}


func QueryUserById(id int) (model.UserDetailView,bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
		
	return dal.QueryUserById(helper, id)	
}

func DeleteUser(id int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	return dal.DeleteUser(helper,id)
}

func SaveUser(id int, account, email string, groups []int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	user := model.UserDetailView{}
	user.Id = id
	user.Account = account
	user.Email = email
	user.Status = model.CREATE
	
	for _, g := range groups {
		group, found := dal.QueryGroupById(helper, g)
		if found {
			user.Groups = append(user.Groups, group)
		}
	}
	
	return dal.SaveUser(helper, user)
}

func UpdateUserDetail(id int, name,email string, status int, groups []int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	user, found := dal.QueryUserById(helper, id)
	if !found {
		return false
	}
	
	user.Name = name
	user.Email = email
	user.Status = status
	
	groupList := []model.Group{}
	for _, g := range groups {
		group, found := dal.QueryGroupById(helper, g)
		if found {
			groupList = append(groupList, group)
		}
	}
	user.Groups = groupList
	
	return dal.SaveUser(helper, user)
}










