package model

// ACL 访问控制列表
type ACL struct {
	ID        int
	URL       string
	Method    string
	Module    string
	Status    int
	AuthGroup []int
}
