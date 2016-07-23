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

func QueryUserByAccount(account string) (model.UserDetail,bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
		
	return dal.QueryUserByAccount(helper, account)
}

func VerifyUserByAccount(account,password string) (model.UserDetail,bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
		
	return dal.VerifyUserByAccount(helper, account,password)
}


func QueryUserById(id int) (model.UserDetail, bool) {
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

func SaveUser(user model.UserDetail) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.SaveUser(helper, user)
}

func CreateUser(account, password, nickName, email string, status int, groups []int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	user := model.UserDetail{}
	
	user.Account = account
	user.Name = nickName
	user.Email = email
	user.Status = status
	user.Groups = groups
		
	return dal.CreateUser(helper, user, password)
}










