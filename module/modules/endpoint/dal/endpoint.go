package dal

import (
	"fmt"

	"github.com/muidea/magicCenter/common/dbhelper"
	"github.com/muidea/magicCommon/foundation/util"
	"github.com/muidea/magicCommon/model"
)

// InsertEndpoint 新增Endpoint记录
func InsertEndpoint(helper dbhelper.DBHelper, id, name, description string, user []int, status int, authToken string) (model.Endpoint, bool) {
	endpoint := model.Endpoint{ID: id, Name: name, Description: description, User: user, Status: status}
	sql := fmt.Sprintf("insert into endpoint_registry (id, name, description, user, status, authToken) values ('%s','%s','%s','%s',%d,'%s')", id, name, description, util.IntArray2Str(user), status, authToken)
	num, ok := helper.Execute(sql)
	if !ok || num != 1 {
		return endpoint, false
	}

	endpoint.AuthToken = authToken

	return endpoint, true
}

// DeleteEndpoint 删除Endpoint记录
func DeleteEndpoint(helper dbhelper.DBHelper, id string) bool {
	sql := fmt.Sprintf("delete from endpoint_registry where id='%s'", id)
	num, ok := helper.Execute(sql)
	return ok && num == 1
}

// UpdateEndpoint 更新Endpoint记录
func UpdateEndpoint(helper dbhelper.DBHelper, endpoint model.Endpoint) (model.Endpoint, bool) {
	sql := fmt.Sprintf("update endpoint_registry set name='%s', description='%s', user='%s', status=%d, authToken='%s' where id='%s'", endpoint.Name, endpoint.Description, util.IntArray2Str(endpoint.User), endpoint.Status, endpoint.AuthToken, endpoint.ID)
	num, ok := helper.Execute(sql)
	if !ok || num != 1 {
		return endpoint, false
	}

	return endpoint, true
}

// QueryAllEndpoint 查询所有Endpoint
func QueryAllEndpoint(helper dbhelper.DBHelper) []model.Endpoint {
	endpoints := []model.Endpoint{}
	sql := fmt.Sprintf("select id, name, description, user, status, authToken from endpoint_registry")

	helper.Query(sql)
	defer helper.Finish()

	for helper.Next() {
		endpoint := model.Endpoint{}
		users := ""
		helper.GetValue(&endpoint.ID, &endpoint.Name, &endpoint.Description, &users, &endpoint.Status, &endpoint.AuthToken)
		endpoint.User, _ = util.Str2IntArray(users)

		endpoints = append(endpoints, endpoint)
	}

	return endpoints
}

// QueryEndpointByID 查询指定Endpoint
func QueryEndpointByID(helper dbhelper.DBHelper, id string) (model.Endpoint, bool) {
	endpoint := model.Endpoint{}
	sql := fmt.Sprintf("select id, name, description, user, status, authToken from endpoint_registry where id='%s'", id)

	helper.Query(sql)
	defer helper.Finish()

	if helper.Next() {
		users := ""
		helper.GetValue(&endpoint.ID, &endpoint.Name, &endpoint.Description, &users, &endpoint.Status, &endpoint.AuthToken)
		endpoint.User, _ = util.Str2IntArray(users)

		return endpoint, true
	}

	return endpoint, false
}
