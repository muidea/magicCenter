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
)

func init() {
	registerRouter()
}

func registerRouter() {
	http.HandleFunc("/auth/login/",loginHandler)
	http.HandleFunc("/auth/logout/",logoutHandler)
    http.HandleFunc("/auth/verify/",verifyHandler)
    
    http.HandleFunc("/account/admin/queryAllInfo/", queryAllAccountInfoHandler)
	http.HandleFunc("/account/admin/queryAllUser/", queryAllUserHandler)
	http.HandleFunc("/account/admin/queryUser/", queryUserHandler)
	http.HandleFunc("/account/admin/deleteUser/", deleteUserHandler)
    http.HandleFunc("/account/user/ajax/", ajaxUserHandler)
    
	http.HandleFunc("/account/admin/queryAllGroup/", queryAllGroupHandler)
	http.HandleFunc("/account/admin/queryGroup/", queryGroupHandler)
	http.HandleFunc("/account/admin/deleteGroup/", deleteGroupHandler)    
    http.HandleFunc("/account/group/ajax/",ajaxGroupHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("loginHandler");
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	session := session.GetSession(w,r)	
    t, err := template.ParseFiles("template/html/auth/login.html")
    if (err != nil) {
        log.Fatal(err)
        
        http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    _, found := session.GetOption(AccountSessionKey)
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
	session.RemoveOption(AccountSessionKey)
	
    http.Redirect(w, r, "/", http.StatusFound)
}


func verifyHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("verifyHandler");
    err := r.ParseForm()
    if err != nil {
    	log.Fatal("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }

	param := VerifyParam{}
	param.account = r.FormValue("login_account")
	param.password = r.FormValue("login_password")
	param.accesscode = r.FormValue("accesscode")
	
	log.Printf("account:%s,password:%s,accesscode:%s", param.account, param.password, param.accesscode)
		
	session := session.GetSession(w,r)
	
    controller := &verifyController{}
    result := controller.Action(&param, session)
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}
	 
    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }

    accessCode := r.FormValue("accesscode")
    
	param := GetAllAccountInfoParam{}
	param.accessCode = accessCode
    param.session = session

    controller := &accountController{}
    result := controller.getAllAccountInfoAction(param)
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}
	 
    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }

    accessCode := r.FormValue("accesscode")
    
	param := GetAllUserParam{}
	param.accessCode = accessCode
    param.session = session

    controller := &accountController{}
    result := controller.getAllUserAction(param)
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}

    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
	
	id := r.FormValue("user-id")
	name := r.FormValue("user-account")
	password := r.FormValue("user-password")
	nickname := r.FormValue("user-nickname")
	email := r.FormValue("user-email")
	group := r.FormValue("user-group")
	
	param := SubmitUserParam{}
	param.id, err = strconv.Atoi(id)
    if err != nil {
    	log.Print("parse id failed, id:%s", id)
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }	
	param.group, err = strconv.Atoi(group)
    if err != nil {
    	log.Print("parse group failed, group:%s", group)
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    param.account = name
    param.password = password
    param.nickname = nickname
    param.email = email    
    param.submitDate = time.Now().Format("2006-01-02 15:04:05")

    controller := &accountController{}
    result := controller.submitUserAction(param)
    
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}

    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }

	var id = ""
	idInfo := r.URL.RawQuery
	if len(idInfo) > 0 {
		parts := strings.Split(idInfo,"=")
		if len(parts) == 2 {
			id = parts[1]
		}
	}
	
	param := GetUserParam{}
	accessCode := r.FormValue("accesscode")
	param.id, err = strconv.Atoi(id)
    if err != nil {
    	log.Printf("convert id failed, id:%s,accessCode:%s", id, accessCode)
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
	param.accessCode = accessCode
	param.session = session
	
	log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
	 
    controller := &accountController{}
    result := controller.getUserAction(param)
    
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}

    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }

	var id = ""
	idInfo := r.URL.RawQuery
	if len(idInfo) > 0 {
		parts := strings.Split(idInfo,"=")
		if len(parts) == 2 {
			id = parts[1]
		}
	}
	
	param := DeleteUserParam{}
	accessCode := r.FormValue("accesscode")
	param.id, err = strconv.Atoi(id)
    if err != nil {
    	log.Printf("convert id failed, id:%s,accessCode:%s", id, accessCode)
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
	param.accessCode = accessCode
	param.session = session
	
	log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
	
    controller := &accountController{}
    result := controller.deleteUserAction(param)
    
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}
	 
    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    accessCode := r.FormValue("accesscode")
    log.Printf("accessCode:%s",accessCode);
	 
	param := GetAllGroupParam{}
	param.accessCode = accessCode
	param.session = session
    controller := &accountController{}
    result := controller.getAllGroupAction(param)
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}

    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
	
	id := r.FormValue("group-id")
	name := r.FormValue("group-name")
	pid := r.FormValue("group-parent")
	
	log.Printf("id:%d,name:%s,pid:%d",id,name,pid)
	
	param := SubmitGroupParam{}
	param.id, err = strconv.Atoi(id)
    if err != nil {
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    param.parent, err = strconv.Atoi(pid)
    if err != nil {
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    param.name = name
    param.submitDate = time.Now().Format("2006-01-02 15:04:05")

    controller := &accountController{}
    result := controller.submitGroupAction(param)
    
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}
	 
	var id = ""
	idInfo := r.URL.RawQuery
	if len(idInfo) > 0 {
		parts := strings.Split(idInfo,"=")
		if len(parts) == 2 {
			id = parts[1]
		}
	}
	
	param := GetGroupParam{}
	accessCode := r.FormValue("accesscode")
	param.id, err = strconv.Atoi(id)
    if err != nil {
    	log.Printf("convert id failed, id:%s,accessCode:%s", id, accessCode)
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
	param.accessCode = accessCode
	param.session = session	 
	 
    controller := &accountController{}
    result := controller.getGroupAction(param)
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}

    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }

	var id = ""
	idInfo := r.URL.RawQuery
	if len(idInfo) > 0 {
		parts := strings.Split(idInfo,"=")
		if len(parts) == 2 {
			id = parts[1]
		}
	}
	
	param := DeleteGroupParam{}
	accessCode := r.FormValue("accesscode")
	param.id, err = strconv.Atoi(id)
    if err != nil {
    	log.Printf("convert id failed, id:%s,accessCode:%s", id, accessCode)
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
	param.accessCode = accessCode
	param.session = session
	
	log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
	
    controller := &accountController{}
    result := controller.deleteGroupAction(param)
    
    b, err := json.Marshal(result)
    if err != nil {
    	log.Fatal("json marshal failed, err:" + err.Error())
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    w.Write(b)
}
