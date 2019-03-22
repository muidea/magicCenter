package common

import (
	common_util "github.com/muidea/magicCommon/foundation/util"
	"github.com/muidea/magicCommon/model"
)

// SystemHandler 系统管理接口
type SystemHandler interface {
	GetSystemProperty() model.SystemProperty

	UpdateSystemProperty(sysProperty model.SystemProperty) bool
	GetSystemStatistics() model.StatisticsView

	GetSystemMenu() (string, bool)

	QuerySyslog(source string, filter *common_util.PageFilter) ([]model.Syslog, int)
	InsertSyslog(user, operation, datetime, source string) bool
}
