package ui

import (
	"encoding/json"
	"html"
	"html/template"
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	"magiccenter/common/model"
	"magiccenter/kernel/modules/content/bll"
	"magiccenter/system"
	"net/http"
	"path"
	"strconv"
	"time"

	"muidea.com/util"
)

// ManageMediaView Media管理视图
type ManageMediaView struct {
	Medias   []model.MediaDetail
	Catalogs []model.Catalog
	Users    []model.User
}

// AllMediaList 全部Media列表
type AllMediaList struct {
	Medias []model.MediaDetail
}

// SingleyMedia 单个Media文件
type SingleyMedia struct {
	common.Result
	Media model.MediaDetail
}

// ManageMediaViewHandler Media管理主界面处理器
// 显示Media列表信息
// 返回html页面
//
func ManageMediaViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageMediaViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/content/media.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ManageMediaView{}
	view.Medias = bll.QueryAllMedia()
	view.Catalogs = bll.QueryAllCatalogList()
	view.Users = commonbll.QueryAllUserList()

	t.Execute(w, view)
}

// QueryAllMediaHandler 查询全部Media
// 返回json
func QueryAllMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryAllMediaHandler")

	result := AllMediaList{}
	result.Medias = bll.QueryAllMedia()

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// QueryMediaHandler 查询指定Media内容
// 返回json
func QueryMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryMediaHandler")

	result := SingleyMedia{}

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

		media, found := bll.QueryMediaByID(aid)
		if !found {
			result.ErrCode = 1
			result.Reason = "查询失败"
			break
		}

		media.Desc = html.UnescapeString(media.Desc)
		result.Media = media
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

// DeleteMediaHandler 删除指定Media
// 返回json
func DeleteMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteMediaHandler")

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

		if !bll.DeleteMediaByID(aid) {
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

// AjaxMediaHandler 保存Media
// 返回json
func AjaxMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxCatalogHandler")

	session := system.GetSession(w, r)
	user, found := session.GetAccount()
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

		id, err := strconv.Atoi(r.FormValue("media-id"))
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		name := r.FormValue("media-name")

		staticPath := system.GetStaticPath()
		filePath := path.Join(staticPath, time.Now().Format("20060102150405"))

		url, fileType, err := util.MultipartFormFile(r, "media-url", filePath)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		desc := html.EscapeString(r.FormValue("media-desc"))
		catalog := r.MultipartForm.Value["media-catalog"]
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

		if !bll.SaveMedia(id, name, url, fileType, desc, user.ID, catalogs) {
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
