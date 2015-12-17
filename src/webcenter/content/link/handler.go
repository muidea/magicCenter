package link

import (
	"net/http"
	"encoding/json"
	"html/template"
	"strings"
	"log"
	"strconv"
	"webcenter/session"
	"webcenter/common"
)

type ManageView struct {
	Accesscode string
	LinkInfo []LinkInfo
}

type EditView struct {
	common.Result
	Accesscode string
	Id int
	Name string
	Url string
	Logo string
	Style int
	Catalog []int
}

func ManageLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageLinkHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	session := session.GetSession(w,r)
    t, err := template.ParseFiles("template/html/admin/content/link.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
	controller := &linkController{}
	info := controller.queryManageInfoAction()
    
    view := ManageView{}
    view.Accesscode = session.AccessToken()
    view.LinkInfo = info.LinkInfo
    
    t.Execute(w, view)    
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
	
	session := session.GetSession(w,r)
	
	for true {
		param := SubmitLinkParam{}
	    err := r.ParseMultipartForm(0)
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
	    catalog := r.MultipartForm.Value["link-catalog"]
	    
	    log.Print(catalog)
	    
		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Print("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    for _, ca := range catalog {
			cid, err := strconv.Atoi(ca)
		    if err != nil {
		    	log.Print("parse catalog failed, catalog:%s", ca)
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
		    }
		    
		    param.catalog = append(param.catalog, cid)
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
		param.creater, _ = session.GetAccountId()

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
	log.Print("EditLinkHandler");
	
	result := EditView{}
	
	for true {
		param := EditLinkParam{}
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
    	lnk := controller.editLinkAction(param)
    	
    	result.ErrCode = lnk.ErrCode
    	result.Reason = lnk.Reason
    	result.Id = lnk.Link.Id()
    	result.Name = lnk.Link.Name()
    	result.Url = lnk.Link.Url()
    	result.Logo = lnk.Link.Logo()
    	result.Style = lnk.Link.Style()
    	for _, c := range lnk.Link.Relative() {
    		result.Catalog = append(result.Catalog, c.Id())
    	}    	
    	
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

