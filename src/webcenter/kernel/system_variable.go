package kernel 

import (
)

var staticPath string = "static"
var resourceFilePath string = "template"
var uploadFilePath string = "upload"

type SystemInfo struct {
	Name string
	Logo string
	Domain string
	MailServer string
	MailAccount string
	MailPassword string
}

var systemInfo SystemInfo

func UpdateSystemInfo(info SystemInfo) {
	systemInfo = info
}

func GetSystemInfo() SystemInfo {
	return systemInfo
}

func Name() string {
	return systemInfo.Name
}

func Logo() string {
	return systemInfo.Logo
}

func Domain() string {
	return systemInfo.Domain
}

func MailServer() string {
	return systemInfo.MailServer
}

func MailAccount() string {
	return systemInfo.MailAccount
}

func MailPassword() string {
	return systemInfo.MailPassword
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



