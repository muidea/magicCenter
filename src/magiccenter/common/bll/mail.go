package bll

/*
 Package bll 邮件辅助，提供发送邮件功能
*/

import "magiccenter/module"

// MailModuleID Mail 模块ID
const MailModuleID = "7fe4a6fa-b73a-401f-bd37-71e76670d18c"

// PostBox 投递邮箱
type PostBox struct {
	UserList []string
	Subject  string
	Content  string
}

// PostMail 发送邮件
func PostMail(usrMail []string, subject, content string) bool {
	mailModule, found := module.FindModule(MailModuleID)
	if !found {
		panic("can't find mail module")
	}

	postBox := PostBox{}
	postBox.UserList = usrMail
	postBox.Subject = subject
	postBox.Content = content

	return mailModule.Invoke(&postBox, nil)
}
