package model

import (

)

/*
Block 用来定义页面块
导航栏、标签云，分类列表等


Id: Block ID
Name: Block 名称
Style: 显示风格，显示内容还是显示链接
Owner: Block所属的Module
Items: Block所拥有的Item列表
*/
type Block struct {
	Id int
	Name string
	Style int
	Owner string
}

type BlockDetail struct {
	Block
	Article []Item
	Catalog []Item
	Link []Item
}
