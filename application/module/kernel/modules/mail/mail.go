package mail

import (
	"log"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/configuration"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/kernel/modulehub"
	"muidea.com/magicCenter/foundation/net"
)

// ID Mail模块ID
const ID = "7fe4a6fa-b73a-401f-bd37-71e76670d18c"

// Name Mail模块名称
const Name = "Magic EMail"

// Description Mail模块描述信息
const Description = "Magic 邮件模块"

// URL Mail模块Url
const URL string = "mail"

// LoadModule 加载Mail模块
func LoadModule(cfg configuration.Configuration, modHub modulehub.ModuleHub) {
	account, _ := cfg.GetOption(model.MailAccount)
	password, _ := cfg.GetOption(model.MailPassword)
	server, _ := cfg.GetOption(model.MailServer)
	instance := &mail{mailAccount: account, mailPassword: password, mailServer: server}

	modHub.RegisterModule(instance)
}

// SendMail 发送邮件
func SendMail(modHub modulehub.ModuleHub, usrMail string, subject, content string) bool {
	mailModule, found := modHub.FindModule(ID)
	if !found {
		panic("can't find mail module")
	}

	endPoint := mailModule.EndPoint()
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
	return ID
}

func (instance *mail) Name() string {
	return Name
}

func (instance *mail) Description() string {
	return Description
}

func (instance *mail) Group() string {
	return "util"
}

func (instance *mail) Type() int {
	return common.KERNEL
}

func (instance *mail) URL() string {
	return URL
}

func (instance *mail) Status() int {
	return 0
}

func (instance *mail) EndPoint() interface{} {
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
