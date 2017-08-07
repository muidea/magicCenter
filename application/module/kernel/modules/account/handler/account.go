package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
)

// CreateAccountHandler 创建Account处理器
func CreateAccountHandler() common.AccountHandler {
	dbhelper, _ := dbhelper.NewHelper()
	i := impl{userHandler: userActionHandler{dbhelper: dbhelper}, groupHandler: groupActionHandler{dbhelper: dbhelper}}

	return &i
}

type impl struct {
	userHandler  userActionHandler
	groupHandler groupActionHandler
}

func (i *impl) GetAllUser() []model.UserDetail {
	return i.userHandler.getAllUser()
}

func (i *impl) FindUserByID(id int) (model.UserDetail, bool) {
	return i.userHandler.findUserByID(id)
}

func (i *impl) FindUserByAccount(account, password string) (model.UserDetail, bool) {
	return i.userHandler.findUserByAccount(account, password)
}

func (i *impl) CreateUser(account, email string) (model.UserDetail, bool) {
	return i.userHandler.createUser(account, email)
}

func (i *impl) SaveUser(user model.UserDetail) (model.UserDetail, bool) {
	return i.userHandler.saveUser(user)
}

func (i *impl) SaveUserWithPassword(user model.UserDetail, password string) (model.UserDetail, bool) {
	return i.userHandler.saveUserWithPassword(user, password)
}

func (i *impl) DestroyUserByID(id int) bool {
	return i.userHandler.destroyUserByID(id)
}

func (i *impl) DestroyUserByAccount(account, password string) bool {
	return i.userHandler.destroyUserByAccount(account, password)
}

func (i *impl) GetAllGroup() []model.Group {
	return i.groupHandler.getAllGroups()
}

func (i *impl) FindGroupByID(id int) (model.Group, bool) {
	return i.groupHandler.findGroupByID(id)
}

func (i *impl) FindGroupByName(name string) (model.Group, bool) {
	return i.groupHandler.findGroupByName(name)
}

func (i *impl) CreateGroup(name, description string) (model.Group, bool) {
	return i.groupHandler.createGroup(name, description)
}

func (i *impl) SaveGroup(group model.Group) (model.Group, bool) {
	return i.groupHandler.saveGroup(group)
}

func (i *impl) DestroyGroup(id int) bool {
	return i.groupHandler.destroyGroup(id)
}
