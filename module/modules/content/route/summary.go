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
	rt := CreateGetSummaryRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	return routes
}

// CreateGetSummaryRoute 查询指定分类的Summary
func CreateGetSummaryRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := summaryGetRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

type summaryGetRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

type summaryGetResult struct {
	common_result.Result
	Summary []model.SummaryView `json:"summary"`
}

func (i *summaryGetRoute) Method() string {
	return common.GET
}

func (i *summaryGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetSummary)
}

func (i *summaryGetRoute) Handler() interface{} {
	return i.getSummaryHandler
}

func (i *summaryGetRoute) AuthGroup() int {
	return common_def.UserAuthGroup.ID
}

func (i *summaryGetRoute) getSummaryHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getSummaryHandler")

	result := summaryGetResult{Summary: []model.SummaryView{}}
	for true {
		catalogStr := r.URL.Query().Get("catalog")
		if len(catalogStr) > 0 {
			id, err := strconv.Atoi(catalogStr)
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

		userStr, ok := r.URL.Query()["user[]"]
		if ok {
			uids := []int{}
			for _, val := range userStr {
				id, err := strconv.Atoi(val)
				if err == nil {
					uids = append(uids, id)
				}
			}
			if len(uids) != len(userStr) {
				result.ErrorCode = common_result.IllegalParam
				result.Reason = "无效参数"
				break
			}

			summarys := i.contentHandler.GetSummaryByUser(uids)
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

		result.ErrorCode = common_result.IllegalParam
		result.Reason = "无效参数"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
