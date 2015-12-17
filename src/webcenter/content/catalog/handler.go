package catalog

import (
	"net/http"
	"encoding/json"
	"html/template"
	"strings"
	"log"
	"strconv"
	"webcenter/session"
)

type ManageView struct {
	Accesscode string
	Catalog []CatalogInfo
}

func ManageCatalogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	log.Print("ManageCatalogHandler");
	
	session := session.GetSession(w,r)
    t, err := template.ParseFiles("template/html/admin/content/catalog.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
	controller := &catalogController{}
	info := controller.queryManageInfoAction()
    
    view := ManageView{}
    view.Accesscode = session.AccessToken()
    view.Catalog = info.Catalog
    
    t.Execute(w, view)    
}


func QueryAllCatalogInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllCatalogInfoHandler");
	
	result := QueryAllCatalogInfoResult{}
	
	for true {
		param := QueryAllCatalogInfoParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
	    accessCode := r.FormValue("accesscode")
		param.accessCode = accessCode

    	controller := &catalogController{}
    	result = controller.queryAllCatalogInfoAction(param)
    	
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


func QueryCatalogInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryCatalogInfoHandler");
	
	result := QueryCatalogInfoResult{}
	
	for true {
		param := QueryCatalogInfoParam{}
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
		    	
    	controller := &catalogController{}
    	result = controller.queryCatalogInfoAction(param)    	
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

func QueryAvalibleParentCatalogInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryAvalibleParentCatalogInfoHandler");
	
	result := QueryAvalibleParentCatalogInfoResult{}
	
	for true {
		param := QueryAvalibleParentCatalogInfoParam{}
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
		    	
    	controller := &catalogController{}
    	result = controller.queryAvalibleParentCatalogInfoAction(param)
    	
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

func QuerySubCatalogInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QuerySubCatalogInfoHandler");
	
	result := QuerySubCatalogInfoResult{}
	
	for true {
		param := QuerySubCatalogInfoParam{}
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
		    	
    	controller := &catalogController{}
    	result = controller.querySubCatalogInfoAction(param)
    	
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

func QueryCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryCatalogHandler");
	
	result := QueryCatalogResult{}
	
	for true {
		param := QueryCatalogParam{}
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
		    	
    	controller := &catalogController{}
    	result = controller.queryCatalogAction(param)
    	
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

func DeleteCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteCatalogHandler");
	
	result := DeleteCatalogResult{}
	
	for true {
		param := DeleteCatalogParam{}
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
		
	    controller := &catalogController{}
	    result = controller.deleteCatalogAction(param)
    	
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

func AjaxCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxCatalogHandler");
	
	result := SubmitCatalogResult{}
	
	session := session.GetSession(w,r)
	for true {
		param := SubmitCatalogParam{}
	    err := r.ParseMultipartForm(0)
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		cId := r.FormValue("catalog-id")
		cName := r.FormValue("catalog-name")
		cParent := r.MultipartForm.Value["catalog-parent"]
	    
		param.id, err = strconv.Atoi(cId)
	    if err != nil {
	    	log.Print("parse id failed, id:%d", cId)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    param.name = cName
	    
	    for _, p := range cParent {
			pid, err := strconv.Atoi(p)
		    if err != nil {
		    	log.Print("parse id failed, pid:%d", cParent)
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
		    }
		    param.pid = append(param.pid, pid)
	    }
	    	    
	    param.creater, _ = session.GetAccountId()

    	controller := &catalogController{}
    	result = controller.submitCatalogAction(param)
    	
    	break
	}
		
    b, err := json.Marshal(result)
    if err != nil {
    	panic("marshal failed")
    }
    
    w.Write(b)	
}

func EditCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("EditCatalogHandler");
	
	result := EditCatalogResult{}
	
	for true {
		param := EditCatalogParam{}
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
		    	
    	controller := &catalogController{}
    	result = controller.editCatalogAction(param)
    	
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
