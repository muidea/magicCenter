package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCommon/model"
)

// CreateModuleRegistryHandler 新建SystemHandler
func CreateModuleRegistryHandler(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) common.ModuleRegistryHandler {
	i := impl{moduleHub: moduleHub}

	moduleList := []model.ModuleDetail{}
	modules := moduleHub.GetAllModule()
	for _, mod := range modules {
		item := model.ModuleDetail{ID: mod.ID(), Name: mod.Name(), Description: mod.Description(), Type: mod.Type(), Status: mod.Status()}

		moduleList = append(moduleList, item)
	}
	i.moduleList = moduleList

	return &i
}

type impl struct {
	moduleHub  common.ModuleHub
	moduleList []model.ModuleDetail
}

func (s *impl) GetModuleList() []model.Module {
	moduleList := []model.Module{}
	for _, val := range s.moduleList {
		moduleList = append(moduleList, model.Module{ID: val.ID, Name: val.Name})
	}

	return moduleList
}

func (s *impl) GetModuleDetailList() []model.ModuleDetail {
	return s.moduleList
}
