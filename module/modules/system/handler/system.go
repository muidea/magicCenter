package handler

import (
	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/common/configuration"
	"muidea.com/magicCommon/model"
)

// CreateSystemHandler 新建SystemHandler
func CreateSystemHandler(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) common.SystemHandler {
	i := impl{moduleHub: moduleHub}

	return &i
}

type impl struct {
	moduleHub common.ModuleHub
}

func (s *impl) GetSystemProperty() model.SystemProperty {
	return configuration.GetSystemProperty()
}

func (s *impl) UpdateSystemProperty(sysProperty model.SystemProperty) bool {
	return configuration.UpdateSystemProperty(sysProperty)
}

func (s *impl) GetSystemStatistics() model.StatisticsView {
	info := model.StatisticsView{}
	contentModule, ok := s.moduleHub.FindModule(common.CotentModuleID)
	if ok {
		contentHandler := contentModule.EntryPoint().(common.ContentHandler)
		info.LastContent = contentHandler.GetLastContent(10)

		contentSummary := contentHandler.GetContentSummary()
		info.SystemSummary = append(info.SystemSummary, contentSummary...)
	}
	accountModule, ok := s.moduleHub.FindModule(common.AccountModuleID)
	if ok {
		accountHandler := accountModule.EntryPoint().(common.AccountHandler)
		info.LastAccount = accountHandler.GetLastAccount(10)

		accountSummary := accountHandler.GetAccountSummary()
		info.SystemSummary = append(info.SystemSummary, accountSummary...)
	}

	return info
}
