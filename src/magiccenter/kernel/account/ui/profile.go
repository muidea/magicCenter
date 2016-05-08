package ui

import (
	"log"
	"net/http"
	"encoding/json"
	"html/template"
	"muidea.com/util"
	"magiccenter/cache"
	"magiccenter/kernel/common"
    "magiccenter/kernel/account/model"
    "magiccenter/kernel/account/bll"

)

type UserProfileView struct {
	Users []model.UserDetailView
	Groups []model.GroupInfo
}

type VerifyUserView struct {
	Id string
	User *model.UserDetail
}

type AjaxUserVerifyResult struct {
	common.Result
}


func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageUserHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/account/user.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
    view := UserProfileView{}
    view.Users = bll.QueryAllUser()
    view.Groups = bll.QueryAllGroupInfo()
    
    t.Execute(w, view)	
}

func UserVerifyHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("UserVerifyHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/user/verify.html")
    if (err != nil) {
    	panic("parse files failed");
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

func AjaxVerifyHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxVerifyHandler");
	
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
    	if bll.CreateUser(userDetail.Account, passWord, nickName, userDetail.Email,model.ACTIVE, userDetail.Groups) {
    		result.ErrCode = 0
    		result.Reason = "激活用户成功"
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
        return
    }
    
    w.Write(b)
}



