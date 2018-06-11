package handler

import (
	"muidea.com/magicCenter/common"
	"muidea.com/magicCommon/model"
)

// CreateModuleRegistryHandler 新建SystemHandler
func CreateModuleRegistryHandler(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) common.ModuleRegistryHandler {
	i := impl{moduleHub: moduleHub}
	return &i
}

type impl struct {
	moduleHub common.ModuleHub
}

func (s *impl) GetModuleList() []model.Module {
	moduleList := []model.Module{}

	modules := s.GetModuleDetailList()
	for _, val := range modules {
		moduleList = append(moduleList, model.Module{ID: val.ID, Name: val.Name})
	}

	return moduleList
}

func (s *impl) GetModuleDetailList() []model.ModuleDetail {
	moduleList := []model.ModuleDetail{}
	modules := s.moduleHub.GetAllModule()
	for _, mod := range modules {
		item := model.ModuleDetail{ID: mod.ID(), Name: mod.Name(), Description: mod.Description(), Type: mod.Type(), Status: mod.Status()}

		moduleList = append(moduleList, item)
	}

	return moduleList
}

func (s *impl) QueryModuleByID(id string) (model.ModuleDetail, bool) {
	detail := model.ModuleDetail{}
	mod, ok := s.moduleHub.FindModule(id)
	if ok {
		detail.ID = mod.ID()
		detail.Name = mod.Name()
		detail.Description = mod.Description()
		detail.Type = mod.Type()
		detail.Status = mod.Status()
	}

	return detail, ok
}
