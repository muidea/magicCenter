package model

// ARTICLE 文章类型
const ARTICLE = "article"

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
