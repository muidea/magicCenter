package ui

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	"magiccenter/common/model"
	"magiccenter/configuration"
	"magiccenter/kernel/modules/content/bll"
	"magiccenter/session"
	"net/http"
	"strconv"

	"muidea.com/util"
)

// ManageLinkView Link管理视图
type ManageLinkView struct {
	Links    []model.Link
	Catalogs []model.Catalog
	Users    []model.User
}

// AllLinkList 全部Link列表
type AllLinkList struct {
	Links []model.Link
}

// SingleLink 单条Link
type SingleLink struct {
	common.Result
	Link model.Link
}

// ManageLinkViewHandler 链接管理视图处理器
// Link管理主界面
// 显示Link列表信息
// 返回html页面
//
func ManageLinkViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageLinkViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/content/link.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ManageLinkView{}
	view.Links = bll.QueryAllLink()
	view.Catalogs = bll.QueryAllCatalogList()
	view.Users = commonbll.QueryAllUserList()

	t.Execute(w, view)
}

// QueryAllLinkHandler 查询全部Link
// 返回json
func QueryAllLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryAllLinkHandler")

	result := AllLinkList{}
	result.Links = bll.QueryAllLink()

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// QueryLinkHandler 查询指定Link内容
// 返回json
func QueryLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryLinkHandler")

	result := SingleLink{}

	for true {
		params := util.SplitParam(r.URL.RawQuery)
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		aid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		image, found := bll.QueryLinkByID(aid)
		if !found {
			result.ErrCode = 1
			result.Reason = "查询失败"
			break
		}

		result.Link = image
		result.ErrCode = 0
		result.Reason = "查询成功"

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// DeleteLinkHandler 删除指定Link
// 返回json
func DeleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteLinkHandler")

	result := common.Result{}

	for true {
		params := util.SplitParam(r.URL.RawQuery)
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		aid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		if !bll.DeleteLinkByID(aid) {
			result.ErrCode = 1
			result.Reason = "删除失败"
			break
		}

		result.ErrCode = 0
		result.Reason = "删除成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// AjaxLinkHandler 保存Link
// 返回json
//
func AjaxLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxCatalogHandler")

	authID, found := configuration.GetOption(configuration.AuthorithID)
	if !found {
		panic("unexpected, can't fetch authorith id")
	}

	session := session.GetSession(w, r)
	user, found := session.GetOption(authID)
	if !found {
		panic("unexpected, must login system first.")
	}

	result := common.Result{}
	for true {
		err := r.ParseMultipartForm(0)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		id, err := strconv.Atoi(r.FormValue("link-id"))
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		name := r.FormValue("link-name")
		url := r.FormValue("link-url")
		logo := r.FormValue("link-logo")
		catalog := r.MultipartForm.Value["link-catalog"]
		catalogs := []int{}
		for _, c := range catalog {
			cid, err := strconv.Atoi(c)
			if err != nil {
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
			}

			catalogs = append(catalogs, cid)
		}

		if !bll.SaveLink(id, name, url, logo, user.(model.UserDetail).ID, catalogs) {
			result.ErrCode = 1
			result.Reason = "保存失败"
			break
		}

		result.ErrCode = 0
		result.Reason = "保存成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
