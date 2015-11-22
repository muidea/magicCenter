package content

import (
	"net/http"
	"encoding/json"
	"strings"
	"log"
	"time"
	"os"
	"io"
	"fmt"
	"strconv"
	"muidea.com/util"
	"webcenter/application"
	"webcenter/session"
)

func init() {
	registerRouter()
}

func registerRouter() {
	application.RegisterPostHandler("/content/admin/queryAllContent/", queryAllContentHandler)
	application.RegisterPostHandler("/content/admin/queryAllArticle/", queryAllArticleHandler)
	application.RegisterPostHandler("/content/admin/queryArticle/", queryArticleHandler)
	application.RegisterPostHandler("/content/admin/editArticle/", editArticleHandler)
	application.RegisterPostHandler("/content/admin/deleteArticle/", deleteArticleHandler)
	application.RegisterPostHandler("/content/article/ajax/", ajaxArticleHandler)
	application.RegisterPostHandler("/content/admin/queryAllCatalog/", queryAllCatalogHandler)
	application.RegisterPostHandler("/content/admin/queryCatalog/", queryCatalogHandler)
	application.RegisterPostHandler("/content/admin/editCatalog/", editCatalogHandler)
	application.RegisterPostHandler("/content/admin/deleteCatalog/", deleteCatalogHandler)
	application.RegisterPostHandler("/content/catalog/ajax/", ajaxCatalogHandler)
	application.RegisterPostHandler("/content/admin/queryAllLink/", queryAllLinkHandler)
	application.RegisterPostHandler("/content/admin/queryLink/", queryLinkHandler)
	application.RegisterPostHandler("/content/admin/editLink/", editLinkHandler)
	application.RegisterPostHandler("/content/admin/deleteLink/", deleteLinkHandler)
	application.RegisterPostHandler("/content/link/ajax/", ajaxLinkHandler)
	application.RegisterPostHandler("/content/admin/queryAllImage/", queryAllImageHandler)
	application.RegisterPostHandler("/content/admin/deleteImage/", deleteImageHandler)
	application.RegisterPostHandler("/content/image/ajax/", ajaxImageHandler)
}

func queryAllContentHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllContentHandler");
	
	result := QueryAllContentResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := QueryAllContentParam{}
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
    	result = controller.queryAllContentAction(param)
    	
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
	
	result := QueryAllArticleResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := QueryAllArticleParam{}
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
    	result = controller.queryAllArticleAction(param)
    	
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
	
	result := QueryArticleResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := QueryArticleParam{}
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
    	result = controller.queryArticleAction(param)
    	
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


func editArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryArticleHandler");
	
	result := EditArticleResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := EditArticleParam{}
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
    	result = controller.editArticleAction(param)
    	
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

func queryAllCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllCatalogHandler");
	
	result := QueryAllCatalogResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := QueryAllCatalogParam{}
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
    	result = controller.queryAllCatalogAction(param)
    	
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
	
	result := QueryCatalogResult{}
	
	for true {
		session := session.GetSession(w,r)
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
		param.session = session
		    	
    	controller := &contentController{}
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

func editCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("editCatalogHandler");
	
	result := EditCatalogResult{}
	
	for true {
		session := session.GetSession(w,r)
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
		param.session = session
		    	
    	controller := &contentController{}
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
	    pid := r.FormValue("catalog-parent")
	    accessCode := r.FormValue("accesscode")
	    
		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Print("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    param.name = name
		param.pid, err = strconv.Atoi(pid)
	    if err != nil {
	    	log.Print("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }

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


func queryAllLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllLinkHandler");
	
	result := QueryAllLinkResult{}
	
	for true {
		session := session.GetSession(w,r)
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
    	param.session = session

    	controller := &contentController{}
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

func queryLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryLinkHandler");
	
	result := QueryLinkResult{}
	
	for true {
		session := session.GetSession(w,r)
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
		param.session = session
		    	
    	controller := &contentController{}
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

func editLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("editLinkHandler");
	
	result := EditLinkResult{}
	
	for true {
		session := session.GetSession(w,r)
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
		param.session = session
		    	
    	controller := &contentController{}
    	result = controller.editLinkAction(param)
    	
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

func deleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteLinkHandler");
	
	result := DeleteLinkResult{}
	
	for true {
		session := session.GetSession(w,r)
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
		param.session = session
		
		log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
		 
	    controller := &contentController{}
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

func ajaxLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxLinkHandler");
	
	result := SubmitLinkResult{}
	
	for true {
		session := session.GetSession(w,r)
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
    	param.session = session

    	controller := &contentController{}
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


func queryAllImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllImageHandler");
	
	result := QueryAllImageResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := QueryAllImageParam{}
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
    	result = controller.queryAllImageAction(param)
    	
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

func deleteImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteImageHandler");
	
	result := DeleteImageResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := DeleteImageParam{}
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
	    result = controller.deleteImageAction(param)
    	
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

func ajaxImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxImageHandler");
	
	result := SubmitImageResult{}
	
	for true {
		session := session.GetSession(w,r)
		param := SubmitImageParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		file, head, err := r.FormFile("image-name")
		if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}
		
		fileName := fmt.Sprintf("%s/%s/%s_%s_%s", application.StaticPath(), application.UploadPath(), time.Now().Format("20060102150405"), util.RandomAlphabetic(16), head.Filename);
		
		log.Print(fileName)
		defer file.Close()
		f,err:=os.Create(fileName)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break			
		}
		
		defer f.Close()
		io.Copy(f,file)		
		
		desc := r.FormValue("image-desc")
	    accessCode := r.FormValue("accesscode")

		staticPath := application.StaticPath()
		param.url = fileName[len(staticPath):]
		param.desc = desc
		param.accessCode = accessCode
    	param.session = session

    	controller := &contentController{}
    	result = controller.submitImageAction(param)
    	
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


