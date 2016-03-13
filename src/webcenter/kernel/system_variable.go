package kernel 

import (
)

var staticPath string = "static"
var resourceFilePath string = "template"
var uploadFilePath string = "upload"

var systemName string
var systemLogo string
var systemDomain string
var systemMailServer string
var systemMailAccount string
var systemMailPassword string


func Name() string {
	return systemName
}

func UpdateName(name string) {
	systemName = name
}

func Logo() string {
	return systemLogo
}

func UpdateLogo(logo string) {
	systemLogo = logo
}

func Domain() string {
	return systemDomain
}

func UpdateDomain(domain string) {
	systemDomain = domain
}

func MailServer() string {
	return systemMailServer
}

func UpdateMailServer(server string) {
	systemMailServer = server
}

func MailAccount() string {
	return systemMailAccount
}

func UpdateMailAccount(account string) {
	systemMailAccount = account
}

func MailPassword() string {
	return systemMailPassword
}

func UpdateMailPassword(password string) {
	systemMailPassword = password
}

func UploadPath() string {
	return uploadFilePath
}

func ResourcePath() string {
	return resourceFilePath
}

func StaticPath() string {
	return staticPath
}



