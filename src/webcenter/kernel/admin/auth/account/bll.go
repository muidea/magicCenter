package account

import (
	"fmt"
	"log"
	"muidea.com/util"
	"webcenter/util/modelhelper"
	"webcenter/kernel"
)

func constructVerifyContent(account, action string) string {
	
	content := fmt.Sprintf("<p>Dear %s:</p> <p>http://%s/user/verify/?account=%s&action=%s</p><p>Best wishes</p><p>admin</p>", account, kernel.Domain(), account, action)
	
	return content
}

func createNewUser(model modelhelper.Model, account, email, groups string) bool {
	usr := newUser()
	usr.Account = account
	usr.Email = email
	usr.Group = groups
	
	if !model.BeginTransaction() {
		return false
	}
	
	if !SaveUser(model,usr) {
		model.Rollback()
		
		return false
	}
	
	user := kernel.MailAccount()
	password := kernel.MailPassword()
	mailSvr := kernel.MailServer()
	subject := "verify email notify"
	content := constructVerifyContent(account, "new")
	
	err := util.SendMail(user, password, mailSvr, email, subject, content, "html")
	if err != nil {
		log.Println(err)
		
		model.Rollback()
		
		return false
	}
	
	model.Commit()	
	return true	
}

func modifyUserMail(model modelhelper.Model, id int, email string) bool {
	usr, found := QueryUserById(model, id)
	if !found {
		return false
	}
	
	usr.Email = email
	if !SaveUser(model,usr) {
		return false
	}	
	
	user := kernel.MailAccount()
	password := kernel.MailPassword()
	mailSvr := kernel.MailServer()
	subject := "modify email notify"
	content := constructVerifyContent(usr.Account, "modify")
	
	err := util.SendMail(user, password, mailSvr, email, subject, content, "")
	if err != nil {
		return false
	}
	
	return true
}

func modifyUserGroup(model modelhelper.Model, id int, group string) bool {
	usr, found := QueryUserById(model, id)
	if !found {
		return false
	}
	
	usr.Group = group
	
	if !SaveUser(model,usr) {
		return false
	}
	
	return true
}

func updateUserInfo(model modelhelper.Model, id int, nickName, password string) bool {
	usr, found := QueryUserById(model, id)
	if !found {
		return false
	}
	
	usr.NickName = nickName
	usr.password = password
	usr.Status = 0
	if !SaveUser(model,usr) {
		return false
	}
	
	return true	
}

