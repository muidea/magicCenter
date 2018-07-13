package handler

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/account/dal"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/model"
)

type userActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *userActionHandler) getUserCount() int {
	return dal.QueryUserCount(i.dbhelper)
}

func (i *userActionHandler) getAllUserIDs() []int {
	return dal.QueryAllUserIDs(i.dbhelper)
}

func (i *userActionHandler) getAllUser() []model.User {
	return dal.QueryAllUser(i.dbhelper)
}

func (i *userActionHandler) getAllUserDetail(filter *common_def.PageFilter) []model.UserDetail {
	return dal.QueryAllUserDetail(i.dbhelper, filter)
}

func (i *userActionHandler) getUsers(ids []int) []model.User {
	return dal.QueryUsers(i.dbhelper, ids)
}

func (i *userActionHandler) findUserByID(id int) (model.UserDetail, bool) {
	return dal.QueryUserByID(i.dbhelper, id)
}

func (i *userActionHandler) findUserByGroup(groupID int) []model.UserDetail {
	return dal.QueryUserByGroup(i.dbhelper, groupID)
}

func (i *userActionHandler) findUserByAccount(account, password string) (model.UserDetail, bool) {
	return dal.QueryUserByAccount(i.dbhelper, account, password)
}

// 新建User
func (i *userActionHandler) createUser(account, password, email string, groups []int) (model.UserDetail, bool) {
	return dal.CreateUser(i.dbhelper, account, password, email, groups)
}

// 保存User
func (i *userActionHandler) saveUser(user model.UserDetail) (model.UserDetail, bool) {
	return dal.SaveUser(i.dbhelper, user)
}

func (i *userActionHandler) saveUserWithPassword(user model.UserDetail, password string) (model.UserDetail, bool) {
	return dal.SaveUserWithPassword(i.dbhelper, user, password)
}

// 销毁User
func (i *userActionHandler) destroyUserByID(id int) bool {
	return dal.DeleteUser(i.dbhelper, id)
}

func (i *userActionHandler) destroyUserByAccount(account, password string) bool {
	return dal.DeleteUserByAccount(i.dbhelper, account, password)
}

func (i *userActionHandler) GetLastRegisterUser(count int) []model.UserDetail {
	return dal.GetLastRegisterUser(i.dbhelper, count)
}
