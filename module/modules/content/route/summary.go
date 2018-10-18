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

	rt = CreateQuerySummaryContentRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateQuerySummaryContentByUserRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	return routes
}

// CreateQuerySummaryRoute 查询指定名称的Summary
func CreateQuerySummaryRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := summaryQueryRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateQuerySummaryContentRoute 查询指定分类的Summary
func CreateQuerySummaryContentRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := summaryContentQueryRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateQuerySummaryContentByUserRoute 查询指定分类的Summary
func CreateQuerySummaryContentByUserRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := summaryContentQueryByUserRoute{contentHandler: contentHandler, accountHandler: accountHandler}
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
	result := common_def.QuerySummaryResult{Summary: model.SummaryView{}}
	for true {
		strictCatalog, err := common_def.DecodeStrictCatalog(r)
		if err != nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal strictCatalog")
			break
		}
		if strictCatalog == nil {
			strictCatalog = common_const.SystemContentCatalog.CatalogUnit()
		}

		summaryName := r.URL.Query().Get("name")
		summaryType := r.URL.Query().Get("type")
		if len(summaryName) == 0 || len(summaryType) == 0 {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal contentType param, summaryName:%s, summaryType:%s", summaryName, summaryType)
			break
		}

		summary, ok := i.contentHandler.QuerySummaryByName(summaryName, summaryType, *strictCatalog)
		if ok {
			result.Summary.Summary = summary
			result.Summary.Catalog = i.contentHandler.GetSummaryByIDs(summary.Catalog)

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

type summaryContentQueryRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

func (i *summaryContentQueryRoute) Method() string {
	return common.GET
}

func (i *summaryContentQueryRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetSummaryDetail)
}

func (i *summaryContentQueryRoute) Handler() interface{} {
	return i.getSummaryDetailHandler
}

func (i *summaryContentQueryRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *summaryContentQueryRoute) getSummaryDetailHandler(w http.ResponseWriter, r *http.Request) {
	result := common_def.QuerySummaryListResult{Summary: []model.SummaryView{}}
	for true {
		filter := &common_def.Filter{}
		filter.Decode(r)

		_, str := net.SplitRESTAPI(r.URL.Path)
		summaryID, err := strconv.Atoi(str)
		if err != nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal id param, id:%s", str)
			break
		}

		summaryType := r.URL.Query().Get("type")
		if len(summaryType) == 0 {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal contentType param, summaryType is null")
			break
		}

		strictCatalog, err := common_def.DecodeStrictCatalog(r)
		if err != nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal strictCatalog")
			break
		}

		summary := model.CatalogUnit{ID: summaryID, Type: summaryType}
		summarys, total := i.contentHandler.QuerySummaryContent(summary, filter)
		for _, v := range summarys {
			if strictCatalog != nil {
				found := false
				for _, sv := range v.Catalog {
					if sv.IsSame(strictCatalog) {
						found = true
						break
					}
				}
				if !found {
					continue
				}
			}

			view := model.SummaryView{}
			view.Summary = v
			view.Catalog = i.contentHandler.GetSummaryByIDs(v.Catalog)

			user, ok := i.accountHandler.FindUserByID(v.Creater)
			if ok {
				view.Creater = user.User
			} else {
				view.Creater = model.User{ID: -1, Name: "未知用户"}
			}

			result.Summary = append(result.Summary, view)
		}
		result.Total = total

		result.ErrorCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type summaryContentQueryByUserRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

func (i *summaryContentQueryByUserRoute) Method() string {
	return common.GET
}

func (i *summaryContentQueryByUserRoute) Pattern() string {
	return net.JoinURL(def.URL, def.QuerySummaryDetail)
}

func (i *summaryContentQueryByUserRoute) Handler() interface{} {
	return i.querySummaryDetailHandler
}

func (i *summaryContentQueryByUserRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *summaryContentQueryByUserRoute) querySummaryDetailHandler(w http.ResponseWriter, r *http.Request) {
	result := common_def.QuerySummaryListResult{Summary: []model.SummaryView{}}
	for true {
		filter := &common_def.Filter{}
		filter.Decode(r)

		strictCatalog, err := common_def.DecodeStrictCatalog(r)
		if err != nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			log.Printf("illegal strictCatalog")
			break
		}

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

		summarys, total := i.contentHandler.GetSummaryByUser(uids, filter)
		for _, v := range summarys {
			if strictCatalog != nil {
				existFlag := false
				for _, sv := range v.Catalog {
					if sv.ID == strictCatalog.ID && sv.Type == strictCatalog.Type {
						existFlag = true
						break
					}
				}
				if !existFlag {
					continue
				}
			}

			view := model.SummaryView{}
			view.Summary = v
			view.Catalog = i.contentHandler.GetSummaryByIDs(v.Catalog)

			user, ok := i.accountHandler.FindUserByID(v.Creater)
			if ok {
				view.Creater = user.User
			} else {
				view.Creater = model.User{ID: -1, Name: "未知用户"}
			}

			result.Summary = append(result.Summary, view)
		}
		result.Total = total

		result.ErrorCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
