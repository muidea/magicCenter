package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/api/def"
	"muidea.com/magicCenter/foundation/net"
)

// AppendUserRoute 追加User Route
func AppendUserRoute(routes []common.Route, modHub common.ModuleHub) []common.Route {

	rt, _ := CreateGetUserRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateGetAllUserRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateCreateUserRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateSaveUserRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateDestroyUserRoute(modHub)
	routes = append(routes, rt)

	return routes
}

// CreateGetUserRoute 新建GetUserRoute
func CreateGetUserRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.AccountModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.AccountHandler:
		i := userGetRoute{accountHandler: endPoint.(common.AccountHandler)}
		return &i, true
	}

	return nil, false
}

// CreateGetAllUserRoute 新建GetAllUser Route
func CreateGetAllUserRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.AccountModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.AccountHandler:
		i := userGetAllRoute{accountHandler: endPoint.(common.AccountHandler)}
		return &i, true
	}

	return nil, false
}

// CreateCreateUserRoute 新建CreateUser Route
func CreateCreateUserRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.AccountModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.AccountHandler:
		i := userCreateRoute{accountHandler: endPoint.(common.AccountHandler)}
		return &i, true
	}

	return nil, false
}

// CreateSaveUserRoute 新建SaveUser Route
func CreateSaveUserRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.AccountModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.AccountHandler:
		i := userSaveRoute{accountHandler: endPoint.(common.AccountHandler)}
		return &i, true
	}

	return nil, false
}

// CreateDestroyUserRoute 新建DestroyUser Route
func CreateDestroyUserRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.AccountModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.AccountHandler:
		i := userDestroyRoute{accountHandler: endPoint.(common.AccountHandler)}
		return &i, true
	}

	return nil, false
}

type userGetRoute struct {
	accountHandler common.AccountHandler
}

type userGetResult struct {
	common.Result
	User model.UserDetail
}

func (i *userGetRoute) Method() string {
	return common.GET
}

func (i *userGetRoute) Pattern() string {
	return net.JoinURL(def.URL, "account/user/[0-9]+/")
}

func (i *userGetRoute) Handler() interface{} {
	return i.getUserHandler
}

func (i *userGetRoute) getUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getUserHandler")

	result := userGetResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		user, ok := i.accountHandler.FindUserByID(id)
		if ok {
			result.User = user
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

type userGetAllRoute struct {
	accountHandler common.AccountHandler
}

type userGetAllResult struct {
	common.Result
	User []model.UserDetail
}

func (i *userGetAllRoute) Method() string {
	return common.GET
}

func (i *userGetAllRoute) Pattern() string {
	return net.JoinURL(def.URL, "account/user/")
}

func (i *userGetAllRoute) Handler() interface{} {
	return i.getAllUserHandler
}

func (i *userGetAllRoute) getAllUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAllUserHandler")

	result := userGetAllResult{}
	for true {
		result.User = i.accountHandler.GetAllUser()
		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type userCreateRoute struct {
	accountHandler common.AccountHandler
}

type userCreateResult struct {
	common.Result
	User model.UserDetail
}

func (i *userCreateRoute) Method() string {
	return common.POST
}

func (i *userCreateRoute) Pattern() string {
	return net.JoinURL(def.URL, "account/user/")
}

func (i *userCreateRoute) Handler() interface{} {
	return i.createUserHandler
}

func (i *userCreateRoute) createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createUserHandler")

	result := userCreateResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		account := r.FormValue("user-account")
		email := r.FormValue("user-email")
		user, ok := i.accountHandler.CreateUser(account, email)
		if !ok {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		result.User = user
		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type userSaveRoute struct {
	accountHandler common.AccountHandler
}

type userSaveResult struct {
	common.Result
	User model.UserDetail
}

func (i *userSaveRoute) Method() string {
	return common.PUT
}

func (i *userSaveRoute) Pattern() string {
	return net.JoinURL(def.URL, "account/user/[0-9]+/")
}

func (i *userSaveRoute) Handler() interface{} {
	return i.saveUserHandler
}

func (i *userSaveRoute) saveUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("saveUserHandler")

	result := userCreateResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		err = r.ParseForm()
		if err != nil {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		account := r.FormValue("user-account")
		email := r.FormValue("user-email")
		nickName := r.FormValue("user-name")
		password := r.FormValue("user-password")
		user := model.UserDetail{Account: account, Email: email}
		user.ID = id
		user.Name = nickName
		ok := false
		if len(password) > 0 {
			user, ok = i.accountHandler.SaveUserWithPassword(user, password)
		} else {
			user, ok = i.accountHandler.SaveUser(user)
		}

		if !ok {
			result.ErrCode = 1
			result.Reason = "更新失败"
			break
		}

		result.User = user
		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type userDestroyRoute struct {
	accountHandler common.AccountHandler
}

type userDestroyResult struct {
	common.Result
}

func (i *userDestroyRoute) Method() string {
	return common.DELETE
}

func (i *userDestroyRoute) Pattern() string {
	return net.JoinURL(def.URL, "account/user/[0-9]+/")
}

func (i *userDestroyRoute) Handler() interface{} {
	return i.destroyUserHandler
}

func (i *userDestroyRoute) destroyUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("destroyUserHandler")

	result := userDestroyResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		ok := i.accountHandler.DestroyUserByID(id)
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
