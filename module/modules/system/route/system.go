package route

import (
	"encoding/json"
	"net/http"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/system/def"
	common_def "muidea.com/magicCommon/common"
	common_result "muidea.com/magicCommon/common"
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

type getSystemConfigRoute struct {
	systemHandler common.SystemHandler
}

type getSystemConfigResult struct {
	common_result.Result
	SystemProperty model.SystemProperty `json:"systemProperty"`
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
	return common_def.MaintainerAuthGroup.ID
}

func (i *getSystemConfigRoute) getSystemConfigHandler(w http.ResponseWriter, r *http.Request) {
	result := getSystemConfigResult{}
	result.SystemProperty = i.systemHandler.GetSystemProperty()
	result.ErrorCode = common_result.Success

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type setSystemConfigRoute struct {
	systemHandler common.SystemHandler
}

type setSystemConfigResult struct {
	common_result.Result
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
	return common_def.MaintainerAuthGroup.ID
}

func (i *setSystemConfigRoute) setSystemConfigHandler(w http.ResponseWriter, r *http.Request) {
	result := setSystemConfigResult{}

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
			result.ErrorCode = common_result.Success
		} else {
			result.ErrorCode = common_result.Failed
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

	common_result.Result
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
	return common_def.MaintainerAuthGroup.ID
}

func (i *getSystemMenuRoute) getSystemMenuHandler(w http.ResponseWriter, r *http.Request) {
	result := getSystemMenuResult{}

	menu, ok := i.systemHandler.GetSystemMenu()
	if ok {
		result.Menu = menu
		result.ErrorCode = common_result.Success
	} else {
		result.ErrorCode = common_result.Failed
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
	common_result.Result
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
	return common_def.MaintainerAuthGroup.ID
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
