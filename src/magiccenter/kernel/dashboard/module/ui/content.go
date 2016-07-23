package ui

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/kernel/common"
	contentBll "magiccenter/kernel/content/bll"
	contentModel "magiccenter/kernel/content/model"
	"magiccenter/kernel/module/bll"
	"magiccenter/kernel/module/model"
	"net/http"
	"strconv"
	"strings"

	"muidea.com/util"
)

type ModuleContentView struct {
	Modules []model.Module
}

type QueryModuleContentResult struct {
	common.Result
	Module   model.ModuleContent
	Articles []contentModel.ArticleSummary
	Catalogs []contentModel.Catalog
	Links    []contentModel.Link
}

type SaveBlockContentResult struct {
	common.Result
	Module model.ModuleContent
}

func ModuleContentHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ModuleContentHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/module/content.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ModuleContentView{}
	view.Modules = bll.QueryAllModules()

	t.Execute(w, view)
}

func QueryModuleContentHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryModuleContentHandler")

	result := QueryModuleContentResult{}

	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		result.Module, found = bll.QueryModuleContent(id)
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		result.Articles = contentBll.QueryAllArticleSummary()
		result.Catalogs = contentBll.QueryAllCatalog()
		result.Links = contentBll.QueryAllLink()

		result.ErrCode = 0
		result.Reason = "查询成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("marshal failed, err:" + err.Error())
	}

	w.Write(b)
}

func SaveBlockContentHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("SaveBlockContentHandler")

	result := SaveBlockContentResult{}

	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		module_id := r.FormValue("module-id")
		block_id, err := strconv.Atoi(r.FormValue("block-id"))
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		articleSlice := strings.Split(r.FormValue("article-list"), ",")
		catalogSlice := strings.Split(r.FormValue("catalog-list"), ",")
		linkSlice := strings.Split(r.FormValue("link-list"), ",")

		articleList := []int{}
		for _, ar := range articleSlice {
			if len(ar) == 0 {
				continue
			}

			val, err := strconv.Atoi(ar)
			if err != nil {
				log.Printf("illegal article, id:=%s", ar)
				continue
			}

			articleList = append(articleList, val)
		}

		catalogList := []int{}
		for _, ca := range catalogSlice {
			if len(ca) == 0 {
				continue
			}

			val, err := strconv.Atoi(ca)
			if err != nil {
				log.Printf("illegal catalog, id:=%s", ca)
				continue
			}

			catalogList = append(catalogList, val)
		}

		linkList := []int{}
		for _, lnk := range linkSlice {
			if len(lnk) == 0 {
				continue
			}

			val, err := strconv.Atoi(lnk)
			if err != nil {
				log.Printf("illegal link, id:=%s", lnk)
				continue
			}

			linkList = append(linkList, val)
		}

		bll.SaveBlockItem(block_id, articleList, catalogList, linkList)

		module, found := bll.QueryModuleContent(module_id)
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		result.Module = module

		result.ErrCode = 0
		result.Reason = "操作成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("marshal failed, err:" + err.Error())
	}

	w.Write(b)
}
