package content

import (
	"net/http"
	"encoding/json"
	"strings"
	"log"
	"time"
	"strconv"
	"webcenter/session"
	"webcenter/auth"
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
}

func queryAllContentHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllContentHandler");
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(auth.AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := auth.NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}
	 
    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }

    accessCode := r.FormValue("accesscode")
    
	param := GetAllContentParam{}
	param.accessCode = accessCode
    param.session = session

    controller := &contentController{}
    result := controller.getAllContentAction(param)
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(auth.AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := auth.NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}

    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
	
	id := r.FormValue("article-id")
	title := r.FormValue("article-title")
	content := r.FormValue("article-content")
	catalog := r.FormValue("article-catalog")
	
	param := SubmitArticleParam{}
	param.id, err = strconv.Atoi(id)
    if err != nil {
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }	
	param.catalog, err = strconv.Atoi(catalog)
    if err != nil {
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    param.title = title
    param.content = content
    param.author = user.Id
    param.submitDate = time.Now().Format("2006-01-02 15:04:05")

    controller := &contentController{}
    result := controller.submitArticleAction(param)
    
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(auth.AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := auth.NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}
	 
    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }

    accessCode := r.FormValue("accesscode")
    
	param := GetAllArticleParam{}
	param.accessCode = accessCode
    param.session = session

    controller := &contentController{}
    result := controller.getAllArticleAction(param)
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(auth.AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := auth.NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}

    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }

	var id = ""
	idInfo := r.URL.RawQuery
	if len(idInfo) > 0 {
		parts := strings.Split(idInfo,"=")
		if len(parts) == 2 {
			id = parts[1]
		}
	}
	
	param := GetArticleParam{}
	accessCode := r.FormValue("accesscode")
	param.id, err = strconv.Atoi(id)
    if err != nil {
    	log.Printf("convert id failed, id:%s,accessCode:%s", id, accessCode)
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
	param.accessCode = accessCode
	param.session = session
	
	log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
	 
    controller := &contentController{}
    result := controller.getArticleAction(param)
    
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(auth.AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := auth.NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}

    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }

	var id = ""
	idInfo := r.URL.RawQuery
	if len(idInfo) > 0 {
		parts := strings.Split(idInfo,"=")
		if len(parts) == 2 {
			id = parts[1]
		}
	}
	
	param := DeleteArticleParam{}
	accessCode := r.FormValue("accesscode")
	param.id, err = strconv.Atoi(id)
    if err != nil {
    	log.Printf("convert id failed, id:%s,accessCode:%s", id, accessCode)
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
	param.accessCode = accessCode
	param.session = session
	
	log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
	
    controller := &contentController{}
    result := controller.deleteArticleAction(param)
    
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(auth.AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := auth.NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}
	 
    err = r.ParseForm()
    if err != nil {
    	log.Print("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    accessCode := r.FormValue("accesscode")
    log.Printf("accessCode:%s",accessCode);
	 
	param := GetAllCatalogParam{}
	param.accessCode = accessCode
	param.session = session
    controller := &contentController{}
    result := controller.getAllCatalogAction(param)
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
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(auth.AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := auth.NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}
	 
	var id = ""
	idInfo := r.URL.RawQuery
	if len(idInfo) > 0 {
		parts := strings.Split(idInfo,"=")
		if len(parts) == 2 {
			id = parts[1]
		}
	}
	
	param := GetCatalogParam{}
	accessCode := r.FormValue("accesscode")
	param.id, err = strconv.Atoi(id)
    if err != nil {
    	log.Printf("convert id failed, id:%s,accessCode:%s", id, accessCode)
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
	param.accessCode = accessCode
	param.session = session	 
	 
    controller := &contentController{}
    result := controller.getCatalogAction(param)
    b, err := json.Marshal(result)
    if err != nil {
    	log.Fatal("json marshal failed, err:" + err.Error())
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    w.Write(b)    
}

