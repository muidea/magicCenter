package model

// OnlineAccountInfo 在线用户信息
type OnlineAccountInfo struct {
	User
	LoginTime  int64  // 登陆时间
	UpdateTime int64  // 更新时间
	Address    string // 访问IP
}

// ACL 访问控制列表
type ACL struct {
	ID        int
	URL       string
	Method    string
	Module    string
	Status    int
	AuthGroup []int
}
