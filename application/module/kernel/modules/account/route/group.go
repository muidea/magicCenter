package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/account/def"
	"muidea.com/magicCenter/foundation/net"
)

// AppendGroupRoute 追加Group Route
func AppendGroupRoute(routes []common.Route, accountHandler common.AccountHandler) []common.Route {

	rt := CreateGetGroupRoute(accountHandler)
	routes = append(routes, rt)

	rt = CreateGetAllGroupRoute(accountHandler)
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

// CreateGetAllGroupRoute 新建GetAllGroup Route
func CreateGetAllGroupRoute(accountHandler common.AccountHandler) common.Route {
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
	common.Result
	Group model.Group
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
	return common.UserAuthGroup.ID
}

func (i *groupGetRoute) getGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getGroupHandler")

	result := groupGetResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		group, ok := i.accountHandler.FindGroupByID(id)
		if ok {
			result.Group = group
			result.ErrCode = 0
		} else {
			result.ErrCode = 1
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
	common.Result
	Group []model.Group
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
	return common.UserAuthGroup.ID
}

func (i *groupGetAllRoute) getAllGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAllGroupHandler")

	result := groupGetAllResult{}
	for true {
		result.Group = i.accountHandler.GetAllGroup()
		result.ErrCode = 0
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
	Name        string
	Description string
}

type groupCreateResult struct {
	common.Result
	Group model.Group
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
	return common.MaintainerAuthGroup.ID
}

func (i *groupCreateRoute) createGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createGroupHandler")

	result := groupCreateResult{}
	for true {
		param := &groupCreateParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		group, ok := i.accountHandler.CreateGroup(param.Name, param.Description)
		if !ok {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		result.Group = group
		result.ErrCode = 0
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
	Name        string
	Description string
}

type groupSaveResult struct {
	common.Result
	Group model.Group
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
	return common.MaintainerAuthGroup.ID
}

func (i *groupSaveRoute) saveGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("saveGroupHandler")

	result := groupCreateResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		param := &groupSaveParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		group := model.Group{ID: id, Name: param.Name, Description: param.Description, Catalog: 0}
		group, ok := i.accountHandler.SaveGroup(group)

		if !ok {
			result.ErrCode = 1
			result.Reason = "更新失败"
			break
		}

		result.Group = group
		result.ErrCode = 0
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
	common.Result
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
	return common.MaintainerAuthGroup.ID
}

func (i *groupDestroyRoute) destroyGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("destroyGroupHandler")

	result := groupDestroyResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		ok := i.accountHandler.DestroyGroup(id)
		if !ok {
			result.ErrCode = 1
			result.Reason = "删除失败"
			break
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
