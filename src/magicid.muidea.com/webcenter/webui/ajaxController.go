package webui
 
import (
    "net/http"
    "fmt"
    "log"
    "encoding/json"
    "muidea.com/dao"
)
 
type Result struct{
    Ret int
    Reason string
    Data interface{}
}
 
type ajaxController struct {
}
 
func (this *ajaxController)AjaxAction(w http.ResponseWriter, r *http.Request) {	
    w.Header().Set("content-type", "application/json")
    err := r.ParseForm()
    if err != nil {
        OutputJson(w, 0, "invalid Param", nil)
        return
    }
     
    admin_name := r.FormValue("admin_name")
    admin_password := r.FormValue("admin_password")
     
    if admin_name == "" || admin_password == "" {
        OutputJson(w, 0, "invalid Param", nil)
        return
    }
    
    dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magicid_db") 
    if err != nil {
    	log.Printf("open database failed, error:%s", err.Error());
    	OutputJson(w, 0, "connect database error", nil)
    	return 
    } 
    defer dao.Release()

	qSql :=  fmt.Sprintf("select * from magicid_db.user where name='%s' and password='%s'", admin_name, admin_password)
    
    if !dao.Query(qSql) {
    	log.Printf("query dataset failed, error:%s, sql:%s", err.Error(), qSql);
    	OutputJson(w, 0, "query database error", nil)
    	return
    }

	if !dao.Next() {
    	OutputJson(w, 0, "invalid name or password", nil)
    	return		
	}
    
    // 存入cookie,使用cookie存储
    cookie := http.Cookie{Name: "admin_name", Path: "/"}
    http.SetCookie(w, &cookie)
     
    OutputJson(w, 1, "Login OK", nil)
    return
}
 
func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
    out := &Result{ret, reason, i}
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}

