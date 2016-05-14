package util

import (
	"net/http"
	"net/smtp"
	"strings"
	"os"
	"path"
	"io"
	"reflect"
)


func SplitParam(params string) map[string]string {
	result := make(map[string]string)
	
	for _, param := range strings.Split(params,"&") {
		items := strings.Split(param, "=")
		if len(items) == 2 {
			result[strings.ToLower(items[0])] = strings.ToLower(items[1])
		}
	}
	
	return result
}

func MultipartFormFile(r *http.Request, field, dstPath string) (string, error) {
	dstFile := ""
	var err error
	
	for true {
		src, head, err := r.FormFile(field)
		if err != nil {
			break
		}
		defer src.Close()
		
		_, err = os.Stat(dstPath)
		if err != nil {
			err = os.MkdirAll(dstPath, os.ModeDir)
		}
		if err != nil {
			break
		}
		
		dstFile = path.Join(dstPath, head.Filename)
		dst,err:=os.Create(dstFile)
		if err != nil {
			break			
		}
		
		defer dst.Close()
		_, err = io.Copy(dst, src)
		break
	}
	
	return dstFile, err
}

/*
 *  user : example@example.com login smtp server user
 *  password: xxxxx login smtp server password
 *  host: smtp.example.com:port   smtp.163.com:25
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */
func SendMail(user, password, host, to, subject, body, mailtype string) error{
    hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", user, password, hp[0])
    var content_type string
    if mailtype == "html" {
        content_type = "Content-Type: text/"+ mailtype + "; charset=UTF-8"
    }else{
        content_type = "Content-Type: text/plain" + "; charset=UTF-8"
    }
 
    msg := []byte("To: " + to + "\r\nFrom: " + user + "<"+ user +">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    send_to := strings.Split(to, ";")
    err := smtp.SendMail(host, auth, user, send_to, msg)
    return err
}

func ValidateFunc(fun interface{}) {
	if reflect.TypeOf(fun).Kind() != reflect.Func {
		panic("fun must be a callable func")
	}
}




