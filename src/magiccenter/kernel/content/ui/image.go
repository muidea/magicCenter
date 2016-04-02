package ui

import (
	"log"
	"strconv"
	"time"
	"path"
	"net/http"
	"encoding/json"
	"html/template"
	"muidea.com/util"
	"magiccenter/configuration"
	"magiccenter/kernel/common"	
	"magiccenter/kernel/content/model"
	"magiccenter/kernel/content/bll"
)


type ManageImageView struct {
	Images []model.ImageDetail
	Catalogs []model.CatalogDetail
}

type QueryAllImageResult struct {
	Images []model.ImageDetail
}

type QueryImageResult struct {
	common.Result
	Image model.ImageDetail
}

type DeleteImageResult struct {
	common.Result
}

type AjaxImageResult struct {
	common.Result
}

type EditImageResult struct {
	common.Result
	Image model.ImageDetail	
}

//
// Image管理主界面
// 显示Image列表信息
// 返回html页面
// 
func ManageImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageImageHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/content/image.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
    view := ManageImageView{}
    view.Images = bll.QueryAllImage()
    
    t.Execute(w, view)    
}

//
// 查询全部Image
// 返回json
//
func QueryAllImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryAllImageHandler");
	
	result := QueryAllImageResult{}
	result.Images = bll.QueryAllImage()
		
    b, err := json.Marshal(result)
    if err != nil {
    	panic("json.Marshal, failed, err:" + err.Error())
    }
        
    w.Write(b)    
}


//
// 查询指定Image内容
// 返回json
//
func QueryImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryImageHandler");
	
	result := QueryImageResult{}
	
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
		
		image, found := bll.QueryImageById(aid)
		if !found {
			result.ErrCode = 1
			result.Reason = "操作失败"
			break			
		}
		 
		result.Image = image
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
// 删除指定Image
// 返回json
//
func DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteImageHandler");
	
	result := DeleteImageResult{}
	
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
		
		if !bll.DeleteImageById(aid) {
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
// 保存Image
// 返回json
//
func AjaxImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxCatalogHandler");
	
	result := AjaxImageResult{}
	
	for true {
		err := r.ParseMultipartForm(0)
    	if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		id, err := strconv.Atoi(r.FormValue("image-id"))
	    if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
	    }
	    
		name := r.FormValue("image-name")
		
		staticPath, _ := configuration.GetOption(configuration.STATIC_PATH)
		uploadPath, _ := configuration.GetOption(configuration.UPLOAD_PATH)
		filePath := path.Join(staticPath, uploadPath,time.Now().Format("20060102150405"))

		url, err := util.MultipartFormFile(r, "image-url", filePath)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}
	    	    
		desc := r.FormValue("image-desc")
	    catalog := r.MultipartForm.Value["image-catalog"]
	    catalogs :=[]int{}
	    for _, c := range catalog {
			cid, err := strconv.Atoi(c)
		    if err != nil {
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
		    }
		    
		    catalogs = append(catalogs, cid)
	    }
	    
	    if !bll.SaveImage(id, name, url, desc,100, catalogs) {
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
// 编辑Image
// 返回Image内容和当前可用Catalog
//
func EditImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("EditImageHandler");
	
	result := EditImageResult{}
	
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
		
		image, found := bll.QueryImageById(aid)
		if !found {
			result.ErrCode = 1
			result.Reason = "操作失败"
			break			
		}		
    	
    	result.Image = image
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



