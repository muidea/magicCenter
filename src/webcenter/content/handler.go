package content

import (
	"net/http"
	"encoding/json"
	"strings"
	"log"
	"time"
	"strconv"
	"webcenter/session"
)

func init() {
	registerRouter()
}

func registerRouter() {
	http.HandleFunc("/content/admin/queryAllContent/", queryAllContentHandler)
	http.HandleFunc("/content/admin/queryAllArticle/", queryAllArticleHandler)
	http.HandleFunc("/content/admin/queryArticle/", queryArticleHandler)
	http.HandleFunc("/content/admin/deleteArticle/", deleteArticleHandler)
	http.HandleFunc("/content/article/ajax/", ajaxArticleHandler)
	http.HandleFunc("/content/admin/queryAllCatalog/", queryAllCatalogHandler)
	http.HandleFunc("/content/admin/queryCatalog/", queryCatalogHandler)
	http.HandleFunc("/content/admin/deleteCatalog/", deleteCatalogHandler)
	http.HandleFunc("/content/catalog/ajax/", ajaxCatalogHandler)
}

func queryAllContentHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllContentHandler");
	
	result := GetAllContentResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := GetAllContentParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
	    accessCode := r.FormValue("accesscode")
		param.accessCode = accessCode
    	param.session = session

    	controller := &contentController{}
    	result = controller.getAllContentAction(param)
    	
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

func queryAllArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllArticleHandler");
	
	result := GetAllArticleResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := GetAllArticleParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
	    accessCode := r.FormValue("accesscode")
		param.accessCode = accessCode
    	param.session = session

    	controller := &contentController{}
    	result = controller.getAllArticleAction(param)
    	
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

func ajaxArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxArticleHandler");
	
	result := SubmitArticleResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := SubmitArticleParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		id := r.FormValue("article-id")
		title := r.FormValue("article-title")
		content := r.FormValue("article-content")
		catalog := r.FormValue("article-catalog")    
	    accessCode := r.FormValue("accesscode")
	    
		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Print("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }	
		param.catalog, err = strconv.Atoi(catalog)
	    if err != nil {
	    	log.Print("parse catalog failed, catalog:%s", catalog)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    
	    param.title = title
	    param.content = content
	    param.submitDate = time.Now().Format("2006-01-02 15:04:05")	    
		param.accessCode = accessCode
    	param.session = session

    	controller := &contentController{}
    	result = controller.submitArticleAction(param)
    	
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

func queryArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryArticleHandler");
	
	result := GetArticleResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := GetArticleParam{}
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
		param.session = session
		    	
    	controller := &contentController{}
    	result = controller.getArticleAction(param)
    	
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


func deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteArticleHandler");
	
	result := DeleteArticleResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := DeleteArticleParam{}
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
		param.session = session
		
		log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
		 
	    controller := &contentController{}
	    result = controller.deleteArticleAction(param)
    	
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

func queryAllCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllCatalogHandler");
	
	result := GetAllCatalogResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := GetAllCatalogParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
	    accessCode := r.FormValue("accesscode")
		param.accessCode = accessCode
    	param.session = session

    	controller := &contentController{}
    	result = controller.getAllCatalogAction(param)
    	
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


func ajaxCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxCatalogHandler");
	
	result := SubmitCatalogResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := SubmitCatalogParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		id := r.FormValue("catalog-id")
		name := r.FormValue("catalog-name")    	
	    accessCode := r.FormValue("accesscode")
	    
		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Print("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    param.name = name
	    param.submitDate = time.Now().Format("2006-01-02 15:04:05")	    
		param.accessCode = accessCode
    	param.session = session

    	controller := &contentController{}
    	result = controller.submitCatalogAction(param)
    	
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

func queryCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryCatalogHandler");
	
	result := GetCatalogResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := GetCatalogParam{}
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
		param.session = session
		    	
    	controller := &contentController{}
    	result = controller.getCatalogAction(param)
    	
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


func deleteCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteCatalogHandler");
	
	result := DeleteCatalogResult{}
	
	for true {
		session := session.GetSession(w,r)
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
		param.session = session
		
		log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
		 
	    controller := &contentController{}
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

