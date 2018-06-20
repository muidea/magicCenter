package account

import (
	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/account/def"
	"muidea.com/magicCenter/module/modules/account/handler"
	"muidea.com/magicCenter/module/modules/account/route"
	common_def "muidea.com/magicCommon/common"
)

// LoadModule 加载模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {

	instance := &account{routes: make([]common.Route, 0), accountHandler: handler.CreateAccountHandler()}

	instance.routes = route.AppendUserRoute(instance.routes, instance.accountHandler)
	instance.routes = route.AppendGroupRoute(instance.routes, instance.accountHandler)

	moduleHub.RegisterModule(instance)
}

type account struct {
	routes         []common.Route
	accountHandler common.AccountHandler
}

func (instance *account) ID() string {
	return def.ID
}

func (instance *account) Name() string {
	return def.Name
}

func (instance *account) Description() string {
	return def.Description
}

func (instance *account) Group() string {
	return "kernel"
}

func (instance *account) Type() int {
	return common_def.KERNEL
}

func (instance *account) Status() int {
	return common_def.ACTIVE
}

func (instance *account) EntryPoint() interface{} {
	return instance.accountHandler
}

// Route Account 路由信息
func (instance *account) Routes() []common.Route {
	return instance.routes
}

// Startup 启动Account模块
func (instance *account) Startup() bool {
	return true
}

// Cleanup 清除Account模块
func (instance *account) Cleanup() {

}
