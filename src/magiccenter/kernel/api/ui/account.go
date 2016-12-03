package ui

import (
	"encoding/json"
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	commonmodel "magiccenter/common/model"
	"net/http"

	"strconv"

	"muidea.com/util"
)

// UserList 用户列表
type UserList struct {
	common.Result
	UserList []commonmodel.User
}

// UserDetail 单用户详情
type UserDetail struct {
	common.Result
	User commonmodel.UserDetail
}

// GetUserActionHandler 获取User列表
func GetUserActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetUserActionHandler")

	params := util.SplitParam(r.URL.RawQuery)
	uid, found := params["id"]
	if !found {
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
	} else {
		result := UserDetail{}
		for true {
			id, err := strconv.Atoi(uid)
			if err != nil {
				result.ErrCode = 1
				result.Reason = "参数非法"
				break
			}

			result.User, found = commonbll.QueryUserDetail(id)
			if !found {
				result.ErrCode = 1
				result.Reason = "无效参数"
				break
			}

			result.ErrCode = 0
			break
		}

		b, err := json.Marshal(result)
		if err != nil {
			panic("json.Marshal, failed, err:" + err.Error())
		}

		w.Write(b)
	}
}

// SingleUser 单用户信息
type SingleUser struct {
	common.Result
	User commonmodel.User
}

// PostUserActionHandler 新建用户
func PostUserActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PostUserActionHandler")
	result := SingleUser{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		account := r.FormValue("account")
		email := r.FormValue("email")

		ret := false
		result.User, ret = commonbll.CreateUser(account, email)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "创建用户失败"
			break
		}

		result.Result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// SingleUserDetail 单用户信息详情
type SingleUserDetail struct {
	common.Result
	User commonmodel.UserDetail
}

// PutUserActionHandler 更新用户
func PutUserActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PutUserActionHandler")
	result := SingleUserDetail{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		uid := r.FormValue("id")
		account := r.FormValue("account")
		email := r.FormValue("email")
		nickname := r.FormValue("nickname")
		status := r.FormValue("status")

		id, err := strconv.Atoi(uid)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		user, found := commonbll.QueryUserDetail(id)
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}
		if user.Account != account {
			// 如果账号信息不一致，说明是非法请求
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}
		if len(email) > 0 {
			user.Email = email
		}
		if len(nickname) > 0 {
			user.Name = nickname
		}
		if len(status) > 0 {
			val, err := strconv.Atoi(status)
			if err != nil {
				result.Result.ErrCode = 1
				result.Result.Reason = "无效参数"
				break
			}
			user.Status = val
		}

		result.User, found = commonbll.UpdateUser(user)
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "更新用户信息失败"
		}

		result.Result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// DeleteUserActionHandler 新建用户
func DeleteUserActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteUserActionHandler")
	result := common.Result{}

	params := util.SplitParam(r.URL.RawQuery)
	uid, found := params["id"]
	for true {
		if !found {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}
		id, err := strconv.Atoi(uid)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "参数非法"
			break
		}

		found = commonbll.DeleteUser(id)
		if !found {
			result.ErrCode = 1
			result.Reason = "删除用户失败"
			break
		}

		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// GroupList 用户列表
type GroupList struct {
	common.Result
	GroupList []commonmodel.Group
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
