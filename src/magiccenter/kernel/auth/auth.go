package auth

import (
	"log"
	"net/http"
	"encoding/json"
	"html/template"
	"magiccenter/router"
	"magiccenter/kernel/common"
    "magiccenter/kernel/account/bll"
)


type VerifyAuthResult struct {
	common.Result
}

func RegisterRouter() {
	router.AddGetRoute("/auth/login/", LoginHandler)
	router.AddGetRoute("/auth/logout/", LogoutHandler)
    router.AddPostRoute("/auth/verify/", VerifyAuthHandler)    
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("loginHandler");
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/auth/login.html")
    if (err != nil) {
    	panic("parse files failed");
    }
        
    t.Execute(w, nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutHandler");
	
	//session := session.GetSession(w,r)
	//session.ResetAccount()
	
    http.Redirect(w, r, "/", http.StatusFound)
}

func VerifyAuthHandler(w http.ResponseWriter, r *http.Request) {	
	log.Print("VerifyAuthHandler");
	
	result := VerifyAuthResult{}
	
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
    	}
		
		account := r.FormValue("login_account")
		password := r.FormValue("login_password")
		
		_, found := bll.VerifyUserByAccount(account, password)
		if !found {
    		result.ErrCode = 1
    		result.Reason = "无效账号"			
		}
				
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
