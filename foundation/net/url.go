package net

import "path"

import "strings"

// JoinURL 合并Url路径
func JoinURL(prefix, subfix string) string {
	if len(subfix) > 0 && subfix[len(subfix)-1] != '/' {
		return path.Join(prefix, subfix)
	}

	return path.Join(prefix, subfix) + "/"
}

// SplitResetAPI 分割出RestAPI的路径和ID
func SplitResetAPI(url string) (string, string) {
	urlPath, urlID := path.Split(url)
	if len(urlID) > 0 {
		return urlPath, urlID
	}

	return path.Split(strings.TrimRight(urlPath, "/"))
}
