package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/account/def"
	"muidea.com/magicCommon/foundation/net"
	common_const "muidea.com/magicCommon/common"
	common_result "muidea.com/magicCommon/common"
	common_model "muidea.com/magicCommon/model"
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
	common_result.Result
	User common_model.UserDetailView `json:"user"`
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
	return common_const.UserAuthGroup.ID
}

func (i *userGetRoute) getUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getUserHandler")

	result := userGetResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		user, ok := i.accountHandler.FindUserByID(id)
		if ok {
			result.User.UserDetail = user
			result.User.Group = i.accountHandler.GetGroups(user.Group)
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

type userGetAllRoute struct {
	accountHandler common.AccountHandler
}

type userGetAllResult struct {
	common_result.Result
	User []common_model.UserDetailView `json:"user"`
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
	return common_const.UserAuthGroup.ID
}

func (i *userGetAllRoute) getAllUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAllUserHandler")

	result := userGetAllResult{}
	for true {
		allUsers := i.accountHandler.GetAllUserDetail()
		for _, val := range allUsers {
			user := common_model.UserDetailView{}
			user.UserDetail = val
			user.Group = i.accountHandler.GetGroups(val.Group)

			result.User = append(result.User, user)
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

type userCreateRoute struct {
	accountHandler common.AccountHandler
}

type userCreateParam struct {
	Account  string               `json:"account"`
	Password string               `json:"password"`
	EMail    string               `json:"email"`
	Group    []common_model.Group `json:"group"`
}

type userCreateResult struct {
	common_result.Result
	User common_model.UserDetailView `json:"user"`
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
	return common_const.VisitorAuthGroup.ID
}

func (i *userCreateRoute) createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createUserHandler")

	result := userCreateResult{}
	for true {
		param := &userCreateParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "非法参数"
			break
		}
		ids := []int{}
		for _, val := range param.Group {
			ids = append(ids, val.ID)
		}

		user, ok := i.accountHandler.CreateUser(param.Account, param.Password, param.EMail, ids)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "创建新用户失败"
			break
		}

		result.User.UserDetail = user
		result.User.Group = i.accountHandler.GetGroups(user.Group)
		result.User.Status = common_const.GetStatus(user.Status)
		result.ErrorCode = common_result.Success
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

type userSaveParam struct {
	Email string               `json:"email"`
	Group []common_model.Group `json:"group"`
}

type userSavePasswordParam struct {
	Password string `json:"password"`
}

type userSaveResult userCreateResult

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
	return common_const.UserAuthGroup.ID
}

func (i *userSaveRoute) saveUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("saveUserHandler")

	result := userCreateResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	action := r.URL.Query().Get("action")
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		user, ok := i.accountHandler.FindUserByID(id)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "非法参数"
			break
		}

		if action == "change_password" {
			param := &userSavePasswordParam{}
			err = net.ParsePostJSON(r, param)
			if err != nil {
				result.ErrorCode = common_result.Failed
				result.Reason = "非法参数"
				break
			}
			if param.Password == "" {
				result.ErrorCode = common_result.Failed
				result.Reason = "非法参数"
				break
			}

			user, ok = i.accountHandler.SaveUserWithPassword(user, param.Password)
		} else {
			param := &userSaveParam{}
			err = net.ParsePostJSON(r, param)
			if err != nil {
				result.ErrorCode = common_result.Failed
				result.Reason = "非法参数"
				break
			}

			if param.Email != "" {
				user.Email = param.Email
			}

			if len(param.Group) > 0 {
				ids := []int{}
				for _, v := range param.Group {
					ids = append(ids, v.ID)
				}
				user.Group = ids
			}

			user, ok = i.accountHandler.SaveUser(user)
		}

		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "更新失败"
			break
		}

		result.User.UserDetail = user
		result.User.Group = i.accountHandler.GetGroups(user.Group)
		result.User.Status = common_const.GetStatus(user.Status)
		result.ErrorCode = common_result.Success
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
	common_result.Result
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
	return common_const.MaintainerAuthGroup.ID
}

func (i *userDestroyRoute) destroyUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("destroyUserHandler")

	result := userDestroyResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		ok := i.accountHandler.DestroyUserByID(id)
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
