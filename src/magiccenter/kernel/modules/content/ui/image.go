package ui

import (
	"encoding/json"
	"html"
	"html/template"
	"log"
	"magiccenter/common"
	"magiccenter/common/model"
	"magiccenter/configuration"
	"magiccenter/kernel/modules/content/bll"
	"magiccenter/session"
	"net/http"
	"path"
	"strconv"
	"time"

	"muidea.com/util"
)

// ManageImageView Image管理视图
type ManageImageView struct {
	Images   []model.ImageDetail
	Catalogs []model.CatalogDetail
}

type QueryAllImageResult struct {
	Images []model.ImageDetail
}

type QueryImageResult struct {
	common.Result
	Image model.ImageDetail
}

type DeleteImageResult struct {
	common.Result
}

type AjaxImageResult struct {
	common.Result
}

type EditImageResult struct {
	common.Result
	Image model.ImageDetail
}

// ManageImageViewHandler Image管理主界面处理器
// 显示Image列表信息
// 返回html页面
//
func ManageImageViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageImageViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/content/image.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ManageImageView{}
	view.Images = bll.QueryAllImage()
	view.Catalogs = bll.QueryAllCatalogDetail()

	t.Execute(w, view)
}

//
// 查询全部Image
// 返回json
//
func QueryAllImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryAllImageHandler")

	result := QueryAllImageResult{}
	result.Images = bll.QueryAllImage()

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

//
// 查询指定Image内容
// 返回json
//
func QueryImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryImageHandler")

	result := QueryImageResult{}

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

		image, found := bll.QueryImageByID(aid)
		if !found {
			result.ErrCode = 1
			result.Reason = "操作失败"
			break
		}

		image.Desc = html.UnescapeString(image.Desc)
		result.Image = image
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
// 删除指定Image
// 返回json
//
func DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteImageHandler")

	result := DeleteImageResult{}

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

		if !bll.DeleteImageByID(aid) {
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
// 保存Image
// 返回json
//
func AjaxImageHandler(w http.ResponseWriter, r *http.Request) {
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

	result := AjaxImageResult{}

	for true {
		err := r.ParseMultipartForm(0)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		id, err := strconv.Atoi(r.FormValue("image-id"))
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		name := r.FormValue("image-name")

		staticPath, _ := configuration.GetOption(configuration.StaticPath)
		uploadPath, _ := configuration.GetOption(configuration.UploadPath)
		filePath := path.Join(staticPath, uploadPath, time.Now().Format("20060102150405"))

		url, err := util.MultipartFormFile(r, "image-url", filePath)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		desc := html.EscapeString(r.FormValue("image-desc"))
		catalog := r.MultipartForm.Value["image-catalog"]
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

		if !bll.SaveImage(id, name, url, desc, user.(model.UserDetail).ID, catalogs) {
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
// 编辑Image
// 返回Image内容和当前可用Catalog
//
func EditImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("EditImageHandler")

	result := EditImageResult{}

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

		image, found := bll.QueryImageByID(aid)
		if !found {
			result.ErrCode = 1
			result.Reason = "操作失败"
			break
		}

		result.Image = image
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
