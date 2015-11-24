package group

import (
	"net/http"
	"encoding/json"
	"log"
	"time"
	"strings"
	"strconv"	
)

func QueryAllGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllGroupHandler");
	
	
	result := GetAllGroupResult{}
	
	for true {
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


func QueryAllSubGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllSubGroupHandler");
	
	result := GetAllSubGroupResult{}
	
	for true {
		param := GetAllSubGroupParam{}
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
    	result = controller.getAllSubGroupAction(param)
    	
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


func QueryGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryGroupHandler");
	
	result := GetGroupResult{}
	
	for true {
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

func AjaxGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxGroupHandler");
	
	result := SubmitGroupResult{}
	for true {
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


func DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteGroupHandler");
	
	result := DeleteGroupResult{}
	
	for true {
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
