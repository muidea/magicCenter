package mail

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/mail/def"
	"muidea.com/magicCenter/application/module/kernel/modules/mail/handler"
)

// LoadModule 加载Mail模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	account, _ := configuration.GetOption(model.MailAccount)
	password, _ := configuration.GetOption(model.MailPassword)
	server, _ := configuration.GetOption(model.MailServer)
	instance := &mail{mailAccount: account, mailPassword: password, mailServer: server, mailHandler: handler.CreateEMailHandler(configuration, sessionRegistry, moduleHub)}

	moduleHub.RegisterModule(instance)
}

type mail struct {
	mailAccount  string
	mailPassword string
	mailServer   string
	mailHandler  common.MailHandler
}

func (instance *mail) ID() string {
	return def.ID
}

func (instance *mail) Name() string {
	return def.Name
}

func (instance *mail) Description() string {
	return def.Description
}

func (instance *mail) Group() string {
	return "util"
}

func (instance *mail) Type() int {
	return common.KERNEL
}

func (instance *mail) Status() int {
	return 0
}

func (instance *mail) EntryPoint() interface{} {
	return instance.mailHandler
}

func (instance *mail) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	return groups
}

// Route Mail 路由信息
func (instance *mail) Routes() []common.Route {
	routes := []common.Route{}

	return routes
}

// Startup 启动Mail模块
func (instance *mail) Startup() bool {
	return true
}

// Cleanup 清除Mail模块
func (instance *mail) Cleanup() {

}
