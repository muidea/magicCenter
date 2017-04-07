package util

import (
	"fmt"
	"strconv"
	"strings"
)

func cleanStr(str string) string {
	size := len(str)
	if size == 0 {
		return ""
	}

	val := str
	if str[0] == ',' {
		val = str[1:]
	}

	if str[size-1] == ',' {
		val = str[:size-1]
	}

	return strings.TrimSpace(val)
}

// Str2IntArray 字符串转换成数字数组
func Str2IntArray(str string) ([]int, bool) {
	ids := []int{}
	vals := strings.Split(cleanStr(str), ",")

	for _, val := range vals {
		if len(val) == 0 {
			continue
		}

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
