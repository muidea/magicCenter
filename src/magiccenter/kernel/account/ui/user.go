package ui

import (
	"log"
	"strconv"
	"net/http"
	"encoding/json"
	"html/template"
	"muidea.com/util"
	"magiccenter/kernel/common"
    "magiccenter/kernel/account/model"
    "magiccenter/kernel/account/bll"
)


type ManageUserView struct {
	Users []model.UserDetail
	Groups []model.GroupInfo
}

type QueryAllUserResult struct {
	common.Result
	Users []model.UserDetail
}

type QueryUserResult struct {
	common.Result
	User model.UserDetail
}

type CreateUserResult struct {
	common.Result
	Users []model.UserDetail	
}

type DeleteUserResult struct {
	CreateUserResult
}

type UpdateUserResult struct {
	common.Result
	User model.UserDetail	
}

func ManageUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageUserHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/account/user.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
    view := ManageUserView{}
    view.Users = bll.QueryAllUser()
    view.Groups = bll.QueryAllGroupInfo()
    
    t.Execute(w, view)
}

func QueryAllUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllUserHandler");
	
	
	result := QueryAllUserResult{}
	result.Users = bll.QueryAllUser()
	result.ErrCode = 0
	result.Reason = "查询成功"
		
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, failed, err:" + err.Error())
    }
    
    w.Write(b)
}

func QueryUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryUserHandler");
	
	result := QueryUserResult{}
	
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
		}
		
		uid, err := strconv.Atoi(id)
		if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;			
		}
		
		result.User, found = bll.QueryUserById(uid)
		if !found {
    		result.ErrCode = 1
    		result.Reason = "指定User不存在"
    		break;
		}
		
		result.ErrCode = 0
		result.Reason = "查询成功"		    	
    	break
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, failed, err:" + err.Error())
    }
    
    w.Write(b)
}

func AjaxUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxUserHandler");
	
	result := CreateUserResult{}
	for true {
	    err := r.ParseMultipartForm(0)
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		account := r.FormValue("user-account")
		email := r.FormValue("user-email")
		groups := r.MultipartForm.Value["user-group"]    	
	    groupList := []int{}
	    for _, g := range groups {
			gid, err := strconv.Atoi(g)
		    if err != nil {
		    	log.Print("parse group id failed, group:%s", g)
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
		    }
		    
		    groupList = append(groupList, gid)	    	
	    }
	    
	    ok := bll.CreateUser(account, email, groupList)
	    if !ok {
			result.ErrCode = 1
			result.Reason = "保存分组失败"
			break	    	
	    }
	    
	    result.Users = bll.QueryAllUser()
		result.ErrCode = 1
		result.Reason = "保存分组成功"
	    break
	}
    
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, failed, err:" + err.Error())
        return
    }
    
    w.Write(b)
}


func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteUserHandler");
	
	result := DeleteUserResult{}
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
		}
		
		uid, err := strconv.Atoi(id)
		if err != nil {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;			
		}
		
		ok := bll.DeleteUser(uid)
		if !ok {
    		result.ErrCode = 1
    		result.Reason = "删除分组失败"
    		break;
		}
		
		result.ErrCode = 0
		result.Reason = "查询成功"
		result.Users = bll.QueryAllUser()
    	break
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, failed, err:" + err.Error())
    }
    
    w.Write(b)
}


func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("UpdateUserHandler");
	
	result := UpdateUserResult{}
	for true {
	    err := r.ParseMultipartForm(0)
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		id, err := strconv.Atoi(r.FormValue("user-id"))
		if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break			
		}
		name := r.FormValue("user-name")
		email := r.FormValue("user-email")
		status, err := strconv.Atoi(r.FormValue("user-status"))
		if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break			
		}		
		groups := r.MultipartForm.Value["user-group"]
	    groupList := []int{}
	    for _, g := range groups {
			gid, err := strconv.Atoi(g)
		    if err != nil {
		    	log.Print("parse group id failed, group:%s", g)
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
		    }
		    
		    groupList = append(groupList, gid)	    	
	    }
	    
	    ok := bll.UpdateUserDetail(id, name, email, status,groupList)
	    if !ok {
			result.ErrCode = 1
			result.Reason = "更新用户信息失败"
			break
	    }
	    
	    result.User, _ = bll.QueryUserById(id)
		result.ErrCode = 0
		result.Reason = "更新用户信息成功"
	    break
	}
    
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, failed, err:" + err.Error())
        return
    }
    
    w.Write(b)
}




