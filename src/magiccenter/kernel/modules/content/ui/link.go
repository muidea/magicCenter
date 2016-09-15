package ui

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/common"
	"magiccenter/common/model"
	"magiccenter/configuration"
	"magiccenter/kernel/modules/content/bll"
	"magiccenter/session"
	"net/http"
	"strconv"

	"muidea.com/util"
)

type ManageLinkView struct {
	Links    []model.Link
	Catalogs []model.CatalogDetail
}

type QueryAllLinkResult struct {
	Links []model.Link
}

type QueryLinkResult struct {
	common.Result
	Link model.Link
}

type DeleteLinkResult struct {
	common.Result
}

type AjaxLinkResult struct {
	common.Result
}

type EditLinkResult struct {
	common.Result
	Link model.Link
}

//
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
	view.Catalogs = bll.QueryAllCatalogDetail()

	t.Execute(w, view)
}

//
// 查询全部Link
// 返回json
//
func QueryAllLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryAllLinkHandler")

	result := QueryAllLinkResult{}
	result.Links = bll.QueryAllLink()

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

//
// 查询指定Link内容
// 返回json
//
func QueryLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryLinkHandler")

	result := QueryLinkResult{}

	for true {
		err := r.ParseForm()
		if err != nil {
			log.Print("paseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

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
			result.Reason = "操作失败"
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

//
// 删除指定Link
// 返回json
//
func DeleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteLinkHandler")

	result := DeleteLinkResult{}

	for true {
		err := r.ParseForm()
		if err != nil {
			log.Print("paseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

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
			result.Reason = "操作失败"
			break
		}

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

//
// 保存Link
// 返回json
//
func AjaxLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxCatalogHandler")

	authId, found := configuration.GetOption(configuration.AuthorithID)
	if !found {
		panic("unexpected, can't fetch authorith id")
	}

	session := session.GetSession(w, r)
	user, found := session.GetOption(authId)
	if !found {
		panic("unexpected, must login system first.")
	}

	result := AjaxLinkResult{}

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
			result.Reason = "操作失败"
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

//
// 编辑Link
// 返回Link内容和当前可用Catalog
//
func EditLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("EditLinkHandler")

	result := EditLinkResult{}

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
			result.Reason = "操作失败"
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
