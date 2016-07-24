package model

// CATALOG 分类类型
const CATALOG = "catalog"

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
