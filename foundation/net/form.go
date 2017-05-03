package net

import (
	"io"
	"net/http"
	"os"
	"path"
)

// MultipartFormFile 接受文件参数
// string 文件名
// error 错误码
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
		dst, err := os.Create(dstFile)
		if err != nil {
			break
		}

		defer dst.Close()
		_, err = io.Copy(dst, src)
		break
	}

	return dstFile, err
}
