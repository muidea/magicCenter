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
	"webcenter/application"	
)

func ManageImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/content/image.html")
    if (err != nil) {
        log.Print(err)
        
        http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
        
    t.Execute(w, nil)
}

func QueryAllImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllImageHandler");
	
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
	log.Print("deleteImageHandler");
	
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
	log.Print("ajaxImageHandler");
	
	result := SubmitImageResult{}
	
	for true {
		param := SubmitImageParam{}
	    err := r.ParseForm()
    	if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
    	}
    	
		file, head, err := r.FormFile("image-name")
		if err != nil {
    		log.Print("paseform failed")
    		
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}
		
		fileName := fmt.Sprintf("%s/%s/%s_%s_%s", application.StaticPath(), application.UploadPath(), time.Now().Format("20060102150405"), util.RandomAlphabetic(16), head.Filename);
		
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
		
		desc := r.FormValue("image-desc")
	    accessCode := r.FormValue("accesscode")

		staticPath := application.StaticPath()
		param.url = fileName[len(staticPath):]
		param.desc = desc
		param.accessCode = accessCode

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
	
}
