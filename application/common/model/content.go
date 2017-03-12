package model

// ARTICLE 文章类型
const ARTICLE = "article"

// CATALOG 分类类型
const CATALOG = "catalog"

// LINK 链接类型
const LINK = "link"

// MEDIA 图像类型
const MEDIA = "media"

// ArticleSummary 文章摘要
type ArticleSummary struct {
	ID         int
	Title      string
	CreateDate string
	Catalog    []int
	Author     int
}

// Article 文章
type Article struct {
	ArticleSummary

	Content string
}

// Catalog 分类
type Catalog struct {
	ID   int
	Name string
}

// CatalogDetail 分类详细信息
type CatalogDetail struct {
	Catalog

	Creater int
	Parent  []int
}

// Link 链接
type Link struct {
	ID      int
	Name    string
	URL     string
	Logo    string
	Creater int

	Catalog []int
}

// MediaDetail 文件信息
// Name 名称
// URL 文件URL
// Desc 文件描述
// Creater 创建者
type MediaDetail struct {
	ID      int
	Name    string
	URL     string
	Desc    string
	Creater int
	Catalog []int
}

// ContentMeta 内容元数据
type ContentMeta struct {
	Subject     string
	Description string
	URL         string
}
