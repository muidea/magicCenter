package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/account/def"
	common_def "muidea.com/magicCommon/common"
	common_result "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/foundation/net"
	"muidea.com/magicCommon/model"
)

// AppendGroupRoute 追加Group Route
func AppendGroupRoute(routes []common.Route, accountHandler common.AccountHandler) []common.Route {

	rt := CreateGetGroupRoute(accountHandler)
	routes = append(routes, rt)

	rt = CreateGetAllGroupDetailRoute(accountHandler)
	routes = append(routes, rt)

	rt = CreateCreateGroupRoute(accountHandler)
	routes = append(routes, rt)

	rt = CreateSaveGroupRoute(accountHandler)
	routes = append(routes, rt)

	rt = CreateDestroyGroupRoute(accountHandler)
	routes = append(routes, rt)

	return routes
}

// CreateGetGroupRoute 新建GetGroupRoute
func CreateGetGroupRoute(accountHandler common.AccountHandler) common.Route {
	i := groupGetRoute{accountHandler: accountHandler}
	return &i
}

// CreateGetAllGroupDetailRoute 新建GetGroupsDetail Route
func CreateGetAllGroupDetailRoute(accountHandler common.AccountHandler) common.Route {
	i := groupGetAllRoute{accountHandler: accountHandler}
	return &i
}

// CreateCreateGroupRoute 新建CreateGroup Route
func CreateCreateGroupRoute(accountHandler common.AccountHandler) common.Route {
	i := groupCreateRoute{accountHandler: accountHandler}
	return &i
}

// CreateSaveGroupRoute 新建SaveGroup Route
func CreateSaveGroupRoute(accountHandler common.AccountHandler) common.Route {
	i := groupSaveRoute{accountHandler: accountHandler}
	return &i
}

// CreateDestroyGroupRoute 新建DestroyGroup Route
func CreateDestroyGroupRoute(accountHandler common.AccountHandler) common.Route {
	i := groupDestroyRoute{accountHandler: accountHandler}
	return &i
}

type groupGetRoute struct {
	accountHandler common.AccountHandler
}

type groupGetResult struct {
	common_result.Result
	Group model.GroupDetailView `json:"group"`
}

func (i *groupGetRoute) Method() string {
	return common.GET
}

func (i *groupGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetGroup)
}

func (i *groupGetRoute) Handler() interface{} {
	return i.getGroupHandler
}

func (i *groupGetRoute) AuthGroup() int {
	return common_def.UserAuthGroup.ID
}

func (i *groupGetRoute) getGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getGroupHandler")

	result := groupGetResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		group, ok := i.accountHandler.FindGroupByID(id)
		if ok {
			result.Group.GroupDetail = group

			catalog, _ := i.accountHandler.FindGroupByID(group.Catalog)
			result.Group.Catalog.ID = catalog.ID
			result.Group.Catalog.Name = catalog.Name

			result.ErrorCode = common_result.Success
		} else {
			result.ErrorCode = common_result.Failed
			result.Reason = "对象不存在"
		}
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type groupGetAllRoute struct {
	accountHandler common.AccountHandler
}

type groupGetAllResult struct {
	common_result.Result
	Group []model.GroupDetailView `json:"group"`
}

func (i *groupGetAllRoute) Method() string {
	return common.GET
}

func (i *groupGetAllRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetAllGroup)
}

func (i *groupGetAllRoute) Handler() interface{} {
	return i.getAllGroupHandler
}

func (i *groupGetAllRoute) AuthGroup() int {
	return common_def.UserAuthGroup.ID
}

func (i *groupGetAllRoute) getAllGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAllGroupHandler")

	result := groupGetAllResult{}
	for true {
		filter := &common_def.PageFilter{}
		filter.Parse(r)

		allGroups := i.accountHandler.GetAllGroupDetail(filter)
		for _, val := range allGroups {
			groupView := model.GroupDetailView{}
			groupView.GroupDetail = val
			catalog, _ := i.accountHandler.FindGroupByID(val.Catalog)
			groupView.Catalog.ID = catalog.ID
			groupView.Catalog.Name = catalog.Name

			result.Group = append(result.Group, groupView)
		}
		result.ErrorCode = common_result.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type groupCreateRoute struct {
	accountHandler common.AccountHandler
}

type groupCreateParam struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Catalog     model.Group `json:"catalog"`
}

type groupCreateResult struct {
	common_result.Result
	Group model.GroupDetailView `json:"group"`
}

func (i *groupCreateRoute) Method() string {
	return common.POST
}

func (i *groupCreateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostGroup)
}

func (i *groupCreateRoute) Handler() interface{} {
	return i.createGroupHandler
}

func (i *groupCreateRoute) AuthGroup() int {
	return common_def.MaintainerAuthGroup.ID
}

func (i *groupCreateRoute) createGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createGroupHandler")

	result := groupCreateResult{}
	for true {
		param := &groupCreateParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "非法参数"
			break
		}

		group, ok := i.accountHandler.CreateGroup(param.Name, param.Description, param.Catalog.ID)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		result.Group.GroupDetail = group
		catalog, _ := i.accountHandler.FindGroupByID(param.Catalog.ID)
		result.Group.Catalog.ID = catalog.ID
		result.Group.Catalog.Name = catalog.Name
		result.ErrorCode = common_result.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type groupSaveRoute struct {
	accountHandler common.AccountHandler
}

type groupSaveParam struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Catalog     model.Group `json:"catalog"`
}

type groupSaveResult struct {
	common_result.Result
	Group model.GroupDetailView `json:"group"`
}

func (i *groupSaveRoute) Method() string {
	return common.PUT
}

func (i *groupSaveRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutGroup)
}

func (i *groupSaveRoute) Handler() interface{} {
	return i.saveGroupHandler
}

func (i *groupSaveRoute) AuthGroup() int {
	return common_def.MaintainerAuthGroup.ID
}

func (i *groupSaveRoute) saveGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("saveGroupHandler")

	result := groupCreateResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		param := &groupSaveParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "非法参数"
			break
		}

		group := model.NewGroup(param.Name, param.Description, param.Catalog.ID)
		group.ID = id

		group, ok := i.accountHandler.SaveGroup(group)

		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "更新失败"
			break
		}

		result.Group.GroupDetail = group
		catalog, _ := i.accountHandler.FindGroupByID(param.Catalog.ID)
		result.Group.Catalog.ID = catalog.ID
		result.Group.Catalog.Name = catalog.Name
		result.ErrorCode = common_result.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type groupDestroyRoute struct {
	accountHandler common.AccountHandler
}

type groupDestroyResult struct {
	common_result.Result
}

func (i *groupDestroyRoute) Method() string {
	return common.DELETE
}

func (i *groupDestroyRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteGroup)
}

func (i *groupDestroyRoute) Handler() interface{} {
	return i.destroyGroupHandler
}

func (i *groupDestroyRoute) AuthGroup() int {
	return common_def.MaintainerAuthGroup.ID
}

func (i *groupDestroyRoute) destroyGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("destroyGroupHandler")

	result := groupDestroyResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		ok := i.accountHandler.DestroyGroup(id)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "删除失败"
			break
		}

		result.ErrorCode = common_result.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
