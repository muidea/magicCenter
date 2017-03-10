package util

import (
	"fmt"
	"strconv"
	"strings"
)

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
