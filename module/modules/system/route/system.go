package route

import (
	"encoding/json"
	"net/http"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/system/def"
	common_const "muidea.com/magicCommon/common"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/foundation/net"
	"muidea.com/magicCommon/model"
)

// AppendSystemRoute 追加SystemRoute
func AppendSystemRoute(routes []common.Route, systemHandler common.SystemHandler) []common.Route {
	rt := GetSystemConfigRoute(systemHandler)
	routes = append(routes, rt)

	rt = SetSystemConfigRoute(systemHandler)
	routes = append(routes, rt)

	rt = GetSystemMenuRoute(systemHandler)
	routes = append(routes, rt)

	rt = GetSystemDashboardRoute(systemHandler)
	routes = append(routes, rt)

	rt = QuerySyslogRoute(systemHandler)
	routes = append(routes, rt)

	rt = InsertSyslogRoute(systemHandler)
	routes = append(routes, rt)

	return routes
}

// GetSystemConfigRoute 新建获取SystemConfig路由
func GetSystemConfigRoute(systemHandler common.SystemHandler) common.Route {
	return &getSystemConfigRoute{systemHandler: systemHandler}
}

// SetSystemConfigRoute 新建获取SystemConfig路由
func SetSystemConfigRoute(systemHandler common.SystemHandler) common.Route {
	return &setSystemConfigRoute{systemHandler: systemHandler}
}

// GetSystemMenuRoute 新建获取SystemMenu路由
func GetSystemMenuRoute(systemHandler common.SystemHandler) common.Route {
	return &getSystemMenuRoute{systemHandler: systemHandler}
}

// GetSystemDashboardRoute 新建获取SystemDashboard路由
func GetSystemDashboardRoute(systemHandler common.SystemHandler) common.Route {
	return &getSystemDashboardRoute{systemHandler: systemHandler}
}

// QuerySyslogRoute 新建查询Syslog路由
func QuerySyslogRoute(systemHandler common.SystemHandler) common.Route {
	return &querySyslogRoute{systemHandler: systemHandler}
}

// InsertSyslogRoute 新建插入日志路由
func InsertSyslogRoute(systemHandler common.SystemHandler) common.Route {
	return &insertSyslogRoute{systemHandler: systemHandler}
}

type getSystemConfigRoute struct {
	systemHandler common.SystemHandler
}

func (i *getSystemConfigRoute) Method() string {
	return common.GET
}

func (i *getSystemConfigRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetSystemProperty)
}

func (i *getSystemConfigRoute) Handler() interface{} {
	return i.getSystemConfigHandler
}

func (i *getSystemConfigRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *getSystemConfigRoute) getSystemConfigHandler(w http.ResponseWriter, r *http.Request) {
	result := common_def.QuerySystemConfigResult{}
	result.SystemProperty = i.systemHandler.GetSystemProperty()
	result.ErrorCode = common_def.Success

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type setSystemConfigRoute struct {
	systemHandler common.SystemHandler
}

func (i *setSystemConfigRoute) Method() string {
	return common.PUT
}

func (i *setSystemConfigRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutSystemConfig)
}

func (i *setSystemConfigRoute) Handler() interface{} {
	return i.setSystemConfigHandler
}

func (i *setSystemConfigRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *setSystemConfigRoute) setSystemConfigHandler(w http.ResponseWriter, r *http.Request) {
	result := common_def.UpdateSystemConfigResult{}

	for true {
		r.ParseForm()

		systemProperty := model.SystemProperty{}
		systemProperty.Name = r.FormValue("name")
		systemProperty.Description = r.FormValue("description")
		systemProperty.Logo = r.FormValue("logo")
		systemProperty.Domain = r.FormValue("domain")
		systemProperty.MailServer = r.FormValue("mailsvr")
		systemProperty.MailAccount = r.FormValue("mailaccount")
		systemProperty.MailPassword = r.FormValue("mailpassword")

		if i.systemHandler.UpdateSystemProperty(systemProperty) {
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
			result.Reason = "更新系统信息失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type getSystemMenuRoute struct {
	systemHandler common.SystemHandler
}

type getSystemMenuResult struct {
	Menu string `json:"menu"`

	common_def.Result
}

func (i *getSystemMenuRoute) Method() string {
	return common.GET
}

func (i *getSystemMenuRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetSystemMenu)
}

func (i *getSystemMenuRoute) Handler() interface{} {
	return i.getSystemMenuHandler
}

func (i *getSystemMenuRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *getSystemMenuRoute) getSystemMenuHandler(w http.ResponseWriter, r *http.Request) {
	result := getSystemMenuResult{}

	menu, ok := i.systemHandler.GetSystemMenu()
	if ok {
		result.Menu = menu
		result.ErrorCode = common_def.Success
	} else {
		result.ErrorCode = common_def.Failed
		result.Reason = "Get Menu failed"
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type getSystemDashboardRoute struct {
	systemHandler common.SystemHandler
}

type getSystemDashboardResult struct {
	model.StatisticsView
	common_def.Result
}

func (i *getSystemDashboardRoute) Method() string {
	return common.GET
}

func (i *getSystemDashboardRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetSystemDashboard)
}

func (i *getSystemDashboardRoute) Handler() interface{} {
	return i.getSystemDashboardHandler
}

func (i *getSystemDashboardRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *getSystemDashboardRoute) getSystemDashboardHandler(w http.ResponseWriter, r *http.Request) {
	result := getSystemDashboardResult{}
	result.ErrorCode = 0
	result.StatisticsView = i.systemHandler.GetSystemStatistics()

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type querySyslogRoute struct {
	systemHandler common.SystemHandler
}

func (i *querySyslogRoute) Method() string {
	return common.GET
}

func (i *querySyslogRoute) Pattern() string {
	return net.JoinURL(def.URL, def.QuerySyslog)
}

func (i *querySyslogRoute) Handler() interface{} {
	return i.querySyslogHandler
}

func (i *querySyslogRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *querySyslogRoute) querySyslogHandler(w http.ResponseWriter, r *http.Request) {
	result := common_def.QuerySyslogResult{}

	filter := &common_def.PageFilter{}
	filter.Decode(r)

	source := r.URL.Query().Get("sourceType")
	result.Syslog, result.Total = i.systemHandler.QuerySyslog(source, filter)

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type insertSyslogRoute struct {
	systemHandler common.SystemHandler
}

func (i *insertSyslogRoute) Method() string {
	return common.POST
}

func (i *insertSyslogRoute) Pattern() string {
	return net.JoinURL(def.URL, def.InsertSyslog)
}

func (i *insertSyslogRoute) Handler() interface{} {
	return i.insertSyslogHandler
}

func (i *insertSyslogRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *insertSyslogRoute) insertSyslogHandler(w http.ResponseWriter, r *http.Request) {
	result := common_def.InsertSyslogResult{}

	param := &common_def.InsertSyslogParam{}
	for {
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		if i.systemHandler.InsertSyslog(param.User, param.Operation, param.DateTime, param.Source) {
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
		}
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
