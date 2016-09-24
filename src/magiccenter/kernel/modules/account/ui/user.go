package ui

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	"magiccenter/common/model"
	"magiccenter/kernel/modules/account/bll"
	"net/http"
	"strconv"

	"muidea.com/util"
)

// UserProfileView 用户Profile视图
type UserProfileView struct {
	Users  []model.UserDetail
	Groups []model.Group
}

// VerifyAccountViewHandler 校验账号信息视图处理器
func VerifyAccountViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("VerifyAccountViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/user/verify.html")
	if err != nil {
		panic("parse files failed")
	}

	result := SingleUserDetail{}

	params := util.SplitParam(r.URL.RawQuery)
	id, found := params["id"]
	for true {
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		data, found := commonbll.FetchOutCache(id)
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		switch data.(type) {
		case model.UserDetail:
			result.User = data.(model.UserDetail)
			result.ErrCode = 0
		default:
			result.ErrCode = 1
			result.Reason = "无效请求数据"
		}

		//commonbll.RemoveCache(id)
		break
	}

	if result.User.Status == model.NEW || result.Fail() {
		t.Execute(w, result)

	} else if result.User.Status == model.DEACTIVE {
		result.User.Status = model.ACTIVE
		bll.SaveUser(result.User)
		commonbll.RemoveCache(id)

		http.Redirect(w, r, "/account/userProfile/", http.StatusFound)
	}
}

// UpdateUserActionHandler 保存用户信息处理器
func UpdateUserActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("UpdateUserActionHandler")

	result := common.Result{}
	for {
		err := r.ParseForm()
		if err != nil {
			log.Print("paseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		uid := -1
		id := r.FormValue("user-id")
		if len(id) > 0 {
			uid, err = strconv.Atoi(id)
			if err != nil {
				log.Print("paseform failed")

				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
			}
		}
		nickName := r.FormValue("user-nickname")
		passWord := r.FormValue("user-password")

		usr, found := bll.QueryUserByID(uid)
		if found {
			// 说明是更新用户信息
			usr.Name = nickName
			usr.Status = model.ACTIVE
			ok := bll.UpdateUserWithPassword(usr, passWord)
			if !ok {
				result.ErrCode = 1
				result.Reason = "更新用户信息失败"
				break
			} else {
				result.ErrCode = 0
				result.Reason = "更新用户信息成功"
			}
		} else {
			result.ErrCode = 1
			result.Reason = "无效用户"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// UserProfileViewHandler 个人空间页面处理器
func UserProfileViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("UserProfileViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/user/profile.html")
	if err != nil {
		panic("parse files failed")
	}

	view := UserProfileView{}

	t.Execute(w, view)
}
