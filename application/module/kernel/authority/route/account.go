package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
)

// CreateAccountLoginRoute 创建AccountLogin Route
func CreateAccountLoginRoute(authorityHandler common.AuthorityHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := authorityAccountLoginRoute{
		authorityHandler: authorityHandler,
		sessionRegistry:  sessionRegistry}
	return &i, true
}

// CreateAccountLogoutRoute 创建AccountLogout Route
func CreateAccountLogoutRoute(authorityHandler common.AuthorityHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := authorityAccountLogoutRoute{
		authorityHandler: authorityHandler,
		sessionRegistry:  sessionRegistry}
	return &i, true
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
		_, found := session.GetAccount()
		if found {
			result.ErrCode = 1
			result.Reason = "重复登陆"
			break
		}

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
	return "/account/"
}

func (i *authorityAccountLogoutRoute) Handler() interface{} {
	return i.logoutHandler
}

func (i *authorityAccountLogoutRoute) logoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := authorityLogoutResult{}
	for true {
		token, ok := r.URL.Query()["token"]
		if !ok || len(token) < 1 {
			result.ErrCode = 1
			result.Reason = "非法请求"
			break
		}

		ok = i.authorityHandler.LogoutAccount(token[0])
		if !ok {
			result.ErrCode = 1
			result.Reason = "无效请求"
			break
		}

		session.ClearAccount()

		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
