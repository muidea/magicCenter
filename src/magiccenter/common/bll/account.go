package bll

import (
	"magiccenter/common"
	commonmodel "magiccenter/common/model"
	"magiccenter/module"
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

// IsAdministrator 用户是否是管理员
func IsAdministrator(user commonmodel.UserDetail) bool {
	accountModule, found := module.FindModule(AccountModuleID)
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
	accountModule, found := module.FindModule(AccountModuleID)
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
func QueryAllUser() []commonmodel.User {
	accountModule, found := module.FindModule(AccountModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryAllUserRequest{}
	response := QueryAllUserResponse{}
	result := accountModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		return response.Users
	}

	return response.Users
}
