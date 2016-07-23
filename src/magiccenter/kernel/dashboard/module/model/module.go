package model

import (

)

/*
Id:标识该Module的字符串ID
Name:该Module的名称，Blog，Shop ect.
Description:该Module的描述信息
Uri:该Module对应的Uri
Enable:是否启用该Module
*/
type Module struct {
	Id string
	Name string
	Description string
	Uri string
	EnableFlag int
}

type ModuleLayout struct {
	Module
	Blocks []Block
	Pages []Page
}

type ModuleContent struct {
	Module
	Blocks []BlockDetail
}

