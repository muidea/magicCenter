package content

import (
	"net/http"
	"encoding/json"
	"log"
	"webcenter/session"
	"webcenter/auth"
)

func init() {
	registerRouter()
}

func registerRouter() {
	http.HandleFunc("/content/admin/getAllArticleInfo/", getAllArticleInfoHandler)
	http.HandleFunc("/content/admin/getAllCatalog/", getAllCatalogHandler)
}

func getAllArticleInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAllArticleInfoHandler");
	
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
	 
    controller := &contentController{}
    result := controller.getAllArticleInfoAction(session)
    b, err := json.Marshal(result)
    if err != nil {
    	log.Fatal("json marshal failed, err:" + err.Error())
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    w.Write(b)    
}


func getAllCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAllCatalogHandler");
	
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
	 
    controller := &contentController{}
    result := controller.getAllCatalogAction(session)
    b, err := json.Marshal(result)
    if err != nil {
    	log.Fatal("json marshal failed, err:" + err.Error())
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    w.Write(b)    
}

