package route

import (
	"encoding/json"
	"net/http"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/module/def"
	common_const "muidea.com/magicCommon/common"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/foundation/net"
	"muidea.com/magicCommon/model"
)

// AppendModuleRegistryRoute 追加SystemRoute
func AppendModuleRegistryRoute(routes []common.Route, moduleHandler common.ModuleRegistryHandler) []common.Route {
	rt := GetModulesRoute(moduleHandler)
	routes = append(routes, rt)

	rt = GetModuleByIDRoute(moduleHandler)
	routes = append(routes, rt)

	return routes
}

// GetModulesRoute 新建获取Modules路由
func GetModulesRoute(moduleHandler common.ModuleRegistryHandler) common.Route {
	return &getModulesRoute{moduleHandler: moduleHandler}
}

// GetModuleByIDRoute 获取指定Module
func GetModuleByIDRoute(moduleHandler common.ModuleRegistryHandler) common.Route {
	return &getModuleByIDRoute{moduleHandler: moduleHandler}
}

type getModulesRoute struct {
	moduleHandler common.ModuleRegistryHandler
}

func (i *getModulesRoute) Method() string {
	return common.GET
}

func (i *getModulesRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetListModule)
}

func (i *getModulesRoute) Handler() interface{} {
	return i.getModulesHandler
}

func (i *getModulesRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *getModulesRoute) getModulesHandler(w http.ResponseWriter, r *http.Request) {
	result := common_def.GetModuleListResult{}

	modules := i.moduleHandler.GetModuleDetailList()
	for _, v := range modules {
		detail := model.ModuleDetailView{}
		detail.ModuleDetail = v
		detail.Type = common_const.GetModuleType(v.Type)
		detail.Status = common_const.GetStatus(v.Status)

		result.Module = append(result.Module, detail)
	}
	result.ErrorCode = common_def.Success

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type getModuleByIDRoute struct {
	moduleHandler common.ModuleRegistryHandler
}

type getModuleByIDResult struct {
	common_def.Result
	Module model.ModuleDetailView `json:"module"`
}

func (i *getModuleByIDRoute) Method() string {
	return common.GET
}

func (i *getModuleByIDRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetModuleByID)
}

func (i *getModuleByIDRoute) Handler() interface{} {
	return i.getModuleByIDHandler
}

func (i *getModuleByIDRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *getModuleByIDRoute) getModuleByIDHandler(w http.ResponseWriter, r *http.Request) {
	result := getModuleByIDResult{}

	for {
		_, id := net.SplitRESTAPI(r.URL.Path)
		detail, ok := i.moduleHandler.QueryModuleByID(id)
		if ok {
			result.Module.ModuleDetail = detail
			result.Module.Type = common_const.GetModuleType(detail.Type)
			result.Module.Status = common_const.GetStatus(detail.Status)
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.NoExist
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
