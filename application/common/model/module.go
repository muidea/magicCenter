package model

// Route 功能入口
type Route struct {
	Pattern string
	Method  string
}

// Module 模块
type Module struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/*
ModuleDetail 模块类型
Id:标识该Module的字符串ID
Name:该Module的名称，Blog，Shop ect.
Description:该Module的描述信息
Type:该Module类型
Status:是否启用该Module
*/
type ModuleDetail struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        int     `json:"type"`
	Status      int     `json:"status"`
	Route       []Route `json:"route"`
}
