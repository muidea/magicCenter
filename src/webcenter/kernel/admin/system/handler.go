package system

import (
	"net/http"
	"html/template"
	"encoding/json"
	"log"
	"strings"
	"reflect"
    "webcenter/module"
    "webcenter/kernel"
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
	Id string
	Name string
	Description string
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
    view.Name = kernel.Name()
    view.Logo = kernel.Logo()
    view.Domain = kernel.Domain()
    view.EMailServer = kernel.MailServer()
    view.EMailAccount = kernel.MailAccount()
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
    
    
    view := ModuleView{}
    modulesList := module.QueryAllEntities()
    for index, _ := range modulesList {
    	m := modulesList[index]
    	
    	log.Println(m)
    	
    	item := ModuleItem{}
    	item.Id = m.ID()
    	item.Name = m.Name()
    	item.Description = m.Description()
    	item.Enable = m.EnableState()
    	item.Default = m.DefaultState()
    	item.Internal = m.Internal()
    	
    	view.ModuleList = append(view.ModuleList,item)
    }
    
    t.Execute(w, view)
}

func ApplyModuleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ApplyModuleHandler");
	
	result := ApplyResult{}
	
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
    	}
		
		param := ApplyParam{}
		enableList := r.FormValue("enableList")
		parts := strings.Split(enableList,",")
		for _, v := range parts {
			if len(v) > 0 {
				param.enableList = append(param.enableList, v)
			}
		}
				
		disableList := r.FormValue("disableList")
		parts = strings.Split(disableList,",")
		for _, v := range parts {
			if len(v) > 0 {
				param.disableList = append(param.disableList, v)
			}
		}
		
		defaultModule := r.FormValue("defaultModule")
		parts = strings.Split(defaultModule,",")
		for _, v := range parts {
			if len(v) > 0 {
				param.defaultModule = append(param.defaultModule, v)
			}
		}
		
		controller := &systemController{}
    	result = controller.ApplyAction(&param)
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

func MaintainModuleHandler(w http.ResponseWriter, r *http.Request) {
	var id = ""
	idInfo := r.URL.RawQuery
	if len(idInfo) > 0 {
		parts := strings.Split(idInfo,"=")
		if len(parts) == 2 {
			id = parts[1]
		}
	}
	
	m, found := module.QueryModule(id)
	if found {
		rts := m.Routes()
		for _, rt := range rts {
			if rt.Pattern() == "/maintain/" {
				fv := reflect.ValueOf(rt.Handler())
				params := make([]reflect.Value, 2)
				params[0] = reflect.ValueOf(w)
				params[1] = reflect.ValueOf(r)
				fv.Call(params)
				//handler := rt.Handler().type( func(http.ResponseWriter, *http.Request) )				
			}
		}
	}
	
}




