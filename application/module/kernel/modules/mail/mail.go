package mail

import (
	"log"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/mail/def"
	"muidea.com/magicCenter/foundation/net"
)

// LoadModule 加载Mail模块
func LoadModule(cfg common.Configuration, modHub common.ModuleHub) {
	account, _ := cfg.GetOption(model.MailAccount)
	password, _ := cfg.GetOption(model.MailPassword)
	server, _ := cfg.GetOption(model.MailServer)
	instance := &mail{mailAccount: account, mailPassword: password, mailServer: server}

	modHub.RegisterModule(instance)
}

// SendMail 发送邮件
func SendMail(modHub common.ModuleHub, usrMail string, subject, content string) bool {
	mailModule, found := modHub.FindModule(def.ID)
	if !found {
		panic("can't find mail module")
	}

	endPoint := mailModule.EntryPoint()
	switch endPoint.(type) {
	case *mail:
		return endPoint.(*mail).postMail(usrMail, subject, content)
	}

	return false
}

type mail struct {
	mailAccount  string
	mailPassword string
	mailServer   string
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
	return instance
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

func (instance *mail) postMail(to, subject, body string) bool {
	err := net.SendMail(instance.mailAccount, instance.mailPassword, instance.mailServer, to, subject, body, "html")
	if err != nil {
		log.Printf("sendMail fail, err:%s", err.Error())
		return false
	}

	return true
}
