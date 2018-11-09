package syslog

import (
	"fmt"

	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCommon/def"
	"muidea.com/magicCommon/model"
)

// QuerySyslog 查询系统日志
func QuerySyslog(helper dbhelper.DBHelper, source string, filter *def.PageFilter) ([]model.Syslog, int) {
	totalCount := 0
	sysLogs := []model.Syslog{}
	sql := fmt.Sprintf(`select count(id) from common_syslog`)
	if source != "" {
		sql = fmt.Sprintf(`select count(id) from common_syslog where source ='%s'`, source)
	}

	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&totalCount)
	}
	helper.Finish()

	limitVal := totalCount
	offsetVal := 0
	if filter != nil {
		limitVal = filter.PageSize
		offsetVal = filter.PageSize * (filter.PageNum - 1)
	}
	if offsetVal < 0 {
		offsetVal = 0
	}
	if offsetVal >= totalCount {
		return sysLogs, totalCount
	}

	sql = fmt.Sprintf(`select id, user, operation, datetime, source from common_syslog order by datetime desc limit %d offset %d`, limitVal, offsetVal)
	if source != "" {
		sql = fmt.Sprintf(`select id, user, operation, datetime, source from common_syslog where source ='%s' order by datetime desc limit %d offset %d`, source, limitVal, offsetVal)
	}

	helper.Query(sql)
	for helper.Next() {
		log := model.Syslog{}
		helper.GetValue(&log.ID, &log.User, &log.Operation, &log.DateTime, &log.Source)
		sysLogs = append(sysLogs, log)
	}
	helper.Finish()

	return sysLogs, totalCount
}

// InsertSyslog 插入一条日志
func InsertSyslog(helper dbhelper.DBHelper, user, operation, datetime, source string) bool {
	sql := fmt.Sprintf(`insert into common_syslog (user,operation,datetime,source) values ('%s','%s','%s','%s')`, user, operation, datetime, source)
	num, ok := helper.Execute(sql)
	return num == 1 && ok
}
