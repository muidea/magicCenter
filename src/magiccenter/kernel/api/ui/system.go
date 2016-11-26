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
