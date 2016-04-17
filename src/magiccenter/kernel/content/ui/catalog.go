package ui

import (
	"log"
	"strconv"
	"net/http"
	"encoding/json"
	"html/template"
	"muidea.com/util"
    "magiccenter/session"
    "magiccenter/configuration"	
	"magiccenter/kernel/common"
	accountModel "magiccenter/kernel/account/model"	
	"magiccenter/kernel/content/model"
	"magiccenter/kernel/content/bll"
)

type ManageCatalogView struct {
	Catalogs []model.CatalogDetail
}

type QueryAllCatalogResult struct {
	Catalogs []model.CatalogDetail
}

type QueryCatalogResult struct {
	common.Result
	Catalog model.CatalogDetail
}

type DeleteCatalogResult struct {
	common.Result
}

type AjaxCatalogResult struct {
	common.Result
}

type EditCatalogResult struct {
	common.Result
	Catalog model.CatalogDetail	
	AvalibleParent []model.Catalog
}

//
// 分类管理主界面
// 显示Catalog列表信息
// 返回html页面
// 
func ManageCatalogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	log.Print("ManageCatalogHandler");
	
    t, err := template.ParseFiles("template/html/admin/content/catalog.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
    view := ManageCatalogView{}
    view.Catalogs = bll.QueryAllCatalogDetail()
    
    t.Execute(w, view)    
}

//
// 查询全部Catalog
// 返回json
//
func QueryAllCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryAllCatalogHandler");
	
	result := QueryAllCatalogResult{}
	result.Catalogs = bll.QueryAllCatalogDetail()
		
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, failed, err:" + err.Error())
    }
        
    w.Write(b)    
}


//
// 查询指定Catalog内容
// 返回json
//
func QueryCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryCatalogHandler");
	
	result := QueryCatalogResult{}
	
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
    	params := util.SplitParam(r.URL.RawQuery)
    	id, found := params["id"]
    	if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
    	aid, err := strconv.Atoi(id)
    	if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break    		
    	}
		
		catalog, found := bll.QueryCatalogById(aid)
		if !found {
			result.ErrCode = 1
			result.Reason = "操作失败"
			break			
		}
		 
		result.Catalog = catalog
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


//
// 删除指定Catalog
// 返回json
//
func DeleteCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteCatalogHandler");
	
	result := DeleteCatalogResult{}
	
	for true {
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
    	params := util.SplitParam(r.URL.RawQuery)
    	id, found := params["id"]
    	if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
    	aid, err := strconv.Atoi(id)
    	if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break    		
    	}
		
		if !bll.DeleteCatalog(aid) {
			result.ErrCode = 1
			result.Reason = "操作失败"
			break
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

//
// 保存Catalog
// 返回json
//
func AjaxCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxCatalogHandler");
	
	authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
	if !found {
		panic("unexpected, can't fetch authorith id")
	}
	
	session := session.GetSession(w, r)
	user, found := session.GetOption(authId)
	if !found {
		panic("unexpected, must login system first.")
	}
	
	result := AjaxCatalogResult{}
	
	for true {
		err := r.ParseMultipartForm(0)
    	if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
    	id := r.FormValue("catalog-id")
		name := r.FormValue("catalog-name")
		parent := r.MultipartForm.Value["catalog-parent"]
		
		aid, err := strconv.Atoi(id)
	    if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    
	    parents :=[]int{}
	    for _, p := range parent {
			pid, err := strconv.Atoi(p)
		    if err != nil {
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
		    }
		    
		    parents = append(parents, pid)
	    }
	    
	    
	    if !bll.SaveCatalog(aid, name, user.(accountModel.UserDetail).Id, parents) {
			result.ErrCode = 1
			result.Reason = "操作失败"
			break	    	
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

//
// 编辑Catalog
// 返回Catalog内容和当前可用Parent Catalog
//
func EditCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("EditCatalogHandler");
	
	result := EditCatalogResult{}
	
	for true {
    	params := util.SplitParam(r.URL.RawQuery)
    	id, found := params["id"]
    	if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
    	aid, err := strconv.Atoi(id)
    	if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break    		
    	}
		
		catalog, found := bll.QueryCatalogById(aid)
		if !found {
			result.ErrCode = 1
			result.Reason = "操作失败"
			break			
		}		
    	
    	result.Catalog = catalog
    	result.AvalibleParent = bll.QueryAvalibleParentCatalog(catalog.Id)
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



