package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCenter/module/modules/cas/def"
	common_const "github.com/muidea/magicCommon/common"
	common_def "github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/foundation/net"
	"github.com/muidea/magicCommon/model"
)

type accountLoginRoute struct {
	casHandler      common.CASHandler
	sessionRegistry common.SessionRegistry
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
	return common_const.VisitorAuthGroup.ID
}

func (i *accountLoginRoute) loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("loginHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.LoginAccountResult{}
	for true {
		param := &common_def.LoginAccountParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			log.Printf("ParsePostJSON failed, err:%s", err.Error())
			result.ErrorCode = common_def.Failed
			result.Reason = "非法请求"
			break
		}

		remoteAddr := r.RemoteAddr
		onlineEntry, authToken, ok := i.casHandler.LoginAccount(param.Account, param.Password, remoteAddr)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "登入失败"
			break
		}

		session.SetAccount(model.User{ID: onlineEntry.Unit.ID, Name: onlineEntry.Unit.Name})
		session.SetOption(common_const.AuthToken, authToken)
		session.Flush()

		result.OnlineEntry = onlineEntry
		result.SessionID = session.ID()
		result.AuthToken = authToken
		result.ErrorCode = common_def.Success
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
	return common_const.UserAuthGroup.ID
}

func (i *accountLogoutRoute) logoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.LogoutAccountResult{}
	for true {
		authToken, ok := session.GetOption(common_const.AuthToken)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "非法请求"
			break
		}

		if !i.casHandler.Logout(authToken.(string), r.RemoteAddr) {
			result.ErrorCode = common_def.Failed
			result.Reason = "非法请求"
			break
		}
		session.ClearAccount()
		session.RemoveOption(common_const.AuthToken)
		session.Flush()

		result.ErrorCode = common_def.Success
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
	return common_const.UserAuthGroup.ID
}

func (i *accountStatusRoute) statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("statusHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.StatusAccountResult{}
	for true {
		authToken, ok := session.GetOption(common_const.AuthToken)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "非法请求"
			break
		}

		onlineEntry, token, found := i.casHandler.VerifyToken(authToken.(string))
		if !found {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效Token"
			break
		}

		result.OnlineEntry = onlineEntry
		result.AuthToken = token
		result.SessionID = session.ID()
		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type accountChangePasswordRoute struct {
	casHandler      common.CASHandler
	sessionRegistry common.SessionRegistry
}

func (i *accountChangePasswordRoute) Method() string {
	return common.PUT
}

func (i *accountChangePasswordRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutUserChangePassword)
}

func (i *accountChangePasswordRoute) Handler() interface{} {
	return i.changePasswordHandler
}

func (i *accountChangePasswordRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *accountChangePasswordRoute) changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("changePasswordHandler")

	result := common_def.ChangeAccountPasswordResult{}
	for true {
		_, value := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		param := &common_def.ChangeAccountPasswordParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			log.Printf("ParsePostJSON failed, err:%s", err.Error())
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		ok := i.casHandler.ChangeAccountPassword(id, param.OldPassword, param.NewPassword)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "更改密码失败"
			break
		}

		result.ErrorCode = common_def.Success
		break
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
