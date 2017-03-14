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
	ID      int
	Name    string
	Catalog []int
}

// ArticleDetail 文章
type ArticleDetail struct {
	Summary

	Content    string
	CreateDate string
	Author     int
}

// CatalogDetail 分类详细信息
type CatalogDetail struct {
	Summary

	Description string
	CreateDate  string
	Author      int
}

// LinkDetail 链接
type LinkDetail struct {
	Summary

	URL        string
	Logo       string
	CreateDate string
	Author     int
}

// MediaDetail 文件信息
type MediaDetail struct {
	Summary
	URL        string
	Desc       string
	CreateDate string
	Author     int
}
