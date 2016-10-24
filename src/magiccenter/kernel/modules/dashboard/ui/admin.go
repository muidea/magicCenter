/*
magiccenter 后台管理

1、提供框架页面显示功能
2、提供后台管理登陆页面&登陆验证
3、提供后台管理登出
4、提供管理导航功能
*/

package ui

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/common"
	"magiccenter/common/model"
	"magiccenter/kernel/modules/account/bll"
	"magiccenter/system"
	"net/http"
)

// AdminView 管理页面视图
// 后台管理主页面信息
// 包含当前登陆的用户信息
type AdminView struct {
	User model.User
}

// VerifyAuthResult 校验结果
// 登陆结果
// 会话Token
type VerifyAuthResult struct {
	common.Result
	Token string
}

// AdminViewHandler 后台管理主页面处理器
func AdminViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("adminViewHandler")

	session := system.GetSession(w, r)
	user, found := session.GetAccount()
	if !found {
		panic("unexpected, must login system first.")
	}

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/index.html")
	if err != nil {
		panic("parse file failed")
	}

	view := AdminView{}
	view.User.ID = user.ID
	view.User.Name = user.Name
	t.Execute(w, view)
}

// LoginViewHandler 后台管理登陆页面处理器
func LoginViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("loginViewHandler")
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	session := system.GetSession(w, r)
	_, found := session.GetAccount()
	if found {
		http.Redirect(w, r, "/admin/", http.StatusFound)
	}

	t, err := template.ParseFiles("template/html/admin/login.html")
	if err != nil {
		panic("parse files failed")
	}

	t.Execute(w, nil)
}

// VerifyAuthActionHandler 校验账号信息处理器
func VerifyAuthActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("VerifyAuthActionHandler")

	result := VerifyAuthResult{}

	for {
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
			group, found := bll.QueryGroupByID(gid)
			if found && group.AdminGroup() {
				isAdmin = true
			}
		}

		if !isAdmin {
			result.ErrCode = 1
			result.Reason = "无效账号"

			break
		}

		session := system.GetSession(w, r)
		session.SetAccount(user)

		result.ErrCode = 0
		result.Reason = "登陆成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// LogoutActionHandler 后台管理登出处理器
func LogoutActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutActionHandler")

	session := system.GetSession(w, r)
	session.ClearAccount()

	http.Redirect(w, r, "/", http.StatusFound)
}
