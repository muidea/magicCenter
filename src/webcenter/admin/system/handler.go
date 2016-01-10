package system

import (
	"net/http"
	"html/template"
	"encoding/json"
	"log"
    "webcenter/application"
    "webcenter/module"
)

type SystemView struct {
	Name string
	Logo string
	Domain string
	EMailServer string
	EMailAccount string
	EMailPassword string
}

type ModuleItem struct {
	Id int
	Name string
	Description string
	Uri string
	Enable bool
	Default bool
	Internal bool
}

type ModuleView struct {
	AccessCode string
	ModuleList []ModuleItem
}

const passwordMark = "******"

func ManageSystemHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageSystemHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/system/system.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
    view := SystemView{}
    view.Name = application.Name()
    view.Logo = application.Logo()
    view.Domain = application.Domain()
    view.EMailServer = application.MailServer()
    view.EMailAccount = application.MailAccount()
    view.EMailPassword = passwordMark
    
    t.Execute(w, view)
}

func UpdateSystemHandler(w http.ResponseWriter, r *http.Request) {	
	log.Print("UpdateSystemHandler");
	
	result := UpdateResult{}
	
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
    	}
		
		param := UpdateParam{}
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
	
		log.Printf("Name:%s,Logo:%s,Domain:%s,server:%s,account:%s,password:%s", param.name, param.logo, param.domain, param.emailServer, param.emailAccount, param.emailPassword)
	
    	controller := &systemController{}
    	result = controller.UpdateAction(&param)
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


func ManageModuleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageModuleHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/system/module.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
    
    view := ModuleView{}
    modulesList := module.QueryAllModules()
    for index, _ := range modulesList {
    	m := modulesList[index]
    	
    	log.Println(m)
    	
    	item := ModuleItem{}
    	item.Id = m.ID()
    	item.Name = m.Name()
    	item.Description = m.Description()
    	item.Uri = m.Uri()
    	item.Enable = m.EnableState()
    	item.Default = m.DefaultState()
    	item.Internal = m.Internal()
    	
    	view.ModuleList = append(view.ModuleList,item)
    }
    
    t.Execute(w, view)
}

func ApplyModuleHandler(w http.ResponseWriter, r *http.Request) {	
	log.Print("ApplyModuleHandler");
	
	result := UpdateResult{}
	
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
    	}
		
		param := UpdateParam{}
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
	
		log.Printf("Name:%s,Logo:%s,Domain:%s,server:%s,account:%s,password:%s", param.name, param.logo, param.domain, param.emailServer, param.emailAccount, param.emailPassword)
	
    	controller := &systemController{}
    	result = controller.UpdateAction(&param)
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

