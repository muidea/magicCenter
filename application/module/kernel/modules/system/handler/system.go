package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/configuration"
	"muidea.com/magicCenter/application/common/model"
)

// CreateSystemHandler 新建SystemHandler
func CreateSystemHandler(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) common.SystemHandler {
	i := impl{}

	return &i
}

type impl struct {
}

func (s *impl) GetSystemConfig() model.SystemInfo {
	return configuration.GetSystemInfo()
}

func (s *impl) UpdateSystemConfig(sysInfo model.SystemInfo) {
	configuration.UpdateSystemInfo(sysInfo)
}
