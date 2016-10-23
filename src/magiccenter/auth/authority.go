package auth

import (
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	commonmodel "magiccenter/common/model"
	"magiccenter/system"
	"martini"
	"net/http"
)

// AuthorithID 登陆会话鉴权ID
const AuthorithID = "@authorith_Id"

type impl struct {
}

// CreateAuthority 创建Authority
func CreateAuthority() common.Authority {
	impl := impl{}

	return &impl
}

// AdminAuthVerify 管理权限校验器
func (intance impl) AdminAuthVerify() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request) bool {
		result := false
		session := system.GetSession(res, req)
		user, found := session.GetOption(AuthorithID)
		if found {
			result = commonbll.IsAdministrator(user.(commonmodel.UserDetail))
		}

		return result
	}
}

// LoginAuthVerify 登陆权限校验器
func (intance impl) LoginAuthVerify() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request) bool {
		session := system.GetSession(res, req)
		_, found := session.GetOption(AuthorithID)
		return found
	}
}

// Authority 权限校验处理器
// 用于在路由过程中进行权限校验
func (intance impl) Authority() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {

		router := system.GetRouter()
		if !router.VerifyAuthority(res, req) {
			http.Redirect(res, req, "/", http.StatusFound)
			return
		}

		c.Next()
	}
}
