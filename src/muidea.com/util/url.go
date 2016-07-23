package util

import "path"

// JoinURL 合并Url路径
func JoinURL(prefix, subfix string) string {
	return path.Join(prefix, subfix)
}
