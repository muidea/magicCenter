package model

// AuthGroup 授权组
type AuthGroup struct {
	ID          int
	Name        string
	Description string
}

// Route 功能入口
type Route struct {
	Pattern string
	Method  string
}

/*
Module 模块类型
Id:标识该Module的字符串ID
Name:该Module的名称，Blog，Shop ect.
Description:该Module的描述信息
Type:该Module类型
Status:是否启用该Module
*/
type Module struct {
	ID          string
	Name        string
	Description string
	Type        int
	Status      int
	Route       []Route
}
