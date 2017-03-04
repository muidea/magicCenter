package ui

import (
	"encoding/json"
	"log"
	"magiccenter/common"
	"magiccenter/system"
	"net/http"
)

const passwordMark = "********"

// SystemInfo 系统信息
type SystemInfo struct {
	SystemInfo common.SystemInfo
}

// GetSystemInfoActionHandler 获取系统信息处理器
func GetSystemInfoActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetSystemInfoActionHandler")

	result := SystemInfo{}

	configuration := system.GetConfiguration()

	result.SystemInfo = configuration.GetSystemInfo()
	result.SystemInfo.MailPassword = passwordMark

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// UpdateSystemInfo 更新系统信息
type UpdateSystemInfo struct {
	Result common.Result
}

// PutSystemInfoActionHandler 更新系统信息处理器
func PutSystemInfoActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PutSystemInfoActionHandler")

	result := UpdateSystemInfo{}

	configuration := system.GetConfiguration()
	systemInfo := configuration.GetSystemInfo()
	for true {
		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		name := r.FormValue("system-name")
		if len(name) > 0 {
			systemInfo.Name = name
		}
		logourl := r.FormValue("system-logo")
		if len(logourl) > 0 {
			systemInfo.Logo = logourl
		}
		domain := r.FormValue("system-domain")
		if len(domain) > 0 {
			systemInfo.Domain = domain
		}
		emailServer := r.FormValue("system-emailserver")
		if len(emailServer) > 0 {
			systemInfo.MailServer = emailServer
		}
		emailAccount := r.FormValue("system-emailaccount")
		if len(emailAccount) > 0 {
			systemInfo.MailAccount = emailAccount
		}
		emailPassword := r.FormValue("system-emailpassword")
		if len(emailPassword) > 0 && emailPassword != passwordMark {
			systemInfo.MailPassword = emailPassword
		}

		if configuration.UpdateSystemInfo(systemInfo) {
			result.Result.ErrCode = 0
		} else {
			result.Result.ErrCode = 1
			result.Result.Reason = "更新系统信息失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
