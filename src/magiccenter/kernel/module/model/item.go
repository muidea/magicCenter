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
Id: 标识Item对象
Name: Item名称，根据实际表示的对象来决定，Article为Title，Catalog为Name ect。
Url: 方位该对象的实际Url
Owner: Item所属的Block
*/
type ItemView struct {
	Id int
	Name string
	Url string
	Owner int
}


