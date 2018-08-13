package handler

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/account/dal"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/model"
)

type userActionHandler struct {
}

func (i *userActionHandler) getUserCount() int {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryUserCount(dbhelper)
}

func (i *userActionHandler) getAllUserIDs() []int {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllUserIDs(dbhelper)
}

func (i *userActionHandler) getAllUser() []model.User {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllUser(dbhelper)
}

func (i *userActionHandler) getAllUserDetail(filter *common_def.PageFilter) []model.UserDetail {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllUserDetail(dbhelper, filter)
}

func (i *userActionHandler) getUsers(ids []int) []model.User {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryUsers(dbhelper, ids)
}

func (i *userActionHandler) findUserByID(id int) (model.UserDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryUserByID(dbhelper, id)
}

func (i *userActionHandler) findUserByGroup(groupID int) []model.UserDetail {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryUserByGroup(dbhelper, groupID)
}

func (i *userActionHandler) findUserByAccount(account, password string) (model.UserDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryUserByAccount(dbhelper, account, password)
}

// 新建User
func (i *userActionHandler) createUser(account, password, email string, groups []int) (model.UserDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.CreateUser(dbhelper, account, password, email, groups)
}

// 保存User
func (i *userActionHandler) saveUser(user model.UserDetail) (model.UserDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.SaveUser(dbhelper, user)
}

func (i *userActionHandler) saveUserWithPassword(user model.UserDetail, password string) (model.UserDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.SaveUserWithPassword(dbhelper, user, password)
}

// 销毁User
func (i *userActionHandler) destroyUserByID(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.DeleteUser(dbhelper, id)
}

func (i *userActionHandler) destroyUserByAccount(account, password string) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.DeleteUserByAccount(dbhelper, account, password)
}

func (i *userActionHandler) GetLastRegisterUser(count int) []model.UserDetail {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.GetLastRegisterUser(dbhelper, count)
}
