package model

import (
	contentModel "magiccenter/kernel/modules/content/model"
)

/*
Page 用来描述Url所对应页面包含的内容

Owner: Page所属Module
Url: Page对应的Url
Blocks: Page包含的Block列表
*/
type Page struct {
	Owner  string
	Url    string
	Blocks []Block
}

type Content struct {
	contentModel.Article
	Url string
}

type PageView struct {
	Owner  string
	Url    string
	Blocks []BlockView
	Posts  []Content
}

type PageContentView struct {
	PageView
	Content contentModel.Article
}

type PageCatalogView struct {
	PageView
	Catalogs []ItemView
}
