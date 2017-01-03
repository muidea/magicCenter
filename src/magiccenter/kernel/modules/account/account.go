package account

import (
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	commonmodel "magiccenter/common/model"
	"magiccenter/kernel/modules/account/bll"
	"magiccenter/kernel/modules/account/ui"
	"magiccenter/system"

	"muidea.com/util"
)

// ID 模块ID
const ID = "b9e35167-b2a3-43ae-8c57-9b4379475e47"

// Name 模块名称
const Name = "Magic Account"

// Description 模块描述信息
const Description = "Magic 账号管理模块"

// URL 模块Url
const URL string = "/account"

type account struct {
}

var instance *account

// LoadModule 加载模块
func LoadModule() {
	if instance == nil {
		instance = &account{}
	}

	modulehub := system.GetModuleHub()
	modulehub.RegisterModule(instance)
}

func (instance *account) ID() string {
	return ID
}

func (instance *account) Name() string {
	return Name
}

func (instance *account) Description() string {
	return Description
}

func (instance *account) Group() string {
	return "kernel"
}

func (instance *account) Type() int {
	return common.KERNEL
}

func (instance *account) URL() string {
	return URL
}

func (instance *account) Status() int {
	return 0
}

func (instance *account) EndPoint() common.EndPoint {
	return nil
}

func (instance *account) AuthGroups() []common.AuthGroup {
	groups := []common.AuthGroup{}

	return groups
}

// Route Account 路由信息
func (instance *account) Routes() []common.Route {
	router := system.GetRouter()
	auth := system.GetAuthority()

	routes := []common.Route{
		// 用户账号信息管理视图
		router.NewRoute(common.GET, "manageUserView/", ui.ManageUserViewHandler, auth.AdminAuthVerify()),

		// 用户分组信息管理视图
		router.NewRoute(common.GET, "manageGroupView/", ui.ManageGroupViewHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动Account模块
func (instance *account) Startup() bool {
	return true
}

// Cleanup 清除Account模块
func (instance *account) Cleanup() {

}

// Invoke 执行外部命令
func (instance *account) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	switch param.(type) {
	case *commonbll.QueryAccountMetaRequest:
		{
			request := param.(*commonbll.QueryAccountMetaRequest)
			if request != nil {
				response := result.(*commonbll.QueryAccountMetaResponse)
				userMeta := commonmodel.AccountMeta{Subject: commonmodel.USER, Description: "账号", URL: "account/user/"}
				response.AccountMetas = append(response.AccountMetas, userMeta)

				groupMeta := commonmodel.AccountMeta{Subject: commonmodel.GROUP, Description: "分组", URL: "account/group/"}
				response.AccountMetas = append(response.AccountMetas, groupMeta)
				return true
			}
			return false
		}
	case *commonbll.QueryUserDetailRequest:
		{
			request := param.(*commonbll.QueryUserDetailRequest)
			response := result.(*commonbll.QueryUserDetailResponse)
			if request != nil && response != nil {
				response.User, response.Result = bll.QueryUserByID(request.ID)
				return true
			}

			return false
		}
	case *commonbll.QueryAllUserRequest:
		{
			request := param.(*commonbll.QueryAllUserRequest)
			response := result.(*commonbll.QueryAllUserResponse)
			if request != nil && response != nil {
				response.Users = bll.QueryAllUser()
				return true
			}

			return false
		}
	case *commonbll.CreateUserRequest:
		{
			request := param.(*commonbll.CreateUserRequest)
			response := result.(*commonbll.CreateUserResponse)
			if request != nil && response != nil {
				response.User, response.Result = bll.CreateUser(request.Account, request.EMail)
				return true
			}

			return false
		}
	case *commonbll.UpdateUserRequest:
		{
			request := param.(*commonbll.UpdateUserRequest)
			response := result.(*commonbll.UpdateUserResponse)
			if request != nil && response != nil {
				response.User, response.Result = bll.UpdateUser(request.User)
				return true
			}

			return false
		}
	case *commonbll.DeleteUserRequest:
		{
			request := param.(*commonbll.DeleteUserRequest)
			response := result.(*commonbll.DeleteUserResponse)
			if request != nil && response != nil {
				response.Result = bll.DeleteUser(request.ID)
				return true
			}

			return false
		}
	case *commonbll.QueryAllGroupRequest:
		{
			request := param.(*commonbll.QueryAllGroupRequest)
			response := result.(*commonbll.QueryAllGroupResponse)
			if request != nil && response != nil {
				response.Groups = bll.QueryAllGroup()
				return true
			}

			return false
		}
	case *commonbll.QueryGroupsRequest:
		{
			request := param.(*commonbll.QueryGroupsRequest)
			response := result.(*commonbll.QueryGroupsResponse)
			if request != nil && response != nil {
				response.Groups = bll.QueryGroups(request.Ids)
				return true
			}

			return false
		}
	case *commonbll.CreateGroupRequest:
		{
			request := param.(*commonbll.CreateGroupRequest)
			response := result.(*commonbll.CreateGroupResponse)
			if request != nil && response != nil {
				response.Group, response.Result = bll.CreateGroup(request.Name, request.Description)
				return true
			}

			return false
		}
	case *commonbll.UpdateGroupRequest:
		{
			request := param.(*commonbll.UpdateGroupRequest)
			response := result.(*commonbll.UpdateGroupResponse)
			if request != nil && response != nil {
				response.Group, response.Result = bll.SaveGroup(request.Group)
				return true
			}

			return false
		}
	case *commonbll.DeleteGroupRequest:
		{
			request := param.(*commonbll.DeleteGroupRequest)
			response := result.(*commonbll.DeleteGroupResponse)
			if request != nil && response != nil {
				response.Result = bll.DeleteGroup(request.ID)
				return true
			}

			return false
		}
	}

	return false
}
