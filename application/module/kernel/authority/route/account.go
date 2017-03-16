package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/foundation/net"
)

// CreateAccountLoginRoute 创建AccountLogin Route
func CreateAccountLoginRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.AuthorityModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := authorityAccountLoginRoute{
			authorityHandler: endPoint.(common.AuthorityHandler),
			sessionRegistry:  sessionRegistry}
		return &i, true
	}

	return nil, false
}

// CreateAccountLogoutRoute 创建AccountLogout Route
func CreateAccountLogoutRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.AuthorityModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := authorityAccountLogoutRoute{
			authorityHandler: endPoint.(common.AuthorityHandler),
			sessionRegistry:  sessionRegistry}
		return &i, true
	}

	return nil, false
}

type authorityAccountLoginRoute struct {
	authorityHandler common.AuthorityHandler
	sessionRegistry  common.SessionRegistry
}

type authorityLoginResult struct {
	common.Result
	AuthToken string
}

func (i *authorityAccountLoginRoute) Type() string {
	return common.POST
}

func (i *authorityAccountLoginRoute) Pattern() string {
	return "/account/"
}

func (i *authorityAccountLoginRoute) Handler() interface{} {
	return i.loginHandler
}

func (i *authorityAccountLoginRoute) loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("loginHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := authorityLoginResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = 1
			result.Reason = "非法请求"
			break
		}

		account := r.FormValue("login-account")
		password := r.FormValue("login-password")
		user, token, ok := i.authorityHandler.LoginAccount(account, password)
		if !ok {
			result.ErrCode = 1
			result.Reason = "登入失败"
			break
		}

		session.SetAccount(user)

		result.ErrCode = 0
		result.AuthToken = token
		break
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type authorityAccountLogoutRoute struct {
	authorityHandler common.AuthorityHandler
	sessionRegistry  common.SessionRegistry
}

type authorityLogoutResult struct {
	common.Result
}

func (i *authorityAccountLogoutRoute) Type() string {
	return common.DELETE
}

func (i *authorityAccountLogoutRoute) Pattern() string {
	return "/account/?token=[a-z0-9A-Z]*/"
}

func (i *authorityAccountLogoutRoute) Handler() interface{} {
	return i.logoutHandler
}

func (i *authorityAccountLogoutRoute) logoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutHandler")

	result := authorityLogoutResult{}
	param := net.SplitParam(r.URL.Path)
	for true {
		token, ok := param["token"]
		if !ok {
			result.ErrCode = 1
			result.Reason = "非法请求"
			break
		}

		ok = i.authorityHandler.LogoutAccount(token)
		if !ok {
			result.ErrCode = 1
			result.Reason = "非法请求"
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
