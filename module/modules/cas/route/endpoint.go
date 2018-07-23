package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/cas/def"
	common_const "muidea.com/magicCommon/common"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/foundation/net"
)

type endpointVerifyRoute struct {
	casHandler      common.CASHandler
	sessionRegistry common.SessionRegistry
}

func (i *endpointVerifyRoute) Method() string {
	return common.GET
}

func (i *endpointVerifyRoute) Pattern() string {
	return net.JoinURL(def.URL, def.VerifyEndpoint)
}

func (i *endpointVerifyRoute) Handler() interface{} {
	return i.verifyHandler
}

func (i *endpointVerifyRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *endpointVerifyRoute) verifyHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("verifyHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.VerifyEndpointResult{}
	for {
		authToken := r.URL.Query().Get(common_const.AuthToken)
		identifyID := r.URL.Query().Get(common_const.IdentifyID)
		userInfo, ok := i.casHandler.LoginEndpoint(identifyID, authToken, r.RemoteAddr)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效授权"
			break
		}

		session.SetAccount(userInfo.User)
		session.SetOption(common_const.AuthToken, authToken)
		session.SetOption(common_const.ExpiryDate, -1)
		session.Flush()

		result.ErrorCode = common_def.Success
		result.SessionID = session.ID()
		break
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
