package route

import (
	"encoding/json"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/system/def"
	"muidea.com/magicCenter/foundation/net"
)

// AppendSystemRoute 追加SystemRoute
func AppendSystemRoute(routes []common.Route, systemHandler common.SystemHandler) []common.Route {
	rt := GetSystemConfigRoute(systemHandler)
	routes = append(routes, rt)

	rt = SetSystemConfigRoute(systemHandler)
	routes = append(routes, rt)

	rt = GetModulesRoute(systemHandler)
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

// GetModulesRoute 新建获取Modules路由
func GetModulesRoute(systemHandler common.SystemHandler) common.Route {
	return &getModulesRoute{systemHandler: systemHandler}
}

type getSystemConfigRoute struct {
	systemHandler common.SystemHandler
}

type getSystemConfigResult struct {
	common.Result
	SystemInfo model.SystemInfo
}

func (i *getSystemConfigRoute) Method() string {
	return common.GET
}

func (i *getSystemConfigRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetSystemConfig)
}

func (i *getSystemConfigRoute) Handler() interface{} {
	return i.getSystemConfigHandler
}

func (i *getSystemConfigRoute) getSystemConfigHandler(w http.ResponseWriter, r *http.Request) {
	result := getSystemConfigResult{}
	result.SystemInfo = i.systemHandler.GetSystemConfig()
	result.ErrCode = common.Success

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
	common.Result
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

func (i *setSystemConfigRoute) setSystemConfigHandler(w http.ResponseWriter, r *http.Request) {
	result := setSystemConfigResult{}

	for true {
		r.ParseForm()

		systemInfo := model.SystemInfo{}
		systemInfo.Name = r.FormValue("system-name")
		systemInfo.Description = r.FormValue("system-description")
		systemInfo.Logo = r.FormValue("system-logo")
		systemInfo.Domain = r.FormValue("system-domain")
		systemInfo.MailServer = r.FormValue("system-mailsvr")
		systemInfo.MailAccount = r.FormValue("system-mailaccount")
		systemInfo.MailPassword = r.FormValue("system-mailpassword")

		if i.systemHandler.UpdateSystemConfig(systemInfo) {
			result.ErrCode = common.Success
		} else {
			result.ErrCode = common.Failed
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

type getModulesRoute struct {
	systemHandler common.SystemHandler
}

type getModulesResult struct {
	common.Result
	Modules []model.Module
}

func (i *getModulesRoute) Method() string {
	return common.GET
}

func (i *getModulesRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetSystemModule)
}

func (i *getModulesRoute) Handler() interface{} {
	return i.getModulesHandler
}

func (i *getModulesRoute) getModulesHandler(w http.ResponseWriter, r *http.Request) {
	result := getModulesResult{}

	result.Modules = i.systemHandler.GetModuleList()
	result.ErrCode = common.Success

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
