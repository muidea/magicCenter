package bll

/* 账号相关处理
1、校验指定用户是否是管理员
2、查询指定用户信息
3、查询所有用户
*/

import (
	"magiccenter/common"
	commonmodel "magiccenter/common/model"
	"magiccenter/system"
)

// AccountModuleID Account 模块ID
const AccountModuleID = "b9e35167-b2a3-43ae-8c57-9b4379475e47"

// VerifyAdministratorRequest 校验用户是否是管理员
type VerifyAdministratorRequest struct {
	ID int
}

// VerifyAdministratorResponse 校验结果
type VerifyAdministratorResponse struct {
	Result common.Result
}

// IsAdministrator 用户是否是管理员
func IsAdministrator(id int) bool {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := VerifyAdministratorRequest{ID: id}

	response := VerifyAdministratorResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		// 如果执行成功，并且返回结果也成功，则说明该用户是管理员
		return true
	}

	return false
}

// QueryUserDetailRequest 查询指定用户请求
type QueryUserDetailRequest struct {
	ID int
}

// QueryUserDetailResponse 查询指定用户结果
type QueryUserDetailResponse struct {
	Result common.Result
	User   commonmodel.UserDetail
}

// QueryUserDetail 查询指定用户
func QueryUserDetail(id int) (commonmodel.UserDetail, bool) {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryUserDetailRequest{}
	request.ID = id

	response := QueryUserDetailResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		return response.User, true
	}

	return response.User, false
}

// QueryAllUserRequest 查询所有用户
type QueryAllUserRequest struct {
}

// QueryAllUserResponse 查询所有用户响应
type QueryAllUserResponse struct {
	Result common.Result
	Users  []commonmodel.User
}

// QueryAllUser 查询所有用户
func QueryAllUser() ([]commonmodel.User, bool) {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryAllUserRequest{}
	response := QueryAllUserResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		return response.Users, true
	}

	return response.Users, false
}

// CreateUserRequest 新建用户请求
type CreateUserRequest struct {
	Account string
	EMail   string
}

// CreateUserResponse 新建用户响应
type CreateUserResponse struct {
	Result common.Result
	User   commonmodel.User
}

// CreateUser 新建用户
func CreateUser(account, email string) (commonmodel.User, bool) {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := CreateUserRequest{Account: account, EMail: email}
	response := CreateUserResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		return response.User, true
	}

	return response.User, false
}

// UpdateUserRequest 更新用户信息请求
type UpdateUserRequest struct {
	User commonmodel.UserDetail
}

// UpdateUserResponse 更新用户信息响应
type UpdateUserResponse struct {
	Result common.Result
	User   commonmodel.UserDetail
}

// UpdateUser 更新用户信息
func UpdateUser(user commonmodel.UserDetail) (commonmodel.UserDetail, bool) {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := UpdateUserRequest{User: user}
	response := UpdateUserResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		return response.User, true
	}

	return response.User, false
}

// DeleteUserRequest 删除用户请求
type DeleteUserRequest struct {
	ID int
}

// DeleteUserResponse 删除用户响应
type DeleteUserResponse struct {
	Result common.Result
}

// DeleteUser 删除指定用户
func DeleteUser(id int) bool {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := DeleteUserRequest{ID: id}
	response := DeleteUserResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		return true
	}

	return false
}

// QueryAllGroupRequest 查询所有分组
type QueryAllGroupRequest struct {
}

// QueryAllGroupResponse 查询所有分组响应
type QueryAllGroupResponse struct {
	Result common.Result
	Groups []commonmodel.Group
}

// QueryAllGroup 查询所有分组
func QueryAllGroup() ([]commonmodel.Group, bool) {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryAllGroupRequest{}
	response := QueryAllGroupResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		return response.Groups, true
	}

	return response.Groups, false
}

// QueryGroupsRequest 查询指定分组请求
type QueryGroupsRequest struct {
	Ids []int
}

// QueryGroupsResponse 查询指定分组响应
type QueryGroupsResponse struct {
	Result common.Result
	Groups []commonmodel.Group
}

// QueryGroups 查询指定分组
func QueryGroups(ids []int) ([]commonmodel.Group, bool) {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryGroupsRequest{Ids: ids}
	response := QueryGroupsResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		return response.Groups, true
	}

	return response.Groups, false
}

// CreateGroupRequest 新建分组请求
type CreateGroupRequest struct {
	Name        string
	Description string
}

// CreateGroupResponse 新建分组响应
type CreateGroupResponse struct {
	Result common.Result
	Group  commonmodel.Group
}

// CreateGroup 新建分组
func CreateGroup(name, description string) (commonmodel.Group, bool) {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := CreateGroupRequest{Name: name, Description: description}
	response := CreateGroupResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		return response.Group, true
	}

	return response.Group, false
}

// UpdateGroupRequest 更新分组信息请求
type UpdateGroupRequest struct {
	Group commonmodel.Group
}

// UpdateGroupResponse 更新分组信息响应
type UpdateGroupResponse struct {
	Result common.Result
	Group  commonmodel.Group
}

// UpdateGroup 更新指定分组
func UpdateGroup(group commonmodel.Group) (commonmodel.Group, bool) {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := UpdateGroupRequest{Group: group}
	response := UpdateGroupResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		return response.Group, true
	}

	return response.Group, false
}

// DeleteGroupRequest 删除指定分组请求
type DeleteGroupRequest struct {
	ID int
}

// DeleteGroupResponse 删除指定分组响应
type DeleteGroupResponse struct {
	Result common.Result
}

// DeleteGroup 更新指定分组
func DeleteGroup(id int) bool {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := DeleteGroupRequest{ID: id}
	response := DeleteGroupResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		return true
	}

	return false
}
