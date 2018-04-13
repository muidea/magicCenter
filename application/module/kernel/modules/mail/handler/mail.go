package handler

import (
	"log"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCommon/model"
	"muidea.com/magicCenter/foundation/net"
)

// CreateEMailHandler 新建MailHandler
func CreateEMailHandler(cfg common.Configuration, sessionRegistry common.SessionRegistry, modHub common.ModuleHub) common.MailHandler {

	return &impl{configuration: cfg}
}

type impl struct {
	configuration common.Configuration
}

func (s *impl) PostMail(mailList []string, subject, content string, attachment []string) {
	account, _ := s.configuration.GetOption(model.MailAccount)
	password, _ := s.configuration.GetOption(model.MailPassword)
	server, _ := s.configuration.GetOption(model.MailServer)
	err := net.SendMail(account, password, server, mailList, subject, content, attachment, "html")
	if err != nil {
		log.Printf("sendMail fail, err:%s", err.Error())
	}
}
