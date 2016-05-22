package model

import (
	contentModel "magiccenter/kernel/content/model"
)

/*
Page 用来描述Url所对应页面包含的内容

Owner: Page所属Module
Url: Page对应的Url
Blocks: Page包含的Block列表
*/
type Page struct {
	Owner string
	Url string
	Blocks []Block
}

type PageView struct {
	Owner string
	Url string
	Blocks []BlockView
	Contents []contentModel.Article
}
