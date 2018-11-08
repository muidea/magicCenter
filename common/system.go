package common

import (
	"muidea.com/magicCommon/def"
	"muidea.com/magicCommon/model"
)

// SystemHandler 系统管理接口
type SystemHandler interface {
	GetSystemProperty() model.SystemProperty

	UpdateSystemProperty(sysProperty model.SystemProperty) bool
	GetSystemStatistics() model.StatisticsView

	GetSystemMenu() (string, bool)

	QuerySyslog(source string, filter *def.PageFilter) ([]*model.Syslog, int)
	InsertSyslog(log *model.Syslog) bool
}
