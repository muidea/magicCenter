package patrol

import (
	"strconv"
	"encoding/json"
	"net/http"
	"html/template"
	"log"
	"webcenter/common"
	"webcenter/session"
)


type AdminPatrolPage struct {
	common.AdminPage
	Routeline []Routeline
}

type patrolController struct {
}

type GetAllPatrolLineResult struct {
	common.Result
	Routelines []Routeline
}

type QueryPatrolLineResult struct {
	common.Result
	Routeline *Routeline
}

type UpdatePatrolLineResult struct {
	common.Result
}

func (this *patrolController)OutputGetAllPatrolLineResult(w http.ResponseWriter, errCode int, reason string, routelines []Routeline ) {
    out := &GetAllPatrolLineResult{}
    out.ErrCode = errCode
    out.Reason = reason
    out.Routelines = routelines
    
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}

func (this *patrolController)OutputQueryPatrolLineResult(w http.ResponseWriter, errCode int, reason string, routeline *Routeline ) {
    out := &QueryPatrolLineResult{}
    out.ErrCode = errCode
    out.Reason = reason
    out.Routeline = routeline
    
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}

func (this *patrolController)OutputUpdatePatrolLineResult(w http.ResponseWriter, errCode int, reason string) {
    out := &UpdatePatrolLineResult{}
    out.ErrCode = errCode
    out.Reason = reason
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}

func (this *patrolController)AdminPatrolAction(w http.ResponseWriter, r *http.Request, session *session.Session) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")	
	access_token := session.AccessToken()
	account, _ := session.GetOption("account")
	
    t, err := template.ParseFiles("template/html/admin_patrol.html")
    if (err != nil) {
        log.Println(err)
        http.Redirect(w, r, "/", http.StatusNotFound)
        return        
    }
 
 	routelines := GetAllRouteLine()
 
 	pageInfo := AdminPatrolPage{}
 	pageInfo.Account = account.(string)
 	pageInfo.AccessToken = access_token
 	pageInfo.Routeline = routelines
    
    t.Execute(w, pageInfo)
}

func (this *patrolController)GetAllPatrolLineAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("charset", "utf-8")
	
	var errCode int
	var reason string
	
	errCode = 0
	routelines := GetAllRouteLine()
	this.OutputGetAllPatrolLineResult(w, errCode, reason, routelines)
}

func (this *patrolController)QueryPatrolLineAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("charset", "utf-8")
	
	var errCode int
	var reason string
	id, err := strconv.Atoi(r.URL.RawQuery)
	if err != nil {
		errCode = -1
		reason = "非法输入"
		this.OutputQueryPatrolLineResult(w, errCode, reason, nil)
		return
	}
	
	routeline,found := FindRouteline(id)
	if found {
		errCode = 0
	} else {
		errCode = 1
		reason = "找不到该路线路线"		
	}
	
	this.OutputQueryPatrolLineResult(w, errCode, reason, &routeline)		
}

func (this *patrolController)ModifyPatrolLineAction(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
    w.Header().Set("charset", "utf-8")
    
	var errCode int
	var reason string    
    err := r.ParseForm()
    if err != nil {
		errCode = -1
		reason = "非法输入"    
        this.OutputUpdatePatrolLineResult(w, errCode, reason)
        return
    }
    
    routelineId := r.FormValue("routeline-id")
    routelineName := r.FormValue("routeline-name")
    routelineDescription := r.FormValue("routeline-description")

    if routelineName == "" || routelineDescription == "" {
		errCode = -1
		reason = "非法输入"    
        this.OutputUpdatePatrolLineResult(w, errCode, reason)
        return
    }

	routeLine := Routeline{}
	routeLine.Id, _ = strconv.Atoi(routelineId)
	routeLine.Name = routelineName
	routeLine.Description = routelineDescription

	_, result := ModifyRouteline(routeLine)
	if result {
		errCode = 0
	} else {
		errCode = 1
		reason = "操作失败"
	}

	reason = "修改成功"
    this.OutputUpdatePatrolLineResult(w, errCode, reason)
}

func (this *patrolController)DeletePatrolLineAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("charset", "utf-8")
	
	var errCode int
	var reason string
	id, err := strconv.Atoi(r.URL.RawQuery)
	if err != nil {
		errCode = -1
		reason = "非法输入"
		this.OutputUpdatePatrolLineResult(w, errCode, reason)
		return
	}
	
	_, result := DeleteRouteline(id)
	if result {
		errCode = 0
	} else {
		errCode = 1
		reason = "操作失败"
	}

    this.OutputUpdatePatrolLineResult(w, errCode, reason)
}
