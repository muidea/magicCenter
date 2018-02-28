package model

// USER 用户类型
const USER = "user"

// GROUP 分组类型
const GROUP = "group"

const (
	// NEW 新建用户，未激活状态
	NEW = iota
	// ACTIVE 用户已经激活
	ACTIVE
	// DEACTIVE 用户未激活
	DEACTIVE
	// DISABLE 用户被禁用
	DISABLE
)

// User 用户信息
type User struct {
	//ID 用户ID,唯一标示该用户
	ID int `json:"id"`
	// Name 用户名称，允许用户修改也允许重名, 如果没有修改，则显示成账号
	Name string `json:"name"`
}

// UserDetail 用户详细信息
type UserDetail struct {
	User

	// Account 用户账号，不允许重复，唯一标示该用户
	Account string `json:"account"`
	//EMail 用户邮箱
	Email string `json:"email"`
	//Groups 所属分组
	Groups []int `json:"groups"`
	// Status 用户状态，预留字段，暂时没有用到
	Status int `json:"status"`
	// 注册时间
	RegisterTime string `json:"registerTime"`
}

const (
	// CommonGroup 普通组
	CommonGroup = iota
	// AdminGroup 管理员组
	AdminGroup
)

// Group 分组信息
// Name 名称
// Description 描述
// Catalog 类型（管理员组，普通组
type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Catalog     int    `json:"catalog"`
}

// AdminGroup 是否是管理员组
func (g *Group) AdminGroup() bool {
	return g.Catalog == AdminGroup
}

// AccountSummary 账号摘要信息
type AccountSummary []SummaryItem

// AccountItem 账号项
type AccountItem struct {
	Name          string `json:"name"`
	RegisterDate  string `json:"registerDate"`
	LastLoginDate string `json:"lastLoginDate"`
}
