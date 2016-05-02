package ui

import (
	"log"
	"net/http"
	"encoding/json"
	"html/template"
	"magiccenter/kernel/common"
	"magiccenter/configuration"
)

const passwordMark = "********"

type ManageSystemView struct {
	SystemInfo configuration.SystemInfo
}

type UpdateSystemResult struct {
	common.Result
}

func ManageSystemHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageSystemHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/system/system.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
    view := ManageSystemView{}
    view.SystemInfo = configuration.GetSystemInfo()
    view.SystemInfo.MailPassword = passwordMark
        
    t.Execute(w, view)
}

func UpdateSystemHandler(w http.ResponseWriter, r *http.Request) {	
	log.Print("UpdateSystemHandler");
	
	result := UpdateSystemResult{}
	
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
    	}
		
		info := configuration.GetSystemInfo()
		name := r.FormValue("system-name")
		if len(name) > 0 {
			info.Name = name
		}
		
		logo := r.FormValue("system-logo")
		if len(logo) > 0 {
			info.Logo = logo
		}
		
		domain := r.FormValue("system-domain")
		if len(domain) > 0 {
			info.Domain = domain
		}
		
		emailServer := r.FormValue("email-server")
		if len(emailServer) > 0 {
			info.MailServer = emailServer
		}
		
		emailAccount := r.FormValue("email-account")
		if len(emailAccount) > 0 {
			info.MailAccount = emailAccount
		}
				
		emailPassword := r.FormValue("email-password")
		if len(emailPassword) > 0 && emailPassword != passwordMark {
			info.MailPassword = emailPassword
		}
		
		if configuration.UpdateSystemInfo(info) {
			result.ErrCode = 0
			result.Reason = "更新系统信息成功"
		} else {
			result.ErrCode = 1
			result.Reason = "更新系统信息失败"			
		}
		
    	break
	}

    b, err := json.Marshal(result)
    if err != nil {
    	panic("marshal failed, err:" + err.Error())
    }
    
    w.Write(b)
}

