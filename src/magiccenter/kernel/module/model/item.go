package model

import (

)


/*
Id: 标识Item对象
Name: Item名称，根据实际表示的对象来决定，Article为Title，Catalog为Name ect。
Url: 方位该对象的实际Url
Owner: Item所属的Block
*/
type Item struct {
	Id int
	Name string
	Url string
	Owner int
}

