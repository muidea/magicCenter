package model

/*
Module 模块类型
Id:标识该Module的字符串ID
Name:该Module的名称，Blog，Shop ect.
Description:该Module的描述信息
URL:该Module对应的URL
Type:该Module类型
Enable:是否启用该Module
*/
type Module struct {
	ID          string
	Name        string
	Description string
	URL         string
	Type        int
	Status      int
}
