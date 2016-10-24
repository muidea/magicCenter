package ui

import (
	"encoding/json"
	"log"
	"magiccenter/common"
	"magiccenter/common/model"
	"magiccenter/kernel/modules/dashboard/content/bll"
	"net/http"
	"strconv"
	"strings"

	"muidea.com/util"
)

// SinglePage 当个Page
type SinglePage struct {
	Result common.Result
	Page   model.Page
}

// GetPageActionHandler 查询Page拥有的Block信息
func GetPageActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetPageActionHandler")

	result := SinglePage{}
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		owner, found := params["module"]
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		url, found := params["url"]
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		result.Page, found = bll.GetModulePage(owner, url)
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		result.Result.ErrCode = 0
		result.Result.Reason = "查询成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// PostPageActionHandler 新建Page
func PostPageActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PostPageActionHandler")

	result := common.Result{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		page := model.Page{}
		page.Owner = r.FormValue("module")
		page.URL = r.FormValue("url")
		blocks := r.FormValue("blocks")
		blockList := strings.Split(blocks, ",")
		for _, ii := range blockList {
			id, err := strconv.Atoi(ii)
			if err == nil {
				page.Blocks = append(page.Blocks, id)
			}
		}

		ret := bll.SaveModulePage(page)
		if !ret {
			result.ErrCode = 1
			result.Reason = "保存Page失败"
			break
		}

		result.ErrCode = 0
		result.Reason = "保存Page成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// DeletePageActionHandler 删除Page
func DeletePageActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeletePageActionHandler")

	result := common.Result{}
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		owner, found := params["module"]
		if !found {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		url, found := params["url"]
		if !found {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		ret := bll.DeleteModulePage(owner, url)
		if !ret {
			result.ErrCode = 1
			result.Reason = "无效参数"
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
