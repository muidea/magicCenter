package model

// AuthGroup 授权组
type AuthGroup struct {
	ID          int
	Name        string
	Description string
	Module      string
}

// CreateAuthGroup 新建AuthGroup
// name 分组名
// description 分组描述
func CreateAuthGroup(name, description, module string) AuthGroup {
	i := AuthGroup{}
	i.ID = -1
	i.Name = name
	i.Description = description
	i.Module = module

	return i
}

// ACL 访问控制列表
type ACL struct {
	ID        int
	URL       string
	AuthGroup []int
}
