package account

import (
	"net/http"
	"encoding/json"
	"html/template"
	"log"
	"time"
	"strings"
	"strconv"
	"webcenter/session"
	"webcenter/auth/group"
)

type ManageView struct {
	Accesscode string
	UserInfo []UserInfo
	GroupInfo []group.GroupInfo
}

func ManageUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageUserHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	session := session.GetSession(w,r)
    t, err := template.ParseFiles("template/html/admin/account/user.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
	controller := &accountController{}
	info := controller.queryManageInfoAction()
    
    view := ManageView{}
    view.Accesscode = session.AccessToken()
    view.UserInfo = info.UserInfo
    view.GroupInfo = info.GroupInfo
    
    t.Execute(w, view)
}

func VerifyAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("VerifyAccountHandler");
	
	result := VerifyAccountResult{}
	for true {
		param := VerifyAccountParam{}
	    err := r.ParseForm()
    	if err != nil {    		
    		panic("paseform failed, err:" + err.Error())
    	}
    	
		account := r.FormValue("user-account")
		accessCode := r.FormValue("accesscode")
				
	    param.account = account
	    param.accessCode = accessCode
	    
	    controller := &accountController{}
	    result = controller.verifyAccountAction(param)
	    
	    break
	}
    
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, err:" + err.Error())
    }
    
    w.Write(b)	
}

func QueryAllUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllUserHandler");
	
	result := QueryAllUserResult{}
	
	for true {
		param := QueryAllUserParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	    	
	    accessCode := r.FormValue("accesscode")
		param.accessCode = accessCode

    	controller := &accountController{}
    	result = controller.queryAllUserAction(param)
    	
    	break
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	log.Fatal("json marshal failed, err:" + err.Error())
    	
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    w.Write(b)
}

func AjaxUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxUserHandler");
	
	result := SubmitUserResult{}
	for true {
		param := SubmitUserParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		id := r.FormValue("user-id")
		name := r.FormValue("user-account")
		password := r.FormValue("user-password")
		nickname := r.FormValue("user-nickname")
		email := r.FormValue("user-email")
		group := r.FormValue("user-group")
		accessCode := r.FormValue("accesscode")
		
		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Print("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }	
		param.group, err = strconv.Atoi(group)
	    if err != nil {
	    	log.Print("parse group failed, group:%s", group)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    param.account = name
	    param.password = password
	    param.nickname = nickname
	    param.email = email    
	    param.submitDate = time.Now().Format("2006-01-02 15:04:05")
	    param.accessCode = accessCode
	    
	    controller := &accountController{}
	    result = controller.submitUserAction(param)
	    
	    break
	}
    
    b, err := json.Marshal(result)
    if err != nil {
    	log.Fatal("json marshal failed, err:" + err.Error())
    	
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    w.Write(b)	
}

func QueryUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryUserHandler");
		
	result := QueryUserResult{}
	
	for true {
		param := QueryUserParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}

		var id = ""
		idInfo := r.URL.RawQuery
		if len(idInfo) > 0 {
			parts := strings.Split(idInfo,"=")
			if len(parts) == 2 {
				id = parts[1]
			}
		}
		
		accessCode := r.FormValue("accesscode")
		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Printf("convert id failed, id:%s,accessCode:%s", id, accessCode)
	    	
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    
		param.accessCode = accessCode
		
	    controller := &accountController{}
	    result = controller.queryUserAction(param)
    	
    	break
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	log.Fatal("json marshal failed, err:" + err.Error())
    	
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    w.Write(b)
}


func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteArticleHandler");
	
	result := DeleteUserResult{}
	
	for true {
		param := DeleteUserParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}

		var id = ""
		idInfo := r.URL.RawQuery
		if len(idInfo) > 0 {
			parts := strings.Split(idInfo,"=")
			if len(parts) == 2 {
				id = parts[1]
			}
		}
		
		accessCode := r.FormValue("accesscode")
		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Printf("convert id failed, id:%s,accessCode:%s", id, accessCode)
	    	
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    
		param.accessCode = accessCode
		
	    controller := &accountController{}
	    result = controller.deleteUserAction(param)
    	
    	break
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	log.Fatal("json marshal failed, err:" + err.Error())
    	
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    w.Write(b)
}


func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("EditUserHandler");
		
	result := QueryUserResult{}
	
	for true {
		param := QueryUserParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}

		var id = ""
		idInfo := r.URL.RawQuery
		if len(idInfo) > 0 {
			parts := strings.Split(idInfo,"=")
			if len(parts) == 2 {
				id = parts[1]
			}
		}
		
		accessCode := r.FormValue("accesscode")
		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Printf("convert id failed, id:%s,accessCode:%s", id, accessCode)
	    	
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    
		param.accessCode = accessCode
		
	    controller := &accountController{}
	    result = controller.queryUserAction(param)
    	
    	break
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	log.Fatal("json marshal failed, err:" + err.Error())
    	
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    w.Write(b)
}


