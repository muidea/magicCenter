package bll

import (
	"muidea.com/magicCenter/application/common"
)

/*
 Package bll 邮件辅助，提供发送邮件功能
*/

// MailModuleID Mail 模块ID
const MailModuleID = "7fe4a6fa-b73a-401f-bd37-71e76670d18c"

// PostBox 投递邮箱
type PostBox struct {
	// 收件人
	UserList []string
	// 邮件主题
	Subject string
	// 邮件内容
	Content string
}

// PostMail 发送邮件
func PostMail(moduleHub common.ModuleHub, usrMail []string, subject, content string) bool {
	mailModule, found := moduleHub.FindModule(MailModuleID)
	if !found {
		panic("can't find mail module")
	}

	postBox := PostBox{}
	postBox.UserList = usrMail
	postBox.Subject = subject
	postBox.Content = content

	return mailModule.Invoke(&postBox, nil)
}
