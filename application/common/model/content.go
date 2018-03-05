package model

// ARTICLE 文章类型
const ARTICLE = "article"

// CATALOG 分类类型
const CATALOG = "catalog"

// LINK 链接类型
const LINK = "link"

// MEDIA 图像类型
const MEDIA = "media"

// Summary 摘要信息
type Summary struct {
	Unit
	Catalog    []int  `json:"catalog"`
	CreateDate string `json:"createDate"`
	Creater    int    `json:"creater"`
}

// SystemSummary 摘要信息显示信息
type SystemSummary struct {
	Summary
	Catalog []Unit `json:"catalog"`
	Creater Unit   `json:"creater"`
}

// ArticleDetail 文章
type ArticleDetail struct {
	Summary

	Content string `json:"content"`
}

// ArticleDetailView 文章显示信息
type ArticleDetailView struct {
	SystemSummary
	Content string `json:"content"`
}

// CatalogDetail 分类详细信息
type CatalogDetail struct {
	Summary

	Description string `json:"id"`
}

// CatalogDetailView 分类详细信息显示信息
type CatalogDetailView struct {
	SystemSummary

	Description string `json:"id"`
}

// LinkDetail 链接
type LinkDetail struct {
	Summary

	URL  string `json:"url"`
	Logo string `json:"logo"`
}

// LinkDetailView 链接显示信息
type LinkDetailView struct {
	SystemSummary

	URL  string `json:"url"`
	Logo string `json:"logo"`
}

// MediaDetail 文件信息
type MediaDetail struct {
	Summary
	URL  string `json:"url"`
	Desc string `json:"desc"`
}

// MediaDetailView 文件信息显示信息
type MediaDetailView struct {
	SystemSummary
	URL  string `json:"url"`
	Desc string `json:"desc"`
}

// ContentSummary 内容摘要信息
type ContentSummary []UnitSummary

// ContentUnit 内容项
type ContentUnit struct {
	Title      string `json:"title"`
	Type       string `json:"type"`
	CreateDate string `json:"createDate"`
}
