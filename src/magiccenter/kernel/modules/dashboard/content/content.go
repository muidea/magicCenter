package content

/*
Content 由三类信息组成
1、Block 功能块，提供基本的信息展示功能，比如导航栏，标签云，链接栏，列表栏，功能块展示有两种方式（1、展示链接、2、展示内容）
2、Item 内容项，功能块里展示的实际内容（目前有这样几种数据组成Article、Catalog、Link、Media），进行内容组织时统一使用Item
3、Page 页面，一个基本的信息展现单元，包含多个Block。每个Module拥有自己的Page页面，数量由各个Module自己决定，每个Page的内容都有多个Block组成

提供对Module内容进行查看，组织和管理的功能
1、提供Content管理视图
2、提供Page页面的Block的管理功能（增加、删除、查看）
3、提供Block内容的管理功能（增加、删除、管理）
*/

import (
	"magiccenter/common"
	commonhandler "magiccenter/common/handler"
	"magiccenter/kernel/modules/dashboard/content/ui"
	"magiccenter/system"
	"net/http"

	"muidea.com/util"
)

// ID 模块ID
const ID = "ffe1c37f-4fea-4a03-a7c3-55a331f5995f"

// Name 模块名称
const Name = "Magic ContentManage"

// Description 模块描述信息
const Description = "Magic 内容管理模块"

// URL 模块Url
const URL string = "/dashboard/content"

type contentmanage struct {
}

var instance *contentmanage

// LoadModule 加载ModuleManage模块
func LoadModule() {
	if instance == nil {
		instance = &contentmanage{}
	}

	modulehub := system.GetModuleHub()

	modulehub.RegisterModule(instance)
}

func (instance *contentmanage) ID() string {
	return ID
}

func (instance *contentmanage) Name() string {
	return Name
}

func (instance *contentmanage) Description() string {
	return Description
}

func (instance *contentmanage) Group() string {
	return "admin contentmanage"
}

func (instance *contentmanage) Type() int {
	return common.KERNEL
}

func (instance *contentmanage) URL() string {
	return URL
}

func (instance *contentmanage) EndPoint() common.EndPoint {
	return nil
}

// Route 路由信息
func (instance *contentmanage) Routes() []common.Route {
	router := system.GetRouter()
	auth := system.GetAuthority()

	routes := []common.Route{
		router.NewRoute(common.GET, "/", viewHandler, auth.AdminAuthVerify()),

		// Content Restfull API接口
		// 获取指定Module的视图信息
		router.NewRoute(common.GET, "module/", ui.ModuleViewHandler, auth.AdminAuthVerify()),

		router.NewRoute(common.POST, "module/", commonhandler.NoSupportActionHandler, auth.AdminAuthVerify()),

		router.NewRoute(common.DELETE, "module/", commonhandler.NoSupportActionHandler, auth.AdminAuthVerify()),

		router.NewRoute(common.PUT, "module/", commonhandler.NoSupportActionHandler, auth.AdminAuthVerify()),

		// 获取指定Page
		router.NewRoute(common.GET, "page/", ui.GetPageActionHandler, auth.AdminAuthVerify()),

		// 新建Page
		router.NewRoute(common.POST, "page/", ui.PostPageActionHandler, auth.AdminAuthVerify()),

		// 删除Page
		router.NewRoute(common.DELETE, "page/", ui.DeletePageActionHandler, auth.AdminAuthVerify()),

		// 更新Page
		router.NewRoute(common.PUT, "page/", commonhandler.NoSupportActionHandler, auth.AdminAuthVerify()),

		// 获取指定Block
		router.NewRoute(common.GET, "block/", ui.GetBlockActionHandler, auth.AdminAuthVerify()),

		// 新建Block
		router.NewRoute(common.POST, "block/", ui.PostBlockActionHandler, auth.AdminAuthVerify()),

		// 删除指定Block
		router.NewRoute(common.DELETE, "block/", ui.DeleteBlockActionHandler, auth.AdminAuthVerify()),

		// 更新指定Block
		router.NewRoute(common.PUT, "block/", ui.PutBlockActionHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动模块
func (instance *contentmanage) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *contentmanage) Cleanup() {

}

// Invoke 执行外部命令
func (instance *contentmanage) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	return false
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	commonhandler.HTMLViewHandler(w, r, "kernel/dashboard/content/content.html")
}
