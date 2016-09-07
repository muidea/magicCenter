package common

// View 基本视图
// ID 视图ID唯一标示该视图
// Name 视图名称
// Style 视图显示方式，链接方式还是内容方式
// URL 视图对应资源的URL
type View struct {
	ID    string
	Name  string
	Style int
	URL   string
}

// BlockView 块视图
// Tag 块视图标签
// Items 块视图包含的资源对象
type BlockView struct {
	View
	Tag   string
	Items []View
}

// PageView 页面视图
// Contents 页面内容信息
// Blocks 页面块信息
type PageView struct {
	View
	Contents []View
	Blocks   []BlockView
}
