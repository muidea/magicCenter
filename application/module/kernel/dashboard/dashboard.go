package dashboard

import (
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	commonhandler "magiccenter/common/handler"
	"magiccenter/kernel/dashboard/ui"
	"magiccenter/system"
	"net/http"

	"muidea.com/util"
)

// ID 模块ID
const ID = "f67123ea-6fe0-5e46-1234-e6ca1af6fe4e"

// Name 模块名称
const Name = "Magic Dashboard"

// Description 模块描述信息
const Description = "Magic Dashboard模块"

// URL 模块Url
const URL = "/dashboard"

// 授权分组属性Key，用于读取和存储授权分组信息
const authGroupKey = "f67123ea-6fe0-5e46-1234-e6ca1af6fe4e_authGroupKey"

type dashboard struct {
	authGroup []common.AuthGroup
}

var instance *dashboard

// LoadModule 加载模块
func LoadModule() {
	if instance == nil {
		instance = &dashboard{}
	}

	modulehub := system.GetModuleHub()
	modulehub.RegisterModule(instance)
}

func (instance *dashboard) ID() string {
	return ID
}

func (instance *dashboard) Name() string {
	return Name
}

func (instance *dashboard) Description() string {
	return Description
}

func (instance *dashboard) Group() string {
	return "admin dashboard"
}

func (instance *dashboard) Type() int {
	return common.KERNEL
}

func (instance *dashboard) URL() string {
	return URL
}

func (instance *dashboard) Status() int {
	return 0
}

func (instance *dashboard) EndPoint() common.EndPoint {
	return nil
}

func (instance *dashboard) AuthGroups() []common.AuthGroup {
	return instance.authGroup
}

// Route 路由信息
func (instance *dashboard) Routes() []common.Route {
	router := system.GetRouter()
	auth := system.GetAuthority()

	routes := []common.Route{
		// Dashboard主页面
		router.NewRoute(common.GET, "/", ui.DashboardViewHandler, auth.AdminAuthVerify()),
		// 登陆页面
		router.NewRoute(common.GET, "login/", ui.LoginViewHandler, nil),
		// 登陆校验
		router.NewRoute(common.POST, "verify/", ui.VerifyAuthActionHandler, nil),
		// 登出校验
		router.NewRoute(common.GET, "logout/", ui.LogoutActionHandler, auth.AdminAuthVerify()),

		router.NewRoute(common.GET, "module/", moduleViewHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.GET, "content/", contentViewHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.GET, "authority/", authorityViewHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.GET, "system/", systemViewHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动模块
func (instance *dashboard) Startup() bool {
	configuration := system.GetConfiguration()
	value, found := configuration.GetOption(authGroupKey)
	if found {
		// fetch data from database
		ids, ok := util.Str2IntArray(value)
		if ok {
			groups, ok := commonbll.QueryGroups(ids)
			if ok {
				for _, g := range groups {
					instance.authGroup = append(instance.authGroup, common.CreateAuthGroup(g.Name, g.Description, g.Type, g.ID))
				}
			}
		}
	} else {
		ids := []int{}
		adminGroup, ok := commonbll.CreateGroup("管理员组", "管理员组，拥有最大的管理权限")
		adminGroup.Type = 1
		adminGroup, ok = commonbll.UpdateGroup(adminGroup)
		if ok {
			ids = append(ids, adminGroup.ID)
			instance.authGroup = append(instance.authGroup, common.CreateAuthGroup(adminGroup.Name, adminGroup.Description, adminGroup.Type, adminGroup.ID))
		}

		configuration.SetOption(authGroupKey, util.IntArray2Str(ids))
	}

	return true
}

// Cleanup 清除模块
func (instance *dashboard) Cleanup() {

}

// Invoke 执行外部命令
func (instance *dashboard) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	return false
}

func moduleViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("dashboard module viewHandler")
	commonhandler.HTMLViewHandler(w, r, "kernel/dashboard/module.html")
}

func contentViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("dashboard content viewHandler")
	commonhandler.HTMLViewHandler(w, r, "kernel/dashboard/content.html")
}

func authorityViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("dashboard authority viewHandler")
	commonhandler.HTMLViewHandler(w, r, "kernel/dashboard/authority.html")
}

func systemViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("dashboard system viewHandler")
	commonhandler.HTMLViewHandler(w, r, "kernel/dashboard/system.html")
}
