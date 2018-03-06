package model

// ARTICLE 文章类型
const ARTICLE = "article"

// CATALOG 分类类型
const CATALOG = "catalog"

// LINK 链接类型
const LINK = "link"

// MEDIA 图像类型
const MEDIA = "media"

// Article 文章
type Article Unit

// Catalog 分类
type Catalog Unit

// Link 链接
type Link Unit

// Media 文件
type Media Unit

// Summary 摘要信息
type Summary struct {
	Unit
	Catalog    []int  `json:"catalog"`
	CreateDate string `json:"createDate"`
	Creater    int    `json:"creater"`
}

// SummaryView 摘要信息显示视图
type SummaryView struct {
	Summary
	Catalog []Catalog `json:"catalog"`
	Creater User      `json:"creater"`
}

// ArticleDetail 文章
type ArticleDetail struct {
	Summary

	Content string `json:"content"`
}

// ArticleDetailView 文章显示信息
type ArticleDetailView struct {
	ArticleDetail
	Catalog []Catalog `json:"catalog"`
	Creater User      `json:"creater"`
}

// CatalogDetail 分类详细信息
type CatalogDetail struct {
	Summary

	Description string `json:"description"`
}

// CatalogDetailView 分类详细信息显示信息
type CatalogDetailView struct {
	CatalogDetail

	Catalog []Catalog `json:"catalog"`
	Creater User      `json:"creater"`
}

// LinkDetail 链接
type LinkDetail struct {
	Summary

	URL  string `json:"url"`
	Logo string `json:"logo"`
}

// LinkDetailView 链接显示信息
type LinkDetailView struct {
	LinkDetail

	Catalog []Catalog `json:"catalog"`
	Creater User      `json:"creater"`
}

// MediaDetail 文件信息
type MediaDetail struct {
	Summary
	URL         string `json:"url"`
	Description string `json:"description"`
}

// MediaDetailView 文件信息显示信息
type MediaDetailView struct {
	MediaDetail

	Catalog []Catalog `json:"catalog"`
	Creater User      `json:"creater"`
}

// ContentSummary 内容摘要信息
type ContentSummary []UnitSummary

// ContentUnit 内容项
type ContentUnit struct {
	Title      string `json:"title"`
	Type       string `json:"type"`
	CreateDate string `json:"createDate"`
}
