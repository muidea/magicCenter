package ui

import (
	"log"
	"net/http"
	"encoding/json"
	"html/template"
	"muidea.com/util"
	"magiccenter/kernel/common"
	"magiccenter/kernel/module/model"
	"magiccenter/kernel/module/bll"
	contentModel "magiccenter/kernel/content/model"
	contentBll "magiccenter/kernel/content/bll"
)


type ModuleContentView struct {
	Modules []model.Module
}

type QueryModuleContentResult struct {
	common.Result
	Module model.ModuleContent
	Articles []contentModel.ArticleSummary
	Catalogs []contentModel.Catalog
	Links []contentModel.Link
}

func ModuleContentHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ModuleContentHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/module/content.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
    
    view := ModuleContentView{}
    view.Modules = bll.QueryAllModules()
        
    t.Execute(w, view)
}


func QueryModuleContentHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryModuleContentHandler");
	
	result := QueryModuleContentResult{}
	
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
		}
		
		result.Module, found = bll.QueryModuleContent(id)
		if !found {
    		result.ErrCode = 1
    		result.Reason = "无效请求数据"
    		break;
		}
		
		result.Articles = contentBll.QueryAllArticleSummary()
		result.Catalogs = contentBll.QueryAllCatalog()
		result.Links = contentBll.QueryAllLink()

		result.ErrCode = 0
		result.Reason = "查询成功"
		break;				
	}
	
    b, err := json.Marshal(result)
    if err != nil {
    	panic("marshal failed, err:" + err.Error())
    }
    
    w.Write(b)
}




