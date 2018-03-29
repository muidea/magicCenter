package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/def"
	"muidea.com/magicCenter/foundation/net"
)

// CreateQueryModuleRoute 新建ModuleUserGetRoute
func CreateQueryModuleRoute(authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler, moduleHub common.ModuleHub) common.Route {
	i := moduleGetRoute{authorityHandler: authorityHandler, accountHandler: accountHandler, moduleHub: moduleHub}
	return &i
}

// CreateGetModuleByIDRoute 新建获取指定Module的用户授权组信息
func CreateGetModuleByIDRoute(authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler, moduleHub common.ModuleHub) common.Route {
	i := moduleGetByIDRoute{authorityHandler: authorityHandler, accountHandler: accountHandler, moduleHub: moduleHub}
	return &i
}

// CreatePutModuleRoute 新建PutModuleUserRoute
func CreatePutModuleRoute(authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler, moduleHub common.ModuleHub) common.Route {
	i := modulePutRoute{authorityHandler: authorityHandler, accountHandler: accountHandler, moduleHub: moduleHub}
	return &i
}

type moduleGetRoute struct {
	authorityHandler common.AuthorityHandler
	accountHandler   common.AccountHandler
	moduleHub        common.ModuleHub
}

type moduleGetResult struct {
	common.Result
	Module []model.ModuleUserInfoView `json:"module"`
}

func (i *moduleGetRoute) Method() string {
	return common.GET
}

func (i *moduleGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.QueryModule)
}

func (i *moduleGetRoute) Handler() interface{} {
	return i.getHandler
}

func (i *moduleGetRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *moduleGetRoute) getHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getHandler")

	result := moduleGetResult{}
	for true {
		moduleUserInfo := i.authorityHandler.QueryAllModuleUser()
		for _, val := range moduleUserInfo {
			view := model.ModuleUserInfoView{}

			mod, _ := i.moduleHub.FindModule(val.Module)
			view.Module.ID = mod.ID()
			view.Module.Name = mod.Name()

			view.User = i.accountHandler.GetUsers(val.User)

			result.Module = append(result.Module, view)
		}

		result.ErrorCode = common.Success

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type moduleGetByIDRoute struct {
	authorityHandler common.AuthorityHandler
	accountHandler   common.AccountHandler
	moduleHub        common.ModuleHub
}

type moduleGetByIDResult struct {
	common.Result
	Module model.ModuleUserAuthGroupView `json:"module"`
}

func (i *moduleGetByIDRoute) Method() string {
	return common.GET
}

func (i *moduleGetByIDRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetModuleByID)
}

func (i *moduleGetByIDRoute) Handler() interface{} {
	return i.getByIDHandler
}

func (i *moduleGetByIDRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *moduleGetByIDRoute) getByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getByIDHandler")

	result := moduleGetByIDResult{}
	for true {
		_, id := net.SplitRESTAPI(r.URL.Path)

		mod, _ := i.moduleHub.FindModule(id)
		result.Module.ID = mod.ID()
		result.Module.Name = mod.Name()
		result.Module.Description = mod.Description()
		result.Module.Type = mod.Type()
		result.Module.Status = mod.Status()

		userAuthGroups := i.authorityHandler.QueryModuleUserAuthGroup(id)
		for _, val := range userAuthGroups {
			user, ok := i.accountHandler.FindUserByID(val.User)

			if ok {
				view := model.UserAuthGroupView{}

				view.User.ID = user.ID
				view.User.Name = user.Name

				switch val.AuthGroup {
				case common.VisitorAuthGroup.ID:
					view.AuthGroup = common.VisitorAuthGroup
				case common.UserAuthGroup.ID:
					view.AuthGroup = common.UserAuthGroup
				case common.MaintainerAuthGroup.ID:
					view.AuthGroup = common.MaintainerAuthGroup
				default:
				}

				result.Module.UserAuthGroup = append(result.Module.UserAuthGroup, view)
			}
		}

		result.ErrorCode = common.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type modulePutRoute struct {
	authorityHandler common.AuthorityHandler
	accountHandler   common.AccountHandler
	moduleHub        common.ModuleHub
}

type modulePutParam struct {
	UserAuthGroup []model.UserAuthGroup `json:"userAuthGroup"`
}

type modulePutResult struct {
	common.Result
}

func (i *modulePutRoute) Method() string {
	return common.PUT
}

func (i *modulePutRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutModule)
}

func (i *modulePutRoute) Handler() interface{} {
	return i.putHandler
}

func (i *modulePutRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *modulePutRoute) putHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putHandler")

	result := modulePutResult{}
	for true {
		_, id := net.SplitRESTAPI(r.URL.Path)

		param := &modulePutParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		ok := i.authorityHandler.UpdateModuleUserAuthGroup(id, param.UserAuthGroup)
		if ok {
			result.ErrorCode = common.Success
		} else {
			result.ErrorCode = common.Failed
			result.Reason = "更新模块用户信息失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
