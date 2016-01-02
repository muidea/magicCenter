package account

import (
	"net/http"
	"encoding/json"
	"html/template"
	"log"
	"fmt"
	"strings"
	"strconv"
	"webcenter/session"
	"webcenter/common"
	"webcenter/auth/group"
)

type ManageView struct {
	Accesscode string
	UserInfo []UserInfo
	GroupInfo []group.GroupInfo
}

type EditView struct {
	common.Result
	Accesscode string
	Id int
	Account string
	Email string
	Group []int
}

func ManageUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageUserHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	session := session.GetSession(w,r)
    t, err := template.ParseFiles("template/html/admin/account/user.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
	controller := &accountController{}
	info := controller.queryManageInfoAction()
    
    view := ManageView{}
    view.Accesscode = session.AccessToken()
    view.UserInfo = info.UserInfo
    view.GroupInfo = info.GroupInfo
    
    t.Execute(w, view)
}

func CheckAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("CheckAccountHandler");
	
	result := CheckAccountResult{}
	for true {
		param := CheckAccountParam{}
	    err := r.ParseForm()
    	if err != nil {
    		panic("paseform failed, err:" + err.Error())
    	}
    	
		account := r.FormValue("user-account")
		accessCode := r.FormValue("accesscode")
				
	    param.account = account
	    param.accessCode = accessCode
	    
	    controller := &accountController{}
	    result = controller.checkAccountAction(param)
	    
	    break
	}
    
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, err:" + err.Error())
    }
    
    w.Write(b)	
}

func QueryAllUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllUserHandler");
	
	result := QueryAllUserResult{}
	
	for true {
		param := QueryAllUserParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	    	
	    accessCode := r.FormValue("accesscode")
		param.accessCode = accessCode

    	controller := &accountController{}
    	result = controller.queryAllUserAction(param)
    	
    	break
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, err:" + err.Error())
    }
    
    w.Write(b)
}

func AjaxUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxUserHandler");
	
	result := SubmitUserResult{}
	for true {
		param := SubmitUserParam{}
	    err := r.ParseMultipartForm(0)
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		id := r.FormValue("user-id")
		name := r.FormValue("user-account")
		email := r.FormValue("user-email")
		groups := r.MultipartForm.Value["user-group"]
		accessCode := r.FormValue("accesscode")
		
		param.id, err = strconv.Atoi(id)
	    if err != nil {
	    	log.Print("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    
	    param.group = ""
	    for _, g := range groups {
			gid, err := strconv.Atoi(g)
		    if err != nil {
		    	log.Print("parse group failed, group:%s", g)
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
		    }
		    
		    if len(param.group) == 0 {
		    	param.group = fmt.Sprintf("%d", gid)
		    } else {
		    	param.group = fmt.Sprintf("%s,%d", param.group, gid)
		    }
	    }
	    param.account = name
	    param.email = email    
	    param.accessCode = accessCode
	    
	    controller := &accountController{}
	    result = controller.submitUserAction(param)
	    
	    break
	}
    
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, err:" + err.Error())
    }
    
    w.Write(b)	
}

func QueryUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryUserHandler");
		
	result := QueryUserResult{}
	
	for true {
		param := QueryUserParam{}
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
		
	    controller := &accountController{}
	    result = controller.queryUserAction(param)
    	
    	break
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, err:" + err.Error())
    }
    
    w.Write(b)
}


func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteArticleHandler");
	
	result := DeleteUserResult{}
	
	for true {
		param := DeleteUserParam{}
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
		
	    controller := &accountController{}
	    result = controller.deleteUserAction(param)
    	
    	break
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, err:" + err.Error())
    }
    
    w.Write(b)
}


func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("EditUserHandler");
		
	result := EditView{}
	
	for true {
		param := QueryUserParam{}
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
		
	    controller := &accountController{}
	    re := controller.queryUserAction(param)
    	
    	result.ErrCode = re.ErrCode
    	result.Reason = re.Reason
    	result.Id = re.User.Id
    	result.Account = re.User.Account
    	result.Email = re.User.Email
    	
    	parts := strings.Split(re.User.Group,",")
    	for _, g := range parts {
    		gid, _ := strconv.Atoi(g)
    		result.Group = append(result.Group, gid)
    	}
    	
    	break
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, err:" + err.Error())
    }
    
    w.Write(b)
}


