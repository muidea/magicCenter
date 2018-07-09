package common

import (
	"muidea.com/magicCommon/common"
	"muidea.com/magicCommon/model"
)

// AccountHandler 账号信息处理Handler
type AccountHandler interface {
	GetAllUserIDs() []int
	GetAllUser() []model.User
	GetAllUserDetail(filter *common.PageFilter) []model.UserDetail
	GetUsers(ids []int) []model.User
	FindUserByID(id int) (model.UserDetail, bool)
	FindUserByGroup(groupID int) []model.UserDetail
	FindUserByAccount(account, password string) (model.UserDetail, bool)
	CreateUser(account, password, email string, groups []int) (model.UserDetail, bool)
	SaveUser(user model.UserDetail) (model.UserDetail, bool)
	SaveUserWithPassword(user model.UserDetail, password string) (model.UserDetail, bool)
	DestroyUserByID(id int) bool
	DestroyUserByAccount(account, password string) bool

	GetAllGroup() []model.Group
	GetAllGroupDetail(filter *common.PageFilter) []model.GroupDetail
	GetGroups(ids []int) []model.Group
	FindGroupByID(id int) (model.GroupDetail, bool)
	FindSubGroup(id int) []model.Group
	FindGroupByName(name string) (model.GroupDetail, bool)
	CreateGroup(name, description string, catalog int) (model.GroupDetail, bool)
	SaveGroup(group model.GroupDetail) (model.GroupDetail, bool)
	DestroyGroup(id int) bool

	GetAccountSummary() model.AccountSummary
	GetLastAccount(count int) model.AccountRecord
}
