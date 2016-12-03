package common

// AuthGroup 授权组
type AuthGroup interface {
	// 组ID
	ID() int
	// 组名称
	Name() string
	// 组描述
	Description() string
	// 组类型
	Type() int
}

type impl struct {
	gid          int
	gname        string
	gdescription string
	gtype        int
}

func (i *impl) ID() int {
	return i.gid
}

func (i *impl) Name() string {
	return i.gname
}

func (i *impl) Description() string {
	return i.gdescription
}

func (i *impl) Type() int {
	return i.gtype
}

// CreateAuthGroup 新建AuthGroup
func CreateAuthGroup(gname, gdescription string, gtype, gid int) AuthGroup {
	i := &impl{}
	i.gid = gid
	i.gname = gname
	i.gdescription = gdescription
	i.gtype = gtype

	return i
}
