package bll

import (
	"magiccenter/common"
	"magiccenter/system"
)

// AuthorityModuleID 模块ID
const AuthorityModuleID = "759a2ee4-147a-4169-ba89-15c0c692bc16"

// QueryAuthGroupRequest 查询指定用户权限分组请求
type QueryAuthGroupRequest struct {
	AID int
}

// QueryAuthGroupResponse 查询指定用户权限分组响应
type QueryAuthGroupResponse struct {
	Result    common.Result
	AuthGroup []int
}

// QueryAuthGroup 查询指定用户分组
func QueryAuthGroup(aid int) ([]int, bool) {
	moduleHub := system.GetModuleHub()
	authorityModule, found := moduleHub.FindModule(AuthorityModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryAuthGroupRequest{AID: aid}

	response := QueryAuthGroupResponse{}
	result := authorityModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		// 如果执行成功，并且返回结果也成功，则说明该用户是管理员
		return response.AuthGroup, true
	}

	return response.AuthGroup, false
}

// InsertAuthGroupRequest 新增用户权限分组请求
type InsertAuthGroupRequest struct {
	AID int
	GID int
}

// InsertAuthGroupResponse 新增用户权限分组响应
type InsertAuthGroupResponse struct {
	Result common.Result
}

// InsertAuthGroup 新增用户分组
func InsertAuthGroup(aid, gid int) bool {
	moduleHub := system.GetModuleHub()
	authorityModule, found := moduleHub.FindModule(AuthorityModuleID)
	if !found {
		panic("can't find account module")
	}

	request := InsertAuthGroupRequest{AID: aid, GID: gid}

	response := InsertAuthGroupResponse{}
	result := authorityModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		// 如果执行成功，并且返回结果也成功，则说明该用户是管理员
		return true
	}

	return false
}

// DeleteAuthGroupRequest 删除用户权限分组请求
type DeleteAuthGroupRequest struct {
	AID int
	GID int
}

// DeleteAuthGroupResponse 删除用户权限分组响应
type DeleteAuthGroupResponse struct {
	Result common.Result
}

// DeleteAuthGroup 删除用户分组
func DeleteAuthGroup(aid, gid int) bool {
	moduleHub := system.GetModuleHub()
	authorityModule, found := moduleHub.FindModule(AuthorityModuleID)
	if !found {
		panic("can't find account module")
	}

	request := DeleteAuthGroupRequest{AID: aid, GID: gid}

	response := DeleteAuthGroupResponse{}
	result := authorityModule.Invoke(&request, &response)
	if result && response.Result.Success() {
		// 如果执行成功，并且返回结果也成功，则说明该用户是管理员
		return true
	}

	return false
}
