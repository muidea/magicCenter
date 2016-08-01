package model

/*
Block 用来定义页面块
导航栏、标签云，分类列表等


ID: Block ID
Name: Block 名称
Tag: Block标记信息，用于客户端进行对象识别
Style: 显示风格，显示内容还是显示链接
Owner: Block所属的Module
Items: Block所拥有的Item列表
*/
type Block struct {
	ID    int
	Name  string
	Tag   string
	Style int
	Owner string
}

// BlockDetail Block 详情
type BlockDetail struct {
	Block
	Items []Item
}
