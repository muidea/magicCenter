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
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

// AppendSummaryRoute 追加Summary Route
func AppendSummaryRoute(routes []common.Route, contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) []common.Route {
	rt := CreateQuerySummaryRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateGetSummaryRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	return routes
}

// CreateQuerySummaryRoute 查询指定名称的Summary
func CreateQuerySummaryRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := summaryQueryRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateGetSummaryRoute 查询指定分类的Summary
func CreateGetSummaryRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := summaryGetRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

type summaryQueryRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

type summaryQueryResult struct {
	common_result.Result
	Summary model.SummaryView `json:"summary"`
}

func (i *summaryQueryRoute) Method() string {
	return common.GET
}

func (i *summaryQueryRoute) Pattern() string {
	return net.JoinURL(def.URL, def.QuerySummary)
}

func (i *summaryQueryRoute) Handler() interface{} {
	return i.querySummaryHandler
}

func (i *summaryQueryRoute) AuthGroup() int {
	return common_def.UserAuthGroup.ID
}

func (i *summaryQueryRoute) querySummaryHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("querySummaryHandler")

	result := summaryQueryResult{Summary: model.SummaryView{}}
	for true {
		summaryName := r.URL.Query().Get("name")
		summaryType := r.URL.Query().Get("type")
		if len(summaryName) == 0 || len(summaryType) == 0 {
			result.ErrorCode = common_result.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal contentType param, summaryName:%s, summaryType:%s", summaryName, summaryType)
			break
		}

		summary, ok := i.contentHandler.QuerySummaryByName(summaryName, summaryType)
		if ok {
			result.Summary.Summary = summary
			result.Summary.Catalog = i.contentHandler.GetCatalogs(summary.Catalog)

			user, ok := i.accountHandler.FindUserByID(summary.Creater)
			if ok {
				result.Summary.Creater = user.User
			} else {
				result.Summary.Creater = model.User{ID: -1, Name: "未知用户"}
			}

			result.ErrorCode = 0
			break
		}

		result.ErrorCode = common_result.NoExist
		result.Reason = "对象不存在"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
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
	return net.JoinURL(def.URL, def.GetSummaryDetail)
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
		_, str := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(str)
		if err != nil {
			result.ErrorCode = common_result.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal id param, id:%s", str)
			break
		}
		contentType := r.URL.Query().Get("type")
		if len(contentType) == 0 {
			result.ErrorCode = common_result.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal contentType param, contentType:%s", contentType)
			break
		}

		uid := -1
		userStr := r.URL.Query().Get("user")
		if len(userStr) > 0 {
			uid, err = strconv.Atoi(userStr)
			if err != nil {
				result.ErrorCode = common_result.IllegalParam
				result.Reason = "非法参数"
				log.Printf("illegal user filter param, user:%s", userStr)
				break
			}
		}
		cid := -1
		catalog := r.URL.Query().Get("catalog")
		if len(catalog) > 0 {
			cid, err = strconv.Atoi(catalog)
			if err != nil {
				result.ErrorCode = common_result.IllegalParam
				result.Reason = "非法参数"
				log.Printf("illegal user filter param, catalog:%s", catalog)
				break
			}
		}

		summarys := i.contentHandler.GetSummaryContent(id, contentType)
		for _, v := range summarys {
			if len(userStr) > 0 {
				if v.Creater != uid {
					continue
				}
			}
			if len(catalog) > 0 {
				if !util.ExistIntArray(cid, v.Catalog) {
					continue
				}
			}
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
