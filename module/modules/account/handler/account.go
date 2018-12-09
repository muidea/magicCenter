package handler

import (
	"muidea.com/magicCenter/common"
	common_const "muidea.com/magicCommon/common"
	common_util "muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

// CreateAccountHandler 创建Account处理器
func CreateAccountHandler() common.AccountHandler {
	i := impl{userHandler: userActionHandler{}, groupHandler: groupActionHandler{}}

	return &i
}

type impl struct {
	userHandler  userActionHandler
	groupHandler groupActionHandler
}

func (i *impl) GetAllUserIDs() []int {
	return i.userHandler.getAllUserIDs()
}

func (i *impl) GetAllUser() []model.User {
	return i.userHandler.getAllUser()
}

func (i *impl) GetAllUserDetail(filter *common_util.PageFilter) []model.UserDetail {
	return i.userHandler.getAllUserDetail(filter)
}

func (i *impl) GetUsers(ids []int) []model.User {
	return i.userHandler.getUsers(ids)
}

func (i *impl) FindUserByID(id int) (model.UserDetail, bool) {
	if id == common_const.SystemAccountUser.ID {
		return common_const.SystemAccountUser, true
	}

	return i.userHandler.findUserByID(id)
}

func (i *impl) FindUserByGroup(groupID int) []model.UserDetail {
	return i.userHandler.findUserByGroup(groupID)
}

func (i *impl) FindUserByAccount(account, password string) (model.UserDetail, bool) {
	return i.userHandler.findUserByAccount(account, password)
}

func (i *impl) CreateUser(account, password, email string, groups []int) (model.UserDetail, bool) {
	if account == common_const.SystemAccountUser.Name {
		return model.UserDetail{}, false
	}

	return i.userHandler.createUser(account, password, email, groups)
}

func (i *impl) SaveUser(user model.UserDetail) (model.UserDetail, bool) {
	if user.Name == common_const.SystemAccountUser.Name {
		return user, false
	}

	if user.ID == common_const.SystemAccountUser.ID {
		return user, false
	}

	return i.userHandler.saveUser(user)
}

func (i *impl) SaveUserWithPassword(user model.UserDetail, password string) (model.UserDetail, bool) {
	return i.userHandler.saveUserWithPassword(user, password)
}

func (i *impl) DestroyUserByID(id int) bool {
	if id == common_const.SystemAccountUser.ID {
		return false
	}

	return i.userHandler.destroyUserByID(id)
}

func (i *impl) DestroyUserByAccount(account, password string) bool {
	return i.userHandler.destroyUserByAccount(account, password)
}

func (i *impl) GetAllGroup() []model.Group {
	return i.groupHandler.getAllGroup()
}

func (i *impl) GetAllGroupDetail(filter *common_util.PageFilter) []model.GroupDetail {
	return i.groupHandler.getAllGroupDetail(filter)
}

func (i *impl) GetGroups(ids []int) []model.Group {
	return i.groupHandler.getGroups(ids)
}

func (i *impl) FindGroupByID(id int) (model.GroupDetail, bool) {
	if id == common_const.SystemAccountGroup.ID {
		return common_const.SystemAccountGroup, true
	}

	return i.groupHandler.findGroupByID(id)
}

func (i *impl) FindSubGroup(id int) []model.Group {
	return i.groupHandler.findSubGroup(id)
}

func (i *impl) FindGroupByName(name string) (model.GroupDetail, bool) {
	if name == common_const.SystemAccountGroup.Name {
		return common_const.SystemAccountGroup, true
	}

	return i.groupHandler.findGroupByName(name)
}

func (i *impl) CreateGroup(name, description string, catalog int) (model.GroupDetail, bool) {
	if name == common_const.SystemAccountGroup.Name {
		return model.GroupDetail{}, false
	}

	return i.groupHandler.createGroup(name, description, catalog)
}

func (i *impl) SaveGroup(group model.GroupDetail) (model.GroupDetail, bool) {
	if group.ID == common_const.SystemAccountGroup.ID {
		return group, false
	}
	if group.Name == common_const.SystemAccountGroup.Name {
		return group, false
	}

	return i.groupHandler.saveGroup(group)
}

func (i *impl) DestroyGroup(id int) bool {
	if id == common_const.SystemAccountGroup.ID {
		return false
	}

	subGroups := i.groupHandler.findSubGroup(id)
	groupUsers := i.userHandler.findUserByGroup(id)
	if len(subGroups) > 0 || len(groupUsers) > 0 {
		return false
	}

	return i.groupHandler.destroyGroup(id)
}

func (i *impl) GetAccountSummary() model.AccountSummary {
	result := model.AccountSummary{}
	userCount := i.userHandler.getUserCount()
	userItem := model.UnitSummary{Name: "用户", Type: "user", Count: userCount}
	result = append(result, userItem)
	groupCount := i.groupHandler.getGroupCount()
	groupItem := model.UnitSummary{Name: "分组", Type: "group", Count: groupCount}
	result = append(result, groupItem)

	return result
}

func (i *impl) GetLastAccount(count int) model.AccountRecord {
	result := model.AccountRecord{}
	users := i.userHandler.GetLastRegisterUser(count)
	for _, v := range users {
		item := model.AccountUnit{Name: v.Name, RegisterDate: v.RegisterTime}

		result = append(result, item)
	}
	return result
}
