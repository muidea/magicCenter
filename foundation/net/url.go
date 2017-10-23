package net

import "path"

// JoinURL 合并Url路径
func JoinURL(prefix, subfix string) string {
	if len(subfix) > 0 && subfix[len(subfix)-1] != '/' {
		return path.Join(prefix, subfix)
	}

	return path.Join(prefix, subfix) + "/"
}

// SplitRESTAPI 分割出RestAPI的路径和ID
func SplitRESTAPI(url string) (string, string) {
	return path.Split(url)
}
