package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
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

// Str2IntArray 字符串转换成数字数组
func Str2IntArray(str string) ([]int, bool) {
	ids := []int{}
	vals := strings.Split(str, ",")
	for _, val := range vals {
		id, err := strconv.Atoi(val)
		if err != nil {
			return ids, false
		}
		ids = append(ids, id)
	}

	return ids, true
}

// IntArray2Str 数字数组转字符串
func IntArray2Str(ids []int) string {
	if len(ids) == 0 {
		return ""
	}

	val := ""
	for _, v := range ids {
		val = fmt.Sprintf("%s,%d", val, v)
	}

	return val[1:]
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
