package ui

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"magiccenter/cache"
	"magiccenter/configuration"
	"magiccenter/kernel/account/bll"
	"magiccenter/kernel/account/model"
	"magiccenter/kernel/common"
	"magiccenter/mail"
	"net/http"
	"strconv"

	"muidea.com/util"
)

type ManageUserView struct {
	Users  []model.UserDetailView
	Groups []model.GroupInfo
}

type QueryAllUserResult struct {
	common.Result
	Users []model.UserDetailView
}

type QueryUserResult struct {
	common.Result
	User model.UserDetail
}

type CheckAccountResult struct {
	common.Result
}

type CreateUserResult struct {
	common.Result
	Users []model.UserDetailView
}

type DeleteUserResult struct {
	CreateUserResult
}

func ManageUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageUserHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/account/user.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ManageUserView{}
	view.Users = bll.QueryAllUser()
	view.Groups = bll.QueryAllGroupInfo()

	t.Execute(w, view)
}

func QueryAllUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllUserHandler")

	result := QueryAllUserResult{}
	result.Users = bll.QueryAllUser()
	result.ErrCode = 0
	result.Reason = "查询成功"

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

func QueryUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryUserHandler")

	result := QueryUserResult{}

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

		result.User, found = bll.QueryUserById(uid)
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

func CheckAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("CheckAccountHandler")

	result := CheckAccountResult{}

	for true {
		err := r.ParseForm()
		if err != nil {
			log.Print("paseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		account := r.FormValue("user-account")

		_, found := bll.QueryUserByAccount(account)
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

func sendVerifyMail(user, email, id string) {
	systemInfo := configuration.GetSystemInfo()

	subject := "MagicCenter用户验证"

	content := fmt.Sprintf("<html><head><title>用户信息验证</title></head><body><p>Hi %s</p><p><a href='http://%s/user/verify/?id=%s'>请点击链接继续验证用户信息</a></p><p>该邮件由MagicCenter自动发送，请勿回复该邮件</p></body></html>", user, systemInfo.Domain, id)

	mail.PostMail(email, subject, content)
}

func AjaxUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxUserHandler")

	result := CreateUserResult{}
	for true {
		err := r.ParseMultipartForm(0)
		if err != nil {
			log.Print("paseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		id, err := strconv.Atoi(r.FormValue("user-id"))
		if err != nil {
			log.Print("paseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}
		account := r.FormValue("user-account")
		email := r.FormValue("user-email")
		groups := r.MultipartForm.Value["user-group"]
		groupList := []int{}
		for _, g := range groups {
			gid, err := strconv.Atoi(g)
			if err != nil {
				log.Print("parse group id failed, group:%s", g)

				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
			}

			groupList = append(groupList, gid)
		}

		usr, found := bll.QueryUserById(id)
		if found {
			// 说明是更新用户信息
			usr.Account = account
			usr.Email = email
			usr.Groups = groupList
			ok := bll.SaveUser(usr)
			if !ok {
				result.ErrCode = 1
				result.Reason = "保存用户信息失败"
				break
			}
		}

		result.Users = bll.QueryAllUser()
		result.ErrCode = 0
		result.Reason = "保存用户信息成功"

		// 如果是新建用户或者用户的Mail变化了，则还需要发送验证邮件
		if !found || usr.Email != email {
			// 说明是新建用户，新建用户临时信息需要保存到Cache中
			usr := &model.UserDetail{}
			usr.Id = id
			usr.Account = account
			usr.Email = email
			usr.Groups = groupList

			cache, found := cache.GetCache()
			if found {
				id := cache.PutIn(usr, 15)

				sendVerifyMail(account, email, id)
			} else {
				result.ErrCode = 1
				result.Reason = "保存用户信息失败"
				break
			}
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
		return
	}

	w.Write(b)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteUserHandler")

	result := DeleteUserResult{}
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
			result.Reason = "删除分组失败"
			break
		}

		result.ErrCode = 0
		result.Reason = "查询成功"
		result.Users = bll.QueryAllUser()
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
