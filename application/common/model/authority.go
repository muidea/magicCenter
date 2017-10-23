package model

// AuthGroup 授权组
type AuthGroup struct {
	ID          int
	Name        string
	Description string
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
