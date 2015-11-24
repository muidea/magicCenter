package link

import (
	"net/http"
	"encoding/json"
	"html/template"
	"strings"
	"log"
	"strconv"
)

func ManageLinkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/content/Link.html")
    if (err != nil) {
        log.Print(err)
        
        http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
        
    t.Execute(w, nil)
}


func QueryAllLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllLinkHandler");
	
	result := QueryAllLinkResult{}
	
	for true {
		param := QueryAllLinkParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
	    accessCode := r.FormValue("accesscode")
		param.accessCode = accessCode

    	controller := &linkController{}
    	result = controller.queryAllLinkAction(param)
    	
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

func QueryLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryLinkHandler");
	
	result := QueryLinkResult{}
	
	for true {
		param := QueryLinkParam{}
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
		    	
    	controller := &linkController{}
    	result = controller.queryLinkAction(param)
    	
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

func DeleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteLinkHandler");
	
	result := DeleteLinkResult{}
	
	for true {
		param := DeleteLinkParam{}
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
		 
	    controller := &linkController{}
	    result = controller.deleteLinkAction(param)
    	
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

func AjaxLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxLinkHandler");
	
	result := SubmitLinkResult{}
	
	for true {
		param := SubmitLinkParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		id := r.FormValue("link-id")
		name := r.FormValue("link-name")
		url := r.FormValue("link-url")
		logo := r.FormValue("link-logo")
		style := r.FormValue("link-style")
	    accessCode := r.FormValue("accesscode")
	    
		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Print("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    param.name = name
	    param.url = url
	    param.logo = logo
		param.style, err = strconv.Atoi(style)
	    if err != nil {
	    	log.Print("parse style failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }

		param.accessCode = accessCode

    	controller := &linkController{}
    	result = controller.submitLinkAction(param)
    	
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

func EditLinkHandler(w http.ResponseWriter, r *http.Request) {
	
}

