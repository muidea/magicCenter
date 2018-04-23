package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCommon/model"
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

func (i *impl) GetUsers(ids []int) []model.User {
	return i.userHandler.getUsers(ids)
}

func (i *impl) FindUserByID(id int) (model.UserDetail, bool) {
	return i.userHandler.findUserByID(id)
}

func (i *impl) FindUserByGroup(groupID int) []model.User {
	return i.userHandler.findUserByGroup(groupID)
}

func (i *impl) FindUserByAccount(account, password string) (model.UserDetail, bool) {
	return i.userHandler.findUserByAccount(account, password)
}

func (i *impl) CreateUser(account, password, email string, groups []int) (model.UserDetail, bool) {
	return i.userHandler.createUser(account, password, email, groups)
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

func (i *impl) GetAllGroup() []model.GroupDetail {
	return i.groupHandler.getAllGroups()
}

func (i *impl) GetGroups(ids []int) []model.Group {
	return i.groupHandler.getGroups(ids)
}

func (i *impl) FindGroupByID(id int) (model.GroupDetail, bool) {
	return i.groupHandler.findGroupByID(id)
}

func (i *impl) FindSubGroup(id int) []model.Group {
	return i.groupHandler.findSubGroup(id)
}

func (i *impl) FindGroupByName(name string) (model.GroupDetail, bool) {
	return i.groupHandler.findGroupByName(name)
}

func (i *impl) CreateGroup(name, description string, catalog int) (model.GroupDetail, bool) {
	return i.groupHandler.createGroup(name, description, catalog)
}

func (i *impl) SaveGroup(group model.GroupDetail) (model.GroupDetail, bool) {
	return i.groupHandler.saveGroup(group)
}

func (i *impl) DestroyGroup(id int) bool {

	subGroups := i.groupHandler.findSubGroup(id)
	groupUsers := i.userHandler.findUserByGroup(id)
	if len(subGroups) > 0 || len(groupUsers) > 0 {
		return false
	}

	return i.groupHandler.destroyGroup(id)
}

func (i *impl) GetAccountSummary() model.AccountSummary {
	result := model.AccountSummary{}
	userCount := len(i.userHandler.getAllUser())
	userItem := model.UnitSummary{Name: "用户", Type: "user", Count: userCount}
	result = append(result, userItem)
	groupCount := len(i.groupHandler.getAllGroups())
	groupItem := model.UnitSummary{Name: "分组", Type: "group", Count: groupCount}
	result = append(result, groupItem)

	return result
}

func (i *impl) GetLastAccount(count int) model.AccountRecord {
	result := model.AccountRecord{}
	users := i.userHandler.GetLastRegisterUser(count)
	for _, v := range users {
		item := model.AccountUnit{Name: v.Account, RegisterDate: v.RegisterTime}

		result = append(result, item)
	}
	return result
}
