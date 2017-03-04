package cms

import (
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	"magiccenter/system"

	"muidea.com/util"
)

// ID CMS Module ID
const ID = "f17133ec-63e9-4b46-8758-e6ca1af6fe3f"

// Name CMS Module Name
const Name = "Magic CMS"

// Description CMS Module Description
const Description = "Magic 内容管理"

// URL CMS Module URL
const URL = "/cms"

// 授权分组属性Key，用于读取和存储授权分组信息
const authGroupKey = "f17133ec-63e9-4b46-8758-e6ca1af6fe3f_authGroupKey"

type cms struct {
	authGroup []common.AuthGroup
}

var instance *cms

// LoadModule 加载CMS模块
func LoadModule() {
	if instance == nil {
		instance = &cms{}
	}

	modulehub := system.GetModuleHub()
	modulehub.RegisterModule(instance)
}

func (c *cms) ID() string {
	return ID
}

func (c *cms) Name() string {
	return Name
}

func (c *cms) Description() string {
	return Description
}

func (c *cms) Group() string {
	return "content"
}

func (c *cms) Type() int {
	return common.INTERNAL
}

func (c *cms) URL() string {
	return URL
}

func (c *cms) Status() int {
	return 0
}

func (c *cms) EndPoint() common.EndPoint {
	return nil
}

func (c *cms) AuthGroups() []common.AuthGroup {
	return c.authGroup
}

func (c *cms) Routes() []common.Route {
	router := system.GetRouter()
	auth := system.GetAuthority()

	routes := []common.Route{
		router.NewRoute(common.GET, "/", indexHandler, nil),
		router.NewRoute(common.GET, "/view/", viewContentHandler, nil),
		router.NewRoute(common.GET, "/catalog/", viewCatalogHandler, nil),
		router.NewRoute(common.GET, "/link/", viewLinkHandler, nil),
		router.NewRoute(common.GET, "/maintain/", MaintainViewHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.POST, "/ajaxMaintain/", MaintainActionHandler, auth.AdminAuthVerify()),
	}

	return routes
}

func (c *cms) Startup() bool {
	configuration := system.GetConfiguration()
	value, found := configuration.GetOption(authGroupKey)
	if found {
		// fetch data from database
		ids, ok := util.Str2IntArray(value)
		if ok {
			groups, ok := commonbll.QueryGroups(ids)
			if ok {
				for _, g := range groups {
					c.authGroup = append(c.authGroup, common.CreateAuthGroup(g.Name, g.Description, g.Type, g.ID))
				}
			}
		}
	} else {
		ids := []int{}
		authorGroup, ok := commonbll.CreateGroup("作者组", "博客文章作者，可编写，更改文章")
		if ok {
			ids = append(ids, authorGroup.ID)
			c.authGroup = append(c.authGroup, common.CreateAuthGroup(authorGroup.Name, authorGroup.Description, authorGroup.Type, authorGroup.ID))
		}
		adminGroup, ok := commonbll.CreateGroup("管理组", "管理博客，维护博客基本信息，分类，文章")
		if ok {
			ids = append(ids, adminGroup.ID)
			c.authGroup = append(c.authGroup, common.CreateAuthGroup(adminGroup.Name, adminGroup.Description, adminGroup.Type, adminGroup.ID))
		}

		configuration.SetOption(authGroupKey, util.IntArray2Str(ids))
	}

	return true
}

func (c *cms) Cleanup() {

}

func (c *cms) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	return false
}
