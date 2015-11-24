package access

import (
	"net/http"
	"html/template"
	"encoding/json"
	"log"
    "webcenter/session"    
)

type LoginView struct {
	Accesscode string
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
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
 
    view := LoginView{}        
    view.Accesscode = session.AccessToken()
    
    t.Execute(w, view)
}


func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutHandler");
	
	session := session.GetSession(w,r)
	session.ResetAccount()
	
    http.Redirect(w, r, "/", http.StatusFound)
}


func VerifyHandler(w http.ResponseWriter, r *http.Request) {	
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
		
    	controller := &accessController{}
    	result = controller.VerifyAction(&param)
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

