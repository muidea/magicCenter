package admin

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/configuration"
	"magiccenter/kernel/account/bll"
	"magiccenter/kernel/account/model"
	"magiccenter/kernel/auth"
	"magiccenter/kernel/common"
	"magiccenter/router"
	"magiccenter/session"
	"net/http"
)

// 后台管理主页面
// 包含当前登陆的用户信息
type AdminView struct {
	User model.User
}

type VerifyAuthResult struct {
	common.Result
	RedirectUrl string
}

func RegisterRouter() {
	router.AddGetRoute("/admin/", AdminHandler, auth.AdminAuthVerify())

	router.AddGetRoute("/admin/login/", LoginHandler, nil)
	router.AddGetRoute("/admin/logout/", LogoutHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/verify/", VerifyAuthHandler, nil)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("adminHandler")

	authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
	if !found {
		panic("unexpected, can't fetch authorith id")
	}

	session := session.GetSession(w, r)
	user, found := session.GetOption(authId)
	if !found {
		panic("unexpected, must login system first.")
	}

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("resources/view/admin/index.html")
	if err != nil {
		panic("parse file failed")
	}

	view := AdminView{}
	view.User.Id = user.(model.UserDetail).Id
	view.User.Name = user.(model.UserDetail).Name
	t.Execute(w, view)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("loginHandler")
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
	if !found {
		panic("unexpected, can't fetch authorith id")
	}

	session := session.GetSession(w, r)
	_, found = session.GetOption(authId)
	if found {
		http.Redirect(w, r, "/admin/", http.StatusFound)
	}

	t, err := template.ParseFiles("template/html/admin/login.html")
	if err != nil {
		panic("parse files failed")
	}

	t.Execute(w, nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutHandler")

	authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
	if !found {
		panic("unexpected, can't fetch authorith id")
	}

	session := session.GetSession(w, r)
	session.RemoveOption(authId)

	http.Redirect(w, r, "/", http.StatusFound)
}

func VerifyAuthHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("VerifyAuthHandler")

	result := VerifyAuthResult{}

	authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
	if !found {
		panic("unexpected, can't fetch authorith id")
	}

	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		account := r.FormValue("login_account")
		password := r.FormValue("login_password")

		user, found := bll.VerifyUserByAccount(account, password)
		if !found {
			result.ErrCode = 1
			result.Reason = "无效账号"

			break
		}

		isAdmin := false
		for _, gid := range user.Groups {
			group, found := bll.QueryGroupById(gid)
			if found && group.AdminGroup() {
				isAdmin = true
			}
		}

		if !isAdmin {
			result.ErrCode = 1
			result.Reason = "无效账号"

			break
		}

		session := session.GetSession(w, r)
		session.SetOption(authId, user)

		result.ErrCode = 0
		result.Reason = "登陆成功"
		result.RedirectUrl = "/admin/"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
