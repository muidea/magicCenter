package configuration

import (
	"log"
	"magiccenter/configuration/bll"
)

const (
	// AppName 应用名称
	AppName = "@application_name"

	// AppDomain 应用域名
	AppDomain = "@application_domain"

	// AppLogo 应用Logo
	AppLogo = "@application_logo"

	// MailServer 邮件服务器地址
	MailServer = "@system_mailServer"

	// MailAccount 邮件账号
	MailAccount = "@system_mailAccount"

	// MailPassword 邮件账号密码
	MailPassword = "@system_mailPassword"

	// SysDefaultModule 系统默认模块
	SysDefaultModule = "@system_defaultModule"

	// StaticPath 静态资源路径
	StaticPath = "@system_staticPath"
	// ResourcePath 应用资源路径
	ResourcePath = "@system_staticPath"

	// UploadPath 上传文件保存路径
	UploadPath = "@system_uploadPath"

	// AuthorithID 登陆会话鉴权ID
	AuthorithID = "@authorith_Id"
)

// SystemInfo 系统信息
type SystemInfo struct {
	Name         string
	Logo         string
	Domain       string
	MailServer   string
	MailAccount  string
	MailPassword string
}

var configInfoMap = map[string]string{}

// LoadConfig 加载系统配置信息
func LoadConfig() {
	log.Println("configuration initialize ...")

	keys := []string{AppName, AppDomain, AppLogo, MailServer, MailAccount, MailPassword, SysDefaultModule}

	configInfoMap = bll.GetConfigurations(keys)

	configInfoMap[StaticPath] = "static"
	configInfoMap[ResourcePath] = "template"
	configInfoMap[UploadPath] = "upload"
	configInfoMap[AuthorithID] = "@@@$$auth_Id@@@"
}

// UpdateSystemInfo 更新系统信息
func UpdateSystemInfo(info SystemInfo) bool {
	configs := map[string]string{}
	configs[AppName] = info.Name
	configs[AppDomain] = info.Domain
	configs[AppLogo] = info.Logo
	configs[MailServer] = info.MailServer
	configs[MailAccount] = info.MailAccount
	configs[MailPassword] = info.MailPassword

	configInfoMap[AppName] = info.Name
	configInfoMap[AppDomain] = info.Domain
	configInfoMap[AppLogo] = info.Logo
	configInfoMap[MailServer] = info.MailServer
	configInfoMap[MailAccount] = info.MailAccount
	configInfoMap[MailPassword] = info.MailPassword

	return bll.UpdateConfigurations(configs)
}

// GetSystemInfo 获取系统信息
func GetSystemInfo() SystemInfo {
	info := SystemInfo{}
	info.Name = configInfoMap[AppName]
	info.Domain = configInfoMap[AppDomain]
	info.Logo = configInfoMap[AppLogo]
	info.MailServer = configInfoMap[MailServer]
	info.MailAccount = configInfoMap[MailAccount]
	info.MailPassword = configInfoMap[MailPassword]

	return info
}

// GetOption 获取指定的配置项
func GetOption(name string) (string, bool) {
	value, found := configInfoMap[name]
	if !found {
		return bll.GetConfiguration(name)
	}

	return value, found
}

// SetOption 设置指定配置项
func SetOption(name, value string) bool {
	// 如果值没有变化则直接返回成功
	oldValue, found := configInfoMap[name]
	if found && oldValue == value {
		return true
	}

	if bll.UpdateConfiguration(name, value) {
		if found {
			// 如果之前已经在内存中Load过了，这里也需要把内存中得信息更新一下
			configInfoMap[name] = value
		}
		return true
	}

	return false
}
