package util

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

// SplitParam 分割URL参数
func SplitParam(params string) map[string]string {
	result := make(map[string]string)

	for _, param := range strings.Split(params, "&") {
		items := strings.Split(param, "=")
		if len(items) == 2 {
			result[strings.ToLower(items[0])] = strings.ToLower(items[1])
		}
	}

	return result
}

// MultipartFormFile 接受文件参数
func MultipartFormFile(r *http.Request, field, dstPath string) (string, string, error) {
	dstFile := ""
	fileType := ""
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

		fileInfo, err := os.Stat(dstFile)
		if err == nil {
			items := strings.Split(fileInfo.Name(), ".")
			cnt := len(items)
			if cnt >= 2 {
				fileType = items[cnt-1]
			} else {
				fileType = "unknown"
			}
		}
		break
	}

	return dstFile, fileType, err
}
