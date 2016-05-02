package ui

import (
	"log"
	"net/http"
	"html/template"
	"muidea.com/util"
	"magiccenter/cache"
    "magiccenter/kernel/account/model"
    "magiccenter/kernel/account/bll"

)

type UserProfileView struct {
	Users []model.UserDetailView
	Groups []model.GroupInfo
}

type VerifyUserView struct {
	User *model.UserDetail
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
    	result = true
    	break
    }
    
    if !result {
    	http.Redirect(w, r, "/", http.StatusFound)
    }

    t.Execute(w, view)
}



