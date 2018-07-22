package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/common/configuration"
	"muidea.com/magicCommon/model"
)

// CreateSystemHandler 新建SystemHandler
func CreateSystemHandler(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) common.SystemHandler {
	i := impl{configuration: configuration, moduleHub: moduleHub}

	return &i
}

type impl struct {
	configuration common.Configuration
	moduleHub     common.ModuleHub
}

func (s *impl) GetSystemProperty() model.SystemProperty {
	return configuration.GetSystemProperty()
}

func (s *impl) UpdateSystemProperty(sysProperty model.SystemProperty) bool {
	return configuration.UpdateSystemProperty(sysProperty)
}

func (s *impl) GetSystemMenu() (string, bool) {
	path, ok := s.configuration.GetOption(model.StaticPath)
	if !ok {
		return "", false
	}

	menuFile := fmt.Sprintf("%s/const/menu.json", path)
	fileHandler, err := os.Open(menuFile)
	if err != nil {
		log.Printf("open menufile failed, err:%s", err.Error())
		return "", false
	}

	data, err := ioutil.ReadAll(fileHandler)
	if err != nil {
		log.Printf("read all menu content failed.")
		return "", false
	}

	return string(data), true
}

func (s *impl) GetSystemStatistics() model.StatisticsView {
	info := model.StatisticsView{SystemSummary: []model.UnitSummary{}, LastContent: []model.ContentUnit{}, LastAccount: []model.AccountUnit{}}
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

	endpointModule, ok := s.moduleHub.FindModule(common.EndpointModuleID)
	if ok {
		endpointHandler := endpointModule.EntryPoint().(common.EndpointHandler)

		endpointSummary := endpointHandler.GetSummary()
		info.SystemSummary = append(info.SystemSummary, endpointSummary...)
	}

	casModule, ok := s.moduleHub.FindModule(common.CASModuleID)
	if ok {
		casHandler := casModule.EntryPoint().(common.CASHandler)

		casSummary := casHandler.GetSummary()
		info.SystemSummary = append(info.SystemSummary, casSummary...)
	}

	return info
}
