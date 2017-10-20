package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/configuration"
	"muidea.com/magicCenter/application/common/model"
)

// CreateSystemHandler 新建SystemHandler
func CreateSystemHandler(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) common.SystemHandler {
	i := impl{moduleHub: moduleHub}

	moduleList := []model.Module{}
	modules := moduleHub.QueryAllModule()
	for _, mod := range modules {
		item := model.Module{ID: mod.ID(), Name: mod.Name(), Description: mod.Description(), Type: mod.Type(), Status: mod.Status()}

		for _, rt := range mod.Routes() {
			r := model.Route{Pattern: rt.Pattern(), Method: rt.Method()}
			item.Route = append(item.Route, r)
		}

		moduleList = append(moduleList, item)
	}
	i.moduleList = moduleList

	return &i
}

type impl struct {
	moduleHub  common.ModuleHub
	moduleList []model.Module
}

func (s *impl) GetSystemConfig() model.SystemInfo {
	return configuration.GetSystemInfo()
}

func (s *impl) UpdateSystemConfig(sysInfo model.SystemInfo) bool {
	return configuration.UpdateSystemInfo(sysInfo)
}

func (s *impl) GetModuleList() []model.Module {
	return s.moduleList
}
