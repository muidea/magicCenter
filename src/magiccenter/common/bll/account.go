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
	User commonmodel.UserDetail
}

// VerifyAdministratorResponse 校验结果
type VerifyAdministratorResponse struct {
	Result common.Result
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

// QueryAllUserRequest 查询所有用户
type QueryAllUserRequest struct {
}

// QueryAllUserResponse 查询所有用户响应
type QueryAllUserResponse struct {
	Result common.Result
	Users  []commonmodel.User
}

// QueryAllGroupRequest 查询所有分组
type QueryAllGroupRequest struct {
}

// QueryAllGroupResponse 查询所有分组响应
type QueryAllGroupResponse struct {
	Result common.Result
	Groups []commonmodel.Group
}

// IsAdministrator 用户是否是管理员
func IsAdministrator(user commonmodel.UserDetail) bool {
	moduleHub := system.GetModuleHub()
	accountModule, found := moduleHub.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := VerifyAdministratorRequest{}
	request.User = user

	response := VerifyAdministratorResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		// 如果执行成功，并且返回结果也成功，则说明该用户是管理员
		return true
	}

	return false
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
