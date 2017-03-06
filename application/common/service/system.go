package service

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
)

// System MagicCenter系统接口
type System interface {
	// StartUp 启动系统
	StartUp() error
	// Run 运行系统
	Run()
	// ShutDown 关闭系统
	ShutDown() error

	// Router 路由器
	Router() Router
	// ModuleHub 模块管理器
	ModuleHub() ModuleHub
	// Configuration 配置管理器
	Configuration() Configuration
	// Authority 权限管理器
	Authority() Authority

	// Session 当前Session
	Session(w http.ResponseWriter, r *http.Request) common.Session
}
