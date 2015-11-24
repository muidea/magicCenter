package account

import (
	"net/http"
	"encoding/json"
	"log"
	"time"
	"strings"
	"strconv"	
)

func QueryAllUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllUserHandler");
	
	result := GetAllUserResult{}
	
	for true {
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

func AjaxUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxUserHandler");
	
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
		
	result := GetUserResult{}
	
	for true {
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
