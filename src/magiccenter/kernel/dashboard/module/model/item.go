package model

import (
	"fmt"

	"muidea.com/util"
)

/*
Item 记录项
ID:表示Item对象
Rid:对应对象的ID
Rtype:对应对象的类型，article,catalog,link
Owner: Item所属的Block
*/
type Item struct {
	ID    int
	Rid   int
	Rtype string
	Owner int
}

// URL 返回资源对应的URL
func (i Item) URL() string {
	param := fmt.Sprintf("?id=%d", i.Rid)

	return util.JoinURL(i.Rtype, param)
}
