package net

import (
	"path"
	"strings"
)

// JoinURL 合并Url路径
func JoinURL(prefix, subfix string) string {
	return path.Join(prefix, subfix) + "/"
}

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

// ParseRestAPIUrl 解析RestAPI Url参数
// 主要用来解析类/user/1/?store=qqq 这样的restAPI url
func ParseRestAPIUrl(url string) (string, map[string]string) {
	subArray := strings.Split(url, "?")

}
