package model

import (

)

/*
Entity 为Module对应的实例

Id: Entity对应的字符串ID
Name: Entity名称
Description: Entity的描述信息
EnableFlag: 是否启用该Entity
DefaultFlag: 是否为默认Entity
Module: 所属Module ID
*/

type Entity struct {
	Id string
	Name string
	Description string
	EnableFlag int
	DefaultFlag int
	Module string
}

