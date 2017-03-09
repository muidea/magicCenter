package model

// AuthGroup 授权组
type AuthGroup struct {
	ID          int
	Name        string
	Description string
	Type        int
}

// CreateAuthGroup 新建AuthGroup
// gname 分组名
// gdescription 分组描述
// gtype 分组类型
func CreateAuthGroup(gname, gdescription string, gtype int) AuthGroup {
	i := AuthGroup{}
	i.ID = -1
	i.Name = gname
	i.Description = gdescription
	i.Type = gtype

	return i
}
