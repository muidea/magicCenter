package common

// 基本HTTP行为定义
const (
	GET    = "get"
	PUT    = "put"
	POST   = "post"
	DELETE = "delete"
)

// Route 路由接口
type Route interface {
	// Type 路由行为GET/PUT/POST/DELETE
	Type() string
	// Pattern 路由规则, 以'/'开始
	Pattern() string
	// Handler 路由处理器
	Handler() interface{}
}
