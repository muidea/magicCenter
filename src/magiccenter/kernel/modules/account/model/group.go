package model

const (
	// AdminGroup 管理员组
	AdminGroup = iota
	// CommonGroup 普通组
	CommonGroup
)

// Group 分组信息
type Group struct {
	ID   int
	Name string
	Type int
}

// AdminGroup 是否是管理员组
func (g *Group) AdminGroup() bool {
	return g.Type == AdminGroup
}
