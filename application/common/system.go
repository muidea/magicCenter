package common

import "muidea.com/magicCenter/application/common/model"

// SystemHandler 系统管理接口
type SystemHandler interface {
	GetSystemConfig() model.SystemInfo

	UpdateSystemConfig(sysInfo model.SystemInfo)
}
