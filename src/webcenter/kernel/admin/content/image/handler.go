package image

import (
	"net/http"
	"encoding/json"
	"html/template"
	"strings"
	"log"
	"time"
	"fmt"
	"os"
	"io"
	"strconv"
	"muidea.com/util"
	"webcenter/kernel"
	"webcenter/util/session"
	"webcenter/kernel/admin/common"		
)

type ManageView struct {
	Accesscode string
	ImageInfo []ImageInfo
}

type EditView struct {
	common.Result
	Accesscode string
	Id int
	Name string
	Url string
	Desc string
	Catalog []int
}

func ManageImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageImageHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	session := session.GetSession(w,r)
    t, err := template.ParseFiles("template/html/admin/content/image.html")
    if (err != nil) {
    	panic("parse files failed");
    }
    
	controller := &imageController{}
	info := controller.queryManageInfoAction()
    
    view := ManageView{}
    view.Accesscode = session.AccessToken()
    view.ImageInfo = info.ImageInfo
    
    t.Execute(w, view)    
}

func QueryAllImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryAllImageHandler");
	
	result := QueryAllImageResult{}
	
	for true {
		param := QueryAllImageParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
	    accessCode := r.FormValue("accesscode")
		param.accessCode = accessCode

    	controller := &imageController{}
    	result = controller.queryAllImageAction(param)
    	
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

func DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteImageHandler");
	
	result := DeleteImageResult{}
	
	for true {
		param := DeleteImageParam{}
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
		
		log.Printf("id:%d, accessCode:%s", param.id, param.accessCode);
		 
	    controller := &imageController{}
	    result = controller.deleteImageAction(param)
    	
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

func AjaxImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxImageHandler");
	
	result := SubmitImageResult{}
	
	session := session.GetSession(w,r)
	for true {
		param := SubmitImageParam{}
	    err := r.ParseMultipartForm(0)
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		file, head, err := r.FormFile("image-url")
		if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}
		
		fileName := fmt.Sprintf("%s/%s/%s_%s_%s", kernel.StaticPath(), kernel.UploadPath(), time.Now().Format("20060102150405"), util.RandomAlphabetic(16), head.Filename);
		
		log.Print(fileName)
		defer file.Close()
		f,err:=os.Create(fileName)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break			
		}
		
		defer f.Close()
		io.Copy(f,file)		
		
		name := r.FormValue("image-name")
		desc := r.FormValue("image-desc")
	    accessCode := r.FormValue("accesscode")
	    catalog := r.MultipartForm.Value["image-catalog"]

		staticPath := kernel.StaticPath()
		param.name = name
		param.url = fileName[len(staticPath):]
		param.desc = desc
		param.accessCode = accessCode
	    for _, ca := range catalog {
			cid, err := strconv.Atoi(ca)
		    if err != nil {
		    	log.Print("parse catalog failed, catalog:%s", ca)
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
		    }
		    
		    param.catalog = append(param.catalog, cid)
	    }
	    param.creater, _ = session.GetAccountId()
	    
    	controller := &imageController{}
    	result = controller.submitImageAction(param)
    	
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

func EditImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("EditImageHandler");
	
	result := EditView{}
	
	for true {
		param := EditImageParam{}
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
		    	
    	controller := &imageController{}
    	img := controller.editImageAction(param)
    	
    	result.ErrCode = img.ErrCode
    	result.Reason = img.Reason
    	result.Id = img.Image.Id()
    	result.Url = img.Image.Url()
    	result.Desc = img.Image.Desc()
    	for _, c := range img.Image.Relative() {
    		result.Catalog = append(result.Catalog, c.Id())
    	}
    	
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
