package ui

import (
	commonbll "magiccenter/common/bll"
	"magiccenter/kernel/modules/authority/bll"
)

// QueryAuthGroup 查询指定用户所属权限组
func QueryAuthGroup(request *commonbll.QueryAuthGroupRequest, response *commonbll.QueryAuthGroupResponse) bool {
	response.Result.ErrCode = 0
	response.AuthGroup = bll.QueryAuthGroups(request.AID)
	return true
}

// InsertAuthGroup 新增权限组
func InsertAuthGroup(request *commonbll.InsertAuthGroupRequest, response *commonbll.InsertAuthGroupResponse) bool {
	response.Result.ErrCode = 0
	ret := bll.InsertAuthGroup(request.AID, request.GID)
	if !ret {
		response.Result.ErrCode = 1
		response.Result.Reason = "新增权限组失败"
	}

	return true
}

// DeleteAuthGroup 删除权限组
func DeleteAuthGroup(request *commonbll.DeleteAuthGroupRequest, response *commonbll.DeleteAuthGroupResponse) bool {
	response.Result.ErrCode = 0
	ret := bll.DeleteAuthGroup(request.AID, request.GID)
	if !ret {
		response.Result.ErrCode = 1
		response.Result.Reason = "删除权限组失败"
	}

	return true
}
