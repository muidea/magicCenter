package group

import (
	"net/http"
	"encoding/json"
	"html/template"
	"log"
	"strings"
	"strconv"
	"webcenter/session"	
	"webcenter/common"
)

type ManageView struct {
	Accesscode string
	GroupInfo []GroupInfo
}

type EditView struct {
	common.Result
	Accesscode string
	Id int
	Name string
	Catalog int
}

func ManageGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageGroupHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	session := session.GetSession(w,r)
    t, err := template.ParseFiles("template/html/admin/account/group.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
	controller := &accountController{}
	info := controller.queryManageInfoAction()
    
    view := ManageView{}
    view.Accesscode = session.AccessToken()
    view.GroupInfo = info.GroupInfo
    
    t.Execute(w, view)
}

func QueryAllGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllGroupHandler");
	
	
	result := QueryAllGroupResult{}
	
	for true {
		param := QueryAllGroupParam{}
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
    	result = controller.queryAllGroupAction(param)
    	
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
	
	result := QueryGroupResult{}
	
	for true {
		param := QueryGroupParam{}
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
	    result = controller.queryGroupAction(param)
    	
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
	    err := r.ParseMultipartForm(0)
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
    	id := r.FormValue("group-id")
		name := r.FormValue("group-name")

		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Printf("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    
	    param.name = name	    
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

func EditGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("EditGroupHandler");
	
	result := EditView{}
	
	for true {
		param := QueryGroupParam{}
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
	    group := controller.queryGroupAction(param)
    	
    	result.ErrCode = group.ErrCode
    	result.Reason = group.Reason
    	result.Id = group.Group.Id
    	result.Name = group.Group.Name
    	result.Catalog = group.Group.Catalog
    	
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

