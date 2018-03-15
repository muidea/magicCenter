package blog

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/common"
	"net/http"
)

const (
	blogTitleID       = "@blog_title"
	blogDescriptionID = "@blog_description"
)

// PageView 页面视图
type PageView struct {
	View model.PageView
}

// MaintainView Maintain视图
type MaintainView struct {
	Title       string `json:"titile"`
	Description string `json:"description"`
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("indexHandler")

	res.Header().Set("content-type", "text/html")
	res.Header().Set("charset", "utf-8")

	view := PageView{}
	/*
		url := req.URL.Path
		view.View, _ = bll.QueryPageView(ID, url)
	*/
	t, err := template.ParseFiles("template/html/modules/blog/index.html")
	if err != nil {
		panic("ParseFiles failed, err:" + err.Error())
	}

	t.Execute(res, view)
}

func viewContentHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("viewContentHandler")

	res.Header().Set("content-type", "text/html")
	res.Header().Set("charset", "utf-8")

	view := PageView{}
	/*
		params := util.SplitParam(req.URL.RawQuery)
		str, found := params["id"]
		if !found {
			panic("illegl param")
		}

		id, err := strconv.Atoi(str)
		if err != nil {
			panic("illegl id, err:" + err.Error())
		}

		view := PageContentView{}

		url := req.URL.Path
		view.View, _ = bll.QueryContentView(ID, url, id)
	*/
	t, err := template.ParseFiles("template/html/modules/blog/view.html")
	if err != nil {
		panic("ParseFiles failed, err:" + err.Error())
	}

	t.Execute(res, view)
}

func viewCatalogHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("viewCatalogHandler")

	res.Header().Set("content-type", "text/html")
	res.Header().Set("charset", "utf-8")

	view := PageView{}
	/*
		params := util.SplitParam(req.URL.RawQuery)
		str, found := params["id"]
		if !found {
			panic("illegl param")
		}

		id, err := strconv.Atoi(str)
		if err != nil {
			panic("illegl id, err:" + err.Error())
		}

		view := PageCatalogView{}

		url := req.URL.Path
		view.View, _ = bll.QueryCatalogView(ID, url, id)
	*/
	t, err := template.ParseFiles("template/html/modules/blog/catalog.html")
	if err != nil {
		panic("ParseFiles failed, err:" + err.Error())
	}

	t.Execute(res, view)
}

func viewLinkHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("viewLinkHandler")
	/*
		params := util.SplitParam(req.URL.RawQuery)
		str, found := params["id"]
		if !found {
			panic("illegl param")
		}

		id, err := strconv.Atoi(str)
		if err != nil {
			panic("illegl id, err:" + err.Error())
		}

		link, found := contentbll.QueryLinkById(id)
		if !found {
			http.Redirect(res, req, "/", http.StatusNotFound)
			return
		}

		http.Redirect(res, req, link.Url, http.StatusFound)
	*/
}

// MaintainViewHandler 管理维护视图处理器
func MaintainViewHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("MaintainViewHandler")

	res.Header().Set("content-type", "text/html")
	res.Header().Set("charset", "utf-8")

	view := MaintainView{}

	t, err := template.ParseFiles("template/html/modules/blog/maintain.html")
	if err != nil {
		panic("ParseFiles failed, err:" + err.Error())
	}

	t.Execute(res, view)
}

// MaintainActionHandler 管理维护信息处理器
func MaintainActionHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("MaintainActionHandler")

	result := common.Result{}

	for {
		err := req.ParseForm()
		if err != nil {
			log.Print("parseform failed")

			result.ErrorCode = common.Failed
			result.Reason = "无效请求数据"
			break
		}

		result.ErrorCode = common.Success
		result.Reason = "更新成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	res.Write(b)
}
