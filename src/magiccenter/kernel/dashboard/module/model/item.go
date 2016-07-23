package model

import (

)

/*
Id:表示Item对象
Rid:对应对象的ID
Rtype:对应对象的类型，article,catalog,link
Owner: Item所属的Block
*/
type Item struct {
	Id int
	Rid int
	Rtype int
	Owner int
}

/*
Id: Item对应真是对象的ID，article，Catalog，。。。
Name: Item名称，根据实际表示的对象来决定，Article为Title，Catalog为Name ect。
Url:访问Item对应的Url
*/
type ItemView struct {
	Id int
	Name string
	Url string
}


