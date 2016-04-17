package model

import (

)

/*
Block 用来定义页面块
导航栏、标签云，分类列表等


Id: Block ID
Name: Block 名称
Owner: Block所属的Module
Items: Block所拥有的Item列表
*/
type Block struct {
	Id int
	Name string
	Owner string
}

type BlockDetail struct {
	Block
	Items []Item
}
