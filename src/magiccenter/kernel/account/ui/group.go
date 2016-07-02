package ui

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/kernel/account/bll"
	"magiccenter/kernel/account/model"
	"magiccenter/kernel/common"
	"net/http"
	"strconv"

	"muidea.com/util"
)

type ManageGroupView struct {
	Groups []model.GroupInfo
}

type QueryAllGroupResult struct {
	common.Result
	Groups []model.GroupInfo
}

type QueryGroupResult struct {
	common.Result
	Group model.Group
}

type SaveGroupResult struct {
	common.Result
	Groups []model.GroupInfo
}

type DeleteGroupResult struct {
	SaveGroupResult
}

func ManageGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageGroupHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/account/group.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ManageGroupView{}
	view.Groups = bll.QueryAllGroupInfo()

	t.Execute(w, view)
}

func QueryAllGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAllGroupHandler")

	result := QueryAllGroupResult{}
	result.Groups = bll.QueryAllGroupInfo()
	result.ErrCode = 0
	result.Reason = "查询成功"

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

func QueryGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryGroupHandler")

	result := QueryGroupResult{}

	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		gid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		result.Group, found = bll.QueryGroupById(gid)
		if !found {
			result.ErrCode = 1
			result.Reason = "指定Group不存在"
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

func AjaxGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ajaxGroupHandler")

	result := SaveGroupResult{}
	for true {
		err := r.ParseMultipartForm(0)
		if err != nil {
			log.Print("paseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		id := r.FormValue("group-id")
		name := r.FormValue("group-name")

		gid, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("parse id failed, id:%s", id)
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		ok := bll.SaveGroup(gid, name)
		if !ok {
			result.ErrCode = 1
			result.Reason = "保存分组失败"
			break
		}

		result.Groups = bll.QueryAllGroupInfo()
		result.ErrCode = 0
		result.Reason = "保存分组成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
		return
	}

	w.Write(b)
}

func DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteGroupHandler")

	result := DeleteGroupResult{}
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		gid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		ok := bll.DeleteGroup(gid)
		if !ok {
			result.ErrCode = 1
			result.Reason = "删除分组失败"
			break
		}

		result.ErrCode = 0
		result.Reason = "查询成功"
		result.Groups = bll.QueryAllGroupInfo()
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
