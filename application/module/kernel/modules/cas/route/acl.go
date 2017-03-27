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

// CreateQueryACLRoute 新建QueryACL 路由
func CreateQueryACLRoute(authorityHandler common.CASHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityACLAddRoute{
		authorityHandler: authorityHandler}
	return &i
}

// CreateAddACLRoute 新建AddACL 路由
func CreateAddACLRoute(authorityHandler common.CASHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityACLAddRoute{
		authorityHandler: authorityHandler}
	return &i
}

type authorityACLQueryRoute struct {
	authorityHandler common.CASHandler
}

type authorityACLQueryResult struct {
	common.Result
	ACLs []model.ACL
}

func (i *authorityACLQueryRoute) Method() string {
	return common.GET
}

func (i *authorityACLQueryRoute) Pattern() string {
	return net.JoinURL(def.URL, "/acl/")
}

func (i *authorityACLQueryRoute) Handler() interface{} {
	return i.queryACLHandler
}

func (i *authorityACLQueryRoute) queryACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryACLHandler")

	result := authorityACLQueryResult{}
	for true {
		modules := r.URL.Query()["module"]
		if len(modules) < 1 {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		acls, ok := i.authorityHandler.QueryACL(modules[0])
		if !ok {
			result.ErrCode = 1
			result.Reason = "查询失败"
			break
		}

		result.ErrCode = 0
		result.ACLs = acls
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type authorityACLAddRoute struct {
	authorityHandler common.CASHandler
}

type authorityACLAddResult struct {
	common.Result
	ACL model.ACL
}

func (i *authorityACLAddRoute) Method() string {
	return common.GET
}

func (i *authorityACLAddRoute) Pattern() string {
	return net.JoinURL(def.URL, "/acl/")
}

func (i *authorityACLAddRoute) Handler() interface{} {
	return i.addACLHandler
}

func (i *authorityACLAddRoute) addACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("addACLHandler")

	result := authorityACLAddResult{}
	for true {
		r.ParseForm()

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
