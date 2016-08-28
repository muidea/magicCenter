package mail

import (
	"log"
	"magiccenter/common"
	"magiccenter/configuration"
	"magiccenter/module"
	"magiccenter/router"

	"muidea.com/util"
)

// ID Mail模块ID
const ID = "17d73fc4-3a77-4eb1-958d-9f0b93ad6a4f"

// Name Mail模块名称
const Name = "Magic EMail"

// Description Mail模块描述信息
const Description = "Magic 邮件模块"

// URL Mail模块Url
const URL string = "mail"

// PostBox 投递邮箱
type PostBox struct {
	UserList []string
	Subject  string
	Content  string
}

type mail struct {
}

var instance *mail

// LoadModule 加载Mail模块
func LoadModule() {
	if instance == nil {
		instance = &mail{}
	}

	module.RegisterModule(instance)
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

func (instance *mail) EndPoint() common.EndPoint {
	return nil
}

// Route Mail 路由信息
func (instance *mail) Routes() []router.Route {
	routes := []router.Route{}

	return routes
}

// Startup 启动Mail模块
func (instance *mail) Startup() bool {
	return true
}

// Cleanup 清除Mail模块
func (instance *mail) Cleanup() {

}

// Invoke 执行外部命令
func (instance *mail) Invoke(param interface{}) bool {
	postBox := param.(PostBox)
	if postBox == nil {
		log.Print("illegal param")
		return false
	}

	go postMails(postBox)
	return true
}

func postMails(postBox PostBox) {
	for _, user := range postBox.UserList {
		postMail(user, postBox.Subject, postBox.Content)
	}
}

func postMail(to, subject, body string) bool {
	systemInfo := configuration.GetSystemInfo()

	err := util.SendMail(systemInfo.MailAccount, systemInfo.MailPassword, systemInfo.MailServer, to, subject, body, "html")
	if err != nil {
		log.Printf("sendMail fail, err:%s", err.Error())
		return false
	}

	return true
}
