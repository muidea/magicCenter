package system

import (
	"net/http"
	"html/template"
	"encoding/json"
	"log"
	"strings"
	"strconv"
	"muidea.com/util"
)

func ManageSystemHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageSystemHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/system/system.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
    view := GetSystemInfoViewAction()
        
    t.Execute(w, view)
}

func UpdateSystemHandler(w http.ResponseWriter, r *http.Request) {	
	log.Print("UpdateSystemHandler");
	
	result := UpdateSystemInfoResult{}
	
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
    	}
		
		param := UpdateSystemInfoParam{}
		name := r.FormValue("system-name")
		if len(name) > 0 {
			param.name = name
		}
		
		logo := r.FormValue("system-logo")
		if len(logo) > 0 {
			param.logo = logo
		}
		
		domain := r.FormValue("system-domain")
		if len(domain) > 0 {
			param.domain = domain
		}
		
		emailServer := r.FormValue("system-emailserver")
		if len(emailServer) > 0 {
			param.emailServer = emailServer
		}
		
		emailAccount := r.FormValue("system-emailaccount")
		if len(emailAccount) > 0 {
			param.emailAccount = emailAccount
		}
		
		emailPassword := r.FormValue("system-emailpassword")
		if len(emailPassword) > 0 && emailPassword != passwordMark {
			param.emailPassword = emailPassword
		}
		
		param.accesscode = r.FormValue("accesscode")
	
    	result = UpdateSystemInfoAction(param)
    	break
	}

    b, err := json.Marshal(result)
    if err != nil {
    	panic("marshal failed, err:" + err.Error())
    }
    
    w.Write(b)
}


func ManageModuleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageModuleHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/system/module.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
    
    view := GetModuleInfoViewAction()    
    t.Execute(w, view)
}

func ApplyModuleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ApplyModuleHandler");
	
	result := ApplyModuleResult{}
	
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
    	}
		
		param := ApplyModuleParam{}
		enableList := r.FormValue("enableList")
		parts := strings.Split(enableList,",")
		for _, v := range parts {
			if len(v) > 0 {
				param.enableList = append(param.enableList, v)
			}
		}

    	result = ApplyModuleAction(param)
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

func QueryModuleDetailHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryModuleDetailHandler");
	
	var id = ""
	idInfo := r.URL.RawQuery
	if len(idInfo) > 0 {
		parts := strings.Split(idInfo,"=")
		if len(parts) == 2 {
			id = parts[1]
		}
	}
	
	param := QueryModuleDetailParam{}
	param.id = id
	
	result := QueryModuleDetailAction(param)
    b, err := json.Marshal(result)
    if err != nil {
    	panic("marshal failed, err:" + err.Error())
    }
    
    w.Write(b)
}


func DeleteModuleBlockHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteModuleBlockHandler");
	
	result := DeleteModuleBlockResult{}
	
	for true {
		rawParams := util.SplitParam(r.URL.RawQuery)
		
		id, found := rawParams["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"    	
			break
		}
		owner, found := rawParams["owner"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"    	
			break
		}
		 
		param := DeleteModuleBlockParam{}
		idValue, err := strconv.Atoi(id)
	    if err == nil {
	    	param.id = idValue
	    	param.owner = owner
			result = DeleteModuleBlockAction(param)
	    } else {
			result.ErrCode = 1
			result.Reason = "无效请求数据"    	
	    }
	    
	    break
	}
    
    b, err := json.Marshal(result)
    if err != nil {
    	panic("marshal failed, err:" + err.Error())
    }
    
    w.Write(b)
}

func SaveModuleBlockHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("SaveModuleBlockHandler");

	result := SaveModuleBlockResult{}
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
    	}
		
		param := SaveModuleBlockParam{}
		param.owner = r.FormValue("module-id")
		param.block = r.FormValue("module-block")
	
    	result = SaveModuleBlockAction(param)
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

func SavePageBlockHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("SavePageBlockHandler");

	result := SavePageBlockResult{}
	for true {
	    err := r.ParseMultipartForm(0)
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    			
		param := SavePageBlockParam{}
		param.owner = r.FormValue("page-owner")
		param.url = r.FormValue("page-url")
		blocks := r.MultipartForm.Value["page-block"]
	    for _, b := range blocks {
			id, err := strconv.Atoi(b)
		    if err != nil {
		    	log.Print("parse page block failed, b:%s", b)
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
		    }
		    
		    param.blocks = append(param.blocks, id)
	    }		
	
    	result = SavePageBlockAction(param)
    	break
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	panic("Marshal failed, err:"  + err.Error())
    }
    
    w.Write(b)
}

