package model

const (
	// CommonGroup 普通组
	CommonGroup = iota
	// AdminGroup 管理员组
	AdminGroup
)

// Group 分组信息
// Name 名称
// Description 描述
// Type 类型（管理员组，普通组
type Group struct {
	ID          int
	Name        string
	Description string
	Type        int
}

// AdminGroup 是否是管理员组
func (g *Group) AdminGroup() bool {
	return g.Type == AdminGroup
}
