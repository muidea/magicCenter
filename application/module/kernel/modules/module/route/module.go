package route

import (
	"encoding/json"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/module/def"
	"muidea.com/magicCenter/foundation/net"
	common_const "muidea.com/magicCommon/common"
	common_result "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/model"
)

// AppendModuleRegistryRoute 追加SystemRoute
func AppendModuleRegistryRoute(routes []common.Route, moduleHandler common.ModuleRegistryHandler) []common.Route {
	rt := GetModulesRoute(moduleHandler)
	routes = append(routes, rt)

	return routes
}

// GetModulesRoute 新建获取Modules路由
func GetModulesRoute(moduleHandler common.ModuleRegistryHandler) common.Route {
	return &getModulesRoute{moduleHandler: moduleHandler}
}

type getModulesRoute struct {
	moduleHandler common.ModuleRegistryHandler
}

type getModulesResult struct {
	common_result.Result
	Module []model.ModuleDetailView `json:"module"`
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
	result := getModulesResult{}

	modules := i.moduleHandler.GetModuleDetailList()
	for _, v := range modules {
		detail := model.ModuleDetailView{}
		detail.ModuleDetail = v
		detail.Type = common_const.GetModuleType(v.Type)
		detail.Status = common_const.GetStatus(v.Status)

		result.Module = append(result.Module, detail)
	}
	result.ErrorCode = common_result.Success

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
