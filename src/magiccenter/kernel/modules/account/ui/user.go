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
	for true {
		id, found := params["id"]
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

	t.Execute(w, result)
}

// SaveUserActionHandler 保存用户信息处理器
func SaveUserActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("SaveUserActionHandler")

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
		account := r.FormValue("user-account")
		nickName := r.FormValue("user-nickname")
		passWord := r.FormValue("user-password")
		email := r.FormValue("user-email")
		groups := r.MultipartForm.Value["user-group"]
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
			// 说明是更新用户信息
			usr.Account = account
			usr.Name = nickName
			usr.Email = email
			if len(groupList) > 0 {
				usr.Groups = groupList
			}
			ok := bll.SaveUser(usr)
			if !ok {
				result.ErrCode = 1
				result.Reason = "保存用户信息失败"
				break
			} else {
				result.ErrCode = 0
				result.Reason = "保存用户信息成功"
			}
		} else {
			ok := bll.CreateUser(account, passWord, nickName, email, model.NEW, groupList)
			if !ok {
				result.ErrCode = 1
				result.Reason = "创建用户失败"
			} else {
				usr, ok = bll.QueryUserByAccount(account)
				if ok {
					result.ErrCode = 0
					result.Reason = "创建用户成功"
				} else {
					result.ErrCode = 1
					result.Reason = "创建用户失败"
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

// UserProfileViewHandler 个人空间页面处理器
//
func UserProfileViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageUserHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/user/profile.html")
	if err != nil {
		panic("parse files failed")
	}

	view := UserProfileView{}

	t.Execute(w, view)
}
