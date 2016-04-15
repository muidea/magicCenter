package auth

import (
	"log"
	"net/http"
	"encoding/json"
	"html/template"
	"martini"
	"magiccenter/router"
	"magiccenter/kernel/common"
    "magiccenter/kernel/account/bll"
    "magiccenter/session"
    "magiccenter/configuration"
)

type VerifyAuthResult struct {
	common.Result
	RedirectUrl string
}

func Auth() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
		
		authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
		if !found {
			panic("unexpected, can't fetch authorith id")
		}
		
		uri := req.URL.Path
		session := session.GetSession(res, req)		
		_, found = session.GetOption(authId)
		if !found {
			if uri != "/auth/login/" && uri != "/auth/verify/" && uri != "/auth/logout/" {
				http.Redirect(res, req, "/auth/login/", http.StatusFound)
				return 
			}
		}
		
		c.Next()
	}
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
	
	authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
	if !found {
		panic("unexpected, can't fetch authorith id")
	}
			
	session := session.GetSession(w,r)
	session.RemoveOption(authId)
	
    http.Redirect(w, r, "/", http.StatusFound)
}

func VerifyAuthHandler(w http.ResponseWriter, r *http.Request) {	
	log.Print("VerifyAuthHandler");
	
	result := VerifyAuthResult{}
	
	authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
	if !found {
		panic("unexpected, can't fetch authorith id")
	}
			
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
    	}
		
		account := r.FormValue("login_account")
		password := r.FormValue("login_password")
		
		user, found := bll.VerifyUserByAccount(account, password)
		if !found {
    		result.ErrCode = 1
    		result.Reason = "无效账号"
    		
    		break;			
		}
		
		session := session.GetSession(w, r)
		session.SetOption(authId, user)
		
		result.ErrCode = 0
		result.Reason = "登陆成功"
		result.RedirectUrl = "/admin/"			
    	break
	}

    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, failed, err:" + err.Error())
    }
    
    w.Write(b)
}
