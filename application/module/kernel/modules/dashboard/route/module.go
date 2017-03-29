package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/dashboard/def"
	"muidea.com/magicCenter/foundation/net"
)

// AppendModuleRoute 追加Module 路由
func AppendModuleRoute(routes []common.Route, modHub common.ModuleHub) []common.Route {
	rt := CreateGetAllModule(modHub)

	routes = append(routes, rt)
	return routes
}

// CreateGetAllModule 新建GetAllModule Route
func CreateGetAllModule(modHub common.ModuleHub) common.Route {
	rt := dashBoardGetAllModuleRoute{moduleHub: modHub}

	return &rt
}

type dashBoardGetAllModuleRoute struct {
	moduleHub common.ModuleHub
}

type dashBoardGetAllModuleResult struct {
	common.Result
	Module []model.Module
}

func (i *dashBoardGetAllModuleRoute) Method() string {
	return common.GET
}

func (i *dashBoardGetAllModuleRoute) Pattern() string {
	return net.JoinURL(def.URL, "/module/")
}

func (i *dashBoardGetAllModuleRoute) Handler() interface{} {
	return i.getAllModuleHandler
}

func (i *dashBoardGetAllModuleRoute) getAllModuleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAllModuleHandler")

	result := dashBoardGetAllModuleResult{}
	for true {
		mod, _ := i.moduleHub.FindModule(common.CASModuleID)
		casHandler := mod.EndPoint().(common.CASHandler)

		modules := i.moduleHub.QueryAllModule()
		for _, v := range modules {
			module := model.Module{ID: v.ID(), Name: v.Name(), Description: v.Description(), Type: v.Type(), Status: v.Status()}

			routes := v.Routes()
			for _, rt := range routes {
				route := model.Route{Pattern: rt.Pattern(), Method: rt.Method()}
				module.Route = append(module.Route, route)
			}
			if module.Route == nil {
				module.Route = []model.Route{}
			}
			module.AuthGroup, _ = casHandler.QueryAuthGroup(v.ID())
			if module.AuthGroup == nil {
				module.AuthGroup = []model.AuthGroup{}
			}

			result.Module = append(result.Module, module)
		}
		result.ErrCode = 0
		break
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
