package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/content/def"
	common_def "muidea.com/magicCommon/common"
	common_result "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/foundation/net"
	"muidea.com/magicCommon/model"
)

// AppendSummaryRoute 追加Summary Route
func AppendSummaryRoute(routes []common.Route, contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) []common.Route {
	rt := CreateGetSummaryByIDRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	return routes
}

// CreateGetSummaryByIDRoute 查询指定分类的Summary
func CreateGetSummaryByIDRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := summaryGetByIDRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

type summaryGetByIDRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

type summaryGetByIDResult struct {
	common_result.Result
	Summary []model.SummaryView `json:"summary"`
}

func (i *summaryGetByIDRoute) Method() string {
	return common.GET
}

func (i *summaryGetByIDRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetSummaryByCatalog)
}

func (i *summaryGetByIDRoute) Handler() interface{} {
	return i.getSummaryHandler
}

func (i *summaryGetByIDRoute) AuthGroup() int {
	return common_def.UserAuthGroup.ID
}

func (i *summaryGetByIDRoute) getSummaryHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getSummaryHandler")

	result := summaryGetByIDResult{Summary: []model.SummaryView{}}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		summarys := i.contentHandler.GetSummaryByCatalog(id)
		for _, v := range summarys {
			view := model.SummaryView{}
			view.Summary = v
			view.Catalog = i.contentHandler.GetCatalogs(v.Catalog)

			user, ok := i.accountHandler.FindUserByID(v.Creater)
			if ok {
				view.Creater = user.User
			} else {
				view.Creater = model.User{ID: -1, Name: "未知用户"}
			}

			result.Summary = append(result.Summary, view)
		}

		result.ErrorCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
