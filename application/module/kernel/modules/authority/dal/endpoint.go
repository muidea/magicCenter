package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

// InsertEndpoint 新增Endpoint记录
func InsertEndpoint(helper dbhelper.DBHelper, id, name, description string, user []int, status int, accessToken string) (model.Endpoint, bool) {
	endpoint := model.Endpoint{ID: id, Name: name, Description: description, User: user, Status: status}
	sql := fmt.Sprintf("insert into authority_endpoint (id, name, description, user, status, accessToken) values ('%s','%s','%s','%s',%d,'%s')", id, name, description, util.IntArray2Str(user), status, accessToken)
	num, ok := helper.Execute(sql)
	if !ok || num != 1 {
		return endpoint, false
	}

	endpoint.AccessToken = accessToken

	return endpoint, true
}

// DeleteEndpoint 删除Endpoint记录
func DeleteEndpoint(helper dbhelper.DBHelper, id string) bool {
	sql := fmt.Sprintf("delete from authority_endpoint where id='%s'", id)
	num, ok := helper.Execute(sql)
	return ok && num == 1
}

// UpdateEndpoint 更新Endpoint记录
func UpdateEndpoint(helper dbhelper.DBHelper, endpoint model.Endpoint) (model.Endpoint, bool) {
	sql := fmt.Sprintf("update authority_endpoint set name='%s', description='%s', user='%s', status=%d, accessToken='%s' where id='%s'", endpoint.Name, endpoint.Description, util.IntArray2Str(endpoint.User), endpoint.Status, endpoint.AccessToken, endpoint.ID)
	num, ok := helper.Execute(sql)
	if !ok || num != 1 {
		return endpoint, false
	}

	return endpoint, true
}

// QueryAllEndpoint 查询所有Endpoint
func QueryAllEndpoint(helper dbhelper.DBHelper) []model.Endpoint {
	endpoints := []model.Endpoint{}
	sql := fmt.Sprintf("select id, name, description, user, status, accessToken from authority_endpoint")

	helper.Query(sql)
	defer helper.Finish()

	for helper.Next() {
		endpoint := model.Endpoint{}
		users := ""
		helper.GetValue(&endpoint.ID, &endpoint.Name, &endpoint.Description, &users, &endpoint.Status, &endpoint.AccessToken)
		endpoint.User, _ = util.Str2IntArray(users)

		endpoints = append(endpoints, endpoint)
	}

	return endpoints
}

// QueryEndpointByID 查询指定Endpoint
func QueryEndpointByID(helper dbhelper.DBHelper, id string) (model.Endpoint, bool) {
	endpoint := model.Endpoint{}
	sql := fmt.Sprintf("select id, name, description, user, status, accessToken from authority_endpoint where id='%s'", id)

	helper.Query(sql)
	defer helper.Finish()

	if helper.Next() {
		users := ""
		helper.GetValue(&endpoint.ID, &endpoint.Name, &endpoint.Description, &users, &endpoint.Status, &endpoint.AccessToken)
		endpoint.User, _ = util.Str2IntArray(users)

		return endpoint, true
	}

	return endpoint, false
}
