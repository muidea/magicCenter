package net

import (
	"path"
	"strings"
)

// JoinURL 合并Url路径
func JoinURL(prefix, subfix string) string {
	if len(subfix) > 0 && subfix[len(subfix)-1] != '/' {
		return path.Join(prefix, subfix)
	}

	return path.Join(prefix, subfix) + "/"
}

// SplitParam 分割URL参数
func SplitParam(params string) map[string]string {
	result := make(map[string]string)

	for _, param := range strings.Split(params, "&") {
		items := strings.Split(param, "=")
		if len(items) == 2 {
			if len(items[0]) > 0 && len(items[1]) > 0 {
				result[strings.ToLower(items[0])] = strings.ToLower(items[1])
			}
		}
	}

	return result
}

// ParseRestAPIUrl 解析RestAPI Url参数
// 主要用来解析类/user/1/?store=qqq 这样的restAPI url
// 返回值 string 对象id
// 返回值 map[string]string 参数对
func ParseRestAPIUrl(url string) (string, map[string]string, bool) {
	id := ""
	param := make(map[string]string)
	ret := false
	subArray := strings.Split(url, "?")
	urlItemArray := strings.Split(subArray[0], "/")
	if len(urlItemArray) < 3 {
		return id, param, ret
	}

	return urlItemArray[len(urlItemArray)-2], SplitParam(subArray[len(subArray)-1]), true
}
