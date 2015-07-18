package util

import (
	"strings"
)


func SplitParam(params string) map[string]string {
	result := make(map[string]string)
	
	for _, param := range strings.Split(params,"&") {
		items := strings.Split(param, "=")
		if len(items) == 2 {
			result[strings.ToLower(items[0])] = strings.ToLower(items[1])
		}
	}
	
	return result
}

