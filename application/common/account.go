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

	GetAllGroup() []model.GroupDetail
	FindGroupByID(id int) (model.GroupDetail, bool)
	FindGroupByName(name string) (model.GroupDetail, bool)
	CreateGroup(name, description string, catalog int) (model.GroupDetail, bool)
	SaveGroup(group model.GroupDetail) (model.GroupDetail, bool)
	DestroyGroup(id int) bool

	GetAccountSummary() model.AccountSummary
	GetLastAccount(count int) model.AccountRecord
}
