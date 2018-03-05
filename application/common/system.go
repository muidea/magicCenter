package common

import "muidea.com/magicCenter/application/common/model"

// SystemHandler 系统管理接口
type SystemHandler interface {
	GetSystemProperty() model.SystemProperty

	UpdateSystemProperty(sysProperty model.SystemProperty) bool

	GetModuleList() []model.Module

	GetSystemStatistics() model.StatisticsView
}
