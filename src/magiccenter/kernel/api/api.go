package api

import (
	"magiccenter/common"
	commonhandler "magiccenter/common/handler"
	"magiccenter/kernel/api/ui"
	"magiccenter/system"

	"muidea.com/util"
)

// ID 模块ID
const ID = "5fa671dc-ccb5-4005-8500-f0e45b13705b"

// Name 模块名称
const Name = "Magic Dashboard API"

// Description 模块描述信息
const Description = "Magic Dashboard API模块"

// URL 模块Url
const URL string = "/api"

type api struct {
}

var instance *api

// LoadModule 加载模块
func LoadModule() {
	if instance == nil {
		instance = &api{}
	}

	modulehub := system.GetModuleHub()
	modulehub.RegisterModule(instance)
}

func (instance *api) ID() string {
	return ID
}

func (instance *api) Name() string {
	return Name
}

func (instance *api) Description() string {
	return Description
}

func (instance *api) Group() string {
	return "admin api"
}

func (instance *api) Type() int {
	return common.KERNEL
}

func (instance *api) URL() string {
	return URL
}

func (instance *api) Status() int {
	return 0
}

func (instance *api) EndPoint() common.EndPoint {
	return nil
}

func (instance *api) AuthGroups() []common.AuthGroup {
	groups := []common.AuthGroup{}

	return groups
}

// Route 路由信息
func (instance *api) Routes() []common.Route {
	router := system.GetRouter()
	auth := system.GetAuthority()

	routes := []common.Route{
		//=============================模块信息=====================================
		// 获取Module列表
		router.NewRoute(common.GET, "module/", ui.GetModuleListActionHandler, auth.AdminAuthVerify()),
		// 获取Module 定义的功能块
		router.NewRoute(common.GET, "module/block/", ui.GetModuleBlockActionHandler, auth.AdminAuthVerify()),
		// 获取Module 指定功能块包含的内容
		router.NewRoute(common.GET, "module/block/item/", ui.GetBlockItemActionHandler, auth.AdminAuthVerify()),
		// 获取Module 定义的授权分组
		router.NewRoute(common.GET, "module/authority/", ui.GetModuleAuthorityGroupActionHandler, auth.AdminAuthVerify()),

		//=============================内容信息=====================================
		// 获取内容元数据列表
		router.NewRoute(common.GET, "content/", ui.GetContentMetadataListActionHandler, auth.AdminAuthVerify()),
		// 获取文章信息
		router.NewRoute(common.GET, "content/article/", ui.GetContentArticleActionHandler, auth.AdminAuthVerify()),
		// 新增文章信息
		router.NewRoute(common.POST, "content/article/", ui.PostContentArticleActionHandler, auth.AdminAuthVerify()),
		// 更新文章信息
		router.NewRoute(common.PUT, "content/article/", ui.PutContentArticleActionHandler, auth.AdminAuthVerify()),
		// 删除文章
		router.NewRoute(common.DELETE, "content/article/", ui.DeleteContentArticleActionHandler, auth.AdminAuthVerify()),

		// 获取分类信息
		router.NewRoute(common.GET, "content/catalog/", ui.GetContentCatalogActionHandler, auth.AdminAuthVerify()),
		// 新增分类信息
		router.NewRoute(common.POST, "content/catalog/", ui.PostContentCatalogActionHandler, auth.AdminAuthVerify()),
		// 更新呢分类信息
		router.NewRoute(common.PUT, "content/catalog/", ui.PutContentCatalogActionHandler, auth.AdminAuthVerify()),
		// 删除分类信息
		router.NewRoute(common.DELETE, "content/catalog/", ui.DeleteContentCatalogActionHandler, auth.AdminAuthVerify()),

		// 获取链接
		router.NewRoute(common.GET, "content/link/", ui.GetContentLinkActionHandler, auth.AdminAuthVerify()),
		// 新增链接
		router.NewRoute(common.POST, "content/link/", ui.PostContentLinkActionHandler, auth.AdminAuthVerify()),
		// 更新链接
		router.NewRoute(common.PUT, "content/link/", ui.PutContentLinkActionHandler, auth.AdminAuthVerify()),
		// 删除链接
		router.NewRoute(common.DELETE, "content/link/", ui.DeleteContentLinkActionHandler, auth.AdminAuthVerify()),

		// 获取文件信息
		router.NewRoute(common.GET, "content/media/", ui.GetContentMediaActionHandler, auth.AdminAuthVerify()),
		// 新增文件信息
		router.NewRoute(common.POST, "content/media/", ui.PostContentMediaActionHandler, auth.AdminAuthVerify()),
		// 更新文件信息
		router.NewRoute(common.PUT, "content/media/", ui.PutContentMediaActionHandler, auth.AdminAuthVerify()),
		// 删除文件信息
		router.NewRoute(common.DELETE, "content/media/", ui.DeleteContentMediaActionHandler, auth.AdminAuthVerify()),

		//=============================账号信息=====================================
		// 获取User信息
		router.NewRoute(common.GET, "account/user/", ui.GetUserActionHandler, auth.AdminAuthVerify()),
		// 新建User
		router.NewRoute(common.POST, "account/user/", ui.PostUserActionHandler, auth.AdminAuthVerify()),
		// 更新User
		router.NewRoute(common.PUT, "account/user/", ui.PutUserActionHandler, auth.AdminAuthVerify()),
		// 删除User
		router.NewRoute(common.DELETE, "account/user/", ui.DeleteUserActionHandler, auth.AdminAuthVerify()),

		// 获取Group信息
		router.NewRoute(common.GET, "account/group/", ui.GetGroupActionHandler, auth.AdminAuthVerify()),
		// 新建Group
		router.NewRoute(common.POST, "account/group/", ui.PostGroupActionHandler, auth.AdminAuthVerify()),
		// 更新Group
		router.NewRoute(common.PUT, "account/group/", ui.PutGroupActionHandler, auth.AdminAuthVerify()),
		// 删除Group
		router.NewRoute(common.DELETE, "account/group/", ui.DeleteGroupActionHandler, auth.AdminAuthVerify()),

		//=============================系统信息=====================================
		// 获取系统信息
		router.NewRoute(common.GET, "system/", ui.GetSystemInfoActionHandler, auth.AdminAuthVerify()),
		// 更新系统信息
		router.NewRoute(common.PUT, "system/", ui.PutSystemInfoActionHandler, auth.AdminAuthVerify()),
		// 不支持Post和Delete
		router.NewRoute(common.POST, "system/", commonhandler.NoSupportActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.DELETE, "system/", commonhandler.NoSupportActionHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动模块
func (instance *api) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *api) Cleanup() {

}

// Invoke 执行外部命令
func (instance *api) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	return false
}