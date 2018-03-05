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

// NewEmptyUser 新建空用户
func NewEmptyUser() UserDetail {
	return UserDetail{}
}

// NewUser 新建用户
func NewUser(account, email string, groups []int, status int) UserDetail {
	return UserDetail{Account: account, Email: email, Group: groups, Status: status}
}

// User 用户信息
type User Unit

// UserDetail 用户详细信息
type UserDetail struct {
	User

	// Account 用户账号，不允许重复，唯一标示该用户
	Account string `json:"account"`
	//EMail 用户邮箱
	Email string `json:"email"`
	//Groups 所属分组
	Group []int `json:"group"`
	// Status 用户状态，预留字段，暂时没有用到
	Status int `json:"status"`
	// 注册时间
	RegisterTime string `json:"registerTime"`
}

// UserDetailView 用户详情显示信息
type UserDetailView struct {
	UserDetail

	Group []Unit `json:"group"`
}

const (
	// InvalidGroup 无效组
	InvalidGroup = iota
	// AdminGroup 管理组
	AdminGroup
	// CommonGroup 普通组
	CommonGroup
)

// NewEmptyGroup 新建空组
func NewEmptyGroup() GroupDetail {
	return GroupDetail{Unit: Unit{ID: InvalidGroup}}
}

// NewGroup 新建组
func NewGroup(name, description string, catalog int) GroupDetail {
	return GroupDetail{Unit: Unit{ID: InvalidGroup, Name: name}, Description: description, Catalog: catalog}
}

// Group 分组信息
type Group Unit

// GroupDetail 分组详情
// Name 名称
// Description 描述
// Catalog 类型（管理组，普通组)
type GroupDetail struct {
	Unit
	Description string `json:"description"`
	Catalog     int    `json:"catalog"`
}

// GroupDetailView 分组详情显示信息
type GroupDetailView struct {
	GroupDetail
	Catalog Unit `json:"catalog"`
}

// AdminGroup 是否是管理员组
func (g *GroupDetail) AdminGroup() bool {
	return g.Catalog == AdminGroup
}

// AccountSummary 账号摘要信息
type AccountSummary []UnitSummary

// AccountUnit 账号项
type AccountUnit struct {
	Name          string `json:"name"`
	RegisterDate  string `json:"registerDate"`
	LastLoginDate string `json:"lastLoginDate"`
}

// AccountRecord 账号记录
type AccountRecord []AccountUnit
