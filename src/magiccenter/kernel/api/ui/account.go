package ui

import (
	"encoding/json"
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	commonmodel "magiccenter/common/model"
	"net/http"
)

// UserList 用户列表
type UserList struct {
	common.Result
	UserList []commonmodel.User
}

// GroupList 用户列表
type GroupList struct {
	common.Result
	GroupList []commonmodel.Group
}

// GetUserListActionHandler 获取User列表
func GetUserListActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetUserListActionHandler")

	result := UserList{}
	found := false

	result.UserList, found = commonbll.QueryAllUser()
	if found {
		result.ErrCode = 0
	} else {
		result.ErrCode = 1
		result.Reason = "查询失败"
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// GetGroupListActionHandler 获取User列表
func GetGroupListActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetGroupListActionHandler")

	result := GroupList{}
	found := false

	result.GroupList, found = commonbll.QueryAllGroup()
	if found {
		result.ErrCode = 0
	} else {
		result.ErrCode = 1
		result.Reason = "查询失败"
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
