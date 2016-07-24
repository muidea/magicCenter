package ui

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/kernel/dashboard/module/bll"
	"magiccenter/kernel/dashboard/module/model"
	"net/http"
	"strconv"
)

type ModulePageView struct {
	Modules []model.Module
}

func ModulePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ModulePageHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/module/page.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ModulePageView{}
	view.Modules = bll.QueryAllModules()

	t.Execute(w, view)
}

func SavePageBlockHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("SavePageBlockHandler")

	result := SavePageBlockResult{}
	for true {
		err := r.ParseMultipartForm(0)
		if err != nil {
			log.Print("paseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		owner := r.FormValue("page-owner")
		url := r.FormValue("page-url")

		blockList := []int{}
		blocks := r.MultipartForm.Value["page-block"]
		for _, b := range blocks {
			id, err := strconv.Atoi(b)
			if err != nil {
				log.Print("parse page block failed, b:%s", b)
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
			}

			blockList = append(blockList, id)
		}

		log.Print(owner)
		log.Print(url)
		log.Print(blockList)

		ret := false
		result.Module, ret = bll.SavePageBlock(owner, url, blockList)
		if !ret {
			result.ErrCode = 1
			result.Reason = "操作失败"
			break
		}

		result.ErrCode = 0
		result.Reason = "操作成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("Marshal failed, err:" + err.Error())
	}

	w.Write(b)
}
