package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/cas/def"
	"muidea.com/magicCenter/foundation/net"
)

// AppendAccountRoute 追加account 路由
func AppendAccountRoute(routes []common.Route, casHandler common.CASHandler, sessionRegistry common.SessionRegistry) []common.Route {
	rt, _ := CreateAccountLoginRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateAccountLogoutRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateAccountLoginRoute 创建AccountLogin Route
func CreateAccountLoginRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := authorityAccountLoginRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}

// CreateAccountLogoutRoute 创建AccountLogout Route
func CreateAccountLogoutRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := authorityAccountLogoutRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}

type authorityAccountLoginRoute struct {
	casHandler      common.CASHandler
	sessionRegistry common.SessionRegistry
}

type authorityLoginResult struct {
	common.Result
	User      string
	AuthToken string
}

func (i *authorityAccountLoginRoute) Method() string {
	return common.POST
}

func (i *authorityAccountLoginRoute) Pattern() string {
	return net.JoinURL(def.URL, "/user/")
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
		user, token, ok := i.casHandler.LoginAccount(account, password)
		if !ok {
			result.ErrCode = 1
			result.Reason = "登入失败"
			break
		}

		usr, found := session.GetAccount()
		if found && usr.Account == account {
			opt, ok := session.GetOption(common.AuthTokenID)
			if ok {
				token = opt.(string)
				result.ErrCode = 0
				result.Reason = "重复登陆"
				result.AuthToken = token
				break
			}
		}

		session.SetAccount(user)
		session.SetOption(common.AuthTokenID, token)

		result.ErrCode = 0
		result.User = user.Name
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
	casHandler      common.CASHandler
	sessionRegistry common.SessionRegistry
}

type authorityLogoutResult struct {
	common.Result
}

func (i *authorityAccountLogoutRoute) Method() string {
	return common.DELETE
}

func (i *authorityAccountLogoutRoute) Pattern() string {
	return net.JoinURL(def.URL, "/user/")
}

func (i *authorityAccountLogoutRoute) Handler() interface{} {
	return i.logoutHandler
}

func (i *authorityAccountLogoutRoute) logoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := authorityLogoutResult{}
	for true {
		token, ok := r.URL.Query()[common.AuthTokenID]
		if !ok || len(token) < 1 {
			result.ErrCode = 1
			result.Reason = "非法请求"
			break
		}

		authToken, ok := session.GetOption(common.AuthTokenID)
		if !ok || authToken != token[0] {
			result.ErrCode = 1
			result.Reason = "非法请求"
			break
		}

		i.casHandler.LogoutAccount(token[0])
		session.ClearAccount()
		session.RemoveOption(common.AuthTokenID)

		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
