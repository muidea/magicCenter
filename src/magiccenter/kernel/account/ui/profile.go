package ui

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/cache"
	"magiccenter/configuration"
	"magiccenter/kernel/account/bll"
	"magiccenter/kernel/account/model"
	"magiccenter/kernel/common"
	"magiccenter/session"
	"net/http"

	"muidea.com/util"
)

type UserProfileView struct {
	Users  []model.UserDetailView
	Groups []model.GroupInfo
}

type VerifyUserView struct {
	Id   string
	User *model.UserDetail
}

type AjaxUserVerifyResult struct {
	common.Result
	RedirectUrl string
}

//
// 个人空间页面
//
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageUserHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/user/profile.html")
	if err != nil {
		panic("parse files failed")
	}

	view := UserProfileView{}
	view.Users = bll.QueryAllUser()
	view.Groups = bll.QueryAllGroupInfo()

	t.Execute(w, view)
}

//
// 获取校验用户信息页面，数据合法后返回提交用户信息页面
// 以便用户输入账号密码信息
//
func UserVerifyHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("UserVerifyHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/user/verify.html")
	if err != nil {
		panic("parse files failed")
	}

	result := false
	view := VerifyUserView{}
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			break
		}

		cache, found := cache.GetCache()
		if !found {
			panic("can't get cache")
		}

		user, found := cache.FetchOut(id)
		if !found {
			log.Printf("can't fetchout user, id:%s", id)
			break
		}

		view.User = user.(*model.UserDetail)
		view.Id = id

		result = true
		break
	}

	if !result {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	t.Execute(w, view)
}

//
// 处理用户信息页面请求，数据处理完成后激活该用户
func AjaxVerifyHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxVerifyHandler")

	result := AjaxUserVerifyResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			log.Print("paseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		id := r.FormValue("user-id")
		nickName := r.FormValue("user-nickname")
		passWord := r.FormValue("user-password")

		cache, found := cache.GetCache()
		if !found {
			panic("can't get cache")
		}

		user, found := cache.FetchOut(id)
		if !found {
			log.Printf("can't fetchout user, id:%s", id)

			result.ErrCode = 1
			result.Reason = "用户信息不存在"

			break
		}

		userDetail := user.(*model.UserDetail)
		if bll.CreateUser(userDetail.Account, passWord, nickName, userDetail.Email, model.ACTIVE, userDetail.Groups) {
			authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
			if !found {
				panic("unexpected, can't fetch authorith id")
			}
			session := session.GetSession(w, r)
			session.SetOption(authId, *userDetail)

			result.ErrCode = 0
			result.Reason = "激活用户成功"
			result.RedirectUrl = "/user/profile/"
		} else {
			result.ErrCode = 1
			result.Reason = "激活用户失败"
		}

		cache.Remove(id)

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
