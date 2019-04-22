package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCenter/module/modules/cas/def"
	common_const "github.com/muidea/magicCommon/common"
	common_def "github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/foundation/net"
	"github.com/muidea/magicCommon/model"
)

type endpointLoginRoute struct {
	casHandler      common.CASHandler
	sessionRegistry common.SessionRegistry
}

func (i *endpointLoginRoute) Method() string {
	return common.POST
}

func (i *endpointLoginRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostEndpointLogin)
}

func (i *endpointLoginRoute) Handler() interface{} {
	return i.loginHandler
}

func (i *endpointLoginRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *endpointLoginRoute) loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("loginHandler")

	session := i.sessionRegistry.GetSession(w, r)
	param := &common_def.LoginEndpointParam{}
	result := common_def.LoginEndpointResult{}
	for {
		err := net.ParseJSONBody(r, param)
		if err != nil {
			log.Printf("ParseJSONBody failed, err:%s", err.Error())
			result.ErrorCode = common_def.Failed
			result.Reason = "非法请求"
			break
		}

		userInfo, authToken, ok := i.casHandler.LoginEndpoint(param.IdentifyID, param.AuthToken, r.RemoteAddr)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效授权"
			break
		}

		session.SetAccount(model.User{ID: userInfo.Unit.ID, Name: userInfo.Unit.Name})
		session.SetOption(common_const.AuthToken, authToken)
		session.SetOption(common_const.ExpiryDate, -1)
		session.Flush()

		result.AuthToken = authToken
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

type endpointLogoutRoute struct {
	casHandler      common.CASHandler
	sessionRegistry common.SessionRegistry
}

func (i *endpointLogoutRoute) Method() string {
	return common.DELETE
}

func (i *endpointLogoutRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteEndpointLogout)
}

func (i *endpointLogoutRoute) Handler() interface{} {
	return i.logoutHandler
}

func (i *endpointLogoutRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *endpointLogoutRoute) logoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.LogoutEndpointResult{}
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

type endpointStatusRoute struct {
	casHandler      common.CASHandler
	sessionRegistry common.SessionRegistry
}

func (i *endpointStatusRoute) Method() string {
	return common.GET
}

func (i *endpointStatusRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetEndpointStatus)
}

func (i *endpointStatusRoute) Handler() interface{} {
	return i.statusHandler
}

func (i *endpointStatusRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *endpointStatusRoute) statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("statusHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.StatusEndpointResult{}
	for true {
		authToken, ok := session.GetOption(common_const.AuthToken)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "非法请求"
			break
		}

		_, token, found := i.casHandler.VerifyToken(authToken.(string))
		if !found {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效Token"
			break
		}

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
