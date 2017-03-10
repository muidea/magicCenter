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

type route struct {
	rType    string
	rPattern string
	rHandler interface{}
}

// Type 路由行为GET/POST
func (r *route) Type() string {
	return r.rType
}

// Pattern 路由规则
func (r *route) Pattern() string {
	return r.rPattern
}

// Handler 路由处理器
func (r *route) Handler() interface{} {
	return r.rHandler
}

// NewRoute 新建一个路由对象
func NewRoute(rType, rPattern string, rHandler interface{}) Route {
	r := route{}
	r.rType = rType
	r.rPattern = rPattern
	r.rHandler = rHandler

	return &r
}
