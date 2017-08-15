package net

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

// MultipartFormFile 接受文件参数
// string 文件名
// error 错误码
func MultipartFormFile(r *http.Request, field, dstPath string) (string, error) {
	dstFilePath := ""
	var retErr error

	for true {
		srcFile, head, err := r.FormFile(field)
		if err != nil {
			log.Printf("get file field failed, field:%s, err:%s", field, err.Error())
			retErr = err
			break
		}
		defer srcFile.Close()

		_, err = os.Stat(dstPath)
		if err != nil {
			err = os.MkdirAll(dstPath, os.ModePerm)
		}

		if err != nil {
			log.Printf("destination path is invalid, err:%s", err.Error())
			retErr = err
			break
		}
		dstFilePath = path.Join(dstPath, head.Filename)
		dstFile, err := os.Create(dstFilePath)
		if err != nil {
			log.Printf("create destination file failed, err:%s", err.Error())
			retErr = err
			break
		}

		defer dstFile.Close()
		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			log.Printf("copy destination file failed, err%s", err.Error())
		}

		retErr = err
		break
	}

	return dstFilePath, retErr
}
