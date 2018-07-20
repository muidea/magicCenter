package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/content/def"
	common_const "muidea.com/magicCommon/common"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/foundation/net"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

// AppendSummaryRoute 追加Summary Route
func AppendSummaryRoute(routes []common.Route, contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) []common.Route {
	rt := CreateQuerySummaryRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateGetSummaryDetailRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateQuerySummaryDetailRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	return routes
}

// CreateQuerySummaryRoute 查询指定名称的Summary
func CreateQuerySummaryRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := summaryQueryRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateGetSummaryDetailRoute 查询指定分类的Summary
func CreateGetSummaryDetailRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := summaryDetailGetRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateQuerySummaryDetailRoute 查询指定分类的Summary
func CreateQuerySummaryDetailRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := summaryDetailQueryRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

type summaryQueryRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
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
	return common_const.UserAuthGroup.ID
}

func (i *summaryQueryRoute) querySummaryHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("querySummaryHandler")

	result := common_def.QuerySummaryResult{Summary: model.SummaryView{}}
	for true {
		summaryName := r.URL.Query().Get("name")
		summaryType := r.URL.Query().Get("type")
		if len(summaryName) == 0 || len(summaryType) == 0 {
			result.ErrorCode = common_def.IllegalParam
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

		result.ErrorCode = common_def.NoExist
		result.Reason = "对象不存在"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type summaryDetailGetRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

func (i *summaryDetailGetRoute) Method() string {
	return common.GET
}

func (i *summaryDetailGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetSummaryDetail)
}

func (i *summaryDetailGetRoute) Handler() interface{} {
	return i.getSummaryDetailHandler
}

func (i *summaryDetailGetRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *summaryDetailGetRoute) getSummaryDetailHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getSummaryDetailHandler")

	result := common_def.QuerySummaryListResult{Summary: []model.SummaryView{}}
	for true {
		_, str := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(str)
		if err != nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal id param, id:%s", str)
			break
		}
		contentType := r.URL.Query().Get("type")
		if len(contentType) == 0 {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal contentType param, contentType:%s", contentType)
			break
		}

		uid := -1
		userStr := r.URL.Query().Get("user")
		if len(userStr) > 0 {
			uid, err = strconv.Atoi(userStr)
			if err != nil {
				result.ErrorCode = common_def.IllegalParam
				result.Reason = "非法参数"
				log.Printf("illegal user filter param, user:%s", userStr)
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

type summaryDetailQueryRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

func (i *summaryDetailQueryRoute) Method() string {
	return common.GET
}

func (i *summaryDetailQueryRoute) Pattern() string {
	return net.JoinURL(def.URL, def.QuerySummaryDetail)
}

func (i *summaryDetailQueryRoute) Handler() interface{} {
	return i.querySummaryDetailHandler
}

func (i *summaryDetailQueryRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *summaryDetailQueryRoute) querySummaryDetailHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("querySummaryDetailHandler")

	result := common_def.QuerySummaryListResult{Summary: []model.SummaryView{}}
	for true {
		userStr := r.URL.Query().Get("user[]")
		if len(userStr) == 0 {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal user filter param, user:%s", userStr)
			break
		}

		uids, ok := util.Str2IntArray(userStr)
		if !ok {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal user filter param, user:%s", userStr)
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

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
