package auth

import (
	"net/http"
	"html/template"
	"encoding/json"
	"log"
	"time"
	"strings"
	"strconv"	
	"webcenter/session"
	"webcenter/application"
)

func init() {
	registerRouter()
}

func registerRouter() {
	application.RegisterGetHandler("/auth/login/", loginHandler)
	application.RegisterGetHandler("/auth/logout/", logoutHandler)
    application.RegisterPostHandler("/auth/verify/", verifyHandler)
    
    application.RegisterPostHandler("/account/admin/queryAllInfo/", queryAllAccountInfoHandler)
	application.RegisterGetHandler("/account/admin/queryAllUser/", queryAllUserHandler)
	application.RegisterGetHandler("/account/admin/queryUser/", queryUserHandler)
	application.RegisterPostHandler("/account/admin/deleteUser/", deleteUserHandler)
    application.RegisterPostHandler("/account/user/ajax/", ajaxUserHandler)
    
	application.RegisterGetHandler("/account/admin/queryAllGroup/", queryAllGroupHandler)
	application.RegisterGetHandler("/account/admin/queryGroup/", queryGroupHandler)
	application.RegisterPostHandler("/account/admin/deleteGroup/", deleteGroupHandler)    
    application.RegisterPostHandler("/account/group/ajax/", ajaxGroupHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("loginHandler");
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	session := session.GetSession(w,r)	
    t, err := template.ParseFiles("template/html/auth/login.html")
    if (err != nil) {
        log.Print(err)
        
        http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    _, found := session.GetAccount()
    if found {
        http.Redirect(w, r, "/admin/", http.StatusFound)
        return
    }
 
    controller := &loginController{}
    view := controller.Action(session)
        
    view.Accesscode = session.AccessToken()
    
    t.Execute(w, view)	
}


func logoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutHandler");
	
	session := session.GetSession(w,r)
	session.ResetAccount()
	
    http.Redirect(w, r, "/", http.StatusFound)
}


func verifyHandler(w http.ResponseWriter, r *http.Request) {	
	log.Print("verifyHandler");
	
	result := VerifyResult{}
	
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
    	}
		
		session := session.GetSession(w,r)
		param := VerifyParam{}
		param.account = r.FormValue("login_account")
		param.password = r.FormValue("login_password")
		param.accesscode = r.FormValue("accesscode")
		param.session =  session
	
		log.Printf("account:%s,password:%s,accesscode:%s", param.account, param.password, param.accesscode)
		
    	controller := &verifyController{}
    	result = controller.Action(&param)
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


func queryAllAccountInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllAccountInfoHandler");
	
	result := GetAllAccountInfoResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := GetAllAccountInfoParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
	    accessCode := r.FormValue("accesscode")
		param.accessCode = accessCode
    	param.session = session

    	controller := &accountController{}
    	result = controller.getAllAccountInfoAction(param)
    	
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


func queryAllUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllUserHandler");
	
	result := GetAllUserResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := GetAllUserParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	    	
	    accessCode := r.FormValue("accesscode")
		param.accessCode = accessCode
    	param.session = session

    	controller := &accountController{}
    	result = controller.getAllUserAction(param)
    	
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

func ajaxUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxUserHandler");
	
	result := SubmitUserResult{}
	for true {
		session := session.GetSession(w,r)
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
	    param.session = session
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

func queryUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryUserHandler");
		
	result := GetUserResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := GetUserParam{}
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
		param.session = session
		
		log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
		 
	    controller := &accountController{}
	    result = controller.getUserAction(param)
    	
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


func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteArticleHandler");
	
	result := DeleteUserResult{}
	
	for true {
		session := session.GetSession(w,r)
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
		param.session = session
		
		log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
		 
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

func queryAllGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllGroupHandler");
	
	
	result := GetAllGroupResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := GetAllGroupParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	    	
	    accessCode := r.FormValue("accesscode")
		param.accessCode = accessCode
    	param.session = session

    	controller := &accountController{}
    	result = controller.getAllGroupAction(param)
    	
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


func ajaxGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxGroupHandler");
	
	result := SubmitGroupResult{}
	for true {
		session := session.GetSession(w,r)
		param := SubmitGroupParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
    	id := r.FormValue("group-id")
		name := r.FormValue("group-name")
		pid := r.FormValue("group-parent")

		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Print("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    
		param.parent, err = strconv.Atoi(pid)
	    if err != nil {
	    	log.Print("parse group pid failed, group:%s", pid)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    
	    param.name = name
    	param.submitDate = time.Now().Format("2006-01-02 15:04:05")
    	param.session = session
	    
	    controller := &accountController{}
	    result = controller.submitGroupAction(param)
	    
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

func queryGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryGroupHandler");
	
	result := GetGroupResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := GetGroupParam{}
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
		param.session = session
		
		log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
		 
	    controller := &accountController{}
	    result = controller.getGroupAction(param)
    	
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


func deleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteGroupHandler");
	
	result := DeleteGroupResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := DeleteGroupParam{}
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
		param.session = session
		
		log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
		 
	    controller := &accountController{}
	    result = controller.deleteGroupAction(param)
    	
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
