package common

import "muidea.com/magicCenter/application/common/model"

// AccountHandler 账号信息处理Handler
type AccountHandler interface {
	GetAllUser() []model.UserDetail
	FindUserByID(id int) (model.UserDetail, bool)
	FindUserByAccount(account, password string) (model.UserDetail, bool)
	CreateUser(account, email string, groups []int) (model.UserDetail, bool)
	SaveUser(user model.UserDetail) (model.UserDetail, bool)
	SaveUserWithPassword(user model.UserDetail, password string) (model.UserDetail, bool)
	DestroyUserByID(id int) bool
	DestroyUserByAccount(account, password string) bool

	GetAllGroup() []model.Group
	FindGroupByID(id int) (model.Group, bool)
	FindGroupByName(name string) (model.Group, bool)
	CreateGroup(name, description string) (model.Group, bool)
	SaveGroup(group model.Group) (model.Group, bool)
	DestroyGroup(id int) bool
}
