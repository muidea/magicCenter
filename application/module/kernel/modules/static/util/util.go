package util

import "strings"
import "path"

// MergePath 合并路径
// 这里是特殊处理一下，必须将url前面模块的名称去掉
func MergePath(rootPath, basePath, url string) string {
	pos := strings.Index(url, "/")
	if pos == 0 {
		pos = strings.Index(url[pos+1:], "/")
	}

	// 这是由于path.Join会自动的将最后一个“/”去掉，所以这里特殊处理一下
	if url[len(url)-1] != '/' {
		return path.Join(rootPath, basePath, url[pos+1:])
	}

	return path.Join(rootPath, basePath, url[pos+1:]) + "/"
}
