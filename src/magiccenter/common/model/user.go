package model

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
	ID int
	// Name 用户名称，允许用户修改也允许重名, 如果没有修改，则显示成账号
	Name string
}

// UserDetail 用户详细信息
type UserDetail struct {
	User

	// Account 用户账号，不允许重复，唯一标示该用户
	Account string
	//EMail 用户邮箱
	Email string
	// Status 用户状态，预留字段，暂时没有用到
	Status int
}
