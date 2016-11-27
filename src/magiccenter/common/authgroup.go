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
