package article

import (
	"net/http"
	"encoding/json"
	"html/template"
	"strings"
	"log"
	"time"
	"strconv"
	"webcenter/session"
)

type ManageView struct {
	Accesscode string
	ArticleInfo []ArticleSummary
}

func ManageArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageArticleHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	session := session.GetSession(w,r)
    t, err := template.ParseFiles("template/html/admin/content/article.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
	controller := &articleController{}
	info := controller.queryManageInfoAction()
    
    view := ManageView{}
    view.Accesscode = session.AccessToken()
    view.ArticleInfo = info.ArticleInfo
    
    t.Execute(w, view)
}


func QueryAllArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryAllArticleHandler");
	
	result := QueryAllArticleResult{}
	
	for true {
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

    	controller := &articleController{}
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

func QueryArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryArticleHandler");
	
	result := QueryArticleResult{}
	
	for true {
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
		    	
    	controller := &articleController{}
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

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteArticleHandler");
	
	result := DeleteArticleResult{}
	
	for true {
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
		
	    controller := &articleController{}
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

func AjaxArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxArticleHandler");
	
	result := SubmitArticleResult{}
	
	session := session.GetSession(w,r)
	
	for true {
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
	    
		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Print("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }	
		cid, err := strconv.Atoi(catalog)
	    if err != nil {
	    	log.Print("parse catalog failed, catalog:%s", catalog)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    param.catalog = append(param.catalog, cid)
	    
	    param.title = title
	    param.content = content
	    param.submitDate = time.Now().Format("2006-01-02 15:04:05")
	    param.author, _ = session.GetAccountId()

    	controller := &articleController{}
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

func EditArticleHandler(w http.ResponseWriter, r *http.Request) {
	
}

