package model

// LINK 链接类型
const LINK = "link"

// Link 链接
type Link struct {
	ID      int
	Name    string
	URL     string
	Logo    string
	Creater int

	Catalog []int
}
