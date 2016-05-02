package mail

import (
	"log"
	"muidea.com/util"
	"magiccenter/configuration"
)

func Startup() bool {
	return true
}

func Cleanup() {
	
}

func PostMail(to, subject, body string) bool {
	systemInfo := configuration.GetSystemInfo()
	
	err := util.SendMail(systemInfo.MailAccount, systemInfo.MailPassword, systemInfo.MailServer, to, subject, body, "html")
	if err != nil {
		log.Printf("sendMail fail, err:%s", err.Error())
		return false
	}
	
	return true
}

