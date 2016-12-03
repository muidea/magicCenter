package ui

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	"magiccenter/common/model"
	"magiccenter/kernel/modules/account/bll"
	"magiccenter/system"
	"net/http"
	"strconv"

	"muidea.com/util"
)

// VerifyAdminUser 校验用户是否是管理员
func VerifyAdminUser(request *commonbll.VerifyAdministratorRequest, response *commonbll.VerifyAdministratorResponse) bool {
	response.Result.ErrCode = 1

	user, found := bll.QueryUserByID(request.ID)
	if !found {
		return false
	}

	groups := user.Groups
	for _, gid := range groups {
		group, found := bll.QueryGroupByID(gid)
		if found && group.AdminGroup() {
			response.Result.ErrCode = 0
			break
		}
	}

	return true
}

// QueryAllUser 查询所有用户
func QueryAllUser(request *commonbll.QueryAllUserRequest, response *commonbll.QueryAllUserResponse) bool {
	response.Result.ErrCode = 0
	response.Users = bll.QueryAllUser()

	return true
}

// QueryUserDetail 查询指定用户
func QueryUserDetail(request *commonbll.QueryUserDetailRequest, response *commonbll.QueryUserDetailResponse) bool {
	response.Result.ErrCode = 0

	ret := false
	response.User, ret = bll.QueryUserByID(request.ID)
	if !ret {
		response.Result.ErrCode = 1
		response.Result.Reason = "指定用户不存在"
	}

	return true
}

// CreateUser 新建用户
func CreateUser(request *commonbll.CreateUserRequest, response *commonbll.CreateUserResponse) bool {
	response.Result.ErrCode = 0

	ret := false
	response.User, ret = bll.CreateUser(request.Account, request.EMail)
	if !ret {
		response.Result.ErrCode = 1
		response.Result.Reason = "新建用户失败"
	}

	return true
}

// UpdateUser 新建用户
func UpdateUser(request *commonbll.UpdateUserRequest, response *commonbll.UpdateUserResponse) bool {
	response.Result.ErrCode = 0

	ret := false
	response.User, ret = bll.UpdateUser(request.User)
	if !ret {
		response.Result.ErrCode = 1
		response.Result.Reason = "新建用户失败"
	}

	return true
}

// DeleteUser 删除用户
func DeleteUser(request *commonbll.DeleteUserRequest, response *commonbll.DeleteUserResponse) bool {
	response.Result.ErrCode = 0

	ret := bll.DeleteUser(request.ID)
	if !ret {
		response.Result.ErrCode = 1
		response.Result.Reason = "删除用户失败"
	}

	return true
}

// ManageUserView 用户管理视图数据
type ManageUserView struct {
	Users  []model.UserDetail
	Groups []model.Group
}

// AllUserList 所有用户结果
type AllUserList struct {
	Users []model.UserDetail
}

// SingleUserDetail 单用户结果
type SingleUserDetail struct {
	common.Result
	User model.UserDetail
}

// ManageUserViewHandler 用户管理视图处理器
func ManageUserViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageUserViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	htmlFile := system.GetHTMLPath("kernel/modules/account/account.html")
	t, err := template.ParseFiles(htmlFile)
	if err != nil {
		panic("parse files failed")
	}

	view := ManageUserView{}
	view.Users = bll.QueryAllUserDetail()
	view.Groups = bll.QueryAllGroup()

	t.Execute(w, view)
}

// QueryAllUserActionHandler 查询所有用户信息处理器
func QueryAllUserActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllUserActionHandler")

	result := AllUserList{}
	result.Users = bll.QueryAllUserDetail()

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// QueryUserActionHandler 查询单个用户信息处理器
func QueryUserActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryUserActionHandler")

	result := SingleUserDetail{}

	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		uid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		result.User, found = bll.QueryUserByID(uid)
		if !found {
			result.ErrCode = 1
			result.Reason = "指定User不存在"
			break
		}

		result.ErrCode = 0
		result.Reason = "查询成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// DeleteUserActionHandler 删除用户处理器
func DeleteUserActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteUserActionHandler")

	result := common.Result{}
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		uid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		ok := bll.DeleteUser(uid)
		if !ok {
			result.ErrCode = 1
			result.Reason = "删除用户失败"
			break
		}

		result.ErrCode = 0
		result.Reason = "删除用户成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

func sendVerifyMail(user, email, id string) {
	configuration := system.GetConfiguration()
	systemInfo := configuration.GetSystemInfo()

	subject := "MagicCenter账号验证"

	content := fmt.Sprintf("<html><head><title>用户信息验证</title></head><body><p>Hi %s</p><p><a href='http://%s/account/verifyAccount/?id=%s'>请点击链接继续验证用户信息</a></p><p>该邮件由MagicCenter自动发送，请勿回复该邮件</p></body></html>", user, systemInfo.Domain, id)

	mailList := []string{}
	mailList = append(mailList, email)
	commonbll.PostMail(mailList, subject, content)
}

// SaveAccountActionHandler 保存账号信息
func SaveAccountActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("SaveAccountActionHandler")

	result := common.Result{}
	for {
		err := r.ParseMultipartForm(0)
		if err != nil {
			log.Print("paseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		uid := -1
		id := r.FormValue("account-id")
		if len(id) > 0 {
			uid, err = strconv.Atoi(id)
			if err != nil {
				log.Print("paseform failed")

				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
			}
		}
		account := r.FormValue("account-account")
		email := r.FormValue("account-email")
		groups := r.MultipartForm.Value["account-group"]
		groupList := []int{}
		for _, g := range groups {
			gid, err := strconv.Atoi(g)
			if err != nil {
				log.Printf("parse group id failed, group:%s", g)

				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
			}

			groupList = append(groupList, gid)
		}

		usr, found := bll.QueryUserByID(uid)
		if found {
			changeFlag := false
			// 说明是更新账号信息
			if usr.Email != email {
				usr.Email = email
				usr.Status = model.DEACTIVE
				changeFlag = true
			}
			usr.Groups = groupList

			_, ok := bll.UpdateUser(usr)
			if !ok {
				result.ErrCode = 1
				result.Reason = "保存账号信息失败"
			} else {
				result.ErrCode = 0
				result.Reason = "保存账号信息成功"
			}

			if changeFlag {
				strID, ok := commonbll.PutInCache(usr, 15) // 有效期15minute
				if ok {
					sendVerifyMail(usr.Name, usr.Email, strID)
				}
			}
		} else {
			// 新建账号
			_, ok := bll.CreateUser(account, email)
			if !ok {
				result.ErrCode = 1
				result.Reason = "创建账号失败"
			} else {
				result.ErrCode = 0
				result.Reason = "创建账号成功"

				usr, _ := bll.QueryUserByAccount(account)
				strID, ok := commonbll.PutInCache(usr, 15) // 有效期15minute
				if ok {
					sendVerifyMail(account, email, strID)
				}
			}
		}
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// CheckAccountActionHandler 检查账号是否可用处理器
func CheckAccountActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("CheckAccountActionHandler")

	result := common.Result{}

	params := util.SplitParam(r.URL.RawQuery)
	for true {
		account, found := params["account"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}
		_, found = bll.QueryUserByAccount(account)
		if !found {
			result.ErrCode = 0
			result.Reason = "该账号可用"
			break
		}

		result.ErrCode = 1
		result.Reason = "该账号不可用"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
