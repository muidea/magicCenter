package common

// 模块类型
const (
	// 内核模块，不能被禁用
	KERNEL = iota
	// 内置模块，属于系统自带可选模块，可以被禁用
	INTERNAL
	// 外部模块，通过外部接口注册进来的模块，可以被禁用
	EXTERNAL
)

// Module 功能模块接口
type Module interface {
	ID() string
	Name() string
	Description() string
	Group() string
	Type() int
	// URL 模块Url，每个模块都对应唯一的Url
	URL() string
	// 状态
	Status() int

	// EndPoint 模块提供的Rest api支持
	EndPoint() EndPoint

	// AuthGroups 授权组信息
	AuthGroups() []AuthGroup

	// Routes 模块支持的路由信息
	Routes() []Route

	//Startup 启动模块
	Startup() bool

	// Cleanup 清除模块
	Cleanup()

	// Invoke 执行指定操作，实际由各个模块具体定义实现
	// interface 返回结果
	Invoke(param interface{}, result interface{}) bool
}
