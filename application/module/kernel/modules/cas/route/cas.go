package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/cas/def"
	"muidea.com/magicCenter/foundation/net"
)

// AppendAccountRoute 追加account 路由
func AppendAccountRoute(routes []common.Route, casHandler common.CASHandler, sessionRegistry common.SessionRegistry) []common.Route {
	rt, _ := CreateAccountLoginRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateAccountLogoutRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateAccountStatusRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateAccountLoginRoute 创建AccountLogin Route
func CreateAccountLoginRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := accountLoginRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}

// CreateAccountLogoutRoute 创建AccountLogout Route
func CreateAccountLogoutRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := accountLogoutRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}

// CreateAccountStatusRoute 创建AccountStatus Route
func CreateAccountStatusRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := accountStatusRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}

type accountLoginRoute struct {
	casHandler      common.CASHandler
	sessionRegistry common.SessionRegistry
}

type accountLoginParam struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type accountLoginResult struct {
	common.Result
	OnlineUser model.AccountOnlineView `json:"onlineUser"`
	SessionID  string                  `json:"sessionID"`
}

func (i *accountLoginRoute) Method() string {
	return common.POST
}

func (i *accountLoginRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostUserLogin)
}

func (i *accountLoginRoute) Handler() interface{} {
	return i.loginHandler
}

func (i *accountLoginRoute) AuthGroup() int {
	return common.VisitorAuthGroup.ID
}

func (i *accountLoginRoute) loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("loginHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := accountLoginResult{}
	for true {
		param := &accountLoginParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			log.Printf("ParsePostJSON failed, err:%s", err.Error())
			result.ErrorCode = common.Failed
			result.Reason = "非法请求"
			break
		}

		remoteAddr := r.RemoteAddr
		onlineUser, ok := i.casHandler.LoginAccount(param.Account, param.Password, remoteAddr)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "登入失败"
			break
		}

		result.ErrorCode = common.Success
		result.OnlineUser = onlineUser
		result.SessionID = session.ID()
		break
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type accountLogoutRoute struct {
	casHandler      common.CASHandler
	sessionRegistry common.SessionRegistry
}

type accountLogoutResult struct {
	common.Result
}

func (i *accountLogoutRoute) Method() string {
	return common.DELETE
}

func (i *accountLogoutRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteUserLogout)
}

func (i *accountLogoutRoute) Handler() interface{} {
	return i.logoutHandler
}

func (i *accountLogoutRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
}

func (i *accountLogoutRoute) logoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := accountLogoutResult{}
	for true {
		authToken, ok := session.GetOption(common.AuthTokenID)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "非法请求"
			break
		}

		if !i.casHandler.Logout(authToken.(string), r.RemoteAddr) {
			result.ErrorCode = common.Failed
			result.Reason = "非法请求"
			break
		}

		result.ErrorCode = common.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type accountStatusRoute struct {
	casHandler      common.CASHandler
	sessionRegistry common.SessionRegistry
}

type accountStatusResult struct {
	common.Result
	OnlineUser model.AccountOnlineView `json:"onlineUser"`
	SessionID  string                  `json:"sessionID"`
}

func (i *accountStatusRoute) Method() string {
	return common.GET
}

func (i *accountStatusRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetUserStatus)
}

func (i *accountStatusRoute) Handler() interface{} {
	return i.statusHandler
}

func (i *accountStatusRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
}

func (i *accountStatusRoute) statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("statusHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := accountStatusResult{}
	for true {
		authToken, ok := session.GetOption(common.AuthTokenID)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "非法请求"
			break
		}

		onlineUser, found := i.casHandler.VerifyToken(authToken.(string))
		if !found {
			result.ErrorCode = common.Failed
			result.Reason = "无效Token"
			break
		}

		result.OnlineUser = onlineUser
		result.SessionID = session.ID()
		result.ErrorCode = common.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
