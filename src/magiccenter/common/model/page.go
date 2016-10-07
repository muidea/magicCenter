package model

/*
Page 用来描述Url所对应页面包含的内容

Owner: Page所属Module
URL: Page对应的Url
Blocks: Page包含的Block列表
*/
type Page struct {
	Owner  string
	URL    string
	Blocks []Block
}