package model

// User 用户信息
type User struct {
	//ID 用户ID,唯一标示该用户
	ID int
	// Name 用户名称，允许用户修改也允许重名
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
	// Groups 用户所属分组
	Groups []int
}
