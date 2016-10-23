package bll

import (
	"magiccenter/common/model"
	"magiccenter/kernel/modules/account/dal"
	"magiccenter/system"
)

// QueryAllUserList 查询全部用户列表
func QueryAllUserList() []model.User {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllUserList(helper)
}

// QueryAllUser 查询全部用户
func QueryAllUser() []model.UserDetail {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllUser(helper)
}

// QueryUserByAccount 查询指定账号的用户信息
func QueryUserByAccount(account string) (model.UserDetail, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryUserByAccount(helper, account)
}

// VerifyUserByAccount 校验指定账号的用户信息
func VerifyUserByAccount(account, password string) (model.UserDetail, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.VerifyUserByAccount(helper, account, password)
}

// QueryUserByID 查询指定账号用户信息
func QueryUserByID(id int) (model.UserDetail, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryUserByID(helper, id)
}

// DeleteUser 删除用户
func DeleteUser(id int) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteUser(helper, id)
}

// SaveUser 保存用户
func SaveUser(user model.UserDetail) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.SaveUser(helper, user)
}

// CreateUser 创建新用户
func CreateUser(account, email string, status int, groups []int) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	user := model.UserDetail{}

	user.Account = account
	user.Name = ""
	user.Email = email
	user.Status = status
	user.Groups = groups

	return dal.CreateUser(helper, user)
}

// UpdateUserWithPassword 跟新用户信息
func UpdateUserWithPassword(usr model.UserDetail, password string) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.SaveUserWithPassword(helper, usr, password)
}
