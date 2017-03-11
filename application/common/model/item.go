package model

/*
Item 记录项
ID:表示Item对象
Rid:对应对象的ID
Rtype:对应对象的类型，article,catalog,link
Owner: Item所属的Block
*/
type Item struct {
	ID    int
	Rid   int
	Rtype string
	Owner int
}
