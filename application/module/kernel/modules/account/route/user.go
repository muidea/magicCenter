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
	"muidea.com/magicCenter/foundation/util"
)

// AppendUserRoute 追加User Route
func AppendUserRoute(routes []common.Route, accountHandler common.AccountHandler) []common.Route {

	rt := CreateGetUserRoute(accountHandler)
	routes = append(routes, rt)

	rt = CreateGetAllUserRoute(accountHandler)
	routes = append(routes, rt)

	rt = CreateCreateUserRoute(accountHandler)
	routes = append(routes, rt)

	rt = CreateSaveUserRoute(accountHandler)
	routes = append(routes, rt)

	rt = CreateDestroyUserRoute(accountHandler)
	routes = append(routes, rt)

	return routes
}

// CreateGetUserRoute 新建GetUserRoute
func CreateGetUserRoute(accountHandler common.AccountHandler) common.Route {
	i := userGetRoute{accountHandler: accountHandler}
	return &i
}

// CreateGetAllUserRoute 新建GetAllUser Route
func CreateGetAllUserRoute(accountHandler common.AccountHandler) common.Route {
	i := userGetAllRoute{accountHandler: accountHandler}
	return &i
}

// CreateCreateUserRoute 新建CreateUser Route
func CreateCreateUserRoute(accountHandler common.AccountHandler) common.Route {
	i := userCreateRoute{accountHandler: accountHandler}
	return &i
}

// CreateSaveUserRoute 新建SaveUser Route
func CreateSaveUserRoute(accountHandler common.AccountHandler) common.Route {
	i := userSaveRoute{accountHandler: accountHandler}
	return &i
}

// CreateDestroyUserRoute 新建DestroyUser Route
func CreateDestroyUserRoute(accountHandler common.AccountHandler) common.Route {
	i := userDestroyRoute{accountHandler: accountHandler}
	return &i
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
	return net.JoinURL(def.URL, def.GetUser)
}

func (i *userGetRoute) Handler() interface{} {
	return i.getUserHandler
}

func (i *userGetRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
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
	return net.JoinURL(def.URL, def.GetAllUser)
}

func (i *userGetAllRoute) Handler() interface{} {
	return i.getAllUserHandler
}

func (i *userGetAllRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
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
	return net.JoinURL(def.URL, def.PostUser)
}

func (i *userCreateRoute) Handler() interface{} {
	return i.createUserHandler
}

func (i *userCreateRoute) AuthGroup() int {
	return common.VisitorAuthGroup.ID
}

func (i *userCreateRoute) createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createUserHandler")

	result := userCreateResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		account := r.FormValue("user-account")
		email := r.FormValue("user-email")
		groups := r.FormValue("user-groups")

		gValues, ok := util.Str2IntArray(groups)
		if !ok {
			result.ErrCode = common.Failed
			result.Reason = "非法用户分组信息"
			break
		}

		user, ok := i.accountHandler.CreateUser(account, email, gValues)
		if !ok {
			result.ErrCode = common.Failed
			result.Reason = "创建新用户失败"
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
	return net.JoinURL(def.URL, def.PutUser)
}

func (i *userSaveRoute) Handler() interface{} {
	return i.saveUserHandler
}

func (i *userSaveRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
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

		user, ok := i.accountHandler.FindUserByID(id)
		if !ok {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		email := r.FormValue("user-email")
		if email != "" {
			user.Email = email
		}
		nickName := r.FormValue("user-name")
		if nickName != "" {
			user.Name = nickName
		}
		groups := r.FormValue("user-groups")
		if groups != "" {
			gValues, ok := util.Str2IntArray(groups)
			if !ok {
				result.ErrCode = common.Failed
				result.Reason = "非法用户分组信息"
				break
			}
			user.Groups = gValues
		}

		password := r.FormValue("user-password")
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
	return net.JoinURL(def.URL, def.DeleteUser)
}

func (i *userDestroyRoute) Handler() interface{} {
	return i.destroyUserHandler
}

func (i *userDestroyRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
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
